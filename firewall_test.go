package indigo

import (
	"reflect"
	"testing"
)

var firewallIDForTest int

func TestCreateFirewall(t *testing.T) {
	var err error
	inbound := []*Rule{
		{
			Type:     "HTTP",
			Protocol: "TCP",
			Port:     "80",
			Source:   "0.0.0.0",
		},
	}
	outbound := []*Rule{
		{
			Type:     "HTTP",
			Protocol: "TCP",
			Port:     "80",
			Source:   "0.0.0.0",
		},
	}
	instanceIDs := []int{}
	firewallIDForTest, err = client.CreateFirewall("Test", inbound, outbound, instanceIDs)
	if err != nil {
		t.Fatalf("CreateFirewall() = %v, want %v", err, "'nil'")
	}
	if firewallIDForTest < 0 {
		t.Fatalf("CreateFirewall() = %d, want >= 0", firewallIDForTest)
	}
}

func TestGetFirewallList(t *testing.T) {
	firewalls, err := client.GetFirewallList()
	if err != nil {
		t.Fatalf("GetFirewallList() = %v, want %v", err, "'nil'")
	}
	if len(firewalls) == 0 {
		t.Fatalf("GetFirewallList() = %v, want %v", firewalls, "not empty")
	}
}

func TestRetrieveFirewall(t *testing.T) {
	firewall, err := client.RetrieveFirewall(firewallIDForTest)
	if err != nil {
		t.Fatalf("RetrieveFirewall() = %v, want %v", err, "'nil'")
	}
	if firewall.ID != firewallIDForTest {
		t.Fatalf("RetrieveFirewall() = %d, want %d", firewall.ID, firewallIDForTest)
	}
}

func TestUpdateFirewall(t *testing.T) {
	instanceIDs := []int{}
	inbound := []*Rule{
		{
			Type:     "Custom",
			Protocol: "TCP",
			Port:     "8080",
			Source:   "0.0.0.0",
		},
		{
			Type:     "HTTP",
			Protocol: "TCP",
			Port:     "80",
			Source:   "0.0.0.0",
		},
	}
	outbound := []*Rule{
		{
			Type:     "Custom",
			Protocol: "TCP",
			Port:     "8080",
			Source:   "0.0.0.0",
		},
	}
	err := client.UpdateFirewall(firewallIDForTest, "Updated", inbound, outbound, instanceIDs)
	if err != nil {
		t.Fatalf("UpdateFirewall() = %v, want %v", err, "'nil'")
	}

	firewall, err := client.RetrieveFirewall(firewallIDForTest)
	if err != nil {
		t.Fatalf("RetrieveFirewall() = %v, want %v", err, "'nil'")
	}

	expected := Firewall{
		ID:       firewall.ID,
		Name:     "Updated",
		Inbound:  inbound,
		Outbound: outbound,
	}
	if reflect.DeepEqual(firewall, expected) {
		t.Fatalf("UpdateFirewall() = %v, want %v", firewall, expected)
	}
}

func TestAssignFirewall(t *testing.T) {
	instance, err := client.CreateInstance(sshKeyForTest.ID, 1, 1, 1, "TempInstance")
	if err != nil {
		t.Fatalf("Failed to create instance for TestAssignFirewall(): %v", err)
	}
	err = client.AssignFirewall(firewallIDForTest, instance.ID)
	if err != nil {
		t.Fatalf("AssignFirewall() = %v, want %v", err, "'not nil'")
	}
	err = client.DeleteInstance(instance.ID)
	if err != nil {
		t.Fatalf("Failed to delete instance for TestAssignFirewall(): %v", err)
	}
}

func TestDeleteFirewall(t *testing.T) {
	err := client.DeleteFirewall(firewallIDForTest)
	if err != nil {
		t.Fatalf("DeleteFirewall() = %v, want %v", err, "nil")
	}
}
