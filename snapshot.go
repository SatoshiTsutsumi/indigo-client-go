package indigo

import (
	"fmt"
	"strconv"
	"time"
)

type SnapshotRequest struct {
	Name       string `json:"name"`
	InstanceID int    `json:"instanceid"`
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
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	ServiceID  string      `json:"service_id"`
	UserIDIf   interface{} `json:"user_id"`
	UserID     int         `json:"-"`
	DiskID     int         `json:"disk_id"`
	Volume     int         `json:"volume"`
	SlotNumber int         `json:"slot_number"`
	Status     string      `json:"status"`
	Size       string      `json:"size"`
	Deleted    int         `json:"deleted"`
	CreatedAt  string      `json:"completed_timestamp"`
	DeletedAt  string      `json:"deleted_timestamp"`
}

func (c *Client) CreateSnapshot(name string, instanceID int) error {
	req := &SnapshotRequest{
		Name:       name,
		InstanceID: instanceID,
		SlotNum:    0,
	}
	res := &SnapshotResultResponse{}
	_, err := requestWithJson(c, "POST", fmt.Sprintf("%s/%s", c.hostURL, PathCreateSnapshot), req, res)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) CreateSnapshotSync(name string, instanceID int) (*Snapshot, error) {
	err := c.CreateSnapshot(name, instanceID)
	if err != nil {
		return nil, err
	}

	return c.WaitForSnapshotReady(instanceID, time.Now())
}

func (c *Client) GetSnapshotList(instanceID int) ([]*Snapshot, error) {
	res := &[]*Snapshot{}
	res, err := requestWithJson[any](c, "GET", fmt.Sprintf("%s/%s/%d", c.hostURL, PathSnapshotList, instanceID), nil, res)
	if err != nil {
		return nil, err
	}

	for _, snapshot := range *res {
		fixSnapshotStruct(snapshot)
	}

	return *res, err
}

func (c *Client) RecreateSnapshot(instanceID, snapshotID int) error {
	req := &SnapshotIDRequest{
		InstanceID: instanceID,
		SnapshotID: snapshotID,
	}
	res := &SnapshotResultResponse{}
	_, err := requestWithJson(c, "POST", fmt.Sprintf("%s/%s", c.hostURL, PathRecreateSnapshot), req, res)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) RecreateSnapshotSync(instanceID, snapshotID int) (*Snapshot, error) {
	err := c.RecreateSnapshot(instanceID, snapshotID)
	if err != nil {
		return nil, err
	}

	return c.WaitForSnapshotReady(instanceID, time.Now())
}

func (c *Client) RestoreSnapshot(instanceID, snapshotID int) error {
	req := &SnapshotIDRequest{
		InstanceID: instanceID,
		SnapshotID: snapshotID,
	}
	res := &SnapshotResultResponse{}
	_, err := requestWithJson(c, "POST", fmt.Sprintf("%s/%s", c.hostURL, PathRestoreSnapshot), req, res)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteSnapshot(id int) error {
	res := &SnapshotResultResponse{}
	_, err := requestWithJson[any](c, "DELETE", fmt.Sprintf("%s/%s/%d", c.hostURL, PathDeleteSnapshot, id), nil, res)
	if err != nil {
		return err
	}

	return nil
}

// NOTE:
// WebARENA Indigo API does NOT return snapshot ID on creation and needs to find it in Snapshot List.
// So set timestamp in Id temporarily and match with nearest one in the list.
func (c *Client) FindSnapshot(instanceID int, calledAt time.Time) (*Snapshot, error) {
	snapshots, err := c.GetSnapshotList(instanceID)
	if err != nil {
		return nil, err
	}

	minDiffSec := 10.0 * 3600 // 10 min.
	minDiffIndex := -1

	for i, snapshot := range snapshots {
		createdAt, _ := time.Parse(time.RFC3339, snapshot.CreatedAt)
		fmt.Printf("DEBUG createdAt %v", createdAt)
		diff := createdAt.Sub(calledAt)
		if diff.Seconds() < minDiffSec {
			minDiffSec = diff.Seconds()
			minDiffIndex = i
		}
	}

	if minDiffIndex < 0 {
		return nil, fmt.Errorf("snapshot created at %v not found", calledAt)
	}

	return snapshots[minDiffIndex], nil
}

func (c *Client) GetSnapshot(instanceID, id int) (*Snapshot, error) {
	snapshots, err := c.GetSnapshotList(instanceID)
	if err != nil {
		return nil, err
	}

	for _, snapshot := range snapshots {
		if snapshot.ID == id {
			return snapshot, nil
		}
	}

	return nil, fmt.Errorf("snapshot (%d) not found", id)
}

func (c *Client) WaitForSnapshotReady(instanceID int, calledAt time.Time) (*Snapshot, error) {
	time.Sleep(CheckIntervalForSnapshotCreation)
	snapshot, err := c.FindSnapshot(instanceID, calledAt)
	if err == nil && snapshot.Status == "created" {
		return snapshot, nil
	}

	t := time.Now()
	timeout := t.Add(TimeoutForSnapshotCreation)
	for timeout.Sub(t).Seconds() > 0 {
		time.Sleep(CheckIntervalForSnapshotCreation)
		snapshot, err := c.GetSnapshot(instanceID, snapshot.ID)
		if err == nil && snapshot.Status == "created" {
			return snapshot, nil
		}
		t = time.Now()
	}
	return nil, fmt.Errorf("WaitForSnapshotReady timeout")
}

func fixSnapshotStruct(snapshot *Snapshot) *Snapshot {
	snapshot.UserID, _ = strconv.Atoi(snapshot.UserIDIf.(string))
	snapshot.CreatedAt = convDate(snapshot.CreatedAt)
	snapshot.DeletedAt = convDate(snapshot.DeletedAt)
	return snapshot
}
