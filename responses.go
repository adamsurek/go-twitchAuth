package go_twitchAuth

type responseStatus int

const (
	// StatusFailure signifies a failed Helix API request status (ex. HTTP 400).
	StatusFailure responseStatus = iota
	// StatusSuccess signifies a successful Helix API request status (ex. HTTP 200).
	StatusSuccess
)

var statusName = map[responseStatus]string{
	StatusFailure: "failure",
	StatusSuccess: "success",
}

func (rs responseStatus) String() string {
	return statusName[rs]
}

// TokenResponse stores the results of a token retrieval request.
type TokenResponse struct {
	TokenRequestStatus responseStatus
	TokenData          *AccessTokenRequestResponse
	FailureData        *FailedRequestResponse
}

// TokenValidationResponse stores the results of a token validation request.
type TokenValidationResponse struct {
	ValidationStatus responseStatus
	ValidationData   *ValidTokenResponse
	FailureData      *FailedRequestResponse
}

// TokenRevocationResponse stores the results of a token revocation request. A successful revocation request returns
// no details beyond a RevocationStatus of StatusSuccess.
type TokenRevocationResponse struct {
	RevocationStatus responseStatus
	FailureData      *FailedRequestResponse
}

// AccessTokenRequestResponse stores the parsed JSON response of an access token request.
type AccessTokenRequestResponse struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	ExpiresIn    int         `json:"expires_in"`
	TokenType    string      `json:"token_type"`
	Scopes       []ScopeType `json:"scopes"`
}

// ValidTokenResponse stores the parsed JSON response of a token validation request on a valid token.
type ValidTokenResponse struct {
	ClientId  string      `json:"client_id"`
	Login     string      `json:"login"`
	Scopes    []ScopeType `json:"scopes"`
	UserId    string      `json:"user_id"`
	ExpiresIn int         `json:"expires_in"`
}

// FailedRequestResponse stores the parsed JSON response of a failed Helix API response.
type FailedRequestResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
