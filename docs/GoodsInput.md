# GoodsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GoodsType** | [**GoodsType**](GoodsType.md) |  | 
**GoodsCategory** | [**GoodsCategory**](GoodsCategory.md) |  | 
**ReferenceGoodsId** | Pointer to **string** |  | [optional] 
**GoodsName** | **string** |  | 
**GoodsDetail** | Pointer to **NullableString** |  | [optional] 
**GoodsUnitAmount** | Pointer to [**NullableGoodsUnitAmountInput**](GoodsUnitAmountInput.md) |  | [optional] 

## Methods

### NewGoodsInput

`func NewGoodsInput(goodsType GoodsType, goodsCategory GoodsCategory, goodsName string, ) *GoodsInput`

NewGoodsInput instantiates a new GoodsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoodsInputWithDefaults

`func NewGoodsInputWithDefaults() *GoodsInput`

NewGoodsInputWithDefaults instantiates a new GoodsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGoodsType

`func (o *GoodsInput) GetGoodsType() GoodsType`

GetGoodsType returns the GoodsType field if non-nil, zero value otherwise.

### GetGoodsTypeOk

`func (o *GoodsInput) GetGoodsTypeOk() (*GoodsType, bool)`

GetGoodsTypeOk returns a tuple with the GoodsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodsType

`func (o *GoodsInput) SetGoodsType(v GoodsType)`

SetGoodsType sets GoodsType field to given value.


### GetGoodsCategory

`func (o *GoodsInput) GetGoodsCategory() GoodsCategory`

GetGoodsCategory returns the GoodsCategory field if non-nil, zero value otherwise.

### GetGoodsCategoryOk

`func (o *GoodsInput) GetGoodsCategoryOk() (*GoodsCategory, bool)`

GetGoodsCategoryOk returns a tuple with the GoodsCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodsCategory

`func (o *GoodsInput) SetGoodsCategory(v GoodsCategory)`

SetGoodsCategory sets GoodsCategory field to given value.


### GetReferenceGoodsId

`func (o *GoodsInput) GetReferenceGoodsId() string`

GetReferenceGoodsId returns the ReferenceGoodsId field if non-nil, zero value otherwise.

### GetReferenceGoodsIdOk

`func (o *GoodsInput) GetReferenceGoodsIdOk() (*string, bool)`

GetReferenceGoodsIdOk returns a tuple with the ReferenceGoodsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceGoodsId

`func (o *GoodsInput) SetReferenceGoodsId(v string)`

SetReferenceGoodsId sets ReferenceGoodsId field to given value.

### HasReferenceGoodsId

`func (o *GoodsInput) HasReferenceGoodsId() bool`

HasReferenceGoodsId returns a boolean if a field has been set.

### GetGoodsName

`func (o *GoodsInput) GetGoodsName() string`

GetGoodsName returns the GoodsName field if non-nil, zero value otherwise.

### GetGoodsNameOk

`func (o *GoodsInput) GetGoodsNameOk() (*string, bool)`

GetGoodsNameOk returns a tuple with the GoodsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodsName

`func (o *GoodsInput) SetGoodsName(v string)`

SetGoodsName sets GoodsName field to given value.


### GetGoodsDetail

`func (o *GoodsInput) GetGoodsDetail() string`

GetGoodsDetail returns the GoodsDetail field if non-nil, zero value otherwise.

### GetGoodsDetailOk

`func (o *GoodsInput) GetGoodsDetailOk() (*string, bool)`

GetGoodsDetailOk returns a tuple with the GoodsDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodsDetail

`func (o *GoodsInput) SetGoodsDetail(v string)`

SetGoodsDetail sets GoodsDetail field to given value.

### HasGoodsDetail

`func (o *GoodsInput) HasGoodsDetail() bool`

HasGoodsDetail returns a boolean if a field has been set.

### SetGoodsDetailNil

`func (o *GoodsInput) SetGoodsDetailNil(b bool)`

 SetGoodsDetailNil sets the value for GoodsDetail to be an explicit nil

### UnsetGoodsDetail
`func (o *GoodsInput) UnsetGoodsDetail()`

UnsetGoodsDetail ensures that no value is present for GoodsDetail, not even an explicit nil
### GetGoodsUnitAmount

`func (o *GoodsInput) GetGoodsUnitAmount() GoodsUnitAmountInput`

GetGoodsUnitAmount returns the GoodsUnitAmount field if non-nil, zero value otherwise.

### GetGoodsUnitAmountOk

`func (o *GoodsInput) GetGoodsUnitAmountOk() (*GoodsUnitAmountInput, bool)`

GetGoodsUnitAmountOk returns a tuple with the GoodsUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodsUnitAmount

`func (o *GoodsInput) SetGoodsUnitAmount(v GoodsUnitAmountInput)`

SetGoodsUnitAmount sets GoodsUnitAmount field to given value.

### HasGoodsUnitAmount

`func (o *GoodsInput) HasGoodsUnitAmount() bool`

HasGoodsUnitAmount returns a boolean if a field has been set.

### SetGoodsUnitAmountNil

`func (o *GoodsInput) SetGoodsUnitAmountNil(b bool)`

 SetGoodsUnitAmountNil sets the value for GoodsUnitAmount to be an explicit nil

### UnsetGoodsUnitAmount
`func (o *GoodsInput) UnsetGoodsUnitAmount()`

UnsetGoodsUnitAmount ensures that no value is present for GoodsUnitAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


