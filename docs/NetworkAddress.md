# NetworkAddress

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
**Address** | **string** |  | 

## Methods

### NewNetworkAddress

`func NewNetworkAddress(id int32, symbol string, address string, ) *NetworkAddress`

NewNetworkAddress instantiates a new NetworkAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkAddressWithDefaults

`func NewNetworkAddressWithDefaults() *NetworkAddress`

NewNetworkAddressWithDefaults instantiates a new NetworkAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAt

`func (o *NetworkAddress) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkAddress) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkAddress) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NetworkAddress) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *NetworkAddress) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *NetworkAddress) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *NetworkAddress) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *NetworkAddress) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetId

`func (o *NetworkAddress) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkAddress) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkAddress) SetId(v int32)`

SetId sets Id field to given value.


### GetSymbol

`func (o *NetworkAddress) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *NetworkAddress) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *NetworkAddress) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetNetwork

`func (o *NetworkAddress) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkAddress) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkAddress) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *NetworkAddress) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetDepositsEnabled

`func (o *NetworkAddress) GetDepositsEnabled() bool`

GetDepositsEnabled returns the DepositsEnabled field if non-nil, zero value otherwise.

### GetDepositsEnabledOk

`func (o *NetworkAddress) GetDepositsEnabledOk() (*bool, bool)`

GetDepositsEnabledOk returns a tuple with the DepositsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositsEnabled

`func (o *NetworkAddress) SetDepositsEnabled(v bool)`

SetDepositsEnabled sets DepositsEnabled field to given value.

### HasDepositsEnabled

`func (o *NetworkAddress) HasDepositsEnabled() bool`

HasDepositsEnabled returns a boolean if a field has been set.

### GetWithdrawalsEnabled

`func (o *NetworkAddress) GetWithdrawalsEnabled() bool`

GetWithdrawalsEnabled returns the WithdrawalsEnabled field if non-nil, zero value otherwise.

### GetWithdrawalsEnabledOk

`func (o *NetworkAddress) GetWithdrawalsEnabledOk() (*bool, bool)`

GetWithdrawalsEnabledOk returns a tuple with the WithdrawalsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalsEnabled

`func (o *NetworkAddress) SetWithdrawalsEnabled(v bool)`

SetWithdrawalsEnabled sets WithdrawalsEnabled field to given value.

### HasWithdrawalsEnabled

`func (o *NetworkAddress) HasWithdrawalsEnabled() bool`

HasWithdrawalsEnabled returns a boolean if a field has been set.

### GetWithdrawalFee

`func (o *NetworkAddress) GetWithdrawalFee() float32`

GetWithdrawalFee returns the WithdrawalFee field if non-nil, zero value otherwise.

### GetWithdrawalFeeOk

`func (o *NetworkAddress) GetWithdrawalFeeOk() (*float32, bool)`

GetWithdrawalFeeOk returns a tuple with the WithdrawalFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalFee

`func (o *NetworkAddress) SetWithdrawalFee(v float32)`

SetWithdrawalFee sets WithdrawalFee field to given value.

### HasWithdrawalFee

`func (o *NetworkAddress) HasWithdrawalFee() bool`

HasWithdrawalFee returns a boolean if a field has been set.

### GetMaxWithdraw

`func (o *NetworkAddress) GetMaxWithdraw() float32`

GetMaxWithdraw returns the MaxWithdraw field if non-nil, zero value otherwise.

### GetMaxWithdrawOk

`func (o *NetworkAddress) GetMaxWithdrawOk() (*float32, bool)`

GetMaxWithdrawOk returns a tuple with the MaxWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWithdraw

`func (o *NetworkAddress) SetMaxWithdraw(v float32)`

SetMaxWithdraw sets MaxWithdraw field to given value.

### HasMaxWithdraw

`func (o *NetworkAddress) HasMaxWithdraw() bool`

HasMaxWithdraw returns a boolean if a field has been set.

### GetMinWithdraw

`func (o *NetworkAddress) GetMinWithdraw() float32`

GetMinWithdraw returns the MinWithdraw field if non-nil, zero value otherwise.

### GetMinWithdrawOk

`func (o *NetworkAddress) GetMinWithdrawOk() (*float32, bool)`

GetMinWithdrawOk returns a tuple with the MinWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWithdraw

`func (o *NetworkAddress) SetMinWithdraw(v float32)`

SetMinWithdraw sets MinWithdraw field to given value.

### HasMinWithdraw

`func (o *NetworkAddress) HasMinWithdraw() bool`

HasMinWithdraw returns a boolean if a field has been set.

### GetMeta

`func (o *NetworkAddress) GetMeta() map[string]Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NetworkAddress) GetMetaOk() (*map[string]Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NetworkAddress) SetMeta(v map[string]Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *NetworkAddress) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetActive

`func (o *NetworkAddress) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *NetworkAddress) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *NetworkAddress) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *NetworkAddress) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetFiat

`func (o *NetworkAddress) GetFiat() bool`

GetFiat returns the Fiat field if non-nil, zero value otherwise.

### GetFiatOk

`func (o *NetworkAddress) GetFiatOk() (*bool, bool)`

GetFiatOk returns a tuple with the Fiat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiat

`func (o *NetworkAddress) SetFiat(v bool)`

SetFiat sets Fiat field to given value.

### HasFiat

`func (o *NetworkAddress) HasFiat() bool`

HasFiat returns a boolean if a field has been set.

### GetScale

`func (o *NetworkAddress) GetScale() int32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *NetworkAddress) GetScaleOk() (*int32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *NetworkAddress) SetScale(v int32)`

SetScale sets Scale field to given value.

### HasScale

`func (o *NetworkAddress) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetAddress

`func (o *NetworkAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NetworkAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NetworkAddress) SetAddress(v string)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


