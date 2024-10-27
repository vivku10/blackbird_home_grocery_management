package api

import (
	"context"
	"net/http"
	"os"
	"time"
  api "api/internal/api"
  "github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	ctrl "sigs.k8s.io/controller-runtime"
)


type apiService struct {
	logger zerolog.Logger
	addr string
	// TODO: optional k8s health checks
}

func NewAPIService() *apiService {
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	logger := zerolog.New(output).With().Timestamp().Logger()
	return &apiService{
		addr:   ":80",
		logger: logger,
	}
}

func (s *apiService) WithLogger(logger zerolog.Logger) *apiService {
	s.logger = logger
	return s
}

func (s *apiService) WithAddr(addr string) *apiService {
	s.addr = addr
	return s
}

// Starts the main loop for the server which manages and spins up one or more sub servers that listen for requests
func (s *apiService) Start() {
	s.logger.Info().Msg("go-simple-api server is starting up")
	ctx := ctrl.SetupSignalHandler()
	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return s.Serve(gCtx)
	})

	// Block here until shutdown signal
	err := g.Wait()
	if err != nil {
		s.logger.Info().Msgf("error occured starting go-simple-api server: %s", err.Error())
	}
	s.logger.Info().Msg("go-simple-api server shutting down...")
}

// Starts the sub server for the api endpoints
func (s *apiService) Serve(ctx context.Context) error {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", s.handle404)
	router.HandleFunc("/about", s.handleAbout)
  apiHandler := api.NewAPIHandler()
	apiHandler.AddRoutesToGorillaMux(router)
  loggedMux := s.logMiddleware(router)

	server := &http.Server{
		Addr:    s.addr,
		Handler: loggedMux,
	}

	errChan := make(chan error)
	go func() {
		s.logger.Info().Msgf("starting go-simple-api server, address: %s", s.addr)
		errChan <- server.ListenAndServe()
	}()

	// wait for shut down or error to occur
	select {
	case <-ctx.Done():
		s.logger.Info().Msg("go-simple-api server graceful shutdown started")
		server.Shutdown(ctx)
		s.logger.Info().Msg("go-simple-api server successfully shutdown")
		return nil
	case err := <-errChan:
		s.logger.Info().Msgf("go-simple-api server shutdown with error: %s", err.Error())
		return err
	}
}
