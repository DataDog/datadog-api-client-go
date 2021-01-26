# TimeSeriesFormulaAndFunctionMetricQueryDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregator** | Pointer to [**FormulaAndFunctionMetricAggregation**](FormulaAndFunctionMetricAggregation.md) |  | [optional] 
**DataSource** | [**FormulaAndFunctionMetricDataSource**](FormulaAndFunctionMetricDataSource.md) |  | 
**Name** | Pointer to **string** | Name of the query for use in formulas. | [optional] 
**Query** | **string** | Metrics query definition. | 

## Methods

### NewTimeSeriesFormulaAndFunctionMetricQueryDefinition

`func NewTimeSeriesFormulaAndFunctionMetricQueryDefinition(dataSource FormulaAndFunctionMetricDataSource, query string, ) *TimeSeriesFormulaAndFunctionMetricQueryDefinition`

NewTimeSeriesFormulaAndFunctionMetricQueryDefinition instantiates a new TimeSeriesFormulaAndFunctionMetricQueryDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesFormulaAndFunctionMetricQueryDefinitionWithDefaults

`func NewTimeSeriesFormulaAndFunctionMetricQueryDefinitionWithDefaults() *TimeSeriesFormulaAndFunctionMetricQueryDefinition`

NewTimeSeriesFormulaAndFunctionMetricQueryDefinitionWithDefaults instantiates a new TimeSeriesFormulaAndFunctionMetricQueryDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregator

`func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) GetAggregator() FormulaAndFunctionMetricAggregation`

GetAggregator returns the Aggregator field if non-nil, zero value otherwise.

### GetAggregatorOk

`func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) GetAggregatorOk() (*FormulaAndFunctionMetricAggregation, bool)`

GetAggregatorOk returns a tuple with the Aggregator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregator

`func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) SetAggregator(v FormulaAndFunctionMetricAggregation)`

SetAggregator sets Aggregator field to given value.

### HasAggregator

`func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) HasAggregator() bool`

HasAggregator returns a boolean if a field has been set.

### GetDataSource

`func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) GetDataSource() FormulaAndFunctionMetricDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) GetDataSourceOk() (*FormulaAndFunctionMetricDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) SetDataSource(v FormulaAndFunctionMetricDataSource)`

SetDataSource sets DataSource field to given value.


### GetName

`func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuery

`func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


