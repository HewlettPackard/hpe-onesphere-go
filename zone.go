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
	"fmt"
	"github.com/HewlettPackard/hpe-onesphere-go/rest"
	"time"
)

type ZoneRequest struct {
	Name            string `json:"name"`
	ProviderURI     string `json:"providerUri"`
	RegionURI       string `json:"regionUri"`
	ZoneTypeURI     string `json:"zoneTypeUri"`
	ApplianceURI    string `json:"applianceUri"`
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
	ResourceProfile struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"resourceProfile"`
	ResourceCapacity int `json:"resourceCapacity"`
	Rates            []struct {
		MetricName string `json:"metricName"`
		RateValue  int    `json:"rateValue"`
	} `json:"rates"`
}

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
	Clusters []*Cluster `json:"clusters"`
	InTransitClusters []*Cluster `json:"inTransitClusters"`
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
	KvmServers   []*KvmServer `json:"kvmServers"`
	InTransitKvmServers []KvmServer `json:"inTransitKvmServers"`
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

// GetZoneByID Retrieve Zone by ID
func (c *Client) GetZoneByID(id string) (Zone, error) {
	var (
		uri        = "/rest/zones/" + id
		zone Zone
	)

	if id == "" {
		return zone, fmt.Errorf("id must not be empty")
	}

	response, err := c.RestAPICall(rest.GET, uri, nil, nil)

	if err != nil {
		return zone, err
	}

	if err := json.Unmarshal([]byte(response), &zone); err != nil {
		return zone, apiResponseError(response, err)
	}

	return zone, err
}

// GetZoneApplianceImage Retrieve Zone Appliance Image URI by Zone.ID
func (c *Client) GetZoneApplianceImage(id string) (string, error) {
	var (
		uri        = "/rest/zones/" + id + "/appliance-image"
		applianceImageURI string
	)

	if id == "" {
		return applianceImageURI, fmt.Errorf("id must not be empty")
	}

	applianceImageURI, err := c.RestAPICall(rest.GET, uri, nil, nil)

	return applianceImageURI, err
}

// GetZoneTaskStatus Retrieve Zone Appliance Image URI by Zone.ID
func (c *Client) GetZoneTaskStatus(id string) (string, error) {
	var (
		uri        = "/rest/zones/" + id + "/task-status"
		taskStatus string
	)

	if id == "" {
		return taskStatus, fmt.Errorf("id must not be empty")
	}

	taskStatus, err := c.RestAPICall(rest.GET, uri, nil, nil)

	return taskStatus, err
}

// GetZoneConnections with optional uuid filter
// leave uuid blank to get all connections
func (c *Client) GetZoneConnections(id, uuid string) (ConnectionList, error) {
	var (
		uri         = "/rest/zones/" + id + "/connections"
		queryParams = createQuery(&map[string]string{
			uuid: "uuid",
		})
		connections ConnectionList
	)

	response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)

	if err != nil {
		return connections, err
	}

	if err := json.Unmarshal([]byte(response), &connections); err != nil {
		return connections, apiResponseError(response, err)
	}

	return connections, nil
}

// CreateZone Creates Zone and returns updated zone
func (c *Client) CreateZone(zoneRequest ZoneRequest) (Zone, error) {
	var (
		uri        = "/rest/zones/"
		zone Zone
	)

	response, err := c.RestAPICall(rest.POST, uri, nil, zoneRequest)

	if err != nil {
		return zone, err
	}

	if err := json.Unmarshal([]byte(response), &zone); err != nil {
		return zone, apiResponseError(response, err)
	}

	return zone, err
}
