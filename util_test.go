package hfm

import (
	"os"
	"testing"
)

func TestIsWritable(t *testing.T) {
	dir, _ := os.Getwd()
	HostsPath = dir + "/testdata/hosts"
	hosts, _ := NewHosts()

	writable := hosts.IsWritable()

	if writable == false {
		t.Error("Failed to file is not writable.")
	}

	HostsPath = dir + "/testdata/denied"
	hosts, _ = NewHosts()

	writable = hosts.IsWritable()

	if writable == true {
		t.Error("Failed to file is writable.")
	}
}
