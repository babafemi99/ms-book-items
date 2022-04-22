package services

import (
	"msitems/domain/item"
	msErrors "msitems/utils/ms_Error"
)

type ItemServiceInterface interface {
	Create(*item.Item) (*item.Item, *msErrors.RestErrors)
	Get(string) (*item.Item, *msErrors.RestErrors)
}
type itemService struct {
}

func (i *itemService) Create(item *item.Item) (*item.Item, *msErrors.RestErrors) {
	err := item.Save()
	if err != nil {
		return nil, err
	}
	return item, nil

}

func (i *itemService) Get(s string) (*item.Item, *msErrors.RestErrors) {
	//TODO implement me
	panic("implement me")
}

func NewItemService() ItemServiceInterface {
	return &itemService{}
}
