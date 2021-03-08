# FormulaAndFunctionMetricQueryDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregator** | Pointer to [**FormulaAndFunctionMetricAggregation**](FormulaAndFunctionMetricAggregation.md) |  | [optional] 
**DataSource** | [**FormulaAndFunctionMetricDataSource**](FormulaAndFunctionMetricDataSource.md) |  | 
**Name** | **string** | Name of the query for use in formulas. | 
**Query** | **string** | Metrics query definition. | 

## Methods

### NewFormulaAndFunctionMetricQueryDefinition

`func NewFormulaAndFunctionMetricQueryDefinition(dataSource FormulaAndFunctionMetricDataSource, name string, query string, ) *FormulaAndFunctionMetricQueryDefinition`

NewFormulaAndFunctionMetricQueryDefinition instantiates a new FormulaAndFunctionMetricQueryDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormulaAndFunctionMetricQueryDefinitionWithDefaults

`func NewFormulaAndFunctionMetricQueryDefinitionWithDefaults() *FormulaAndFunctionMetricQueryDefinition`

NewFormulaAndFunctionMetricQueryDefinitionWithDefaults instantiates a new FormulaAndFunctionMetricQueryDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregator

`func (o *FormulaAndFunctionMetricQueryDefinition) GetAggregator() FormulaAndFunctionMetricAggregation`

GetAggregator returns the Aggregator field if non-nil, zero value otherwise.

### GetAggregatorOk

`func (o *FormulaAndFunctionMetricQueryDefinition) GetAggregatorOk() (*FormulaAndFunctionMetricAggregation, bool)`

GetAggregatorOk returns a tuple with the Aggregator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregator

`func (o *FormulaAndFunctionMetricQueryDefinition) SetAggregator(v FormulaAndFunctionMetricAggregation)`

SetAggregator sets Aggregator field to given value.

### HasAggregator

`func (o *FormulaAndFunctionMetricQueryDefinition) HasAggregator() bool`

HasAggregator returns a boolean if a field has been set.

### GetDataSource

`func (o *FormulaAndFunctionMetricQueryDefinition) GetDataSource() FormulaAndFunctionMetricDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *FormulaAndFunctionMetricQueryDefinition) GetDataSourceOk() (*FormulaAndFunctionMetricDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *FormulaAndFunctionMetricQueryDefinition) SetDataSource(v FormulaAndFunctionMetricDataSource)`

SetDataSource sets DataSource field to given value.


### GetName

`func (o *FormulaAndFunctionMetricQueryDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormulaAndFunctionMetricQueryDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormulaAndFunctionMetricQueryDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetQuery

`func (o *FormulaAndFunctionMetricQueryDefinition) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *FormulaAndFunctionMetricQueryDefinition) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *FormulaAndFunctionMetricQueryDefinition) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


