# PairBaseClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCurrency** | **string** |  | 
**QuoteCurrency** | **string** |  | 

## Methods

### NewPairBaseClient

`func NewPairBaseClient(baseCurrency string, quoteCurrency string, ) *PairBaseClient`

NewPairBaseClient instantiates a new PairBaseClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPairBaseClientWithDefaults

`func NewPairBaseClientWithDefaults() *PairBaseClient`

NewPairBaseClientWithDefaults instantiates a new PairBaseClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseCurrency

`func (o *PairBaseClient) GetBaseCurrency() string`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *PairBaseClient) GetBaseCurrencyOk() (*string, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *PairBaseClient) SetBaseCurrency(v string)`

SetBaseCurrency sets BaseCurrency field to given value.


### GetQuoteCurrency

`func (o *PairBaseClient) GetQuoteCurrency() string`

GetQuoteCurrency returns the QuoteCurrency field if non-nil, zero value otherwise.

### GetQuoteCurrencyOk

`func (o *PairBaseClient) GetQuoteCurrencyOk() (*string, bool)`

GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCurrency

`func (o *PairBaseClient) SetQuoteCurrency(v string)`

SetQuoteCurrency sets QuoteCurrency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


