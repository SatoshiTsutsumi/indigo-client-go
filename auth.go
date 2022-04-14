package indigo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type AccessTokenRequest struct {
	GrantType    string `json:"grantType"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Code         string `json:"code"`
}

type AccessToken struct {
	Token     string `json:"accessToken"`
	TokenType string `json:"tokenType"`
	ExpiresIn string `json:"expiresIn"`
	Scope     string `json:"scope"`
	IssuedAt  string `json:"issuedAt"`
}

type APIKey struct {
	APIKey    string `json:"apiKey"`
	APISecret string `json:"apiSecret"`
}

func (c *Client) GenerateAccessToken(apiKey, apiSecret string) (*AccessToken, error) {
	payload := &AccessTokenRequest{
		GrantType:    "client_credentials",
		ClientID:     apiKey,
		ClientSecret: apiSecret,
		Code:         "",
	}
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", c.HostURL, PathAccessToken), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	res := AccessToken{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
