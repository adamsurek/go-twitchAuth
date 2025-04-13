package go_twitchAuth

import (
	"encoding/json"
	"errors"
	"fmt"
	. "go-twitchAuth/constants"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type AuthorizationCodeGrantAuthenticator struct {
	RequestedScopesTypes []ScopeType
	clientId             string
	clientSecret         string
	forceVerify          bool
	redirectUri          string
	scopeNames           []string
	state                string
	grantType            string
	responseType         string
}

func NewAuthorizationCodeGrantAuthenticator(clientId string, clientSecret string, forceVerify bool, redirectUri string, scopes []ScopeType, state string) *AuthorizationCodeGrantAuthenticator {
	return &AuthorizationCodeGrantAuthenticator{
		RequestedScopesTypes: scopes,
		clientId:             clientId,
		clientSecret:         clientSecret,
		forceVerify:          forceVerify,
		redirectUri:          redirectUri,
		state:                state,
		grantType:            "authorization_code",
		responseType:         "code",
	}
}

func (a *AuthorizationCodeGrantAuthenticator) GenerateAuthorizationUrl() (*url.URL, error) {
	authUrl, err := url.Parse(AuthorizationUrl)
	if err != nil {
		return nil, err
	}

	q := authUrl.Query()
	q.Add("client_id", a.clientId)
	q.Add("force_verify", strconv.FormatBool(a.forceVerify))
	q.Add("redirect_uri", a.redirectUri)
	q.Add("response_type", a.responseType)
	q.Add("scope", strings.Join(a.getScopeNames(), " "))

	if a.state != "" {
		q.Add("state", a.state)
	}

	authUrl.RawQuery = q.Encode()

	return authUrl, err
}

func (a *AuthorizationCodeGrantAuthenticator) getScopeNames() []string {
	var scopeNames []string
	for _, s := range a.RequestedScopesTypes {
		scopeNames = append(scopeNames, ScopeTypeName[s])
	}

	return scopeNames
}

func (a *AuthorizationCodeGrantAuthenticator) GetToken(code string) (*AccessTokenRequestResponse, *FailedRequestResponse, error) {
	var t AccessTokenRequestResponse
	var f FailedRequestResponse

	req, err := http.NewRequest("POST", TokenUrl, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	q := req.URL.Query()
	q.Add("client_id", a.clientId)
	q.Add("client_secret", a.clientSecret)
	q.Add("code", code)
	q.Add("grant_type", a.grantType)
	q.Add("redirect_uri", a.redirectUri)
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

func (a *AuthorizationCodeGrantAuthenticator) UpdateScopes(scopes []ScopeType) (*url.URL, error) {
	a.RequestedScopesTypes = scopes
	return a.GenerateAuthorizationUrl()
}
