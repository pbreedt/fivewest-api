# Rfq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** |  | 
**Rate** | **float32** |  | 
**TradeId** | Pointer to **string** |  | [optional] 
**PairId** | Pointer to **string** |  | [optional] 
**Pair** | [**PairBaseClient**](PairBaseClient.md) |  | 
**BaseAmount** | **float32** |  | 
**QuoteAmount** | **float32** |  | 
**Side** | [**TradeSide**](TradeSide.md) |  | 
**Status** | [**QuoteStatus**](QuoteStatus.md) |  | 
**QuoteLifetimeSeconds** | **int32** |  | 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewRfq

`func NewRfq(uid string, rate float32, pair PairBaseClient, baseAmount float32, quoteAmount float32, side TradeSide, status QuoteStatus, quoteLifetimeSeconds int32, id string, createdAt time.Time, ) *Rfq`

NewRfq instantiates a new Rfq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRfqWithDefaults

`func NewRfqWithDefaults() *Rfq`

NewRfqWithDefaults instantiates a new Rfq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *Rfq) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Rfq) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Rfq) SetUid(v string)`

SetUid sets Uid field to given value.


### GetRate

`func (o *Rfq) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *Rfq) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *Rfq) SetRate(v float32)`

SetRate sets Rate field to given value.


### GetTradeId

`func (o *Rfq) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *Rfq) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *Rfq) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *Rfq) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.

### GetPairId

`func (o *Rfq) GetPairId() string`

GetPairId returns the PairId field if non-nil, zero value otherwise.

### GetPairIdOk

`func (o *Rfq) GetPairIdOk() (*string, bool)`

GetPairIdOk returns a tuple with the PairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairId

`func (o *Rfq) SetPairId(v string)`

SetPairId sets PairId field to given value.

### HasPairId

`func (o *Rfq) HasPairId() bool`

HasPairId returns a boolean if a field has been set.

### GetPair

`func (o *Rfq) GetPair() PairBaseClient`

GetPair returns the Pair field if non-nil, zero value otherwise.

### GetPairOk

`func (o *Rfq) GetPairOk() (*PairBaseClient, bool)`

GetPairOk returns a tuple with the Pair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPair

`func (o *Rfq) SetPair(v PairBaseClient)`

SetPair sets Pair field to given value.


### GetBaseAmount

`func (o *Rfq) GetBaseAmount() float32`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *Rfq) GetBaseAmountOk() (*float32, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *Rfq) SetBaseAmount(v float32)`

SetBaseAmount sets BaseAmount field to given value.


### GetQuoteAmount

`func (o *Rfq) GetQuoteAmount() float32`

GetQuoteAmount returns the QuoteAmount field if non-nil, zero value otherwise.

### GetQuoteAmountOk

`func (o *Rfq) GetQuoteAmountOk() (*float32, bool)`

GetQuoteAmountOk returns a tuple with the QuoteAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAmount

`func (o *Rfq) SetQuoteAmount(v float32)`

SetQuoteAmount sets QuoteAmount field to given value.


### GetSide

`func (o *Rfq) GetSide() TradeSide`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *Rfq) GetSideOk() (*TradeSide, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *Rfq) SetSide(v TradeSide)`

SetSide sets Side field to given value.


### GetStatus

`func (o *Rfq) GetStatus() QuoteStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Rfq) GetStatusOk() (*QuoteStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Rfq) SetStatus(v QuoteStatus)`

SetStatus sets Status field to given value.


### GetQuoteLifetimeSeconds

`func (o *Rfq) GetQuoteLifetimeSeconds() int32`

GetQuoteLifetimeSeconds returns the QuoteLifetimeSeconds field if non-nil, zero value otherwise.

### GetQuoteLifetimeSecondsOk

`func (o *Rfq) GetQuoteLifetimeSecondsOk() (*int32, bool)`

GetQuoteLifetimeSecondsOk returns a tuple with the QuoteLifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteLifetimeSeconds

`func (o *Rfq) SetQuoteLifetimeSeconds(v int32)`

SetQuoteLifetimeSeconds sets QuoteLifetimeSeconds field to given value.


### GetId

`func (o *Rfq) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Rfq) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Rfq) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Rfq) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Rfq) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Rfq) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


