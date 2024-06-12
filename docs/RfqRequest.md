# RfqRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAmount** | Pointer to **float32** |  | [optional] 
**ToAmount** | Pointer to **float32** |  | [optional] 
**FromCurrency** | **string** |  | 
**ToCurrency** | **string** |  | 

## Methods

### NewRfqRequest

`func NewRfqRequest(fromCurrency string, toCurrency string, ) *RfqRequest`

NewRfqRequest instantiates a new RfqRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRfqRequestWithDefaults

`func NewRfqRequestWithDefaults() *RfqRequest`

NewRfqRequestWithDefaults instantiates a new RfqRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAmount

`func (o *RfqRequest) GetFromAmount() float32`

GetFromAmount returns the FromAmount field if non-nil, zero value otherwise.

### GetFromAmountOk

`func (o *RfqRequest) GetFromAmountOk() (*float32, bool)`

GetFromAmountOk returns a tuple with the FromAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAmount

`func (o *RfqRequest) SetFromAmount(v float32)`

SetFromAmount sets FromAmount field to given value.

### HasFromAmount

`func (o *RfqRequest) HasFromAmount() bool`

HasFromAmount returns a boolean if a field has been set.

### GetToAmount

`func (o *RfqRequest) GetToAmount() float32`

GetToAmount returns the ToAmount field if non-nil, zero value otherwise.

### GetToAmountOk

`func (o *RfqRequest) GetToAmountOk() (*float32, bool)`

GetToAmountOk returns a tuple with the ToAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAmount

`func (o *RfqRequest) SetToAmount(v float32)`

SetToAmount sets ToAmount field to given value.

### HasToAmount

`func (o *RfqRequest) HasToAmount() bool`

HasToAmount returns a boolean if a field has been set.

### GetFromCurrency

`func (o *RfqRequest) GetFromCurrency() string`

GetFromCurrency returns the FromCurrency field if non-nil, zero value otherwise.

### GetFromCurrencyOk

`func (o *RfqRequest) GetFromCurrencyOk() (*string, bool)`

GetFromCurrencyOk returns a tuple with the FromCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCurrency

`func (o *RfqRequest) SetFromCurrency(v string)`

SetFromCurrency sets FromCurrency field to given value.


### GetToCurrency

`func (o *RfqRequest) GetToCurrency() string`

GetToCurrency returns the ToCurrency field if non-nil, zero value otherwise.

### GetToCurrencyOk

`func (o *RfqRequest) GetToCurrencyOk() (*string, bool)`

GetToCurrencyOk returns a tuple with the ToCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCurrency

`func (o *RfqRequest) SetToCurrency(v string)`

SetToCurrency sets ToCurrency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


