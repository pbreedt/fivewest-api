# \CurrencyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllApiV1CurrencyGet**](CurrencyAPI.md#GetAllApiV1CurrencyGet) | **Get** /payments/api/v1/currency | Get All
[**GetAllSupportedCurrencyInformationApiV1CurrencyGet**](CurrencyAPI.md#GetAllSupportedCurrencyInformationApiV1CurrencyGet) | **Get** /wallet/api/v1/currency | Get All Supported Currency Information
[**GetAllSupportedSymbolsApiV1CurrencySymbolsGet**](CurrencyAPI.md#GetAllSupportedSymbolsApiV1CurrencySymbolsGet) | **Get** /wallet/api/v1/currency/symbols | Get All Supported Symbols
[**GetApiV1CurrencyIdGet**](CurrencyAPI.md#GetApiV1CurrencyIdGet) | **Get** /payments/api/v1/currency/{id} | Get
[**SymbolsApiV1CurrencySymbolsGet**](CurrencyAPI.md#SymbolsApiV1CurrencySymbolsGet) | **Get** /payments/api/v1/currency/symbols | Symbols



## GetAllApiV1CurrencyGet

> []CurrencyResponse GetAllApiV1CurrencyGet(ctx).Page(page).PageSize(pageSize).Symbol(symbol).Network(network).Fiat(fiat).DefaultsOnly(defaultsOnly).Execute()

Get All



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
	network := "network_example" // string |  (optional)
	fiat := true // bool |  (optional)
	defaultsOnly := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrencyAPI.GetAllApiV1CurrencyGet(context.Background()).Page(page).PageSize(pageSize).Symbol(symbol).Network(network).Fiat(fiat).DefaultsOnly(defaultsOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.GetAllApiV1CurrencyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllApiV1CurrencyGet`: []CurrencyResponse
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.GetAllApiV1CurrencyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllApiV1CurrencyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 1]
 **symbol** | **string** |  | 
 **network** | **string** |  | 
 **fiat** | **bool** |  | 
 **defaultsOnly** | **bool** |  | 

### Return type

[**[]CurrencyResponse**](CurrencyResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSupportedCurrencyInformationApiV1CurrencyGet

> []Currency GetAllSupportedCurrencyInformationApiV1CurrencyGet(ctx).Page(page).PageSize(pageSize).Symbol(symbol).Network(network).Fiat(fiat).Execute()

Get All Supported Currency Information



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
	network := "network_example" // string |  (optional)
	fiat := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrencyAPI.GetAllSupportedCurrencyInformationApiV1CurrencyGet(context.Background()).Page(page).PageSize(pageSize).Symbol(symbol).Network(network).Fiat(fiat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.GetAllSupportedCurrencyInformationApiV1CurrencyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllSupportedCurrencyInformationApiV1CurrencyGet`: []Currency
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.GetAllSupportedCurrencyInformationApiV1CurrencyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSupportedCurrencyInformationApiV1CurrencyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 1]
 **symbol** | **string** |  | 
 **network** | **string** |  | 
 **fiat** | **bool** |  | 

### Return type

[**[]Currency**](Currency.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSupportedSymbolsApiV1CurrencySymbolsGet

> []string GetAllSupportedSymbolsApiV1CurrencySymbolsGet(ctx).Network(network).Fiat(fiat).Execute()

Get All Supported Symbols



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
	network := "network_example" // string |  (optional)
	fiat := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrencyAPI.GetAllSupportedSymbolsApiV1CurrencySymbolsGet(context.Background()).Network(network).Fiat(fiat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.GetAllSupportedSymbolsApiV1CurrencySymbolsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllSupportedSymbolsApiV1CurrencySymbolsGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.GetAllSupportedSymbolsApiV1CurrencySymbolsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSupportedSymbolsApiV1CurrencySymbolsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **network** | **string** |  | 
 **fiat** | **bool** |  | 

### Return type

**[]string**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1CurrencyIdGet

> CurrencyResponse GetApiV1CurrencyIdGet(ctx, id).Execute()

Get



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
	resp, r, err := apiClient.CurrencyAPI.GetApiV1CurrencyIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.GetApiV1CurrencyIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1CurrencyIdGet`: CurrencyResponse
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.GetApiV1CurrencyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1CurrencyIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CurrencyResponse**](CurrencyResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SymbolsApiV1CurrencySymbolsGet

> []string SymbolsApiV1CurrencySymbolsGet(ctx).Network(network).Fiat(fiat).Execute()

Symbols



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
	network := "network_example" // string |  (optional)
	fiat := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrencyAPI.SymbolsApiV1CurrencySymbolsGet(context.Background()).Network(network).Fiat(fiat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.SymbolsApiV1CurrencySymbolsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SymbolsApiV1CurrencySymbolsGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.SymbolsApiV1CurrencySymbolsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSymbolsApiV1CurrencySymbolsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **network** | **string** |  | 
 **fiat** | **bool** |  | 

### Return type

**[]string**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

