# Balance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Id** | **int32** |  | 
**Uid** | **string** |  | 
**Symbol** | **string** |  | 
**Free** | Pointer to **float32** |  | [optional] [default to 0]
**Locked** | Pointer to **float32** |  | [optional] [default to 0]
**Wallet** | [**Wallet**](Wallet.md) |  | 

## Methods

### NewBalance

`func NewBalance(id int32, uid string, symbol string, wallet Wallet, ) *Balance`

NewBalance instantiates a new Balance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceWithDefaults

`func NewBalanceWithDefaults() *Balance`

NewBalanceWithDefaults instantiates a new Balance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAt

`func (o *Balance) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Balance) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Balance) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Balance) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Balance) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Balance) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Balance) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Balance) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetId

`func (o *Balance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Balance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Balance) SetId(v int32)`

SetId sets Id field to given value.


### GetUid

`func (o *Balance) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Balance) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Balance) SetUid(v string)`

SetUid sets Uid field to given value.


### GetSymbol

`func (o *Balance) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Balance) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Balance) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetFree

`func (o *Balance) GetFree() float32`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *Balance) GetFreeOk() (*float32, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *Balance) SetFree(v float32)`

SetFree sets Free field to given value.

### HasFree

`func (o *Balance) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetLocked

`func (o *Balance) GetLocked() float32`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Balance) GetLockedOk() (*float32, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Balance) SetLocked(v float32)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *Balance) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetWallet

`func (o *Balance) GetWallet() Wallet`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *Balance) GetWalletOk() (*Wallet, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *Balance) SetWallet(v Wallet)`

SetWallet sets Wallet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


