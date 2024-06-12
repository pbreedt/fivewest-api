# \BalanceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllWalletBalancesApiV1BalanceGet**](BalanceAPI.md#GetAllWalletBalancesApiV1BalanceGet) | **Get** /wallet/api/v1/balance | Get All Wallet Balances
[**GetBalanceForSymbolApiV1BalanceSymbolGet**](BalanceAPI.md#GetBalanceForSymbolApiV1BalanceSymbolGet) | **Get** /wallet/api/v1/balance/{symbol} | Get Balance For Symbol



## GetAllWalletBalancesApiV1BalanceGet

> []Balance GetAllWalletBalancesApiV1BalanceGet(ctx).Page(page).PageSize(pageSize).Symbol(symbol).Execute()

Get All Wallet Balances



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
	page := int32(56) // int32 |  (optional) (default to 0)
	pageSize := int32(56) // int32 |  (optional) (default to 1)
	symbol := "symbol_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BalanceAPI.GetAllWalletBalancesApiV1BalanceGet(context.Background()).Page(page).PageSize(pageSize).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BalanceAPI.GetAllWalletBalancesApiV1BalanceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllWalletBalancesApiV1BalanceGet`: []Balance
	fmt.Fprintf(os.Stdout, "Response from `BalanceAPI.GetAllWalletBalancesApiV1BalanceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllWalletBalancesApiV1BalanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 1]
 **symbol** | **string** |  | 

### Return type

[**[]Balance**](Balance.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalanceForSymbolApiV1BalanceSymbolGet

> Balance GetBalanceForSymbolApiV1BalanceSymbolGet(ctx, symbol).Execute()

Get Balance For Symbol



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
	symbol := "symbol_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BalanceAPI.GetBalanceForSymbolApiV1BalanceSymbolGet(context.Background(), symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BalanceAPI.GetBalanceForSymbolApiV1BalanceSymbolGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalanceForSymbolApiV1BalanceSymbolGet`: Balance
	fmt.Fprintf(os.Stdout, "Response from `BalanceAPI.GetBalanceForSymbolApiV1BalanceSymbolGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceForSymbolApiV1BalanceSymbolGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Balance**](Balance.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

