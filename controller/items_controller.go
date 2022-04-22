package controller

import (
	"encoding/json"
	"msitems/domain/item"
	"msitems/services"
	"msitems/utils/oauth/oauth"
	"net/http"
)

type ItemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
}

var (
	AuthLib     *oauth.Oauthservice
	itemService = services.NewItemService()
)

type itemsController struct {
}

func (i *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	AuthLib = oauth.NewOauth(r)
	expired := AuthLib.IsExpired()
	if expired {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Session Timeout"})
		return
	}

	var item item.Item

	reqErr := json.NewDecoder(r.Body).Decode(&item)
	if reqErr != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Bad Request !"})
		return
	}
	defer r.Body.Close()

	item.Id = string(AuthLib.GetUserDetails().UserID)
	res, err := itemService.Create(&item)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func NewItemsController() ItemsControllerInterface {
	return &itemsController{}
}
