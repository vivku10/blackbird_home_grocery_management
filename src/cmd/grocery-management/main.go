package main

import (
	api "api/internal/grocery-management"
)

func main() {

	api.NewAPIService().Start()
}
