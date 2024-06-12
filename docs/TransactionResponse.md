# TransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**Id** | **string** |  | 
**PaymentAmount** | **string** |  | 
**Status** | Pointer to [**TransactionStatus**](TransactionStatus.md) |  | [optional] [default to TRANSACTIONSTATUS_PENDING]
**Currency** | [**CurrencyResponse**](CurrencyResponse.md) |  | 
**ChargeId** | **string** |  | 
**Late** | **bool** |  | 
**Hash** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**SentAmount** | Pointer to **NullableString** |  | [optional] 
**PayInstructions** | [**PayInstructions**](PayInstructions.md) |  | 

## Methods

### NewTransactionResponse

`func NewTransactionResponse(id string, paymentAmount string, currency CurrencyResponse, chargeId string, late bool, createdAt time.Time, payInstructions PayInstructions, ) *TransactionResponse`

NewTransactionResponse instantiates a new TransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionResponseWithDefaults

`func NewTransactionResponseWithDefaults() *TransactionResponse`

NewTransactionResponseWithDefaults instantiates a new TransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAt

`func (o *TransactionResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TransactionResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TransactionResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TransactionResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *TransactionResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *TransactionResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetId

`func (o *TransactionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetPaymentAmount

`func (o *TransactionResponse) GetPaymentAmount() string`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *TransactionResponse) GetPaymentAmountOk() (*string, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *TransactionResponse) SetPaymentAmount(v string)`

SetPaymentAmount sets PaymentAmount field to given value.


### GetStatus

`func (o *TransactionResponse) GetStatus() TransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionResponse) GetStatusOk() (*TransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionResponse) SetStatus(v TransactionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransactionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCurrency

`func (o *TransactionResponse) GetCurrency() CurrencyResponse`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionResponse) GetCurrencyOk() (*CurrencyResponse, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionResponse) SetCurrency(v CurrencyResponse)`

SetCurrency sets Currency field to given value.


### GetChargeId

`func (o *TransactionResponse) GetChargeId() string`

GetChargeId returns the ChargeId field if non-nil, zero value otherwise.

### GetChargeIdOk

`func (o *TransactionResponse) GetChargeIdOk() (*string, bool)`

GetChargeIdOk returns a tuple with the ChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeId

`func (o *TransactionResponse) SetChargeId(v string)`

SetChargeId sets ChargeId field to given value.


### GetLate

`func (o *TransactionResponse) GetLate() bool`

GetLate returns the Late field if non-nil, zero value otherwise.

### GetLateOk

`func (o *TransactionResponse) GetLateOk() (*bool, bool)`

GetLateOk returns a tuple with the Late field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLate

`func (o *TransactionResponse) SetLate(v bool)`

SetLate sets Late field to given value.


### GetHash

`func (o *TransactionResponse) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *TransactionResponse) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *TransactionResponse) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *TransactionResponse) HasHash() bool`

HasHash returns a boolean if a field has been set.

### SetHashNil

`func (o *TransactionResponse) SetHashNil(b bool)`

 SetHashNil sets the value for Hash to be an explicit nil

### UnsetHash
`func (o *TransactionResponse) UnsetHash()`

UnsetHash ensures that no value is present for Hash, not even an explicit nil
### GetCreatedAt

`func (o *TransactionResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransactionResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransactionResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSentAmount

`func (o *TransactionResponse) GetSentAmount() string`

GetSentAmount returns the SentAmount field if non-nil, zero value otherwise.

### GetSentAmountOk

`func (o *TransactionResponse) GetSentAmountOk() (*string, bool)`

GetSentAmountOk returns a tuple with the SentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAmount

`func (o *TransactionResponse) SetSentAmount(v string)`

SetSentAmount sets SentAmount field to given value.

### HasSentAmount

`func (o *TransactionResponse) HasSentAmount() bool`

HasSentAmount returns a boolean if a field has been set.

### SetSentAmountNil

`func (o *TransactionResponse) SetSentAmountNil(b bool)`

 SetSentAmountNil sets the value for SentAmount to be an explicit nil

### UnsetSentAmount
`func (o *TransactionResponse) UnsetSentAmount()`

UnsetSentAmount ensures that no value is present for SentAmount, not even an explicit nil
### GetPayInstructions

`func (o *TransactionResponse) GetPayInstructions() PayInstructions`

GetPayInstructions returns the PayInstructions field if non-nil, zero value otherwise.

### GetPayInstructionsOk

`func (o *TransactionResponse) GetPayInstructionsOk() (*PayInstructions, bool)`

GetPayInstructionsOk returns a tuple with the PayInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayInstructions

`func (o *TransactionResponse) SetPayInstructions(v PayInstructions)`

SetPayInstructions sets PayInstructions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


