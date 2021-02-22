# TimeSeriesFormulaAndFunctionEventQueryDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | [**TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute**](TimeSeriesFormulaAndFunctionEventQueryDefinition_compute.md) |  | 
**DataSource** | [**FormulaAndFunctionEventsDataSource**](FormulaAndFunctionEventsDataSource.md) |  | 
**GroupBy** | Pointer to [**[]TimeSeriesFormulaAndFunctionEventQueryGroupBy**](TimeSeriesFormulaAndFunctionEventQueryGroupBy.md) | Group by options. | [optional] 
**Indexes** | Pointer to **[]string** | An array of index names to query in the stream. Omit or use &#x60;[]&#x60; to query all indexes at once. | [optional] 
**Name** | Pointer to **string** | Name of the query for use in formulas. | [optional] 
**Search** | Pointer to [**TimeSeriesFormulaAndFunctionEventQueryDefinitionSearch**](TimeSeriesFormulaAndFunctionEventQueryDefinition_search.md) |  | [optional] 

## Methods

### NewTimeSeriesFormulaAndFunctionEventQueryDefinition

`func NewTimeSeriesFormulaAndFunctionEventQueryDefinition(compute TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute, dataSource FormulaAndFunctionEventsDataSource, ) *TimeSeriesFormulaAndFunctionEventQueryDefinition`

NewTimeSeriesFormulaAndFunctionEventQueryDefinition instantiates a new TimeSeriesFormulaAndFunctionEventQueryDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesFormulaAndFunctionEventQueryDefinitionWithDefaults

`func NewTimeSeriesFormulaAndFunctionEventQueryDefinitionWithDefaults() *TimeSeriesFormulaAndFunctionEventQueryDefinition`

NewTimeSeriesFormulaAndFunctionEventQueryDefinitionWithDefaults instantiates a new TimeSeriesFormulaAndFunctionEventQueryDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) GetCompute() TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) GetComputeOk() (*TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) SetCompute(v TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute)`

SetCompute sets Compute field to given value.


### GetDataSource

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) GetDataSource() FormulaAndFunctionEventsDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) GetDataSourceOk() (*FormulaAndFunctionEventsDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) SetDataSource(v FormulaAndFunctionEventsDataSource)`

SetDataSource sets DataSource field to given value.


### GetGroupBy

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) GetGroupBy() []TimeSeriesFormulaAndFunctionEventQueryGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) GetGroupByOk() (*[]TimeSeriesFormulaAndFunctionEventQueryGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) SetGroupBy(v []TimeSeriesFormulaAndFunctionEventQueryGroupBy)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetIndexes

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) GetIndexes() []string`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) GetIndexesOk() (*[]string, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) SetIndexes(v []string)`

SetIndexes sets Indexes field to given value.

### HasIndexes

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) HasIndexes() bool`

HasIndexes returns a boolean if a field has been set.

### GetName

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSearch

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) GetSearch() TimeSeriesFormulaAndFunctionEventQueryDefinitionSearch`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) GetSearchOk() (*TimeSeriesFormulaAndFunctionEventQueryDefinitionSearch, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) SetSearch(v TimeSeriesFormulaAndFunctionEventQueryDefinitionSearch)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinition) HasSearch() bool`

HasSearch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


