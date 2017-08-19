## hfm

> hfm is Golang Hosts File Maneger

### API

#### Install

```sh
go get github.com/kazu69/hosts_file_manager
```

#### list

```go
package main

import (
    "fmt"
    "strings"
    "github.com/kazu69/hosts_file_manager"
)

func main() {
    HostsPath = `/etc/hosts`
    hosts := hfm.NewHosts(HostsPath)

    for _, record := range hosts.records {
        host =  strings.Join(record.Hosts[:], " ")
        fmt. Printf("%s %s", record.IP , host)
    }
}
```

#### Host record Add

```go
package main

import (
    "fmt"
    "github.com/kazu69/hosts_file_manager"
)

func main() {
    HostsPath = `/etc/hosts`
    hosts := hfm.NewHosts(HostsPath)
    addedRecord, err := hosts.Add("127.0.0.1", "exampe.com", "exampe.net"))

    if err == nil {
        fmt.Println(addedRecord)
    }
}
```

#### Host record remove

```go
package main

import (
    "fmt"
    "github.com/kazu69/hosts_file_manager"
)

func main() {
    HostsPath = `/etc/hosts`
    hosts := hfm.NewHosts(HostsPath)
    removedRecord, err := hosts.Remove("127.0.0.1")

    if err == nil {
        fmt.Println(removedRecord)
    }
}
```

#### Host record update


```go
package main

import (
    "fmt"
    "github.com/kazu69/hosts_file_manager"
)

func main() {
    HostsPath = `/etc/hosts`
    hosts := hfm.NewHosts(HostsPath)
    updatedRecord, err := hosts.Update("127.0.0.1", "example.com", "example.net")

    if err == nil {
        fmt.Println(updatedRecord)
    }
}
```

#### Find host record

```go
package main

import (
    "fmt"
    "github.com/kazu69/hosts_file_manager"
)

func main() {
    HostsPath = `/etc/hosts`
    hosts := hfm.NewHosts(HostsPath)
    record, exist := hosts.Find("127.0.0.1")

    if exist == true {
        fmt.Println(record)
    }
}
```

### Tests

Unit testing is done according to standard methods

```sh
go test
```

### License

MIT license.