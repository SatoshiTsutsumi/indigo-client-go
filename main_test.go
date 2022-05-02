package indigo

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var client *Client
var sshKeyForTest *SSHKey
var instanceForTest *Instance

func TestMain(m *testing.M) {
	// SetUp
	fmt.Print("SetUp (requires about 5 min)...\n")
	var err error
	client, err = NewClient("https://api.customer.jp", os.Getenv("API_KEY"), os.Getenv("API_SECRET"), true)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}
	fmt.Print("Client created.\n")

	fmt.Print("Creating SSHKey...\n")
	sshKeyForTest, err = client.CreateSSHKey(
		"instanceTestSSHKey",
		sshKeyString1,
	)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}
	fmt.Print("SSHKey created.\n")

	fmt.Print("Creating Instance...\n")
	var exitStatus int = 0
	instanceForTest, err = client.CreateInstance(sshKeyForTest.ID, 1, 1, 1, "instanceForTest")
	if err != nil {
		fmt.Printf("%v", err)
		goto tearDown
	}
	// Wait for instance created
	time.Sleep(time.Minute * 5)
	fmt.Print("Instance created.\n")

	exitStatus = m.Run()

tearDown:
	// tearDown
	fmt.Print("Delete instance.\n")
	client.DeleteInstance(instanceForTest.ID)
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Print("Delete SSHKey.\n")
	err = client.DeleteSSHKey(sshKeyForTest.ID)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}

	os.Exit(exitStatus)
}
