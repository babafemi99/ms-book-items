package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"msitems/domain/item"
	"msitems/domain/queries"
	"msitems/services"
	"msitems/utils/oauth/oauth"
	"net/http"
	"strconv"
)

type ItemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}

var (
	AuthLib     *oauth.Oauthservice
	itemService = services.NewItemService()
)

type itemsController struct {
}

func (i *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	var query queries.EsQuery
	reqErr := json.NewDecoder(r.Body).Decode(&query)
	if reqErr != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Bad Request !"})
		return
	}
	defer r.Body.Close()
	items, searchErr := itemService.Search(query)
	if searchErr != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(map[string]string{"error": "Bad Request !"})
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(items)
}

func (i *itemsController) Get(w http.ResponseWriter, r *http.Request) {

	//AuthLib = oauth.NewOauth(r)
	//nexp := AuthLib.IsExpired()
	//fmt.Println(nexp)
	//if nexp == false {
	//	w.Header().Set("content-type", "application/json")
	//	w.WriteHeader(http.StatusBadRequest)
	//	json.NewEncoder(w).Encode(map[string]string{"error": "token is invalid"})
	//	return
	//}

	vars := mux.Vars(r)
	itemId := vars["id"]
	get, getErr := itemService.Get(itemId)
	if getErr != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(getErr.Status)
		json.NewEncoder(w).Encode(getErr)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(get)

}

func (i *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	AuthLib = oauth.NewOauth(r)
	nexp := AuthLib.IsExpired()
	fmt.Println(nexp)
	if nexp == false {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "token is invalid"})
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

	item.Seller = strconv.Itoa(int(AuthLib.GetUserDetails().UserID))
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
