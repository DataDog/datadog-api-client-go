# SyntheticsAPITestResultData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Cert** | Pointer to [**SyntheticsSSLCertificate**](SyntheticsSSLCertificate.md) |  | [optional] 
**ErrorCode** | Pointer to [**SyntheticsErrorCode**](SyntheticsErrorCode.md) |  | [optional] 
**ErrorMessage** | Pointer to **string** | The API test error message. | [optional] 
**EventType** | Pointer to [**SyntheticsTestProcessStatus**](SyntheticsTestProcessStatus.md) |  | [optional] 
**HttpStatusCode** | Pointer to **int64** | The API test HTTP status code. | [optional] 
**RequestHeaders** | Pointer to **map[string]interface{}** | Request header object used for the API test. | [optional] 
**ResponseBody** | Pointer to **string** | Response body returned for the API test. | [optional] 
**ResponseHeaders** | Pointer to **map[string]interface{}** | Response headers returned for the API test. | [optional] 
**ResponseSize** | Pointer to **int64** | Global size in byte of the API test response. | [optional] 
**Timings** | Pointer to [**SyntheticsTiming**](SyntheticsTiming.md) |  | [optional] 

## Methods

### NewSyntheticsAPITestResultData

`func NewSyntheticsAPITestResultData() *SyntheticsAPITestResultData`

NewSyntheticsAPITestResultData instantiates a new SyntheticsAPITestResultData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsAPITestResultDataWithDefaults

`func NewSyntheticsAPITestResultDataWithDefaults() *SyntheticsAPITestResultData`

NewSyntheticsAPITestResultDataWithDefaults instantiates a new SyntheticsAPITestResultData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCert

`func (o *SyntheticsAPITestResultData) GetCert() SyntheticsSSLCertificate`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *SyntheticsAPITestResultData) GetCertOk() (*SyntheticsSSLCertificate, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *SyntheticsAPITestResultData) SetCert(v SyntheticsSSLCertificate)`

SetCert sets Cert field to given value.

### HasCert

`func (o *SyntheticsAPITestResultData) HasCert() bool`

HasCert returns a boolean if a field has been set.

### GetErrorCode

`func (o *SyntheticsAPITestResultData) GetErrorCode() SyntheticsErrorCode`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *SyntheticsAPITestResultData) GetErrorCodeOk() (*SyntheticsErrorCode, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *SyntheticsAPITestResultData) SetErrorCode(v SyntheticsErrorCode)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *SyntheticsAPITestResultData) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *SyntheticsAPITestResultData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SyntheticsAPITestResultData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *SyntheticsAPITestResultData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *SyntheticsAPITestResultData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetEventType

`func (o *SyntheticsAPITestResultData) GetEventType() SyntheticsTestProcessStatus`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SyntheticsAPITestResultData) GetEventTypeOk() (*SyntheticsTestProcessStatus, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SyntheticsAPITestResultData) SetEventType(v SyntheticsTestProcessStatus)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *SyntheticsAPITestResultData) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetHttpStatusCode

`func (o *SyntheticsAPITestResultData) GetHttpStatusCode() int64`

GetHttpStatusCode returns the HttpStatusCode field if non-nil, zero value otherwise.

### GetHttpStatusCodeOk

`func (o *SyntheticsAPITestResultData) GetHttpStatusCodeOk() (*int64, bool)`

GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatusCode

`func (o *SyntheticsAPITestResultData) SetHttpStatusCode(v int64)`

SetHttpStatusCode sets HttpStatusCode field to given value.

### HasHttpStatusCode

`func (o *SyntheticsAPITestResultData) HasHttpStatusCode() bool`

HasHttpStatusCode returns a boolean if a field has been set.

### GetRequestHeaders

`func (o *SyntheticsAPITestResultData) GetRequestHeaders() map[string]interface{}`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *SyntheticsAPITestResultData) GetRequestHeadersOk() (*map[string]interface{}, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *SyntheticsAPITestResultData) SetRequestHeaders(v map[string]interface{})`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *SyntheticsAPITestResultData) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### GetResponseBody

`func (o *SyntheticsAPITestResultData) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *SyntheticsAPITestResultData) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *SyntheticsAPITestResultData) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *SyntheticsAPITestResultData) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetResponseHeaders

`func (o *SyntheticsAPITestResultData) GetResponseHeaders() map[string]interface{}`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *SyntheticsAPITestResultData) GetResponseHeadersOk() (*map[string]interface{}, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *SyntheticsAPITestResultData) SetResponseHeaders(v map[string]interface{})`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *SyntheticsAPITestResultData) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### GetResponseSize

`func (o *SyntheticsAPITestResultData) GetResponseSize() int64`

GetResponseSize returns the ResponseSize field if non-nil, zero value otherwise.

### GetResponseSizeOk

`func (o *SyntheticsAPITestResultData) GetResponseSizeOk() (*int64, bool)`

GetResponseSizeOk returns a tuple with the ResponseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSize

`func (o *SyntheticsAPITestResultData) SetResponseSize(v int64)`

SetResponseSize sets ResponseSize field to given value.

### HasResponseSize

`func (o *SyntheticsAPITestResultData) HasResponseSize() bool`

HasResponseSize returns a boolean if a field has been set.

### GetTimings

`func (o *SyntheticsAPITestResultData) GetTimings() SyntheticsTiming`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *SyntheticsAPITestResultData) GetTimingsOk() (*SyntheticsTiming, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimings

`func (o *SyntheticsAPITestResultData) SetTimings(v SyntheticsTiming)`

SetTimings sets Timings field to given value.

### HasTimings

`func (o *SyntheticsAPITestResultData) HasTimings() bool`

HasTimings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


