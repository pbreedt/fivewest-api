# \ViewTradesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllTradesApiV1TradeGet**](ViewTradesAPI.md#GetAllTradesApiV1TradeGet) | **Get** /trading/api/v1/trade | Get All Trades
[**GetTradeByIDApiV1TradeIdGet**](ViewTradesAPI.md#GetTradeByIDApiV1TradeIdGet) | **Get** /trading/api/v1/trade/{id} | Get Trade By Id



## GetAllTradesApiV1TradeGet

> PageTrade GetAllTradesApiV1TradeGet(ctx).BaseCurrency(baseCurrency).QuoteCurrency(quoteCurrency).Type_(type_).Mode(mode).StartTime(startTime).EndTime(endTime).Page(page).PageSize(pageSize).Execute()

Get All Trades



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
	baseCurrency := "baseCurrency_example" // string |  (optional)
	quoteCurrency := "quoteCurrency_example" // string |  (optional)
	type_ := openapiclient.TradeType("pts") // TradeType |  (optional)
	mode := openapiclient.Mode("otc") // Mode |  (optional)
	startTime := time.Now() // time.Time |  (optional)
	endTime := time.Now() // time.Time |  (optional)
	page := int32(56) // int32 |  (optional) (default to 0)
	pageSize := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewTradesAPI.GetAllTradesApiV1TradeGet(context.Background()).BaseCurrency(baseCurrency).QuoteCurrency(quoteCurrency).Type_(type_).Mode(mode).StartTime(startTime).EndTime(endTime).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewTradesAPI.GetAllTradesApiV1TradeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllTradesApiV1TradeGet`: PageTrade
	fmt.Fprintf(os.Stdout, "Response from `ViewTradesAPI.GetAllTradesApiV1TradeGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTradesApiV1TradeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseCurrency** | **string** |  | 
 **quoteCurrency** | **string** |  | 
 **type_** | [**TradeType**](TradeType.md) |  | 
 **mode** | [**Mode**](Mode.md) |  | 
 **startTime** | **time.Time** |  | 
 **endTime** | **time.Time** |  | 
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 1]

### Return type

[**PageTrade**](PageTrade.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradeByIDApiV1TradeIdGet

> []Trade GetTradeByIDApiV1TradeIdGet(ctx, id).Execute()

Get Trade By Id



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewTradesAPI.GetTradeByIDApiV1TradeIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewTradesAPI.GetTradeByIDApiV1TradeIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradeByIDApiV1TradeIdGet`: []Trade
	fmt.Fprintf(os.Stdout, "Response from `ViewTradesAPI.GetTradeByIDApiV1TradeIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeByIDApiV1TradeIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Trade**](Trade.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

