# \AuthenticationApi

All URIs are relative to *https://penapi.pacnetconnect.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**100AuthGeneratetokenPost**](AuthenticationApi.md#100AuthGeneratetokenPost) | **Post** /1.0.0/auth/generatetoken | Create an authentication token
[**100AuthValidatetokenGet**](AuthenticationApi.md#100AuthValidatetokenGet) | **Get** /1.0.0/auth/validatetoken | Validate authentication token


# **100AuthGeneratetokenPost**
> Model100AuthGeneratetokenResponse 100AuthGeneratetokenPost($grantType, $username, $password)

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

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **100AuthValidatetokenGet**
> Model100AuthValidatetokenResponse 100AuthValidatetokenGet()

Validate authentication token

Validate the authentication token and get information about the user (roles, permissions, etc.)


### Parameters
This endpoint does not need any parameter.

### Return type

[**Model100AuthValidatetokenResponse**](100AuthValidatetokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

