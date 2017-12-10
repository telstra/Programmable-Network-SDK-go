# \ContractsApi

All URIs are relative to *https://penapi.pacnetconnect.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewContractOnSpecifiedLink**](ContractsApi.md#CreateNewContractOnSpecifiedLink) | **Post** /1.0.0/inventory/links/{linkid}/contract | Create new Contract on specified link
[**GetActiveContractByContractID**](ContractsApi.md#GetActiveContractByContractID) | **Get** /1.0.0/inventory/links/{linkid}/contract/{contractid} | Get active Contract by ContractID
[**UpdateActiveContractByContractID**](ContractsApi.md#UpdateActiveContractByContractID) | **Put** /1.0.0/inventory/links/{linkid}/contract/{contractid} | Update active Contract by ContractID


# **CreateNewContractOnSpecifiedLink**
> Model100InventoryLinksContractResponse36 CreateNewContractOnSpecifiedLink($linkid, $body)

Create new Contract on specified link

Create new Contract on specified link


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkid** | **string**| Unique identifier representing a specific link | 
 **body** | [**Model100InventoryLinksContractRequest35**](Model100InventoryLinksContractRequest35.md)|  | [optional] 

### Return type

[**Model100InventoryLinksContractResponse36**](100InventoryLinksContractResponse36.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActiveContractByContractID**
> Model100InventoryLinksContractResponse GetActiveContractByContractID($linkid, $contractid)

Get active Contract by ContractID

Get active Contract by ContractID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkid** | **string**| Unique identifier representing a specific link | 
 **contractid** | **string**| Unique identifier representing a specific contract | 

### Return type

[**Model100InventoryLinksContractResponse**](100InventoryLinksContractResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateActiveContractByContractID**
> Model100InventoryLinksContractResponse31 UpdateActiveContractByContractID($linkid, $contractid, $body)

Update active Contract by ContractID

Update active Contract by ContractID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkid** | **string**| Unique identifier representing a specific link | 
 **contractid** | **string**| Unique identifier representing a specific contract | 
 **body** | [**Model100InventoryLinksContractRequest**](Model100InventoryLinksContractRequest.md)|  | [optional] 

### Return type

[**Model100InventoryLinksContractResponse31**](100InventoryLinksContractResponse31.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

