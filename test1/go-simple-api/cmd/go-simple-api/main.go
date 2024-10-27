package main

import (
	api "api/internal/go-simple-api"
)

func main() {

	api.NewAPIService().Start()
}
