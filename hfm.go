package hfm

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type Record struct {
	Hosts []string
	IP    string
	Line  int
	Raw   string
}

type Records []Record

type Hosts struct {
	records Records
	path    string
}

// Return new `Hosts` instance
func NewHosts() (Hosts, error) {
	
	hosts := Hosts{path: HostsPath}

	_, err := hosts.Readfile()

	if err != nil {
		return hosts, err
	}

	return hosts, nil
}

// Parse hosts file
func (h *Hosts) Readfile() (Records, error) {
	file, err := os.Open(h.path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	records := Records{}
	scanner := bufio.NewScanner(file)
	line := 0

	for scanner.Scan() {
		line++

		record := Record{
			Line: line,
			Raw:  scanner.Text(),
		}

		fields := strings.Fields(record.Raw)

		// ignore comment out line `#`
		reg := regexp.MustCompile("^#")
		if len(fields) != 0 && !reg.MatchString(fields[0]) {
			record.IP = fields[0]
			record.Hosts = fields[1:]
		}

		records = append(records, record)
	}

	h.records = records
	return records, nil
}

// Find record from hosts file
func (h *Hosts) Find(ip string) (r Record, exist bool) {
	record := Record{}

	for _, r := range h.records {
		if r.IP == ip {
			record = r
		}
	}

	if record.IP != "" {
		return record, true
	}

	return record, false
}

// Write record to hosts file
func (h *Hosts) Write() error {
	if !h.IsWritable() {
		return fmt.Errorf("Error: %s is not Writables", h.path)
	}

	// ファイルに追加
	file, err := os.OpenFile(h.path, os.O_WRONLY|os.O_APPEND, 0660)

	if err != nil {
		return err
	}

	defer file.Close()

	lines := []string{}
	for _, record := range h.records {
		lines = append(lines, record.Raw)
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(h.path, []byte(output), 0644)

	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

// show all hosts list
func (h *Hosts) List() Records {
	records := Records{}
	for _, r := range h.records {
		if r.IP != "" {
			record := Record{
				IP:    r.IP,
				Hosts: r.Hosts[:],
				Raw:   r.Raw,
				Line:  r.Line,
			}

			records = append(records, record)
		}
	}

	return records
}

// Add Record to hosts
func (h *Hosts) Add(ip string, host ...string) (Record, error) {
	record := Record{}

	_, exist := h.Find(ip)

	if exist == false {
		hosts := strings.Join(host[:], " ")
		raw := strings.Join([]string{ip, hosts}[:], " ")

		record = Record{
			IP:    ip,
			Hosts: host,
			Line:  len(h.records) + 1,
			Raw:   raw,
		}

		// レコード登録
		h.records = append(h.records, record)
	}

	// ファイルに書き込み
	err := h.Write()

	if err != nil {
		return Record{}, err
	}

	return record, nil
}

// Remove record from hosts
func (h *Hosts) Remove(ip string) (Record, error) {
	record := Record{}

	_, exist := h.Find(ip)

	if !exist {
		return record, fmt.Errorf("Error: %s is not Exist", ip)
	}

	index := 0
	for i, r := range h.records {
		if r.IP == ip {
			index = i
		}

		if i > index {
			r.Line++
		}
	}

	// 該当のレコードを除いた新しいsliceを作成する
	record = h.records[index]
	h.records = append(h.records[:index], h.records[index+1:]...)

	// ファイルに書き込み
	err := h.Write()

	if err != nil {
		log.Fatalln(err)
	}

	return record, nil
}

// Update record
func (h *Hosts) Update(ip string, host ...string) (Record, error) {
	record := Record{}

	_, exist := h.Find(ip)

	if !exist {
		return record, fmt.Errorf("Error: %s is not Exist", ip)
	}

	hosts := strings.Join(host[:], " ")
	raw := strings.Join([]string{ip, hosts}[:], " ")

	for i, r := range h.records {
		if r.IP == ip {

			record = Record{
				IP:    r.IP,
				Hosts: host[:],
				Raw:   raw,
				Line:  r.Line,
			}

			h.records[i] = record
		}
	}

	// ファイルに書き込み
	err := h.Write()

	if err != nil {
		log.Fatalln(err)
	}

	return record, nil
}
