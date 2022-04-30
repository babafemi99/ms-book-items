package oauth

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"msitems/utils/oauth/utils"
	"net/http"
	"testing"
)

var testReq, _ = http.NewRequest("GET", "items", nil)
var privateReq, _ = http.NewRequest("GET", "items", nil)
var publicReq, _ = http.NewRequest("GET", "items", nil)

func TestNewOauthConstant(t *testing.T) {
	assert.Equal(t, "X-Public", publicXHeader)
}

func TestCreateNewRequest(t *testing.T) {
	testReq.Header.Set("Authorization", "Bearer hldvabsnmdoibdncncmdscbnwebwln")
	NewReqSrv := NewOauth(testReq)
	fmt.Printf("%#v", NewReqSrv.request)
	assert.NotNil(t, NewReqSrv.request)
}

func TestOauthservice_GetUserDetails(t *testing.T) {
	type fields struct {
		request *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		want   *utils.SignedDetails
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Oauthservice{
				request: tt.fields.request,
			}
			assert.Equalf(t, tt.want, o.GetUserDetails(), "GetUserDetails()")
		})
	}
}

func TestOauthservice_IsExpired(t *testing.T) {
	type fields struct {
		request *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Oauthservice{
				request: tt.fields.request,
			}
			assert.Equalf(t, tt.want, o.IsExpired(), "IsExpired()")
		})
	}
}

func TestOauthservice_IsPrivate(t *testing.T) {
	publicReq.Header.Set(publicXHeader, "true")
	type fields struct {
		request *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{

		{
			name: "public request",
			fields: fields{
				request: publicReq,
			},
			want: false,
		},
		{
			name: "private request",
			fields: fields{
				request: privateReq,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Oauthservice{
				request: tt.fields.request,
			}
			assert.Equalf(t, tt.want, o.IsPrivate(), "IsPrivate()")
		})
	}
}

func TestOauthservice_IsPublic(t *testing.T) {
	publicReq.Header.Set(publicXHeader, "true")
	type fields struct {
		request *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "public request",
			fields: fields{
				request: publicReq,
			},
			want: true,
		},
		{
			name: "private request",
			fields: fields{
				request: privateReq,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Oauthservice{
				request: tt.fields.request,
			}
			assert.Equalf(t, tt.want, o.IsPublic(), "IsPublic()")
		})
	}
}
