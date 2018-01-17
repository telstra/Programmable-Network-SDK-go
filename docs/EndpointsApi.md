# \EndpointsApi

All URIs are relative to *https://penapi.pacnetconnect.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Eis100EndpointsAssignTopologyTagByEndpointuuidPost**](EndpointsApi.md#Eis100EndpointsAssignTopologyTagByEndpointuuidPost) | **Post** /eis/1.0.0/endpoints/{endpointuuid}/assign_topology_tag | Assign a Topology Tag to an Endpoint
[**InventoryEndpointByEndpointuuidGet**](EndpointsApi.md#InventoryEndpointByEndpointuuidGet) | **Get** /1.0.0/inventory/endpoint/{endpointuuid} | Get information about the specified endpoint
[**InventoryEndpointsCustomeruuidByCustomeruuidGet**](EndpointsApi.md#InventoryEndpointsCustomeruuidByCustomeruuidGet) | **Get** /1.0.0/inventory/endpoints/customeruuid/{customeruuid} | Get list of endpoints for a customer
[**InventoryRegularendpointPost**](EndpointsApi.md#InventoryRegularendpointPost) | **Post** /1.0.0/inventory/regularendpoint | Create Physical (Port) Endpoint
[**InventoryVnfendpointPost**](EndpointsApi.md#InventoryVnfendpointPost) | **Post** /1.0.0/inventory/vnfendpoint | Create VNF Endpoint


# **Eis100EndpointsAssignTopologyTagByEndpointuuidPost**
> []SuccessFragment Eis100EndpointsAssignTopologyTagByEndpointuuidPost(ctx, endpointuuid, optional)
Assign a Topology Tag to an Endpoint

Assign a Topology Tag to an Endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **endpointuuid** | **string**| Unique identifier representing a specific endpoint | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointuuid** | **string**| Unique identifier representing a specific endpoint | 
 **body** | [**Eis100EndpointsAssignTopologyTagRequest**](Eis100EndpointsAssignTopologyTagRequest.md)|  | 

### Return type

[**[]SuccessFragment**](SuccessFragment.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryEndpointByEndpointuuidGet**
> InventoryEndpointResponse InventoryEndpointByEndpointuuidGet(ctx, endpointuuid)
Get information about the specified endpoint

Get information about the specified endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **endpointuuid** | **string**| Unique identifier representing a specific endpoint | 

### Return type

[**InventoryEndpointResponse**](InventoryEndpointResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryEndpointsCustomeruuidByCustomeruuidGet**
> InventoryEndpointsCustomeruuidResponse InventoryEndpointsCustomeruuidByCustomeruuidGet(ctx, customeruuid)
Get list of endpoints for a customer

Get list of endpoints for a customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **customeruuid** | **string**| Unique identifier representing a specific customer | 

### Return type

[**InventoryEndpointsCustomeruuidResponse**](InventoryEndpointsCustomeruuidResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryRegularendpointPost**
> []InventoryRegularendpointResponse InventoryRegularendpointPost(ctx, optional)
Create Physical (Port) Endpoint

Create Physical (Port) Endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryRegularendpointRequest**](InventoryRegularendpointRequest.md)|  | 

### Return type

[**[]InventoryRegularendpointResponse**](InventoryRegularendpointResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryVnfendpointPost**
> []InventoryVnfendpointResponse InventoryVnfendpointPost(ctx, optional)
Create VNF Endpoint

Create VNF Endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryVnfendpointRequest**](InventoryVnfendpointRequest.md)|  | 

### Return type

[**[]InventoryVnfendpointResponse**](InventoryVnfendpointResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

