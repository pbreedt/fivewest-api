# PairInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCurrency** | **string** |  | 
**QuoteCurrency** | **string** |  | 
**MinSizeQuoteAmount** | **float32** |  | 
**MaxSizeQuoteAmount** | **float32** |  | 
**RfqDailyLimitQuoteAmount** | **float32** |  | 
**IsActive** | **bool** |  | 
**QuoteLifetimeSeconds** | **int32** |  | 

## Methods

### NewPairInfo

`func NewPairInfo(baseCurrency string, quoteCurrency string, minSizeQuoteAmount float32, maxSizeQuoteAmount float32, rfqDailyLimitQuoteAmount float32, isActive bool, quoteLifetimeSeconds int32, ) *PairInfo`

NewPairInfo instantiates a new PairInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPairInfoWithDefaults

`func NewPairInfoWithDefaults() *PairInfo`

NewPairInfoWithDefaults instantiates a new PairInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseCurrency

`func (o *PairInfo) GetBaseCurrency() string`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *PairInfo) GetBaseCurrencyOk() (*string, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *PairInfo) SetBaseCurrency(v string)`

SetBaseCurrency sets BaseCurrency field to given value.


### GetQuoteCurrency

`func (o *PairInfo) GetQuoteCurrency() string`

GetQuoteCurrency returns the QuoteCurrency field if non-nil, zero value otherwise.

### GetQuoteCurrencyOk

`func (o *PairInfo) GetQuoteCurrencyOk() (*string, bool)`

GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCurrency

`func (o *PairInfo) SetQuoteCurrency(v string)`

SetQuoteCurrency sets QuoteCurrency field to given value.


### GetMinSizeQuoteAmount

`func (o *PairInfo) GetMinSizeQuoteAmount() float32`

GetMinSizeQuoteAmount returns the MinSizeQuoteAmount field if non-nil, zero value otherwise.

### GetMinSizeQuoteAmountOk

`func (o *PairInfo) GetMinSizeQuoteAmountOk() (*float32, bool)`

GetMinSizeQuoteAmountOk returns a tuple with the MinSizeQuoteAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSizeQuoteAmount

`func (o *PairInfo) SetMinSizeQuoteAmount(v float32)`

SetMinSizeQuoteAmount sets MinSizeQuoteAmount field to given value.


### GetMaxSizeQuoteAmount

`func (o *PairInfo) GetMaxSizeQuoteAmount() float32`

GetMaxSizeQuoteAmount returns the MaxSizeQuoteAmount field if non-nil, zero value otherwise.

### GetMaxSizeQuoteAmountOk

`func (o *PairInfo) GetMaxSizeQuoteAmountOk() (*float32, bool)`

GetMaxSizeQuoteAmountOk returns a tuple with the MaxSizeQuoteAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSizeQuoteAmount

`func (o *PairInfo) SetMaxSizeQuoteAmount(v float32)`

SetMaxSizeQuoteAmount sets MaxSizeQuoteAmount field to given value.


### GetRfqDailyLimitQuoteAmount

`func (o *PairInfo) GetRfqDailyLimitQuoteAmount() float32`

GetRfqDailyLimitQuoteAmount returns the RfqDailyLimitQuoteAmount field if non-nil, zero value otherwise.

### GetRfqDailyLimitQuoteAmountOk

`func (o *PairInfo) GetRfqDailyLimitQuoteAmountOk() (*float32, bool)`

GetRfqDailyLimitQuoteAmountOk returns a tuple with the RfqDailyLimitQuoteAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfqDailyLimitQuoteAmount

`func (o *PairInfo) SetRfqDailyLimitQuoteAmount(v float32)`

SetRfqDailyLimitQuoteAmount sets RfqDailyLimitQuoteAmount field to given value.


### GetIsActive

`func (o *PairInfo) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PairInfo) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PairInfo) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetQuoteLifetimeSeconds

`func (o *PairInfo) GetQuoteLifetimeSeconds() int32`

GetQuoteLifetimeSeconds returns the QuoteLifetimeSeconds field if non-nil, zero value otherwise.

### GetQuoteLifetimeSecondsOk

`func (o *PairInfo) GetQuoteLifetimeSecondsOk() (*int32, bool)`

GetQuoteLifetimeSecondsOk returns a tuple with the QuoteLifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteLifetimeSeconds

`func (o *PairInfo) SetQuoteLifetimeSeconds(v int32)`

SetQuoteLifetimeSeconds sets QuoteLifetimeSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


