# HTTPLogItem

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Ddsource** | Pointer to **string** | The integration name associated with your log: the technology from which the log originated. When it matches an integration name, Datadog automatically installs the corresponding parsers and facets. See [reserved attributes](https://docs.datadoghq.com/logs/log_collection/#reserved-attributes). | [optional] 
**Ddtags** | Pointer to **string** | Tags associated with your logs. | [optional] 
**Hostname** | Pointer to **string** | The name of the originating host of the log. | [optional] 
**Message** | Pointer to **string** | The message [reserved attribute](https://docs.datadoghq.com/logs/log_collection/#reserved-attributes) of your log. By default, Datadog ingests the value of the message attribute as the body of the log entry. That value is then highlighted and displayed in the Logstream, where it is indexed for full text search. | [optional] 
**Service** | Pointer to **string** | The name of the application or service generating the log events. It is used to switch from Logs to APM, so make sure you define the same value when you use both products. See [reserved attributes](https://docs.datadoghq.com/logs/log_collection/#reserved-attributes). | [optional] 

## Methods

### NewHTTPLogItem

`func NewHTTPLogItem() *HTTPLogItem`

NewHTTPLogItem instantiates a new HTTPLogItem object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewHTTPLogItemWithDefaults

`func NewHTTPLogItemWithDefaults() *HTTPLogItem`

NewHTTPLogItemWithDefaults instantiates a new HTTPLogItem object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDdsource

`func (o *HTTPLogItem) GetDdsource() string`

GetDdsource returns the Ddsource field if non-nil, zero value otherwise.

### GetDdsourceOk

`func (o *HTTPLogItem) GetDdsourceOk() (*string, bool)`

GetDdsourceOk returns a tuple with the Ddsource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdsource

`func (o *HTTPLogItem) SetDdsource(v string)`

SetDdsource sets Ddsource field to given value.

### HasDdsource

`func (o *HTTPLogItem) HasDdsource() bool`

HasDdsource returns a boolean if a field has been set.

### GetDdtags

`func (o *HTTPLogItem) GetDdtags() string`

GetDdtags returns the Ddtags field if non-nil, zero value otherwise.

### GetDdtagsOk

`func (o *HTTPLogItem) GetDdtagsOk() (*string, bool)`

GetDdtagsOk returns a tuple with the Ddtags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdtags

`func (o *HTTPLogItem) SetDdtags(v string)`

SetDdtags sets Ddtags field to given value.

### HasDdtags

`func (o *HTTPLogItem) HasDdtags() bool`

HasDdtags returns a boolean if a field has been set.

### GetHostname

`func (o *HTTPLogItem) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *HTTPLogItem) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *HTTPLogItem) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *HTTPLogItem) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetMessage

`func (o *HTTPLogItem) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HTTPLogItem) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HTTPLogItem) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HTTPLogItem) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetService

`func (o *HTTPLogItem) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *HTTPLogItem) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *HTTPLogItem) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *HTTPLogItem) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


