package indigo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	AccessTokenURL         = "oauth/v1/accesstokens"
	SSHKeyURL              = "webarenaIndigo/v1/vm/sshkey"
	ActiveSSHKeyURL        = "webarenaIndigo/v1/vm/sshkey/active/status"
	GetRegionURL           = "webarenaIndigo/v1/vm/getregion?instanceTypeId=1"
	GetInstanceTypeListURL = "webarenaIndigo/v1/vm/instancetypes"
	GetOSListURL           = "webarenaIndigo/v1/vm/oslist?instanceTypeId=1"
	GetInstanceSpecListURL = "webarenaIndigo/v1/vm/getinstancespec?instanceTypeId=1&osId=1"
	CreateInstanceURL      = "webarenaIndigo/v1/vm/createinstance"
	GetInstanceListURL     = "webarenaIndigo/v1/vm/getinstancelist"
	StatusUpdateURL        = "webarenaIndigo/v1/vm/instance/statusupdate"
)

type Client struct {
	HostURL     string
	HTTPClient  *http.Client
	AccessToken *AccessToken
}

type Date struct {
	Date         string `json:"date"`
	TimeZoneType int    `json:"timezone_type"`
	TimeZone     string `json:"timezone"`
}

func NewClient(host, clientID, clientSecret string) (*Client, error) {
	if host == "" || clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("Invalid parameter")
	}

	c := Client{
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		HostURL:    host,
	}

	token, err := c.GenerateAccessToken(clientID, clientSecret)
	if err != nil {
		return nil, err
	}
	c.AccessToken = token
	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	if c.AccessToken != nil {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken.Token))
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Printf("----------\n")
	fmt.Printf("%s\n", body)
	fmt.Printf("----------\n")

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
