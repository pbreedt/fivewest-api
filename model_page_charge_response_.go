/*
FiveWest Client API

API Documentation for FiveWest's wallet and RfQ trading services. Please visit dashboard.fivewest.co.za to sign up. Once your account has been verified to the sufficient level necessary for the given service, you can create an API key and secret with fine-grained service permissions in the 'API Management' section of the 'Profile' tab. These credentials must be provided in the /profile/api/v1/auth/token route in order to acquire a JWT to make further requests with. This JWT is valid for one week; the credentials do not expire but may be deleted.

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fivewestapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PageChargeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageChargeResponse{}

// PageChargeResponse struct for PageChargeResponse
type PageChargeResponse struct {
	Items []ChargeResponse `json:"items"`
	Total int32 `json:"total"`
	Pages int32 `json:"pages"`
	PageSize int32 `json:"page_size"`
	HasNext *bool `json:"has_next,omitempty"`
	HasPrev *bool `json:"has_prev,omitempty"`
}

type _PageChargeResponse PageChargeResponse

// NewPageChargeResponse instantiates a new PageChargeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageChargeResponse(items []ChargeResponse, total int32, pages int32, pageSize int32) *PageChargeResponse {
	this := PageChargeResponse{}
	this.Items = items
	this.Total = total
	this.Pages = pages
	this.PageSize = pageSize
	var hasNext bool = false
	this.HasNext = &hasNext
	var hasPrev bool = false
	this.HasPrev = &hasPrev
	return &this
}

// NewPageChargeResponseWithDefaults instantiates a new PageChargeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageChargeResponseWithDefaults() *PageChargeResponse {
	this := PageChargeResponse{}
	var hasNext bool = false
	this.HasNext = &hasNext
	var hasPrev bool = false
	this.HasPrev = &hasPrev
	return &this
}

// GetItems returns the Items field value
func (o *PageChargeResponse) GetItems() []ChargeResponse {
	if o == nil {
		var ret []ChargeResponse
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *PageChargeResponse) GetItemsOk() ([]ChargeResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *PageChargeResponse) SetItems(v []ChargeResponse) {
	o.Items = v
}

// GetTotal returns the Total field value
func (o *PageChargeResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *PageChargeResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *PageChargeResponse) SetTotal(v int32) {
	o.Total = v
}

// GetPages returns the Pages field value
func (o *PageChargeResponse) GetPages() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pages
}

// GetPagesOk returns a tuple with the Pages field value
// and a boolean to check if the value has been set.
func (o *PageChargeResponse) GetPagesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pages, true
}

// SetPages sets field value
func (o *PageChargeResponse) SetPages(v int32) {
	o.Pages = v
}

// GetPageSize returns the PageSize field value
func (o *PageChargeResponse) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *PageChargeResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *PageChargeResponse) SetPageSize(v int32) {
	o.PageSize = v
}

// GetHasNext returns the HasNext field value if set, zero value otherwise.
func (o *PageChargeResponse) GetHasNext() bool {
	if o == nil || IsNil(o.HasNext) {
		var ret bool
		return ret
	}
	return *o.HasNext
}

// GetHasNextOk returns a tuple with the HasNext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageChargeResponse) GetHasNextOk() (*bool, bool) {
	if o == nil || IsNil(o.HasNext) {
		return nil, false
	}
	return o.HasNext, true
}

// HasHasNext returns a boolean if a field has been set.
func (o *PageChargeResponse) HasHasNext() bool {
	if o != nil && !IsNil(o.HasNext) {
		return true
	}

	return false
}

// SetHasNext gets a reference to the given bool and assigns it to the HasNext field.
func (o *PageChargeResponse) SetHasNext(v bool) {
	o.HasNext = &v
}

// GetHasPrev returns the HasPrev field value if set, zero value otherwise.
func (o *PageChargeResponse) GetHasPrev() bool {
	if o == nil || IsNil(o.HasPrev) {
		var ret bool
		return ret
	}
	return *o.HasPrev
}

// GetHasPrevOk returns a tuple with the HasPrev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageChargeResponse) GetHasPrevOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPrev) {
		return nil, false
	}
	return o.HasPrev, true
}

// HasHasPrev returns a boolean if a field has been set.
func (o *PageChargeResponse) HasHasPrev() bool {
	if o != nil && !IsNil(o.HasPrev) {
		return true
	}

	return false
}

// SetHasPrev gets a reference to the given bool and assigns it to the HasPrev field.
func (o *PageChargeResponse) SetHasPrev(v bool) {
	o.HasPrev = &v
}

func (o PageChargeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageChargeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["total"] = o.Total
	toSerialize["pages"] = o.Pages
	toSerialize["page_size"] = o.PageSize
	if !IsNil(o.HasNext) {
		toSerialize["has_next"] = o.HasNext
	}
	if !IsNil(o.HasPrev) {
		toSerialize["has_prev"] = o.HasPrev
	}
	return toSerialize, nil
}

func (o *PageChargeResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
		"total",
		"pages",
		"page_size",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPageChargeResponse := _PageChargeResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPageChargeResponse)

	if err != nil {
		return err
	}

	*o = PageChargeResponse(varPageChargeResponse)

	return err
}

type NullablePageChargeResponse struct {
	value *PageChargeResponse
	isSet bool
}

func (v NullablePageChargeResponse) Get() *PageChargeResponse {
	return v.value
}

func (v *NullablePageChargeResponse) Set(val *PageChargeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePageChargeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePageChargeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageChargeResponse(val *PageChargeResponse) *NullablePageChargeResponse {
	return &NullablePageChargeResponse{value: val, isSet: true}
}

func (v NullablePageChargeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageChargeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


