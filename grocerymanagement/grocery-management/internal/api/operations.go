package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"

	"github.com/jackc/pgx/v5"
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
	// Database connection string
	//DATABASE_URL=postgresql://postgres:BlackbirdAPIashaygayathrivivek1011@database-grocery-blackbird.c95hdma6vizc.us-east-1.rds.amazonaws.com:5432/postgres

	// Database connection string
	//connString := "postgres://username:password@localhost:5432/dbname"
	connString := "postgresql://postgres:BlackbirdAPIashaygayathrivivek1011@database-grocery-blackbird.c95hdma6vizc.us-east-1.rds.amazonaws.com:5432/postgres"

	// Establish a connection
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		//log.Fatalf("Unable to connect to database: %v\n", err)
		return NewResponse(404, ErrorMsg{fmt.Sprintf("Unable to connect to database: %v\n", err)}, "application/json", nil), nil
	}
	defer conn.Close(context.Background())

	// Query the database
	rows, err := conn.Query(context.Background(), "SELECT id, name, category, quantity, expiration_date FROM grocery_items")
	if err != nil {
		//log.Fatalf("Query failed: %v\n", err)
		return NewResponse(404, ErrorMsg{fmt.Sprintf("Query failed: %v\n", err)}, "application/json", nil), nil
	}
	defer rows.Close()

	// Iterate over the rows and print each item
	var items []GroceryItem
	for rows.Next() {
		var item GroceryItem
		err := rows.Scan(&item.id, &item.name, &item.category, &item.quantity, &item.expiration_date)
		if err != nil {
			//log.Fatalf("Failed to scan row: %v\n", err)
			return NewResponse(404, ErrorMsg{fmt.Sprintf("Failed to scan row: %v\n", err)}, "application/json", nil), nil
		}
		items = append(items, item)
	}

	if rows.Err() != nil {
		//log.Fatalf("Error iterating rows: %v\n", rows.Err())
		return NewResponse(404, ErrorMsg{fmt.Sprintf("Error iterating rows: %v\n", rows.Err())}, "application/json", nil), nil
	}

	// Convert items slice to JSON
	jsonData, err := json.Marshal(items)
	if err != nil {
		//log.Fatalf("Failed to convert to JSON: %v", err)
		return NewResponse(404, ErrorMsg{fmt.Sprintf("Failed to convert to JSON: %v", err)}, "application/json", nil), nil
	}

	// Convert JSON byte slice to a string and print it
	jsonString := string(jsonData)
	//fmt.Println(jsonString)
	return NewResponse(200, jsonString, "application/json", nil), nil

	// Print retrieved items
	//fmt.Println("Items in database:")
	//for _, item := range items {
	//	fmt.Printf("ID: %d, Name: %s, Quantity: %d\n", item.id, item.name, item.category, item.quantity, item.expirationDate)
	//}

	//return NewResponse(404, ErrorMsg{"ListItems API failed"}, "application/json", nil), nil

	// TODO: implement the ListItems function to return the following responses

	/*
		itemX := Item{
			ExpirationDate: "2025-12-27",
			Id:             "2",
			Name:           "Curd",
			Quantity:       5,
		}

		return NewResponse(200, itemX, "application/json", nil), nil
	*/
	// return NewResponse(404, {}, "", responseHeaders), nil

	// return NewResponse(500, {}, "", responseHeaders), nil

	//return NewResponse(http.StatusNotImplemented, ErrorMsg{"XXXXXXXXXXXX - listItems operation has not been implemented yet"}, "application/json", nil), nil
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
