# Currency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Id** | **int32** |  | 
**Symbol** | **string** |  | 
**Network** | Pointer to **string** |  | [optional] 
**DepositsEnabled** | Pointer to **bool** |  | [optional] [default to true]
**WithdrawalsEnabled** | Pointer to **bool** |  | [optional] [default to true]
**WithdrawalFee** | Pointer to **float32** |  | [optional] [default to 0]
**MaxWithdraw** | Pointer to **float32** |  | [optional] 
**MinWithdraw** | Pointer to **float32** |  | [optional] [default to 0]
**Meta** | Pointer to [**map[string]Meta**](Meta.md) |  | [optional] [default to {}]
**Active** | Pointer to **bool** |  | [optional] [default to false]
**Fiat** | Pointer to **bool** |  | [optional] [default to false]
**Scale** | Pointer to **int32** |  | [optional] [default to 8]

## Methods

### NewCurrency

`func NewCurrency(id int32, symbol string, ) *Currency`

NewCurrency instantiates a new Currency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyWithDefaults

`func NewCurrencyWithDefaults() *Currency`

NewCurrencyWithDefaults instantiates a new Currency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAt

`func (o *Currency) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Currency) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Currency) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Currency) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Currency) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Currency) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Currency) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Currency) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetId

`func (o *Currency) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Currency) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Currency) SetId(v int32)`

SetId sets Id field to given value.


### GetSymbol

`func (o *Currency) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Currency) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Currency) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetNetwork

`func (o *Currency) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Currency) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Currency) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Currency) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetDepositsEnabled

`func (o *Currency) GetDepositsEnabled() bool`

GetDepositsEnabled returns the DepositsEnabled field if non-nil, zero value otherwise.

### GetDepositsEnabledOk

`func (o *Currency) GetDepositsEnabledOk() (*bool, bool)`

GetDepositsEnabledOk returns a tuple with the DepositsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositsEnabled

`func (o *Currency) SetDepositsEnabled(v bool)`

SetDepositsEnabled sets DepositsEnabled field to given value.

### HasDepositsEnabled

`func (o *Currency) HasDepositsEnabled() bool`

HasDepositsEnabled returns a boolean if a field has been set.

### GetWithdrawalsEnabled

`func (o *Currency) GetWithdrawalsEnabled() bool`

GetWithdrawalsEnabled returns the WithdrawalsEnabled field if non-nil, zero value otherwise.

### GetWithdrawalsEnabledOk

`func (o *Currency) GetWithdrawalsEnabledOk() (*bool, bool)`

GetWithdrawalsEnabledOk returns a tuple with the WithdrawalsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalsEnabled

`func (o *Currency) SetWithdrawalsEnabled(v bool)`

SetWithdrawalsEnabled sets WithdrawalsEnabled field to given value.

### HasWithdrawalsEnabled

`func (o *Currency) HasWithdrawalsEnabled() bool`

HasWithdrawalsEnabled returns a boolean if a field has been set.

### GetWithdrawalFee

`func (o *Currency) GetWithdrawalFee() float32`

GetWithdrawalFee returns the WithdrawalFee field if non-nil, zero value otherwise.

### GetWithdrawalFeeOk

`func (o *Currency) GetWithdrawalFeeOk() (*float32, bool)`

GetWithdrawalFeeOk returns a tuple with the WithdrawalFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalFee

`func (o *Currency) SetWithdrawalFee(v float32)`

SetWithdrawalFee sets WithdrawalFee field to given value.

### HasWithdrawalFee

`func (o *Currency) HasWithdrawalFee() bool`

HasWithdrawalFee returns a boolean if a field has been set.

### GetMaxWithdraw

`func (o *Currency) GetMaxWithdraw() float32`

GetMaxWithdraw returns the MaxWithdraw field if non-nil, zero value otherwise.

### GetMaxWithdrawOk

`func (o *Currency) GetMaxWithdrawOk() (*float32, bool)`

GetMaxWithdrawOk returns a tuple with the MaxWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWithdraw

`func (o *Currency) SetMaxWithdraw(v float32)`

SetMaxWithdraw sets MaxWithdraw field to given value.

### HasMaxWithdraw

`func (o *Currency) HasMaxWithdraw() bool`

HasMaxWithdraw returns a boolean if a field has been set.

### GetMinWithdraw

`func (o *Currency) GetMinWithdraw() float32`

GetMinWithdraw returns the MinWithdraw field if non-nil, zero value otherwise.

### GetMinWithdrawOk

`func (o *Currency) GetMinWithdrawOk() (*float32, bool)`

GetMinWithdrawOk returns a tuple with the MinWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWithdraw

`func (o *Currency) SetMinWithdraw(v float32)`

SetMinWithdraw sets MinWithdraw field to given value.

### HasMinWithdraw

`func (o *Currency) HasMinWithdraw() bool`

HasMinWithdraw returns a boolean if a field has been set.

### GetMeta

`func (o *Currency) GetMeta() map[string]Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Currency) GetMetaOk() (*map[string]Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Currency) SetMeta(v map[string]Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Currency) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetActive

`func (o *Currency) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Currency) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Currency) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Currency) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetFiat

`func (o *Currency) GetFiat() bool`

GetFiat returns the Fiat field if non-nil, zero value otherwise.

### GetFiatOk

`func (o *Currency) GetFiatOk() (*bool, bool)`

GetFiatOk returns a tuple with the Fiat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiat

`func (o *Currency) SetFiat(v bool)`

SetFiat sets Fiat field to given value.

### HasFiat

`func (o *Currency) HasFiat() bool`

HasFiat returns a boolean if a field has been set.

### GetScale

`func (o *Currency) GetScale() int32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *Currency) GetScaleOk() (*int32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *Currency) SetScale(v int32)`

SetScale sets Scale field to given value.

### HasScale

`func (o *Currency) HasScale() bool`

HasScale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


