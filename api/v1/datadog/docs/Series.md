# Series

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | The name of the host that produced the metric. | [optional] 
**Interval** | Pointer to **NullableInt64** | If the type of the metric is rate or count, define the corresponding interval. | [optional] 
**Metric** | Pointer to **string** | The name of the timeseries. | 
**Points** | Pointer to [**[][]float64**](array.md) |  | 
**Tags** | Pointer to **[]string** | A list of tags associated with the metric. | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "gauge"]

## Methods

### NewSeries

`func NewSeries(metric string, points [][]float64, ) *Series`

NewSeries instantiates a new Series object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesWithDefaults

`func NewSeriesWithDefaults() *Series`

NewSeriesWithDefaults instantiates a new Series object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *Series) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Series) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Series) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *Series) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetInterval

`func (o *Series) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Series) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Series) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *Series) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *Series) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *Series) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetMetric

`func (o *Series) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *Series) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *Series) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetPoints

`func (o *Series) GetPoints() [][]float64`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *Series) GetPointsOk() (*[][]float64, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *Series) SetPoints(v [][]float64)`

SetPoints sets Points field to given value.


### GetTags

`func (o *Series) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Series) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Series) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Series) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *Series) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Series) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Series) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Series) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


