package indigo

import (
	"fmt"
	"testing"
	"time"
)

func TestGetRegionList(t *testing.T) {
	time.Sleep(time.Second * 10)
	regions, err := client.GetRegionList()
	if err != nil || len(regions) == 0 {
		t.Errorf("GetRegionList() = %v, want %v", err, "'not nil'")
	}
}

func TestGetInstanceTypeList(t *testing.T) {
	time.Sleep(time.Second * 10)
	types, err := client.GetInstanceTypeList()
	if err != nil {
		t.Fatalf("GetInstanceTypeList() = %v, want %v", err, "'nil'")
	}
	fmt.Printf("%v", types)
}

func TestGetOSList(t *testing.T) {
	time.Sleep(time.Second * 10)
	oss, err := client.GetOSList()
	if err != nil {
		t.Fatalf("GetOSList() = %v, want %v", err, "'nil'")
	}
	fmt.Printf("%v", oss)
}

func TestGetInstanceSpecList(t *testing.T) {
	time.Sleep(time.Second * 10)
	oss, err := client.GetInstanceSpecList()
	if err != nil {
		t.Fatalf("GetInstanceSpecList() = %v, want %v", err, "'nil'")
	}
	for i, os := range oss {
		fmt.Printf("%d %v\n", i, os)
	}
}

var instance *NewInstance

func TestCreateInstance(t *testing.T) {
	var err error
	time.Sleep(time.Second * 10)
	instance, err = client.CreateInstance(instanceKey.ID, 1, 4, 1, "VM00")
	if err != nil {
		t.Fatalf("CreateInstance() = %v, want %v", err, "'nil'")
	}
	fmt.Printf("%v", instance)
}

func TestGetInstanceList(t *testing.T) {
	time.Sleep(time.Second * 10)
	instances, err := client.GetInstanceList()
	if err != nil {
		t.Fatalf("GetInstanceList() = %v, want %v", err, "'nil'")
	}
	fmt.Printf("%v", instances)
}

func TestUpdateInstanceStatus(t *testing.T) {
	time.Sleep(time.Second * 10)
	err := client.UpdateInstanceStatus(instance.ID, "destroy")
	if err != nil {
		t.Fatalf("UpdateInstanceStatus() = %v, want %v", err, "'nil'")
	}
}
