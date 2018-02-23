# GO Language Binding for HPE OneSphere APIs

The GO language binding package for GO developers to call HPE OneSphere APIs

## Prerequisites

go1.9.2 and above. 
You can install the latest version from:

```
https://golang.org/dl/
```

## Usage

Copy the onesphere folder into your GO project folder.

Example:

```
package main

import (
    osbinding "./onesphere"
    "fmt"
)

func main() {
    osbinding.Connect("https://onesphere-host-url", "username", "password")
    fmt.Println("Token:", osbinding.Token)

    fmt.Println("Status:", osbinding.GetStatus())
    fmt.Println("Session:", osbinding.GetSession("full"))
    fmt.Println("Account:", osbinding.GetAccount("full"))
    fmt.Println("ProviderTypes:", osbinding.GetProviderTypes())
    fmt.Println("ZoneTypes:", osbinding.GetZoneTypes())
    fmt.Println("ServiceTypes:", osbinding.GetServiceTypes())
    fmt.Println("Roles:", osbinding.GetRoles())
    fmt.Println("Users:", osbinding.GetUsers())
    fmt.Println("TagKeys:", osbinding.GetTagKeys("full"))
    fmt.Println("Tags:", osbinding.GetTags("full"))

    osbinding.Disconnect()
}
```

## APIs

All the APIs return data in JSON format the same as those returned from HPE OneSphere composable APIs.

### Not Implemented Yet

Some APIs are not yet implemented.
