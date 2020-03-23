# HistoryServiceLevelObjectiveMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Denominator** | Pointer to [**HistoryServiceLevelObjectiveMetricsSeries**](HistoryServiceLevelObjectiveMetricsSeries.md) |  | 
**Interval** | Pointer to **int64** | The aggregated query interval for the series data. It&#39;s implicit based on the query time window. | 
**Message** | Pointer to **string** | Optional message if there are specific query issues/warnings. | [optional] 
**Numerator** | Pointer to [**HistoryServiceLevelObjectiveMetricsSeries**](HistoryServiceLevelObjectiveMetricsSeries.md) |  | 
**Query** | Pointer to **string** | The combined numerator &amp;&amp; denominator query CSV. | 
**ResType** | Pointer to **string** | The series result type. This mimics &#x60;batch_query&#x60; response type. | 
**RespVersion** | Pointer to **int64** | The series response version type. This mimics &#x60;batch_query&#x60; response type. | 
**Times** | Pointer to **[]float64** | The query timestamps in epoch milliseconds | 

## Methods

### NewHistoryServiceLevelObjectiveMetrics

`func NewHistoryServiceLevelObjectiveMetrics(denominator HistoryServiceLevelObjectiveMetricsSeries, interval int64, numerator HistoryServiceLevelObjectiveMetricsSeries, query string, resType string, respVersion int64, times []float64, ) *HistoryServiceLevelObjectiveMetrics`

NewHistoryServiceLevelObjectiveMetrics instantiates a new HistoryServiceLevelObjectiveMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryServiceLevelObjectiveMetricsWithDefaults

`func NewHistoryServiceLevelObjectiveMetricsWithDefaults() *HistoryServiceLevelObjectiveMetrics`

NewHistoryServiceLevelObjectiveMetricsWithDefaults instantiates a new HistoryServiceLevelObjectiveMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDenominator

`func (o *HistoryServiceLevelObjectiveMetrics) GetDenominator() HistoryServiceLevelObjectiveMetricsSeries`

GetDenominator returns the Denominator field if non-nil, zero value otherwise.

### GetDenominatorOk

`func (o *HistoryServiceLevelObjectiveMetrics) GetDenominatorOk() (HistoryServiceLevelObjectiveMetricsSeries, bool)`

GetDenominatorOk returns a tuple with the Denominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDenominator

`func (o *HistoryServiceLevelObjectiveMetrics) HasDenominator() bool`

HasDenominator returns a boolean if a field has been set.

### SetDenominator

`func (o *HistoryServiceLevelObjectiveMetrics) SetDenominator(v HistoryServiceLevelObjectiveMetricsSeries)`

SetDenominator gets a reference to the given HistoryServiceLevelObjectiveMetricsSeries and assigns it to the Denominator field.

### GetInterval

`func (o *HistoryServiceLevelObjectiveMetrics) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *HistoryServiceLevelObjectiveMetrics) GetIntervalOk() (int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInterval

`func (o *HistoryServiceLevelObjectiveMetrics) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetInterval

`func (o *HistoryServiceLevelObjectiveMetrics) SetInterval(v int64)`

SetInterval gets a reference to the given int64 and assigns it to the Interval field.

### GetMessage

`func (o *HistoryServiceLevelObjectiveMetrics) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HistoryServiceLevelObjectiveMetrics) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *HistoryServiceLevelObjectiveMetrics) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *HistoryServiceLevelObjectiveMetrics) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.

### GetNumerator

`func (o *HistoryServiceLevelObjectiveMetrics) GetNumerator() HistoryServiceLevelObjectiveMetricsSeries`

GetNumerator returns the Numerator field if non-nil, zero value otherwise.

### GetNumeratorOk

`func (o *HistoryServiceLevelObjectiveMetrics) GetNumeratorOk() (HistoryServiceLevelObjectiveMetricsSeries, bool)`

GetNumeratorOk returns a tuple with the Numerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNumerator

`func (o *HistoryServiceLevelObjectiveMetrics) HasNumerator() bool`

HasNumerator returns a boolean if a field has been set.

### SetNumerator

`func (o *HistoryServiceLevelObjectiveMetrics) SetNumerator(v HistoryServiceLevelObjectiveMetricsSeries)`

SetNumerator gets a reference to the given HistoryServiceLevelObjectiveMetricsSeries and assigns it to the Numerator field.

### GetQuery

`func (o *HistoryServiceLevelObjectiveMetrics) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *HistoryServiceLevelObjectiveMetrics) GetQueryOk() (string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQuery

`func (o *HistoryServiceLevelObjectiveMetrics) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQuery

`func (o *HistoryServiceLevelObjectiveMetrics) SetQuery(v string)`

SetQuery gets a reference to the given string and assigns it to the Query field.

### GetResType

`func (o *HistoryServiceLevelObjectiveMetrics) GetResType() string`

GetResType returns the ResType field if non-nil, zero value otherwise.

### GetResTypeOk

`func (o *HistoryServiceLevelObjectiveMetrics) GetResTypeOk() (string, bool)`

GetResTypeOk returns a tuple with the ResType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResType

`func (o *HistoryServiceLevelObjectiveMetrics) HasResType() bool`

HasResType returns a boolean if a field has been set.

### SetResType

`func (o *HistoryServiceLevelObjectiveMetrics) SetResType(v string)`

SetResType gets a reference to the given string and assigns it to the ResType field.

### GetRespVersion

`func (o *HistoryServiceLevelObjectiveMetrics) GetRespVersion() int64`

GetRespVersion returns the RespVersion field if non-nil, zero value otherwise.

### GetRespVersionOk

`func (o *HistoryServiceLevelObjectiveMetrics) GetRespVersionOk() (int64, bool)`

GetRespVersionOk returns a tuple with the RespVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRespVersion

`func (o *HistoryServiceLevelObjectiveMetrics) HasRespVersion() bool`

HasRespVersion returns a boolean if a field has been set.

### SetRespVersion

`func (o *HistoryServiceLevelObjectiveMetrics) SetRespVersion(v int64)`

SetRespVersion gets a reference to the given int64 and assigns it to the RespVersion field.

### GetTimes

`func (o *HistoryServiceLevelObjectiveMetrics) GetTimes() []float64`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *HistoryServiceLevelObjectiveMetrics) GetTimesOk() ([]float64, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimes

`func (o *HistoryServiceLevelObjectiveMetrics) HasTimes() bool`

HasTimes returns a boolean if a field has been set.

### SetTimes

`func (o *HistoryServiceLevelObjectiveMetrics) SetTimes(v []float64)`

SetTimes gets a reference to the given []float64 and assigns it to the Times field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


