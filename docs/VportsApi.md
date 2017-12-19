# \VportsApi

All URIs are relative to *https://penapi.pacnetconnect.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**100InventoryRegularvportPost**](VportsApi.md#100InventoryRegularvportPost) | **Post** /1.0.0/inventory/regularvport | Create VPort for physical endpoint
[**100InventoryVnfVportPost**](VportsApi.md#100InventoryVnfVportPost) | **Post** /1.0.0/inventory/vnf/vport | Create VNF VPort
[**100InventoryVportByVportuuidGet**](VportsApi.md#100InventoryVportByVportuuidGet) | **Get** /1.0.0/inventory/vport/{vportuuid} | Get information about the specified VPort


# **100InventoryRegularvportPost**
> Model100InventoryRegularvportResponse 100InventoryRegularvportPost($body)

Create VPort for physical endpoint

Create VPort representing a VLAN on a Physical Ethernet Port


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Model100InventoryRegularvportRequest**](Model100InventoryRegularvportRequest.md)|  | [optional] 

### Return type

[**Model100InventoryRegularvportResponse**](100InventoryRegularvportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **100InventoryVnfVportPost**
> Model100InventoryVnfVportResponse 100InventoryVnfVportPost($body)

Create VNF VPort

Create VNF VPort


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Model100InventoryVnfVportRequest**](Model100InventoryVnfVportRequest.md)|  | [optional] 

### Return type

[**Model100InventoryVnfVportResponse**](100InventoryVnfVportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **100InventoryVportByVportuuidGet**
> EndpointPort 100InventoryVportByVportuuidGet($vportuuid)

Get information about the specified VPort

Get information about the specified VPort


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vportuuid** | **string**| Unique identifier representing a specific virtual port | 

### Return type

[**EndpointPort**](EndpointPort.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

