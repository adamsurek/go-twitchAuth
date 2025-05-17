package go_twitchAuth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

/*
AuthorizationCodeGrantAuthenticator allows for the generation of an authorization URL following Twitch's
OAuth client credentials grant flow.

New instances of AuthorizationCodeGrantAuthenticator should be created via
NewAuthorizationCodeGrantAuthenticator.

Twitch docs: https://dev.twitch.tv/docs/authentication/getting-tokens-oauth/#authorization-code-grant-flow
*/
type AuthorizationCodeGrantAuthenticator struct {
	requestedScopes []ScopeType
	clientId        string
	clientSecret    string
	forceVerify     bool
	redirectUri     string
	scopeNames      []string
	state           string
	grantType       string
	responseType    string
}

// NewAuthorizationCodeGrantAuthenticator generates a new AuthorizationCodeGrantAuthenticator instance.
func NewAuthorizationCodeGrantAuthenticator(clientId string, clientSecret string, forceVerify bool, redirectUri string, scopes []ScopeType, state string) *AuthorizationCodeGrantAuthenticator {
	return &AuthorizationCodeGrantAuthenticator{
		requestedScopes: scopes,
		clientId:        clientId,
		clientSecret:    clientSecret,
		forceVerify:     forceVerify,
		redirectUri:     redirectUri,
		state:           state,
		grantType:       "authorization_code",
		responseType:    "code",
	}
}

// GenerateAuthorizationUrl builds a url.URL that allows a user to authorize a Twitch app and generate
// a bearer token.
func (a *AuthorizationCodeGrantAuthenticator) GenerateAuthorizationUrl() (*url.URL, error) {
	authUrl, err := url.Parse(authorizationUrl)
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

// getScopeNames retrieves the string version of the ScopeType(s) supplied to the AuthorizationCodeGrantAuthenticator.
func (a *AuthorizationCodeGrantAuthenticator) getScopeNames() []string {
	var scopeNames []string
	for _, s := range a.requestedScopes {
		scopeNames = append(scopeNames, ScopeTypeName[s])
	}

	return scopeNames
}

// GetToken retrieves a new bearer token via the Twitch Helix API using the auth code generated when the user
// follows the authorization URL.
func (a *AuthorizationCodeGrantAuthenticator) GetToken(code string) (*TokenResponse, error) {
	var t TokenResponse

	req, err := http.NewRequest("POST", tokenUrl, nil)
	if err != nil {
		return nil, err
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

// UpdateScopes replaces the original array of ScopeType provided during initialization
func (a *AuthorizationCodeGrantAuthenticator) UpdateScopes(scopes []ScopeType) (*url.URL, error) {
	a.requestedScopes = scopes
	return a.GenerateAuthorizationUrl()
}

/*
GetScopes retrieves the currently requested list of scopes. It's important to note that the scopes returned
are only what has been supplied to the authenticator - not what the end user has authorized.

To retrieve the scopes that the user has authorized, you can use the ValidateToken function.
*/
func (a *AuthorizationCodeGrantAuthenticator) GetScopes() []ScopeType {
	return a.requestedScopes
}
