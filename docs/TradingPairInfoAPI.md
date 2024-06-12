# \TradingPairInfoAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllTradingPairsApiV1PairGet**](TradingPairInfoAPI.md#GetAllTradingPairsApiV1PairGet) | **Get** /trading/api/v1/pair | Get All Trading Pairs



## GetAllTradingPairsApiV1PairGet

> []PairInfo GetAllTradingPairsApiV1PairGet(ctx).BaseCurrency(baseCurrency).QuoteCurrency(quoteCurrency).IsActive(isActive).Page(page).PageSize(pageSize).Execute()

Get All Trading Pairs



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
	baseCurrency := "baseCurrency_example" // string |  (optional)
	quoteCurrency := "quoteCurrency_example" // string |  (optional)
	isActive := true // bool |  (optional)
	page := int32(56) // int32 |  (optional) (default to 0)
	pageSize := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingPairInfoAPI.GetAllTradingPairsApiV1PairGet(context.Background()).BaseCurrency(baseCurrency).QuoteCurrency(quoteCurrency).IsActive(isActive).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingPairInfoAPI.GetAllTradingPairsApiV1PairGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllTradingPairsApiV1PairGet`: []PairInfo
	fmt.Fprintf(os.Stdout, "Response from `TradingPairInfoAPI.GetAllTradingPairsApiV1PairGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTradingPairsApiV1PairGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseCurrency** | **string** |  | 
 **quoteCurrency** | **string** |  | 
 **isActive** | **bool** |  | 
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 1]

### Return type

[**[]PairInfo**](PairInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

