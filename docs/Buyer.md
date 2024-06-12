# Buyer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceBuyerId** | Pointer to **NullableString** |  | [optional] 
**BuyerName** | Pointer to [**NullableName**](Name.md) |  | [optional] 
**BuyerPhoneCountryCode** | Pointer to **NullableString** |  | [optional] 
**BuyerPhoneNo** | Pointer to **NullableString** |  | [optional] 
**BuyerEmail** | Pointer to **NullableString** |  | [optional] 
**BuyerRegistrationTime** | Pointer to **NullableTime** |  | [optional] 
**BuyerBrowserLanguage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBuyer

`func NewBuyer() *Buyer`

NewBuyer instantiates a new Buyer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyerWithDefaults

`func NewBuyerWithDefaults() *Buyer`

NewBuyerWithDefaults instantiates a new Buyer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceBuyerId

`func (o *Buyer) GetReferenceBuyerId() string`

GetReferenceBuyerId returns the ReferenceBuyerId field if non-nil, zero value otherwise.

### GetReferenceBuyerIdOk

`func (o *Buyer) GetReferenceBuyerIdOk() (*string, bool)`

GetReferenceBuyerIdOk returns a tuple with the ReferenceBuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceBuyerId

`func (o *Buyer) SetReferenceBuyerId(v string)`

SetReferenceBuyerId sets ReferenceBuyerId field to given value.

### HasReferenceBuyerId

`func (o *Buyer) HasReferenceBuyerId() bool`

HasReferenceBuyerId returns a boolean if a field has been set.

### SetReferenceBuyerIdNil

`func (o *Buyer) SetReferenceBuyerIdNil(b bool)`

 SetReferenceBuyerIdNil sets the value for ReferenceBuyerId to be an explicit nil

### UnsetReferenceBuyerId
`func (o *Buyer) UnsetReferenceBuyerId()`

UnsetReferenceBuyerId ensures that no value is present for ReferenceBuyerId, not even an explicit nil
### GetBuyerName

`func (o *Buyer) GetBuyerName() Name`

GetBuyerName returns the BuyerName field if non-nil, zero value otherwise.

### GetBuyerNameOk

`func (o *Buyer) GetBuyerNameOk() (*Name, bool)`

GetBuyerNameOk returns a tuple with the BuyerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerName

`func (o *Buyer) SetBuyerName(v Name)`

SetBuyerName sets BuyerName field to given value.

### HasBuyerName

`func (o *Buyer) HasBuyerName() bool`

HasBuyerName returns a boolean if a field has been set.

### SetBuyerNameNil

`func (o *Buyer) SetBuyerNameNil(b bool)`

 SetBuyerNameNil sets the value for BuyerName to be an explicit nil

### UnsetBuyerName
`func (o *Buyer) UnsetBuyerName()`

UnsetBuyerName ensures that no value is present for BuyerName, not even an explicit nil
### GetBuyerPhoneCountryCode

`func (o *Buyer) GetBuyerPhoneCountryCode() string`

GetBuyerPhoneCountryCode returns the BuyerPhoneCountryCode field if non-nil, zero value otherwise.

### GetBuyerPhoneCountryCodeOk

`func (o *Buyer) GetBuyerPhoneCountryCodeOk() (*string, bool)`

GetBuyerPhoneCountryCodeOk returns a tuple with the BuyerPhoneCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerPhoneCountryCode

`func (o *Buyer) SetBuyerPhoneCountryCode(v string)`

SetBuyerPhoneCountryCode sets BuyerPhoneCountryCode field to given value.

### HasBuyerPhoneCountryCode

`func (o *Buyer) HasBuyerPhoneCountryCode() bool`

HasBuyerPhoneCountryCode returns a boolean if a field has been set.

### SetBuyerPhoneCountryCodeNil

`func (o *Buyer) SetBuyerPhoneCountryCodeNil(b bool)`

 SetBuyerPhoneCountryCodeNil sets the value for BuyerPhoneCountryCode to be an explicit nil

### UnsetBuyerPhoneCountryCode
`func (o *Buyer) UnsetBuyerPhoneCountryCode()`

UnsetBuyerPhoneCountryCode ensures that no value is present for BuyerPhoneCountryCode, not even an explicit nil
### GetBuyerPhoneNo

`func (o *Buyer) GetBuyerPhoneNo() string`

GetBuyerPhoneNo returns the BuyerPhoneNo field if non-nil, zero value otherwise.

### GetBuyerPhoneNoOk

`func (o *Buyer) GetBuyerPhoneNoOk() (*string, bool)`

GetBuyerPhoneNoOk returns a tuple with the BuyerPhoneNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerPhoneNo

`func (o *Buyer) SetBuyerPhoneNo(v string)`

SetBuyerPhoneNo sets BuyerPhoneNo field to given value.

### HasBuyerPhoneNo

`func (o *Buyer) HasBuyerPhoneNo() bool`

HasBuyerPhoneNo returns a boolean if a field has been set.

### SetBuyerPhoneNoNil

`func (o *Buyer) SetBuyerPhoneNoNil(b bool)`

 SetBuyerPhoneNoNil sets the value for BuyerPhoneNo to be an explicit nil

### UnsetBuyerPhoneNo
`func (o *Buyer) UnsetBuyerPhoneNo()`

UnsetBuyerPhoneNo ensures that no value is present for BuyerPhoneNo, not even an explicit nil
### GetBuyerEmail

`func (o *Buyer) GetBuyerEmail() string`

GetBuyerEmail returns the BuyerEmail field if non-nil, zero value otherwise.

### GetBuyerEmailOk

`func (o *Buyer) GetBuyerEmailOk() (*string, bool)`

GetBuyerEmailOk returns a tuple with the BuyerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerEmail

`func (o *Buyer) SetBuyerEmail(v string)`

SetBuyerEmail sets BuyerEmail field to given value.

### HasBuyerEmail

`func (o *Buyer) HasBuyerEmail() bool`

HasBuyerEmail returns a boolean if a field has been set.

### SetBuyerEmailNil

`func (o *Buyer) SetBuyerEmailNil(b bool)`

 SetBuyerEmailNil sets the value for BuyerEmail to be an explicit nil

### UnsetBuyerEmail
`func (o *Buyer) UnsetBuyerEmail()`

UnsetBuyerEmail ensures that no value is present for BuyerEmail, not even an explicit nil
### GetBuyerRegistrationTime

`func (o *Buyer) GetBuyerRegistrationTime() time.Time`

GetBuyerRegistrationTime returns the BuyerRegistrationTime field if non-nil, zero value otherwise.

### GetBuyerRegistrationTimeOk

`func (o *Buyer) GetBuyerRegistrationTimeOk() (*time.Time, bool)`

GetBuyerRegistrationTimeOk returns a tuple with the BuyerRegistrationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerRegistrationTime

`func (o *Buyer) SetBuyerRegistrationTime(v time.Time)`

SetBuyerRegistrationTime sets BuyerRegistrationTime field to given value.

### HasBuyerRegistrationTime

`func (o *Buyer) HasBuyerRegistrationTime() bool`

HasBuyerRegistrationTime returns a boolean if a field has been set.

### SetBuyerRegistrationTimeNil

`func (o *Buyer) SetBuyerRegistrationTimeNil(b bool)`

 SetBuyerRegistrationTimeNil sets the value for BuyerRegistrationTime to be an explicit nil

### UnsetBuyerRegistrationTime
`func (o *Buyer) UnsetBuyerRegistrationTime()`

UnsetBuyerRegistrationTime ensures that no value is present for BuyerRegistrationTime, not even an explicit nil
### GetBuyerBrowserLanguage

`func (o *Buyer) GetBuyerBrowserLanguage() string`

GetBuyerBrowserLanguage returns the BuyerBrowserLanguage field if non-nil, zero value otherwise.

### GetBuyerBrowserLanguageOk

`func (o *Buyer) GetBuyerBrowserLanguageOk() (*string, bool)`

GetBuyerBrowserLanguageOk returns a tuple with the BuyerBrowserLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerBrowserLanguage

`func (o *Buyer) SetBuyerBrowserLanguage(v string)`

SetBuyerBrowserLanguage sets BuyerBrowserLanguage field to given value.

### HasBuyerBrowserLanguage

`func (o *Buyer) HasBuyerBrowserLanguage() bool`

HasBuyerBrowserLanguage returns a boolean if a field has been set.

### SetBuyerBrowserLanguageNil

`func (o *Buyer) SetBuyerBrowserLanguageNil(b bool)`

 SetBuyerBrowserLanguageNil sets the value for BuyerBrowserLanguage to be an explicit nil

### UnsetBuyerBrowserLanguage
`func (o *Buyer) UnsetBuyerBrowserLanguage()`

UnsetBuyerBrowserLanguage ensures that no value is present for BuyerBrowserLanguage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


