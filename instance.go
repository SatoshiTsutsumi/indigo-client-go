package indigo

import (
	"fmt"
	"reflect"
	"strconv"
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
	ID                 int         `json:"id"`
	InstanceName       string      `json:"instance_name"`
	InstanceTypeID     int         `json:"instance_type"`
	SetNo              int         `json:"set_no"`
	VPSKindIf          interface{} `json:"vps_kind"` // NOTE: For Unmarshal as API returns string when creating else number!!
	VPSKind            int         `json:"-"`
	SequenceID         int         `json:"sequence_id"`
	UserID             int         `json:"user_id"`
	ServiceID          string      `json:"service_id"`
	Status             string      `json:"status"`
	SSHKeyID           int         `json:"sshkey_id"`
	SnapshotID         int         `json:"snapshot_id"`
	CreatedAtIf        interface{} `json:"created_at"` // NOTE: For Unmarshal as API returns Date when creating else string!!
	CreatedAt          string      `json:"-"`
	StartDateIf        interface{} `json:"start_date"` // NOTE: For Unmarshal as API returns Date when creating else string!!
	StartDate          string      `json:"-"`
	HostID             int         `json:"host_id"`
	Plan               string      `json:"plan"`
	PlanID             int         `json:"plan_id"`
	DiskPoint          int         `json:"disk_point"`
	MemSize            int         `json:"memsize"`
	CPUs               int         `json:"cpus"`
	OSID               int         `json:"os_id"`
	OtherStatus        int         `json:"otherstatus"`
	UUID               string      `json:"uuid"`
	UIDGID             int         `json:"uidgid"`
	VNCPort            int         `json:"vnc_port"`
	VNCPasswd          string      `json:"vnc_passwd"`
	ARPAName           string      `json:"arpaname"`
	ARPADate           int         `json:"arpadate"`
	StartedAtIf        interface{} `json:"started_at,omitempty"`
	StartedAt          string      `json:"-"`
	ClosedAtIf         interface{} `json:"closed_at,omitempty"`
	ClosedAt           string      `json:"-"`
	StatusChangeDateIf interface{} `json:"status_change_date"` // NOTE: For Unmarshal as API returns Date when creating else string!!
	StatusChangeDate   string      `json:"-"`
	UpdatedAtIf        interface{} `json:"updated_at"`
	UpdatedAt          string      `json:"-"`
	VMRevert           int         `json:"vm_revert"`
	IPAddress          string      `json:"ipaddress,omitempty"`
	MACAddress         string      `json:"macaddress,omitempty"`
	ImportInstance     int         `json:"import_instance"`
	ContainerID        int         `json:"container_id,omitempty"`
	DaemonStatus       string      `json:"daemonstatus"`
	OutOfStock         int         `json:"outofstock"`
	IPAddressType      string      `json:"ipaddress_type"`
	VEID               string      `json:"VEID,omitempty"`
	OS                 *OS         `json:"os,omitempty"`
	IP                 string      `json:"ip,omitempty"`
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

	return fixInstanceStruct(res.Instance), nil
}

func (c *Client) CreateInstanceSync(sshKeyID int, regionID int, osID int, plan int, name string) (*Instance, error) {
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

	time.Sleep(time.Minute * 4)
	return c.GetInstance(res.Instance.ID)
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

	return fixInstanceStruct(res.Instance), nil
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

	return fixInstanceStruct(res.Instance), nil
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

	return fixInstanceStruct(res.Instance), nil
}

func (c *Client) GetInstanceList() ([]*Instance, error) {
	res := &[]*Instance{}
	res, err := requestWithJson[any](c, "GET", fmt.Sprintf("%s/%s", c.hostURL, PathGetInstanceList), nil, res)
	if err != nil {
		return nil, err
	}

	for _, v := range *res {
		fixInstanceStruct(v)
	}
	return *res, nil
}

func (c *Client) GetInstance(id int) (*Instance, error) {
	instances, err := c.GetInstanceList()
	if err != nil {
		return nil, err
	}

	for _, instance := range instances {
		if instance.ID == id {
			return instance, nil
		}
	}
	return nil, fmt.Errorf("instance (%d) not found", id)
}

func (c *Client) UpdateInstanceStatus(instanceID int, status string) error {
	req := &InstanceStatusUpdateRequest{
		InstanceID: instanceID,
		Status:     status,
	}
	res := &InstanceStatusResponse{}
	_, err := requestWithJson(c, "POST", fmt.Sprintf("%s/%s", c.hostURL, PathInstanceStatusUpdate), req, res)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteInstance(instanceID int) error {
	return c.UpdateInstanceStatus(instanceID, "destroy")
}

func fixInstanceStruct(instance *Instance) *Instance {

	if reflect.TypeOf(instance.VPSKindIf).String() == "string" {
		instance.VPSKind, _ = strconv.Atoi(instance.VPSKindIf.(string))
	} else {
		instance.VPSKind = int(instance.VPSKindIf.(float64))
	}
	instance.CreatedAt = convDate(instance.CreatedAtIf)
	instance.StartDate = convDate(instance.StartDateIf)
	instance.StatusChangeDate = convDate(instance.StatusChangeDateIf)
	instance.StartedAt = convDate(instance.StartedAtIf)
	instance.ClosedAt = convDate(instance.ClosedAtIf)
	instance.UpdatedAt = convDate(instance.UpdatedAtIf)

	return instance
}
