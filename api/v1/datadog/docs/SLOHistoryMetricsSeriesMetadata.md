# SLOHistoryMetricsSeriesMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggr** | Pointer to **string** | Query aggregator function. | [optional] 
**Expression** | Pointer to **string** | Query expression. | [optional] 
**Metric** | Pointer to **string** | Query metric used. | [optional] 
**QueryIndex** | Pointer to **int64** | Query index from original combined query. | [optional] 
**Scope** | Pointer to **string** | Query scope. | [optional] 
**Unit** | Pointer to [**[]SLOHistoryMetricsSeriesMetadataUnit**](SLOHistoryMetricsSeriesMetadataUnit.md) | An array of metric units that contains up to two unit objects. For example, bytes represents one unit object and bytes per second represents two unit objects. If a metric query only has one unit object, the second array element is null. | [optional] 

## Methods

### NewSLOHistoryMetricsSeriesMetadata

`func NewSLOHistoryMetricsSeriesMetadata() *SLOHistoryMetricsSeriesMetadata`

NewSLOHistoryMetricsSeriesMetadata instantiates a new SLOHistoryMetricsSeriesMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLOHistoryMetricsSeriesMetadataWithDefaults

`func NewSLOHistoryMetricsSeriesMetadataWithDefaults() *SLOHistoryMetricsSeriesMetadata`

NewSLOHistoryMetricsSeriesMetadataWithDefaults instantiates a new SLOHistoryMetricsSeriesMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggr

`func (o *SLOHistoryMetricsSeriesMetadata) GetAggr() string`

GetAggr returns the Aggr field if non-nil, zero value otherwise.

### GetAggrOk

`func (o *SLOHistoryMetricsSeriesMetadata) GetAggrOk() (*string, bool)`

GetAggrOk returns a tuple with the Aggr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggr

`func (o *SLOHistoryMetricsSeriesMetadata) SetAggr(v string)`

SetAggr sets Aggr field to given value.

### HasAggr

`func (o *SLOHistoryMetricsSeriesMetadata) HasAggr() bool`

HasAggr returns a boolean if a field has been set.

### GetExpression

`func (o *SLOHistoryMetricsSeriesMetadata) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *SLOHistoryMetricsSeriesMetadata) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *SLOHistoryMetricsSeriesMetadata) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *SLOHistoryMetricsSeriesMetadata) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetMetric

`func (o *SLOHistoryMetricsSeriesMetadata) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *SLOHistoryMetricsSeriesMetadata) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *SLOHistoryMetricsSeriesMetadata) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *SLOHistoryMetricsSeriesMetadata) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetQueryIndex

`func (o *SLOHistoryMetricsSeriesMetadata) GetQueryIndex() int64`

GetQueryIndex returns the QueryIndex field if non-nil, zero value otherwise.

### GetQueryIndexOk

`func (o *SLOHistoryMetricsSeriesMetadata) GetQueryIndexOk() (*int64, bool)`

GetQueryIndexOk returns a tuple with the QueryIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryIndex

`func (o *SLOHistoryMetricsSeriesMetadata) SetQueryIndex(v int64)`

SetQueryIndex sets QueryIndex field to given value.

### HasQueryIndex

`func (o *SLOHistoryMetricsSeriesMetadata) HasQueryIndex() bool`

HasQueryIndex returns a boolean if a field has been set.

### GetScope

`func (o *SLOHistoryMetricsSeriesMetadata) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SLOHistoryMetricsSeriesMetadata) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SLOHistoryMetricsSeriesMetadata) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SLOHistoryMetricsSeriesMetadata) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetUnit

`func (o *SLOHistoryMetricsSeriesMetadata) GetUnit() []SLOHistoryMetricsSeriesMetadataUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *SLOHistoryMetricsSeriesMetadata) GetUnitOk() (*[]SLOHistoryMetricsSeriesMetadataUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *SLOHistoryMetricsSeriesMetadata) SetUnit(v []SLOHistoryMetricsSeriesMetadataUnit)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *SLOHistoryMetricsSeriesMetadata) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *SLOHistoryMetricsSeriesMetadata) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *SLOHistoryMetricsSeriesMetadata) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


