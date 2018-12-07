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

type Zone struct {
	Created     time.Time `json:"created"`
	ID          string    `json:"id"`
	Metrics     []*Metric `json:"metrics"`
	Modified    time.Time `json:"modified"`
	Name        string    `json:"name"`
	ProviderURI string    `json:"providerUri"`
	RegionURI   string    `json:"regionUri"`
	Error       *Error `json:"error"`
	Status       string `json:"status"`
	State        string `json:"state"`
	CurrentTasks []struct {
		TaskName   string `json:"taskName"`
		TaskState  string `json:"taskState"`
		TaskStatus string `json:"taskStatus"`
	} `json:"currentTasks"`
	Clusters []struct {
		Created  time.Time `json:"created"`
		ID       string    `json:"id"`
		Name     string    `json:"name"`
		Modified time.Time `json:"modified"`
		Status   string    `json:"status"`
		State    string    `json:"state"`
		Error    *Error `json:"error"`
		Hosts []struct {
			ID            string   `json:"id"`
			Name          string   `json:"name"`
			OsName        string   `json:"osName"`
			OsVersion     string   `json:"osVersion"`
			TotalMemoryGb int      `json:"totalMemoryGb"`
			FreeMemoryGb  int      `json:"freeMemoryGb"`
			TotalCPUGhz   int      `json:"totalCpuGhz"`
			FreeCPUGhz    int      `json:"freeCpuGhz"`
			Datastores    []string `json:"datastores"`
		} `json:"hosts"`
		Datastores []struct {
			Created  time.Time `json:"created"`
			ID       string    `json:"id"`
			Name     string    `json:"name"`
			SizeGiB  int       `json:"sizeGiB"`
			Type     string    `json:"type"`
			Modified time.Time `json:"modified"`
			Status   string    `json:"status"`
			State    string    `json:"state"`
			Error    *Error `json:"error"`
		} `json:"datastores"`
	} `json:"clusters"`
	InTransitClusters []struct {
		Created  time.Time `json:"created"`
		ID       string    `json:"id"`
		Name     string    `json:"name"`
		Modified time.Time `json:"modified"`
		Status   string    `json:"status"`
		State    string    `json:"state"`
		Error    *Error `json:"error"`
		Hosts []struct {
			ID            string   `json:"id"`
			Name          string   `json:"name"`
			OsName        string   `json:"osName"`
			OsVersion     string   `json:"osVersion"`
			TotalMemoryGb int      `json:"totalMemoryGb"`
			FreeMemoryGb  int      `json:"freeMemoryGb"`
			TotalCPUGhz   int      `json:"totalCpuGhz"`
			FreeCPUGhz    int      `json:"freeCpuGhz"`
			Datastores    []string `json:"datastores"`
		} `json:"hosts"`
		Datastores []struct {
			Created  time.Time `json:"created"`
			ID       string    `json:"id"`
			Name     string    `json:"name"`
			SizeGiB  int       `json:"sizeGiB"`
			Type     string    `json:"type"`
			Modified time.Time `json:"modified"`
			Status   string    `json:"status"`
			State    string    `json:"state"`
			Error    *Error `json:"error"`
		} `json:"datastores"`
	} `json:"inTransitClusters"`
	EsxLcmTask struct {
		URI             string   `json:"uri"`
		Name            string   `json:"name"`
		Type            string   `json:"type"`
		UUID            string   `json:"uuid"`
		ParentID        string   `json:"parentId"`
		ChildTasks      []string `json:"childTasks"`
		PercentComplete int      `json:"percentComplete"`
		ProgressUpdates []struct {
			StatusUpdate string    `json:"StatusUpdate"`
			TimeStamp    time.Time `json:"TimeStamp"`
		} `json:"ProgressUpdates"`
		AssociatedResourceInstanceURI  string `json:"associatedResourceInstanceUri"`
		AssociatedResourceInstanceID   string `json:"associatedResourceInstanceId"`
		AssociatedResourceInstanceType string `json:"associatedResourceInstanceType"`
		State                          string `json:"state"`
		Status                         string `json:"status"`
		Error                          *Error `json:"error"`
		TaskFailed bool   `json:"taskFailed"`
		Created    string `json:"created"`
		Modified   string `json:"modified"`
	} `json:"esxLcmTask"`
	NetworkSettings struct {
		NcsManagementNetwork string   `json:"ncsManagementNetwork"`
		EsxManagementNetwork string   `json:"esxManagementNetwork"`
		StorageNetwork       string   `json:"storageNetwork"`
		VMotionNetwork       string   `json:"vMotionNetwork"`
		ProductionNetwork    []string `json:"productionNetwork"`
		PhysicalNetworks     []struct {
			Name        string `json:"name"`
			NetworkType string `json:"networkType"`
		} `json:"physicalNetworks"`
	} `json:"networkSettings"`
	VcenterSettings struct {
		IPAddress string `json:"ipAddress"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		Port      int    `json:"port"`
	} `json:"vcenterSettings"`
	Managed         bool       `json:"managed"`
	URI             string     `json:"uri"`
	ZoneTypeURI     string     `json:"zoneTypeUri"`
	ProjectUris     []string   `json:"projectUris"`
	Projects        []*Project `json:"projects"`
	ResourceProfile struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"resourceProfile"`
	Default      bool   `json:"default"`
	ApplianceURI string `json:"applianceUri"`
	KvmServers   []struct {
		ServerURI  string    `json:"serverUri"`
		Name       string    `json:"name"`
		Created    time.Time `json:"created"`
		Modified   time.Time `json:"modified"`
		Status     string    `json:"status"`
		State      string    `json:"state"`
		Datastores []struct {
			Created  time.Time `json:"created"`
			ID       string    `json:"id"`
			Name     string    `json:"name"`
			SizeGiB  int       `json:"sizeGiB"`
			Type     string    `json:"type"`
			Modified time.Time `json:"modified"`
			Status   string    `json:"status"`
			State    string    `json:"state"`
			Error    *Error `json:"error"`
		} `json:"datastores"`
		Host struct {
			ID            string   `json:"id"`
			Name          string   `json:"name"`
			OsName        string   `json:"osName"`
			OsVersion     string   `json:"osVersion"`
			TotalMemoryGb int      `json:"totalMemoryGb"`
			FreeMemoryGb  int      `json:"freeMemoryGb"`
			TotalCPUGhz   int      `json:"totalCpuGhz"`
			FreeCPUGhz    int      `json:"freeCpuGhz"`
			Datastores    []string `json:"datastores"`
		} `json:"host"`
		Roles []string `json:"roles"`
		Error *Error `json:"error"`
	} `json:"kvmServers"`
	InTransitKvmServers []struct {
		ServerURI  string    `json:"serverUri"`
		Name       string    `json:"name"`
		Created    time.Time `json:"created"`
		Modified   time.Time `json:"modified"`
		Status     string    `json:"status"`
		State      string    `json:"state"`
		Datastores []struct {
			Created  time.Time `json:"created"`
			ID       string    `json:"id"`
			Name     string    `json:"name"`
			SizeGiB  int       `json:"sizeGiB"`
			Type     string    `json:"type"`
			Modified time.Time `json:"modified"`
			Status   string    `json:"status"`
			State    string    `json:"state"`
			Error    *Error `json:"error"`
		} `json:"datastores"`
		Host struct {
			ID            string   `json:"id"`
			Name          string   `json:"name"`
			OsName        string   `json:"osName"`
			OsVersion     string   `json:"osVersion"`
			TotalMemoryGb int      `json:"totalMemoryGb"`
			FreeMemoryGb  int      `json:"freeMemoryGb"`
			TotalCPUGhz   int      `json:"totalCpuGhz"`
			FreeCPUGhz    int      `json:"freeCpuGhz"`
			Datastores    []string `json:"datastores"`
		} `json:"host"`
		Roles []string `json:"roles"`
		Error *Error `json:"error"`
	} `json:"inTransitKvmServers"`
}

type ZoneList struct {
	Total       int           `json:"total"`
	Members     []Zone  `json:"members"`
}

/* GetZones with optional query, and filters by regionUri, providerUri, applianceUri
leave query and filter blank to get all zones

query supports equality comparison against one or more properties using a
"name EQ value" syntax. Multiple comparisons can be combined
using a "name1 EQ value1 AND name2 EQ value2" syntax.

example query: "providerUri EQ /rest/providers/xxxx"

example view: "full"
*/
func (c *Client) GetZones(query, regionUri, providerUri, applianceUri, view string) (ZoneList, error) {
	var (
		uri         = "/rest/zones"
		queryParams = createQuery(&map[string]string{
			query: "query",
			regionUri: "regionUri",
			providerUri: "providerUri",
			applianceUri: "applianceUri",
			view: "view",
		})
		zones ZoneList
	)

	response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)

	if err != nil {
		return zones, err
	}

	if err := json.Unmarshal([]byte(response), &zones); err != nil {
		return zones, apiResponseError(response, err)
	}

	return zones, nil
}
