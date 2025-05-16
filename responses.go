package go_twitchAuth

type responseStatus int

const (
	SuccessStatus responseStatus = iota
	FailureStatus
)

var statusName = map[responseStatus]string{
	SuccessStatus: "success",
	FailureStatus: "failure",
}

func (rs responseStatus) String() string {
	return statusName[rs]
}

type tokenResponse struct {
	Status      responseStatus
	TokenData   accessTokenRequestResponse
	FailureData failedRequestResponse
}

type tokenValidationResponse struct {
	Status         responseStatus
	ValidationData validTokenResponse
	FailureData    failedRequestResponse
}

type tokenRevocationResponse struct {
	Status      responseStatus
	FailureData failedRequestResponse
}

type accessTokenRequestResponse struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	ExpiresIn    int         `json:"expires_in"`
	TokenType    string      `json:"token_type"`
	Scopes       []ScopeType `json:"scopes"`
}

type validTokenResponse struct {
	ClientId  string      `json:"client_id"`
	Login     string      `json:"login"`
	Scopes    []ScopeType `json:"scopes"`
	UserId    string      `json:"user_id"`
	ExpiresIn int         `json:"expires_in"`
}

type failedRequestResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
