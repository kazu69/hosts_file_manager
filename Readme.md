## hfm

[![Build Status](https://travis-ci.org/kazu69/hosts_file_manager.svg?branch=master)](https://travis-ci.org/kazu69/hosts_file_manager)

> hfm is Golang Hosts File Maneger

### Install

```sh
go get -u github.com/kazu69/hosts_file_manager
```

### CLI

#### list

```sh
hfm list

127.0.0.1 localhost
255.255.255.255 broadcasthost
::1 localhost

hfm list --format json

[{"Hosts":["localhost"],"IP":"127.0.0.1","Line":8,"Raw":"127.0.0.1\tlocalhost"},{"Hosts":["broadcasthost"],"IP":"255.255.255.255","Line":9,"Raw":"255.255.255.255\tbroadcasthost"},{"Hosts":["localhost"],"IP":"::1","Line":10,"Raw":"::1 localhost "}]
```


```

#### Add

```sh
hfm add 0.0.0.0 hoge.com (--format json)

Added 0.0.0.0 hoge.com
```

#### Update

```sh
hfm update 0.0.0.0 hoge.com huga.com (--format json)

Updated 0.0.0.0 hoge.com huga.com
```

#### Remove

```sh
hfm remove 0.0.0.0 (--format json)

Removed 0.0.0.0 hoge.com huga.com
```

### API

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
