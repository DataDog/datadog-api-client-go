# MetricsQueryMetadata

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Aggr** | Pointer to **string** | Aggregation type. | [optional] [readonly] 
**DisplayName** | Pointer to **string** | Display name of the metric. | [optional] [readonly] 
**End** | Pointer to **int64** | End of the time window, milliseconds since Unix epoch. | [optional] [readonly] 
**Expression** | Pointer to **string** | Metric expression. | [optional] [readonly] 
**Interval** | Pointer to **int64** | Number of seconds between data samples. | [optional] [readonly] 
**Length** | Pointer to **int64** | Number of data samples. | [optional] [readonly] 
**Metric** | Pointer to **string** | Metric name. | [optional] [readonly] 
**Pointlist** | Pointer to **[][]float64** | List of points of the time series. | [optional] [readonly] 
**QueryIndex** | Pointer to **int64** | The index of the series&#39; query within the request. | [optional] [readonly] 
**Scope** | Pointer to **string** | Metric scope, comma separated list of tags. | [optional] [readonly] 
**Start** | Pointer to **int64** | Start of the time window, milliseconds since Unix epoch. | [optional] [readonly] 
**TagSet** | Pointer to **[]string** | Unique tags identifying this series. | [optional] [readonly] 
**Unit** | Pointer to [**[]MetricsQueryUnit**](MetricsQueryUnit.md) | Detailed information about the metric unit. First element describes the \&quot;primary unit\&quot; (for example, &#x60;bytes&#x60; in &#x60;bytes per second&#x60;), second describes the \&quot;per unit\&quot; (for example, &#x60;second&#x60; in &#x60;bytes per second&#x60;). | [optional] [readonly] 

## Methods

### NewMetricsQueryMetadata

`func NewMetricsQueryMetadata() *MetricsQueryMetadata`

NewMetricsQueryMetadata instantiates a new MetricsQueryMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsQueryMetadataWithDefaults

`func NewMetricsQueryMetadataWithDefaults() *MetricsQueryMetadata`

NewMetricsQueryMetadataWithDefaults instantiates a new MetricsQueryMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggr

`func (o *MetricsQueryMetadata) GetAggr() string`

GetAggr returns the Aggr field if non-nil, zero value otherwise.

### GetAggrOk

`func (o *MetricsQueryMetadata) GetAggrOk() (*string, bool)`

GetAggrOk returns a tuple with the Aggr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggr

`func (o *MetricsQueryMetadata) SetAggr(v string)`

SetAggr sets Aggr field to given value.

### HasAggr

`func (o *MetricsQueryMetadata) HasAggr() bool`

HasAggr returns a boolean if a field has been set.

### GetDisplayName

`func (o *MetricsQueryMetadata) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MetricsQueryMetadata) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MetricsQueryMetadata) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MetricsQueryMetadata) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnd

`func (o *MetricsQueryMetadata) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *MetricsQueryMetadata) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *MetricsQueryMetadata) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *MetricsQueryMetadata) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetExpression

`func (o *MetricsQueryMetadata) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *MetricsQueryMetadata) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *MetricsQueryMetadata) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *MetricsQueryMetadata) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetInterval

`func (o *MetricsQueryMetadata) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *MetricsQueryMetadata) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *MetricsQueryMetadata) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *MetricsQueryMetadata) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetLength

`func (o *MetricsQueryMetadata) GetLength() int64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *MetricsQueryMetadata) GetLengthOk() (*int64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *MetricsQueryMetadata) SetLength(v int64)`

SetLength sets Length field to given value.

### HasLength

`func (o *MetricsQueryMetadata) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetMetric

`func (o *MetricsQueryMetadata) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *MetricsQueryMetadata) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *MetricsQueryMetadata) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *MetricsQueryMetadata) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetPointlist

`func (o *MetricsQueryMetadata) GetPointlist() [][]float64`

GetPointlist returns the Pointlist field if non-nil, zero value otherwise.

### GetPointlistOk

`func (o *MetricsQueryMetadata) GetPointlistOk() (*[][]float64, bool)`

GetPointlistOk returns a tuple with the Pointlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointlist

`func (o *MetricsQueryMetadata) SetPointlist(v [][]float64)`

SetPointlist sets Pointlist field to given value.

### HasPointlist

`func (o *MetricsQueryMetadata) HasPointlist() bool`

HasPointlist returns a boolean if a field has been set.

### GetQueryIndex

`func (o *MetricsQueryMetadata) GetQueryIndex() int64`

GetQueryIndex returns the QueryIndex field if non-nil, zero value otherwise.

### GetQueryIndexOk

`func (o *MetricsQueryMetadata) GetQueryIndexOk() (*int64, bool)`

GetQueryIndexOk returns a tuple with the QueryIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryIndex

`func (o *MetricsQueryMetadata) SetQueryIndex(v int64)`

SetQueryIndex sets QueryIndex field to given value.

### HasQueryIndex

`func (o *MetricsQueryMetadata) HasQueryIndex() bool`

HasQueryIndex returns a boolean if a field has been set.

### GetScope

`func (o *MetricsQueryMetadata) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *MetricsQueryMetadata) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *MetricsQueryMetadata) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *MetricsQueryMetadata) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStart

`func (o *MetricsQueryMetadata) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *MetricsQueryMetadata) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *MetricsQueryMetadata) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *MetricsQueryMetadata) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTagSet

`func (o *MetricsQueryMetadata) GetTagSet() []string`

GetTagSet returns the TagSet field if non-nil, zero value otherwise.

### GetTagSetOk

`func (o *MetricsQueryMetadata) GetTagSetOk() (*[]string, bool)`

GetTagSetOk returns a tuple with the TagSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagSet

`func (o *MetricsQueryMetadata) SetTagSet(v []string)`

SetTagSet sets TagSet field to given value.

### HasTagSet

`func (o *MetricsQueryMetadata) HasTagSet() bool`

HasTagSet returns a boolean if a field has been set.

### GetUnit

`func (o *MetricsQueryMetadata) GetUnit() []MetricsQueryUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricsQueryMetadata) GetUnitOk() (*[]MetricsQueryUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricsQueryMetadata) SetUnit(v []MetricsQueryUnit)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricsQueryMetadata) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


