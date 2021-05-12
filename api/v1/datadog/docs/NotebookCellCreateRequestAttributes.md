# NotebookCellCreateRequestAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Definition** | [**LogStreamWidgetDefinition**](LogStreamWidgetDefinition.md) |  | 
**GraphSize** | Pointer to [**NotebookGraphSize**](NotebookGraphSize.md) |  | [optional] 
**SplitBy** | Pointer to [**NotebookSplitBy**](NotebookSplitBy.md) |  | [optional] 
**Time** | Pointer to [**NullableNotebookCellTime**](NotebookCellTime.md) |  | [optional] 

## Methods

### NewNotebookCellCreateRequestAttributes

`func NewNotebookCellCreateRequestAttributes(definition LogStreamWidgetDefinition, ) *NotebookCellCreateRequestAttributes`

NewNotebookCellCreateRequestAttributes instantiates a new NotebookCellCreateRequestAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotebookCellCreateRequestAttributesWithDefaults

`func NewNotebookCellCreateRequestAttributesWithDefaults() *NotebookCellCreateRequestAttributes`

NewNotebookCellCreateRequestAttributesWithDefaults instantiates a new NotebookCellCreateRequestAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinition

`func (o *NotebookCellCreateRequestAttributes) GetDefinition() LogStreamWidgetDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *NotebookCellCreateRequestAttributes) GetDefinitionOk() (*LogStreamWidgetDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *NotebookCellCreateRequestAttributes) SetDefinition(v LogStreamWidgetDefinition)`

SetDefinition sets Definition field to given value.


### GetGraphSize

`func (o *NotebookCellCreateRequestAttributes) GetGraphSize() NotebookGraphSize`

GetGraphSize returns the GraphSize field if non-nil, zero value otherwise.

### GetGraphSizeOk

`func (o *NotebookCellCreateRequestAttributes) GetGraphSizeOk() (*NotebookGraphSize, bool)`

GetGraphSizeOk returns a tuple with the GraphSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphSize

`func (o *NotebookCellCreateRequestAttributes) SetGraphSize(v NotebookGraphSize)`

SetGraphSize sets GraphSize field to given value.

### HasGraphSize

`func (o *NotebookCellCreateRequestAttributes) HasGraphSize() bool`

HasGraphSize returns a boolean if a field has been set.

### GetSplitBy

`func (o *NotebookCellCreateRequestAttributes) GetSplitBy() NotebookSplitBy`

GetSplitBy returns the SplitBy field if non-nil, zero value otherwise.

### GetSplitByOk

`func (o *NotebookCellCreateRequestAttributes) GetSplitByOk() (*NotebookSplitBy, bool)`

GetSplitByOk returns a tuple with the SplitBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitBy

`func (o *NotebookCellCreateRequestAttributes) SetSplitBy(v NotebookSplitBy)`

SetSplitBy sets SplitBy field to given value.

### HasSplitBy

`func (o *NotebookCellCreateRequestAttributes) HasSplitBy() bool`

HasSplitBy returns a boolean if a field has been set.

### GetTime

`func (o *NotebookCellCreateRequestAttributes) GetTime() NotebookCellTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *NotebookCellCreateRequestAttributes) GetTimeOk() (*NotebookCellTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *NotebookCellCreateRequestAttributes) SetTime(v NotebookCellTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *NotebookCellCreateRequestAttributes) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTimeNil

`func (o *NotebookCellCreateRequestAttributes) SetTimeNil(b bool)`

 SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime
`func (o *NotebookCellCreateRequestAttributes) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


