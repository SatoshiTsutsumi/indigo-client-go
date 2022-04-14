package indigo

import (
	"fmt"
)

type SnapshotRequest struct {
	Name       string `json:"name"`
	InstanceID int    `json:"instanceId"`
	SlotNum    int    `json:"slotnum"`
}

type SnapshotIDRequest struct {
	InstanceID int `json:"instanceid"`
	SnapshotID int `json:"snapshotid"`
}

type SnapshotResultResponse struct {
	Status int `json:"STATUS"`
}

type Snapshot struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ServiceID  string `json:"service_id"`
	UserID     int    `json:"user_id"`
	DiskID     int    `json:"disk_id"`
	Volume     int    `json:"volume"`
	SlotNumber int    `json:"slot_number"`
	Status     string `json:"status"`
	Size       string `json:"size"`
	Deleted    int    `json:"deleted"`
	CreatedAt  string `json:"completed_timestamp"`
	DeletedAt  string `json:"deleted_timestamp"`
}

func (c *Client) CreateSnapshot(name string, instanceID int) error {
	req := &SnapshotRequest{
		Name:       name,
		InstanceID: instanceID,
		SlotNum:    0,
	}
	res := &SnapshotResultResponse{}
	res, err := requestWithJson(c, "POST", fmt.Sprintf("%s/%s", c.hostURL, PathCreateSnapshot), req, res)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetSnapshotList(instanceID int) ([]*Snapshot, error) {
	res := &[]*Snapshot{}
	res, err := requestWithJson[any](c, "GET", fmt.Sprintf("%s/%s/%d", c.hostURL, PathSnapshotList, instanceID), nil, res)
	if err != nil {
		return nil, err
	}

	return *res, err
}

func (c *Client) RecreateSnapshot(instanceID, snapshotID int) error {
	req := &SnapshotIDRequest{
		InstanceID: instanceID,
		SnapshotID: snapshotID,
	}
	res := &SnapshotResultResponse{}
	res, err := requestWithJson(c, "POST", fmt.Sprintf("%s/%s", c.hostURL, PathRecreateSnapshot), req, res)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) RestoreSnapshot(instanceID, snapshotID int) error {
	req := &SnapshotIDRequest{
		InstanceID: instanceID,
		SnapshotID: snapshotID,
	}
	res := &SnapshotResultResponse{}
	res, err := requestWithJson(c, "POST", fmt.Sprintf("%s/%s", c.hostURL, PathRestoreSnapshot), req, res)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteSnapshot(id int) error {
	res := &SnapshotResultResponse{}
	res, err := requestWithJson[any](c, "DELETE", fmt.Sprintf("%s/%s/%d", c.hostURL, PathDeleteSnapshot, id), nil, res)
	if err != nil {
		return err
	}

	return nil
}
