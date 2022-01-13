# ApmStatsQueryColumnType

## Properties

| Name                | Type                                                                       | Description                           | Notes      |
| ------------------- | -------------------------------------------------------------------------- | ------------------------------------- | ---------- |
| **Alias**           | Pointer to **string**                                                      | A user-assigned alias for the column. | [optional] |
| **CellDisplayMode** | Pointer to [**TableWidgetCellDisplayMode**](TableWidgetCellDisplayMode.md) |                                       | [optional] |
| **Name**            | **string**                                                                 | Column name.                          |
| **Order**           | Pointer to [**WidgetSort**](WidgetSort.md)                                 |                                       | [optional] |

## Methods

### NewApmStatsQueryColumnType

`func NewApmStatsQueryColumnType(name string) *ApmStatsQueryColumnType`

NewApmStatsQueryColumnType instantiates a new ApmStatsQueryColumnType object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewApmStatsQueryColumnTypeWithDefaults

`func NewApmStatsQueryColumnTypeWithDefaults() *ApmStatsQueryColumnType`

NewApmStatsQueryColumnTypeWithDefaults instantiates a new ApmStatsQueryColumnType object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAlias

`func (o *ApmStatsQueryColumnType) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ApmStatsQueryColumnType) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ApmStatsQueryColumnType) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ApmStatsQueryColumnType) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetCellDisplayMode

`func (o *ApmStatsQueryColumnType) GetCellDisplayMode() TableWidgetCellDisplayMode`

GetCellDisplayMode returns the CellDisplayMode field if non-nil, zero value otherwise.

### GetCellDisplayModeOk

`func (o *ApmStatsQueryColumnType) GetCellDisplayModeOk() (*TableWidgetCellDisplayMode, bool)`

GetCellDisplayModeOk returns a tuple with the CellDisplayMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellDisplayMode

`func (o *ApmStatsQueryColumnType) SetCellDisplayMode(v TableWidgetCellDisplayMode)`

SetCellDisplayMode sets CellDisplayMode field to given value.

### HasCellDisplayMode

`func (o *ApmStatsQueryColumnType) HasCellDisplayMode() bool`

HasCellDisplayMode returns a boolean if a field has been set.

### GetName

`func (o *ApmStatsQueryColumnType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApmStatsQueryColumnType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApmStatsQueryColumnType) SetName(v string)`

SetName sets Name field to given value.

### GetOrder

`func (o *ApmStatsQueryColumnType) GetOrder() WidgetSort`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ApmStatsQueryColumnType) GetOrderOk() (*WidgetSort, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ApmStatsQueryColumnType) SetOrder(v WidgetSort)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ApmStatsQueryColumnType) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
