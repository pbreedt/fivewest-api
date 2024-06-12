# ChargeMetaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GoodsDetails** | Pointer to [**[]GoodsInput**](GoodsInput.md) |  | [optional] [default to [{"goods_type":"01","goods_category":"Z000","reference_goods_id":"5e1fc69a-a4d7-44f6-89f0-8a5952bc8d0e","goods_name":""}]]
**Shipping** | Pointer to [**NullableShippingInput**](ShippingInput.md) |  | [optional] 
**Buyer** | Pointer to [**NullableBuyer**](Buyer.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] [default to "Miscellaneous"]

## Methods

### NewChargeMetaInput

`func NewChargeMetaInput() *ChargeMetaInput`

NewChargeMetaInput instantiates a new ChargeMetaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeMetaInputWithDefaults

`func NewChargeMetaInputWithDefaults() *ChargeMetaInput`

NewChargeMetaInputWithDefaults instantiates a new ChargeMetaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGoodsDetails

`func (o *ChargeMetaInput) GetGoodsDetails() []GoodsInput`

GetGoodsDetails returns the GoodsDetails field if non-nil, zero value otherwise.

### GetGoodsDetailsOk

`func (o *ChargeMetaInput) GetGoodsDetailsOk() (*[]GoodsInput, bool)`

GetGoodsDetailsOk returns a tuple with the GoodsDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodsDetails

`func (o *ChargeMetaInput) SetGoodsDetails(v []GoodsInput)`

SetGoodsDetails sets GoodsDetails field to given value.

### HasGoodsDetails

`func (o *ChargeMetaInput) HasGoodsDetails() bool`

HasGoodsDetails returns a boolean if a field has been set.

### GetShipping

`func (o *ChargeMetaInput) GetShipping() ShippingInput`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *ChargeMetaInput) GetShippingOk() (*ShippingInput, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *ChargeMetaInput) SetShipping(v ShippingInput)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *ChargeMetaInput) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### SetShippingNil

`func (o *ChargeMetaInput) SetShippingNil(b bool)`

 SetShippingNil sets the value for Shipping to be an explicit nil

### UnsetShipping
`func (o *ChargeMetaInput) UnsetShipping()`

UnsetShipping ensures that no value is present for Shipping, not even an explicit nil
### GetBuyer

`func (o *ChargeMetaInput) GetBuyer() Buyer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *ChargeMetaInput) GetBuyerOk() (*Buyer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *ChargeMetaInput) SetBuyer(v Buyer)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *ChargeMetaInput) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### SetBuyerNil

`func (o *ChargeMetaInput) SetBuyerNil(b bool)`

 SetBuyerNil sets the value for Buyer to be an explicit nil

### UnsetBuyer
`func (o *ChargeMetaInput) UnsetBuyer()`

UnsetBuyer ensures that no value is present for Buyer, not even an explicit nil
### GetDescription

`func (o *ChargeMetaInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChargeMetaInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChargeMetaInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChargeMetaInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


