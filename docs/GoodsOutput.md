# GoodsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GoodsType** | [**GoodsType**](GoodsType.md) |  | 
**GoodsCategory** | [**GoodsCategory**](GoodsCategory.md) |  | 
**ReferenceGoodsId** | Pointer to **string** |  | [optional] 
**GoodsName** | **string** |  | 
**GoodsDetail** | Pointer to **NullableString** |  | [optional] 
**GoodsUnitAmount** | Pointer to [**NullableGoodsUnitAmountOutput**](GoodsUnitAmountOutput.md) |  | [optional] 

## Methods

### NewGoodsOutput

`func NewGoodsOutput(goodsType GoodsType, goodsCategory GoodsCategory, goodsName string, ) *GoodsOutput`

NewGoodsOutput instantiates a new GoodsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoodsOutputWithDefaults

`func NewGoodsOutputWithDefaults() *GoodsOutput`

NewGoodsOutputWithDefaults instantiates a new GoodsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGoodsType

`func (o *GoodsOutput) GetGoodsType() GoodsType`

GetGoodsType returns the GoodsType field if non-nil, zero value otherwise.

### GetGoodsTypeOk

`func (o *GoodsOutput) GetGoodsTypeOk() (*GoodsType, bool)`

GetGoodsTypeOk returns a tuple with the GoodsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodsType

`func (o *GoodsOutput) SetGoodsType(v GoodsType)`

SetGoodsType sets GoodsType field to given value.


### GetGoodsCategory

`func (o *GoodsOutput) GetGoodsCategory() GoodsCategory`

GetGoodsCategory returns the GoodsCategory field if non-nil, zero value otherwise.

### GetGoodsCategoryOk

`func (o *GoodsOutput) GetGoodsCategoryOk() (*GoodsCategory, bool)`

GetGoodsCategoryOk returns a tuple with the GoodsCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodsCategory

`func (o *GoodsOutput) SetGoodsCategory(v GoodsCategory)`

SetGoodsCategory sets GoodsCategory field to given value.


### GetReferenceGoodsId

`func (o *GoodsOutput) GetReferenceGoodsId() string`

GetReferenceGoodsId returns the ReferenceGoodsId field if non-nil, zero value otherwise.

### GetReferenceGoodsIdOk

`func (o *GoodsOutput) GetReferenceGoodsIdOk() (*string, bool)`

GetReferenceGoodsIdOk returns a tuple with the ReferenceGoodsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceGoodsId

`func (o *GoodsOutput) SetReferenceGoodsId(v string)`

SetReferenceGoodsId sets ReferenceGoodsId field to given value.

### HasReferenceGoodsId

`func (o *GoodsOutput) HasReferenceGoodsId() bool`

HasReferenceGoodsId returns a boolean if a field has been set.

### GetGoodsName

`func (o *GoodsOutput) GetGoodsName() string`

GetGoodsName returns the GoodsName field if non-nil, zero value otherwise.

### GetGoodsNameOk

`func (o *GoodsOutput) GetGoodsNameOk() (*string, bool)`

GetGoodsNameOk returns a tuple with the GoodsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodsName

`func (o *GoodsOutput) SetGoodsName(v string)`

SetGoodsName sets GoodsName field to given value.


### GetGoodsDetail

`func (o *GoodsOutput) GetGoodsDetail() string`

GetGoodsDetail returns the GoodsDetail field if non-nil, zero value otherwise.

### GetGoodsDetailOk

`func (o *GoodsOutput) GetGoodsDetailOk() (*string, bool)`

GetGoodsDetailOk returns a tuple with the GoodsDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodsDetail

`func (o *GoodsOutput) SetGoodsDetail(v string)`

SetGoodsDetail sets GoodsDetail field to given value.

### HasGoodsDetail

`func (o *GoodsOutput) HasGoodsDetail() bool`

HasGoodsDetail returns a boolean if a field has been set.

### SetGoodsDetailNil

`func (o *GoodsOutput) SetGoodsDetailNil(b bool)`

 SetGoodsDetailNil sets the value for GoodsDetail to be an explicit nil

### UnsetGoodsDetail
`func (o *GoodsOutput) UnsetGoodsDetail()`

UnsetGoodsDetail ensures that no value is present for GoodsDetail, not even an explicit nil
### GetGoodsUnitAmount

`func (o *GoodsOutput) GetGoodsUnitAmount() GoodsUnitAmountOutput`

GetGoodsUnitAmount returns the GoodsUnitAmount field if non-nil, zero value otherwise.

### GetGoodsUnitAmountOk

`func (o *GoodsOutput) GetGoodsUnitAmountOk() (*GoodsUnitAmountOutput, bool)`

GetGoodsUnitAmountOk returns a tuple with the GoodsUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodsUnitAmount

`func (o *GoodsOutput) SetGoodsUnitAmount(v GoodsUnitAmountOutput)`

SetGoodsUnitAmount sets GoodsUnitAmount field to given value.

### HasGoodsUnitAmount

`func (o *GoodsOutput) HasGoodsUnitAmount() bool`

HasGoodsUnitAmount returns a boolean if a field has been set.

### SetGoodsUnitAmountNil

`func (o *GoodsOutput) SetGoodsUnitAmountNil(b bool)`

 SetGoodsUnitAmountNil sets the value for GoodsUnitAmount to be an explicit nil

### UnsetGoodsUnitAmount
`func (o *GoodsOutput) UnsetGoodsUnitAmount()`

UnsetGoodsUnitAmount ensures that no value is present for GoodsUnitAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


