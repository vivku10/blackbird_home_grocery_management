package api

import (
	"github.com/gorilla/mux"
	"net/http"
)


type Route struct {
	Name    string
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type Routes []Route

func (h *APIHandler) AddRoutesToGorillaMux(router *mux.Router) {
	for _, route := range h.GetRoutes() {
		router.
			Name(route.Name).
			Path(route.Path).
			Methods(route.Method).
			Handler(route.Handler)
	}
}

func (h *APIHandler) GetRoutes() Routes {
	return Routes{
		{
			"listItems",
			"/items",
			"GET",
			h.HandleListItems,
		},{
			"addItem",
			"/items",
			"POST",
			h.HandleAddItem,
		},{
			"getExpiredItems",
			"/items/expired",
			"GET",
			h.HandleGetExpiredItems,
		},{
			"searchItem",
			"/items/search",
			"GET",
			h.HandleSearchItem,
		},{
			"deleteItem",
			"/items/{itemId}",
			"DELETE",
			h.HandleDeleteItem,
		},{
			"getItem",
			"/items/{itemId}",
			"GET",
			h.HandleGetItem,
		},{
			"updateItem",
			"/items/{itemId}",
			"PUT",
			h.HandleUpdateItem,
		},
	}
}

