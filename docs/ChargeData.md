# ChargeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubMerchantUid** | Pointer to **NullableString** |  | [optional] 
**OrderAmount** | [**OrderAmount**](OrderAmount.md) |  | 
**OrderId** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**NullableChargeMetaInput**](ChargeMetaInput.md) |  | [optional] 
**OrderCurrencyId** | Pointer to **NullableString** |  | [optional] 
**Symbol** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewChargeData

`func NewChargeData(orderAmount OrderAmount, ) *ChargeData`

NewChargeData instantiates a new ChargeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeDataWithDefaults

`func NewChargeDataWithDefaults() *ChargeData`

NewChargeDataWithDefaults instantiates a new ChargeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubMerchantUid

`func (o *ChargeData) GetSubMerchantUid() string`

GetSubMerchantUid returns the SubMerchantUid field if non-nil, zero value otherwise.

### GetSubMerchantUidOk

`func (o *ChargeData) GetSubMerchantUidOk() (*string, bool)`

GetSubMerchantUidOk returns a tuple with the SubMerchantUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantUid

`func (o *ChargeData) SetSubMerchantUid(v string)`

SetSubMerchantUid sets SubMerchantUid field to given value.

### HasSubMerchantUid

`func (o *ChargeData) HasSubMerchantUid() bool`

HasSubMerchantUid returns a boolean if a field has been set.

### SetSubMerchantUidNil

`func (o *ChargeData) SetSubMerchantUidNil(b bool)`

 SetSubMerchantUidNil sets the value for SubMerchantUid to be an explicit nil

### UnsetSubMerchantUid
`func (o *ChargeData) UnsetSubMerchantUid()`

UnsetSubMerchantUid ensures that no value is present for SubMerchantUid, not even an explicit nil
### GetOrderAmount

`func (o *ChargeData) GetOrderAmount() OrderAmount`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *ChargeData) GetOrderAmountOk() (*OrderAmount, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *ChargeData) SetOrderAmount(v OrderAmount)`

SetOrderAmount sets OrderAmount field to given value.


### GetOrderId

`func (o *ChargeData) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ChargeData) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ChargeData) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ChargeData) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetMeta

`func (o *ChargeData) GetMeta() ChargeMetaInput`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ChargeData) GetMetaOk() (*ChargeMetaInput, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ChargeData) SetMeta(v ChargeMetaInput)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ChargeData) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *ChargeData) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *ChargeData) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetOrderCurrencyId

`func (o *ChargeData) GetOrderCurrencyId() string`

GetOrderCurrencyId returns the OrderCurrencyId field if non-nil, zero value otherwise.

### GetOrderCurrencyIdOk

`func (o *ChargeData) GetOrderCurrencyIdOk() (*string, bool)`

GetOrderCurrencyIdOk returns a tuple with the OrderCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCurrencyId

`func (o *ChargeData) SetOrderCurrencyId(v string)`

SetOrderCurrencyId sets OrderCurrencyId field to given value.

### HasOrderCurrencyId

`func (o *ChargeData) HasOrderCurrencyId() bool`

HasOrderCurrencyId returns a boolean if a field has been set.

### SetOrderCurrencyIdNil

`func (o *ChargeData) SetOrderCurrencyIdNil(b bool)`

 SetOrderCurrencyIdNil sets the value for OrderCurrencyId to be an explicit nil

### UnsetOrderCurrencyId
`func (o *ChargeData) UnsetOrderCurrencyId()`

UnsetOrderCurrencyId ensures that no value is present for OrderCurrencyId, not even an explicit nil
### GetSymbol

`func (o *ChargeData) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ChargeData) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ChargeData) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ChargeData) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### SetSymbolNil

`func (o *ChargeData) SetSymbolNil(b bool)`

 SetSymbolNil sets the value for Symbol to be an explicit nil

### UnsetSymbol
`func (o *ChargeData) UnsetSymbol()`

UnsetSymbol ensures that no value is present for Symbol, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


