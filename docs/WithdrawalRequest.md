# WithdrawalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Id** | **int32** |  | 
**Currency** | [**Currency**](Currency.md) |  | 
**CurrencyId** | **int32** |  | 
**TransferId** | Pointer to **int32** |  | [optional] 
**Uid** | **string** |  | 
**Amount** | **float32** |  | 
**Status** | Pointer to [**WithdrawalRequestStatus**](WithdrawalRequestStatus.md) |  | [optional] [default to WITHDRAWALREQUESTSTATUS__0]
**Destination** | [**Destination**](Destination.md) |  | 
**Wallet** | [**Wallet**](Wallet.md) |  | 
**Reference** | Pointer to **string** |  | [optional] 

## Methods

### NewWithdrawalRequest

`func NewWithdrawalRequest(id int32, currency Currency, currencyId int32, uid string, amount float32, destination Destination, wallet Wallet, ) *WithdrawalRequest`

NewWithdrawalRequest instantiates a new WithdrawalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawalRequestWithDefaults

`func NewWithdrawalRequestWithDefaults() *WithdrawalRequest`

NewWithdrawalRequestWithDefaults instantiates a new WithdrawalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAt

`func (o *WithdrawalRequest) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WithdrawalRequest) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WithdrawalRequest) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WithdrawalRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *WithdrawalRequest) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *WithdrawalRequest) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *WithdrawalRequest) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *WithdrawalRequest) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetId

`func (o *WithdrawalRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WithdrawalRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WithdrawalRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetCurrency

`func (o *WithdrawalRequest) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *WithdrawalRequest) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *WithdrawalRequest) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetCurrencyId

`func (o *WithdrawalRequest) GetCurrencyId() int32`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *WithdrawalRequest) GetCurrencyIdOk() (*int32, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *WithdrawalRequest) SetCurrencyId(v int32)`

SetCurrencyId sets CurrencyId field to given value.


### GetTransferId

`func (o *WithdrawalRequest) GetTransferId() int32`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *WithdrawalRequest) GetTransferIdOk() (*int32, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *WithdrawalRequest) SetTransferId(v int32)`

SetTransferId sets TransferId field to given value.

### HasTransferId

`func (o *WithdrawalRequest) HasTransferId() bool`

HasTransferId returns a boolean if a field has been set.

### GetUid

`func (o *WithdrawalRequest) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *WithdrawalRequest) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *WithdrawalRequest) SetUid(v string)`

SetUid sets Uid field to given value.


### GetAmount

`func (o *WithdrawalRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WithdrawalRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WithdrawalRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetStatus

`func (o *WithdrawalRequest) GetStatus() WithdrawalRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WithdrawalRequest) GetStatusOk() (*WithdrawalRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WithdrawalRequest) SetStatus(v WithdrawalRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WithdrawalRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDestination

`func (o *WithdrawalRequest) GetDestination() Destination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *WithdrawalRequest) GetDestinationOk() (*Destination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *WithdrawalRequest) SetDestination(v Destination)`

SetDestination sets Destination field to given value.


### GetWallet

`func (o *WithdrawalRequest) GetWallet() Wallet`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *WithdrawalRequest) GetWalletOk() (*Wallet, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *WithdrawalRequest) SetWallet(v Wallet)`

SetWallet sets Wallet field to given value.


### GetReference

`func (o *WithdrawalRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *WithdrawalRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *WithdrawalRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *WithdrawalRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


