package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/tupt0101/student-api-service/src/utils/errors"
)

const (
	headerXPublic   = "X-Public"
	headerXCallerId = "X-Caller-Id"

	paramAccessToken = "access_token"
)

var (
	oauthRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:9092",
		Timeout: 200 * time.Millisecond,
	}
)

type accessToken struct {
	Id       string `json:"id"`
	UserId   int64  `json:"user_id"`
	ClientId int64  `json:"client_id"`
}

func AuthenticateRequest(request *http.Request) *errors.RestErr {
	if request == nil {
		return nil
	}

	cleanRequest(request)

	accessTokenId := request.URL.Query().Get(paramAccessToken)
	if accessTokenId == "" {
		return nil
	}

	at, err := getAccessToken(accessTokenId)
	if err != nil {
		return err
	}

	// request.Header.Add(headerXClientId, fmt.Sprintf("%v", at.ClientId))
	request.Header.Add(headerXCallerId, fmt.Sprintf("%v", at.UserId))

	return nil
}

func cleanRequest(request *http.Request) {
	if request == nil {
		return
	}
	// request.Header.Del(headerXClientId)
	// request.Header.Del(headerXCallerId)
}

func getAccessToken(accessTokenId string) (*accessToken, *errors.RestErr) {
	response := oauthRestClient.Get(fmt.Sprintf("/oauth/access_token/%s", accessTokenId))
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to get access token")
	}

	if response.StatusCode > 299 {
		var restErr *errors.RestErr
		if err := json.Unmarshal(response.Bytes(), &restErr); err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to get access token")
		}
		return nil, restErr
	}

	var at accessToken
	if err := json.Unmarshal(response.Bytes(), &at); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal access token response")
	}
	return &at, nil
}
