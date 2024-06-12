# \DepositAddressAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllDepositAddressesForSymbolApiV1DepositAddressSymbolGet**](DepositAddressAPI.md#GetAllDepositAddressesForSymbolApiV1DepositAddressSymbolGet) | **Get** /wallet/api/v1/deposit_address/{symbol} | Get All Deposit Addresses For Symbol



## GetAllDepositAddressesForSymbolApiV1DepositAddressSymbolGet

> DepositAddresses GetAllDepositAddressesForSymbolApiV1DepositAddressSymbolGet(ctx, symbol).Execute()

Get All Deposit Addresses For Symbol



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
	resp, r, err := apiClient.DepositAddressAPI.GetAllDepositAddressesForSymbolApiV1DepositAddressSymbolGet(context.Background(), symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepositAddressAPI.GetAllDepositAddressesForSymbolApiV1DepositAddressSymbolGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllDepositAddressesForSymbolApiV1DepositAddressSymbolGet`: DepositAddresses
	fmt.Fprintf(os.Stdout, "Response from `DepositAddressAPI.GetAllDepositAddressesForSymbolApiV1DepositAddressSymbolGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDepositAddressesForSymbolApiV1DepositAddressSymbolGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DepositAddresses**](DepositAddresses.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

