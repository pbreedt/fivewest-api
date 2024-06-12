# \TransferAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllTransfersApiV1TransferGet**](TransferAPI.md#GetAllTransfersApiV1TransferGet) | **Get** /wallet/api/v1/transfer | Get All Transfers



## GetAllTransfersApiV1TransferGet

> Page GetAllTransfersApiV1TransferGet(ctx).Symbol(symbol).Origin(origin).Destination(destination).External(external).Page(page).PageSize(pageSize).Start(start).End(end).Execute()

Get All Transfers



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
	symbol := "symbol_example" // string |  (optional)
	origin := openapiclient.TransferLocation(0) // TransferLocation |  (optional)
	destination := openapiclient.TransferLocation(0) // TransferLocation |  (optional)
	external := true // bool |  (optional)
	page := int32(56) // int32 |  (optional) (default to 0)
	pageSize := int32(56) // int32 |  (optional) (default to 1)
	start := time.Now() // time.Time |  (optional)
	end := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransferAPI.GetAllTransfersApiV1TransferGet(context.Background()).Symbol(symbol).Origin(origin).Destination(destination).External(external).Page(page).PageSize(pageSize).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferAPI.GetAllTransfersApiV1TransferGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllTransfersApiV1TransferGet`: Page
	fmt.Fprintf(os.Stdout, "Response from `TransferAPI.GetAllTransfersApiV1TransferGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTransfersApiV1TransferGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **origin** | [**TransferLocation**](TransferLocation.md) |  | 
 **destination** | [**TransferLocation**](TransferLocation.md) |  | 
 **external** | **bool** |  | 
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 1]
 **start** | **time.Time** |  | 
 **end** | **time.Time** |  | 

### Return type

[**Page**](Page.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

