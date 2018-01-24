# GO Language Binding for HPE OneSphere APIs

The GO language binding package for GO developers to call HPE OneSphere APIs

## Prerequisites

go1.9.2 and above. 
You can install the latest version from:

```
https://golang.org/dl/
```

## Usage

Copy the ncsmodule into your GO project folder.

Example:

```
package main

import (
    ncsmodule "./ncsmodule"
    "fmt"
)

func main() {
    ncsmodule.Connect("https://onespheretme1.stackbeta.hpe.com", "peng.liu@hpe.com", "Passw0rd!")
    fmt.Println("Token:", ncsmodule.Token)

    fmt.Println("Status:", ncsmodule.GetStatus())
    //fmt.Println("ConnectApp:", ncsmodule.GetConnectApp("windows"))
    fmt.Println("Session:", ncsmodule.GetSession("full"))
    fmt.Println("ProviderTypes:", ncsmodule.GetProviderTypes())
    fmt.Println("ZoneTypes:", ncsmodule.GetZoneTypes())
    fmt.Println("Roles:", ncsmodule.GetRoles())
    fmt.Println("Users:", ncsmodule.GetUsers())
}
```

## APIs

All the APIs return data in JSON format the same as those returned from HPE OneSphere composable APIs.

### Not Implemented Yet

Some APIs are not yet implemented.
