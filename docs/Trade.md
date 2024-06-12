# Trade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** |  | [optional] 
**ExchangeAccountId** | Pointer to **int32** |  | [optional] 
**ExchangeTradeId** | Pointer to **string** |  | [optional] 
**BaseCurrency** | **string** |  | 
**QuoteCurrency** | **string** |  | 
**Rate** | **float32** |  | 
**BaseAmount** | **float32** |  | 
**QuoteAmount** | **float32** |  | 
**Side** | [**TradeSide**](TradeSide.md) |  | 
**Mode** | [**Mode**](Mode.md) |  | 
**Type** | Pointer to [**TradeType**](TradeType.md) |  | [optional] 
**ConfirmedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewTrade

`func NewTrade(baseCurrency string, quoteCurrency string, rate float32, baseAmount float32, quoteAmount float32, side TradeSide, mode Mode, id string, createdAt time.Time, ) *Trade`

NewTrade instantiates a new Trade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradeWithDefaults

`func NewTradeWithDefaults() *Trade`

NewTradeWithDefaults instantiates a new Trade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *Trade) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Trade) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Trade) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Trade) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetExchangeAccountId

`func (o *Trade) GetExchangeAccountId() int32`

GetExchangeAccountId returns the ExchangeAccountId field if non-nil, zero value otherwise.

### GetExchangeAccountIdOk

`func (o *Trade) GetExchangeAccountIdOk() (*int32, bool)`

GetExchangeAccountIdOk returns a tuple with the ExchangeAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeAccountId

`func (o *Trade) SetExchangeAccountId(v int32)`

SetExchangeAccountId sets ExchangeAccountId field to given value.

### HasExchangeAccountId

`func (o *Trade) HasExchangeAccountId() bool`

HasExchangeAccountId returns a boolean if a field has been set.

### GetExchangeTradeId

`func (o *Trade) GetExchangeTradeId() string`

GetExchangeTradeId returns the ExchangeTradeId field if non-nil, zero value otherwise.

### GetExchangeTradeIdOk

`func (o *Trade) GetExchangeTradeIdOk() (*string, bool)`

GetExchangeTradeIdOk returns a tuple with the ExchangeTradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTradeId

`func (o *Trade) SetExchangeTradeId(v string)`

SetExchangeTradeId sets ExchangeTradeId field to given value.

### HasExchangeTradeId

`func (o *Trade) HasExchangeTradeId() bool`

HasExchangeTradeId returns a boolean if a field has been set.

### GetBaseCurrency

`func (o *Trade) GetBaseCurrency() string`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *Trade) GetBaseCurrencyOk() (*string, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *Trade) SetBaseCurrency(v string)`

SetBaseCurrency sets BaseCurrency field to given value.


### GetQuoteCurrency

`func (o *Trade) GetQuoteCurrency() string`

GetQuoteCurrency returns the QuoteCurrency field if non-nil, zero value otherwise.

### GetQuoteCurrencyOk

`func (o *Trade) GetQuoteCurrencyOk() (*string, bool)`

GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCurrency

`func (o *Trade) SetQuoteCurrency(v string)`

SetQuoteCurrency sets QuoteCurrency field to given value.


### GetRate

`func (o *Trade) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *Trade) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *Trade) SetRate(v float32)`

SetRate sets Rate field to given value.


### GetBaseAmount

`func (o *Trade) GetBaseAmount() float32`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *Trade) GetBaseAmountOk() (*float32, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *Trade) SetBaseAmount(v float32)`

SetBaseAmount sets BaseAmount field to given value.


### GetQuoteAmount

`func (o *Trade) GetQuoteAmount() float32`

GetQuoteAmount returns the QuoteAmount field if non-nil, zero value otherwise.

### GetQuoteAmountOk

`func (o *Trade) GetQuoteAmountOk() (*float32, bool)`

GetQuoteAmountOk returns a tuple with the QuoteAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAmount

`func (o *Trade) SetQuoteAmount(v float32)`

SetQuoteAmount sets QuoteAmount field to given value.


### GetSide

`func (o *Trade) GetSide() TradeSide`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *Trade) GetSideOk() (*TradeSide, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *Trade) SetSide(v TradeSide)`

SetSide sets Side field to given value.


### GetMode

`func (o *Trade) GetMode() Mode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Trade) GetModeOk() (*Mode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Trade) SetMode(v Mode)`

SetMode sets Mode field to given value.


### GetType

`func (o *Trade) GetType() TradeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Trade) GetTypeOk() (*TradeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Trade) SetType(v TradeType)`

SetType sets Type field to given value.

### HasType

`func (o *Trade) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfirmedAt

`func (o *Trade) GetConfirmedAt() time.Time`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *Trade) GetConfirmedAtOk() (*time.Time, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *Trade) SetConfirmedAt(v time.Time)`

SetConfirmedAt sets ConfirmedAt field to given value.

### HasConfirmedAt

`func (o *Trade) HasConfirmedAt() bool`

HasConfirmedAt returns a boolean if a field has been set.

### GetId

`func (o *Trade) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Trade) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Trade) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Trade) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Trade) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Trade) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


