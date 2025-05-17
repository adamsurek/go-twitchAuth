package go_twitchAuth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

/*
ClientCredentialsGrantAuthenticator allows for the generation of an authorization URL following Twitch's
OAuth client credentials grant flow.

New instances of ClientCredentialsGrantAuthenticator should be created via
NewClientCredentialsGrantAuthenticator.

Twitch docs: https://dev.twitch.tv/docs/authentication/getting-tokens-oauth/#client-credentials-grant-flow
*/
type ClientCredentialsGrantAuthenticator struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

// NewClientCredentialsGrantAuthenticator generates a new ClientCredentialsGrantAuthenticator instance.
func NewClientCredentialsGrantAuthenticator(clientId string, clientSecret string) *ClientCredentialsGrantAuthenticator {
	return &ClientCredentialsGrantAuthenticator{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		GrantType:    "client_credentials",
	}
}

// GetToken retrieves a new bearer token via the Twitch Helix API.
func (a *ClientCredentialsGrantAuthenticator) GetToken() (*TokenResponse, error) {
	var t TokenResponse

	req, err := http.NewRequest("POST", tokenUrl, nil)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		t.TokenRequestStatus = StatusFailure
		err = json.Unmarshal(b, &t.FailureData)
		if err != nil {
			e := fmt.Sprintf("error while parsing failed request response: %s", err)
			return nil, errors.New(e)
		}
		return &t, nil
	}

	t.TokenRequestStatus = StatusSuccess
	err = json.Unmarshal(b, &t.TokenData)
	if err != nil {
		e := fmt.Sprintf("error while parsing token response: %s", err)
		return nil, errors.New(e)
	}

	return &t, nil
}
