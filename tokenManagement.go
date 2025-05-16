package go_twitchAuth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func ValidateToken(token string) (*TokenValidationResponse, error) {
	var t TokenValidationResponse

	req, err := http.NewRequest("GET", validationUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
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
		t.ValidationStatus = FailureStatus
		err = json.Unmarshal(b, &t.FailureData)
		if err != nil {
			e := fmt.Sprintf("error while parsing failed request response: %s", err)
			return nil, errors.New(e)
		}

		return &t, nil
	}

	t.ValidationStatus = SuccessStatus
	err = json.Unmarshal(b, &t.ValidationData)
	if err != nil {
		e := fmt.Sprintf("error while parsing valid token response: %s", err)
		return nil, errors.New(e)
	}

	return &t, nil
}

func RevokeToken(token string, clientId string) (*TokenRevocationResponse, error) {
	var t TokenRevocationResponse

	req, err := http.NewRequest("POST", revocationUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	q := req.URL.Query()
	q.Add("token", token)
	q.Add("client_id", clientId)
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
		t.ValidationStatus = FailureStatus
		err = json.Unmarshal(b, &t.FailureData)
		if err != nil {
			e := fmt.Sprintf("error while parsing failed request response: %s", err)
			return nil, errors.New(e)
		}
		return &t, nil
	}

	t.ValidationStatus = SuccessStatus
	return &t, nil
}
