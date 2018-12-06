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
	BillingAccountURI string    `json:"billingAccountUri"`
	ProjectUris       []string  `json:"projectUris"`
	Regions           []*Region `json:"regions"`
	Metrics           []*Metric `json:"metrics"`
	Created           time.Time `json:"created"`
	Modified          time.Time `json:"modified"`
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

// DeleteProvider Deletes Provider
func (c *Client) DeleteProvider(provider Provider) error {
	if provider.ID == "" {
		return fmt.Errorf("Provider must have a non-empty ID")
	}

	var uri = "/rest/providers/" + provider.ID

	response, err := c.RestAPICall(rest.DELETE, uri, nil, nil)

	if err != nil {
		return apiResponseError(response, err)
	}

	return nil
}
