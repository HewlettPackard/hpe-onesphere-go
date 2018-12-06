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

// NamedUriIdentifier defines JSON Struct for { id, name, uri }
type NamedUriIdentifier struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Uri  string `json:"uri"`
}

// AddressWithType defines JSON Struct for { address, addressType }
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

func (c *Client) RestAPICall(method rest.Method, path string, queryParams map[string]string, values interface{}) (string, error) {
	jsonValue, err := json.Marshal(values)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(method.String(), c.buildURL(path), bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.Auth.Token)

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

// Membership Roles APIs

func (c *Client) GetMembershipRoles() (string, error) {
	return c.callHTTPRequest("GET", "/rest/membership-roles", nil, nil)
}

// Memberships APIs

func (c *Client) GetMemberships(query string) (string, error) {
	params := map[string]string{"query": query}
	return c.callHTTPRequest("GET", "/rest/memberships", params, nil)
}

func (c *Client) CreateMembership(userUri, roleUri, projectUri string) (string, error) {
	values := map[string]string{"userUri": userUri, "membershipRoleUri": roleUri, "projectUri": projectUri}
	return c.callHTTPRequest("POST", "/rest/memberships", nil, values)
}

func (c *Client) DeleteMembership(userUri, roleUri, projectUri string) (string, error) {
	values := map[string]string{"userUri": userUri, "membershipRoleUri": roleUri, "projectUri": projectUri}
	return c.callHTTPRequest("DELETE", "/rest/memberships", nil, values)
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

// Networks APIs

func (c *Client) GetNetwork(networkID string) (string, error) {
	return c.callHTTPRequest("GET", "/rest/networks/"+networkID, nil, nil)
}

// infoArray: [{op, path, value}]
func (c *Client) UpdateNetwork(networkID string, infoArray []string) (string, error) {
	values := infoArray
	return c.callHTTPRequest("PUT", "/rest/networks/"+networkID, nil, values)
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

// Projects APIs

func (c *Client) GetProjects(userQuery, view string) (string, error) {
	params := map[string]string{}
	if strings.TrimSpace(userQuery) != "" {
		params["userQuery"] = userQuery
	}
	if strings.TrimSpace(view) != "" {
		params["view"] = view
	}
	return c.callHTTPRequest("GET", "/rest/projects", params, nil)
}

func (c *Client) CreateProject(name, description string, tagUris []string) (string, error) {
	values := map[string]interface{}{
		"name":        name,
		"description": description,
		"tagUris":     tagUris}
	return c.callHTTPRequest("POST", "/rest/projects", nil, values)
}

func (c *Client) GetProject(projectID, view string) (string, error) {
	params := map[string]string{"view": view}
	return c.callHTTPRequest("GET", "/rest/projects/"+projectID, params, nil)
}

func (c *Client) DeleteProject(projectID string) (string, error) {
	// return c.callHTTPRequest("DELETE", "/rest/projects/"+projectID, nil, nil)
	return "", c.notImplementedError(rest.DELETE, "/rest/projects", "projects")
}

func (c *Client) UpdateProject(projectID, name, description string, tagUris []string) (string, error) {
	values := map[string]interface{}{
		"name":        name,
		"description": description,
		"tagUris":     tagUris}
	return c.callHTTPRequest("PUT", "/rest/projects/"+projectID, nil, values)
}

// Provider Types APIs

func (c *Client) GetProviderTypes() (string, error) {
	return c.callHTTPRequest("GET", "/rest/provider-types", nil, nil)
}

// Providers APIs

func (c *Client) GetProviders(query string) (string, error) {
	params := map[string]string{"query": query}
	return c.callHTTPRequest("GET", "/rest/providers", params, nil)
}

// state: "Enabled|Disabled"
func (c *Client) CreateProvider(providerID, providerTypeUri, accessKey, secretKey string,
	paymentProvider bool,
	s3CostBucket, masterUri, state string) (string, error) {
	values := map[string]interface{}{
		"id":              providerID,
		"providerTypeUri": providerTypeUri,
		"accessKey":       accessKey,
		"secretKey":       secretKey,
		"paymentProvider": paymentProvider,
		"s3CostBucket":    s3CostBucket,
		"masterUri":       masterUri,
		"state":           state}
	return c.callHTTPRequest("POST", "/rest/providers", nil, values)
}

// view="full"
// discover: boolean
func (c *Client) GetProvider(providerID, view string, discover bool) (string, error) {
	params := map[string]string{
		"view":     view,
		"discover": strconv.FormatBool(discover)}
	return c.callHTTPRequest("GET", "/rest/providers/"+providerID, params, nil)
}

func (c *Client) DeleteProvider(providerID string) (string, error) {
	return c.callHTTPRequest("DELETE", "/rest/providers/"+providerID, nil, nil)
}

// infoArray: [{op, path, value}]
// op: "add|replace|remove"
func (c *Client) UpdateProvider(providerID, infoArray string) (string, error) {
	return c.callHTTPRequest("PUT", "/rest/providers/"+providerID, nil, infoArray)
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

// Regions APIs

func (c *Client) GetRegions(query, view string) (string, error) {
	params := map[string]string{"query": query, "view": view}
	return c.callHTTPRequest("GET", "/rest/regions", params, nil)
}

func (c *Client) CreateRegion(name, providerUri, locLatitude, locLongitude string) (string, error) {
	values := map[string]interface{}{
		"location": map[string]interface{}{
			"latitude":  locLatitude,
			"longitude": locLongitude},
		"name":        name,
		"providerUri": providerUri}
	return c.callHTTPRequest("POST", "/rest/regions", nil, values)
}

func (c *Client) GetRegion(regionID, view string, discover bool) (string, error) {
	params := map[string]string{"view": view, "discover": strconv.FormatBool(discover)}
	return c.callHTTPRequest("GET", "/rest/regions/"+regionID, params, nil)
}

func (c *Client) DeleteRegion(regionID string, force bool) (string, error) {
	params := map[string]string{"force": strconv.FormatBool(force)}
	return c.callHTTPRequest("DELETE", "/rest/regions/"+regionID, params, nil)
}

// infoArray: [{op, path, value}]
// op: "add|replace"
// path: "/name|/location"
func (c *Client) PatchRegion(regionID string, infoArray []string) (string, error) {
	return c.callHTTPRequest("PUT", "/rest/regions/"+regionID, nil, infoArray)
}

func (c *Client) UpdateRegion(regionID, region string) (string, error) {
	return c.callHTTPRequest("PUT", "/rest/regions/"+regionID, nil, region)
}

func (c *Client) GetRegionConnection(regionID string) (string, error) {
	return c.callHTTPRequest("GET", "/rest/regions/"+regionID+"/connection", nil, nil)
}

// state: "Enabling|Enabled|Disabling|Disabled"
func (c *Client) CreateRegionConnection(regionID, endpointUuid, name, ipAddress, username, password string,
	port int,
	state, uri string) (string, error) {
	values := map[string]interface{}{
		"endpointUuid": endpointUuid,
		"name":         name,
		"location": map[string]interface{}{
			"ipAddress": ipAddress,
			"username":  username,
			"password":  password,
			"port":      strconv.Itoa(port)},
		"state": state,
		"uri":   uri}
	return c.callHTTPRequest("POST", "/rest/regions/"+regionID+"/connection", nil, values)
}

func (c *Client) DeleteRegionConnection(regionID string) (string, error) {
	return c.callHTTPRequest("DELETE", "/rest/regions/"+regionID+"/connection", nil, nil)
}

func (c *Client) GetRegionConnectorImage(regionID string) (string, error) {
	return c.callHTTPRequest("GET", "/rest/regions/"+regionID+"/connector-image", nil, nil)
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

// Service Types APIs

func (c *Client) GetServiceTypes() (string, error) {
	return c.callHTTPRequest("GET", "/rest/service-types", nil, nil)
}

func (c *Client) GetServiceType(serviceTypeID string) (string, error) {
	return c.callHTTPRequest("GET", "/rest/service-types/"+serviceTypeID, nil, nil)
}

// Services APIs

func (c *Client) GetServices(query, userQuery, view string) (string, error) {
	params := map[string]string{"query": query, "userQuery": userQuery, "view": view}
	return c.callHTTPRequest("GET", "/rest/services", params, nil)
}

// view: "full|deployment"
func (c *Client) GetService(serviceID, view string) (string, error) {
	params := map[string]string{"view": view}
	return c.callHTTPRequest("GET", "/rest/services/"+serviceID, params, nil)
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

// Tag Keys APIs

// view: "full"
func (c *Client) GetTagKeys(view string) (string, error) {
	params := map[string]string{"view": view}
	return c.callHTTPRequest("GET", "/rest/tag-keys", params, nil)
}

func (c *Client) CreateTagKey(name string) (string, error) {
	values := map[string]string{"name": name}
	return c.callHTTPRequest("POST", "/rest/tag-keys", nil, values)
}

// view: "full"
func (c *Client) GetTagKey(tagKeyID, view string) (string, error) {
	params := map[string]string{"view": view}
	return c.callHTTPRequest("GET", "/rest/tag-keys/"+tagKeyID, params, nil)
}

func (c *Client) DeleteTagKey(tagKeyID string) (string, error) {
	return c.callHTTPRequest("DELETE", "/rest/tag-keys/"+tagKeyID, nil, nil)
}

// Tags APIs

// view: "full"
func (c *Client) GetTags(view string) (string, error) {
	params := map[string]string{"view": view}
	return c.callHTTPRequest("GET", "/rest/tags", params, nil)
}

func (c *Client) CreateTag(name, tagKeyUri string) (string, error) {
	values := map[string]string{"name": name, "tagKeyUri": tagKeyUri}
	return c.callHTTPRequest("POST", "/rest/tags", nil, values)
}

// view: "full"
func (c *Client) GetTag(tagID, view string) (string, error) {
	params := map[string]string{"view": view}
	return c.callHTTPRequest("GET", "/rest/tags/"+tagID, params, nil)
}

func (c *Client) DeleteTag(tagID string) (string, error) {
	return c.callHTTPRequest("DELETE", "/rest/tags/"+tagID, nil, nil)
}

// Users APIs

func (c *Client) GetUsers(userQuery string) (string, error) {
	params := map[string]string{"userQuery": userQuery}
	return c.callHTTPRequest("GET", "/rest/users", params, nil)
}

// role: "administrator|analyst|consumer|project-creator"
func (c *Client) CreateUser(email, name, password, role string) (string, error) {
	values := map[string]string{"email": email, "name": name, "password": password, "role": role}
	return c.callHTTPRequest("POST", "/rest/users", nil, values)
}

func (c *Client) GetUser(userID string) (string, error) {
	return c.callHTTPRequest("GET", "/rest/users/"+userID, nil, nil)
}

// role: "administrator|analyst|consumer|project-creator"
func (c *Client) UpdateUser(userID, email, name, password, role string) (string, error) {
	values := map[string]string{"email": email, "name": name, "password": password, "role": role}
	return c.callHTTPRequest("PUT", "/rest/users/"+userID, nil, values)
}

func (c *Client) DeleteUser(userID string) (string, error) {
	return c.callHTTPRequest("DELETE", "/rest/users/"+userID, nil, nil)
}

// Versions APIs

func (c *Client) GetVersions() (string, error) {
	return c.callHTTPRequest("GET", "/rest/about/versions", nil, nil)
}

// Virtual Machine Profiles APIs

func (c *Client) GetVirtualMachineProfiles(zoneUri, serviceUri string) (string, error) {
	params := map[string]string{"zoneUri": zoneUri, "serviceUri": serviceUri}
	return c.callHTTPRequest("GET", "/rest/virtual-machine-profiles", params, nil)
}

func (c *Client) GetVirtualMachineProfile(vmProfileID string) (string, error) {
	return c.callHTTPRequest("GET", "/rest/virtual-machine-profiles/"+vmProfileID, nil, nil)
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

// Zone Types APIs

func (c *Client) GetZoneTypes() (string, error) {
	return c.callHTTPRequest("GET", "/rest/zone-types", nil, nil)
}

func (c *Client) GetZoneTypeResourceProfiles(zoneTypeID string) (string, error) {
	return c.callHTTPRequest("GET", "/rest/zone-types/"+zoneTypeID+"/resource-profiles", nil, nil)
}

// Zones APIs

func (c *Client) GetZones(query, regionUri, applianceUri string) (string, error) {
	params := map[string]string{"query": query, "regionUri": regionUri, "applianceUri": applianceUri}
	return c.callHTTPRequest("GET", "/rest/zones", params, nil)
}

func (c *Client) CreateZone(zoneData string) (string, error) {
	return c.callHTTPRequest("POST", "/rest/zones", nil, zoneData)
}

// view: "full"
func (c *Client) GetZone(zoneID, view string) (string, error) {
	params := map[string]string{"view": view}
	return c.callHTTPRequest("GET", "/rest/zones/"+zoneID, params, nil)
}

// op: "add|replace|remove"
func (c *Client) UpdateZone(zoneID, op, path string, value interface{}) (string, error) {
	values := map[string]interface{}{"op": op, "path": path, "value": value}
	return c.callHTTPRequest("PUT", "/rest/zones/"+zoneID, nil, values)
}

func (c *Client) DeleteZone(zoneID string, force bool) (string, error) {
	params := map[string]string{"force": strconv.FormatBool(force)}
	return c.callHTTPRequest("DELETE", "/rest/zones/"+zoneID, params, nil)
}

// actionType: "reset|add-capacity|reduce-capacity"
// resourceType: "compute|storage"
func (c *Client) ActionOnZone(zoneID, actionType, resourceType string, resourceCapacity int) (string, error) {
	values := map[string]interface{}{
		"type": actionType,
		"resourceOp": map[string]interface{}{
			"resourceType":     resourceType,
			"resourceCapacity": resourceCapacity}}
	return c.callHTTPRequest("POST", "/rest/zones/"+zoneID+"/actions", nil, values)
}

func (c *Client) GetZoneApplianceImage(zoneID string) (string, error) {
	return c.callHTTPRequest("GET", "/rest/zones/"+zoneID+"/appliance-image", nil, nil)
}

func (c *Client) GetZoneTaskStatus(zoneID string) (string, error) {
	return c.callHTTPRequest("GET", "/rest/zones/"+zoneID+"/task-status", nil, nil)
}

func (c *Client) GetZoneConnections(zoneID, uuid string) (string, error) {
	params := map[string]string{"uuid": uuid}
	return c.callHTTPRequest("GET", "/rest/zones/"+zoneID+"/connections", params, nil)
}

// state: "Enabling|Enabled|Disabling|Disabled"
func (c *Client) CreateZoneConnection(zoneID, uuid, name, ipAddress, username, password string,
	port int, state string) (string, error) {
	values := map[string]interface{}{
		"uuid": uuid,
		"name": name,
		"location": map[string]interface{}{
			"ipAddress": ipAddress,
			"username":  username,
			"password":  password,
			"port":      port},
		"state": state}
	return c.callHTTPRequest("POST", "/rest/zones/"+zoneID+"/connections", nil, values)
}

func (c *Client) DeleteZoneConnection(zoneID, uuid string) (string, error) {
	return c.callHTTPRequest("DELETE", "/rest/zones/"+zoneID+"/connections/"+uuid, nil, nil)
}

// op: "add|replace|remove"
func (c *Client) UpdateZoneConnection(zoneID, uuid, op, path string, value interface{}) (string, error) {
	values := map[string]interface{}{"op": op, "path": path, "value": value}
	return c.callHTTPRequest("PUT", "/rest/zones/"+zoneID+"/connections", nil, values)
}
