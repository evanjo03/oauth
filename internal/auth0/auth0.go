package auth0

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type (
	Request struct {
		GrantType    string `json:"grant_type"`
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		Audience     string `json:"audience"`
		Scope        string `json:"scope"`
	}
	Response struct {
		AccessToken string `json:"access_token"`
	}
)

func GetToken(request Request, domain string) (string, error) {
	payloadBytes, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s/oauth/token", domain)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenResponse Response
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}
