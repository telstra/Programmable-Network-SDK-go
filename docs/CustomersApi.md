# \CustomersApi

All URIs are relative to *https://penapi.pacnetconnect.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**100AccountByCustomeruuidGet**](CustomersApi.md#100AccountByCustomeruuidGet) | **Get** /1.0.0/account/{customeruuid} | Get account information details
[**100AccountUserByCustomeruuidGet**](CustomersApi.md#100AccountUserByCustomeruuidGet) | **Get** /1.0.0/account/{customeruuid}/user | List users


# **100AccountByCustomeruuidGet**
> Model100AccountResponse 100AccountByCustomeruuidGet($customeruuid)

Get account information details

Get the account information for the specified customer


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customeruuid** | **string**| Unique identifier representing a specific customer | 

### Return type

[**Model100AccountResponse**](100AccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **100AccountUserByCustomeruuidGet**
> []User 100AccountUserByCustomeruuidGet($customeruuid)

List users

List all users associated with the specified customer


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customeruuid** | **string**| Unique identifier representing a specific customer | 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

