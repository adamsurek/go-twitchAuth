package go_twitchAuth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	. "go-twitchAuth/constants"
	"io"
	"net/http"
	"net/url"
	"time"
)

type ValidTokenResponse struct {
	ClientId  string      `json:"client_id"`
	Login     string      `json:"login"`
	Scopes    []ScopeType `json:"scopes"`
	UserId    string      `json:"user_id"`
	ExpiresIn int         `json:"expires_in"`
}

type FailedRequestResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ImplicitGrantAuthenticator struct {
	RequestedScopesTypes []ScopeType `url:"-"`
	ClientId             string      `url:"client_id"`
	ForceVerify          bool        `url:"force_verify,omitempty"`
	RedirectUri          string      `url:"redirect_uri"`
	ScopeNames           []string    `url:"scope,space"`
	State                string      `url:"state,omitempty"`
	ResponseType         string      `url:"response_type"`
}

func NewImplicitGrantAuthenticator(clientId string, forceVerify bool, redirectUri string, scopes []ScopeType, state string) *ImplicitGrantAuthenticator {
	return &ImplicitGrantAuthenticator{
		RequestedScopesTypes: scopes,
		ClientId:             clientId,
		ForceVerify:          forceVerify,
		RedirectUri:          redirectUri,
		State:                state,
		ResponseType:         "token",
	}
}

func (a *ImplicitGrantAuthenticator) AuthorizationUri() (*url.URL, error) {
	authUrl, err := url.Parse(BASE_OAUTH_URL + "/authorize")
	if err != nil {
		return nil, err
	}

	authUrl.RawQuery, err = a.encodeQueryParams()
	return authUrl, err
}

func (a *ImplicitGrantAuthenticator) encodeQueryParams() (string, error) {
	a.ScopeNames = a.getScopeNames()

	v, err := query.Values(a)
	if err != nil {
		return "", err
	}

	return v.Encode(), nil
}

func (a *ImplicitGrantAuthenticator) getScopeNames() []string {
	var scopeNames []string
	for _, s := range a.RequestedScopesTypes {
		scopeNames = append(scopeNames, ScopeTypeName[s])
	}

	return scopeNames
}

func (a *ImplicitGrantAuthenticator) UpdateScopes(scopes []ScopeType) (*url.URL, error) {
	a.RequestedScopesTypes = scopes
	return a.AuthorizationUri()
}

func (a *ImplicitGrantAuthenticator) ValidateToken(token string) (*ValidTokenResponse, *FailedRequestResponse, error) {
	var v ValidTokenResponse
	var f FailedRequestResponse

	req, err := http.NewRequest("GET", BASE_OAUTH_URL+"/validate", nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
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

	err = json.Unmarshal(b, &v)
	if err != nil {
		e := fmt.Sprintf("error while parsing valid token response: %s", err)
		return nil, nil, errors.New(e)
	}

	return &v, nil, nil
}

func (a *ImplicitGrantAuthenticator) RevokeToken(token string) (*FailedRequestResponse, error) {
	var f FailedRequestResponse

	req, err := http.NewRequest("POST", BASE_OAUTH_URL+"/revoke", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	q := req.URL.Query()
	q.Add("token", token)
	q.Add("client_id", a.ClientId)
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
		err = json.Unmarshal(b, &f)
		if err != nil {
			e := fmt.Sprintf("error while parsing failed request response: %s", err)
			return nil, errors.New(e)
		}
		return &f, nil
	}

	return nil, nil
}
