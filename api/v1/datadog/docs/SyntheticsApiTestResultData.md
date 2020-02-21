# SyntheticsAPITestResultData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cert** | Pointer to [**SyntheticsSSLCertificate**](SyntheticsSSLCertificate.md) |  | [optional] 
**ErrorCode** | Pointer to [**SyntheticsErrorCode**](SyntheticsErrorCode.md) |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**EventType** | Pointer to [**SyntheticsTestProcessStatus**](SyntheticsTestProcessStatus.md) |  | [optional] 
**HttpStatusCode** | Pointer to **int64** |  | [optional] 
**RequestHeaders** | Pointer to [**map[string]interface{}**](interface{}.md) |  | [optional] 
**ResponseBody** | Pointer to **string** |  | [optional] 
**ResponseHeaders** | Pointer to [**map[string]interface{}**](interface{}.md) |  | [optional] 
**ResponseSize** | Pointer to **int64** |  | [optional] 
**Timings** | Pointer to [**SyntheticsTiming**](SyntheticsTiming.md) |  | [optional] 

## Methods

### NewSyntheticsAPITestResultData

`func NewSyntheticsAPITestResultData() *SyntheticsAPITestResultData`

NewSyntheticsAPITestResultData instantiates a new SyntheticsAPITestResultData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsAPITestResultDataWithDefaults

`func NewSyntheticsAPITestResultDataWithDefaults() *SyntheticsAPITestResultData`

NewSyntheticsAPITestResultDataWithDefaults instantiates a new SyntheticsAPITestResultData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCert

`func (o *SyntheticsAPITestResultData) GetCert() SyntheticsSSLCertificate`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *SyntheticsAPITestResultData) GetCertOk() (SyntheticsSSLCertificate, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCert

`func (o *SyntheticsAPITestResultData) HasCert() bool`

HasCert returns a boolean if a field has been set.

### SetCert

`func (o *SyntheticsAPITestResultData) SetCert(v SyntheticsSSLCertificate)`

SetCert gets a reference to the given SyntheticsSSLCertificate and assigns it to the Cert field.

### GetErrorCode

`func (o *SyntheticsAPITestResultData) GetErrorCode() SyntheticsErrorCode`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *SyntheticsAPITestResultData) GetErrorCodeOk() (SyntheticsErrorCode, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrorCode

`func (o *SyntheticsAPITestResultData) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCode

`func (o *SyntheticsAPITestResultData) SetErrorCode(v SyntheticsErrorCode)`

SetErrorCode gets a reference to the given SyntheticsErrorCode and assigns it to the ErrorCode field.

### GetErrorMessage

`func (o *SyntheticsAPITestResultData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SyntheticsAPITestResultData) GetErrorMessageOk() (string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrorMessage

`func (o *SyntheticsAPITestResultData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessage

`func (o *SyntheticsAPITestResultData) SetErrorMessage(v string)`

SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.

### GetEventType

`func (o *SyntheticsAPITestResultData) GetEventType() SyntheticsTestProcessStatus`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SyntheticsAPITestResultData) GetEventTypeOk() (SyntheticsTestProcessStatus, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventType

`func (o *SyntheticsAPITestResultData) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### SetEventType

`func (o *SyntheticsAPITestResultData) SetEventType(v SyntheticsTestProcessStatus)`

SetEventType gets a reference to the given SyntheticsTestProcessStatus and assigns it to the EventType field.

### GetHttpStatusCode

`func (o *SyntheticsAPITestResultData) GetHttpStatusCode() int64`

GetHttpStatusCode returns the HttpStatusCode field if non-nil, zero value otherwise.

### GetHttpStatusCodeOk

`func (o *SyntheticsAPITestResultData) GetHttpStatusCodeOk() (int64, bool)`

GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHttpStatusCode

`func (o *SyntheticsAPITestResultData) HasHttpStatusCode() bool`

HasHttpStatusCode returns a boolean if a field has been set.

### SetHttpStatusCode

`func (o *SyntheticsAPITestResultData) SetHttpStatusCode(v int64)`

SetHttpStatusCode gets a reference to the given int64 and assigns it to the HttpStatusCode field.

### GetRequestHeaders

`func (o *SyntheticsAPITestResultData) GetRequestHeaders() map[string]interface{}`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *SyntheticsAPITestResultData) GetRequestHeadersOk() (map[string]interface{}, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequestHeaders

`func (o *SyntheticsAPITestResultData) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### SetRequestHeaders

`func (o *SyntheticsAPITestResultData) SetRequestHeaders(v map[string]interface{})`

SetRequestHeaders gets a reference to the given map[string]interface{} and assigns it to the RequestHeaders field.

### GetResponseBody

`func (o *SyntheticsAPITestResultData) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *SyntheticsAPITestResultData) GetResponseBodyOk() (string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseBody

`func (o *SyntheticsAPITestResultData) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### SetResponseBody

`func (o *SyntheticsAPITestResultData) SetResponseBody(v string)`

SetResponseBody gets a reference to the given string and assigns it to the ResponseBody field.

### GetResponseHeaders

`func (o *SyntheticsAPITestResultData) GetResponseHeaders() map[string]interface{}`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *SyntheticsAPITestResultData) GetResponseHeadersOk() (map[string]interface{}, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseHeaders

`func (o *SyntheticsAPITestResultData) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### SetResponseHeaders

`func (o *SyntheticsAPITestResultData) SetResponseHeaders(v map[string]interface{})`

SetResponseHeaders gets a reference to the given map[string]interface{} and assigns it to the ResponseHeaders field.

### GetResponseSize

`func (o *SyntheticsAPITestResultData) GetResponseSize() int64`

GetResponseSize returns the ResponseSize field if non-nil, zero value otherwise.

### GetResponseSizeOk

`func (o *SyntheticsAPITestResultData) GetResponseSizeOk() (int64, bool)`

GetResponseSizeOk returns a tuple with the ResponseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseSize

`func (o *SyntheticsAPITestResultData) HasResponseSize() bool`

HasResponseSize returns a boolean if a field has been set.

### SetResponseSize

`func (o *SyntheticsAPITestResultData) SetResponseSize(v int64)`

SetResponseSize gets a reference to the given int64 and assigns it to the ResponseSize field.

### GetTimings

`func (o *SyntheticsAPITestResultData) GetTimings() SyntheticsTiming`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *SyntheticsAPITestResultData) GetTimingsOk() (SyntheticsTiming, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimings

`func (o *SyntheticsAPITestResultData) HasTimings() bool`

HasTimings returns a boolean if a field has been set.

### SetTimings

`func (o *SyntheticsAPITestResultData) SetTimings(v SyntheticsTiming)`

SetTimings gets a reference to the given SyntheticsTiming and assigns it to the Timings field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


