# SLOHistoryMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Denominator** | Pointer to [**SLOHistoryMetricsSeries**](SLOHistoryMetricsSeries.md) |  | 
**Interval** | Pointer to **int64** | The aggregated query interval for the series data. It&#39;s implicit based on the query time window. | 
**Message** | Pointer to **string** | Optional message if there are specific query issues/warnings. | [optional] 
**Numerator** | Pointer to [**SLOHistoryMetricsSeries**](SLOHistoryMetricsSeries.md) |  | 
**Query** | Pointer to **string** | The combined numerator and denominator query CSV. | 
**ResType** | Pointer to **string** | The series result type. This mimics &#x60;batch_query&#x60; response type. | 
**RespVersion** | Pointer to **int64** | The series response version type. This mimics &#x60;batch_query&#x60; response type. | 
**Times** | Pointer to **[]float64** | An array of query timestamps in EPOCH milliseconds | 

## Methods

### NewSLOHistoryMetrics

`func NewSLOHistoryMetrics(denominator SLOHistoryMetricsSeries, interval int64, numerator SLOHistoryMetricsSeries, query string, resType string, respVersion int64, times []float64, ) *SLOHistoryMetrics`

NewSLOHistoryMetrics instantiates a new SLOHistoryMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLOHistoryMetricsWithDefaults

`func NewSLOHistoryMetricsWithDefaults() *SLOHistoryMetrics`

NewSLOHistoryMetricsWithDefaults instantiates a new SLOHistoryMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDenominator

`func (o *SLOHistoryMetrics) GetDenominator() SLOHistoryMetricsSeries`

GetDenominator returns the Denominator field if non-nil, zero value otherwise.

### GetDenominatorOk

`func (o *SLOHistoryMetrics) GetDenominatorOk() (*SLOHistoryMetricsSeries, bool)`

GetDenominatorOk returns a tuple with the Denominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominator

`func (o *SLOHistoryMetrics) SetDenominator(v SLOHistoryMetricsSeries)`

SetDenominator sets Denominator field to given value.


### GetInterval

`func (o *SLOHistoryMetrics) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *SLOHistoryMetrics) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *SLOHistoryMetrics) SetInterval(v int64)`

SetInterval sets Interval field to given value.


### GetMessage

`func (o *SLOHistoryMetrics) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SLOHistoryMetrics) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SLOHistoryMetrics) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SLOHistoryMetrics) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNumerator

`func (o *SLOHistoryMetrics) GetNumerator() SLOHistoryMetricsSeries`

GetNumerator returns the Numerator field if non-nil, zero value otherwise.

### GetNumeratorOk

`func (o *SLOHistoryMetrics) GetNumeratorOk() (*SLOHistoryMetricsSeries, bool)`

GetNumeratorOk returns a tuple with the Numerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerator

`func (o *SLOHistoryMetrics) SetNumerator(v SLOHistoryMetricsSeries)`

SetNumerator sets Numerator field to given value.


### GetQuery

`func (o *SLOHistoryMetrics) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SLOHistoryMetrics) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SLOHistoryMetrics) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetResType

`func (o *SLOHistoryMetrics) GetResType() string`

GetResType returns the ResType field if non-nil, zero value otherwise.

### GetResTypeOk

`func (o *SLOHistoryMetrics) GetResTypeOk() (*string, bool)`

GetResTypeOk returns a tuple with the ResType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResType

`func (o *SLOHistoryMetrics) SetResType(v string)`

SetResType sets ResType field to given value.


### GetRespVersion

`func (o *SLOHistoryMetrics) GetRespVersion() int64`

GetRespVersion returns the RespVersion field if non-nil, zero value otherwise.

### GetRespVersionOk

`func (o *SLOHistoryMetrics) GetRespVersionOk() (*int64, bool)`

GetRespVersionOk returns a tuple with the RespVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespVersion

`func (o *SLOHistoryMetrics) SetRespVersion(v int64)`

SetRespVersion sets RespVersion field to given value.


### GetTimes

`func (o *SLOHistoryMetrics) GetTimes() []float64`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *SLOHistoryMetrics) GetTimesOk() (*[]float64, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *SLOHistoryMetrics) SetTimes(v []float64)`

SetTimes sets Times field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


