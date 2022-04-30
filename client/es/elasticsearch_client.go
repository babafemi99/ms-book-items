package es

import (
	"context"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	"msitems/logger"
)

type EsClient interface {
	IndexClient(index string, doc interface{}) (*elastic.IndexResponse, error)
	GetById(index, id string) (*elastic.GetResult, error)
	Search(index string, query elastic.Query) (*elastic.SearchResult, error)
}

type esclient struct {
	client *elastic.Client
}

func (cl *esclient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()
	result, err := cl.client.Search(index).Query(query).RestTotalHitsAsInt(true).Do(ctx)
	if err != nil {
		fmtErr := fmt.Sprintf("Error while searchin for document: %s", index)
		logger.Error(fmtErr, err)
		return nil, err
	}
	return result, nil
}

func (cl *esclient) GetById(index, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	res, err := cl.client.Get().Index(index).Id(id).Do(ctx)
	if err != nil {
		fmtErr := fmt.Sprintf("Error while indexing document: %s", index)
		logger.Error(fmtErr, err)
		return nil, err
	}
	if !res.Found {
		return nil, errors.New("no result found")
	}
	return res, nil
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
		fmtErr := fmt.Sprintf("Error while finding document in index: %s", index)
		logger.Error(fmtErr, err)
		return nil, err
	}
	return res, nil
}
