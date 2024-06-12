# ShippingAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** |  | 
**State** | Pointer to **NullableString** |  | [optional] 
**City** | Pointer to **NullableString** |  | [optional] 
**Address** | Pointer to **NullableString** |  | [optional] 
**ZipCode** | Pointer to **NullableString** |  | [optional] 
**ShippingAddressType** | Pointer to [**NullableShippingAddressType**](ShippingAddressType.md) |  | [optional] 

## Methods

### NewShippingAddress

`func NewShippingAddress(region string, ) *ShippingAddress`

NewShippingAddress instantiates a new ShippingAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingAddressWithDefaults

`func NewShippingAddressWithDefaults() *ShippingAddress`

NewShippingAddressWithDefaults instantiates a new ShippingAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *ShippingAddress) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ShippingAddress) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ShippingAddress) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetState

`func (o *ShippingAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ShippingAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ShippingAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ShippingAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ShippingAddress) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ShippingAddress) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCity

`func (o *ShippingAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ShippingAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ShippingAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ShippingAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *ShippingAddress) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *ShippingAddress) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetAddress

`func (o *ShippingAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ShippingAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ShippingAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ShippingAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *ShippingAddress) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *ShippingAddress) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetZipCode

`func (o *ShippingAddress) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *ShippingAddress) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *ShippingAddress) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *ShippingAddress) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *ShippingAddress) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *ShippingAddress) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
### GetShippingAddressType

`func (o *ShippingAddress) GetShippingAddressType() ShippingAddressType`

GetShippingAddressType returns the ShippingAddressType field if non-nil, zero value otherwise.

### GetShippingAddressTypeOk

`func (o *ShippingAddress) GetShippingAddressTypeOk() (*ShippingAddressType, bool)`

GetShippingAddressTypeOk returns a tuple with the ShippingAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddressType

`func (o *ShippingAddress) SetShippingAddressType(v ShippingAddressType)`

SetShippingAddressType sets ShippingAddressType field to given value.

### HasShippingAddressType

`func (o *ShippingAddress) HasShippingAddressType() bool`

HasShippingAddressType returns a boolean if a field has been set.

### SetShippingAddressTypeNil

`func (o *ShippingAddress) SetShippingAddressTypeNil(b bool)`

 SetShippingAddressTypeNil sets the value for ShippingAddressType to be an explicit nil

### UnsetShippingAddressType
`func (o *ShippingAddress) UnsetShippingAddressType()`

UnsetShippingAddressType ensures that no value is present for ShippingAddressType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


