package indigo

import (
	"fmt"
	"testing"
)

func TestGetRegionList(t *testing.T) {
	regions, err := client.GetRegionList()
	if err != nil || len(regions) == 0 {
		t.Errorf("GetRegionList() = %v, want %v", err, "'not nil'")
	}
}

func TestGetInstanceTypeList(t *testing.T) {
	types, err := client.GetInstanceTypeList()
	if err != nil {
		t.Fatalf("GetInstanceTypeList() = %v, want %v", err, "'nil'")
	}
	fmt.Printf("%v", types)
}

func TestGetOSList(t *testing.T) {
	oss, err := client.GetOSList()
	if err != nil {
		t.Fatalf("GetOSList() = %v, want %v", err, "'nil'")
	}
	fmt.Printf("%v", oss)
}

func TestGetInstanceSpecList(t *testing.T) {
	oss, err := client.GetInstanceSpecList()
	if err != nil {
		t.Fatalf("GetInstanceSpecList() = %v, want %v", err, "'nil'")
	}
	for i, os := range oss {
		fmt.Printf("%d %v\n", i, os)
	}
}

var instanceForTestInstance *Instance

func TestCreateInstance(t *testing.T) {
	var err error
	instanceForTestInstance, err = client.CreateInstance(sshKeyForTest.ID, 1, 1, 1, "VM00")
	if err != nil {
		t.Fatalf("CreateInstance() = %v, want %v", err, "'nil'")
	}
	fmt.Printf("%v", instanceForTestInstance)
}

func TestGetInstanceList(t *testing.T) {
	instances, err := client.GetInstanceList()
	if err != nil {
		t.Fatalf("GetInstanceList() = %v, want %v", err, "'nil'")
	}
	fmt.Printf("%v", instances)
}

func TestUpdateInstanceStatus(t *testing.T) {
	err := client.UpdateInstanceStatus(instanceForTestInstance.ID, "destroy")
	if err != nil {
		t.Fatalf("UpdateInstanceStatus() = %v, want %v", err, "'nil'")
	}
}
