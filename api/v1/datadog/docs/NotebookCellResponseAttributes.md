# NotebookCellResponseAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Definition** | [**LogStreamWidgetDefinition**](LogStreamWidgetDefinition.md) |  | 
**GraphSize** | Pointer to [**NotebookGraphSize**](NotebookGraphSize.md) |  | [optional] 
**SplitBy** | Pointer to [**NotebookSplitBy**](NotebookSplitBy.md) |  | [optional] 
**Time** | Pointer to [**NullableNotebookCellTime**](NotebookCellTime.md) |  | [optional] 

## Methods

### NewNotebookCellResponseAttributes

`func NewNotebookCellResponseAttributes(definition LogStreamWidgetDefinition, ) *NotebookCellResponseAttributes`

NewNotebookCellResponseAttributes instantiates a new NotebookCellResponseAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotebookCellResponseAttributesWithDefaults

`func NewNotebookCellResponseAttributesWithDefaults() *NotebookCellResponseAttributes`

NewNotebookCellResponseAttributesWithDefaults instantiates a new NotebookCellResponseAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinition

`func (o *NotebookCellResponseAttributes) GetDefinition() LogStreamWidgetDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *NotebookCellResponseAttributes) GetDefinitionOk() (*LogStreamWidgetDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *NotebookCellResponseAttributes) SetDefinition(v LogStreamWidgetDefinition)`

SetDefinition sets Definition field to given value.


### GetGraphSize

`func (o *NotebookCellResponseAttributes) GetGraphSize() NotebookGraphSize`

GetGraphSize returns the GraphSize field if non-nil, zero value otherwise.

### GetGraphSizeOk

`func (o *NotebookCellResponseAttributes) GetGraphSizeOk() (*NotebookGraphSize, bool)`

GetGraphSizeOk returns a tuple with the GraphSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphSize

`func (o *NotebookCellResponseAttributes) SetGraphSize(v NotebookGraphSize)`

SetGraphSize sets GraphSize field to given value.

### HasGraphSize

`func (o *NotebookCellResponseAttributes) HasGraphSize() bool`

HasGraphSize returns a boolean if a field has been set.

### GetSplitBy

`func (o *NotebookCellResponseAttributes) GetSplitBy() NotebookSplitBy`

GetSplitBy returns the SplitBy field if non-nil, zero value otherwise.

### GetSplitByOk

`func (o *NotebookCellResponseAttributes) GetSplitByOk() (*NotebookSplitBy, bool)`

GetSplitByOk returns a tuple with the SplitBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitBy

`func (o *NotebookCellResponseAttributes) SetSplitBy(v NotebookSplitBy)`

SetSplitBy sets SplitBy field to given value.

### HasSplitBy

`func (o *NotebookCellResponseAttributes) HasSplitBy() bool`

HasSplitBy returns a boolean if a field has been set.

### GetTime

`func (o *NotebookCellResponseAttributes) GetTime() NotebookCellTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *NotebookCellResponseAttributes) GetTimeOk() (*NotebookCellTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *NotebookCellResponseAttributes) SetTime(v NotebookCellTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *NotebookCellResponseAttributes) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTimeNil

`func (o *NotebookCellResponseAttributes) SetTimeNil(b bool)`

 SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime
`func (o *NotebookCellResponseAttributes) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


