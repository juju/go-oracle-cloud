# \RefreshTokenApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRefreshToken**](RefreshTokenApi.md#GetRefreshToken) | **Get** /refresh/ | Refresh an Authentication Token


# **GetRefreshToken**
> GetRefreshToken()

Refresh an Authentication Token

Authentication tokens expire in 30 minutes. This request extends the expiry of the authentication token by 30 minutes from the time you run the command. It extends the expiry of the current authentication token, but not beyond the session expiry time, which is 3 hours.


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

