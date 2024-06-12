# BaseWithdrawalRequestInsert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** |  | 
**Destination** | [**Destination**](Destination.md) |  | 
**CurrencyId** | **int32** |  | 

## Methods

### NewBaseWithdrawalRequestInsert

`func NewBaseWithdrawalRequestInsert(amount float32, destination Destination, currencyId int32, ) *BaseWithdrawalRequestInsert`

NewBaseWithdrawalRequestInsert instantiates a new BaseWithdrawalRequestInsert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseWithdrawalRequestInsertWithDefaults

`func NewBaseWithdrawalRequestInsertWithDefaults() *BaseWithdrawalRequestInsert`

NewBaseWithdrawalRequestInsertWithDefaults instantiates a new BaseWithdrawalRequestInsert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BaseWithdrawalRequestInsert) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BaseWithdrawalRequestInsert) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BaseWithdrawalRequestInsert) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDestination

`func (o *BaseWithdrawalRequestInsert) GetDestination() Destination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *BaseWithdrawalRequestInsert) GetDestinationOk() (*Destination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *BaseWithdrawalRequestInsert) SetDestination(v Destination)`

SetDestination sets Destination field to given value.


### GetCurrencyId

`func (o *BaseWithdrawalRequestInsert) GetCurrencyId() int32`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *BaseWithdrawalRequestInsert) GetCurrencyIdOk() (*int32, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *BaseWithdrawalRequestInsert) SetCurrencyId(v int32)`

SetCurrencyId sets CurrencyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


