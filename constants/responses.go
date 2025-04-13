package constants

type AccessTokenRequestResponse struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	ExpiresIn    int         `json:"expires_in"`
	TokenType    string      `json:"token_type"`
	Scopes       []ScopeType `json:"scopes"`
}

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
