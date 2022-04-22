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
}
