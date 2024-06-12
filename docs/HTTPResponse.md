# HTTPResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | **string** |  | 
**StatusCode** | **int32** |  | 
**Headers** | **map[string]interface{}** |  | 

## Methods

### NewHTTPResponse

`func NewHTTPResponse(detail string, statusCode int32, headers map[string]interface{}, ) *HTTPResponse`

NewHTTPResponse instantiates a new HTTPResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPResponseWithDefaults

`func NewHTTPResponseWithDefaults() *HTTPResponse`

NewHTTPResponseWithDefaults instantiates a new HTTPResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *HTTPResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *HTTPResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *HTTPResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetStatusCode

`func (o *HTTPResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *HTTPResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *HTTPResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetHeaders

`func (o *HTTPResponse) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HTTPResponse) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HTTPResponse) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


