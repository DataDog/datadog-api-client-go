# NotebookLogStreamCellAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Definition** | [**LogStreamWidgetDefinition**](LogStreamWidgetDefinition.md) |  | 
**GraphSize** | Pointer to [**NotebookGraphSize**](NotebookGraphSize.md) |  | [optional] 
**Time** | Pointer to [**NullableNotebookCellTime**](NotebookCellTime.md) |  | [optional] 

## Methods

### NewNotebookLogStreamCellAttributes

`func NewNotebookLogStreamCellAttributes(definition LogStreamWidgetDefinition, ) *NotebookLogStreamCellAttributes`

NewNotebookLogStreamCellAttributes instantiates a new NotebookLogStreamCellAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotebookLogStreamCellAttributesWithDefaults

`func NewNotebookLogStreamCellAttributesWithDefaults() *NotebookLogStreamCellAttributes`

NewNotebookLogStreamCellAttributesWithDefaults instantiates a new NotebookLogStreamCellAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinition

`func (o *NotebookLogStreamCellAttributes) GetDefinition() LogStreamWidgetDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *NotebookLogStreamCellAttributes) GetDefinitionOk() (*LogStreamWidgetDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *NotebookLogStreamCellAttributes) SetDefinition(v LogStreamWidgetDefinition)`

SetDefinition sets Definition field to given value.


### GetGraphSize

`func (o *NotebookLogStreamCellAttributes) GetGraphSize() NotebookGraphSize`

GetGraphSize returns the GraphSize field if non-nil, zero value otherwise.

### GetGraphSizeOk

`func (o *NotebookLogStreamCellAttributes) GetGraphSizeOk() (*NotebookGraphSize, bool)`

GetGraphSizeOk returns a tuple with the GraphSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphSize

`func (o *NotebookLogStreamCellAttributes) SetGraphSize(v NotebookGraphSize)`

SetGraphSize sets GraphSize field to given value.

### HasGraphSize

`func (o *NotebookLogStreamCellAttributes) HasGraphSize() bool`

HasGraphSize returns a boolean if a field has been set.

### GetTime

`func (o *NotebookLogStreamCellAttributes) GetTime() NotebookCellTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *NotebookLogStreamCellAttributes) GetTimeOk() (*NotebookCellTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *NotebookLogStreamCellAttributes) SetTime(v NotebookCellTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *NotebookLogStreamCellAttributes) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTimeNil

`func (o *NotebookLogStreamCellAttributes) SetTimeNil(b bool)`

 SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime
`func (o *NotebookLogStreamCellAttributes) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


