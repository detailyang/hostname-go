package hostname

import (
	"os"
	"testing"
)

func TestHostname(t *testing.T) {
	h, err := os.Hostname()
	if err != nil {
		t.Fatal(err)
	}

	if Get() != h {
		t.Errorf("Expect %s but got %s", h, Get())
	}
}
