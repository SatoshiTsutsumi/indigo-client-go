package indigo

import (
	"fmt"
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

func (c *Client) RefreshAccessToken() error {
	c.accessToken = nil
	req := &AccessTokenRequest{
		GrantType:    "client_credentials",
		ClientID:     c.apiKey,
		ClientSecret: c.apiSecret,
		Code:         "",
	}
	res := &AccessToken{}
	res, err := requestWithJsonNoRefresh(c, "POST", fmt.Sprintf("%s/%s", c.hostURL, PathAccessToken), req, res)
	if err != nil {
		return err
	}

	c.accessToken = res

	return nil
}
