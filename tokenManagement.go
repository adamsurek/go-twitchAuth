package go_twitchAuth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func ValidateToken(token string) (*ValidTokenResponse, *FailedRequestResponse, error) {
	var v ValidTokenResponse
	var f FailedRequestResponse

	req, err := http.NewRequest("GET", ValidationUrl, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
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

	err = json.Unmarshal(b, &v)
	if err != nil {
		e := fmt.Sprintf("error while parsing valid token response: %s", err)
		return nil, nil, errors.New(e)
	}

	return &v, nil, nil
}

func RevokeToken(token string, clientId string) (*FailedRequestResponse, error) {
	var f FailedRequestResponse

	req, err := http.NewRequest("POST", RevocationUrl, nil)
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
		err = json.Unmarshal(b, &f)
		if err != nil {
			e := fmt.Sprintf("error while parsing failed request response: %s", err)
			return nil, errors.New(e)
		}
		return &f, nil
	}

	return nil, nil
}
