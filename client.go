package indigo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	PathAccessToken          = "oauth/v1/accesstokens"
	PathSSHKey               = "webarenaIndigo/v1/vm/sshkey"
	PathActiveSSHKey         = "webarenaIndigo/v1/vm/sshkey/active/status"
	PathGetRegion            = "webarenaIndigo/v1/vm/getregion?instanceTypeId=1"
	PathInstanceTypeList     = "webarenaIndigo/v1/vm/instancetypes"
	PathOSList               = "webarenaIndigo/v1/vm/oslist?instanceTypeId=1"
	PathGetInstanceSpecList  = "webarenaIndigo/v1/vm/getinstancespec?instanceTypeId=1&osId=1"
	PathCreateInstance       = "webarenaIndigo/v1/vm/createinstance"
	PathGetInstanceList      = "webarenaIndigo/v1/vm/getinstancelist"
	PathInstanceStatusUpdate = "webarenaIndigo/v1/vm/instance/statusupdate"
	PathCreateFirewall       = "webarenaIndigo/v1/nw/createfirewall"
	PathGetFirewallList      = "webarenaIndigo/v1/nw/getfirewalllist"
	PathRetrieveFirewall     = "webarenaIndigo/v1/nw/gettemplate"
	PathUpdateFirewall       = "webarenaIndigo/v1/nw/updatefirewall"
	PathAssignFirewall       = "webarenaIndigo/v1/nw/assign"
	PathDeleteFirewall       = "webarenaIndigo/v1/nw/deletefirewall"
	PathCreateSnapshot       = "webarenaIndigo/v1/disk/takesnapshot"
	PathRecreateSnapshot     = "webarenaIndigo/v1/disk/retakesnapshot"
	PathRestoreSnapshot      = "webarenaIndigo/v1/disk/restoresnapshot"
	PathDeleteSnapshot       = "webarenaIndigo/v1/disk/deletesnapshot"
	PathSnapshotList         = "webarenaIndigo/v1/disk/snapshotlist"
	PathAddDomain            = "webarenaIndigo/v1/dns/registerdomain"
	PathGetDomainList        = "webarenaIndigo/v1/dns/getdomainlist"
	PathGetDomain            = "webarenaIndigo/v1/dns/getdomainrecord"
	PathDeleteDomain         = "webarenaIndigo/v1/dns/canceldomain"
)

type Client struct {
	hostURL       string
	httpClient    *http.Client
	apiKey        string
	apiSecret     string
	accessToken   *AccessToken
	AutoRateLimit bool
}

type Date struct {
	Date         string `json:"date"`
	TimeZoneType int    `json:"timezone_type"`
	TimeZone     string `json:"timezone"`
}

func NewClient(host, apiKey, apiSecret string, autoRateLimit bool) (*Client, error) {
	if host == "" || apiKey == "" || apiSecret == "" {
		return nil, fmt.Errorf("error NewClient(): Invalid parameter(host, apiKey, apiSecret required)")
	}

	c := Client{
		hostURL:       host,
		httpClient:    &http.Client{Timeout: 30 * time.Second},
		apiKey:        apiKey,
		apiSecret:     apiSecret,
		AutoRateLimit: autoRateLimit,
	}

	err := c.RefreshAccessToken()
	if err != nil {
		return nil, err
	}
  
	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	if c.accessToken != nil {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken.Token))
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

func requestWithJsonNoRefresh[Request, Response any](c *Client, method string, url string, request *Request, response *Response) (*Response, error) {
	var req *http.Request
	var err error

	if request != nil {
		var rb []byte
		rb, err = json.Marshal(request)
		if err != nil {
			return nil, err
		}
		nr := strings.NewReader(string(rb))
		req, err = http.NewRequest(method, url, nr)
		req.Header.Add("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	if response != nil {
		err = json.Unmarshal(body, response)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func requestWithJson[Request, Response any](c *Client, method string, url string, request *Request, response *Response) (*Response, error) {
	// NOTE: Avoiding "Too many requests error" or "Spike arrest violation error".
	//   API server often returns the errors even if the client calls an API with more than 10s interval!
	//   API server says
	//     "Allowed rate : MessageRate{messagesPerPeriod=2, periodInMicroseconds=1000000, maxBurstMessageCount=1.0}",
	//     which means 2 messages per 1s is allowed??
	//   It seems using refreshed AccessToken alleviates this issue even though requests increase.
	if c.AutoRateLimit {
		time.Sleep(time.Second * 12)

		err := c.RefreshAccessToken()
		if err != nil {
			return nil, err
		}
		time.Sleep(time.Second * 6)
	}

	return requestWithJsonNoRefresh(c, method, url, request, response)
}
