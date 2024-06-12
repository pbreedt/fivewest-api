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

// checks if the HTTPResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HTTPResponse{}

// HTTPResponse struct for HTTPResponse
type HTTPResponse struct {
	Detail string `json:"detail"`
	StatusCode int32 `json:"status_code"`
	Headers map[string]interface{} `json:"headers"`
}

type _HTTPResponse HTTPResponse

// NewHTTPResponse instantiates a new HTTPResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHTTPResponse(detail string, statusCode int32, headers map[string]interface{}) *HTTPResponse {
	this := HTTPResponse{}
	this.Detail = detail
	this.StatusCode = statusCode
	this.Headers = headers
	return &this
}

// NewHTTPResponseWithDefaults instantiates a new HTTPResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHTTPResponseWithDefaults() *HTTPResponse {
	this := HTTPResponse{}
	return &this
}

// GetDetail returns the Detail field value
func (o *HTTPResponse) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *HTTPResponse) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *HTTPResponse) SetDetail(v string) {
	o.Detail = v
}

// GetStatusCode returns the StatusCode field value
func (o *HTTPResponse) GetStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *HTTPResponse) GetStatusCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *HTTPResponse) SetStatusCode(v int32) {
	o.StatusCode = v
}

// GetHeaders returns the Headers field value
func (o *HTTPResponse) GetHeaders() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *HTTPResponse) GetHeadersOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Headers, true
}

// SetHeaders sets field value
func (o *HTTPResponse) SetHeaders(v map[string]interface{}) {
	o.Headers = v
}

func (o HTTPResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HTTPResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["detail"] = o.Detail
	toSerialize["status_code"] = o.StatusCode
	toSerialize["headers"] = o.Headers
	return toSerialize, nil
}

func (o *HTTPResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"detail",
		"status_code",
		"headers",
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

	varHTTPResponse := _HTTPResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHTTPResponse)

	if err != nil {
		return err
	}

	*o = HTTPResponse(varHTTPResponse)

	return err
}

type NullableHTTPResponse struct {
	value *HTTPResponse
	isSet bool
}

func (v NullableHTTPResponse) Get() *HTTPResponse {
	return v.value
}

func (v *NullableHTTPResponse) Set(val *HTTPResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHTTPResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHTTPResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHTTPResponse(val *HTTPResponse) *NullableHTTPResponse {
	return &NullableHTTPResponse{value: val, isSet: true}
}

func (v NullableHTTPResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHTTPResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


