# SyntheticsApiTestResultData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cert** | Pointer to [**SyntheticsSslCertificate**](SyntheticsSSLCertificate.md) |  | [optional] 
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

### NewSyntheticsApiTestResultData

`func NewSyntheticsApiTestResultData() *SyntheticsApiTestResultData`

NewSyntheticsApiTestResultData instantiates a new SyntheticsApiTestResultData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsApiTestResultDataWithDefaults

`func NewSyntheticsApiTestResultDataWithDefaults() *SyntheticsApiTestResultData`

NewSyntheticsApiTestResultDataWithDefaults instantiates a new SyntheticsApiTestResultData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCert

`func (o *SyntheticsApiTestResultData) GetCert() SyntheticsSslCertificate`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *SyntheticsApiTestResultData) GetCertOk() (SyntheticsSslCertificate, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCert

`func (o *SyntheticsApiTestResultData) HasCert() bool`

HasCert returns a boolean if a field has been set.

### SetCert

`func (o *SyntheticsApiTestResultData) SetCert(v SyntheticsSslCertificate)`

SetCert gets a reference to the given SyntheticsSslCertificate and assigns it to the Cert field.

### GetErrorCode

`func (o *SyntheticsApiTestResultData) GetErrorCode() SyntheticsErrorCode`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *SyntheticsApiTestResultData) GetErrorCodeOk() (SyntheticsErrorCode, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrorCode

`func (o *SyntheticsApiTestResultData) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCode

`func (o *SyntheticsApiTestResultData) SetErrorCode(v SyntheticsErrorCode)`

SetErrorCode gets a reference to the given SyntheticsErrorCode and assigns it to the ErrorCode field.

### GetErrorMessage

`func (o *SyntheticsApiTestResultData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SyntheticsApiTestResultData) GetErrorMessageOk() (string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrorMessage

`func (o *SyntheticsApiTestResultData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessage

`func (o *SyntheticsApiTestResultData) SetErrorMessage(v string)`

SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.

### GetEventType

`func (o *SyntheticsApiTestResultData) GetEventType() SyntheticsTestProcessStatus`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SyntheticsApiTestResultData) GetEventTypeOk() (SyntheticsTestProcessStatus, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventType

`func (o *SyntheticsApiTestResultData) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### SetEventType

`func (o *SyntheticsApiTestResultData) SetEventType(v SyntheticsTestProcessStatus)`

SetEventType gets a reference to the given SyntheticsTestProcessStatus and assigns it to the EventType field.

### GetHttpStatusCode

`func (o *SyntheticsApiTestResultData) GetHttpStatusCode() int64`

GetHttpStatusCode returns the HttpStatusCode field if non-nil, zero value otherwise.

### GetHttpStatusCodeOk

`func (o *SyntheticsApiTestResultData) GetHttpStatusCodeOk() (int64, bool)`

GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHttpStatusCode

`func (o *SyntheticsApiTestResultData) HasHttpStatusCode() bool`

HasHttpStatusCode returns a boolean if a field has been set.

### SetHttpStatusCode

`func (o *SyntheticsApiTestResultData) SetHttpStatusCode(v int64)`

SetHttpStatusCode gets a reference to the given int64 and assigns it to the HttpStatusCode field.

### GetRequestHeaders

`func (o *SyntheticsApiTestResultData) GetRequestHeaders() map[string]interface{}`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *SyntheticsApiTestResultData) GetRequestHeadersOk() (map[string]interface{}, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequestHeaders

`func (o *SyntheticsApiTestResultData) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### SetRequestHeaders

`func (o *SyntheticsApiTestResultData) SetRequestHeaders(v map[string]interface{})`

SetRequestHeaders gets a reference to the given map[string]interface{} and assigns it to the RequestHeaders field.

### GetResponseBody

`func (o *SyntheticsApiTestResultData) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *SyntheticsApiTestResultData) GetResponseBodyOk() (string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseBody

`func (o *SyntheticsApiTestResultData) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### SetResponseBody

`func (o *SyntheticsApiTestResultData) SetResponseBody(v string)`

SetResponseBody gets a reference to the given string and assigns it to the ResponseBody field.

### GetResponseHeaders

`func (o *SyntheticsApiTestResultData) GetResponseHeaders() map[string]interface{}`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *SyntheticsApiTestResultData) GetResponseHeadersOk() (map[string]interface{}, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseHeaders

`func (o *SyntheticsApiTestResultData) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### SetResponseHeaders

`func (o *SyntheticsApiTestResultData) SetResponseHeaders(v map[string]interface{})`

SetResponseHeaders gets a reference to the given map[string]interface{} and assigns it to the ResponseHeaders field.

### GetResponseSize

`func (o *SyntheticsApiTestResultData) GetResponseSize() int64`

GetResponseSize returns the ResponseSize field if non-nil, zero value otherwise.

### GetResponseSizeOk

`func (o *SyntheticsApiTestResultData) GetResponseSizeOk() (int64, bool)`

GetResponseSizeOk returns a tuple with the ResponseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseSize

`func (o *SyntheticsApiTestResultData) HasResponseSize() bool`

HasResponseSize returns a boolean if a field has been set.

### SetResponseSize

`func (o *SyntheticsApiTestResultData) SetResponseSize(v int64)`

SetResponseSize gets a reference to the given int64 and assigns it to the ResponseSize field.

### GetTimings

`func (o *SyntheticsApiTestResultData) GetTimings() SyntheticsTiming`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *SyntheticsApiTestResultData) GetTimingsOk() (SyntheticsTiming, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimings

`func (o *SyntheticsApiTestResultData) HasTimings() bool`

HasTimings returns a boolean if a field has been set.

### SetTimings

`func (o *SyntheticsApiTestResultData) SetTimings(v SyntheticsTiming)`

SetTimings gets a reference to the given SyntheticsTiming and assigns it to the Timings field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


