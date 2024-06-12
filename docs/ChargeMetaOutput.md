# ChargeMetaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GoodsDetails** | Pointer to [**[]GoodsOutput**](GoodsOutput.md) |  | [optional] [default to [{"goods_type":"01","goods_category":"Z000","reference_goods_id":"5e1fc69a-a4d7-44f6-89f0-8a5952bc8d0e","goods_name":""}]]
**Shipping** | Pointer to [**NullableShippingOutput**](ShippingOutput.md) |  | [optional] 
**Buyer** | Pointer to [**NullableBuyer**](Buyer.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] [default to "Miscellaneous"]

## Methods

### NewChargeMetaOutput

`func NewChargeMetaOutput() *ChargeMetaOutput`

NewChargeMetaOutput instantiates a new ChargeMetaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeMetaOutputWithDefaults

`func NewChargeMetaOutputWithDefaults() *ChargeMetaOutput`

NewChargeMetaOutputWithDefaults instantiates a new ChargeMetaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGoodsDetails

`func (o *ChargeMetaOutput) GetGoodsDetails() []GoodsOutput`

GetGoodsDetails returns the GoodsDetails field if non-nil, zero value otherwise.

### GetGoodsDetailsOk

`func (o *ChargeMetaOutput) GetGoodsDetailsOk() (*[]GoodsOutput, bool)`

GetGoodsDetailsOk returns a tuple with the GoodsDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodsDetails

`func (o *ChargeMetaOutput) SetGoodsDetails(v []GoodsOutput)`

SetGoodsDetails sets GoodsDetails field to given value.

### HasGoodsDetails

`func (o *ChargeMetaOutput) HasGoodsDetails() bool`

HasGoodsDetails returns a boolean if a field has been set.

### GetShipping

`func (o *ChargeMetaOutput) GetShipping() ShippingOutput`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *ChargeMetaOutput) GetShippingOk() (*ShippingOutput, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *ChargeMetaOutput) SetShipping(v ShippingOutput)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *ChargeMetaOutput) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### SetShippingNil

`func (o *ChargeMetaOutput) SetShippingNil(b bool)`

 SetShippingNil sets the value for Shipping to be an explicit nil

### UnsetShipping
`func (o *ChargeMetaOutput) UnsetShipping()`

UnsetShipping ensures that no value is present for Shipping, not even an explicit nil
### GetBuyer

`func (o *ChargeMetaOutput) GetBuyer() Buyer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *ChargeMetaOutput) GetBuyerOk() (*Buyer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *ChargeMetaOutput) SetBuyer(v Buyer)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *ChargeMetaOutput) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### SetBuyerNil

`func (o *ChargeMetaOutput) SetBuyerNil(b bool)`

 SetBuyerNil sets the value for Buyer to be an explicit nil

### UnsetBuyer
`func (o *ChargeMetaOutput) UnsetBuyer()`

UnsetBuyer ensures that no value is present for Buyer, not even an explicit nil
### GetDescription

`func (o *ChargeMetaOutput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChargeMetaOutput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChargeMetaOutput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChargeMetaOutput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


