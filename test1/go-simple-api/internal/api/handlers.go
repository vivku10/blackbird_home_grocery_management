package api

import (
	"net/http"
)

// HandleGetHello handles parsing input to pass to the GetHello operation and sends responses back to the client
func (h *APIHandler) HandleGetHello(w http.ResponseWriter, r *http.Request) {
	var err error
	response, err := h.GetHello(r.Context())
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("GetHello returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("GetHello was unable to send it's response, err: %s", err)
	}
}

// HandlePostEcho handles parsing input to pass to the PostEcho operation and sends responses back to the client
func (h *APIHandler) HandlePostEcho(w http.ResponseWriter, r *http.Request) {
	var err error
	response, err := h.PostEcho(r.Context())
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("PostEcho returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("PostEcho was unable to send it's response, err: %s", err)
	}
}

