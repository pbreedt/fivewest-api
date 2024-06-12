# PageTrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Trade**](Trade.md) |  | 
**Total** | **int32** |  | 
**Pages** | **int32** |  | 
**Size** | Pointer to **int32** |  | [optional] 
**HasNext** | Pointer to **bool** |  | [optional] [default to false]
**HasPrev** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewPageTrade

`func NewPageTrade(items []Trade, total int32, pages int32, ) *PageTrade`

NewPageTrade instantiates a new PageTrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageTradeWithDefaults

`func NewPageTradeWithDefaults() *PageTrade`

NewPageTradeWithDefaults instantiates a new PageTrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PageTrade) GetItems() []Trade`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PageTrade) GetItemsOk() (*[]Trade, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PageTrade) SetItems(v []Trade)`

SetItems sets Items field to given value.


### GetTotal

`func (o *PageTrade) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PageTrade) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PageTrade) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPages

`func (o *PageTrade) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *PageTrade) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *PageTrade) SetPages(v int32)`

SetPages sets Pages field to given value.


### GetSize

`func (o *PageTrade) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PageTrade) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PageTrade) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PageTrade) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetHasNext

`func (o *PageTrade) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *PageTrade) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *PageTrade) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *PageTrade) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### GetHasPrev

`func (o *PageTrade) GetHasPrev() bool`

GetHasPrev returns the HasPrev field if non-nil, zero value otherwise.

### GetHasPrevOk

`func (o *PageTrade) GetHasPrevOk() (*bool, bool)`

GetHasPrevOk returns a tuple with the HasPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrev

`func (o *PageTrade) SetHasPrev(v bool)`

SetHasPrev sets HasPrev field to given value.

### HasHasPrev

`func (o *PageTrade) HasHasPrev() bool`

HasHasPrev returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


