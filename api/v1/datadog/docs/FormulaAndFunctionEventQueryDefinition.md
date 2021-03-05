# FormulaAndFunctionEventQueryDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | [**FormulaAndFunctionEventQueryDefinitionCompute**](FormulaAndFunctionEventQueryDefinition_compute.md) |  | 
**DataSource** | [**FormulaAndFunctionEventsDataSource**](FormulaAndFunctionEventsDataSource.md) |  | 
**GroupBy** | Pointer to [**[]FormulaAndFunctionEventQueryGroupBy**](FormulaAndFunctionEventQueryGroupBy.md) | Group by options. | [optional] 
**Indexes** | Pointer to **[]string** | An array of index names to query in the stream. Omit or use &#x60;[]&#x60; to query all indexes at once. | [optional] 
**Name** | Pointer to **string** | Name of the query for use in formulas. | [optional] 
**Search** | Pointer to [**FormulaAndFunctionEventQueryDefinitionSearch**](FormulaAndFunctionEventQueryDefinition_search.md) |  | [optional] 

## Methods

### NewFormulaAndFunctionEventQueryDefinition

`func NewFormulaAndFunctionEventQueryDefinition(compute FormulaAndFunctionEventQueryDefinitionCompute, dataSource FormulaAndFunctionEventsDataSource, ) *FormulaAndFunctionEventQueryDefinition`

NewFormulaAndFunctionEventQueryDefinition instantiates a new FormulaAndFunctionEventQueryDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormulaAndFunctionEventQueryDefinitionWithDefaults

`func NewFormulaAndFunctionEventQueryDefinitionWithDefaults() *FormulaAndFunctionEventQueryDefinition`

NewFormulaAndFunctionEventQueryDefinitionWithDefaults instantiates a new FormulaAndFunctionEventQueryDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *FormulaAndFunctionEventQueryDefinition) GetCompute() FormulaAndFunctionEventQueryDefinitionCompute`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *FormulaAndFunctionEventQueryDefinition) GetComputeOk() (*FormulaAndFunctionEventQueryDefinitionCompute, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *FormulaAndFunctionEventQueryDefinition) SetCompute(v FormulaAndFunctionEventQueryDefinitionCompute)`

SetCompute sets Compute field to given value.


### GetDataSource

`func (o *FormulaAndFunctionEventQueryDefinition) GetDataSource() FormulaAndFunctionEventsDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *FormulaAndFunctionEventQueryDefinition) GetDataSourceOk() (*FormulaAndFunctionEventsDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *FormulaAndFunctionEventQueryDefinition) SetDataSource(v FormulaAndFunctionEventsDataSource)`

SetDataSource sets DataSource field to given value.


### GetGroupBy

`func (o *FormulaAndFunctionEventQueryDefinition) GetGroupBy() []FormulaAndFunctionEventQueryGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *FormulaAndFunctionEventQueryDefinition) GetGroupByOk() (*[]FormulaAndFunctionEventQueryGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *FormulaAndFunctionEventQueryDefinition) SetGroupBy(v []FormulaAndFunctionEventQueryGroupBy)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *FormulaAndFunctionEventQueryDefinition) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetIndexes

`func (o *FormulaAndFunctionEventQueryDefinition) GetIndexes() []string`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *FormulaAndFunctionEventQueryDefinition) GetIndexesOk() (*[]string, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *FormulaAndFunctionEventQueryDefinition) SetIndexes(v []string)`

SetIndexes sets Indexes field to given value.

### HasIndexes

`func (o *FormulaAndFunctionEventQueryDefinition) HasIndexes() bool`

HasIndexes returns a boolean if a field has been set.

### GetName

`func (o *FormulaAndFunctionEventQueryDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormulaAndFunctionEventQueryDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormulaAndFunctionEventQueryDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FormulaAndFunctionEventQueryDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSearch

`func (o *FormulaAndFunctionEventQueryDefinition) GetSearch() FormulaAndFunctionEventQueryDefinitionSearch`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *FormulaAndFunctionEventQueryDefinition) GetSearchOk() (*FormulaAndFunctionEventQueryDefinitionSearch, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *FormulaAndFunctionEventQueryDefinition) SetSearch(v FormulaAndFunctionEventQueryDefinitionSearch)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *FormulaAndFunctionEventQueryDefinition) HasSearch() bool`

HasSearch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


