package services

import (
	"msitems/domain/item"
	"msitems/domain/queries"
	msErrors "msitems/utils/ms_Error"
)

type ItemServiceInterface interface {
	Create(*item.Item) (*item.Item, *msErrors.RestErrors)
	Get(string) (*item.Item, *msErrors.RestErrors)
	Search(query queries.EsQuery) (*[]item.Item, *msErrors.RestErrors)
}
type itemService struct {
}

func (i *itemService) Search(query queries.EsQuery) (*[]item.Item, *msErrors.RestErrors) {
	dao := item.Item{}
	return dao.Search(query)
}

func (i *itemService) Create(item *item.Item) (*item.Item, *msErrors.RestErrors) {
	err := item.Save()
	if err != nil {
		return nil, err
	}
	return item, nil

}

func (i *itemService) Get(s string) (*item.Item, *msErrors.RestErrors) {
	item := &item.Item{Id: s}
	err := item.GetById()
	if err != nil {
		return nil, err
	}
	return item, nil
}

func NewItemService() ItemServiceInterface {
	return &itemService{}
}
