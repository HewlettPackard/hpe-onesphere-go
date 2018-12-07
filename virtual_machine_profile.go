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

import (
	"encoding/json"
	"github.com/HewlettPackard/hpe-onesphere-go/rest"
	"time"
)

type VirtualMachineProfileList struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	URI          string    `json:"uri"`
	Description  string    `json:"description"`
	RegionURI    string    `json:"regionUri"`
	ZoneURI      string    `json:"zoneUri"`
	CPUCount     int       `json:"cpuCount"`
	MemorySizeGB int       `json:"memorySizeGB"`
	DiskSizeGB   int       `json:"diskSizeGB"`
	Created      time.Time `json:"created"`
	Modified     time.Time `json:"modified"`
}

// GetVirtualMachineProfiles returns VirtualMachineProfileList with optional query for zoneUri and serviceUri
// leave query blank to get all virtualMachineProfiles
// example query for serviceUri: "serviceUri EQ /rest/services/2F8bbc7abe-2ae1-a366-a4dd-f065618063a6"
// example query for zoneUri: "zoneUri EQ /rest/zones/b1d0b94b-b3e2-459f-95ed-a1b4d4645338"
// example query for both: "serviceUri EQ /rest/services/2F8bbc7abe-2ae1-a366-a4dd-f065618063a6 AND zoneUri EQ /rest/zones/b1d0b94b-b3e2-459f-95ed-a1b4d4645338"
func (c *Client) GetVirtualMachineProfiles(query string) (VirtualMachineProfileList, error) {
	var (
		uri      = "/rest/virtual-machine-profiles"
		queryParams = createQuery(&map[string]string{
			"query":     query,
		})
		virtualMachineProfiles VirtualMachineProfileList
	)

	response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)

	if err != nil {
		return virtualMachineProfiles, err
	}

	if err := json.Unmarshal([]byte(response), &virtualMachineProfiles); err != nil {
		return virtualMachineProfiles, apiResponseError(response, err)
	}

	return virtualMachineProfiles, err
}

// GetVirtualMachineProfilesByServiceURI returns VirtualMachineProfileList by serviceUri
// example: client.GetVirtualMachineProfilesByServiceURI("/rest/services/2F8bbc7abe-2ae1-a366-a4dd-f065618063a6")
func (c *Client) GetVirtualMachineProfilesByServiceURI(serviceURI string) (VirtualMachineProfileList, error) {
	return c.GetVirtualMachineProfiles("serviceUri EQ "+serviceURI)
}
