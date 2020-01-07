# Series

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | The name of the host that produced the metric. | [optional] 
**Interval** | Pointer to **NullableInt32** | If the type of the metric is rate or count, define the corresponding interval. | [optional] 
**Metric** | Pointer to **string** | The name of the timeseries | 
**Points** | Pointer to [**[][]float64**](array.md) |  | 
**Tags** | Pointer to **[]string** | A list of tags associated with the metric. | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to gauge]

## Methods

### GetHost

`func (o *Series) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Series) GetHostOk() (string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHost

`func (o *Series) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHost

`func (o *Series) SetHost(v string)`

SetHost gets a reference to the given string and assigns it to the Host field.

### GetInterval

`func (o *Series) GetInterval() NullableInt32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Series) GetIntervalOk() (NullableInt32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInterval

`func (o *Series) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetInterval

`func (o *Series) SetInterval(v NullableInt32)`

SetInterval gets a reference to the given NullableInt32 and assigns it to the Interval field.

### SetIntervalExplicitNull

`func (o *Series) SetIntervalExplicitNull(b bool)`

SetIntervalExplicitNull (un)sets Interval to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Interval value is set to nil even if false is passed
### GetMetric

`func (o *Series) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *Series) GetMetricOk() (string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMetric

`func (o *Series) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### SetMetric

`func (o *Series) SetMetric(v string)`

SetMetric gets a reference to the given string and assigns it to the Metric field.

### GetPoints

`func (o *Series) GetPoints() [][]float64`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *Series) GetPointsOk() ([][]float64, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPoints

`func (o *Series) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### SetPoints

`func (o *Series) SetPoints(v [][]float64)`

SetPoints gets a reference to the given [][]float64 and assigns it to the Points field.

### GetTags

`func (o *Series) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Series) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *Series) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *Series) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.

### GetType

`func (o *Series) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Series) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *Series) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *Series) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


