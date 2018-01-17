# \VportsApi

All URIs are relative to *https://penapi.pacnetconnect.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryRegularvportPost**](VportsApi.md#InventoryRegularvportPost) | **Post** /1.0.0/inventory/regularvport | Create VPort for physical endpoint
[**InventoryVnfVportPost**](VportsApi.md#InventoryVnfVportPost) | **Post** /1.0.0/inventory/vnf/vport | Create VNF VPort
[**InventoryVportByVportuuidGet**](VportsApi.md#InventoryVportByVportuuidGet) | **Get** /1.0.0/inventory/vport/{vportuuid} | Get information about the specified VPort


# **InventoryRegularvportPost**
> InventoryRegularvportResponse InventoryRegularvportPost(ctx, optional)
Create VPort for physical endpoint

Create VPort representing a VLAN on a Physical Ethernet Port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryRegularvportRequest**](InventoryRegularvportRequest.md)|  | 

### Return type

[**InventoryRegularvportResponse**](InventoryRegularvportResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryVnfVportPost**
> InventoryVnfVportResponse InventoryVnfVportPost(ctx, optional)
Create VNF VPort

Create VNF VPort

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryVnfVportRequest**](InventoryVnfVportRequest.md)|  | 

### Return type

[**InventoryVnfVportResponse**](InventoryVnfVportResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryVportByVportuuidGet**
> []EndpointPort InventoryVportByVportuuidGet(ctx, vportuuid)
Get information about the specified VPort

Get information about the specified VPort

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **vportuuid** | **string**| Unique identifier representing a specific virtual port | 

### Return type

[**[]EndpointPort**](EndpointPort.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

