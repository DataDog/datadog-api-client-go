# Log

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**LogAttributes**](Log_attributes.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID of the Log. | [optional] 
**Type** | Pointer to [**LogType**](LogType.md) |  | [optional] [default to "log"]

## Methods

### NewLog

`func NewLog() *Log`

NewLog instantiates a new Log object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogWithDefaults

`func NewLogWithDefaults() *Log`

NewLogWithDefaults instantiates a new Log object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *Log) GetAttributes() LogAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Log) GetAttributesOk() (*LogAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Log) SetAttributes(v LogAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Log) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

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

### GetType

`func (o *Log) GetType() LogType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Log) GetTypeOk() (*LogType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Log) SetType(v LogType)`

SetType sets Type field to given value.

### HasType

`func (o *Log) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


