# LogsGroupByHistogram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | **float64** | The bin size of the histogram buckets | 
**Max** | **float64** | The maximum value for the measure used in the histogram (values greater than this one are filtered out) | 
**Min** | **float64** | The minimum value for the measure used in the histogram (values smaller than this one are filtered out) | 

## Methods

### NewLogsGroupByHistogram

`func NewLogsGroupByHistogram(interval float64, max float64, min float64, ) *LogsGroupByHistogram`

NewLogsGroupByHistogram instantiates a new LogsGroupByHistogram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsGroupByHistogramWithDefaults

`func NewLogsGroupByHistogramWithDefaults() *LogsGroupByHistogram`

NewLogsGroupByHistogramWithDefaults instantiates a new LogsGroupByHistogram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *LogsGroupByHistogram) GetInterval() float64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *LogsGroupByHistogram) GetIntervalOk() (*float64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *LogsGroupByHistogram) SetInterval(v float64)`

SetInterval sets Interval field to given value.


### GetMax

`func (o *LogsGroupByHistogram) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *LogsGroupByHistogram) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *LogsGroupByHistogram) SetMax(v float64)`

SetMax sets Max field to given value.


### GetMin

`func (o *LogsGroupByHistogram) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *LogsGroupByHistogram) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *LogsGroupByHistogram) SetMin(v float64)`

SetMin sets Min field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


