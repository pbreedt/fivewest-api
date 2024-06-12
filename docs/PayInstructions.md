# PayInstructions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to [**NullableDepositAddress**](DepositAddress.md) |  | [optional] 
**BpayQrContent** | Pointer to **NullableString** |  | [optional] 
**BpayCheckoutUrl** | Pointer to **NullableString** |  | [optional] 
**BpayDeeplink** | Pointer to **NullableString** |  | [optional] 
**BpayUniversalUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPayInstructions

`func NewPayInstructions() *PayInstructions`

NewPayInstructions instantiates a new PayInstructions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayInstructionsWithDefaults

`func NewPayInstructionsWithDefaults() *PayInstructions`

NewPayInstructionsWithDefaults instantiates a new PayInstructions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *PayInstructions) GetDestination() DepositAddress`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PayInstructions) GetDestinationOk() (*DepositAddress, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PayInstructions) SetDestination(v DepositAddress)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *PayInstructions) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### SetDestinationNil

`func (o *PayInstructions) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *PayInstructions) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetBpayQrContent

`func (o *PayInstructions) GetBpayQrContent() string`

GetBpayQrContent returns the BpayQrContent field if non-nil, zero value otherwise.

### GetBpayQrContentOk

`func (o *PayInstructions) GetBpayQrContentOk() (*string, bool)`

GetBpayQrContentOk returns a tuple with the BpayQrContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpayQrContent

`func (o *PayInstructions) SetBpayQrContent(v string)`

SetBpayQrContent sets BpayQrContent field to given value.

### HasBpayQrContent

`func (o *PayInstructions) HasBpayQrContent() bool`

HasBpayQrContent returns a boolean if a field has been set.

### SetBpayQrContentNil

`func (o *PayInstructions) SetBpayQrContentNil(b bool)`

 SetBpayQrContentNil sets the value for BpayQrContent to be an explicit nil

### UnsetBpayQrContent
`func (o *PayInstructions) UnsetBpayQrContent()`

UnsetBpayQrContent ensures that no value is present for BpayQrContent, not even an explicit nil
### GetBpayCheckoutUrl

`func (o *PayInstructions) GetBpayCheckoutUrl() string`

GetBpayCheckoutUrl returns the BpayCheckoutUrl field if non-nil, zero value otherwise.

### GetBpayCheckoutUrlOk

`func (o *PayInstructions) GetBpayCheckoutUrlOk() (*string, bool)`

GetBpayCheckoutUrlOk returns a tuple with the BpayCheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpayCheckoutUrl

`func (o *PayInstructions) SetBpayCheckoutUrl(v string)`

SetBpayCheckoutUrl sets BpayCheckoutUrl field to given value.

### HasBpayCheckoutUrl

`func (o *PayInstructions) HasBpayCheckoutUrl() bool`

HasBpayCheckoutUrl returns a boolean if a field has been set.

### SetBpayCheckoutUrlNil

`func (o *PayInstructions) SetBpayCheckoutUrlNil(b bool)`

 SetBpayCheckoutUrlNil sets the value for BpayCheckoutUrl to be an explicit nil

### UnsetBpayCheckoutUrl
`func (o *PayInstructions) UnsetBpayCheckoutUrl()`

UnsetBpayCheckoutUrl ensures that no value is present for BpayCheckoutUrl, not even an explicit nil
### GetBpayDeeplink

`func (o *PayInstructions) GetBpayDeeplink() string`

GetBpayDeeplink returns the BpayDeeplink field if non-nil, zero value otherwise.

### GetBpayDeeplinkOk

`func (o *PayInstructions) GetBpayDeeplinkOk() (*string, bool)`

GetBpayDeeplinkOk returns a tuple with the BpayDeeplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpayDeeplink

`func (o *PayInstructions) SetBpayDeeplink(v string)`

SetBpayDeeplink sets BpayDeeplink field to given value.

### HasBpayDeeplink

`func (o *PayInstructions) HasBpayDeeplink() bool`

HasBpayDeeplink returns a boolean if a field has been set.

### SetBpayDeeplinkNil

`func (o *PayInstructions) SetBpayDeeplinkNil(b bool)`

 SetBpayDeeplinkNil sets the value for BpayDeeplink to be an explicit nil

### UnsetBpayDeeplink
`func (o *PayInstructions) UnsetBpayDeeplink()`

UnsetBpayDeeplink ensures that no value is present for BpayDeeplink, not even an explicit nil
### GetBpayUniversalUrl

`func (o *PayInstructions) GetBpayUniversalUrl() string`

GetBpayUniversalUrl returns the BpayUniversalUrl field if non-nil, zero value otherwise.

### GetBpayUniversalUrlOk

`func (o *PayInstructions) GetBpayUniversalUrlOk() (*string, bool)`

GetBpayUniversalUrlOk returns a tuple with the BpayUniversalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpayUniversalUrl

`func (o *PayInstructions) SetBpayUniversalUrl(v string)`

SetBpayUniversalUrl sets BpayUniversalUrl field to given value.

### HasBpayUniversalUrl

`func (o *PayInstructions) HasBpayUniversalUrl() bool`

HasBpayUniversalUrl returns a boolean if a field has been set.

### SetBpayUniversalUrlNil

`func (o *PayInstructions) SetBpayUniversalUrlNil(b bool)`

 SetBpayUniversalUrlNil sets the value for BpayUniversalUrl to be an explicit nil

### UnsetBpayUniversalUrl
`func (o *PayInstructions) UnsetBpayUniversalUrl()`

UnsetBpayUniversalUrl ensures that no value is present for BpayUniversalUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


