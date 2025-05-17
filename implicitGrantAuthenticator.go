package go_twitchAuth

import (
	"net/url"
	"strconv"
	"strings"
)

/*
ImplicitGrantAuthenticator allows for the generation of an authorization URL following Twitch's
OAuth implicit grant flow.

New instances of ImplicitGrantAuthenticator should be created via
NewImplicitGrantAuthenticator.

Twitch docs: https://dev.twitch.tv/docs/authentication/getting-tokens-oauth/#implicit-grant-flow
*/
type ImplicitGrantAuthenticator struct {
	requestedScopes []ScopeType
	clientId        string
	forceVerify     bool
	redirectUri     string
	scopeNames      []string
	state           string
	responseType    string
}

// NewImplicitGrantAuthenticator generates a new ImplicitGrantAuthenticator instance.
func NewImplicitGrantAuthenticator(clientId string, forceVerify bool, redirectUri string, scopes []ScopeType, state string) *ImplicitGrantAuthenticator {
	return &ImplicitGrantAuthenticator{
		requestedScopes: scopes,
		clientId:        clientId,
		forceVerify:     forceVerify,
		redirectUri:     redirectUri,
		state:           state,
		responseType:    "token",
	}
}

// GenerateAuthorizationUrl builds a url.URL that allows a user to authorize a Twitch app and generate
// a bearer token.
func (a *ImplicitGrantAuthenticator) GenerateAuthorizationUrl() (*url.URL, error) {
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

// getScopeNames retrieves the string version of the ScopeType(s) supplied to the ImplicitGrantAuthenticator.
func (a *ImplicitGrantAuthenticator) getScopeNames() []string {
	var scopeNames []string
	for _, s := range a.requestedScopes {
		scopeNames = append(scopeNames, scopeTypeName[s])
	}

	return scopeNames
}

// UpdateScopes replaces the original array of ScopeType provided during initialization
func (a *ImplicitGrantAuthenticator) UpdateScopes(scopes []ScopeType) {
	a.requestedScopes = scopes
}

/*
GetScopes retrieves the currently requested list of scopes. It's important to note that the scopes returned
are only what has been supplied to the authenticator - not what the end user has authorized.

To retrieve the scopes that the user has authorized, you can use the ValidateToken function.
*/
func (a *ImplicitGrantAuthenticator) GetScopes() []ScopeType {
	return a.requestedScopes
}
