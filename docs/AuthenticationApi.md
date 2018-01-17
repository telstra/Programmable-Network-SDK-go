# \AuthenticationApi

All URIs are relative to *https://penapi.pacnetconnect.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthGeneratetokenPost**](AuthenticationApi.md#AuthGeneratetokenPost) | **Post** /1.0.0/auth/generatetoken | Create an authentication token
[**AuthValidatetokenGet**](AuthenticationApi.md#AuthValidatetokenGet) | **Get** /1.0.0/auth/validatetoken | Validate authentication token


# **AuthGeneratetokenPost**
> AuthGeneratetokenResponse AuthGeneratetokenPost(ctx, grantType, username, password)
Create an authentication token

Create an authentication token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **grantType** | **string**|  | 
  **username** | **string**|  | 
  **password** | **string**|  | 

### Return type

[**AuthGeneratetokenResponse**](AuthGeneratetokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthValidatetokenGet**
> AuthValidatetokenResponse AuthValidatetokenGet(ctx, )
Validate authentication token

Validate the authentication token and get information about the user (roles, permissions, etc.)

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AuthValidatetokenResponse**](AuthValidatetokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

