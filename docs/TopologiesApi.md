# \TopologiesApi

All URIs are relative to *https://penapi.pacnetconnect.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateANamedTopologyTag**](TopologiesApi.md#CreateANamedTopologyTag) | **Post** /ttms/1.0.0/topology_tag | Create a named topology tag
[**GetInformationAboutTheSpecifiedTopologyTag**](TopologiesApi.md#GetInformationAboutTheSpecifiedTopologyTag) | **Get** /ttms/1.0.0/topology_tag/{topotaguuid}/ | Get information about the specified topology tag
[**ListAllTopologyTags**](TopologiesApi.md#ListAllTopologyTags) | **Get** /ttms/1.0.0/topology_tag | List all topology tags
[**ListObjectsForTopology**](TopologiesApi.md#ListObjectsForTopology) | **Get** /ttms/1.0.0/topology_tag/{topotaguuid}/objects/ | List objects for Topology


# **CreateANamedTopologyTag**
> Topology CreateANamedTopologyTag($body)

Create a named topology tag

Create a named topology tag


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Ttms100TopologyTagRequest**](Ttms100TopologyTagRequest.md)|  | [optional] 

### Return type

[**Topology**](Topology.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInformationAboutTheSpecifiedTopologyTag**
> Topology GetInformationAboutTheSpecifiedTopologyTag($topotaguuid)

Get information about the specified topology tag

Get information about the specified topology tag


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topotaguuid** | **string**| Unique identifier representing a specific topology tag | 

### Return type

[**Topology**](Topology.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllTopologyTags**
> []Topology ListAllTopologyTags()

List all topology tags

List all topology tags


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]Topology**](Topology.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListObjectsForTopology**
> Ttms100TopologyTagObjectsResponse ListObjectsForTopology($topotaguuid)

List objects for Topology

List all objects (Endpoints, Links, VPorts, etc.) associated with the topology tag.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topotaguuid** | **string**| Unique identifier representing a specific topology tag | 

### Return type

[**Ttms100TopologyTagObjectsResponse**](Ttms100TopologyTagObjectsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

