# \LinksApi

All URIs are relative to *https://penapi.pacnetconnect.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**100InventoryLinkPost**](LinksApi.md#100InventoryLinkPost) | **Post** /1.0.0/inventory/link | Create Link and initial Contract
[**100InventoryLinksByLinkidGet**](LinksApi.md#100InventoryLinksByLinkidGet) | **Get** /1.0.0/inventory/links/{linkid} | Get details of specified link
[**100InventoryLinksCustomerByCustomeruuidGet**](LinksApi.md#100InventoryLinksCustomerByCustomeruuidGet) | **Get** /1.0.0/inventory/links/customer/{customeruuid} | Get active Links
[**100InventoryLinksHistoryByLinkidGet**](LinksApi.md#100InventoryLinksHistoryByLinkidGet) | **Get** /1.0.0/inventory/links/history/{linkid} | Get Link history


# **100InventoryLinkPost**
> Model100InventoryLinkResponse 100InventoryLinkPost($body)

Create Link and initial Contract

Create Link and initial Contract


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Model100InventoryLinkRequest**](Model100InventoryLinkRequest.md)|  | [optional] 

### Return type

[**Model100InventoryLinkResponse**](100InventoryLinkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **100InventoryLinksByLinkidGet**
> Model100InventoryLinksResponse 100InventoryLinksByLinkidGet($linkid)

Get details of specified link

Get details of specified link


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkid** | **string**| Unique identifier representing a specific link | 

### Return type

[**Model100InventoryLinksResponse**](100InventoryLinksResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **100InventoryLinksCustomerByCustomeruuidGet**
> []Link 100InventoryLinksCustomerByCustomeruuidGet($customeruuid)

Get active Links

Get active Links


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customeruuid** | **string**| Unique identifier representing a specific customer | 

### Return type

[**[]Link**](Link.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **100InventoryLinksHistoryByLinkidGet**
> Model100InventoryLinksHistoryResponse 100InventoryLinksHistoryByLinkidGet($linkid)

Get Link history

Get Link history


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkid** | **string**| Unique identifier representing a specific link | 

### Return type

[**Model100InventoryLinksHistoryResponse**](100InventoryLinksHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

