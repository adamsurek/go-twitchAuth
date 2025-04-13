package go_twitchAuth

import (
	. "go-twitchAuth/constants"
	"net/url"
	"strconv"
	"strings"
)

type ImplicitGrantAuthenticator struct {
	RequestedScopes []ScopeType
	clientId        string
	forceVerify     bool
	redirectUri     string
	scopeNames      []string
	state           string
	responseType    string
}

func NewImplicitGrantAuthenticator(clientId string, forceVerify bool, redirectUri string, scopes []ScopeType, state string) *ImplicitGrantAuthenticator {
	return &ImplicitGrantAuthenticator{
		RequestedScopes: scopes,
		clientId:        clientId,
		forceVerify:     forceVerify,
		redirectUri:     redirectUri,
		state:           state,
		responseType:    "token",
	}
}

func (a *ImplicitGrantAuthenticator) GenerateAuthorizationUri() (*url.URL, error) {
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

func (a *ImplicitGrantAuthenticator) getScopeNames() []string {
	var scopeNames []string
	for _, s := range a.RequestedScopes {
		scopeNames = append(scopeNames, ScopeTypeName[s])
	}

	return scopeNames
}

func (a *ImplicitGrantAuthenticator) UpdateScopes(scopes []ScopeType) (*url.URL, error) {
	a.RequestedScopes = scopes
	return a.GenerateAuthorizationUri()
}
