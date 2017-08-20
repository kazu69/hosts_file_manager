package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
	"github.com/kazu69/hosts_file_manager"
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
	add (a)    - hfm add <IP> <HOSTS...>
	remove (r) - hfm remove <IP>
	update (u) - hfm update <IP> <HOSTS...>
	list (l)   - hfm list`
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a hosts record to hosts file",
			Action: func(c *cli.Context) error {

				ip := c.Args().Get(0)
				hosts := c.Args()[1:]
				record, err := hfm.Add(ip, hosts...)

				if err != nil {
					return err
				}

				cyan := chalk.Cyan.NewStyle()
				fmt.Printf("%sAdded %s %s\n", cyan, record.IP, strings.Join(record.Hosts, " "))

				return nil
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"r"},
			Usage:   "remove a hosts record to hosts file",
			Action: func(c *cli.Context) error {

				ip := c.Args().Get(0)
				record, err := hfm.Remove(ip)

				if err != nil {
					return err
				}

				red := chalk.Red.NewStyle()
				fmt.Printf("%sRemoved %s %s\n", red, record.IP, strings.Join(record.Hosts, " "))
				return nil
			},
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "update a hosts record to hosts file",
			Action: func(c *cli.Context) error {

				ip := c.Args().Get(0)
				hosts := c.Args()[1:]
				record, err := hfm.Update(ip, hosts...)

				if err != nil {
					return err
				}

				green := chalk.Green.NewStyle()
				fmt.Printf("%sUpdated %s %s\n", green, record.IP, strings.Join(record.Hosts, " "))
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "lits hosts records to hosts file",
			Action: func(c *cli.Context) error {
				records := hfm.List()

				if err != nil {
					return err
				}

				for _, r := range records {
					fmt.Printf("%s %s\n", r.IP, strings.Join(r.Hosts, " "))
				}

				return nil
			},
		},
	}

	app.Run(os.Args)
}
