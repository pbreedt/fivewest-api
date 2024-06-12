# PageTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]TransactionResponse**](TransactionResponse.md) |  | 
**Total** | **int32** |  | 
**Pages** | **int32** |  | 
**PageSize** | **int32** |  | 
**HasNext** | Pointer to **bool** |  | [optional] [default to false]
**HasPrev** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewPageTransactionResponse

`func NewPageTransactionResponse(items []TransactionResponse, total int32, pages int32, pageSize int32, ) *PageTransactionResponse`

NewPageTransactionResponse instantiates a new PageTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageTransactionResponseWithDefaults

`func NewPageTransactionResponseWithDefaults() *PageTransactionResponse`

NewPageTransactionResponseWithDefaults instantiates a new PageTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PageTransactionResponse) GetItems() []TransactionResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PageTransactionResponse) GetItemsOk() (*[]TransactionResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PageTransactionResponse) SetItems(v []TransactionResponse)`

SetItems sets Items field to given value.


### GetTotal

`func (o *PageTransactionResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PageTransactionResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PageTransactionResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPages

`func (o *PageTransactionResponse) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *PageTransactionResponse) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *PageTransactionResponse) SetPages(v int32)`

SetPages sets Pages field to given value.


### GetPageSize

`func (o *PageTransactionResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PageTransactionResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PageTransactionResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetHasNext

`func (o *PageTransactionResponse) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *PageTransactionResponse) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *PageTransactionResponse) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *PageTransactionResponse) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### GetHasPrev

`func (o *PageTransactionResponse) GetHasPrev() bool`

GetHasPrev returns the HasPrev field if non-nil, zero value otherwise.

### GetHasPrevOk

`func (o *PageTransactionResponse) GetHasPrevOk() (*bool, bool)`

GetHasPrevOk returns a tuple with the HasPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrev

`func (o *PageTransactionResponse) SetHasPrev(v bool)`

SetHasPrev sets HasPrev field to given value.

### HasHasPrev

`func (o *PageTransactionResponse) HasHasPrev() bool`

HasHasPrev returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


