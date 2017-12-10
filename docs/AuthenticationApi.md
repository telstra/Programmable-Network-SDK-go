# \AuthenticationApi

All URIs are relative to *https://penapi.pacnetconnect.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAnAuthenticationToken**](AuthenticationApi.md#CreateAnAuthenticationToken) | **Post** /1.0.0/auth/generatetoken | Create an authentication token
[**ValidateAuthenticationToken**](AuthenticationApi.md#ValidateAuthenticationToken) | **Get** /1.0.0/auth/validatetoken | Validate authentication token


# **CreateAnAuthenticationToken**
> Model100AuthGeneratetokenResponse CreateAnAuthenticationToken($grantType, $username, $password)

Create an authentication token

Create an authentication token


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string**|  | 
 **username** | **string**|  | 
 **password** | **string**|  | 

### Return type

[**Model100AuthGeneratetokenResponse**](100AuthGeneratetokenResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateAuthenticationToken**
> Model100AuthValidatetokenResponse ValidateAuthenticationToken()

Validate authentication token

Validate the authentication token and get information about the user (roles, permissions, etc.)


### Parameters
This endpoint does not need any parameter.

### Return type

[**Model100AuthValidatetokenResponse**](100AuthValidatetokenResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

