# NotebookDistributionCellAttributes

## Properties

| Name           | Type                                                                | Description | Notes      |
| -------------- | ------------------------------------------------------------------- | ----------- | ---------- |
| **Definition** | [**DistributionWidgetDefinition**](DistributionWidgetDefinition.md) |             |
| **GraphSize**  | Pointer to [**NotebookGraphSize**](NotebookGraphSize.md)            |             | [optional] |
| **SplitBy**    | Pointer to [**NotebookSplitBy**](NotebookSplitBy.md)                |             | [optional] |
| **Time**       | Pointer to [**NullableNotebookCellTime**](NotebookCellTime.md)      |             | [optional] |

## Methods

### NewNotebookDistributionCellAttributes

`func NewNotebookDistributionCellAttributes(definition DistributionWidgetDefinition) *NotebookDistributionCellAttributes`

NewNotebookDistributionCellAttributes instantiates a new NotebookDistributionCellAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewNotebookDistributionCellAttributesWithDefaults

`func NewNotebookDistributionCellAttributesWithDefaults() *NotebookDistributionCellAttributes`

NewNotebookDistributionCellAttributesWithDefaults instantiates a new NotebookDistributionCellAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDefinition

`func (o *NotebookDistributionCellAttributes) GetDefinition() DistributionWidgetDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *NotebookDistributionCellAttributes) GetDefinitionOk() (*DistributionWidgetDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *NotebookDistributionCellAttributes) SetDefinition(v DistributionWidgetDefinition)`

SetDefinition sets Definition field to given value.

### GetGraphSize

`func (o *NotebookDistributionCellAttributes) GetGraphSize() NotebookGraphSize`

GetGraphSize returns the GraphSize field if non-nil, zero value otherwise.

### GetGraphSizeOk

`func (o *NotebookDistributionCellAttributes) GetGraphSizeOk() (*NotebookGraphSize, bool)`

GetGraphSizeOk returns a tuple with the GraphSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphSize

`func (o *NotebookDistributionCellAttributes) SetGraphSize(v NotebookGraphSize)`

SetGraphSize sets GraphSize field to given value.

### HasGraphSize

`func (o *NotebookDistributionCellAttributes) HasGraphSize() bool`

HasGraphSize returns a boolean if a field has been set.

### GetSplitBy

`func (o *NotebookDistributionCellAttributes) GetSplitBy() NotebookSplitBy`

GetSplitBy returns the SplitBy field if non-nil, zero value otherwise.

### GetSplitByOk

`func (o *NotebookDistributionCellAttributes) GetSplitByOk() (*NotebookSplitBy, bool)`

GetSplitByOk returns a tuple with the SplitBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitBy

`func (o *NotebookDistributionCellAttributes) SetSplitBy(v NotebookSplitBy)`

SetSplitBy sets SplitBy field to given value.

### HasSplitBy

`func (o *NotebookDistributionCellAttributes) HasSplitBy() bool`

HasSplitBy returns a boolean if a field has been set.

### GetTime

`func (o *NotebookDistributionCellAttributes) GetTime() NotebookCellTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *NotebookDistributionCellAttributes) GetTimeOk() (*NotebookCellTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *NotebookDistributionCellAttributes) SetTime(v NotebookCellTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *NotebookDistributionCellAttributes) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTimeNil

`func (o *NotebookDistributionCellAttributes) SetTimeNil(b bool)`

SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime

`func (o *NotebookDistributionCellAttributes) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
