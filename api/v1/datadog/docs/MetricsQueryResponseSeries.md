# MetricsQueryResponseSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggr** | Pointer to **string** | Aggregation type. | [optional] [readonly] 
**DisplayName** | Pointer to **string** | Display name of the metric. | [optional] [readonly] 
**End** | Pointer to **int64** | End of the time window, milliseconds since Unix epoch. | [optional] [readonly] 
**Expression** | Pointer to **string** | Metric expression. | [optional] [readonly] 
**Interval** | Pointer to **int64** | Number of seconds between data samples. | [optional] [readonly] 
**Length** | Pointer to **int64** | Number of data samples. | [optional] [readonly] 
**Metric** | Pointer to **string** | Metric name. | [optional] [readonly] 
**Pointlist** | Pointer to **[][]float64** | List of points of the time series. | [optional] [readonly] 
**Scope** | Pointer to **string** | Metric scope, comma separated list of tags. | [optional] [readonly] 
**Start** | Pointer to **int64** | Start of the time window, milliseconds since Unix epoch. | [optional] [readonly] 
**Unit** | Pointer to [**[]MetricsQueryResponseUnit**](MetricsQueryResponseUnit.md) | Detailed information about the metric unit. First element describes the \&quot;primary unit\&quot; (for example, &#x60;bytes&#x60; in &#x60;bytes per second&#x60;), second describes the \&quot;per unit\&quot; (for example, &#x60;second&#x60; in &#x60;bytes per second&#x60;). | [optional] [readonly] 

## Methods

### NewMetricsQueryResponseSeries

`func NewMetricsQueryResponseSeries() *MetricsQueryResponseSeries`

NewMetricsQueryResponseSeries instantiates a new MetricsQueryResponseSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsQueryResponseSeriesWithDefaults

`func NewMetricsQueryResponseSeriesWithDefaults() *MetricsQueryResponseSeries`

NewMetricsQueryResponseSeriesWithDefaults instantiates a new MetricsQueryResponseSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggr

`func (o *MetricsQueryResponseSeries) GetAggr() string`

GetAggr returns the Aggr field if non-nil, zero value otherwise.

### GetAggrOk

`func (o *MetricsQueryResponseSeries) GetAggrOk() (*string, bool)`

GetAggrOk returns a tuple with the Aggr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggr

`func (o *MetricsQueryResponseSeries) SetAggr(v string)`

SetAggr sets Aggr field to given value.

### HasAggr

`func (o *MetricsQueryResponseSeries) HasAggr() bool`

HasAggr returns a boolean if a field has been set.

### GetDisplayName

`func (o *MetricsQueryResponseSeries) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MetricsQueryResponseSeries) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MetricsQueryResponseSeries) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MetricsQueryResponseSeries) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnd

`func (o *MetricsQueryResponseSeries) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *MetricsQueryResponseSeries) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *MetricsQueryResponseSeries) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *MetricsQueryResponseSeries) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetExpression

`func (o *MetricsQueryResponseSeries) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *MetricsQueryResponseSeries) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *MetricsQueryResponseSeries) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *MetricsQueryResponseSeries) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetInterval

`func (o *MetricsQueryResponseSeries) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *MetricsQueryResponseSeries) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *MetricsQueryResponseSeries) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *MetricsQueryResponseSeries) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetLength

`func (o *MetricsQueryResponseSeries) GetLength() int64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *MetricsQueryResponseSeries) GetLengthOk() (*int64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *MetricsQueryResponseSeries) SetLength(v int64)`

SetLength sets Length field to given value.

### HasLength

`func (o *MetricsQueryResponseSeries) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetMetric

`func (o *MetricsQueryResponseSeries) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *MetricsQueryResponseSeries) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *MetricsQueryResponseSeries) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *MetricsQueryResponseSeries) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetPointlist

`func (o *MetricsQueryResponseSeries) GetPointlist() [][]float64`

GetPointlist returns the Pointlist field if non-nil, zero value otherwise.

### GetPointlistOk

`func (o *MetricsQueryResponseSeries) GetPointlistOk() (*[][]float64, bool)`

GetPointlistOk returns a tuple with the Pointlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointlist

`func (o *MetricsQueryResponseSeries) SetPointlist(v [][]float64)`

SetPointlist sets Pointlist field to given value.

### HasPointlist

`func (o *MetricsQueryResponseSeries) HasPointlist() bool`

HasPointlist returns a boolean if a field has been set.

### GetScope

`func (o *MetricsQueryResponseSeries) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *MetricsQueryResponseSeries) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *MetricsQueryResponseSeries) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *MetricsQueryResponseSeries) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStart

`func (o *MetricsQueryResponseSeries) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *MetricsQueryResponseSeries) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *MetricsQueryResponseSeries) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *MetricsQueryResponseSeries) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUnit

`func (o *MetricsQueryResponseSeries) GetUnit() []MetricsQueryResponseUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricsQueryResponseSeries) GetUnitOk() (*[]MetricsQueryResponseUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricsQueryResponseSeries) SetUnit(v []MetricsQueryResponseUnit)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricsQueryResponseSeries) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


