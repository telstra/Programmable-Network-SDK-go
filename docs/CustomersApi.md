# \CustomersApi

All URIs are relative to *https://penapi.pacnetconnect.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountByCustomeruuidGet**](CustomersApi.md#AccountByCustomeruuidGet) | **Get** /1.0.0/account/{customeruuid} | Get account information details
[**AccountUserByCustomeruuidGet**](CustomersApi.md#AccountUserByCustomeruuidGet) | **Get** /1.0.0/account/{customeruuid}/user | List users


# **AccountByCustomeruuidGet**
> []AccountResponse AccountByCustomeruuidGet(ctx, customeruuid)
Get account information details

Get the account information for the specified customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **customeruuid** | **string**| Unique identifier representing a specific customer | 

### Return type

[**[]AccountResponse**](AccountResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountUserByCustomeruuidGet**
> []User AccountUserByCustomeruuidGet(ctx, customeruuid)
List users

List all users associated with the specified customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **customeruuid** | **string**| Unique identifier representing a specific customer | 

### Return type

[**[]User**](User.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

