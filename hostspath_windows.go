package hfm

import (
	"os"
)

var HostsPath = os.Getenv("SystemRoot") + `\System32\drivers\etc\hosts`
