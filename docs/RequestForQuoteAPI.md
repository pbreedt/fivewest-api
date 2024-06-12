# \RequestForQuoteAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptQuoteApiV1RfqAcceptQuoteIdPost**](RequestForQuoteAPI.md#AcceptQuoteApiV1RfqAcceptQuoteIdPost) | **Post** /trading/api/v1/rfq/accept_quote/{id} | Accept Quote
[**GenerateQuoteApiV1RfqPost**](RequestForQuoteAPI.md#GenerateQuoteApiV1RfqPost) | **Post** /trading/api/v1/rfq | Generate Quote
[**GetAllQuoteRequestsApiV1RfqGet**](RequestForQuoteAPI.md#GetAllQuoteRequestsApiV1RfqGet) | **Get** /trading/api/v1/rfq | Get All Quote Requests
[**GetAllTradableCurrencyPairsApiV1RfqPairsBidAskGet**](RequestForQuoteAPI.md#GetAllTradableCurrencyPairsApiV1RfqPairsBidAskGet) | **Get** /trading/api/v1/rfq/pairs_bid_ask | Get All Tradable Currency Pairs
[**GetQuoteRequestApiV1RfqIdGet**](RequestForQuoteAPI.md#GetQuoteRequestApiV1RfqIdGet) | **Get** /trading/api/v1/rfq/{id} | Get Quote Request



## AcceptQuoteApiV1RfqAcceptQuoteIdPost

> Rfq AcceptQuoteApiV1RfqAcceptQuoteIdPost(ctx, id).Execute()

Accept Quote



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
	resp, r, err := apiClient.RequestForQuoteAPI.AcceptQuoteApiV1RfqAcceptQuoteIdPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestForQuoteAPI.AcceptQuoteApiV1RfqAcceptQuoteIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcceptQuoteApiV1RfqAcceptQuoteIdPost`: Rfq
	fmt.Fprintf(os.Stdout, "Response from `RequestForQuoteAPI.AcceptQuoteApiV1RfqAcceptQuoteIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptQuoteApiV1RfqAcceptQuoteIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Rfq**](Rfq.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateQuoteApiV1RfqPost

> Rfq GenerateQuoteApiV1RfqPost(ctx).RfqRequest(rfqRequest).Execute()

Generate Quote



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
	rfqRequest := *openapiclient.NewRfqRequest("FromCurrency_example", "ToCurrency_example") // RfqRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestForQuoteAPI.GenerateQuoteApiV1RfqPost(context.Background()).RfqRequest(rfqRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestForQuoteAPI.GenerateQuoteApiV1RfqPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateQuoteApiV1RfqPost`: Rfq
	fmt.Fprintf(os.Stdout, "Response from `RequestForQuoteAPI.GenerateQuoteApiV1RfqPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateQuoteApiV1RfqPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rfqRequest** | [**RfqRequest**](RfqRequest.md) |  | 

### Return type

[**Rfq**](Rfq.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllQuoteRequestsApiV1RfqGet

> PageRfq GetAllQuoteRequestsApiV1RfqGet(ctx).QuoteStatus(quoteStatus).Page(page).PageSize(pageSize).Execute()

Get All Quote Requests



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
	quoteStatus := openapiclient.QuoteStatus("accepted") // QuoteStatus |  (optional)
	page := int32(56) // int32 |  (optional) (default to 0)
	pageSize := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestForQuoteAPI.GetAllQuoteRequestsApiV1RfqGet(context.Background()).QuoteStatus(quoteStatus).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestForQuoteAPI.GetAllQuoteRequestsApiV1RfqGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllQuoteRequestsApiV1RfqGet`: PageRfq
	fmt.Fprintf(os.Stdout, "Response from `RequestForQuoteAPI.GetAllQuoteRequestsApiV1RfqGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllQuoteRequestsApiV1RfqGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quoteStatus** | [**QuoteStatus**](QuoteStatus.md) |  | 
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 1]

### Return type

[**PageRfq**](PageRfq.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllTradableCurrencyPairsApiV1RfqPairsBidAskGet

> []BidAsk GetAllTradableCurrencyPairsApiV1RfqPairsBidAskGet(ctx).Execute()

Get All Tradable Currency Pairs



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestForQuoteAPI.GetAllTradableCurrencyPairsApiV1RfqPairsBidAskGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestForQuoteAPI.GetAllTradableCurrencyPairsApiV1RfqPairsBidAskGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllTradableCurrencyPairsApiV1RfqPairsBidAskGet`: []BidAsk
	fmt.Fprintf(os.Stdout, "Response from `RequestForQuoteAPI.GetAllTradableCurrencyPairsApiV1RfqPairsBidAskGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTradableCurrencyPairsApiV1RfqPairsBidAskGetRequest struct via the builder pattern


### Return type

[**[]BidAsk**](BidAsk.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuoteRequestApiV1RfqIdGet

> Rfq GetQuoteRequestApiV1RfqIdGet(ctx, id).Execute()

Get Quote Request



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
	resp, r, err := apiClient.RequestForQuoteAPI.GetQuoteRequestApiV1RfqIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestForQuoteAPI.GetQuoteRequestApiV1RfqIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuoteRequestApiV1RfqIdGet`: Rfq
	fmt.Fprintf(os.Stdout, "Response from `RequestForQuoteAPI.GetQuoteRequestApiV1RfqIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteRequestApiV1RfqIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Rfq**](Rfq.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

