# \CustomersApi

All URIs are relative to *https://penapi.pacnetconnect.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountInformationDetails**](CustomersApi.md#GetAccountInformationDetails) | **Get** /1.0.0/account/{customeruuid} | Get account information details
[**ListUsers**](CustomersApi.md#ListUsers) | **Get** /1.0.0/account/{customeruuid}/user | List users


# **GetAccountInformationDetails**
> Model100AccountResponse GetAccountInformationDetails($customeruuid)

Get account information details

Get the account information for the specified customer


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customeruuid** | **string**| Unique identifier representing a specific customer | 

### Return type

[**Model100AccountResponse**](100AccountResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsers**
> []User ListUsers($customeruuid)

List users

List all users associated with the specified customer


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customeruuid** | **string**| Unique identifier representing a specific customer | 

### Return type

[**[]User**](User.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

