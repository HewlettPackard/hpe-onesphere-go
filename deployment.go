// (C) Copyright 2018 Hewlett Packard Enterprise Development LP.
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL
// THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR
// OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package onesphere

type Deployment struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	ZoneUri string `json:"zoneUri"`
	Zone    struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Uri  string `json:"uri"`
	} `json:"zone"`
	RegionUri  string `json:"regionUri"`
	ServiceUri string `json:"serviceUri"`
	Service    struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Uri  string `json:"uri"`
	} `json:"service"`
	ServiceTypeUri      string `json:"serviceTypeUri"`
	Version             string `json:"version"`
	Status              string `json:"status"`
	State               string `json:"state"`
	Uri                 string `json:"uri"`
	ProjectUri          string `json:"projectUri"`
	DeploymentEndpoints []struct {
		Address     string `json:"address"`
		AddressType string `json:"addressType"`
	} `json:"deploymentEndpoints"`
	AppDeploymentInfo string `json:"appDeploymentInfo"`
	HasConsole        bool   `json:"hasConsole"`
	CloudPlatformId   string `json:"cloudPlatformId"`
	Created           string `json:"created"`
	Modified          string `json:"modified"`
}
