package api

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// APIHandler is a type to give the api functions below access to a common logger
// any any other shared objects
type APIHandler struct {
	// Zerolog was chosen as the default logger, but you can replace it with any logger of your choice
	logger zerolog.Logger

	// Note: if you need to pass in a client for your database, this would be a good place to include it
}

func NewAPIHandler() *APIHandler {
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	logger := zerolog.New(output).With().Timestamp().Logger()
	return &APIHandler{logger: logger}
}

func (h *APIHandler) WithLogger(logger zerolog.Logger) *APIHandler {
	h.logger = logger
	return h
}

// Get a hello message
func (h *APIHandler) GetHello(ctx context.Context) (Response, error) {
	// TODO: implement the GetHello function to return the following responses

	// return NewResponse(200, {}, "application/json", responseHeaders), nil

	// return NewResponse(4XX, {}, "", responseHeaders), nil

	// return NewResponse(5XX, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"XXXXXX - getHello operation has not been implemented yet"}, "application/json", nil), nil
}

// Echo back the message
func (h *APIHandler) PostEcho(ctx context.Context) (Response, error) {
	// TODO: implement the PostEcho function to return the following responses

	// return NewResponse(200, {}, "application/json", responseHeaders), nil

	// return NewResponse(4XX, {}, "", responseHeaders), nil

	// return NewResponse(5XX, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"XXXXXX - postEcho operation has not been implemented yet"}, "application/json", nil), nil
}
