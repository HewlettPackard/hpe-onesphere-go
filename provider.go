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
	"strconv"
	"time"
)

type ProviderRequest struct {
	ID                string `json:"id"`
	ProviderTypeURI   string `json:"providerTypeUri"`
	AccessKey         string `json:"accessKey"`
	SecretKey         string `json:"secretKey"`
	PaymentProvider   bool   `json:"paymentProvider"`
	S3CostBucket      string `json:"s3CostBucket"`
	MasterURI         string `json:"masterUri"`
	SubscriptionID    string `json:"subscriptionId"`
	DirectoryURI      string `json:"directoryUri"`
	TenantID          string `json:"tenantId"`
	UniqueName        string `json:"uniqueName"`
	FamilyName        string `json:"familyName"`
	GivenName         string `json:"givenName"`
	BillingAccountURI string `json:"billingAccountUri"`
	State             string `json:"state"`
}

type Provider struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	URI             string `json:"uri"`
	ProviderTypeURI string `json:"providerTypeUri"`
	Status          string `json:"status"`
	State           string `json:"state"`
	AccessKey       string `json:"accessKey"`
	SecretKey       string `json:"secretKey"`
	PaymentProvider bool   `json:"paymentProvider"`
	S3CostBucket    string `json:"s3CostBucket"`
	SubscriptionID  string `json:"subscriptionId"`
	DirectoryURI    string `json:"directoryUri"`
	TenantID        string `json:"tenantId"`
	UniqueName      string `json:"uniqueName"`
	FamilyName      string `json:"familyName"`
	GivenName       string `json:"givenName"`
	ClientID        string `json:"clientID"`
	ClientSecret    string `json:"clientSecret"`
	Children        []struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		URI             string `json:"uri"`
		ProviderTypeURI string `json:"providerTypeUri"`
		Status          string `json:"status"`
		State           string `json:"state"`
	} `json:"children"`
	BillingAccountURI string   `json:"billingAccountUri"`
	ProjectUris       []string `json:"projectUris"`
	Regions           []struct {
		ID      string `json:"id"`
		Metrics []struct {
			ResourceURI string `json:"resourceUri"`
			Resource    struct {
				Value   string `json:"value"`
				Name    string `json:"name"`
				URI     string `json:"uri"`
				Project struct {
					Name string `json:"name"`
					URI  string `json:"uri"`
				} `json:"project"`
				Zone struct {
					Name   string `json:"name"`
					URI    string `json:"uri"`
					Region struct {
						Name     string `json:"name"`
						URI      string `json:"uri"`
						Provider struct {
							Name         string `json:"name"`
							URI          string `json:"uri"`
							ProviderType struct {
								Name string `json:"name"`
								URI  string `json:"uri"`
							} `json:"providerType"`
						} `json:"provider"`
					} `json:"region"`
				} `json:"zone"`
			} `json:"resource"`
			Name        string `json:"name"`
			Units       string `json:"units"`
			Description string `json:"description"`
			Values      []struct {
				Value int       `json:"value"`
				Start time.Time `json:"start"`
				End   time.Time `json:"end"`
			} `json:"values"`
			Total        int `json:"total"`
			Start        int `json:"start"`
			Count        int `json:"count"`
			Associations []struct {
				Category string `json:"category"`
				Name     string `json:"name"`
				URI      string `json:"uri"`
			} `json:"associations"`
		} `json:"metrics"`
		Created  time.Time `json:"created"`
		Modified time.Time `json:"modified"`
		Name     string    `json:"name"`
		Location struct {
			Latitude  int `json:"latitude"`
			Longitude int `json:"longitude"`
		} `json:"location"`
		ProviderURI string `json:"providerUri"`
		Zones       []struct {
			Created time.Time `json:"created"`
			ID      string    `json:"id"`
			Metrics []struct {
				ResourceURI string `json:"resourceUri"`
				Resource    struct {
					Value   string `json:"value"`
					Name    string `json:"name"`
					URI     string `json:"uri"`
					Project struct {
						Name string `json:"name"`
						URI  string `json:"uri"`
					} `json:"project"`
					Zone struct {
						Name   string `json:"name"`
						URI    string `json:"uri"`
						Region struct {
							Name     string `json:"name"`
							URI      string `json:"uri"`
							Provider struct {
								Name         string `json:"name"`
								URI          string `json:"uri"`
								ProviderType struct {
									Name string `json:"name"`
									URI  string `json:"uri"`
								} `json:"providerType"`
							} `json:"provider"`
						} `json:"region"`
					} `json:"zone"`
				} `json:"resource"`
				Name        string `json:"name"`
				Units       string `json:"units"`
				Description string `json:"description"`
				Values      []struct {
					Value int       `json:"value"`
					Start time.Time `json:"start"`
					End   time.Time `json:"end"`
				} `json:"values"`
				Total        int `json:"total"`
				Start        int `json:"start"`
				Count        int `json:"count"`
				Associations []struct {
					Category string `json:"category"`
					Name     string `json:"name"`
					URI      string `json:"uri"`
				} `json:"associations"`
			} `json:"metrics"`
			Modified    time.Time `json:"modified"`
			Name        string    `json:"name"`
			ProviderURI string    `json:"providerUri"`
			RegionURI   string    `json:"regionUri"`
			Error       struct {
				Message            string   `json:"message"`
				Details            string   `json:"details"`
				RecommendedActions []string `json:"recommendedActions"`
				NestedErrors       string   `json:"nestedErrors"`
				ErrorSource        string   `json:"errorSource"`
				ErrorCode          string   `json:"errorCode"`
				Data               string   `json:"data"`
				CanForce           bool     `json:"canForce"`
			} `json:"error"`
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
				Error    struct {
					Message            string   `json:"message"`
					Details            string   `json:"details"`
					RecommendedActions []string `json:"recommendedActions"`
					NestedErrors       string   `json:"nestedErrors"`
					ErrorSource        string   `json:"errorSource"`
					ErrorCode          string   `json:"errorCode"`
					Data               string   `json:"data"`
					CanForce           bool     `json:"canForce"`
				} `json:"error"`
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
					Error    struct {
						Message            string   `json:"message"`
						Details            string   `json:"details"`
						RecommendedActions []string `json:"recommendedActions"`
						NestedErrors       string   `json:"nestedErrors"`
						ErrorSource        string   `json:"errorSource"`
						ErrorCode          string   `json:"errorCode"`
						Data               string   `json:"data"`
						CanForce           bool     `json:"canForce"`
					} `json:"error"`
				} `json:"datastores"`
			} `json:"clusters"`
			InTransitClusters []struct {
				Created  time.Time `json:"created"`
				ID       string    `json:"id"`
				Name     string    `json:"name"`
				Modified time.Time `json:"modified"`
				Status   string    `json:"status"`
				State    string    `json:"state"`
				Error    struct {
					Message            string   `json:"message"`
					Details            string   `json:"details"`
					RecommendedActions []string `json:"recommendedActions"`
					NestedErrors       string   `json:"nestedErrors"`
					ErrorSource        string   `json:"errorSource"`
					ErrorCode          string   `json:"errorCode"`
					Data               string   `json:"data"`
					CanForce           bool     `json:"canForce"`
				} `json:"error"`
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
					Error    struct {
						Message            string   `json:"message"`
						Details            string   `json:"details"`
						RecommendedActions []string `json:"recommendedActions"`
						NestedErrors       string   `json:"nestedErrors"`
						ErrorSource        string   `json:"errorSource"`
						ErrorCode          string   `json:"errorCode"`
						Data               string   `json:"data"`
						CanForce           bool     `json:"canForce"`
					} `json:"error"`
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
				Error                          struct {
					Message            string   `json:"message"`
					Details            string   `json:"details"`
					RecommendedActions []string `json:"recommendedActions"`
					NestedErrors       string   `json:"nestedErrors"`
					ErrorSource        string   `json:"errorSource"`
					ErrorCode          string   `json:"errorCode"`
					Data               string   `json:"data"`
					CanForce           bool     `json:"canForce"`
				} `json:"error"`
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
			Managed     bool     `json:"managed"`
			URI         string   `json:"uri"`
			ZoneTypeURI string   `json:"zoneTypeUri"`
			ProjectUris []string `json:"projectUris"`
			Projects    []struct {
				Created     time.Time `json:"created"`
				Deployments struct {
					Members []struct {
						ActiveUsers []struct {
							Email   string `json:"email"`
							Name    string `json:"name"`
							Role    string `json:"role"`
							IsLocal string `json:"isLocal"`
							URI     string `json:"uri"`
						} `json:"activeUsers"`
						ClusterURI string    `json:"clusterUri"`
						Created    time.Time `json:"created"`
						Endpoints  []struct {
							Name        string `json:"name"`
							Address     string `json:"address"`
							AddressType string `json:"addressType"`
						} `json:"endpoints"`
						HasConsole               bool      `json:"hasConsole"`
						ID                       string    `json:"id"`
						Modified                 time.Time `json:"modified"`
						Name                     string    `json:"name"`
						ZoneURI                  string    `json:"zoneUri"`
						ServiceURI               string    `json:"serviceUri"`
						RegionURI                string    `json:"regionUri"`
						VirtualMachineProfileURI string    `json:"virtualMachineProfileUri"`
						Service                  struct {
							ID      string `json:"id"`
							Name    string `json:"name"`
							URI     string `json:"uri"`
							Version string `json:"version"`
							Icon    string `json:"icon"`
						} `json:"service"`
						Region struct {
							ID   string `json:"id"`
							Name string `json:"name"`
							URI  string `json:"uri"`
						} `json:"region"`
						Status       string   `json:"status"`
						State        string   `json:"state"`
						MemorySizeGB int      `json:"memorySizeGB"`
						CPUCount     int      `json:"cpuCount"`
						CPUGhz       int      `json:"cpuGhz"`
						DiskSizeGB   int      `json:"diskSizeGB"`
						URI          string   `json:"uri"`
						ProjectURI   string   `json:"projectUri"`
						VolumeURIs   []string `json:"volumeURIs"`
						Firewall     []struct {
							AllowedIPs string `json:"allowedIPs"`
							Ports      []int  `json:"ports"`
						} `json:"firewall"`
						Volumes []struct {
							ID      string `json:"id"`
							Name    string `json:"name"`
							URI     string `json:"uri"`
							Status  string `json:"status"`
							SizeGiB int    `json:"sizeGiB"`
						} `json:"volumes"`
						ErrorMessage string `json:"errorMessage"`
					} `json:"members"`
					Total int `json:"total"`
				} `json:"deployments"`
				ID        string    `json:"id"`
				Modified  time.Time `json:"modified"`
				Name      string    `json:"name"`
				Protected bool      `json:"protected"`
				URI       string    `json:"uri"`
				TagUris   []string  `json:"tagUris"`
			} `json:"projects"`
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
					Error    struct {
						Message            string   `json:"message"`
						Details            string   `json:"details"`
						RecommendedActions []string `json:"recommendedActions"`
						NestedErrors       string   `json:"nestedErrors"`
						ErrorSource        string   `json:"errorSource"`
						ErrorCode          string   `json:"errorCode"`
						Data               string   `json:"data"`
						CanForce           bool     `json:"canForce"`
					} `json:"error"`
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
				Error struct {
					Message            string   `json:"message"`
					Details            string   `json:"details"`
					RecommendedActions []string `json:"recommendedActions"`
					NestedErrors       string   `json:"nestedErrors"`
					ErrorSource        string   `json:"errorSource"`
					ErrorCode          string   `json:"errorCode"`
					Data               string   `json:"data"`
					CanForce           bool     `json:"canForce"`
				} `json:"error"`
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
					Error    struct {
						Message            string   `json:"message"`
						Details            string   `json:"details"`
						RecommendedActions []string `json:"recommendedActions"`
						NestedErrors       string   `json:"nestedErrors"`
						ErrorSource        string   `json:"errorSource"`
						ErrorCode          string   `json:"errorCode"`
						Data               string   `json:"data"`
						CanForce           bool     `json:"canForce"`
					} `json:"error"`
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
				Error struct {
					Message            string   `json:"message"`
					Details            string   `json:"details"`
					RecommendedActions []string `json:"recommendedActions"`
					NestedErrors       string   `json:"nestedErrors"`
					ErrorSource        string   `json:"errorSource"`
					ErrorCode          string   `json:"errorCode"`
					Data               string   `json:"data"`
					CanForce           bool     `json:"canForce"`
				} `json:"error"`
			} `json:"inTransitKvmServers"`
		} `json:"zones"`
		Status string `json:"status"`
		State  string `json:"state"`
		URI    string `json:"uri"`
	} `json:"regions"`
	Metrics []struct {
		ResourceURI string `json:"resourceUri"`
		Resource    struct {
			Value   string `json:"value"`
			Name    string `json:"name"`
			URI     string `json:"uri"`
			Project struct {
				Name string `json:"name"`
				URI  string `json:"uri"`
			} `json:"project"`
			Zone struct {
				Name   string `json:"name"`
				URI    string `json:"uri"`
				Region struct {
					Name     string `json:"name"`
					URI      string `json:"uri"`
					Provider struct {
						Name         string `json:"name"`
						URI          string `json:"uri"`
						ProviderType struct {
							Name string `json:"name"`
							URI  string `json:"uri"`
						} `json:"providerType"`
					} `json:"provider"`
				} `json:"region"`
			} `json:"zone"`
		} `json:"resource"`
		Name        string `json:"name"`
		Units       string `json:"units"`
		Description string `json:"description"`
		Values      []struct {
			Value int       `json:"value"`
			Start time.Time `json:"start"`
			End   time.Time `json:"end"`
		} `json:"values"`
		Total        int `json:"total"`
		Start        int `json:"start"`
		Count        int `json:"count"`
		Associations []struct {
			Category string `json:"category"`
			Name     string `json:"name"`
			URI      string `json:"uri"`
		} `json:"associations"`
	} `json:"metrics"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

type ProviderList struct {
	Total   int        `json:"total"`
	Members []Provider `json:"members"`
}

// GetProviders returns ProviderList with optional query
// AWS credentials are not included when listing providers.
// leave filter blank to get all providers
// example query: "providerTypeUri EQ /rest/provider-types/aws"
func (c *Client) GetProviders(query string) (ProviderList, error) {
	var (
		uri         = "/rest/providers"
		queryParams = createQuery(&map[string]string{
			"query": query,
		})
		providers ProviderList
	)

	response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)

	if err != nil {
		return providers, err
	}

	if err := json.Unmarshal([]byte(response), &providers); err != nil {
		return providers, apiResponseError(response, err)
	}

	return providers, err
}

/* GetProviderByID returns an Provider by id
example view: "full"
discover: Will return the merged set of regions from AWS and existing regions in Onesphere.
*/
func (c *Client) GetProviderByID(id, view string, discover bool) (Provider, error) {
	var (
		uri         = "/rest/providers/" + id
		queryParams = createQuery(&map[string]string{
			"view":     view,
			"discover": strconv.FormatBool(discover),
		})
		provider Provider
	)

	if id == "" {
		return provider, fmt.Errorf("id must not be empty")
	}

	response, err := c.RestAPICall(rest.GET, uri, queryParams, nil)

	if err != nil {
		return provider, err
	}

	if err := json.Unmarshal([]byte(response), &provider); err != nil {
		return provider, apiResponseError(response, err)
	}

	return provider, err
}

// CreateProvider Creates a new Master provider or Member provider and returns updated Provider
// use GetProviderTypes() for ProviderTypeURI
func (c *Client) CreateProvider(providerRequest ProviderRequest) (Provider, error) {
	var (
		uri      = "/rest/providers"
		provider Provider
	)

	response, err := c.RestAPICall(rest.POST, uri, nil, providerRequest)

	if err != nil {
		return provider, err
	}

	if err := json.Unmarshal([]byte(response), &provider); err != nil {
		return provider, apiResponseError(response, err)
	}

	return provider, err
}

/* UpdateProvider using []*PatchOp returns updated provider on success

Allowed Ops for PATCH of providers: add | replace | remove

example:

Op: replace
Path: /name
Value: new name

*/
func (c *Client) UpdateProvider(provider Provider, updates []*PatchOp) (Provider, error) {
	if provider.ID == "" {
		return provider, fmt.Errorf("Provider must have a non-empty ID")
	}

	allowedOps := []string{"add", "replace", "remove"}

	for _, pb := range updates {
		fieldIsValid := false

		for _, allowedOp := range allowedOps {
			if pb.Op == allowedOp {
				fieldIsValid = true
			}
		}

		if !fieldIsValid {
			return provider, fmt.Errorf("UpdateProvider received invalid Op for update.\nReceived Op: %s\nValid Ops: %v\n", pb.Op, allowedOps)
		}
	}

	var (
		uri             = "/rest/providers/" + provider.ID
		updatedProvider Provider
	)

	response, err := c.RestAPICall(rest.PATCH, uri, nil, updates)

	if err != nil {
		return provider, err
	}

	if err := json.Unmarshal([]byte(response), &updatedProvider); err != nil {
		return provider, apiResponseError(response, err)
	}

	return updatedProvider, err
}
