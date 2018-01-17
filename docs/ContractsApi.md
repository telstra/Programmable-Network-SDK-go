# \ContractsApi

All URIs are relative to *https://penapi.pacnetconnect.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryLinksContractByLinkidAndContractidGet**](ContractsApi.md#InventoryLinksContractByLinkidAndContractidGet) | **Get** /1.0.0/inventory/links/{linkid}/contract/{contractid} | Get active Contract by ContractID
[**InventoryLinksContractByLinkidAndContractidPut**](ContractsApi.md#InventoryLinksContractByLinkidAndContractidPut) | **Put** /1.0.0/inventory/links/{linkid}/contract/{contractid} | Update active Contract by ContractID
[**InventoryLinksContractByLinkidPost**](ContractsApi.md#InventoryLinksContractByLinkidPost) | **Post** /1.0.0/inventory/links/{linkid}/contract | Create new Contract on specified link


# **InventoryLinksContractByLinkidAndContractidGet**
> InventoryLinksContractResponse InventoryLinksContractByLinkidAndContractidGet(ctx, linkid, contractid)
Get active Contract by ContractID

Get active Contract by ContractID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **linkid** | **string**| Unique identifier representing a specific link | 
  **contractid** | **string**| Unique identifier representing a specific contract | 

### Return type

[**InventoryLinksContractResponse**](InventoryLinksContractResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryLinksContractByLinkidAndContractidPut**
> InventoryLinksContractResponse33 InventoryLinksContractByLinkidAndContractidPut(ctx, linkid, contractid, optional)
Update active Contract by ContractID

Update active Contract by ContractID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **linkid** | **string**| Unique identifier representing a specific link | 
  **contractid** | **string**| Unique identifier representing a specific contract | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkid** | **string**| Unique identifier representing a specific link | 
 **contractid** | **string**| Unique identifier representing a specific contract | 
 **body** | [**InventoryLinksContractRequest**](InventoryLinksContractRequest.md)|  | 

### Return type

[**InventoryLinksContractResponse33**](InventoryLinksContractResponse33.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryLinksContractByLinkidPost**
> []InventoryLinksContractResponse38 InventoryLinksContractByLinkidPost(ctx, linkid, optional)
Create new Contract on specified link

Create new Contract on specified link

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **linkid** | **string**| Unique identifier representing a specific link | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkid** | **string**| Unique identifier representing a specific link | 
 **body** | [**InventoryLinksContractRequest37**](InventoryLinksContractRequest37.md)|  | 

### Return type

[**[]InventoryLinksContractResponse38**](InventoryLinksContractResponse38.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

