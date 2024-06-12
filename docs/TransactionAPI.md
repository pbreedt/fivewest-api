# \TransactionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1TransactionPost**](TransactionAPI.md#CreateApiV1TransactionPost) | **Post** /payments/api/v1/transaction | Create
[**GetAllApiV1TransactionGet**](TransactionAPI.md#GetAllApiV1TransactionGet) | **Get** /payments/api/v1/transaction | Get All
[**GetOneApiV1TransactionTransactionIdGet**](TransactionAPI.md#GetOneApiV1TransactionTransactionIdGet) | **Get** /payments/api/v1/transaction/{transaction_id} | Get One
[**GetOrUpdateApiV1TransactionPut**](TransactionAPI.md#GetOrUpdateApiV1TransactionPut) | **Put** /payments/api/v1/transaction | Get Or Update



## CreateApiV1TransactionPost

> TransactionResponse CreateApiV1TransactionPost(ctx).ChargeId(chargeId).Symbol(symbol).Network(network).CurrencyId(currencyId).Execute()

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
	chargeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	symbol := "symbol_example" // string |  (optional)
	network := "network_example" // string |  (optional)
	currencyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.CreateApiV1TransactionPost(context.Background()).ChargeId(chargeId).Symbol(symbol).Network(network).CurrencyId(currencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.CreateApiV1TransactionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApiV1TransactionPost`: TransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.CreateApiV1TransactionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1TransactionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chargeId** | **string** |  | 
 **symbol** | **string** |  | 
 **network** | **string** |  | 
 **currencyId** | **string** |  | 

### Return type

[**TransactionResponse**](TransactionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllApiV1TransactionGet

> PageTransactionResponse GetAllApiV1TransactionGet(ctx).ChargeId(chargeId).CurrencyId(currencyId).Symbol(symbol).Network(network).Status(status).Late(late).Page(page).PageSize(pageSize).Start(start).End(end).Execute()

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
	chargeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	currencyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	symbol := "symbol_example" // string |  (optional)
	network := "network_example" // string |  (optional)
	status := openapiclient.TransactionStatus("complete") // TransactionStatus |  (optional)
	late := true // bool |  (optional)
	page := int32(56) // int32 |  (optional) (default to 0)
	pageSize := int32(56) // int32 |  (optional) (default to 1)
	start := time.Now() // time.Time |  (optional)
	end := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.GetAllApiV1TransactionGet(context.Background()).ChargeId(chargeId).CurrencyId(currencyId).Symbol(symbol).Network(network).Status(status).Late(late).Page(page).PageSize(pageSize).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.GetAllApiV1TransactionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllApiV1TransactionGet`: PageTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.GetAllApiV1TransactionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllApiV1TransactionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chargeId** | **string** |  | 
 **currencyId** | **string** |  | 
 **symbol** | **string** |  | 
 **network** | **string** |  | 
 **status** | [**TransactionStatus**](TransactionStatus.md) |  | 
 **late** | **bool** |  | 
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 1]
 **start** | **time.Time** |  | 
 **end** | **time.Time** |  | 

### Return type

[**PageTransactionResponse**](PageTransactionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOneApiV1TransactionTransactionIdGet

> TransactionResponse GetOneApiV1TransactionTransactionIdGet(ctx, transactionId).Execute()

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
	transactionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.GetOneApiV1TransactionTransactionIdGet(context.Background(), transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.GetOneApiV1TransactionTransactionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOneApiV1TransactionTransactionIdGet`: TransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.GetOneApiV1TransactionTransactionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOneApiV1TransactionTransactionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransactionResponse**](TransactionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrUpdateApiV1TransactionPut

> TransactionResponse GetOrUpdateApiV1TransactionPut(ctx).TransactionId(transactionId).ChargeId(chargeId).Symbol(symbol).Network(network).CurrencyId(currencyId).Execute()

Get Or Update



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
	transactionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	chargeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	symbol := "symbol_example" // string |  (optional)
	network := "network_example" // string |  (optional)
	currencyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.GetOrUpdateApiV1TransactionPut(context.Background()).TransactionId(transactionId).ChargeId(chargeId).Symbol(symbol).Network(network).CurrencyId(currencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.GetOrUpdateApiV1TransactionPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrUpdateApiV1TransactionPut`: TransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.GetOrUpdateApiV1TransactionPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrUpdateApiV1TransactionPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **string** |  | 
 **chargeId** | **string** |  | 
 **symbol** | **string** |  | 
 **network** | **string** |  | 
 **currencyId** | **string** |  | 

### Return type

[**TransactionResponse**](TransactionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

