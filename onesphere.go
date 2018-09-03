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
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// API contains all the methods needed to interact with the OneSphere API
// use Connect() to return an *API
type API struct {
	Auth *Auth
}

// Auth contains the Token and HostURL of the OneSphere API connection
type Auth struct {
	Token   string
	HostURL string
}

var auth *Auth

// Connect provides an interface to make calls to the OneSphere API
func Connect(hostURL, user, password string) (*API, error) {
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
	defer resp.Body.Close()

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

	return &API{
		Auth: &Auth{
			HostURL: hostURL,
			Token:   dat["token"],
		},
	}, nil

}

func (api *API) callHTTPRequest(method, path string, params map[string]string, values interface{}) (string, error) {
	jsonValue, err := json.Marshal(values)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(method, api.buildURL(path), bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", api.Auth.Token)

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
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	bodyStr := string(bodyBytes)
	return bodyStr, nil
}

func (api *API) buildURL(path string) string {
	return api.Auth.HostURL + path
}

func (api *API) notImplementedError(method, endpoint, path string) error {
	return fmt.Errorf("%s %s is not yet implemented.\nSee: %s/docs/api/endpoint?&path=%%2F%s", method, endpoint, api.Auth.HostURL, path)
}

func (api *API) Disconnect() {
	api.callHTTPRequest("DELETE", "/rest/session", nil, nil)
}

// Account APIs

// view="full"
func (api *API) GetAccount(view string) (string, error) {
	// params := map[string]string{"view": view}
	// return api.callHTTPRequest("GET", "/rest/account", params, nil)
	return "", api.notImplementedError("GET", "/rest/account", "account")
}

// Appliances APIs

func (api *API) GetAppliances(name, regionUri string) (string, error) {
	params := map[string]string{}
	if strings.TrimSpace(name) != "" {
		params["name"] = name
	}
	if strings.TrimSpace(regionUri) != "" {
		params["regionUri"] = regionUri
	}
	return api.callHTTPRequest("GET", "/rest/appliances", params, nil)
}

func (api *API) CreateAppliance(epAddress, epUsername, epPassword,
	name, regionUri, applianceType string) (string, error) {
	values := map[string]interface{}{
		"endpoint": map[string]interface{}{
			"address":  epAddress,
			"password": epPassword,
			"username": epUsername},
		"name":      name,
		"regionUri": regionUri,
		"type":      applianceType}
	return api.callHTTPRequest("POST", "/rest/appliances", nil, values)
}

func (api *API) GetAppliance(applianceID string) (string, error) {
	return api.callHTTPRequest("GET", "/rest/appliances/"+applianceID, nil, nil)
}

func (api *API) DeleteAppliance(applianceID string) (string, error) {
	return api.callHTTPRequest("DELETE", "/rest/appliances/"+applianceID, nil, nil)
}

// infoArray: [{op, path, value}]
// op: "replace|remove"
func (api *API) UpdateAppliance(applianceID string, infoArray []string) (string, error) {
	values := infoArray
	return api.callHTTPRequest("PUT", "/rest/appliances/"+applianceID, nil, values)
}

// Billing Accounts APIs

func (api *API) GetBillingAccounts(query, view string) (string, error) {
	params := map[string]string{}
	if strings.TrimSpace(query) != "" {
		params["query"] = query
	}
	if strings.TrimSpace(view) != "" {
		params["view"] = view
	}
	return api.callHTTPRequest("GET", "/rest/billing-accounts", params, nil)
}

func (api *API) CreateBillingAccount(apiAccessKey, description, directoryUri, enrollmentNumber, name, providerTypeUri string) (string, error) {
	values := map[string]string{
		"apiAccessKey":     apiAccessKey,
		"description":      description,
		"directoryUri":     directoryUri,
		"enrollmentNumber": enrollmentNumber,
		"name":             name,
		"providerTypeUri":  providerTypeUri,
	}
	return api.callHTTPRequest("POST", "/rest/billing-accounts", nil, values)
}

func (api *API) GetBillingAccount(id string) (string, error) {
	return api.callHTTPRequest("GET", "/rest/billing-accounts/"+id, nil, nil)
}

func (api *API) DeleteBillingAccount(id string) (string, error) {
	return api.callHTTPRequest("DELETE", "/rest/billing-accounts/"+id, nil, nil)
}

// UpdateBillingAccount sends PATCH with Op: "add|replace|remove"
func (api *API) UpdateBillingAccount(id string, patchPayload []*PatchOp) (string, error) {
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

	return api.callHTTPRequest("PATCH", "/rest/billing-accounts/"+id, nil, patchPayload)
}

// Catalog Types APIs

func (api *API) GetCatalogTypes() (string, error) {
	return api.callHTTPRequest("GET", "/rest/catalog-types", nil, nil)
}

// Catalogs APIs

func (api *API) GetCatalogs(userQuery, view string) (string, error) {
	params := map[string]string{}
	if strings.TrimSpace(userQuery) != "" {
		params["userQuery"] = userQuery
	}
	if strings.TrimSpace(view) != "" {
		params["view"] = view
	}
	return api.callHTTPRequest("GET", "/rest/catalogs", params, nil)
}

/* CreateCatalog sends POST with catalogTypeUri:
  - /rest/catalog-types/aws-az
	- /rest/catalog-types/vcenter
	- /rest/catalog-types/kvm
	- /rest/catalog-types/helm-charts-repository
	- /rest/catalog-types/docker-hub
	- /rest/catalog-types/docker-registry
	- /rest/catalog-types/docker-trusted-registry
	- /rest/catalog-types/private-docker-registry
	- /rest/catalog-types/amazon-ecr
	- /rest/catalog-types/azure-container-registry
	- /rest/catalog-types/hpe-managed
*/
func (api *API) CreateCatalog(accessKey, catalogTypeUri, name, password, regionName, secretKey, url, username string) (string, error) {
	validCatalogTypeUris := []string{
		"/rest/catalog-types/aws-az",
		"/rest/catalog-types/vcenter",
		"/rest/catalog-types/kvm",
		"/rest/catalog-types/helm-charts-repository",
		"/rest/catalog-types/docker-hub",
		"/rest/catalog-types/docker-registry",
		"/rest/catalog-types/docker-trusted-registry",
		"/rest/catalog-types/private-docker-registry",
		"/rest/catalog-types/amazon-ecr",
		"/rest/catalog-types/azure-container-registry",
		"/rest/catalog-types/hpe-managed",
	}
	catalogTypeUriIsValid := false
	for _, validCatalogTypeUri := range validCatalogTypeUris {
		if catalogTypeUri == validCatalogTypeUri {
			catalogTypeUriIsValid = true
		}
	}
	if !catalogTypeUriIsValid {
		return "", fmt.Errorf("CreateCatalog received invalid catalogTypeUri.\nReceived catalogTypeUri: %s\nValid catalogTypeUri values: %v\n", catalogTypeUri, validCatalogTypeUris)
	}

	values := map[string]string{
		"accessKey":      accessKey,
		"catalogTypeUri": catalogTypeUri,
		"name":           name,
		"password":       password,
		"regionName":     regionName,
		"secretKey":      secretKey,
		"url":            url,
		"username":       username}
	return api.callHTTPRequest("POST", "/rest/catalogs", nil, values)
}

func (api *API) GetCatalog(catalogID, view string) (string, error) {
	params := map[string]string{"view": view}
	return api.callHTTPRequest("GET", "/rest/catalogs/"+catalogID, params, nil)
}

func (api *API) DeleteCatalog(catalogID string) (string, error) {
	// return api.callHTTPRequest("DELETE", "/rest/catalogs/"+catalogID, nil, nil)
	return "", api.notImplementedError("DELETE", "/rest/catalogs/"+catalogID, "catalogs")
}

/* UpdateCatalog allowed fields to update:
	[Op] => Field
  - add => /name, /password, /accessKey, /secretKey, /state
	- replace => /name, /state
*/
func (api *API) UpdateCatalog(catalogID string, patchPayload []*PatchOp) (string, error) {
	allowedFields := map[string][]string{
		"add":     []string{"/name", "/password", "/accessKey", "/secretKey", "/state"},
		"replace": []string{"/name", "/state"},
	}

	for _, pb := range patchPayload {
		fieldIsValid := false

		if allowedPaths, ok := allowedFields[pb.Op]; ok {
			for _, allowedPath := range allowedPaths {
				if pb.Path == allowedPath {
					fieldIsValid = true
				}
			}
		}

		if !fieldIsValid {
			return "", fmt.Errorf("UpdateCatalog received invalid Field for update.\nReceived Op: %s\nReceived Path: %s\nValid Fields: %v\n", pb.Op, pb.Path, allowedFields)
		}
	}

	return api.callHTTPRequest("PATCH", "/rest/catalogs/"+catalogID, nil, patchPayload)
}

// Connect App APIs

// os="windows" or os="mac"
func (api *API) GetConnectApp(os string) (string, error) {
	params := map[string]string{"os": os}
	return api.callHTTPRequest("GET", "/rest/connect-app", params, nil)
}

// Deployments APIs

func (api *API) GetDeployments(query, userQuery, view string) (string, error) {
	params := map[string]string{"query": query, "userQuery": userQuery, "view": view}
	return api.callHTTPRequest("GET", "/rest/deployments", params, nil)
}

func (api *API) CreateDeployment(info string) (string, error) {
	return api.callHTTPRequest("POST", "/rest/deployments", nil, info)
}

func (api *API) GetDeployment(deploymentID, view string) (string, error) {
	values := map[string]string{"view": view}
	return api.callHTTPRequest("GET", "/rest/deployments/"+deploymentID, nil, values)
}

func (api *API) UpdateDeployment(deploymentID, info string) (string, error) {
	return api.callHTTPRequest("PUT", "/rest/deployments/"+deploymentID, nil, info)
}

func (api *API) DeleteDeployment(deploymentID string) (string, error) {
	return api.callHTTPRequest("DELETE", "/rest/deployments/"+deploymentID, nil, nil)
}

func (api *API) ActionOnDeployment(deploymentID, actionType string, force bool) (string, error) {
	values := map[string]interface{}{"force": force, "type": actionType}
	return api.callHTTPRequest("POST", "/rest/deployments/"+deploymentID+"/actions", nil, values)
}

func (api *API) GetDeploymentConsole(deploymentID string) (string, error) {
	return api.callHTTPRequest("GET", "/rest/deployments/"+deploymentID+"/console", nil, nil)
}

// Events APIs

func (api *API) GetEvents(resourceUri string) (string, error) {
	params := map[string]string{"resourceUri": resourceUri}
	return api.callHTTPRequest("GET", "/rest/events", params, nil)
}

// Keypairs APIs

func (api *API) GetKeyPair(regionUri, projectUri string) (string, error) {
	params := map[string]string{"regionUri": regionUri, "projectUri": projectUri}
	return api.callHTTPRequest("GET", "/rest/keypairs", params, nil)
}

// Membership Roles APIs

func (api *API) GetMembershipRoles() (string, error) {
	return api.callHTTPRequest("GET", "/rest/membership-roles", nil, nil)
}

// Memberships APIs

func (api *API) GetMemberships(query string) (string, error) {
	params := map[string]string{"query": query}
	return api.callHTTPRequest("GET", "/rest/memberships", params, nil)
}

func (api *API) CreateMembership(userUri, roleUri, projectUri string) (string, error) {
	values := map[string]string{"userUri": userUri, "membershipRoleUri": roleUri, "projectUri": projectUri}
	return api.callHTTPRequest("POST", "/rest/memberships", nil, values)
}

func (api *API) DeleteMembership(userUri, roleUri, projectUri string) (string, error) {
	values := map[string]string{"userUri": userUri, "membershipRoleUri": roleUri, "projectUri": projectUri}
	return api.callHTTPRequest("DELETE", "/rest/memberships", nil, values)
}

// Metrics APIs

func (api *API) GetMetrics(
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
	return api.callHTTPRequest("GET", "/rest/metrics", params, nil)
}

// Networks APIs

func (api *API) GetNetworks(query string) (string, error) {
	params := map[string]string{"query": query}
	return api.callHTTPRequest("GET", "/rest/networks", params, nil)
}

func (api *API) GetNetwork(networkID string) (string, error) {
	return api.callHTTPRequest("GET", "/rest/networks/"+networkID, nil, nil)
}

// infoArray: [{op, path, value}]
func (api *API) UpdateNetwork(networkID string, infoArray []string) (string, error) {
	values := infoArray
	return api.callHTTPRequest("PUT", "/rest/networks/"+networkID, nil, values)
}

// Password Reset APIs

func (api *API) ResetSingleUsePassword(email string) (string, error) {
	values := map[string]string{"email": email}
	return api.callHTTPRequest("POST", "/rest/password-reset", nil, values)
}

func (api *API) ChangePassword(password, token string) (string, error) {
	values := map[string]string{"password": password, "token": token}
	return api.callHTTPRequest("POST", "/rest/password-reset/change", nil, values)
}

// Projects APIs

func (api *API) GetProjects(userQuery, view string) (string, error) {
	params := map[string]string{}
	if strings.TrimSpace(userQuery) != "" {
		params["userQuery"] = userQuery
	}
	if strings.TrimSpace(view) != "" {
		params["view"] = view
	}
	return api.callHTTPRequest("GET", "/rest/projects", params, nil)
}

func (api *API) CreateProject(name, description string, tagUris []string) (string, error) {
	values := map[string]interface{}{
		"name":        name,
		"description": description,
		"tagUris":     tagUris}
	return api.callHTTPRequest("POST", "/rest/projects", nil, values)
}

func (api *API) GetProject(projectID, view string) (string, error) {
	params := map[string]string{"view": view}
	return api.callHTTPRequest("GET", "/rest/projects/"+projectID, params, nil)
}

func (api *API) DeleteProject(projectID string) (string, error) {
	return api.callHTTPRequest("DELETE", "/rest/projects/"+projectID, nil, nil)
}

func (api *API) UpdateProject(projectID, name, description string, tagUris []string) (string, error) {
	values := map[string]interface{}{
		"name":        name,
		"description": description,
		"tagUris":     tagUris}
	return api.callHTTPRequest("PUT", "/rest/projects/"+projectID, nil, values)
}

// Provider Types APIs

func (api *API) GetProviderTypes() (string, error) {
	return api.callHTTPRequest("GET", "/rest/provider-types", nil, nil)
}

// Providers APIs

func (api *API) GetProviders(query string) (string, error) {
	params := map[string]string{"query": query}
	return api.callHTTPRequest("GET", "/rest/providers", params, nil)
}

// state: "Enabled|Disabled"
func (api *API) CreateProvider(providerID, providerTypeUri, accessKey, secretKey string,
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
	return api.callHTTPRequest("POST", "/rest/providers", nil, values)
}

// view="full"
// discover: boolean
func (api *API) GetProvider(providerID, view string, discover bool) (string, error) {
	params := map[string]string{
		"view":     view,
		"discover": strconv.FormatBool(discover)}
	return api.callHTTPRequest("GET", "/rest/providers/"+providerID, params, nil)
}

func (api *API) DeleteProvider(providerID string) (string, error) {
	return api.callHTTPRequest("DELETE", "/rest/providers/"+providerID, nil, nil)
}

// infoArray: [{op, path, value}]
// op: "add|replace|remove"
func (api *API) UpdateProvider(providerID, infoArray string) (string, error) {
	return api.callHTTPRequest("PUT", "/rest/providers/"+providerID, nil, infoArray)
}

// Rates APIs

func (api *API) GetRates(resourceUri, effectiveForDate, effectiveDate, metricName string,
	active bool, start, count int) (string, error) {
	params := map[string]string{
		"resourceUri":      resourceUri,
		"effectiveForDate": effectiveForDate,
		"effectiveDate":    effectiveDate,
		"metricName":       metricName,
		"active":           strconv.FormatBool(active),
		"start":            strconv.Itoa(start),
		"count":            strconv.Itoa(count)}
	return api.callHTTPRequest("GET", "/rest/rates", params, nil)
}

func (api *API) GetRate(rateID string) (string, error) {
	return api.callHTTPRequest("GET", "/rest/rates/"+rateID, nil, nil)
}

// Regions APIs

func (api *API) GetRegions(query, view string) (string, error) {
	params := map[string]string{"query": query, "view": view}
	return api.callHTTPRequest("GET", "/rest/regions", params, nil)
}

func (api *API) CreateRegion(name, providerUri, locLatitude, locLongitude string) (string, error) {
	values := map[string]interface{}{
		"location": map[string]interface{}{
			"latitude":  locLatitude,
			"longitude": locLongitude},
		"name":        name,
		"providerUri": providerUri}
	return api.callHTTPRequest("POST", "/rest/regions", nil, values)
}

func (api *API) GetRegion(regionID, view string, discover bool) (string, error) {
	params := map[string]string{"view": view, "discover": strconv.FormatBool(discover)}
	return api.callHTTPRequest("GET", "/rest/regions/"+regionID, params, nil)
}

func (api *API) DeleteRegion(regionID string, force bool) (string, error) {
	params := map[string]string{"force": strconv.FormatBool(force)}
	return api.callHTTPRequest("DELETE", "/rest/regions/"+regionID, params, nil)
}

// infoArray: [{op, path, value}]
// op: "add|replace"
// path: "/name|/location"
func (api *API) PatchRegion(regionID string, infoArray []string) (string, error) {
	return api.callHTTPRequest("PUT", "/rest/regions/"+regionID, nil, infoArray)
}

func (api *API) UpdateRegion(regionID, region string) (string, error) {
	return api.callHTTPRequest("PUT", "/rest/regions/"+regionID, nil, region)
}

func (api *API) GetRegionConnection(regionID string) (string, error) {
	return api.callHTTPRequest("GET", "/rest/regions/"+regionID+"/connection", nil, nil)
}

// state: "Enabling|Enabled|Disabling|Disabled"
func (api *API) CreateRegionConnection(regionID, endpointUuid, name, ipAddress, username, password string,
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
	return api.callHTTPRequest("POST", "/rest/regions/"+regionID+"/connection", nil, values)
}

func (api *API) DeleteRegionConnection(regionID string) (string, error) {
	return api.callHTTPRequest("DELETE", "/rest/regions/"+regionID+"/connection", nil, nil)
}

func (api *API) GetRegionConnectorImage(regionID string) (string, error) {
	return api.callHTTPRequest("GET", "/rest/regions/"+regionID+"/connector-image", nil, nil)
}

// Roles APIs

func (api *API) GetRoles() (string, error) {
	return api.callHTTPRequest("GET", "/rest/roles", nil, nil)
}

// Service Types APIs

func (api *API) GetServiceTypes() (string, error) {
	return api.callHTTPRequest("GET", "/rest/service-types", nil, nil)
}

func (api *API) GetServiceType(serviceTypeID string) (string, error) {
	return api.callHTTPRequest("GET", "/rest/service-types/"+serviceTypeID, nil, nil)
}

// Services APIs

func (api *API) GetServices(query, userQuery, view string) (string, error) {
	params := map[string]string{"query": query, "userQuery": userQuery, "view": view}
	return api.callHTTPRequest("GET", "/rest/services", params, nil)
}

// view: "full|deployment"
func (api *API) GetService(serviceID, view string) (string, error) {
	params := map[string]string{"view": view}
	return api.callHTTPRequest("GET", "/rest/services/"+serviceID, params, nil)
}

// Session APIs

// view: "full"
func (api *API) GetSession(view string) (string, error) {
	params := map[string]string{"view": view}
	return api.callHTTPRequest("GET", "/rest/session", params, nil)
}

func (api *API) GetSessionIdp(userName string) (string, error) {
	params := map[string]string{"userName": userName}
	return api.callHTTPRequest("GET", "/rest/session/idp", params, nil)
}

// GetStatus calls the /rest/status endpoint
func (api *API) GetStatus() (string, error) {
	return api.callHTTPRequest("GET", "/rest/status", nil, nil)
}

// Tag Keys APIs

// view: "full"
func (api *API) GetTagKeys(view string) (string, error) {
	params := map[string]string{"view": view}
	return api.callHTTPRequest("GET", "/rest/tag-keys", params, nil)
}

func (api *API) CreateTagKey(name string) (string, error) {
	values := map[string]string{"name": name}
	return api.callHTTPRequest("POST", "/rest/tag-keys", nil, values)
}

// view: "full"
func (api *API) GetTagKey(tagKeyID, view string) (string, error) {
	params := map[string]string{"view": view}
	return api.callHTTPRequest("GET", "/rest/tag-keys/"+tagKeyID, params, nil)
}

func (api *API) DeleteTagKey(tagKeyID string) (string, error) {
	return api.callHTTPRequest("DELETE", "/rest/tag-keys/"+tagKeyID, nil, nil)
}

// Tags APIs

// view: "full"
func (api *API) GetTags(view string) (string, error) {
	params := map[string]string{"view": view}
	return api.callHTTPRequest("GET", "/rest/tags", params, nil)
}

func (api *API) CreateTag(name, tagKeyUri string) (string, error) {
	values := map[string]string{"name": name, "tagKeyUri": tagKeyUri}
	return api.callHTTPRequest("POST", "/rest/tags", nil, values)
}

// view: "full"
func (api *API) GetTag(tagID, view string) (string, error) {
	params := map[string]string{"view": view}
	return api.callHTTPRequest("GET", "/rest/tags/"+tagID, params, nil)
}

func (api *API) DeleteTag(tagID string) (string, error) {
	return api.callHTTPRequest("DELETE", "/rest/tags/"+tagID, nil, nil)
}

// Users APIs

func (api *API) GetUsers(userQuery string) (string, error) {
	params := map[string]string{"userQuery": userQuery}
	return api.callHTTPRequest("GET", "/rest/users", params, nil)
}

// role: "administrator|analyst|consumer|project-creator"
func (api *API) CreateUser(email, name, password, role string) (string, error) {
	values := map[string]string{"email": email, "name": name, "password": password, "role": role}
	return api.callHTTPRequest("POST", "/rest/users", nil, values)
}

func (api *API) GetUser(userID string) (string, error) {
	return api.callHTTPRequest("GET", "/rest/users/"+userID, nil, nil)
}

// role: "administrator|analyst|consumer|project-creator"
func (api *API) UpdateUser(userID, email, name, password, role string) (string, error) {
	values := map[string]string{"email": email, "name": name, "password": password, "role": role}
	return api.callHTTPRequest("PUT", "/rest/users/"+userID, nil, values)
}

func (api *API) DeleteUser(userID string) (string, error) {
	return api.callHTTPRequest("DELETE", "/rest/users/"+userID, nil, nil)
}

// Virtual Machine Profiles APIs

func (api *API) GetVirtualMachineProfiles(zoneUri, serviceUri string) (string, error) {
	params := map[string]string{"zoneUri": zoneUri, "serviceUri": serviceUri}
	return api.callHTTPRequest("GET", "/rest/virtual-machine-profiles", params, nil)
}

func (api *API) GetVirtualMachineProfile(vmProfileID string) (string, error) {
	return api.callHTTPRequest("GET", "/rest/virtual-machine-profiles/"+vmProfileID, nil, nil)
}

// Volumes APIs

// view: "full"
func (api *API) GetVolumes(query, view string) (string, error) {
	params := map[string]string{"query": query, "view": view}
	return api.callHTTPRequest("GET", "/rest/volumes", params, nil)
}

func (api *API) CreateVolume(name string, sizeGiB int, zoneUri, projectUri string) (string, error) {
	values := map[string]interface{}{
		"name":       name,
		"sizeGiB":    strconv.Itoa(sizeGiB),
		"zoneUri":    zoneUri,
		"projectUri": projectUri}
	return api.callHTTPRequest("POST", "/rest/volumes", nil, values)
}

func (api *API) GetVolume(volumeID string) (string, error) {
	return api.callHTTPRequest("GET", "/rest/volumes/"+volumeID, nil, nil)
}

func (api *API) UpdateVolume(volumeID, name string, sizeGiB int) (string, error) {
	values := map[string]interface{}{
		"name":    name,
		"sizeGiB": strconv.Itoa(sizeGiB)}
	return api.callHTTPRequest("PUT", "/rest/volumes/"+volumeID, nil, values)
}

func (api *API) DeleteVolume(volumeID string) (string, error) {
	return api.callHTTPRequest("DELETE", "/rest/volumes/"+volumeID, nil, nil)
}

// Zone Types APIs

func (api *API) GetZoneTypes() (string, error) {
	return api.callHTTPRequest("GET", "/rest/zone-types", nil, nil)
}

func (api *API) GetZoneTypeResourceProfiles(zoneTypeID string) (string, error) {
	return api.callHTTPRequest("GET", "/rest/zone-types/"+zoneTypeID+"/resource-profiles", nil, nil)
}

// Zones APIs

func (api *API) GetZones(query, regionUri, applianceUri string) (string, error) {
	params := map[string]string{"query": query, "regionUri": regionUri, "applianceUri": applianceUri}
	return api.callHTTPRequest("GET", "/rest/zones", params, nil)
}

func (api *API) CreateZone(zoneData string) (string, error) {
	return api.callHTTPRequest("POST", "/rest/zones", nil, zoneData)
}

// view: "full"
func (api *API) GetZone(zoneID, view string) (string, error) {
	params := map[string]string{"view": view}
	return api.callHTTPRequest("GET", "/rest/zones/"+zoneID, params, nil)
}

// op: "add|replace|remove"
func (api *API) UpdateZone(zoneID, op, path string, value interface{}) (string, error) {
	values := map[string]interface{}{"op": op, "path": path, "value": value}
	return api.callHTTPRequest("PUT", "/rest/zones/"+zoneID, nil, values)
}

func (api *API) DeleteZone(zoneID string, force bool) (string, error) {
	params := map[string]string{"force": strconv.FormatBool(force)}
	return api.callHTTPRequest("DELETE", "/rest/zones/"+zoneID, params, nil)
}

// actionType: "reset|add-capacity|reduce-capacity"
// resourceType: "compute|storage"
func (api *API) ActionOnZone(zoneID, actionType, resourceType string, resourceCapacity int) (string, error) {
	values := map[string]interface{}{
		"type": actionType,
		"resourceOp": map[string]interface{}{
			"resourceType":     resourceType,
			"resourceCapacity": resourceCapacity}}
	return api.callHTTPRequest("POST", "/rest/zones/"+zoneID+"/actions", nil, values)
}

func (api *API) GetZoneApplianceImage(zoneID string) (string, error) {
	return api.callHTTPRequest("GET", "/rest/zones/"+zoneID+"/appliance-image", nil, nil)
}

func (api *API) GetZoneTaskStatus(zoneID string) (string, error) {
	return api.callHTTPRequest("GET", "/rest/zones/"+zoneID+"/task-status", nil, nil)
}

func (api *API) GetZoneConnections(zoneID, uuid string) (string, error) {
	params := map[string]string{"uuid": uuid}
	return api.callHTTPRequest("GET", "/rest/zones/"+zoneID+"/connections", params, nil)
}

// state: "Enabling|Enabled|Disabling|Disabled"
func (api *API) CreateZoneConnection(zoneID, uuid, name, ipAddress, username, password string,
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
	return api.callHTTPRequest("POST", "/rest/zones/"+zoneID+"/connections", nil, values)
}

func (api *API) DeleteZoneConnection(zoneID, uuid string) (string, error) {
	return api.callHTTPRequest("DELETE", "/rest/zones/"+zoneID+"/connections/"+uuid, nil, nil)
}

// op: "add|replace|remove"
func (api *API) UpdateZoneConnection(zoneID, uuid, op, path string, value interface{}) (string, error) {
	values := map[string]interface{}{"op": op, "path": path, "value": value}
	return api.callHTTPRequest("PUT", "/rest/zones/"+zoneID+"/connections", nil, values)
}
