# \LinksApi

All URIs are relative to *https://penapi.pacnetconnect.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryLinkPost**](LinksApi.md#InventoryLinkPost) | **Post** /1.0.0/inventory/link | Create Link and initial Contract
[**InventoryLinksByLinkidGet**](LinksApi.md#InventoryLinksByLinkidGet) | **Get** /1.0.0/inventory/links/{linkid} | Get details of specified link
[**InventoryLinksCustomerByCustomeruuidGet**](LinksApi.md#InventoryLinksCustomerByCustomeruuidGet) | **Get** /1.0.0/inventory/links/customer/{customeruuid} | Get active Links
[**InventoryLinksHistoryByLinkidGet**](LinksApi.md#InventoryLinksHistoryByLinkidGet) | **Get** /1.0.0/inventory/links/history/{linkid} | Get Link history


# **InventoryLinkPost**
> InventoryLinkResponse InventoryLinkPost(ctx, optional)
Create Link and initial Contract

Create Link and initial Contract

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryLinkRequest**](InventoryLinkRequest.md)|  | 

### Return type

[**InventoryLinkResponse**](InventoryLinkResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryLinksByLinkidGet**
> InventoryLinksResponse InventoryLinksByLinkidGet(ctx, linkid)
Get details of specified link

Get details of specified link

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **linkid** | **string**| Unique identifier representing a specific link | 

### Return type

[**InventoryLinksResponse**](InventoryLinksResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryLinksCustomerByCustomeruuidGet**
> []Link InventoryLinksCustomerByCustomeruuidGet(ctx, customeruuid)
Get active Links

Get active Links

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **customeruuid** | **string**| Unique identifier representing a specific customer | 

### Return type

[**[]Link**](Link.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryLinksHistoryByLinkidGet**
> InventoryLinksHistoryResponse InventoryLinksHistoryByLinkidGet(ctx, linkid)
Get Link history

Get Link history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **linkid** | **string**| Unique identifier representing a specific link | 

### Return type

[**InventoryLinksHistoryResponse**](InventoryLinksHistoryResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

