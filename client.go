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
	/*
		fmt.Printf("----------\n")
		fmt.Printf("%s\n", body)
		fmt.Printf("----------\n")
	*/
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

func requestWithJson[Request, Response any](c *Client, method string, url string, request *Request, response *Response) (*Response, error) {
	var nr *strings.Reader
	if request != nil {
		rb, err := json.Marshal(request)
		if err != nil {
			return nil, err
		}
		nr = strings.NewReader(string(rb))
	}

	req, err := http.NewRequest(method, url, nr)
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
