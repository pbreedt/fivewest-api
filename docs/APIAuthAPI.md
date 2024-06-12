# \APIAuthAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoginApiV1AuthTokenPost**](APIAuthAPI.md#LoginApiV1AuthTokenPost) | **Post** /profile/api/v1/auth/token | Get OAuth2 access token from API key and secret



## LoginApiV1AuthTokenPost

> Token LoginApiV1AuthTokenPost(ctx).OAuth2ClientCredentialsRequestForm(oAuth2ClientCredentialsRequestForm).Execute()

Get OAuth2 access token from API key and secret



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pbreedt/fivewestapi"
)

func main() {
	oAuth2ClientCredentialsRequestForm := *openapiclient.NewOAuth2ClientCredentialsRequestForm("Key_example", "Secret_example") // OAuth2ClientCredentialsRequestForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIAuthAPI.LoginApiV1AuthTokenPost(context.Background()).OAuth2ClientCredentialsRequestForm(oAuth2ClientCredentialsRequestForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIAuthAPI.LoginApiV1AuthTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoginApiV1AuthTokenPost`: Token
	fmt.Fprintf(os.Stdout, "Response from `APIAuthAPI.LoginApiV1AuthTokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginApiV1AuthTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oAuth2ClientCredentialsRequestForm** | [**OAuth2ClientCredentialsRequestForm**](OAuth2ClientCredentialsRequestForm.md) |  | 

### Return type

[**Token**](Token.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

