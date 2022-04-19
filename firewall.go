package indigo

import (
	"fmt"
)

type FirewallRequest struct {
	TemplateID  int     `json:"templateid,omitempty"`
	Name        string  `json:"name"`
	Inbound     []*Rule `json:"inbound"`
	Outbound    []*Rule `json:"outbound"`
	InstanceIDs []int   `json:"instances"`
}

type FirewallAssignRequest struct {
	TemplateID int `json:"templateid"`
	InstanceID int `json:"instanceid"`
}

type FirewallResponse struct {
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	SuccessCode string `json:"successCode"`
	ID          int    `json:"firewallId"`
}

type FirewallOperationResponse struct {
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	SuccessCode string `json:"successCode"`
}

type Rule struct {
	Type     string `json:"type"`
	Protocol string `json:"protocol"`
	Port     string `json:"port"`
	Source   string `json:"source"`
}

type FirewallIndex struct {
	ID        int    `json:"id"`
	ServiceID string `json:"service_id"`
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type FirewallEntry struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Direction string `json:"direction"`
	Type      string `json:"type"`
	Protocol  string `json:"protocol"`
	Port      string `json:"port"`
	Source    string `json:"source"`
}

type Firewall struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Inbound  []*Rule `json:"inbound"`
	Outbound []*Rule `json:"outbound"`
}

func (c *Client) CreateFirewall(name string, inbound []*Rule, outbound []*Rule, instanceIDs []int) (int, error) {
	req := &FirewallRequest{
		Name:        name,
		Inbound:     inbound,
		Outbound:    outbound,
		InstanceIDs: instanceIDs,
	}
	res := &FirewallResponse{}
	res, err := requestWithJson(c, "POST", fmt.Sprintf("%s/%s", c.hostURL, PathCreateFirewall), req, res)
	if err != nil {
		return -1, err
	}

	return res.ID, nil
}

func (c *Client) GetFirewallList() ([]*FirewallIndex, error) {

	res := &[]*FirewallIndex{}
	res, err := requestWithJson[any](c, "GET", fmt.Sprintf("%s/%s", c.hostURL, PathGetFirewallList), nil, res)
	if err != nil {
		return nil, err
	}

	return *res, nil
}

func (c *Client) RetrieveFirewall(id int) (*Firewall, error) {
	res := &[]*FirewallEntry{}
	res, err := requestWithJson[any](c, "GET", fmt.Sprintf("%s/%s/%d", c.hostURL, PathRetrieveFirewall, id), nil, res)
	if err != nil {
		return nil, err
	}

	firewall := &Firewall{
		ID: id,
	}

	for _, e := range *res {
		firewall.Name = e.Name
		rule := Rule{
			Type:     e.Type,
			Protocol: e.Protocol,
			Port:     e.Port,
			Source:   e.Source,
		}
		if e.Direction == "in" {
			firewall.Inbound = append(firewall.Inbound, &rule)
		} else {
			firewall.Outbound = append(firewall.Outbound, &rule)
		}
	}
	return firewall, nil
}

func (c *Client) UpdateFirewall(id int, name string, inbound []*Rule, outbound []*Rule, instanceIDs []int) error {
	req := &FirewallRequest{
		TemplateID:  id,
		Name:        name,
		Inbound:     inbound,
		Outbound:    outbound,
		InstanceIDs: instanceIDs,
	}
	res := &FirewallResponse{}
	res, err := requestWithJson(c, "PUT", fmt.Sprintf("%s/%s", c.hostURL, PathUpdateFirewall), req, res)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) AssignFirewall(templateID, instanceID int) error {
	req := &FirewallAssignRequest{
		TemplateID: templateID,
		InstanceID: instanceID,
	}
	res := &FirewallOperationResponse{}
	res, err := requestWithJson(c, "POST", fmt.Sprintf("%s/%s", c.hostURL, PathAssignFirewall), req, res)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteFirewall(id int) error {
	res := &FirewallOperationResponse{}
	res, err := requestWithJson[any](c, "DELETE", fmt.Sprintf("%s/%s/%d", c.hostURL, PathDeleteFirewall, id), nil, res)
	if err != nil {
		return err
	}

	return nil
}
