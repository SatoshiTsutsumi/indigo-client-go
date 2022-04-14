package indigo

import (
	"os"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	time.Sleep(time.Second * 8)
	c, err := NewClient("https://api.customer.jp", os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	if c == nil {
		t.Fatalf("NewClient() = %v, want %v (%v)", c, "'not nil'", err)
	}
}
