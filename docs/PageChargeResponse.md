# PageChargeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ChargeResponse**](ChargeResponse.md) |  | 
**Total** | **int32** |  | 
**Pages** | **int32** |  | 
**PageSize** | **int32** |  | 
**HasNext** | Pointer to **bool** |  | [optional] [default to false]
**HasPrev** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewPageChargeResponse

`func NewPageChargeResponse(items []ChargeResponse, total int32, pages int32, pageSize int32, ) *PageChargeResponse`

NewPageChargeResponse instantiates a new PageChargeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageChargeResponseWithDefaults

`func NewPageChargeResponseWithDefaults() *PageChargeResponse`

NewPageChargeResponseWithDefaults instantiates a new PageChargeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PageChargeResponse) GetItems() []ChargeResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PageChargeResponse) GetItemsOk() (*[]ChargeResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PageChargeResponse) SetItems(v []ChargeResponse)`

SetItems sets Items field to given value.


### GetTotal

`func (o *PageChargeResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PageChargeResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PageChargeResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPages

`func (o *PageChargeResponse) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *PageChargeResponse) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *PageChargeResponse) SetPages(v int32)`

SetPages sets Pages field to given value.


### GetPageSize

`func (o *PageChargeResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PageChargeResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PageChargeResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetHasNext

`func (o *PageChargeResponse) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *PageChargeResponse) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *PageChargeResponse) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *PageChargeResponse) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### GetHasPrev

`func (o *PageChargeResponse) GetHasPrev() bool`

GetHasPrev returns the HasPrev field if non-nil, zero value otherwise.

### GetHasPrevOk

`func (o *PageChargeResponse) GetHasPrevOk() (*bool, bool)`

GetHasPrevOk returns a tuple with the HasPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrev

`func (o *PageChargeResponse) SetHasPrev(v bool)`

SetHasPrev sets HasPrev field to given value.

### HasHasPrev

`func (o *PageChargeResponse) HasHasPrev() bool`

HasHasPrev returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


