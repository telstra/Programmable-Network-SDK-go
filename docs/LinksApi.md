# \LinksApi

All URIs are relative to *https://penapi.pacnetconnect.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLinkAndInitialContract**](LinksApi.md#CreateLinkAndInitialContract) | **Post** /1.0.0/inventory/link | Create Link and initial Contract
[**GetActiveLinks**](LinksApi.md#GetActiveLinks) | **Get** /1.0.0/inventory/links/customer/{customeruuid} | Get active Links
[**GetDetailsOfSpecifiedLink**](LinksApi.md#GetDetailsOfSpecifiedLink) | **Get** /1.0.0/inventory/links/{linkid} | Get details of specified link
[**GetLinkHistory**](LinksApi.md#GetLinkHistory) | **Get** /1.0.0/inventory/links/history/{linkid} | Get Link history


# **CreateLinkAndInitialContract**
> Model100InventoryLinkResponse CreateLinkAndInitialContract($body)

Create Link and initial Contract

Create Link and initial Contract


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Model100InventoryLinkRequest**](Model100InventoryLinkRequest.md)|  | [optional] 

### Return type

[**Model100InventoryLinkResponse**](100InventoryLinkResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActiveLinks**
> []Link GetActiveLinks($customeruuid)

Get active Links

Get active Links


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customeruuid** | **string**| Unique identifier representing a specific customer | 

### Return type

[**[]Link**](Link.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetailsOfSpecifiedLink**
> Model100InventoryLinksResponse GetDetailsOfSpecifiedLink($linkid)

Get details of specified link

Get details of specified link


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkid** | **string**| Unique identifier representing a specific link | 

### Return type

[**Model100InventoryLinksResponse**](100InventoryLinksResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLinkHistory**
> Model100InventoryLinksHistoryResponse GetLinkHistory($linkid)

Get Link history

Get Link history


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkid** | **string**| Unique identifier representing a specific link | 

### Return type

[**Model100InventoryLinksHistoryResponse**](100InventoryLinksHistoryResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

