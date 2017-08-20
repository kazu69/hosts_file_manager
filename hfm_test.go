package hfm

import (
	"os"
	"testing"
)

func TestNewHosts(t *testing.T) {
	dir, _ := os.Getwd()
	HostsPath = dir + "/testdata/hosts"
	hosts, err := NewHosts()

	if err != nil {
		t.Error("Failed to load file.")
	}

	if hosts.path == "" {
		t.Error("Failed to set path.")
	}
}

func TestReadfile(t *testing.T) {
	dir, _ := os.Getwd()
	HostsPath = dir + "/testdata/hosts"
	hosts, _ := NewHosts()
	records, err := hosts.Readfile()

	if err != nil {
		t.Error("Failed to open file.")
	}

	if len(records) != 5 {
		t.Error("Faild read file inaccuracy.")
	}

	record := records[1]

	if record.IP != "127.0.0.1" {
		t.Error("Faild parse hosts file.")
	}

	if record.Hosts[0] != "localhost" {
		t.Error("Faild parse hosts file.")
	}

	if record.Line != 2 {
		t.Error("Faild parse hosts file.")
	}

}

func TestFind(t *testing.T) {
	dir, _ := os.Getwd()
	HostsPath = dir + "/testdata/hosts"
	hosts, _ := NewHosts()

	ip := "192.168.1.1"
	record, exist := hosts.Find(ip)

	if exist == false {
		t.Error("Faild record not exist.")
	}

	if record.IP != ip {
		t.Error("Faild record not exist.")
	}

	ip = "192.168.1.2"
	record, exist = hosts.Find(ip)

	if exist == true {
		t.Error("Faild record exist.")
	}

	if record.IP == ip {
		t.Error("Faild record exist.")
	}
}

func TestWrite(t *testing.T) {
	dir, _ := os.Getwd()
	HostsPath = dir + "/testdata/hosts"
	hosts, _ := NewHosts()

	hosts.Add("123.123.123.123", "hoge.com")

	_, exist := hosts.Find("123.123.123.123")

	if exist == false {
		t.Error("Faild record writed.")
	}

	hosts.Remove("123.123.123.123")
}

func TestList(t *testing.T) {
	dir, _ := os.Getwd()
	HostsPath = dir + "/testdata/hosts"
	hosts, _ := NewHosts()

	records := hosts.List()

	if len(records) != 3 {
		t.Error("Faild records number not correct.")
	}

	record := records[0]

	if record.IP != "127.0.0.1" && record.Hosts[0] != "localhost" {
		t.Error("Faild record parse not correct.")
	}

	record = records[1]

	if record.IP != "::1" && record.Hosts[0] != "localhost" {
		t.Error("Faild record parse not correct.")
	}

	record = records[2]

	if record.IP != "192.168.1.1" && record.Hosts[0] != "example.com" {
		t.Error("Faild record parse not correct.")
	}
}

func TestAdd(t *testing.T) {
	dir, _ := os.Getwd()
	HostsPath = dir + "/testdata/hosts"
	hosts, _ := NewHosts()

	_, err := hosts.Add("invalid", "hoge.com")

	if err == nil {
		t.Error("Faild Add method validation error.")
	}

	added, err := hosts.Add("123.123.123.123", "hoge.com")

	if err != nil {
		t.Error("Faild Add method has error.")
	}

	if added.IP != "123.123.123.123" {
		t.Error("Faild Add method return incorrectly.")
	}

	if added.Hosts[0] != "hoge.com" {
		t.Error("Faild Remove method return incorrectly.")
	}

	_, exist := hosts.Find("123.123.123.123")

	if exist == false {
		t.Error("Faild record not added.")
	}

	hosts.Remove("123.123.123.123")
}

func TestRemove(t *testing.T) {
	dir, _ := os.Getwd()
	HostsPath = dir + "/testdata/hosts"
	hosts, _ := NewHosts()

	hosts.Add("123.123.123.123", "hoge.com")
	removed, err := hosts.Remove("123.123.123.123")

	if err != nil {
		t.Error("Faild Remove method has error.")
	}

	if removed.IP != "123.123.123.123" {
		t.Error("Faild Remove method return incorrectly.")
	}

	if removed.Hosts[0] != "hoge.com" {
		t.Error("Faild Remove method return incorrectly.")
	}

	_, exist := hosts.Find("123.123.123.123")

	if exist == true {
		t.Error("Faild record not removed.")
	}
}

func TestUpdatee(t *testing.T) {
	dir, _ := os.Getwd()
	HostsPath = dir + "/testdata/hosts"
	hosts, _ := NewHosts()

	hosts.Add("123.123.123.123", "hoge.com")
	updated, err := hosts.Update("123.123.123.123", "fuga.com")

	if err != nil {
		t.Error("Faild Update method has error.")
	}

	if updated.Hosts[0] != "fuga.com" {
		t.Error("Faild Update method return incorrectly.")
	}

	hosts.Remove("123.123.123.123")
}
