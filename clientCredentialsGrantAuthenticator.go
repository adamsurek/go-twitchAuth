package go_twitchAuth

import (
	"encoding/json"
	"errors"
	"fmt"
	. "go-twitchAuth/constants"
	"io"
	"net/http"
	"time"
)

type ClientCredentialsGrantAuthenticator struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

func NewClientCredentialsGrantAuthenticator(clientId string, clientSecret string) *ClientCredentialsGrantAuthenticator {
	return &ClientCredentialsGrantAuthenticator{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		GrantType:    "client_credentials",
	}
}

func (a *ClientCredentialsGrantAuthenticator) GetToken() (*AccessTokenRequestResponse, *FailedRequestResponse, error) {
	var t AccessTokenRequestResponse
	var f FailedRequestResponse

	req, err := http.NewRequest("POST", TokenUrl, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	q := req.URL.Query()
	q.Add("client_id", a.ClientId)
	q.Add("client_secret", a.ClientSecret)
	q.Add("grant_type", a.GrantType)
	req.URL.RawQuery = q.Encode()

	client := http.Client{Timeout: 60 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	if res.StatusCode != 200 {
		err = json.Unmarshal(b, &f)
		if err != nil {
			e := fmt.Sprintf("error while parsing failed request response: %s", err)
			return nil, nil, errors.New(e)
		}
		return nil, &f, nil
	}

	err = json.Unmarshal(b, &t)
	if err != nil {
		e := fmt.Sprintf("error while parsing valid token response: %s", err)
		return nil, nil, errors.New(e)
	}

	return &t, nil, nil
}
