# MonitorFormulaAndFunctionEventQueryDefinition

## Properties

| Name           | Type                                                                                                                         | Description                                                                                              | Notes      |
| -------------- | ---------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | ---------- |
| **Compute**    | [**MonitorFormulaAndFunctionEventQueryDefinitionCompute**](MonitorFormulaAndFunctionEventQueryDefinitionCompute.md)          |                                                                                                          |
| **DataSource** | [**MonitorFormulaAndFunctionEventsDataSource**](MonitorFormulaAndFunctionEventsDataSource.md)                                |                                                                                                          |
| **GroupBy**    | Pointer to [**[]MonitorFormulaAndFunctionEventQueryGroupBy**](MonitorFormulaAndFunctionEventQueryGroupBy.md)                 | Group by options.                                                                                        | [optional] |
| **Indexes**    | Pointer to **[]string**                                                                                                      | An array of index names to query in the stream. Omit or use &#x60;[]&#x60; to query all indexes at once. | [optional] |
| **Name**       | **string**                                                                                                                   | Name of the query for use in formulas.                                                                   |
| **Search**     | Pointer to [**MonitorFormulaAndFunctionEventQueryDefinitionSearch**](MonitorFormulaAndFunctionEventQueryDefinitionSearch.md) |                                                                                                          | [optional] |

## Methods

### NewMonitorFormulaAndFunctionEventQueryDefinition

`func NewMonitorFormulaAndFunctionEventQueryDefinition(compute MonitorFormulaAndFunctionEventQueryDefinitionCompute, dataSource MonitorFormulaAndFunctionEventsDataSource, name string) *MonitorFormulaAndFunctionEventQueryDefinition`

NewMonitorFormulaAndFunctionEventQueryDefinition instantiates a new MonitorFormulaAndFunctionEventQueryDefinition object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMonitorFormulaAndFunctionEventQueryDefinitionWithDefaults

`func NewMonitorFormulaAndFunctionEventQueryDefinitionWithDefaults() *MonitorFormulaAndFunctionEventQueryDefinition`

NewMonitorFormulaAndFunctionEventQueryDefinitionWithDefaults instantiates a new MonitorFormulaAndFunctionEventQueryDefinition object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCompute

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetCompute() MonitorFormulaAndFunctionEventQueryDefinitionCompute`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetComputeOk() (*MonitorFormulaAndFunctionEventQueryDefinitionCompute, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) SetCompute(v MonitorFormulaAndFunctionEventQueryDefinitionCompute)`

SetCompute sets Compute field to given value.

### GetDataSource

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetDataSource() MonitorFormulaAndFunctionEventsDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetDataSourceOk() (*MonitorFormulaAndFunctionEventsDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) SetDataSource(v MonitorFormulaAndFunctionEventsDataSource)`

SetDataSource sets DataSource field to given value.

### GetGroupBy

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetGroupBy() []MonitorFormulaAndFunctionEventQueryGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetGroupByOk() (*[]MonitorFormulaAndFunctionEventQueryGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) SetGroupBy(v []MonitorFormulaAndFunctionEventQueryGroupBy)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetIndexes

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetIndexes() []string`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetIndexesOk() (*[]string, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) SetIndexes(v []string)`

SetIndexes sets Indexes field to given value.

### HasIndexes

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) HasIndexes() bool`

HasIndexes returns a boolean if a field has been set.

### GetName

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) SetName(v string)`

SetName sets Name field to given value.

### GetSearch

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetSearch() MonitorFormulaAndFunctionEventQueryDefinitionSearch`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetSearchOk() (*MonitorFormulaAndFunctionEventQueryDefinitionSearch, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) SetSearch(v MonitorFormulaAndFunctionEventQueryDefinitionSearch)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *MonitorFormulaAndFunctionEventQueryDefinition) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
