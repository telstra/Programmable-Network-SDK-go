# \ContractsApi

All URIs are relative to *https://penapi.pacnetconnect.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**100InventoryLinksContractByLinkidAndContractidGet**](ContractsApi.md#100InventoryLinksContractByLinkidAndContractidGet) | **Get** /1.0.0/inventory/links/{linkid}/contract/{contractid} | Get active Contract by ContractID
[**100InventoryLinksContractByLinkidAndContractidPut**](ContractsApi.md#100InventoryLinksContractByLinkidAndContractidPut) | **Put** /1.0.0/inventory/links/{linkid}/contract/{contractid} | Update active Contract by ContractID
[**100InventoryLinksContractByLinkidPost**](ContractsApi.md#100InventoryLinksContractByLinkidPost) | **Post** /1.0.0/inventory/links/{linkid}/contract | Create new Contract on specified link


# **100InventoryLinksContractByLinkidAndContractidGet**
> Model100InventoryLinksContractResponse 100InventoryLinksContractByLinkidAndContractidGet($linkid, $contractid)

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

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **100InventoryLinksContractByLinkidAndContractidPut**
> Model100InventoryLinksContractResponse33 100InventoryLinksContractByLinkidAndContractidPut($linkid, $contractid, $body)

Update active Contract by ContractID

Update active Contract by ContractID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkid** | **string**| Unique identifier representing a specific link | 
 **contractid** | **string**| Unique identifier representing a specific contract | 
 **body** | [**Model100InventoryLinksContractRequest**](Model100InventoryLinksContractRequest.md)|  | [optional] 

### Return type

[**Model100InventoryLinksContractResponse33**](100InventoryLinksContractResponse33.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **100InventoryLinksContractByLinkidPost**
> Model100InventoryLinksContractResponse38 100InventoryLinksContractByLinkidPost($linkid, $body)

Create new Contract on specified link

Create new Contract on specified link


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkid** | **string**| Unique identifier representing a specific link | 
 **body** | [**Model100InventoryLinksContractRequest37**](Model100InventoryLinksContractRequest37.md)|  | [optional] 

### Return type

[**Model100InventoryLinksContractResponse38**](100InventoryLinksContractResponse38.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

