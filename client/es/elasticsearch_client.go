package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"msitems/logger"
)

type EsClient interface {
	IndexClient(index string, doc interface{}) (*elastic.IndexResponse, error)
}

type esclient struct {
	client *elastic.Client
}

func NewEsclient() EsClient {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		panic(err)
	}
	return &esclient{client: client}
}

func (cl *esclient) IndexClient(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	res, err := cl.client.Index().Index(index).BodyJson(doc).Human(true).Do(ctx)
	if err != nil {
		fmtErr := fmt.Sprintf("Error while indexing document: %s", index)
		logger.Error(fmtErr, err)
		return nil, err
	}
	return res, nil
}
