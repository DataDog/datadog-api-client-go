# HTTPLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ddsource** | Pointer to **string** | The integration name associated with your log: the technology from which the log originated. When it matches an integration name, Datadog automatically installs the corresponding parsers and facets. See [reserved attribute](https://docs.datadoghq.com/logs/log_collection/#reserved-attributes). | [optional] 
**Ddtags** | Pointer to **string** | Tags associated with your logs. | [optional] 
**Hostname** | Pointer to **string** | The name of the originating host of the log. | [optional] 
**Message** | Pointer to **string** | The message [reserved attribute](https://docs.datadoghq.com/logs/log_collection/#reserved-attributes) of your log. By default, Datadog ingests the value of the message attribute as the body of the log entry. That value is then highlighted and displayed in the Logstream, where it is indexed for full text search. | [optional] 

## Methods

### NewHTTPLog

`func NewHTTPLog() *HTTPLog`

NewHTTPLog instantiates a new HTTPLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPLogWithDefaults

`func NewHTTPLogWithDefaults() *HTTPLog`

NewHTTPLogWithDefaults instantiates a new HTTPLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDdsource

`func (o *HTTPLog) GetDdsource() string`

GetDdsource returns the Ddsource field if non-nil, zero value otherwise.

### GetDdsourceOk

`func (o *HTTPLog) GetDdsourceOk() (string, bool)`

GetDdsourceOk returns a tuple with the Ddsource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDdsource

`func (o *HTTPLog) HasDdsource() bool`

HasDdsource returns a boolean if a field has been set.

### SetDdsource

`func (o *HTTPLog) SetDdsource(v string)`

SetDdsource gets a reference to the given string and assigns it to the Ddsource field.

### GetDdtags

`func (o *HTTPLog) GetDdtags() string`

GetDdtags returns the Ddtags field if non-nil, zero value otherwise.

### GetDdtagsOk

`func (o *HTTPLog) GetDdtagsOk() (string, bool)`

GetDdtagsOk returns a tuple with the Ddtags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDdtags

`func (o *HTTPLog) HasDdtags() bool`

HasDdtags returns a boolean if a field has been set.

### SetDdtags

`func (o *HTTPLog) SetDdtags(v string)`

SetDdtags gets a reference to the given string and assigns it to the Ddtags field.

### GetHostname

`func (o *HTTPLog) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *HTTPLog) GetHostnameOk() (string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHostname

`func (o *HTTPLog) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostname

`func (o *HTTPLog) SetHostname(v string)`

SetHostname gets a reference to the given string and assigns it to the Hostname field.

### GetMessage

`func (o *HTTPLog) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HTTPLog) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *HTTPLog) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *HTTPLog) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


