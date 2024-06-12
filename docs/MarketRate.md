# MarketRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | **string** |  | 
**Quote** | **string** |  | 
**Exchange** | **string** |  | 
**Bid** | **float32** |  | 
**Ask** | **float32** |  | 
**GeneratedAt** | **time.Time** |  | 

## Methods

### NewMarketRate

`func NewMarketRate(base string, quote string, exchange string, bid float32, ask float32, generatedAt time.Time, ) *MarketRate`

NewMarketRate instantiates a new MarketRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketRateWithDefaults

`func NewMarketRateWithDefaults() *MarketRate`

NewMarketRateWithDefaults instantiates a new MarketRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *MarketRate) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *MarketRate) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *MarketRate) SetBase(v string)`

SetBase sets Base field to given value.


### GetQuote

`func (o *MarketRate) GetQuote() string`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *MarketRate) GetQuoteOk() (*string, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *MarketRate) SetQuote(v string)`

SetQuote sets Quote field to given value.


### GetExchange

`func (o *MarketRate) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *MarketRate) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *MarketRate) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetBid

`func (o *MarketRate) GetBid() float32`

GetBid returns the Bid field if non-nil, zero value otherwise.

### GetBidOk

`func (o *MarketRate) GetBidOk() (*float32, bool)`

GetBidOk returns a tuple with the Bid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBid

`func (o *MarketRate) SetBid(v float32)`

SetBid sets Bid field to given value.


### GetAsk

`func (o *MarketRate) GetAsk() float32`

GetAsk returns the Ask field if non-nil, zero value otherwise.

### GetAskOk

`func (o *MarketRate) GetAskOk() (*float32, bool)`

GetAskOk returns a tuple with the Ask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsk

`func (o *MarketRate) SetAsk(v float32)`

SetAsk sets Ask field to given value.


### GetGeneratedAt

`func (o *MarketRate) GetGeneratedAt() time.Time`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *MarketRate) GetGeneratedAtOk() (*time.Time, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *MarketRate) SetGeneratedAt(v time.Time)`

SetGeneratedAt sets GeneratedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


