# LogContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**map[string]interface{}**](interface{}.md) | TODO. | [optional] 
**Host** | Pointer to **string** | Name of the machine from where the logs are being sent. | [optional] 
**Message** | Pointer to **string** | TODO. | [optional] 
**Service** | Pointer to **string** | TODO. | [optional] 
**Tags** | Pointer to [**[]interface{}**](interface{}.md) | TODO. | [optional] 
**Timestamp** | Pointer to [**time.Time**](time.Time.md) | TODO. | [optional] 

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

`func (o *LogContent) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *LogContent) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *LogContent) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetHost

`func (o *LogContent) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *LogContent) GetHostOk() (string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHost

`func (o *LogContent) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHost

`func (o *LogContent) SetHost(v string)`

SetHost gets a reference to the given string and assigns it to the Host field.

### GetMessage

`func (o *LogContent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LogContent) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *LogContent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *LogContent) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.

### GetService

`func (o *LogContent) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *LogContent) GetServiceOk() (string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasService

`func (o *LogContent) HasService() bool`

HasService returns a boolean if a field has been set.

### SetService

`func (o *LogContent) SetService(v string)`

SetService gets a reference to the given string and assigns it to the Service field.

### GetTags

`func (o *LogContent) GetTags() []interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LogContent) GetTagsOk() ([]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *LogContent) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *LogContent) SetTags(v []interface{})`

SetTags gets a reference to the given []interface{} and assigns it to the Tags field.

### GetTimestamp

`func (o *LogContent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LogContent) GetTimestampOk() (time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimestamp

`func (o *LogContent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestamp

`func (o *LogContent) SetTimestamp(v time.Time)`

SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


