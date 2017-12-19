# \EndpointsApi

All URIs are relative to *https://penapi.pacnetconnect.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**100InventoryEndpointByEndpointuuidGet**](EndpointsApi.md#100InventoryEndpointByEndpointuuidGet) | **Get** /1.0.0/inventory/endpoint/{endpointuuid} | Get information about the specified endpoint
[**100InventoryEndpointsCustomeruuidByCustomeruuidGet**](EndpointsApi.md#100InventoryEndpointsCustomeruuidByCustomeruuidGet) | **Get** /1.0.0/inventory/endpoints/customeruuid/{customeruuid} | Get list of endpoints for a customer
[**100InventoryRegularendpointPost**](EndpointsApi.md#100InventoryRegularendpointPost) | **Post** /1.0.0/inventory/regularendpoint | Create Physical (Port) Endpoint
[**100InventoryVnfendpointPost**](EndpointsApi.md#100InventoryVnfendpointPost) | **Post** /1.0.0/inventory/vnfendpoint | Create VNF Endpoint
[**Eis100EndpointsAssignTopologyTagByEndpointuuidPost**](EndpointsApi.md#Eis100EndpointsAssignTopologyTagByEndpointuuidPost) | **Post** /eis/1.0.0/endpoints/{endpointuuid}/assign_topology_tag | Assign a Topology Tag to an Endpoint


# **100InventoryEndpointByEndpointuuidGet**
> Model100InventoryEndpointResponse 100InventoryEndpointByEndpointuuidGet($endpointuuid)

Get information about the specified endpoint

Get information about the specified endpoint


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointuuid** | **string**| Unique identifier representing a specific endpoint | 

### Return type

[**Model100InventoryEndpointResponse**](100InventoryEndpointResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **100InventoryEndpointsCustomeruuidByCustomeruuidGet**
> Model100InventoryEndpointsCustomeruuidResponse 100InventoryEndpointsCustomeruuidByCustomeruuidGet($customeruuid)

Get list of endpoints for a customer

Get list of endpoints for a customer


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customeruuid** | **string**| Unique identifier representing a specific customer | 

### Return type

[**Model100InventoryEndpointsCustomeruuidResponse**](100InventoryEndpointsCustomeruuidResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **100InventoryRegularendpointPost**
> Model100InventoryRegularendpointResponse 100InventoryRegularendpointPost($body)

Create Physical (Port) Endpoint

Create Physical (Port) Endpoint


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Model100InventoryRegularendpointRequest**](Model100InventoryRegularendpointRequest.md)|  | [optional] 

### Return type

[**Model100InventoryRegularendpointResponse**](100InventoryRegularendpointResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **100InventoryVnfendpointPost**
> Model100InventoryVnfendpointResponse 100InventoryVnfendpointPost($body)

Create VNF Endpoint

Create VNF Endpoint


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Model100InventoryVnfendpointRequest**](Model100InventoryVnfendpointRequest.md)|  | [optional] 

### Return type

[**Model100InventoryVnfendpointResponse**](100InventoryVnfendpointResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Eis100EndpointsAssignTopologyTagByEndpointuuidPost**
> SuccessFragment Eis100EndpointsAssignTopologyTagByEndpointuuidPost($endpointuuid, $body)

Assign a Topology Tag to an Endpoint

Assign a Topology Tag to an Endpoint


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointuuid** | **string**| Unique identifier representing a specific endpoint | 
 **body** | [**Eis100EndpointsAssignTopologyTagRequest**](Eis100EndpointsAssignTopologyTagRequest.md)|  | [optional] 

### Return type

[**SuccessFragment**](SuccessFragment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

