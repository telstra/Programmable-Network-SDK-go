# \EndpointsApi

All URIs are relative to *https://penapi.pacnetconnect.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePhysicalPortEndpoint**](EndpointsApi.md#CreatePhysicalPortEndpoint) | **Post** /1.0.0/inventory/regularendpoint | Create Physical (Port) Endpoint
[**CreateVNFEndpoint**](EndpointsApi.md#CreateVNFEndpoint) | **Post** /1.0.0/inventory/vnfendpoint | Create VNF Endpoint
[**GetInformationAboutTheSpecifiedEndpoint**](EndpointsApi.md#GetInformationAboutTheSpecifiedEndpoint) | **Get** /1.0.0/inventory/endpoint/{endpointuuid} | Get information about the specified endpoint
[**GetListOfEndpointsForACustomer**](EndpointsApi.md#GetListOfEndpointsForACustomer) | **Get** /1.0.0/inventory/endpoints/customeruuid/{customeruuid} | Get list of endpoints for a customer


# **CreatePhysicalPortEndpoint**
> Model100InventoryRegularendpointResponse CreatePhysicalPortEndpoint($body)

Create Physical (Port) Endpoint

Create Physical (Port) Endpoint


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Model100InventoryRegularendpointRequest**](Model100InventoryRegularendpointRequest.md)|  | [optional] 

### Return type

[**Model100InventoryRegularendpointResponse**](100InventoryRegularendpointResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVNFEndpoint**
> Model100InventoryVnfendpointResponse CreateVNFEndpoint($body)

Create VNF Endpoint

Create VNF Endpoint


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Model100InventoryVnfendpointRequest**](Model100InventoryVnfendpointRequest.md)|  | [optional] 

### Return type

[**Model100InventoryVnfendpointResponse**](100InventoryVnfendpointResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInformationAboutTheSpecifiedEndpoint**
> Model100InventoryEndpointResponse GetInformationAboutTheSpecifiedEndpoint($endpointuuid)

Get information about the specified endpoint

Get information about the specified endpoint


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointuuid** | **string**| Unique identifier representing a specific endpoint | 

### Return type

[**Model100InventoryEndpointResponse**](100InventoryEndpointResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListOfEndpointsForACustomer**
> Model100InventoryEndpointsCustomeruuidResponse GetListOfEndpointsForACustomer($customeruuid)

Get list of endpoints for a customer

Get list of endpoints for a customer


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customeruuid** | **string**| Unique identifier representing a specific customer | 

### Return type

[**Model100InventoryEndpointsCustomeruuidResponse**](100InventoryEndpointsCustomeruuidResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

