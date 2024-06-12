# CurrencyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**Id** | **string** |  | 
**Symbol** | **string** |  | 
**Network** | **string** |  | 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Fiat** | Pointer to **bool** |  | [optional] [default to false]
**ExpiryDelta** | Pointer to **int32** |  | [optional] 
**Scale** | **int32** |  | 

## Methods

### NewCurrencyResponse

`func NewCurrencyResponse(id string, symbol string, network string, scale int32, ) *CurrencyResponse`

NewCurrencyResponse instantiates a new CurrencyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyResponseWithDefaults

`func NewCurrencyResponseWithDefaults() *CurrencyResponse`

NewCurrencyResponseWithDefaults instantiates a new CurrencyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAt

`func (o *CurrencyResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CurrencyResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CurrencyResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CurrencyResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *CurrencyResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *CurrencyResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetId

`func (o *CurrencyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurrencyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurrencyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSymbol

`func (o *CurrencyResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CurrencyResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CurrencyResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetNetwork

`func (o *CurrencyResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CurrencyResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CurrencyResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetMeta

`func (o *CurrencyResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CurrencyResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CurrencyResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CurrencyResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *CurrencyResponse) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *CurrencyResponse) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetFiat

`func (o *CurrencyResponse) GetFiat() bool`

GetFiat returns the Fiat field if non-nil, zero value otherwise.

### GetFiatOk

`func (o *CurrencyResponse) GetFiatOk() (*bool, bool)`

GetFiatOk returns a tuple with the Fiat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiat

`func (o *CurrencyResponse) SetFiat(v bool)`

SetFiat sets Fiat field to given value.

### HasFiat

`func (o *CurrencyResponse) HasFiat() bool`

HasFiat returns a boolean if a field has been set.

### GetExpiryDelta

`func (o *CurrencyResponse) GetExpiryDelta() int32`

GetExpiryDelta returns the ExpiryDelta field if non-nil, zero value otherwise.

### GetExpiryDeltaOk

`func (o *CurrencyResponse) GetExpiryDeltaOk() (*int32, bool)`

GetExpiryDeltaOk returns a tuple with the ExpiryDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDelta

`func (o *CurrencyResponse) SetExpiryDelta(v int32)`

SetExpiryDelta sets ExpiryDelta field to given value.

### HasExpiryDelta

`func (o *CurrencyResponse) HasExpiryDelta() bool`

HasExpiryDelta returns a boolean if a field has been set.

### GetScale

`func (o *CurrencyResponse) GetScale() int32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *CurrencyResponse) GetScaleOk() (*int32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *CurrencyResponse) SetScale(v int32)`

SetScale sets Scale field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


