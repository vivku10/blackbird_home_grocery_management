package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

// HandleAddItem handles parsing input to pass to the AddItem operation and sends responses back to the client
func (h *APIHandler) HandleAddItem(w http.ResponseWriter, r *http.Request) {
	var err error
	reqBody := Item{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reqBody); err != nil {
		ErrorResponseWithMsg(http.StatusBadRequest, "request body was not able to be parsed successfully 'Item'", w)
		return
	}
	if err := reqBody.Validate(); err != nil {
		errMsg := fmt.Errorf("request body was parsed successfully but failed validation, err: %w", err)
		ErrorResponseWithMsg(http.StatusBadRequest, errMsg.Error(), w)
		return
	}

	response, err := h.AddItem(r.Context(), reqBody)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("AddItem returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("AddItem was unable to send it's response, err: %s", err)
	}
}

// HandleDeleteItem handles parsing input to pass to the DeleteItem operation and sends responses back to the client
func (h *APIHandler) HandleDeleteItem(w http.ResponseWriter, r *http.Request) {
	var err error
	pathParams := mux.Vars(r)

	var itemId string
	itemId = pathParams["itemId"]
	if itemId == "" {
		ErrorResponseWithMsg(http.StatusBadRequest, "request is missing required path parameter 'itemId'", w)
		return
	}

	response, err := h.DeleteItem(r.Context(), itemId)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("DeleteItem returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("DeleteItem was unable to send it's response, err: %s", err)
	}
}

// HandleGetExpiredItems handles parsing input to pass to the GetExpiredItems operation and sends responses back to the client
func (h *APIHandler) HandleGetExpiredItems(w http.ResponseWriter, r *http.Request) {
	var err error
	response, err := h.GetExpiredItems(r.Context())
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("GetExpiredItems returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("GetExpiredItems was unable to send it's response, err: %s", err)
	}
}

// HandleGetItem handles parsing input to pass to the GetItem operation and sends responses back to the client
func (h *APIHandler) HandleGetItem(w http.ResponseWriter, r *http.Request) {
	var err error
	pathParams := mux.Vars(r)

	var itemId string
	itemId = pathParams["itemId"]
	if itemId == "" {
		ErrorResponseWithMsg(http.StatusBadRequest, "request is missing required path parameter 'itemId'", w)
		return
	}

	response, err := h.GetItem(r.Context(), itemId)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("GetItem returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("GetItem was unable to send it's response, err: %s", err)
	}
}

// HandleListItems handles parsing input to pass to the ListItems operation and sends responses back to the client
func (h *APIHandler) HandleListItems(w http.ResponseWriter, r *http.Request) {
	var err error
	response, err := h.ListItems(r.Context())
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("ListItems returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("ListItems was unable to send it's response, err: %s", err)
	}
}

// HandleSearchItem handles parsing input to pass to the SearchItem operation and sends responses back to the client
func (h *APIHandler) HandleSearchItem(w http.ResponseWriter, r *http.Request) {
	var err error
	queryVals, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("SearchItem was unable to parse the query parameters, err: %s", err)
		return
	}

	var query string
	if queryVals.Has("query") {
		query = queryVals.Get("query")
	} else {
		ErrorResponseWithMsg(http.StatusBadRequest, "request is missing required query parameter 'query'", w)
		return
	}

	response, err := h.SearchItem(r.Context(), query)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("SearchItem returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("SearchItem was unable to send it's response, err: %s", err)
	}
}

// HandleUpdateItem handles parsing input to pass to the UpdateItem operation and sends responses back to the client
func (h *APIHandler) HandleUpdateItem(w http.ResponseWriter, r *http.Request) {
	var err error
	pathParams := mux.Vars(r)

	var itemId string
	itemId = pathParams["itemId"]
	if itemId == "" {
		ErrorResponseWithMsg(http.StatusBadRequest, "request is missing required path parameter 'itemId'", w)
		return
	}

	reqBody := Item{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reqBody); err != nil {
		ErrorResponseWithMsg(http.StatusBadRequest, "request body was not able to be parsed successfully 'Item'", w)
		return
	}
	if err := reqBody.Validate(); err != nil {
		errMsg := fmt.Errorf("request body was parsed successfully but failed validation, err: %w", err)
		ErrorResponseWithMsg(http.StatusBadRequest, errMsg.Error(), w)
		return
	}

	response, err := h.UpdateItem(r.Context(), itemId, reqBody)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("UpdateItem returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("UpdateItem was unable to send it's response, err: %s", err)
	}
}
