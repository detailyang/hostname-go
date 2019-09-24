// Package hostname returns the host name reported by kernel.
package hostname

import "os"

var h string

func init() {
	var err error
	h, err = os.Hostname()
	if err != nil {
		h = "unknown hostname"
	}
}

// Get returns the host name reported by the kernel.
func Get() string {
	return h
}
