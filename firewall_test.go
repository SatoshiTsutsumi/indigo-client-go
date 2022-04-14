package indigo

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateFirewall(t *testing.T) {
	time.Sleep(time.Second * 10)
	inbound := []Rule{Rule{
		Type:     "HTTP",
		Protocol: "TCP",
		Port:     80,
		Source:   "0.0.0.0",
	}}
	outbound := []Rule{Rule{
		Type:     "HTTP",
		Protocol: "TCP",
		Port:     80,
		Source:   "0.0.0.0",
	}}
	instanceIDs := []int{}
	firewallID, err := client.CreateFirewall("Test", inbound, outbound, instanceIDs)
	if err != nil {
		t.Fatalf("CreateFirewall() = %v, want %v", err, "'nil'")
	}
	fmt.Printf("%v", firewallID)
}

func TestGetFirewallList(t *testing.T) {
	time.Sleep(time.Second * 10)
	firewalls, err := client.GetFirewallList()
	if err != nil {
		t.Fatalf("GetFirewallList() = %v, want %v", err, "'nil'")
	}
	fmt.Printf("%v", firewalls)
}
