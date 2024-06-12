# ChargeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**Id** | **string** |  | 
**SubMerchantUid** | Pointer to **NullableString** |  | [optional] 
**OrderId** | **string** |  | 
**OrderAmount** | **string** |  | 
**OrderCurrency** | [**CurrencyResponse**](CurrencyResponse.md) |  | 
**Fee** | **string** |  | 
**Meta** | Pointer to [**NullableChargeMetaOutput**](ChargeMetaOutput.md) |  | [optional] 
**Status** | Pointer to [**ChargeStatus**](ChargeStatus.md) |  | [optional] [default to CHARGESTATUS_ACTIVE]
**SettlementAmount** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewChargeResponse

`func NewChargeResponse(id string, orderId string, orderAmount string, orderCurrency CurrencyResponse, fee string, ) *ChargeResponse`

NewChargeResponse instantiates a new ChargeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeResponseWithDefaults

`func NewChargeResponseWithDefaults() *ChargeResponse`

NewChargeResponseWithDefaults instantiates a new ChargeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAt

`func (o *ChargeResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ChargeResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ChargeResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ChargeResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ChargeResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ChargeResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetId

`func (o *ChargeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChargeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChargeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSubMerchantUid

`func (o *ChargeResponse) GetSubMerchantUid() string`

GetSubMerchantUid returns the SubMerchantUid field if non-nil, zero value otherwise.

### GetSubMerchantUidOk

`func (o *ChargeResponse) GetSubMerchantUidOk() (*string, bool)`

GetSubMerchantUidOk returns a tuple with the SubMerchantUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantUid

`func (o *ChargeResponse) SetSubMerchantUid(v string)`

SetSubMerchantUid sets SubMerchantUid field to given value.

### HasSubMerchantUid

`func (o *ChargeResponse) HasSubMerchantUid() bool`

HasSubMerchantUid returns a boolean if a field has been set.

### SetSubMerchantUidNil

`func (o *ChargeResponse) SetSubMerchantUidNil(b bool)`

 SetSubMerchantUidNil sets the value for SubMerchantUid to be an explicit nil

### UnsetSubMerchantUid
`func (o *ChargeResponse) UnsetSubMerchantUid()`

UnsetSubMerchantUid ensures that no value is present for SubMerchantUid, not even an explicit nil
### GetOrderId

`func (o *ChargeResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ChargeResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ChargeResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetOrderAmount

`func (o *ChargeResponse) GetOrderAmount() string`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *ChargeResponse) GetOrderAmountOk() (*string, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *ChargeResponse) SetOrderAmount(v string)`

SetOrderAmount sets OrderAmount field to given value.


### GetOrderCurrency

`func (o *ChargeResponse) GetOrderCurrency() CurrencyResponse`

GetOrderCurrency returns the OrderCurrency field if non-nil, zero value otherwise.

### GetOrderCurrencyOk

`func (o *ChargeResponse) GetOrderCurrencyOk() (*CurrencyResponse, bool)`

GetOrderCurrencyOk returns a tuple with the OrderCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCurrency

`func (o *ChargeResponse) SetOrderCurrency(v CurrencyResponse)`

SetOrderCurrency sets OrderCurrency field to given value.


### GetFee

`func (o *ChargeResponse) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ChargeResponse) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ChargeResponse) SetFee(v string)`

SetFee sets Fee field to given value.


### GetMeta

`func (o *ChargeResponse) GetMeta() ChargeMetaOutput`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ChargeResponse) GetMetaOk() (*ChargeMetaOutput, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ChargeResponse) SetMeta(v ChargeMetaOutput)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ChargeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *ChargeResponse) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *ChargeResponse) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetStatus

`func (o *ChargeResponse) GetStatus() ChargeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChargeResponse) GetStatusOk() (*ChargeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChargeResponse) SetStatus(v ChargeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChargeResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSettlementAmount

`func (o *ChargeResponse) GetSettlementAmount() string`

GetSettlementAmount returns the SettlementAmount field if non-nil, zero value otherwise.

### GetSettlementAmountOk

`func (o *ChargeResponse) GetSettlementAmountOk() (*string, bool)`

GetSettlementAmountOk returns a tuple with the SettlementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementAmount

`func (o *ChargeResponse) SetSettlementAmount(v string)`

SetSettlementAmount sets SettlementAmount field to given value.

### HasSettlementAmount

`func (o *ChargeResponse) HasSettlementAmount() bool`

HasSettlementAmount returns a boolean if a field has been set.

### SetSettlementAmountNil

`func (o *ChargeResponse) SetSettlementAmountNil(b bool)`

 SetSettlementAmountNil sets the value for SettlementAmount to be an explicit nil

### UnsetSettlementAmount
`func (o *ChargeResponse) UnsetSettlementAmount()`

UnsetSettlementAmount ensures that no value is present for SettlementAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


