package indigo

import (
	"fmt"
	"os"
	"testing"
)

var client *Client
var key *SSHKey

func TestMain(m *testing.M) {
	var err error
	client, err = NewClient("https://api.customer.jp", os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	if err != nil {
		os.Exit(-1)
	}
	key, err = client.CreateSSHKey(
		"testkey",
		sshKeyString1,
	)
	fmt.Printf("%v", key)
	if err != nil {
		os.Exit(-1)
	}
	exitStatus := m.Run()
	os.Exit(exitStatus)
}
