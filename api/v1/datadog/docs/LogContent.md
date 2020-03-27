# LogContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]interface{}** | JSON object of attributes from your log. | [optional] 
**Host** | Pointer to **string** | Name of the machine from where the logs are being sent. | [optional] 
**Message** | Pointer to **string** | The message [reserved attribute](https://docs.datadoghq.com/logs/log_collection/#reserved-attributes) of your log. By default, Datadog ingests the value of the message attribute as the body of the log entry. That value is then highlighted and displayed in the Logstream, where it is indexed for full text search. | [optional] 
**Service** | Pointer to **string** | The name of the application or service generating the log events. It is used to switch from Logs to APM, so make sure you define the same value when you use both products. | [optional] 
**Tags** | Pointer to **[]interface{}** | Array of tags associated with your log. | [optional] 
**Timestamp** | Pointer to [**time.Time**](time.Time.md) | Timestamp of your log. | [optional] 

## Methods

### NewLogContent

`func NewLogContent() *LogContent`

NewLogContent instantiates a new LogContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogContentWithDefaults

`func NewLogContentWithDefaults() *LogContent`

NewLogContentWithDefaults instantiates a new LogContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *LogContent) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LogContent) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LogContent) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *LogContent) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetHost

`func (o *LogContent) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *LogContent) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *LogContent) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *LogContent) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetMessage

`func (o *LogContent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LogContent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LogContent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LogContent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetService

`func (o *LogContent) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *LogContent) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *LogContent) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *LogContent) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTags

`func (o *LogContent) GetTags() []interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LogContent) GetTagsOk() (*[]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LogContent) SetTags(v []interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *LogContent) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimestamp

`func (o *LogContent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LogContent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LogContent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LogContent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


