# \WithdrawalRequestAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelWithdrawalRequestApiV1WithdrawalRequestIdPatch**](WithdrawalRequestAPI.md#CancelWithdrawalRequestApiV1WithdrawalRequestIdPatch) | **Patch** /wallet/api/v1/withdrawal_request/{id} | Cancel Withdrawal Request
[**CreateWithdrawalRequestApiV1WithdrawalRequestPost**](WithdrawalRequestAPI.md#CreateWithdrawalRequestApiV1WithdrawalRequestPost) | **Post** /wallet/api/v1/withdrawal_request | Create Withdrawal Request
[**GetAllWithdrawalRequestsApiV1WithdrawalRequestGet**](WithdrawalRequestAPI.md#GetAllWithdrawalRequestsApiV1WithdrawalRequestGet) | **Get** /wallet/api/v1/withdrawal_request | Get All Withdrawal Requests



## CancelWithdrawalRequestApiV1WithdrawalRequestIdPatch

> CancelWithdrawalRequestApiV1WithdrawalRequestIdPatch(ctx, id).Execute()

Cancel Withdrawal Request



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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WithdrawalRequestAPI.CancelWithdrawalRequestApiV1WithdrawalRequestIdPatch(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WithdrawalRequestAPI.CancelWithdrawalRequestApiV1WithdrawalRequestIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelWithdrawalRequestApiV1WithdrawalRequestIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWithdrawalRequestApiV1WithdrawalRequestPost

> WithdrawalRequest CreateWithdrawalRequestApiV1WithdrawalRequestPost(ctx).BaseWithdrawalRequestInsert(baseWithdrawalRequestInsert).Execute()

Create Withdrawal Request



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
	baseWithdrawalRequestInsert := *openapiclient.NewBaseWithdrawalRequestInsert(float32(123), *openapiclient.NewDestination(), int32(123)) // BaseWithdrawalRequestInsert | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WithdrawalRequestAPI.CreateWithdrawalRequestApiV1WithdrawalRequestPost(context.Background()).BaseWithdrawalRequestInsert(baseWithdrawalRequestInsert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WithdrawalRequestAPI.CreateWithdrawalRequestApiV1WithdrawalRequestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWithdrawalRequestApiV1WithdrawalRequestPost`: WithdrawalRequest
	fmt.Fprintf(os.Stdout, "Response from `WithdrawalRequestAPI.CreateWithdrawalRequestApiV1WithdrawalRequestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWithdrawalRequestApiV1WithdrawalRequestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseWithdrawalRequestInsert** | [**BaseWithdrawalRequestInsert**](BaseWithdrawalRequestInsert.md) |  | 

### Return type

[**WithdrawalRequest**](WithdrawalRequest.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllWithdrawalRequestsApiV1WithdrawalRequestGet

> []WithdrawalRequest GetAllWithdrawalRequestsApiV1WithdrawalRequestGet(ctx).CurrencyId(currencyId).Symbol(symbol).Network(network).Status(status).Page(page).PageSize(pageSize).Start(start).End(end).Execute()

Get All Withdrawal Requests



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/pbreedt/fivewestapi"
)

func main() {
	currencyId := int32(56) // int32 |  (optional)
	symbol := "symbol_example" // string |  (optional)
	network := "network_example" // string |  (optional)
	status := openapiclient.WithdrawalRequestStatus(0) // WithdrawalRequestStatus |  (optional)
	page := int32(56) // int32 |  (optional) (default to 0)
	pageSize := int32(56) // int32 |  (optional) (default to 1)
	start := time.Now() // time.Time |  (optional)
	end := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WithdrawalRequestAPI.GetAllWithdrawalRequestsApiV1WithdrawalRequestGet(context.Background()).CurrencyId(currencyId).Symbol(symbol).Network(network).Status(status).Page(page).PageSize(pageSize).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WithdrawalRequestAPI.GetAllWithdrawalRequestsApiV1WithdrawalRequestGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllWithdrawalRequestsApiV1WithdrawalRequestGet`: []WithdrawalRequest
	fmt.Fprintf(os.Stdout, "Response from `WithdrawalRequestAPI.GetAllWithdrawalRequestsApiV1WithdrawalRequestGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllWithdrawalRequestsApiV1WithdrawalRequestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currencyId** | **int32** |  | 
 **symbol** | **string** |  | 
 **network** | **string** |  | 
 **status** | [**WithdrawalRequestStatus**](WithdrawalRequestStatus.md) |  | 
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 1]
 **start** | **time.Time** |  | 
 **end** | **time.Time** |  | 

### Return type

[**[]WithdrawalRequest**](WithdrawalRequest.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

