# DepositAddresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** |  | 
**Free** | Pointer to **float32** |  | [optional] [default to 0]
**Locked** | Pointer to **float32** |  | [optional] [default to 0]
**NetworkAddresses** | [**[]NetworkAddress**](NetworkAddress.md) |  | 

## Methods

### NewDepositAddresses

`func NewDepositAddresses(symbol string, networkAddresses []NetworkAddress, ) *DepositAddresses`

NewDepositAddresses instantiates a new DepositAddresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositAddressesWithDefaults

`func NewDepositAddressesWithDefaults() *DepositAddresses`

NewDepositAddressesWithDefaults instantiates a new DepositAddresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *DepositAddresses) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *DepositAddresses) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *DepositAddresses) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetFree

`func (o *DepositAddresses) GetFree() float32`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *DepositAddresses) GetFreeOk() (*float32, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *DepositAddresses) SetFree(v float32)`

SetFree sets Free field to given value.

### HasFree

`func (o *DepositAddresses) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetLocked

`func (o *DepositAddresses) GetLocked() float32`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *DepositAddresses) GetLockedOk() (*float32, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *DepositAddresses) SetLocked(v float32)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *DepositAddresses) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetNetworkAddresses

`func (o *DepositAddresses) GetNetworkAddresses() []NetworkAddress`

GetNetworkAddresses returns the NetworkAddresses field if non-nil, zero value otherwise.

### GetNetworkAddressesOk

`func (o *DepositAddresses) GetNetworkAddressesOk() (*[]NetworkAddress, bool)`

GetNetworkAddressesOk returns a tuple with the NetworkAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddresses

`func (o *DepositAddresses) SetNetworkAddresses(v []NetworkAddress)`

SetNetworkAddresses sets NetworkAddresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


