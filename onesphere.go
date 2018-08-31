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
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var Token string
var HostUrl string

func callHttpRequest(method, url string, params map[string]string, values interface{}) string {
	jsonValue, err := json.Marshal(values)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", Token)

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
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	bodyStr := string(bodyBytes)
	return bodyStr
}

func Connect(hostUrl, user, password string) error {
	HostUrl = hostUrl
	fullUrl := hostUrl + "/rest/session"
	values := map[string]string{"userName": user, "password": password}
	jsonValue, err := json.Marshal(values)
	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	//bodyStr := string(body)
	var dat map[string]string
	err = json.Unmarshal(body, &dat)
	if err != nil {
		return err
	}

	Token = dat["token"]
	return nil
}

func Disconnect() {
	fullUrl := HostUrl + "/rest/session"
	callHttpRequest("DELETE", fullUrl, nil, nil)
}

// Account APIs

// view="full"
func GetAccount(view string) string {
	fullUrl := HostUrl + "/rest/account"
	params := map[string]string{"view": view}
	return callHttpRequest("GET", fullUrl, params, nil)
}

// Appliances APIs

func GetAppliances(name, regionUri string) string {
	fullUrl := HostUrl + "/rest/appliances"
	params := map[string]string{}
	if strings.TrimSpace(name) != "" {
		params["name"] = name
	}
	if strings.TrimSpace(regionUri) != "" {
		params["regionUri"] = regionUri
	}
	return callHttpRequest("GET", fullUrl, params, nil)
}

func CreateAppliance(epAddress, epUsername, epPassword,
	name, regionUri, applianceType string) string {
	fullUrl := HostUrl + "/rest/appliances"
	values := map[string]interface{}{
		"endpoint": map[string]interface{}{
			"address":  epAddress,
			"password": epPassword,
			"username": epUsername},
		"name":      name,
		"regionUri": regionUri,
		"type":      applianceType}
	return callHttpRequest("POST", fullUrl, nil, values)
}

func GetAppliance(applianceID string) string {
	fullUrl := HostUrl + "/rest/appliances/" + applianceID
	return callHttpRequest("GET", fullUrl, nil, nil)
}

func DeleteAppliance(applianceID string) string {
	fullUrl := HostUrl + "/rest/appliances/" + applianceID
	return callHttpRequest("DELETE", fullUrl, nil, nil)
}

// infoArray: [{op, path, value}]
// op: "replace|remove"
func UpdateAppliance(applianceID string, infoArray []string) string {
	fullUrl := HostUrl + "/rest/appliances/" + applianceID
	values := infoArray
	return callHttpRequest("PUT", fullUrl, nil, values)
}

// Catalog Types APIs

func GetCatalogTypes() string {
	fullUrl := HostUrl + "/rest/catalog-types"
	return callHttpRequest("GET", fullUrl, nil, nil)
}

// Catalogs APIs

func GetCatalogs(userQuery, view string) string {
	fullUrl := HostUrl + "/rest/catalogs"
	params := map[string]string{}
	if strings.TrimSpace(userQuery) != "" {
		params["userQuery"] = userQuery
	}
	if strings.TrimSpace(view) != "" {
		params["view"] = view
	}
	return callHttpRequest("GET", fullUrl, params, nil)
}

func CreateCatalog(accessKey, catalogTypeUri, name, password, regionName, secretKey, url, username string) string {
	fullUrl := HostUrl + "/rest/catalogs"
	values := map[string]string{
		"accessKey":      accessKey,
		"catalogTypeUri": catalogTypeUri,
		"name":           name,
		"password":       password,
		"regionName":     regionName,
		"secretKey":      secretKey,
		"url":            url,
		"username":       username}
	return callHttpRequest("POST", fullUrl, nil, values)
}

func GetCatalog(catalogID, view string) string {
	fullUrl := HostUrl + "/rest/catalogs/" + catalogID
	params := map[string]string{"view": view}
	return callHttpRequest("GET", fullUrl, params, nil)
}

func DeleteCatalog(catalogID string) string {
	fullUrl := HostUrl + "/rest/catalogs/" + catalogID
	return callHttpRequest("DELETE", fullUrl, nil, nil)
}

func UpdateCatalog(catalogID, name, password, accessKey, secretKey, regionName, state string) string {
	fullUrl := HostUrl + "/rest/catalogs/" + catalogID
	values := map[string]interface{}{
		"name": name, "password": password, "accessKey": accessKey,
		"secretKey": secretKey, "regionName": regionName, "state": state}
	return callHttpRequest("PUT", fullUrl, nil, values)
}

// Connect App APIs

// os="windows" or os="mac"
func GetConnectApp(os string) string {
	fullUrl := HostUrl + "/rest/connect-app"
	params := map[string]string{"os": os}
	return callHttpRequest("GET", fullUrl, params, nil)
}

// Deployments APIs

func GetDeployments(query, userQuery, view string) string {
	fullUrl := HostUrl + "/rest/deployments"
	params := map[string]string{"query": query, "userQuery": userQuery, "view": view}
	return callHttpRequest("GET", fullUrl, params, nil)
}

func CreateDeployment(info string) string {
	fullUrl := HostUrl + "/rest/deployments"
	return callHttpRequest("POST", fullUrl, nil, info)
}

func GetDeployment(deploymentID, view string) string {
	fullUrl := HostUrl + "/rest/deployments/" + deploymentID
	values := map[string]string{"view": view}
	return callHttpRequest("GET", fullUrl, nil, values)
}

func UpdateDeployment(deploymentID, info string) string {
	fullUrl := HostUrl + "/rest/deployments/" + deploymentID
	return callHttpRequest("PUT", fullUrl, nil, info)
}

func DeleteDeployment(deploymentID string) string {
	fullUrl := HostUrl + "/rest/deployments/" + deploymentID
	return callHttpRequest("DELETE", fullUrl, nil, nil)
}

func ActionOnDeployment(deploymentID, actionType string, force bool) string {
	fullUrl := HostUrl + "/rest/deployments/" + deploymentID + "/actions"
	values := map[string]interface{}{"force": force, "type": actionType}
	return callHttpRequest("POST", fullUrl, nil, values)
}

func GetDeploymentConsole(deploymentID string) string {
	fullUrl := HostUrl + "/rest/deployments/" + deploymentID + "/console"
	return callHttpRequest("GET", fullUrl, nil, nil)
}

// Events APIs

func GetEvents(resourceUri string) string {
	fullUrl := HostUrl + "/rest/events"
	params := map[string]string{"resourceUri": resourceUri}
	return callHttpRequest("GET", fullUrl, params, nil)
}

// Keypairs APIs

func GetKeyPair(regionUri, projectUri string) string {
	fullUrl := HostUrl + "/rest/keypairs"
	params := map[string]string{"regionUri": regionUri, "projectUri": projectUri}
	return callHttpRequest("GET", fullUrl, params, nil)
}

// Membership Roles APIs

func GetMembershipRoles() string {
	fullUrl := HostUrl + "/rest/membership-roles"
	return callHttpRequest("GET", fullUrl, nil, nil)
}

// Memberships APIs

func GetMemberships(query string) string {
	fullUrl := HostUrl + "/rest/memberships"
	params := map[string]string{"query": query}
	return callHttpRequest("GET", fullUrl, params, nil)
}

func CreateMembership(userUri, roleUri, projectUri string) string {
	fullUrl := HostUrl + "/rest/memberships"
	values := map[string]string{"userUri": userUri, "membershipRoleUri": roleUri, "projectUri": projectUri}
	return callHttpRequest("POST", fullUrl, nil, values)
}

func DeleteMembership(userUri, roleUri, projectUri string) string {
	fullUrl := HostUrl + "/rest/memberships"
	values := map[string]string{"userUri": userUri, "membershipRoleUri": roleUri, "projectUri": projectUri}
	return callHttpRequest("DELETE", fullUrl, nil, values)
}

// Metrics APIs

func GetMetrics(
	resourceUri, category, groupBy, query, name string,
	periodStart, period string,
	periodCount int,
	view string,
	start, count int) string {
	fullUrl := HostUrl + "/rest/metrics"
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
	return callHttpRequest("GET", fullUrl, params, nil)
}

// Networks APIs

func GetNetworks(query string) string {
	fullUrl := HostUrl + "/rest/networks"
	params := map[string]string{"query": query}
	return callHttpRequest("GET", fullUrl, params, nil)
}

func GetNetwork(networkID string) string {
	fullUrl := HostUrl + "/rest/networks/" + networkID
	return callHttpRequest("GET", fullUrl, nil, nil)
}

// infoArray: [{op, path, value}]
func UpdateNetwork(networkID string, infoArray []string) string {
	fullUrl := HostUrl + "/rest/networks/" + networkID
	values := infoArray
	return callHttpRequest("PUT", fullUrl, nil, values)
}

// Password Reset APIs

func ResetSingleUsePassword(email string) string {
	fullUrl := HostUrl + "/rest/password-reset"
	values := map[string]string{"email": email}
	return callHttpRequest("POST", fullUrl, nil, values)
}

func ChangePassword(password, token string) string {
	fullUrl := HostUrl + "/rest/password-reset/change"
	values := map[string]string{"password": password, "token": token}
	return callHttpRequest("POST", fullUrl, nil, values)
}

// Projects APIs

func GetProjects(userQuery, view string) string {
	fullUrl := HostUrl + "/rest/projects"
	params := map[string]string{}
	if strings.TrimSpace(userQuery) != "" {
		params["userQuery"] = userQuery
	}
	if strings.TrimSpace(view) != "" {
		params["view"] = view
	}
	return callHttpRequest("GET", fullUrl, params, nil)
}

func CreateProject(name, description string, tagUris []string) string {
	fullUrl := HostUrl + "/rest/projects"
	values := map[string]interface{}{
		"name":        name,
		"description": description,
		"tagUris":     tagUris}
	return callHttpRequest("POST", fullUrl, nil, values)
}

func GetProject(projectID, view string) string {
	fullUrl := HostUrl + "/rest/projects/" + projectID
	params := map[string]string{"view": view}
	return callHttpRequest("GET", fullUrl, params, nil)
}

func DeleteProject(projectID string) string {
	fullUrl := HostUrl + "/rest/projects/" + projectID
	return callHttpRequest("DELETE", fullUrl, nil, nil)
}

func UpdateProject(projectID, name, description string, tagUris []string) string {
	fullUrl := HostUrl + "/rest/projects/" + projectID
	values := map[string]interface{}{
		"name":        name,
		"description": description,
		"tagUris":     tagUris}
	return callHttpRequest("PUT", fullUrl, nil, values)
}

// Provider Types APIs

func GetProviderTypes() string {
	fullUrl := HostUrl + "/rest/provider-types"
	return callHttpRequest("GET", fullUrl, nil, nil)
}

// Providers APIs

func GetProviders(query string) string {
	fullUrl := HostUrl + "/rest/providers"
	params := map[string]string{"query": query}
	return callHttpRequest("GET", fullUrl, params, nil)
}

// state: "Enabled|Disabled"
func CreateProvider(providerID, providerTypeUri, accessKey, secretKey string,
	paymentProvider bool,
	s3CostBucket, masterUri, state string) string {
	fullUrl := HostUrl + "/rest/providers"
	values := map[string]interface{}{
		"id":              providerID,
		"providerTypeUri": providerTypeUri,
		"accessKey":       accessKey,
		"secretKey":       secretKey,
		"paymentProvider": paymentProvider,
		"s3CostBucket":    s3CostBucket,
		"masterUri":       masterUri,
		"state":           state}
	return callHttpRequest("POST", fullUrl, nil, values)
}

// view="full"
// discover: boolean
func GetProvider(providerID, view string, discover bool) string {
	fullUrl := HostUrl + "/rest/providers/" + providerID
	params := map[string]string{
		"view":     view,
		"discover": strconv.FormatBool(discover)}
	return callHttpRequest("GET", fullUrl, params, nil)
}

func DeleteProvider(providerID string) string {
	fullUrl := HostUrl + "/rest/providers/" + providerID
	return callHttpRequest("DELETE", fullUrl, nil, nil)
}

// infoArray: [{op, path, value}]
// op: "add|replace|remove"
func UpdateProvider(providerID, infoArray string) string {
	fullUrl := HostUrl + "/rest/providers/" + providerID
	return callHttpRequest("PUT", fullUrl, nil, infoArray)
}

// Rates APIs

func GetRates(resourceUri, effectiveForDate, effectiveDate, metricName string,
	active bool, start, count int) string {
	fullUrl := HostUrl + "/rest/rates"
	params := map[string]string{
		"resourceUri":      resourceUri,
		"effectiveForDate": effectiveForDate,
		"effectiveDate":    effectiveDate,
		"metricName":       metricName,
		"active":           strconv.FormatBool(active),
		"start":            strconv.Itoa(start),
		"count":            strconv.Itoa(count)}
	return callHttpRequest("GET", fullUrl, params, nil)
}

func GetRate(rateID string) string {
	fullUrl := HostUrl + "/rest/rates/" + rateID
	return callHttpRequest("GET", fullUrl, nil, nil)
}

// Regions APIs

func GetRegions(query, view string) string {
	fullUrl := HostUrl + "/rest/regions"
	params := map[string]string{"query": query, "view": view}
	return callHttpRequest("GET", fullUrl, params, nil)
}

func CreateRegion(name, providerUri, locLatitude, locLongitude string) string {
	fullUrl := HostUrl + "/rest/regions"
	values := map[string]interface{}{
		"location": map[string]interface{}{
			"latitude":  locLatitude,
			"longitude": locLongitude},
		"name":        name,
		"providerUri": providerUri}
	return callHttpRequest("POST", fullUrl, nil, values)
}

func GetRegion(regionID, view string, discover bool) string {
	fullUrl := HostUrl + "/rest/regions/" + regionID
	params := map[string]string{"view": view, "discover": strconv.FormatBool(discover)}
	return callHttpRequest("GET", fullUrl, params, nil)
}

func DeleteRegion(regionID string, force bool) string {
	fullUrl := HostUrl + "/rest/regions/" + regionID
	params := map[string]string{"force": strconv.FormatBool(force)}
	return callHttpRequest("DELETE", fullUrl, params, nil)
}

// infoArray: [{op, path, value}]
// op: "add|replace"
// path: "/name|/location"
func PatchRegion(regionID string, infoArray []string) string {
	fullUrl := HostUrl + "/rest/regions/" + regionID
	return callHttpRequest("PUT", fullUrl, nil, infoArray)
}

func UpdateRegion(regionID, region string) string {
	fullUrl := HostUrl + "/rest/regions/" + regionID
	return callHttpRequest("PUT", fullUrl, nil, region)
}

func GetRegionConnection(regionID string) string {
	fullUrl := HostUrl + "/rest/regions/" + regionID + "/connection"
	return callHttpRequest("GET", fullUrl, nil, nil)
}

// state: "Enabling|Enabled|Disabling|Disabled"
func CreateRegionConnection(regionID, endpointUuid, name, ipAddress, username, password string,
	port int,
	state, uri string) string {
	fullUrl := HostUrl + "/rest/regions/" + regionID + "/connection"
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
	return callHttpRequest("POST", fullUrl, nil, values)
}

func DeleteRegionConnection(regionID string) string {
	fullUrl := HostUrl + "/rest/regions/" + regionID + "/connection"
	return callHttpRequest("DELETE", fullUrl, nil, nil)
}

func GetRegionConnectorImage(regionID string) string {
	fullUrl := HostUrl + "/rest/regions/" + regionID + "/connector-image"
	return callHttpRequest("GET", fullUrl, nil, nil)
}

// Roles APIs

func GetRoles() string {
	fullUrl := HostUrl + "/rest/roles"
	return callHttpRequest("GET", fullUrl, nil, nil)
}

// Service Types APIs

func GetServiceTypes() string {
	fullUrl := HostUrl + "/service-types"
	return callHttpRequest("GET", fullUrl, nil, nil)
}

func GetServiceType(serviceTypeID string) string {
	fullUrl := HostUrl + "/service-types/" + serviceTypeID
	return callHttpRequest("GET", fullUrl, nil, nil)
}

// Services APIs

func GetServices(query, userQuery, view string) string {
	fullUrl := HostUrl + "/rest/services"
	params := map[string]string{"query": query, "userQuery": userQuery, "view": view}
	return callHttpRequest("GET", fullUrl, params, nil)
}

// view: "full|deployment"
func GetService(serviceID, view string) string {
	fullUrl := HostUrl + "/rest/services/" + serviceID
	params := map[string]string{"view": view}
	return callHttpRequest("GET", fullUrl, params, nil)
}

// Session APIs

// view: "full"
func GetSession(view string) string {
	fullUrl := HostUrl + "/rest/session"
	params := map[string]string{"view": view}
	return callHttpRequest("GET", fullUrl, params, nil)
}

func GetSessionIdp(userName string) string {
	fullUrl := HostUrl + "/rest/session/idp"
	params := map[string]string{"userName": userName}
	return callHttpRequest("GET", fullUrl, params, nil)
}

// Status APIs

func GetStatus() string {
	fullUrl := HostUrl + "/rest/status"
	return callHttpRequest("GET", fullUrl, nil, nil)
}

// Tag Keys APIs

// view: "full"
func GetTagKeys(view string) string {
	fullUrl := HostUrl + "/rest/tag-keys"
	params := map[string]string{"view": view}
	return callHttpRequest("GET", fullUrl, params, nil)
}

func CreateTagKey(name string) string {
	fullUrl := HostUrl + "/rest/tag-keys"
	values := map[string]string{"name": name}
	return callHttpRequest("POST", fullUrl, nil, values)
}

// view: "full"
func GetTagKey(tagKeyID, view string) string {
	fullUrl := HostUrl + "/rest/tag-keys/" + tagKeyID
	params := map[string]string{"view": view}
	return callHttpRequest("GET", fullUrl, params, nil)
}

func DeleteTagKey(tagKeyID string) string {
	fullUrl := HostUrl + "/rest/tag-keys/" + tagKeyID
	return callHttpRequest("DELETE", fullUrl, nil, nil)
}

// Tags APIs

// view: "full"
func GetTags(view string) string {
	fullUrl := HostUrl + "/rest/tags"
	params := map[string]string{"view": view}
	return callHttpRequest("GET", fullUrl, params, nil)
}

func CreateTag(name, tagKeyUri string) string {
	fullUrl := HostUrl + "/rest/tags"
	values := map[string]string{"name": name, "tagKeyUri": tagKeyUri}
	return callHttpRequest("POST", fullUrl, nil, values)
}

// view: "full"
func GetTag(tagID, view string) string {
	fullUrl := HostUrl + "/rest/tags/" + tagID
	params := map[string]string{"view": view}
	return callHttpRequest("GET", fullUrl, params, nil)
}

func DeleteTag(tagID string) string {
	fullUrl := HostUrl + "/rest/tags/" + tagID
	return callHttpRequest("DELETE", fullUrl, nil, nil)
}

// Users APIs

func GetUsers(userQuery string) string {
	fullUrl := HostUrl + "/rest/users"
	params := map[string]string{"userQuery": userQuery}
	return callHttpRequest("GET", fullUrl, params, nil)
}

// role: "administrator|analyst|consumer|project-creator"
func CreateUser(email, name, password, role string) string {
	fullUrl := HostUrl + "/rest/users"
	values := map[string]string{"email": email, "name": name, "password": password, "role": role}
	return callHttpRequest("POST", fullUrl, nil, values)
}

func GetUser(userID string) string {
	fullUrl := HostUrl + "/rest/users/" + userID
	return callHttpRequest("GET", fullUrl, nil, nil)
}

// role: "administrator|analyst|consumer|project-creator"
func UpdateUser(userID, email, name, password, role string) string {
	fullUrl := HostUrl + "/rest/users/" + userID
	values := map[string]string{"email": email, "name": name, "password": password, "role": role}
	return callHttpRequest("PUT", fullUrl, nil, values)
}

func DeleteUser(userID string) string {
	fullUrl := HostUrl + "/rest/users/" + userID
	return callHttpRequest("DELETE", fullUrl, nil, nil)
}

// Virtual Machine Profiles APIs

func GetVirtualMachineProfiles(zoneUri, serviceUri string) string {
	fullUrl := HostUrl + "/rest/virtual-machine-profiles"
	params := map[string]string{"zoneUri": zoneUri, "serviceUri": serviceUri}
	return callHttpRequest("GET", fullUrl, params, nil)
}

func GetVirtualMachineProfile(vmProfileID string) string {
	fullUrl := HostUrl + "/rest/virtual-machine-profiles/" + vmProfileID
	return callHttpRequest("GET", fullUrl, nil, nil)
}

// Volumes APIs

// view: "full"
func GetVolumes(query, view string) string {
	fullUrl := HostUrl + "/rest/volumes"
	params := map[string]string{"query": query, "view": view}
	return callHttpRequest("GET", fullUrl, params, nil)
}

func CreateVolume(name string, sizeGiB int, zoneUri, projectUri string) string {
	fullUrl := HostUrl + "/rest/volumes"
	values := map[string]interface{}{
		"name":       name,
		"sizeGiB":    strconv.Itoa(sizeGiB),
		"zoneUri":    zoneUri,
		"projectUri": projectUri}
	return callHttpRequest("POST", fullUrl, nil, values)
}

func GetVolume(volumeID string) string {
	fullUrl := HostUrl + "/rest/volumes/" + volumeID
	return callHttpRequest("GET", fullUrl, nil, nil)
}

func UpdateVolume(volumeID, name string, sizeGiB int) string {
	fullUrl := HostUrl + "/rest/volumes/" + volumeID
	values := map[string]interface{}{
		"name":    name,
		"sizeGiB": strconv.Itoa(sizeGiB)}
	return callHttpRequest("PUT", fullUrl, nil, values)
}

func DeleteVolume(volumeID string) string {
	fullUrl := HostUrl + "/rest/volumes/" + volumeID
	return callHttpRequest("DELETE", fullUrl, nil, nil)
}

// Zone Types APIs

func GetZoneTypes() string {
	fullUrl := HostUrl + "/rest/zone-types"
	return callHttpRequest("GET", fullUrl, nil, nil)
}

func GetZoneTypeResourceProfiles(zoneTypeID string) string {
	fullUrl := HostUrl + "/rest/zone-types/" + zoneTypeID + "/resource-profiles"
	return callHttpRequest("GET", fullUrl, nil, nil)
}

// Zones APIs

func GetZones(query, regionUri, applianceUri string) string {
	fullUrl := HostUrl + "/rest/zones"
	params := map[string]string{"query": query, "regionUri": regionUri, "applianceUri": applianceUri}
	return callHttpRequest("GET", fullUrl, params, nil)
}

func CreateZone(zoneData string) string {
	fullUrl := HostUrl + "/rest/zones"
	return callHttpRequest("POST", fullUrl, nil, zoneData)
}

// view: "full"
func GetZone(zoneID, view string) string {
	fullUrl := HostUrl + "/rest/zones/" + zoneID
	params := map[string]string{"view": view}
	return callHttpRequest("GET", fullUrl, params, nil)
}

// op: "add|replace|remove"
func UpdateZone(zoneID, op, path string, value interface{}) string {
	fullUrl := HostUrl + "/rest/zones/" + zoneID
	values := map[string]interface{}{"op": op, "path": path, "value": value}
	return callHttpRequest("PUT", fullUrl, nil, values)
}

func DeleteZone(zoneID string, force bool) string {
	fullUrl := HostUrl + "/rest/zones/" + zoneID
	params := map[string]string{"force": strconv.FormatBool(force)}
	return callHttpRequest("DELETE", fullUrl, params, nil)
}

// actionType: "reset|add-capacity|reduce-capacity"
// resourceType: "compute|storage"
func ActionOnZone(zoneID, actionType, resourceType string, resourceCapacity int) string {
	fullUrl := HostUrl + "/rest/zones/" + zoneID + "/actions"
	values := map[string]interface{}{
		"type": actionType,
		"resourceOp": map[string]interface{}{
			"resourceType":     resourceType,
			"resourceCapacity": resourceCapacity}}
	return callHttpRequest("POST", fullUrl, nil, values)
}

func GetZoneApplianceImage(zoneID string) string {
	fullUrl := HostUrl + "/rest/zones/" + zoneID + "/appliance-image"
	return callHttpRequest("GET", fullUrl, nil, nil)
}

func GetZoneTaskStatus(zoneID string) string {
	fullUrl := HostUrl + "/rest/zones/" + zoneID + "/task-status"
	return callHttpRequest("GET", fullUrl, nil, nil)
}

func GetZoneConnections(zoneID, uuid string) string {
	fullUrl := HostUrl + "/rest/zones/" + zoneID + "/connections"
	params := map[string]string{"uuid": uuid}
	return callHttpRequest("GET", fullUrl, params, nil)
}

// state: "Enabling|Enabled|Disabling|Disabled"
func CreateZoneConnection(zoneID, uuid, name, ipAddress, username, password string,
	port int, state string) string {
	fullUrl := HostUrl + "/rest/zones/" + zoneID + "/connections"
	values := map[string]interface{}{
		"uuid": uuid,
		"name": name,
		"location": map[string]interface{}{
			"ipAddress": ipAddress,
			"username":  username,
			"password":  password,
			"port":      port},
		"state": state}
	return callHttpRequest("POST", fullUrl, nil, values)
}

func DeleteZoneConnection(zoneID, uuid string) string {
	fullUrl := HostUrl + "/rest/zones/" + zoneID + "/connections/" + uuid
	return callHttpRequest("DELETE", fullUrl, nil, nil)
}

// op: "add|replace|remove"
func UpdateZoneConnection(zoneID, uuid, op, path string, value interface{}) string {
	fullUrl := HostUrl + "/rest/zones/" + zoneID + "/connections"
	values := map[string]interface{}{"op": op, "path": path, "value": value}
	return callHttpRequest("PUT", fullUrl, nil, values)
}
