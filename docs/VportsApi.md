# \VportsApi

All URIs are relative to *https://penapi.pacnetconnect.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVNFVPort**](VportsApi.md#CreateVNFVPort) | **Post** /1.0.0/inventory/vnf/vport | Create VNF VPort
[**CreateVPortForPhysicalEndpoint**](VportsApi.md#CreateVPortForPhysicalEndpoint) | **Post** /1.0.0/inventory/regularvport | Create VPort for physical endpoint
[**GetInformationAboutTheSpecifiedVPort**](VportsApi.md#GetInformationAboutTheSpecifiedVPort) | **Get** /1.0.0/inventory/vport/{vportuuid} | Get information about the specified VPort


# **CreateVNFVPort**
> Model100InventoryVnfVportResponse CreateVNFVPort($body)

Create VNF VPort

Create VNF VPort


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Model100InventoryVnfVportRequest**](Model100InventoryVnfVportRequest.md)|  | [optional] 

### Return type

[**Model100InventoryVnfVportResponse**](100InventoryVnfVportResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVPortForPhysicalEndpoint**
> Model100InventoryRegularvportResponse CreateVPortForPhysicalEndpoint($body)

Create VPort for physical endpoint

Create VPort representing a VLAN on a Physical Ethernet Port


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Model100InventoryRegularvportRequest**](Model100InventoryRegularvportRequest.md)|  | [optional] 

### Return type

[**Model100InventoryRegularvportResponse**](100InventoryRegularvportResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInformationAboutTheSpecifiedVPort**
> EndpointPort GetInformationAboutTheSpecifiedVPort($vportuuid)

Get information about the specified VPort

Get information about the specified VPort


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vportuuid** | **string**| Unique identifier representing a specific virtual port | 

### Return type

[**EndpointPort**](EndpointPort.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

