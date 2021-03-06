package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"encoding/json"

	"github.com/kazu69/hosts_file_manager"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

var (
	version string
)

func main() {

	hfm, err := hfm.NewHosts()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app := cli.NewApp()
	app.Name = "hfm - Hostds File Maneger"
	app.Version = version
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name: "kazu69",
		},
	}
	app.Usage = "hosts file management"
	app.UsageText = `
	add (a)    - hfm add <IP> <HOSTS...> [--format json]
	remove (r) - hfm remove <IP> [--format json]
	update (u) - hfm update <IP> <HOSTS...> [--format jsos 
	list (l)   - hfm list [--format json] `
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a hosts record to hosts file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "format, f",
					Usage: "output format type",
				},
			},
			Action: func(c *cli.Context) error {
				ip := c.Args().Get(0)
				hosts := c.Args()[1:]
				record, err := hfm.Add(ip, hosts...)

				if err != nil {
					return err
				}

				format := c.String("format")

				if format == "json" {
					json := ToJSON(record)
					fmt.Println(json)
				} else {
					cyan := chalk.Cyan.NewStyle()
					fmt.Printf("%sAdded %s %s\n", cyan, record.IP, strings.Join(record.Hosts, " "))
				}

				return nil
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"r"},
			Usage:   "remove a hosts record to hosts file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "format, f",
					Usage: "output format type",
				},
			},
			Action: func(c *cli.Context) error {
				ip := c.Args().Get(0)
				record, err := hfm.Remove(ip)

				if err != nil {
					return err
				}

				format := c.String("format")

				if format == "json" {
					json := ToJSON(record)
					fmt.Println(json)
				} else {
					red := chalk.Red.NewStyle()
					fmt.Printf("%sRemoved %s %s\n", red, record.IP, strings.Join(record.Hosts, " "))
				}

				return nil
			},
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "update a hosts record to hosts file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "format, f",
					Usage: "output format type",
				},
			},
			Action: func(c *cli.Context) error {

				ip := c.Args().Get(0)
				hosts := c.Args()[1:]
				record, err := hfm.Update(ip, hosts...)

				if err != nil {
					return err
				}

				format := c.String("format")

				if format == "json" {
					json := ToJSON(record)
					fmt.Println(json)
				} else {
					green := chalk.Green.NewStyle()
					fmt.Printf("%sUpdated %s %s\n", green, record.IP, strings.Join(record.Hosts, " "))
				}

				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "lits hosts records to hosts file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "format, f",
					Usage: "output format type",
				},
			},
			Action: func(c *cli.Context) error {
				records := hfm.List()

				if err != nil {
					return err
				}

				format := c.String("format")

				if format == "json" {
					json := ToJSON(records)
					fmt.Println(json)
				} else {
					for _, r := range records {
						fmt.Printf("%s %s\n", r.IP, strings.Join(r.Hosts, " "))
					}
				}

				return nil
			},
		},
	}

	app.Run(os.Args)
}

func ToJSON(records interface{}) string {
	b, _ := json.Marshal(records)
	return string(b)
}
