package indigo

import (
	"fmt"
)

type InstanceRequest struct {
	SSHKeyID     int    `json:"sshKeyId,omitempty"`
	WinPassword  string `json:"winPassword,omitempty"`
	ImportURL    string `json:"importUrl,omitempty"`
	SnapshotID   int    `json:"snapshotId,omitempty"`
	RegionID     int    `json:"regionId"`
	OSID         int    `json:"osId"`
	InstancePlan int    `json:"instancePlan"`
	InstanceName string `json:"instanceName"`
}

type InstanceStatusUpdateRequest struct {
	InstanceID int    `json:"instanceId"`
	Status     string `json:"status"`
}

type RegionListResponse struct {
	Success bool      `json:"success"`
	Total   int       `json:"total"`
	Regions []*Region `json:"regionlist"`
}

type InstanceTypeListResponse struct {
	Success       bool            `json:"success"`
	Total         int             `json:"total"`
	InstanceTypes []*InstanceType `json:"instanceTypes"`
}

type OSListResponse struct {
	Success      bool          `json:"success"`
	Total        int           `json:"total"`
	OSCategories []*OSCategory `json:"osCategory"`
}

type InstanceSpecListResponse struct {
	Success bool            `json:"success"`
	Total   int             `json:"total"`
	Specs   []*InstanceSpec `json:"speclist"`
}

type InstanceResponse struct {
	Success  bool      `json:"success"`
	Message  string    `json:"message"`
	Instance *Instance `json:"vms"`
}

type InstanceStatusResponse struct {
	Success        bool   `json:"success"`
	Message        string `json:"message"`
	SuccessCode    string `json:"successCode"`
	InstanceStatus string `json:"instanceStatus"`
}

type Region struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	OEMID           int    `json:"oem_id"`
	UsePossibleDate string `json:"use_possible_date"`
}

type InstanceType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type OSCategory struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Logo   string  `json:"logo"`
	OSType *OSType `json:"osLists"`
}

type OSType struct {
	ID             int    `json:"id"`
	CategoryID     int    `json:"categoryid"`
	Name           string `json:"name"`
	ViewName       string `json:"viewname"`
	InstanceTypeID int    `json:"instancetype_id"`
}

type InstanceSpec struct {
	ID              int           `json:"id"`
	Code            string        `json:"code"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	UsePossibleDate string        `json:"use_possible_date"`
	InstanceTypeID  int           `json:"instancetype_id"`
	IPAddressType   string        `json:"ipaddress_type"`
	CreatedAt       string        `json:"created_at"`
	UpdatedAt       string        `json:"updated_at"`
	InstanceType    *InstanceType `json:"instance_type"`
	KVMResources    *KVMResources `json:"kvm_resources"`
}

type KVMResources struct {
	ID       int    `json:"id"`
	PlanID   int    `json:"plan_id"`
	Name     string `json:"name"`
	Param    string `json:"param"`
	LimitNum int    `json:"limitnum"`
}

type OS struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ViewName string `json:"viewname"`
}

// FIXME
type Instance struct {
	ID               int    `json:"id"`
	InstanceName     string `json:"instance_name"`
	InstanceTypeID   int    `json:"instance_type"`
	SetNo            int    `json:"set_no"`
	VPSKind          string `json:"vps_kind"`
	SequenceID       int    `json:"sequence_id"`
	UserID           int    `json:"user_id"`
	ServiceID        string `json:"service_id"`
	Status           string `json:"status"`
	SSHKeyID         int    `json:"sshkey_id"`
	SnapshotID       int    `json:"snapshot_id"`
	CreatedAt        *Date  `json:"created_at"`
	StartDate        *Date  `json:"start_date"`
	HostID           int    `json:"host_id"`
	Plan             string `json:"plan"`
	PlanID           int    `json:"plan_id"`
	DiskPoint        int    `json:"disk_point"`
	MemSize          int    `json:"memsize"`
	CPUs             int    `json:"cpus"`
	OSID             int    `json:"os_id"`
	OtherStatus      int    `json:"otherstatus"`
	UUID             string `json:"uuid"`
	UIDGID           int    `json:"uidgid"`
	VNCPort          int    `json:"vnc_port"`
	VNCPasswd        string `json:"vnc_passwd"`
	ARPAName         string `json:"arpaname"`
	ARPADate         int    `json:"arpadate"`
	StartedAt        *Date  `json:"started_at,omitempty"`
	ClosedAt         *Date  `json:"closed_at,omitempty"`
	StatusChangeDate *Date  `json:"status_change_date"`
	UpdatedAt        *Date  `json:"updated_at"`
	VMRevert         int    `json:"vm_revert"`
	IPAddress        string `json:"ipaddress,omitempty"`
	MACAddress       string `json:"macaddress,omitempty"`
	ImportInstance   int    `json:"import_instance"`
	ContainerID      int    `json:"container_id,omitempty"`
	DaemonStatus     string `json:"daemonstatus"`
	OutOfStock       int    `json:"outofstock"`
	IPAddressType    string `json:"ipaddress_type"`
	VEID             string `json:"VEID,omitempty"`
	OS               *OS    `json:"os,omitempty"`
	IP               string `json:"ip,omitempty"`
}

func (c *Client) GetRegionList() ([]*Region, error) {
	res := &RegionListResponse{}
	res, err := requestWithJson[any](c, "GET", fmt.Sprintf("%s/%s", c.hostURL, PathGetRegion), nil, res)
	if err != nil {
		return nil, err
	}

	return res.Regions, nil
}

func (c *Client) GetInstanceTypeList() ([]*InstanceType, error) {
	res := &InstanceTypeListResponse{}
	res, err := requestWithJson[any](c, "GET", fmt.Sprintf("%s/%s", c.hostURL, PathInstanceTypeList), nil, res)
	if err != nil {
		return nil, err
	}

	return res.InstanceTypes, nil
}

func (c *Client) GetOSList() ([]*OSCategory, error) {
	res := &OSListResponse{}
	res, err := requestWithJson[any](c, "GET", fmt.Sprintf("%s/%s", c.hostURL, PathOSList), nil, res)
	if err != nil {
		return nil, err
	}

	return res.OSCategories, nil
}

func (c *Client) GetInstanceSpecList() ([]*InstanceSpec, error) {
	res := &InstanceSpecListResponse{}
	res, err := requestWithJson[any](c, "GET", fmt.Sprintf("%s/%s", c.hostURL, PathGetInstanceSpecList), nil, res)
	if err != nil {
		return nil, err
	}

	return res.Specs, nil
}

func (c *Client) CreateInstance(sshKeyID int, regionID int, osID int, plan int, name string) (*Instance, error) {
	req := &InstanceRequest{
		SSHKeyID:     sshKeyID,
		RegionID:     regionID,
		OSID:         osID,
		InstancePlan: plan,
		InstanceName: name,
	}
	res := &InstanceResponse{}
	res, err := requestWithJson(c, "POST", fmt.Sprintf("%s/%s", c.hostURL, PathCreateInstance), req, res)
	if err != nil {
		return nil, err
	}

	return res.Instance, nil
}

func (c *Client) CreateWindowsInstance(winPassword string, regionID int, osID int, plan int, name string) (*Instance, error) {
	req := &InstanceRequest{
		WinPassword:  winPassword,
		RegionID:     regionID,
		OSID:         osID,
		InstancePlan: plan,
		InstanceName: name,
	}
	res := &InstanceResponse{}
	res, err := requestWithJson(c, "POST", fmt.Sprintf("%s/%s", c.hostURL, PathCreateInstance), req, res)
	if err != nil {
		return nil, err
	}

	return res.Instance, nil
}

func (c *Client) CreateImportInstance(url string, regionID int, osID int, plan int, name string) (*Instance, error) {
	req := &InstanceRequest{
		ImportURL:    url,
		RegionID:     regionID,
		OSID:         osID,
		InstancePlan: plan,
		InstanceName: name,
	}
	res := &InstanceResponse{}
	res, err := requestWithJson(c, "POST", fmt.Sprintf("%s/%s", c.hostURL, PathCreateInstance), req, res)
	if err != nil {
		return nil, err
	}

	return res.Instance, nil
}

func (c *Client) CreateSnapshotInstance(sshKeyID int, snapshotID int, plan int, name string) (*Instance, error) {
	req := &InstanceRequest{
		SSHKeyID:     sshKeyID,
		SnapshotID:   snapshotID,
		InstancePlan: plan,
		InstanceName: name,
	}
	res := &InstanceResponse{}
	res, err := requestWithJson(c, "POST", fmt.Sprintf("%s/%s", c.hostURL, PathCreateInstance), req, res)
	if err != nil {
		return nil, err
	}

	return res.Instance, nil
}

func (c *Client) GetInstanceList() ([]*Instance, error) {
	res := &[]*Instance{}
	res, err := requestWithJson[any](c, "GET", fmt.Sprintf("%s/%s", c.hostURL, PathGetInstanceList), nil, res)
	if err != nil {
		return nil, err
	}

	return *res, nil
}

func (c *Client) UpdateInstanceStatus(instanceID int, status string) error {
	req := &InstanceStatusUpdateRequest{
		InstanceID: instanceID,
		Status:     status,
	}
	res := &InstanceStatusResponse{}
	_, err := requestWithJson(c, "POST", fmt.Sprintf("%s/%s", c.hostURL, PathGetInstanceList), req, res)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteInstance(instanceID int) error {
	return c.UpdateInstanceStatus(instanceID, "destroy")
}
