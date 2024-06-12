# \RatesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadLatestApiV1RatesLatestGet**](RatesAPI.md#ReadLatestApiV1RatesLatestGet) | **Get** /markets/api/v1/rates/latest | Read Latest
[**StreamLatestApiV1RatesStreamLatestGet**](RatesAPI.md#StreamLatestApiV1RatesStreamLatestGet) | **Get** /markets/api/v1/rates/stream/latest | Stream Latest



## ReadLatestApiV1RatesLatestGet

> map[string][]interface{} ReadLatestApiV1RatesLatestGet(ctx).Base(base).Quote(quote).Symbol(symbol).Exchange(exchange).Start(start).Execute()

Read Latest



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
	base := "base_example" // string |  (optional)
	quote := "quote_example" // string |  (optional)
	symbol := "symbol_example" // string |  (optional)
	exchange := openapiclient.Exchange("ftx") // Exchange |  (optional)
	start := float32(8.14) // float32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RatesAPI.ReadLatestApiV1RatesLatestGet(context.Background()).Base(base).Quote(quote).Symbol(symbol).Exchange(exchange).Start(start).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatesAPI.ReadLatestApiV1RatesLatestGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadLatestApiV1RatesLatestGet`: map[string][]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RatesAPI.ReadLatestApiV1RatesLatestGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadLatestApiV1RatesLatestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **base** | **string** |  | 
 **quote** | **string** |  | 
 **symbol** | **string** |  | 
 **exchange** | [**Exchange**](Exchange.md) |  | 
 **start** | **float32** |  | 

### Return type

[**map[string][]interface{}**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamLatestApiV1RatesStreamLatestGet

> MarketRate StreamLatestApiV1RatesStreamLatestGet(ctx).Base(base).Quote(quote).Exchange(exchange).Execute()

Stream Latest

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
	base := "base_example" // string | 
	quote := "quote_example" // string | 
	exchange := "exchange_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RatesAPI.StreamLatestApiV1RatesStreamLatestGet(context.Background()).Base(base).Quote(quote).Exchange(exchange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatesAPI.StreamLatestApiV1RatesStreamLatestGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamLatestApiV1RatesStreamLatestGet`: MarketRate
	fmt.Fprintf(os.Stdout, "Response from `RatesAPI.StreamLatestApiV1RatesStreamLatestGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStreamLatestApiV1RatesStreamLatestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **base** | **string** |  | 
 **quote** | **string** |  | 
 **exchange** | **string** |  | 

### Return type

[**MarketRate**](MarketRate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

