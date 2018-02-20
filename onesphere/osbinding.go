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

package osbinding

import (
    //"fmt"
    "strings"
    "strconv"
    "bytes"
    //"io"
    "io/ioutil"
    "net/http"
    //"net/url"
    "encoding/json"
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

func Connect(hostUrl, user, password string) {
    HostUrl = hostUrl
    fullUrl := hostUrl + "/rest/session"
    values := map[string]string{"userName": user, "password": password}
    jsonValue, err := json.Marshal(values)
    req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(jsonValue))
    if err != nil {
        panic(err)
    }
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    //bodyStr := string(body)
    var dat map[string]string
    err = json.Unmarshal(body, &dat)
    if err != nil {
        panic(err)
    }

    Token = dat["token"]
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
            "address": epAddress,
            "password": epPassword,
            "username": epUsername},
        "name": name,
        "regionUri": regionUri,
        "type": applianceType}
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

// infoArray: [{os, path, value}]
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
                "accessKey": accessKey,
                "catalogTypeUri": catalogTypeUri,
                "name": name,
                "password": password,
                "regionName": regionName,
                "secretKey": secretKey,
                "url": url,
                "username": username}
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
        "category": category,
        "groupBy": groupBy,
        "query": query,
        "nameArray": name,
        "periodStart": periodStart,
        "period": period,
        "periodCount": strconv.Itoa(periodCount),
        "view": view,
        "start": strconv.Itoa(start),
        "count": strconv.Itoa(count)}
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
                "name": name,
                "description": description,
                "tagUris": tagUris}
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
                "name": name,
                "description": description,
                "tagUris": tagUris}
    return callHttpRequest("PUT", fullUrl, nil, values)
}

// Provider Types APIs

func GetProviderTypes() string {
    fullUrl := HostUrl + "/rest/provider-types"
    return callHttpRequest("GET", fullUrl, nil, nil)
}

// Providers APIs

func GetProviders(parentUri, providerTypeUri string) string {
    fullUrl := HostUrl + "/rest/providers"
    values := map[string]string{"parentUri": parentUri, "providerTypeUri": providerTypeUri}
    return callHttpRequest("GET", fullUrl, nil, values)
}

func CreateProvider(providerID, providerTypeUri, accessKey, secretKey, 
                    s3CostBucket, parentUri, state string, 
                    paymentProvider bool) string {
    fullUrl := HostUrl + "/rest/providers"
    values := map[string]interface{}{
        "id": providerID,
        "providerTypeUri": providerTypeUri,
        "accessKey": accessKey,
        "secretKey": secretKey,
        "paymentProvider": paymentProvider,
        "s3CostBucket": s3CostBucket,
        "parentUri": parentUri,
        "state": state}
    return callHttpRequest("POST", fullUrl, nil, values)
}

// view="full"
func GetProvider(providerID, view string, discover bool) string {
    fullUrl := HostUrl + "/rest/providers/" + providerID
    values := map[string]interface{}{
        "view": view,
        "discover": discover}
    return callHttpRequest("GET", fullUrl, nil, values)
}

func DeleteProvider(providerID string) string {
    fullUrl := HostUrl + "/rest/providers/" + providerID
    return callHttpRequest("DELETE", fullUrl, nil, nil)
}

func UpdateProvider(providerID, infoArray string) string {
    fullUrl := HostUrl + "/rest/providers/" + providerID
    return callHttpRequest("PUT", fullUrl, nil, infoArray)
}

// Status APIs

func GetStatus() string {
    fullUrl := HostUrl + "/rest/status"
    return callHttpRequest("GET", fullUrl, nil, nil)
}

// Session APIs

// view="full"
func GetSession(view string) string {
    fullUrl := HostUrl + "/rest/session"
    params := map[string]string{"view": view}
    return callHttpRequest("GET", fullUrl, params, nil)
}

func GetSessionIdp(username string) string {
    fullUrl := HostUrl + "/rest/session/idp"
    values := map[string]string{"userName": username}
    return callHttpRequest("GET", fullUrl, nil, values)
}

// Regions APIs

func GetRegions(providerUri, view string) string {
    fullUrl := HostUrl + "/rest/regions"
    values := map[string]string{"providerUri": providerUri, "view": view}
    return callHttpRequest("GET", fullUrl, nil, values)
}

func CreateRegion(name, providerUri, locLatitude, locLongitude string) string {
    fullUrl := HostUrl + "/rest/regions"
    values := map[string]interface{}{
        "location": map[string]interface{}{
            "latitude": locLatitude,
            "longitude": locLongitude},
        "name": name,
        "providerUri": providerUri}
    return callHttpRequest("POST", fullUrl, nil, values)
}

func GetRegion(regionID, view string, discover bool) string {
    fullUrl := HostUrl + "/rest/regions/" + regionID
    values := map[string]interface{}{"view": view, "discover": discover}
    return callHttpRequest("GET", fullUrl, nil, values)
}

func UpdateRegion(regionID, info string) string {
    fullUrl := HostUrl + "/rest/regions/" + regionID
    return callHttpRequest("PUT", fullUrl, nil, info)
}

// Zone Types APIs

func GetZoneTypes() string {
    fullUrl := HostUrl + "/rest/zone-types"
    return callHttpRequest("GET", fullUrl, nil, nil)
}

// Zones APIs

func GetZones(regionUri, query string) string {
    fullUrl := HostUrl + "/rest/zones"
    values := map[string]string{"regionUri": regionUri, "query": query}
    return callHttpRequest("GET", fullUrl, nil, values)
}

func CreateZone(name, providerUri, regionUri, zoneTypeUri string) string {
    fullUrl := HostUrl + "/rest/zones"
    values := map[string]string{"name": name, "providerUri": providerUri, 
                                "regionUri": regionUri, "zoneTypeUri": zoneTypeUri}
    return callHttpRequest("POST", fullUrl, nil, values)
}

func GetZone(zoneID, view string) string {
    fullUrl := HostUrl + "/rest/zones/" + zoneID
    values := map[string]string{"view": view}
    return callHttpRequest("GET", fullUrl, nil, values)
}

func UpdateZone(zoneID, infoArray string) string {
    fullUrl := HostUrl + "/rest/zones/" + zoneID
    return callHttpRequest("PUT", fullUrl, nil, infoArray)
}

func DeleteZone(zoneID string, force bool) string {
    fullUrl := HostUrl + "/rest/zones/" + zoneID
    values := map[string]bool{"force": force}
    return callHttpRequest("DELETE", fullUrl, nil, values)
}

func ActionOnZone(zoneID, action string) string {
    fullUrl := HostUrl + "/rest/zones/" + zoneID + "/actions"
    values := map[string]string{"type": action}
    return callHttpRequest("POST", fullUrl, nil, values)
}

func GetZoneApplianceImage(zoneID string) string {
    fullUrl := HostUrl + "/rest/zones/" + zoneID + "/appliance-image"
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

func GetServices(query, userQuery, serviceTypeUri, zoneUri, 
                 workspaceUri, catalogUri, view string) string {
    fullUrl := HostUrl + "/rest/services"
    values := map[string]string{
        "query": query, 
        "userQuery": userQuery, 
        "serviceTypeUri": serviceTypeUri, 
        "zoneUri": zoneUri, 
        "workspaceUri": workspaceUri, 
        "catalogUri": catalogUri, 
        "view": view}
    return callHttpRequest("GET", fullUrl, nil, values)
}

func GetService(serviceID, view string) string {
    fullUrl := HostUrl + "/rest/services/" + serviceID
    values := map[string]string{"view": view}
    return callHttpRequest("GET", fullUrl, nil, values)
}

// Virtual Machine Profiles APIs

func GetVirtualMachineProfiles(query, zoneUri, serviceUri string) string {
    fullUrl := HostUrl + "/rest/virtual-machine-profiles"
    values := map[string]string{"q": query, "zoneUri": zoneUri, "serviceUri": serviceUri}
    return callHttpRequest("GET", fullUrl, nil, values)
}

func GetVirtualMachineProfile(vmProfileID string) string {
    fullUrl := HostUrl + "/rest/virtual-machine-profiles/" + vmProfileID
    return callHttpRequest("GET", fullUrl, nil, nil)
}

// Workspaces APIs

func GetWorkspaces(query, view string) string {
    fullUrl := HostUrl + "/rest/workspaces"
    values := map[string]string{"q": query, "view": view}
    return callHttpRequest("GET", fullUrl, nil, values)
}

func CreateWorkspace(name, description, tagUrisArray string) string {
    fullUrl := HostUrl + "/rest/workspaces"
    values := map[string]string{"name": name, "description": description, "tagUris": tagUrisArray}
    return callHttpRequest("POST", fullUrl, nil, values)
}

func GetWorkspace(workspaceID, view string) string {
    fullUrl := HostUrl + "/rest/workspaces/" + workspaceID
    values := map[string]string{"view": view}
    return callHttpRequest("GET", fullUrl, nil, values)
}

func UpdateWorkspace(workspaceID, name, description, tagUrisArray string) string {
    fullUrl := HostUrl + "/rest/workspaces/" + workspaceID
    values := map[string]string{"name": name, "description": description, "tagUris": tagUrisArray}
    return callHttpRequest("PUT", fullUrl, nil, values)
}

func DeleteWorkspace(workspaceID string) string {
    fullUrl := HostUrl + "/rest/workspaces/" + workspaceID
    return callHttpRequest("DELETE", fullUrl, nil, nil)
}

// Roles APIs

func GetRoles() string {
    fullUrl := HostUrl + "/rest/roles"
    return callHttpRequest("GET", fullUrl, nil, nil)
}

// Users APIs

func GetUsers() string {
    fullUrl := HostUrl + "/rest/users"
    return callHttpRequest("GET", fullUrl, nil, nil)
}

func CreateUser(name, password, email, role string) string {
    fullUrl := HostUrl + "/rest/users"
    values := map[string]string{"name": name, "email": email, "password": password, "role": role}
    return callHttpRequest("POST", fullUrl, nil, values)
}

func GetUser(userID string) string {
    fullUrl := HostUrl + "/rest/users/" + userID
    return callHttpRequest("GET", fullUrl, nil, nil)
}

func UpdateUser(userID, name, password, email, role string) string {
    fullUrl := HostUrl + "/rest/users/" + userID
    values := map[string]string{"name": name, "email": email, "password": password, "role": role}
    return callHttpRequest("PUT", fullUrl, nil, values)
}

func DeleteUser(userID string) string {
    fullUrl := HostUrl + "/rest/users/" + userID
    return callHttpRequest("DELETE", fullUrl, nil, nil)
}

// Volumes APIs

func GetVolumes(query string) string {
    fullUrl := HostUrl + "/rest/volumes"
    values := map[string]string{"query": query}
    return callHttpRequest("GET", fullUrl, nil, values)
}

func GetVolume(volumeID string) string {
    fullUrl := HostUrl + "/rest/volumes/" + volumeID
    return callHttpRequest("GET", fullUrl, nil, nil)
}

// Tag Keys APIs

func GetTagKeys(view string) string {
    fullUrl := HostUrl + "/rest/tag-keys"
    values := map[string]string{"view": view}
    return callHttpRequest("GET", fullUrl, nil, values)
}

func CreateTagKey(name string) string {
    fullUrl := HostUrl + "/rest/tag-keys"
    values := map[string]string{"name": name}
    return callHttpRequest("POST", fullUrl, nil, values)
}

func GetTagKey(tagKeyID, view string) string {
    fullUrl := HostUrl + "/rest/tag-keys/" + tagKeyID
    values := map[string]string{"view": view}
    return callHttpRequest("GET", fullUrl, nil, values)
}

func DeleteTagKey(tagKeyID string) string {
    fullUrl := HostUrl + "/rest/tag-keys/" + tagKeyID
    return callHttpRequest("DELETE", fullUrl, nil, nil)
}

// Tags APIs

func GetTags(view string) string {
    fullUrl := HostUrl + "/rest/tags"
    values := map[string]string{"view": view}
    return callHttpRequest("GET", fullUrl, nil, values)
}

func CreateTag(name, tagKeyUri string) string {
    fullUrl := HostUrl + "/rest/tags"
    values := map[string]string{"name": name, "tagKeyUri": tagKeyUri}
    return callHttpRequest("POST", fullUrl, nil, values)
}

func GetTag(tagID, view string) string {
    fullUrl := HostUrl + "/rest/tags/" + tagID
    values := map[string]string{"view": view}
    return callHttpRequest("GET", fullUrl, nil, values)
}

func DeleteTag(tagID string) string {
    fullUrl := HostUrl + "/rest/tags/" + tagID
    return callHttpRequest("DELETE", fullUrl, nil, nil)
}

