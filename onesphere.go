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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/HewlettPackard/hpe-onesphere-go/rest"
)

// Client contains all the methods needed to interact with the OneSphere API
// use Connect() to return a *Client
type Client struct {
	Auth *Auth
}

// Auth contains the Token and HostURL of the OneSphere API connection
type Auth struct {
	Token   string
	HostURL string
}

// NamedUri defines JSON struct for { name, uri }
type NamedUri struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}

// NamedUriIdentifier defines JSON struct for { id, name, uri }
type NamedUriIdentifier struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URI  string `json:"uri"`
}

// AddressWithType defines JSON struct for { address, addressType }
type AddressWithType struct {
	Address     string `json:"address"`
	AddressType string `json:"addressType"`
}

func closer(r io.Closer, funcName string) {
	err := r.Close()
	if err != nil {
		fmt.Printf("Error closing response body reader in %s\n%v\n", funcName, err)
	}
}

// Connect provides an interface to make calls to the OneSphere API
func Connect(hostURL, user, password string) (*Client, error) {
	fullUrl := hostURL + "/rest/session"
	values := map[string]string{"userName": user, "password": password}
	jsonValue, err := json.Marshal(values)
	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer closer(resp.Body, fmt.Sprintf("onesphere.Connect(%s,%s,#masked password#)", hostURL, user))

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//bodyStr := string(body)
	var dat map[string]string
	err = json.Unmarshal(body, &dat)
	if err != nil {
		return nil, err
	}

	return &Client{
		Auth: &Auth{
			HostURL: hostURL,
			Token:   dat["token"],
		},
	}, nil

}

func (c *Client) callHTTPRequest(method, path string, params map[string]string, values interface{}) (string, error) {
	jsonValue, err := json.Marshal(values)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(method, c.buildURL(path), bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.Auth.Token)

	if params != nil && len(params) > 0 {
		q := req.URL.Query()
		for key, value := range params {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer closer(resp.Body, fmt.Sprintf("onesphere.callHTTPRequest(%s,%s,%v,%v)", method, path, params, values))

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	bodyStr := string(bodyBytes)
	return bodyStr, nil
}

// createQuery returns a map for passing to Client.RestAPICall
// if values are "", it will be omitted from the map
func createQuery(params *map[string]string) map[string]string {
	for k, v := range *params {
		if len(v) == 0 {
			delete(*params, k)
		}
	}
	return *params
}

// apiResponseError returns a more helpful error including the response payload
func apiResponseError(response string, err error) error {
	return fmt.Errorf("Unmarshal Error:\n\tRaw Response: %v\n\tError: %v", response, err)
}

func (c *Client) RestAPICallCustomHeaders(method rest.Method, customHeaders map[string]string, path string, queryParams map[string]string, values interface{}) (string, error) {

	jsonValue, err := json.Marshal(values)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(method.String(), c.buildURL(path), bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", c.Auth.Token)
	for key, value := range customHeaders {
		req.Header.Set(key, value)

	}

	if queryParams != nil && len(queryParams) > 0 {
		q := req.URL.Query()
		for key, value := range queryParams {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer closer(resp.Body, fmt.Sprintf("onesphere.RestAPICall(%v,%s,%v,%v)", method, path, queryParams, values))

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	bodyStr := string(bodyBytes)
	return bodyStr, nil
}

func (c *Client) RestAPICall(method rest.Method, path string, queryParams map[string]string, values interface{}) (string, error) {
	return c.RestAPICallCustomHeaders(method, map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}, path, queryParams, values)
}

func (c *Client) RestAPICallPatch(path string, queryParams map[string]string, values interface{}) (string, error) {
	return c.RestAPICallCustomHeaders(rest.PATCH, map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json-patch+json",
	}, path, queryParams, values)
}

func (c *Client) buildURL(path string) string {
	return c.Auth.HostURL + path
}

func (c *Client) notImplementedError(method rest.Method, endpoint, path string) error {
	return fmt.Errorf("%s %s is not yet implemented.\nSee: %s/docs/api/endpoint?&path=%%2F%s", method, endpoint, c.Auth.HostURL, path)
}

func (c *Client) Disconnect() {
	_, err := c.callHTTPRequest("DELETE", "/rest/session", nil, nil)
	if err != nil {
		fmt.Printf("Error logging out of OneSphere api in onesphere.Disconnect()\n%v\n", err)
	}
}

// Billing Accounts APIs

func (c *Client) GetBillingAccounts(query, view string) (string, error) {
	params := map[string]string{}
	if strings.TrimSpace(query) != "" {
		params["query"] = query
	}
	if strings.TrimSpace(view) != "" {
		params["view"] = view
	}
	return c.callHTTPRequest("GET", "/rest/billing-accounts", params, nil)
}

func (c *Client) CreateBillingAccount(apiAccessKey, description, directoryUri, enrollmentNumber, name, providerTypeUri string) (string, error) {
	values := map[string]string{
		"apiAccessKey":     apiAccessKey,
		"description":      description,
		"directoryUri":     directoryUri,
		"enrollmentNumber": enrollmentNumber,
		"name":             name,
		"providerTypeUri":  providerTypeUri,
	}
	return c.callHTTPRequest("POST", "/rest/billing-accounts", nil, values)
}

func (c *Client) GetBillingAccount(id string) (string, error) {
	return c.callHTTPRequest("GET", "/rest/billing-accounts/"+id, nil, nil)
}

func (c *Client) DeleteBillingAccount(id string) (string, error) {
	return c.callHTTPRequest("DELETE", "/rest/billing-accounts/"+id, nil, nil)
}

// UpdateBillingAccount sends PATCH with Op: "add|replace|remove"
func (c *Client) UpdateBillingAccount(id string, patchPayload []*PatchOp) (string, error) {
	validOps := []string{"add", "replace", "remove"}
	for _, pb := range patchPayload {
		opIsValid := false
		for _, validOp := range validOps {
			if pb.Op == validOp {
				opIsValid = true
			}
		}
		if !opIsValid {
			return "", fmt.Errorf("UpdateBillingAccount received invalid Op in patchBodies.\nReceived Op: %s\nValid Ops: %v\n", pb.Op, validOps)
		}
	}

	return c.callHTTPRequest("PATCH", "/rest/billing-accounts/"+id, nil, patchPayload)
}

// Connect App APIs

// GetConnectApp allowed operating systems: ["windows", "mac"]
func (c *Client) GetConnectApp(os string) (string, error) {
	validOperatingSystems := []string{
		"windows",
		"mac",
	}
	osIsValid := false
	for _, validOperatingSystem := range validOperatingSystems {
		if os == validOperatingSystem {
			osIsValid = true
		}
	}
	if !osIsValid {
		return "", fmt.Errorf("GetConnectApp received invalid os parameter.\nReceived os: %s\nValid os values: %v\n", os, validOperatingSystems)
	}

	params := map[string]string{"os": os}
	return c.callHTTPRequest("GET", "/rest/connect-app", params, nil)
}

// Events APIs

func (c *Client) GetEvents(resourceUri string) (string, error) {
	// params := map[string]string{"resourceUri": resourceUri}
	// return c.callHTTPRequest("GET", "/rest/events", params, nil)
	return "", c.notImplementedError(rest.GET, "/rest/events", "events")
}

// Keypairs APIs

func (c *Client) GetKeyPair(regionUri, projectUri string) (string, error) {
	params := map[string]string{"regionUri": regionUri, "projectUri": projectUri}
	return c.callHTTPRequest("GET", "/rest/keypairs", params, nil)
}

// Metrics APIs

func (c *Client) GetMetrics(
	resourceUri, category, groupBy, query, name string,
	periodStart, period string,
	periodCount int,
	view string,
	start, count int) (string, error) {
	params := map[string]string{
		"resourceUri": resourceUri,
		"category":    category,
		"groupBy":     groupBy,
		"query":       query,
		"nameArray":   name,
		"periodStart": periodStart,
		"period":      period,
		"periodCount": strconv.Itoa(periodCount),
		"view":        view,
		"start":       strconv.Itoa(start),
		"count":       strconv.Itoa(count)}
	return c.callHTTPRequest("GET", "/rest/metrics", params, nil)
}

// Onboarding APIs

func (c *Client) GetAzureLoginProperties() (string, error) {
	return c.callHTTPRequest("GET", "/rest/onboarding/azure/properties", nil, nil)
}

func (c *Client) GetAzureProviderInfo(directoryUri, location string) (string, error) {
	params := map[string]string{"directoryUri": directoryUri, "location": location}
	return c.callHTTPRequest("GET", "/rest/onboarding/azure/provider-info", params, nil)
}

func (c *Client) GetAzureSubscriptions(directoryUri, location string) (string, error) {
	params := map[string]string{"directoryUri": directoryUri, "location": location}
	return c.callHTTPRequest("GET", "/rest/onboarding/azure/subscriptions", params, nil)
}

/* UpdateAzureSubscription allowed Ops in patchPayload:
  - add
	- replace
*/
func (c *Client) UpdateAzureSubscription(directoryUri, location, subscriptionId string, patchPayload []*PatchOp) (string, error) {
	allowedOps := []string{"add", "replace"}

	for _, pb := range patchPayload {
		opIsValid := false

		for _, allowedOp := range allowedOps {
			if pb.Op == allowedOp {
				opIsValid = true
			}
		}

		if !opIsValid {
			return "", fmt.Errorf("UpdateAzureSubscription received invalid Op for update.\nReceived Op: %s\nValid Ops: %v\n", pb.Op, allowedOps)
		}
	}

	params := map[string]string{"directoryUri": directoryUri, "location": location}
	values := map[string][]*PatchOp{"items": patchPayload}
	return c.callHTTPRequest("PATCH", "/rest/onboarding/azure/subscriptions/"+subscriptionId, params, values)
}

// Password Reset APIs

func (c *Client) ResetSingleUsePassword(email string) (string, error) {
	values := map[string]string{"email": email}
	return c.callHTTPRequest("POST", "/rest/password-reset", nil, values)
}

func (c *Client) ChangePassword(password, token string) (string, error) {
	values := map[string]string{"password": password, "token": token}
	return c.callHTTPRequest("POST", "/rest/password-reset/change", nil, values)
}

// Rates APIs

func (c *Client) GetRates(resourceUri, effectiveForDate, effectiveDate, metricName string,
	active bool, start, count int) (string, error) {
	params := map[string]string{
		"resourceUri":      resourceUri,
		"effectiveForDate": effectiveForDate,
		"effectiveDate":    effectiveDate,
		"metricName":       metricName,
		"active":           strconv.FormatBool(active),
		"start":            strconv.Itoa(start),
		"count":            strconv.Itoa(count)}
	return c.callHTTPRequest("GET", "/rest/rates", params, nil)
}

func (c *Client) GetRate(rateID string) (string, error) {
	return c.callHTTPRequest("GET", "/rest/rates/"+rateID, nil, nil)
}

// Roles APIs

func (c *Client) GetRoles() (string, error) {
	return c.callHTTPRequest("GET", "/rest/roles", nil, nil)
}

// Servers APIs

func (c *Client) GetServers(regionUri, applianceUri, zoneUri string) (string, error) {
	params := map[string]string{}
	if regionUri != "" {
		params["regionUri"] = regionUri
	}
	if applianceUri != "" {
		params["applianceUri"] = applianceUri
	}
	if zoneUri != "" {
		params["zoneUri"] = zoneUri
	}
	return c.callHTTPRequest("GET", "/rest/servers", params, nil)
}

func (c *Client) CreateServer(server *Server) (string, error) {
	values := map[string]*Server{
		"server": server,
	}
	return c.callHTTPRequest("POST", "/rest/servers", nil, values)
}

func (c *Client) DeleteServer(serverID string, force bool) (string, error) {
	params := map[string]string{}
	if force {
		params["force"] = "true"
	}
	return c.callHTTPRequest("DELETE", "/rest/servers/"+serverID, params, nil)
}

func (c *Client) GetServer(serverID string) (string, error) {
	return c.callHTTPRequest("GET", "/rest/servers/"+serverID, nil, nil)
}

/* UpdateServer allowed Ops in patchPayload:
- replace
- remove
*/
func (c *Client) UpdateServer(serverID string, patchPayload []*PatchOp) (string, error) {
	allowedOps := []string{"replace", "remove"}

	for _, pb := range patchPayload {
		opIsValid := false

		for _, allowedOp := range allowedOps {
			if pb.Op == allowedOp {
				opIsValid = true
			}
		}

		if !opIsValid {
			return "", fmt.Errorf("UpdateServer received invalid Op for update.\nReceived Op: %s\nValid Ops: %v\n", pb.Op, allowedOps)
		}
	}

	values := map[string][]*PatchOp{"body": patchPayload}
	return c.callHTTPRequest("PATCH", "/rest/servers/"+serverID, nil, values)
}

// Session APIs

// view: "full"
func (c *Client) GetSession(view string) (string, error) {
	params := map[string]string{"view": view}
	return c.callHTTPRequest("GET", "/rest/session", params, nil)
}

func (c *Client) GetSessionIdp(userName string) (string, error) {
	// params := map[string]string{"userName": userName}
	// return c.callHTTPRequest("GET", "/rest/session/idp", params, nil)
	return "", c.notImplementedError(rest.GET, "/rest/account", "account")
}

// GetStatus calls the /rest/status endpoint
func (c *Client) GetStatus() (string, error) {
	return c.callHTTPRequest("GET", "/rest/status", nil, nil)
}

// Versions APIs

func (c *Client) GetVersions() (string, error) {
	return c.callHTTPRequest("GET", "/rest/about/versions", nil, nil)
}

// Volumes APIs

// view: "full"
func (c *Client) GetVolumes(query, view string) (string, error) {
	params := map[string]string{"query": query, "view": view}
	return c.callHTTPRequest("GET", "/rest/volumes", params, nil)
}

func (c *Client) CreateVolume(name string, sizeGiB int, zoneUri, projectUri string) (string, error) {
	values := map[string]interface{}{
		"name":       name,
		"sizeGiB":    strconv.Itoa(sizeGiB),
		"zoneUri":    zoneUri,
		"projectUri": projectUri}
	return c.callHTTPRequest("POST", "/rest/volumes", nil, values)
}

func (c *Client) GetVolume(volumeID string) (string, error) {
	return c.callHTTPRequest("GET", "/rest/volumes/"+volumeID, nil, nil)
}

func (c *Client) UpdateVolume(volumeID, name string, sizeGiB int) (string, error) {
	values := map[string]interface{}{
		"name":    name,
		"sizeGiB": strconv.Itoa(sizeGiB)}
	return c.callHTTPRequest("PUT", "/rest/volumes/"+volumeID, nil, values)
}

func (c *Client) DeleteVolume(volumeID string) (string, error) {
	return c.callHTTPRequest("DELETE", "/rest/volumes/"+volumeID, nil, nil)
}
