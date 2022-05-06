package indigo

import (
	"flag"
	"fmt"
	"os"
	"testing"
	"time"
)

var client *Client
var sshKeyForTest *SSHKey
var instanceForTest *Instance

func setUp() bool {
	var err error
	fmt.Print("SetUp (requires about 5 min)...\n")
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
	instanceForTest, err = client.CreateInstance(sshKeyForTest.ID, 1, 1, 1, "instanceForTest")
	if err != nil {
		fmt.Printf("%v", err)
		return true
	}
	// Wait for instance created
	time.Sleep(time.Minute * 5)
	fmt.Print("Instance created.\n")
	return false
}

func tearDown() {
	var err error
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
}

func TestMain(m *testing.M) {
	var err error
	skipFixture := flag.Bool("skip", false, "skip fixture")
	flag.Parse()

	client, err = NewClient("https://api.customer.jp", os.Getenv("API_KEY"), os.Getenv("API_SECRET"), true)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}
	fmt.Print("Client created.\n")

	var exitStatus int
	if !*skipFixture {
		if setUp() {
			goto tearDown
		}
	}

	exitStatus = m.Run()

tearDown:
	// tearDown
	if !*skipFixture {
		tearDown()
	}

	os.Exit(exitStatus)
}
