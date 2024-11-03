package api

import (
	"context"
	"fmt"
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

// Create a connection to the database.
func dbConnection() (*pgx.Conn, error) {
	connStr := "postgresql://postgres:BlackbirdAPIashaygayathrivivek1011@database-grocery-blackbird.c95hdma6vizc.us-east-1.rds.amazonaws.com:5432/postgres"

	// Open the database connection
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return conn, nil
}

// getGroceryItems retrieves all items from the `grocery_items` table.
func getGroceryItems() ([]Item, error) {
	// Get the database connection
	conn, err := dbConnection()
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background()) // Ensure the connection is closed after usage

	// Query the database
	rows, err := conn.Query(context.Background(), "SELECT id, name, category, quantity, expiration_date FROM grocery_items")
	if err != nil {
		//log.Fatalf("Query failed: %v\n", err)
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	// Process the rows into a slice of Item structs
	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.Id, &item.Name, &item.Category, &item.Quantity, &item.ExpirationDate); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		items = append(items, item)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return items, nil
}

// getExpiredGroceryItems retrieves all items from the `grocery_items` table which has passed expiration date.
func getExpiredGroceryItems() ([]Item, error) {
	// Get the database connection
	conn, err := dbConnection()
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background()) // Ensure the connection is closed after usage

	// Query the database
	now := time.Now().Format(time.RFC3339) //.String()
	query := "SELECT id, name, category, quantity, expiration_date FROM grocery_items WHERE expiration_date < $1"

	rows, err := conn.Query(context.Background(), query, now)
	if err != nil {
		//log.Fatalf("Query failed: %v\n", err)
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	// Process the rows into a slice of Item structs
	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.Id, &item.Name, &item.Category, &item.Quantity, &item.ExpirationDate); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		items = append(items, item)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return items, nil
}

// deleteGroceryItemById deletes the grocery details of an item by its ID.
func deleteGroceryItemById(itemId string) (*Item, error) {
	// Get the database connection
	conn, err := dbConnection()
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background()) // Ensure the connection is closed after usage

	// Prepare the query using a variable
	query := "DELETE FROM grocery_items WHERE id = $1"

	// Execute the DELETE statement
	result, err := conn.Exec(context.Background(), query, itemId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete row: %w", err)
	}

	// Check if any row was deleted
	if result.RowsAffected() == 0 {
		return nil, fmt.Errorf("no row found with id %s", itemId)
	}

	return nil, nil
}

// getGroceryItemById retrieves the grocery details of an item by its ID.
func getGroceryItemById(itemId string) (*Item, error) {
	// Get the database connection
	conn, err := dbConnection()
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background()) // Ensure the connection is closed after usage

	// Prepare the query using a variable
	var item Item
	query := "SELECT id, name, category, quantity, expiration_date FROM grocery_items WHERE id = $1"

	// Execute the query with the id variable
	err = conn.QueryRow(context.Background(), query, itemId).Scan(&item.Id, &item.Name, &item.Category, &item.Quantity, &item.ExpirationDate)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	return &item, nil
}

// getGroceryItemByName retrieves the grocery details of an item by its name.
func getGroceryItemByName(name string) (*Item, error) {
	// Get the database connection
	conn, err := dbConnection()
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background()) // Ensure the connection is closed after usage

	// Prepare the query using a variable
	var item Item
	query := "SELECT id, name, category, quantity, expiration_date FROM grocery_items WHERE name = $1"

	// Execute the query with the id variable
	err = conn.QueryRow(context.Background(), query, name).Scan(&item.Id, &item.Name, &item.Category, &item.Quantity, &item.ExpirationDate)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	return &item, nil
}

// updateGroceryItemById updates the grocery details of an item by its Id.
func updateGroceryItemById(itemId string, itemDetails Item) (*Item, error) {
	// Get the database connection
	conn, err := dbConnection()
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background()) // Ensure the connection is closed after usage

	// Update query
	query := "UPDATE grocery_items SET name = $1, category = $2, quantity = $3, expiration_date = $4 WHERE id = $5"
	_, err = conn.Exec(context.Background(), query, itemDetails.Name, itemDetails.Category, itemDetails.Quantity, itemDetails.ExpirationDate, itemId)
	if err != nil {
		return nil, fmt.Errorf("failed to update item: %w", err)
	}

	return nil, nil
}

// addGroceryItem adds a grocery item.
func addGroceryItem(itemDetails Item) (*Item, error) {
	// Get the database connection
	conn, err := dbConnection()
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}
	defer conn.Close(context.Background()) // Ensure the connection is closed after usage

	// Insert query
	query := "INSERT INTO grocery_items(name, category, quantity, expiration_date) VALUES($1, $2, $3, $4)"

	// Execute the INSERT statement
	_, err = conn.Exec(context.Background(), query, itemDetails.Name, itemDetails.Category, itemDetails.Quantity, itemDetails.ExpirationDate)
	if err != nil {
		return nil, fmt.Errorf("failed to insert item: %w", err)
	}

	return nil, nil
}

// Adds a new grocery item to the refrigerator.
// Add a new grocery item
func (h *APIHandler) AddItem(ctx context.Context, reqBody Item) (Response, error) {

	item, err := addGroceryItem(reqBody)
	if err != nil {
		return NewResponse(404, ErrorMsg{fmt.Sprintf("%v", err)}, "application/json", nil), nil
	}

	return NewResponse(200, item, "application/json", nil), nil
}

// Deletes a grocery item from the refrigerator.
// Delete a grocery item
func (h *APIHandler) DeleteItem(ctx context.Context, itemId string) (Response, error) {

	item, err := deleteGroceryItemById(itemId)
	if err != nil {
		return NewResponse(404, ErrorMsg{fmt.Sprintf("%v", err)}, "application/json", nil), nil
	}

	return NewResponse(200, item, "application/json", nil), nil
}

// Retrieves a list of all expired grocery items in the refrigerator.
// Get expired grocery items
func (h *APIHandler) GetExpiredItems(ctx context.Context) (Response, error) {

	items, err := getExpiredGroceryItems()
	if err != nil {
		return NewResponse(404, ErrorMsg{fmt.Sprintf("%v", err)}, "application/json", nil), nil
	}

	return NewResponse(200, items, "application/json", nil), nil
}

// Retrieves a specific grocery item by its ID.
// Get a grocery item by ID
func (h *APIHandler) GetItem(ctx context.Context, itemId string) (Response, error) {

	item, err := getGroceryItemById(itemId)
	if err != nil {
		return NewResponse(404, ErrorMsg{fmt.Sprintf("%v", err)}, "application/json", nil), nil
	}

	return NewResponse(200, item, "application/json", nil), nil
}

// List all items
// Retrieves a list of all grocery items in the refrigerator along with their quantities and expiration dates.
func (h *APIHandler) ListItems(ctx context.Context) (Response, error) {

	items, err := getGroceryItems()
	if err != nil {
		return NewResponse(404, ErrorMsg{fmt.Sprintf("%v", err)}, "application/json", nil), nil
	}

	return NewResponse(200, items, "application/json", nil), nil
}

// Searches for grocery items based on a query string.
// Search for a grocery item
func (h *APIHandler) SearchItem(ctx context.Context, query string) (Response, error) {

	item, err := getGroceryItemByName(query)
	if err != nil {
		return NewResponse(404, ErrorMsg{fmt.Sprintf("%v", err)}, "application/json", nil), nil
	}

	return NewResponse(200, item, "application/json", nil), nil
}

// Updates the details of an existing grocery item.
// Update a grocery item
func (h *APIHandler) UpdateItem(ctx context.Context, itemId string, reqBody Item) (Response, error) {

	item, err := updateGroceryItemById(itemId, reqBody)
	if err != nil {
		return NewResponse(404, ErrorMsg{fmt.Sprintf("%v", err)}, "application/json", nil), nil
	}

	return NewResponse(200, item, "application/json", nil), nil
}
