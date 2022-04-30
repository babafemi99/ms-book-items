package oauth

import (
	"errors"
	"fmt"
	msErrors "msitems/utils/ms_Error"
	"msitems/utils/oauth/utils"

	"net/http"
	"time"
)

const (
	publicXHeader = "X-Public"
)

type Oauthservice struct {
	request *http.Request
}

func (o *Oauthservice) Authenticate() *msErrors.RestErrors {
	if o.request == nil {
		return msErrors.NewBadRequest("Bad request", errors.New("check if request is empty"))
	}
	return nil
}

func NewOauth(request *http.Request) *Oauthservice {
	return &Oauthservice{request: request}
}

func (o *Oauthservice) IsExpired() bool {
	user := o.GetUserDetails()
	if user == nil {
		fmt.Println("user is nil ")
		return false
	}
	if user.VerifyExpiresAt(time.Now(), true) {
		return true
	}
	return true
}

func (o *Oauthservice) IsPublic() bool {
	if o.request == nil {
		return true
	}
	return o.request.Header.Get(publicXHeader) == "true"
}

func (o *Oauthservice) IsPrivate() bool {
	return !o.IsPublic()
}

func (o *Oauthservice) GetUserDetails() *utils.SignedDetails {
	user, msg := utils.ValidateToken(o.request)
	if msg != "" {
		fmt.Println("message is ", msg)
		return nil
	}
	return user
}
