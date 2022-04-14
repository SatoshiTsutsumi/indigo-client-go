package indigo

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var client *Client
var instanceKey *SSHKey

func TestMain(m *testing.M) {
	var err error
	client, err = NewClient("https://api.customer.jp", os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}
	time.Sleep(time.Second * 12)
	instanceKey, err = client.CreateSSHKey(
		"instanceTestSSHKey",
		sshKeyString1,
	)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}
	exitStatus := m.Run()

	time.Sleep(time.Second * 15)
	err = client.DeleteSSHKey(instanceKey.ID)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}

	os.Exit(exitStatus)
}
