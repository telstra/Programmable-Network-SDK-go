# \TopologiesApi

All URIs are relative to *https://penapi.pacnetconnect.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Ttms100TopologyTagByTopotaguuidDelete**](TopologiesApi.md#Ttms100TopologyTagByTopotaguuidDelete) | **Delete** /ttms/1.0.0/topology_tag/{topotaguuid}/ | Delete a topology tag
[**Ttms100TopologyTagByTopotaguuidGet**](TopologiesApi.md#Ttms100TopologyTagByTopotaguuidGet) | **Get** /ttms/1.0.0/topology_tag/{topotaguuid}/ | Get information about the specified topology tag
[**Ttms100TopologyTagByTopotaguuidPut**](TopologiesApi.md#Ttms100TopologyTagByTopotaguuidPut) | **Put** /ttms/1.0.0/topology_tag/{topotaguuid}/ | Update a topology tag&#39;s name and/or description
[**Ttms100TopologyTagGet**](TopologiesApi.md#Ttms100TopologyTagGet) | **Get** /ttms/1.0.0/topology_tag | List all topology tags
[**Ttms100TopologyTagObjectsByTopotaguuidGet**](TopologiesApi.md#Ttms100TopologyTagObjectsByTopotaguuidGet) | **Get** /ttms/1.0.0/topology_tag/{topotaguuid}/objects/ | List objects for Topology
[**Ttms100TopologyTagPost**](TopologiesApi.md#Ttms100TopologyTagPost) | **Post** /ttms/1.0.0/topology_tag | Create a named topology tag


# **Ttms100TopologyTagByTopotaguuidDelete**
> Ttms100TopologyTagByTopotaguuidDelete($topotaguuid)

Delete a topology tag

Delete a topology tag


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topotaguuid** | **string**| Unique identifier representing a specific topology tag | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Ttms100TopologyTagByTopotaguuidGet**
> Topology Ttms100TopologyTagByTopotaguuidGet($topotaguuid)

Get information about the specified topology tag

Get information about the specified topology tag


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topotaguuid** | **string**| Unique identifier representing a specific topology tag | 

### Return type

[**Topology**](Topology.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Ttms100TopologyTagByTopotaguuidPut**
> Topology Ttms100TopologyTagByTopotaguuidPut($topotaguuid, $body)

Update a topology tag's name and/or description

Update a topology tag's name and/or description


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topotaguuid** | **string**| Unique identifier representing a specific topology tag | 
 **body** | [**Ttms100TopologyTagRequest**](Ttms100TopologyTagRequest.md)|  | [optional] 

### Return type

[**Topology**](Topology.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Ttms100TopologyTagGet**
> []Topology Ttms100TopologyTagGet()

List all topology tags

List all topology tags


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]Topology**](Topology.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Ttms100TopologyTagObjectsByTopotaguuidGet**
> Ttms100TopologyTagObjectsResponse Ttms100TopologyTagObjectsByTopotaguuidGet($topotaguuid)

List objects for Topology

List all objects (Endpoints, Links, VPorts, etc.) associated with the topology tag.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topotaguuid** | **string**| Unique identifier representing a specific topology tag | 

### Return type

[**Ttms100TopologyTagObjectsResponse**](Ttms100TopologyTagObjectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Ttms100TopologyTagPost**
> Topology Ttms100TopologyTagPost($body)

Create a named topology tag

Create a named topology tag


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Ttms100TopologyTagRequest**](Ttms100TopologyTagRequest.md)|  | [optional] 

### Return type

[**Topology**](Topology.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

