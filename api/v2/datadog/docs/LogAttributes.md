# LogAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | Pointer to **map[string]interface{}** | JSON object of attributes from your log. | [optional] 
**Host** | Pointer to **string** | Name of the machine from where the logs are being sent. | [optional] 
**Message** | Pointer to **string** | The message [reserved attribute](https://docs.datadoghq.com/logs/log_collection/#reserved-attributes) of your log. By default, Datadog ingests the value of the message attribute as the body of the log entry. That value is then highlighted and displayed in the Logstream, where it is indexed for full text search. | [optional] 
**Service** | Pointer to **string** | The name of the application or service generating the log events. It is used to switch from Logs to APM, so make sure you define the same value when you use both products. | [optional] 
**Status** | Pointer to **string** | Status of the message associated with your log. | [optional] 
**Tags** | Pointer to **[]string** | Array of tags associated with your log. | [optional] 
**Timestamp** | Pointer to **time.Time** | Timestamp of your log. | [optional] 

## Methods

### NewLogAttributes

`func NewLogAttributes() *LogAttributes`

NewLogAttributes instantiates a new LogAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogAttributesWithDefaults

`func NewLogAttributesWithDefaults() *LogAttributes`

NewLogAttributesWithDefaults instantiates a new LogAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *LogAttributes) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LogAttributes) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LogAttributes) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *LogAttributes) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetHost

`func (o *LogAttributes) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *LogAttributes) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *LogAttributes) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *LogAttributes) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetMessage

`func (o *LogAttributes) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LogAttributes) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LogAttributes) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LogAttributes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetService

`func (o *LogAttributes) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *LogAttributes) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *LogAttributes) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *LogAttributes) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *LogAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LogAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LogAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LogAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *LogAttributes) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LogAttributes) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LogAttributes) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LogAttributes) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimestamp

`func (o *LogAttributes) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LogAttributes) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LogAttributes) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LogAttributes) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


