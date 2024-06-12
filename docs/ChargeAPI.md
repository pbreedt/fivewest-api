# \ChargeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1ChargePost**](ChargeAPI.md#CreateApiV1ChargePost) | **Post** /payments/api/v1/charge | Create
[**DeactivateApiV1ChargeIdPatch**](ChargeAPI.md#DeactivateApiV1ChargeIdPatch) | **Patch** /payments/api/v1/charge/{id} | Deactivate
[**GetAllApiV1ChargeGet**](ChargeAPI.md#GetAllApiV1ChargeGet) | **Get** /payments/api/v1/charge | Get All
[**GetOneApiV1ChargeIdGet**](ChargeAPI.md#GetOneApiV1ChargeIdGet) | **Get** /payments/api/v1/charge/{id} | Get One



## CreateApiV1ChargePost

> ChargeResponse CreateApiV1ChargePost(ctx).ChargeData(chargeData).Execute()

Create



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
	chargeData := *openapiclient.NewChargeData(*openapiclient.NewOrderAmount()) // ChargeData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChargeAPI.CreateApiV1ChargePost(context.Background()).ChargeData(chargeData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargeAPI.CreateApiV1ChargePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApiV1ChargePost`: ChargeResponse
	fmt.Fprintf(os.Stdout, "Response from `ChargeAPI.CreateApiV1ChargePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1ChargePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chargeData** | [**ChargeData**](ChargeData.md) |  | 

### Return type

[**ChargeResponse**](ChargeResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateApiV1ChargeIdPatch

> DeactivateApiV1ChargeIdPatch(ctx, id).Execute()

Deactivate



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
	id := *openapiclient.NewId1() // Id1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChargeAPI.DeactivateApiV1ChargeIdPatch(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargeAPI.DeactivateApiV1ChargeIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**Id1**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateApiV1ChargeIdPatchRequest struct via the builder pattern


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


## GetAllApiV1ChargeGet

> PageChargeResponse GetAllApiV1ChargeGet(ctx).OrderCurrencyId(orderCurrencyId).Status(status).Page(page).PageSize(pageSize).TransactionCurrencyId(transactionCurrencyId).TransactionStatus(transactionStatus).TransactionHash(transactionHash).Start(start).End(end).Execute()

Get All



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
	orderCurrencyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	status := openapiclient.ChargeStatus("complete") // ChargeStatus |  (optional)
	page := int32(56) // int32 |  (optional) (default to 0)
	pageSize := int32(56) // int32 |  (optional) (default to 1)
	transactionCurrencyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	transactionStatus := openapiclient.TransactionStatus("complete") // TransactionStatus |  (optional)
	transactionHash := "transactionHash_example" // string |  (optional)
	start := time.Now() // time.Time |  (optional)
	end := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChargeAPI.GetAllApiV1ChargeGet(context.Background()).OrderCurrencyId(orderCurrencyId).Status(status).Page(page).PageSize(pageSize).TransactionCurrencyId(transactionCurrencyId).TransactionStatus(transactionStatus).TransactionHash(transactionHash).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargeAPI.GetAllApiV1ChargeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllApiV1ChargeGet`: PageChargeResponse
	fmt.Fprintf(os.Stdout, "Response from `ChargeAPI.GetAllApiV1ChargeGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllApiV1ChargeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderCurrencyId** | **string** |  | 
 **status** | [**ChargeStatus**](ChargeStatus.md) |  | 
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 1]
 **transactionCurrencyId** | **string** |  | 
 **transactionStatus** | [**TransactionStatus**](TransactionStatus.md) |  | 
 **transactionHash** | **string** |  | 
 **start** | **time.Time** |  | 
 **end** | **time.Time** |  | 

### Return type

[**PageChargeResponse**](PageChargeResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOneApiV1ChargeIdGet

> ChargeResponse GetOneApiV1ChargeIdGet(ctx, id).Execute()

Get One



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
	id := *openapiclient.NewId() // Id | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChargeAPI.GetOneApiV1ChargeIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargeAPI.GetOneApiV1ChargeIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOneApiV1ChargeIdGet`: ChargeResponse
	fmt.Fprintf(os.Stdout, "Response from `ChargeAPI.GetOneApiV1ChargeIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**Id**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOneApiV1ChargeIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChargeResponse**](ChargeResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

