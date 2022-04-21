package indigo

import (
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient("https://api.customer.jp", os.Getenv("API_KEY"), os.Getenv("API_SECRET"), false)
	if c == nil {
		t.Fatalf("NewClient() = %v, want %v (%v)", c, "'not nil'", err)
	}
}
