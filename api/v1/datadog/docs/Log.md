# Log

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Content** | Pointer to [**LogContent**](LogContent.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID of the Log. | [optional] 

## Methods

### NewLog

`func NewLog() *Log`

NewLog instantiates a new Log object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogWithDefaults

`func NewLogWithDefaults() *Log`

NewLogWithDefaults instantiates a new Log object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetContent

`func (o *Log) GetContent() LogContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Log) GetContentOk() (*LogContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Log) SetContent(v LogContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *Log) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetId

`func (o *Log) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Log) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Log) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Log) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


