# Golang Twitch Authentication Wrapper

A wrapper around Twitch's primary OAuth authentication flows.

## Features

- Support for the following OAuth grant flows:
  - Implicit Code grant flow
  - Authorization Code grant flow
  - Client Credentials grant flow
- Token validation and revocation

> [!IMPORTANT]
> Device Code Grant Flow is currently not supported. My intention with this library is to provide
> a fairly simple wrapper around the more common authentication flows.

## Project Status

Still a work in progress. I am still actively working on cleaning up, simplifying, and documenting various
pieces of this library.

There are also a handful of items that I intend to implement in the near future:

- TESTS
- Token refreshes when following the Authorization Code Grant flow

## Installation

`go get github.com/adamsurek/go-twitchAuth`

## Usage

### Implicit Code Grant Flow

```go
package main

import (
  ta "github.com/adamsurek/go-twitchAuth"
  "log"
)

func main() {
  // Define the scopes you'd like the user to authorize
  s := []ta.ScopeType{
    ta.ScopeUserReadChat,
    ta.ScopeUserReadEmotes,
    ta.ScopeUserReadFollows,
    ta.ScopeUserReadSubscriptions,
    ta.ScopeUserManageWhispers,
  }

  // Initialize the authenticator
  a := ta.NewImplicitGrantAuthenticator(
    "{YOUR_CLIENT_ID}",    // Client ID
    true,                  // Force Verify?
    "{YOUR_REDIRECT_URI}", // Redirect URI
    s,                     // Scopes
    "{STATE}",             // State
  )

  // Generate an authorization URL that the user can follow to authorize your app
  u, err := a.GenerateAuthorizationUrl()
  if err != nil {
    log.Fatalf("failed to generate auth url: %s", err)
  }

  // Provide URL to user
  log.Println(u.String())
  
  // ...Retrieve token after user has authorized app
}

```

### Authorization Code Grant Flow

```go
package main

import (
  ta "github.com/adamsurek/go-twitchAuth"
  "log"
)

func main() {
  // Define the scopes you'd like the user to authorize
  s := []ta.ScopeType{
    ta.ScopeUserReadChat,
    ta.ScopeUserReadEmotes,
    ta.ScopeUserReadFollows,
    ta.ScopeUserReadSubscriptions,
    ta.ScopeUserManageWhispers,
  }

  // Initialize the authenticator
  a := ta.NewAuthorizationCodeGrantAuthenticator(
    "{YOUR_CLIENT_ID}",     // Client ID
    "{YOUR_CLIENT_SECRET}", // Client Secret
    true,                   // Force Verify?
    "{YOUR_REDIRECT_URI}",  // Redirect URI
    s,                      // Scopes
    "{STATE}",              // State
  )

  // Generate an authorization URL that the user can follow to authorize your app
  u, err := a.GenerateAuthorizationUrl()
  if err != nil {
    log.Fatalf("failed to generate auth url: %s", err)
  }

  log.Println(u.String())

  // ...Provide auth URL to user
  // ...Retrieve authorization code from Twitch to your redirect URL

  // Exchange code for token
  t, err := a.GetToken("{CODE_FROM_TWITCH}")
  if err != nil {
    log.Fatalf("failed to send code exchange request: %s", err)
  }

  // Ensure that the exchange returned a token
  if t.TokenRequestStatus == ta.StatusSuccess {
    log.Println(t.TokenData.AccessToken)
  } else {
    // ex.: 400 - Invalid authorization code
    log.Fatalf("exchange did not succeed: %d - %s", t.FailureData.Status, t.FailureData.Message)
  }
}

```

### Client Credentials Grant Flow

```go
package main

import (
	ta "github.com/adamsurek/go-twitchAuth"
	"log"
)

func main() {
	// Initialize the authenticator
	a := ta.NewClientCredentialsGrantAuthenticator(
		"{YOUR_CLIENT_ID}",     // Client ID
		"{YOUR_CLIENT_SECRET}", // Client Secret
	)

	// Request token from Twitch API
	t, err := a.GetToken()
	if err != nil {
		log.Fatalf("failed to send token request: %s", err)
	}

	// Ensure that the response contains a token
	if t.TokenRequestStatus == ta.StatusSuccess {
		log.Println(t.TokenData.AccessToken)
	} else {
		// ex.: 403 - invalid client secret
		log.Fatalf("token request did not succeed: %d - %s", t.FailureData.Status, t.FailureData.Message)
	}
}
```

### Validating and Revoking Tokens

```go
package main

import (
  ta "github.com/adamsurek/go-twitchAuth"
  "log"
)

func main() {
  // Send token to Twitch API for validation
  v, err := ta.ValidateToken("{YOUR_TOKEN}")
  if err != nil {
    log.Fatalf("failed to send token validation request: %s", err)
  }

  // Confirm that token is valid
  if v.ValidationStatus == ta.StatusSuccess {
    log.Printf("token valid for user %s. expires in %d seconds",
      v.ValidationData.Login, v.ValidationData.ExpiresIn)
  } else {
    // ex.: 401 - invalid access token
    log.Fatalf("token invalid: %d - %s", v.FailureData.Status, v.FailureData.Message)
  }

  // Send token revocation request
  r, err := ta.RevokeToken("{YOUR_CLIENT_ID}", "{YOUR_TOKEN")
  if err != nil {
    log.Fatalf("failed to send token revocation request: %s", err)
  }

  // Confirm that revocation was successful
  if r.RevocationStatus == ta.StatusSuccess {
    log.Println("token revoked")
  } else {
    // ex.: 400 - token Invalid token
    log.Fatalf("failed to revoke token: %s", err)
  }
}
```