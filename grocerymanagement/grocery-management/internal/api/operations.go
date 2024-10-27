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

// Adds a new grocery item to the refrigerator.
// Add a new grocery item
func (h *APIHandler) AddItem(ctx context.Context, reqBody Item) (Response, error) {
	// TODO: implement the AddItem function to return the following responses

	// return NewResponse(201, Item{}, "application/json", responseHeaders), nil

	// return NewResponse(400, {}, "", responseHeaders), nil

	// return NewResponse(500, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"addItem operation has not been implemented yet"}, "application/json", nil), nil
}

// Deletes a grocery item from the refrigerator.
// Delete a grocery item
func (h *APIHandler) DeleteItem(ctx context.Context, itemId string) (Response, error) {
	// TODO: implement the DeleteItem function to return the following responses

	// return NewResponse(204, {}, "", responseHeaders), nil

	// return NewResponse(404, {}, "", responseHeaders), nil

	// return NewResponse(500, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"deleteItem operation has not been implemented yet"}, "application/json", nil), nil
}

// Retrieves a list of all expired grocery items in the refrigerator.
// Get expired grocery items
func (h *APIHandler) GetExpiredItems(ctx context.Context) (Response, error) {
	// TODO: implement the GetExpiredItems function to return the following responses

	// return NewResponse(200, []Item, "application/json", responseHeaders), nil

	// return NewResponse(404, {}, "", responseHeaders), nil

	// return NewResponse(500, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"getExpiredItems operation has not been implemented yet"}, "application/json", nil), nil
}

// Retrieves a specific grocery item by its ID.
// Get a grocery item by ID
func (h *APIHandler) GetItem(ctx context.Context, itemId string) (Response, error) {
	// TODO: implement the GetItem function to return the following responses

	// return NewResponse(200, Item{}, "application/json", responseHeaders), nil

	// return NewResponse(404, {}, "", responseHeaders), nil

	// return NewResponse(500, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"getItem operation has not been implemented yet"}, "application/json", nil), nil
}

// Retrieves a list of all grocery items in the refrigerator along with their quantities and expiration dates.
// List all grocery items
func (h *APIHandler) ListItems(ctx context.Context) (Response, error) {
	// TODO: implement the ListItems function to return the following responses

	// return NewResponse(200, []Item, "application/json", responseHeaders), nil

	// return NewResponse(404, {}, "", responseHeaders), nil

	// return NewResponse(500, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"XXXXXXXXXXXX - listItems operation has not been implemented yet"}, "application/json", nil), nil
}

// Searches for grocery items based on a query string.
// Search for a grocery item
func (h *APIHandler) SearchItem(ctx context.Context, query string) (Response, error) {
	// TODO: implement the SearchItem function to return the following responses

	// return NewResponse(200, []Item, "application/json", responseHeaders), nil

	// return NewResponse(404, {}, "", responseHeaders), nil

	// return NewResponse(500, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"XXXXXXXXX -- searchItem operation has not been implemented yet"}, "application/json", nil), nil
}

// Updates the details of an existing grocery item.
// Update a grocery item
func (h *APIHandler) UpdateItem(ctx context.Context, itemId string, reqBody Item) (Response, error) {
	// TODO: implement the UpdateItem function to return the following responses

	// return NewResponse(200, Item{}, "application/json", responseHeaders), nil

	// return NewResponse(404, {}, "", responseHeaders), nil

	// return NewResponse(500, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"updateItem operation has not been implemented yet"}, "application/json", nil), nil
}
