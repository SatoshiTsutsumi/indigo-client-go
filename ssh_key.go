package indigo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type SSHKeyRequest struct {
	SSHName string `json:"sshName"`
	SSHKey  string `json:"sshKey"`
}

type SSHKeyUpdateRequest struct {
	SSHName      string `json:"sshName"`
	SSHKey       string `json:"sshKey"`
	SSHKeyStatus string `json:"sshKeyStatus"`
}

type SSHKeyResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	SSHKey  SSHKey `json:"sshKey"`
}

type SSHKeyRetrieveResponse struct {
	Success bool     `json:"success"`
	SSHKeys []SSHKey `json:"sshKey"`
}

type SSHKeyListResponse struct {
	Success bool     `json:"success"`
	Total   int      `json:"total"`
	SSHKeys []SSHKey `json:"sshkeys"`
}

type SSHKey struct {
	Name      string `json:"name"`
	Key       string `json:"sshkey"`
	Status    string `json:"status"`
	UserID    int    `json:"user_id"`
	ServiceID string `json:"service_id"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
	ID        int    `json:"id"`
}

func (c *Client) CreateSSHKey(sshName, sshKey string) (*SSHKey, error) {
	payload := &SSHKeyRequest{
		SSHName: sshName,
		SSHKey:  sshKey,
	}
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", c.HostURL, SSHKeyURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	res := SSHKeyResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res.SSHKey, nil
}

func (c *Client) GetSSHKeyList() ([]SSHKey, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.HostURL, SSHKeyURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	res := SSHKeyListResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return res.SSHKeys, nil
}

func (c *Client) GetActiveSSHKeyList() ([]SSHKey, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.HostURL, ActiveSSHKeyURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	res := SSHKeyListResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return res.SSHKeys, nil
}

func (c *Client) RetrieveSSHKey(sshKeyID int) (*SSHKey, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%d", c.HostURL, SSHKeyURL, sshKeyID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	res := SSHKeyRetrieveResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	if len(res.SSHKeys) != 1 {
		return nil, fmt.Errorf("Invalid response: %v", res.SSHKeys)
	}

	return &res.SSHKeys[0], nil
}

func (c *Client) UpdateSSHKey(sshKeyID int, sshName, sshKey, sshKeyStatus string) error {
	payload := &SSHKeyUpdateRequest{
		SSHName:      sshName,
		SSHKey:       sshKey,
		SSHKeyStatus: sshKeyStatus,
	}
	rb, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s/%d", c.HostURL, SSHKeyURL, sshKeyID), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DestroySSHKey(sshKeyID int) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s/%d", c.HostURL, SSHKeyURL, sshKeyID), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
