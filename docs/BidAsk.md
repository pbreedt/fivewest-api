# BidAsk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCurrency** | **string** |  | 
**QuoteCurrency** | **string** |  | 
**Bid** | **float32** |  | 
**Ask** | **float32** |  | 

## Methods

### NewBidAsk

`func NewBidAsk(baseCurrency string, quoteCurrency string, bid float32, ask float32, ) *BidAsk`

NewBidAsk instantiates a new BidAsk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBidAskWithDefaults

`func NewBidAskWithDefaults() *BidAsk`

NewBidAskWithDefaults instantiates a new BidAsk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseCurrency

`func (o *BidAsk) GetBaseCurrency() string`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *BidAsk) GetBaseCurrencyOk() (*string, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *BidAsk) SetBaseCurrency(v string)`

SetBaseCurrency sets BaseCurrency field to given value.


### GetQuoteCurrency

`func (o *BidAsk) GetQuoteCurrency() string`

GetQuoteCurrency returns the QuoteCurrency field if non-nil, zero value otherwise.

### GetQuoteCurrencyOk

`func (o *BidAsk) GetQuoteCurrencyOk() (*string, bool)`

GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCurrency

`func (o *BidAsk) SetQuoteCurrency(v string)`

SetQuoteCurrency sets QuoteCurrency field to given value.


### GetBid

`func (o *BidAsk) GetBid() float32`

GetBid returns the Bid field if non-nil, zero value otherwise.

### GetBidOk

`func (o *BidAsk) GetBidOk() (*float32, bool)`

GetBidOk returns a tuple with the Bid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBid

`func (o *BidAsk) SetBid(v float32)`

SetBid sets Bid field to given value.


### GetAsk

`func (o *BidAsk) GetAsk() float32`

GetAsk returns the Ask field if non-nil, zero value otherwise.

### GetAskOk

`func (o *BidAsk) GetAskOk() (*float32, bool)`

GetAskOk returns a tuple with the Ask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsk

`func (o *BidAsk) SetAsk(v float32)`

SetAsk sets Ask field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


