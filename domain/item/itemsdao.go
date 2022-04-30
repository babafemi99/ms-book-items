package item

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"msitems/client/es"
	"msitems/domain/queries"
	msErrors "msitems/utils/ms_Error"
)

const (
	index = "items"
)

var (
	elasticService = es.NewEsclient()
)

func (item *Item) Save() *msErrors.RestErrors {
	res, err := elasticService.IndexClient(index, item)
	if err != nil {
		return msErrors.NewInternalServerError("error when trying to index document", errors.New("DB ERROR"))
	}
	item.Id = res.Id
	return nil
}

func (item *Item) GetById() *msErrors.RestErrors {
	itemId := item.Id
	byId, err := elasticService.GetById(index, item.Id)
	if err != nil {
		return msErrors.NewNotFoundRequestError("unable to find document", err)
	}
	if !byId.Found {
		return msErrors.NewNotFoundRequestError("error while trying to find element", errors.New(""))
	}
	jsons, _ := byId.Source.MarshalJSON()
	_ = json.Unmarshal(jsons, item)
	item.Id = itemId
	return nil
}
func (item *Item) Search(query queries.EsQuery) (*[]Item, *msErrors.RestErrors) {
	result, err := elasticService.Search(index, query.Build())
	if err != nil {
		return nil, msErrors.NewInternalServerError("error when trying to search document", errors.New("db error"))
	}
	var items []Item
	if result.Hits.TotalHits.Value > 0 {
		for _, hit := range result.Hits.Hits {
			var item Item
			err := json.Unmarshal(hit.Source, &item)
			if err != nil {
				return nil, msErrors.NewInternalServerError("unable to marshal error", err)
			}
			items = append(items, item)
		}
	}
	fmt.Println(items)
	return &items, nil
}
