package go_twitchAuth

import (
	"github.com/google/go-querystring/query"
	"net/url"
)

const baseOAuthUrl = "https://id.twitch.tv/oauth2"

type ImplicitGrantAuthenticator struct {
	oAuthBaseUrl string `url:"-"`

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
		oAuthBaseUrl:         baseOAuthUrl,
		RequestedScopesTypes: scopes,
		ClientId:             clientId,
		ForceVerify:          forceVerify,
		RedirectUri:          redirectUri,
		State:                state,
		ResponseType:         "token",
	}
}

func (a *ImplicitGrantAuthenticator) AuthorizationUri() (*url.URL, error) {
	authUrl, err := url.Parse(a.oAuthBaseUrl + "/authorize")
	if err != nil {
		return nil, err
	}

	authUrl.RawQuery, err = a.encodeQueryParams()
	return authUrl, err
}

func (a *ImplicitGrantAuthenticator) UpdateScopes(scopes []ScopeType) (*url.URL, error) {
	a.RequestedScopesTypes = scopes
	return a.AuthorizationUri()
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

func (a *ImplicitGrantAuthenticator) ValidateToken(token string) (bool, error) {
	return true, nil
}

func (a *ImplicitGrantAuthenticator) RevokeToken(token string) (bool, error) {
	return true, nil
}
