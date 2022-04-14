package indigo

import (
	"fmt"
)

type SSHKeyRequest struct {
	Name      string `json:"sshName"`
	Key       string `json:"sshKey"`
	KeyStatus string `json:"sshKeyStatus,omitempty"`
}

type SSHKeyResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message,omitempty"`
	Key     *SSHKey `json:"sshKey"`
}

type SSHKeyListResponse struct {
	Success bool      `json:"success"`
	Total   int       `json:"total"`
	Keys    []*SSHKey `json:"sshkeys"`
}

type SSHKeyRetrieveResponse struct {
	Success bool      `json:"success"`
	Keys    []*SSHKey `json:"sshKey"`
}

type SSHKeyBoolResponse struct {
	Success bool   `json:"-"`
	Message string `json:"-"`
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
	req := &SSHKeyRequest{
		Name: sshName,
		Key:  sshKey,
	}
	res := &SSHKeyResponse{}
	res, err := requestWithJson(c, "POST", fmt.Sprintf("%s/%s", c.hostURL, PathSSHKey), req, res)
	if err != nil {
		return nil, err
	}

	return res.Key, nil
}

func (c *Client) GetSSHKeyList() ([]*SSHKey, error) {
	res := &SSHKeyListResponse{}
	res, err := requestWithJson[any](c, "GET", fmt.Sprintf("%s/%s", c.hostURL, PathSSHKey), nil, res)
	if err != nil {
		return nil, err
	}

	return res.Keys, nil
}

func (c *Client) GetActiveSSHKeyList() ([]*SSHKey, error) {
	res := &SSHKeyListResponse{}
	res, err := requestWithJson[any](c, "GET", fmt.Sprintf("%s/%s", c.hostURL, PathActiveSSHKey), nil, res)
	if err != nil {
		return nil, err
	}

	return res.Keys, nil
}

func (c *Client) RetrieveSSHKey(sshKeyID int) (*SSHKey, error) {
	res := &SSHKeyRetrieveResponse{}
	res, err := requestWithJson[any](c, "GET", fmt.Sprintf("%s/%s/%d", c.hostURL, PathSSHKey, sshKeyID), nil, res)
	if err != nil {
		return nil, err
	}
	if len(res.Keys) != 1 {
		return nil, fmt.Errorf("Invalid response: %v", res)
	}

	return res.Keys[0], nil
}

func (c *Client) UpdateSSHKey(sshKeyID int, sshName, sshKey, sshKeyStatus string) error {
	req := &SSHKeyRequest{
		Name:      sshName,
		Key:       sshKey,
		KeyStatus: sshKeyStatus,
	}
	res := &SSHKeyBoolResponse{}
	res, err := requestWithJson(c, "PUT", fmt.Sprintf("%s/%s/%d", c.hostURL, PathSSHKey, sshKeyID), req, res)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteSSHKey(sshKeyID int) error {
	res := &SSHKeyBoolResponse{}
	res, err := requestWithJson[any](c, "DELETE", fmt.Sprintf("%s/%s/%d", c.hostURL, PathSSHKey, sshKeyID), nil, res)
	if err != nil {
		return err
	}

	return nil
}
