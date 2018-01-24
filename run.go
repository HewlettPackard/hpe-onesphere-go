// (C) Copyright 2018 Hewlett Packard Enterprise Development LP.

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

