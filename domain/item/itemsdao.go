package item

import (
	"github.com/pkg/errors"
	"msitems/client/es"
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
