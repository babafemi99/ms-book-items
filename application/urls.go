package application

import (
	"msitems/controller"
	"net/http"
)

var (
	itemControl = controller.NewItemsController()
)

func mapUrls() {
	router.HandleFunc("/items", itemControl.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", itemControl.Get).Methods(http.MethodGet)
	router.HandleFunc("/items/search", itemControl.Search).Methods(http.MethodPost)
}
