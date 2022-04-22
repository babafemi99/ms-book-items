package oauth

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var testReq, _ = http.NewRequest("GET", "items", nil)

func TestCreateNewRequest(t *testing.T) {
	testReq.Header.Set("Authorization", "Bearer hldvabsnmdoibdncncmdscbnwebwln")
	NewReqSrv := NewOauth(testReq)
	fmt.Printf("%#v", NewReqSrv.request)
	assert.NotNil(t, NewReqSrv.request)
}
