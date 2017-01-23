# \AuthenticateApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAuthenticate**](AuthenticateApi.md#AddAuthenticate) | **Post** /authenticate/ | Authenticate User


# **AddAuthenticate**
> AddAuthenticate($body)

Authenticate User

<b>Note:</b> This request returns an authentication token in the <code>Set-Cookie</code> response header. The token expires after 30 minutes. A valid (that is, unexpired) authentication token must be included in every request to the service, in the <code>Cookie</code>: request header. The client making the API call must examine the cookie expiry time and discard it if the cookie has expired. Requests sent with expired cookies will result in an <code>Unauthorized</code> error in the response.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Authenticate**](Authenticate.md)|  | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

