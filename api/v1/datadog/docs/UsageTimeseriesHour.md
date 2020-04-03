# UsageTimeseriesHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hour** | Pointer to [**time.Time**](time.Time.md) | The hour for the usage. | [optional] 
**NumCustomInputTimeseries** | Pointer to **int64** | Contains the number of custom metrics that are inputs for aggregations (metric configured is custom). | [optional] 
**NumCustomOutputTimeseries** | Pointer to **int64** | Contains the number of custom metrics that are outputs for aggregations (metric configured is custom). | [optional] 
**NumCustomTimeseries** | Pointer to **int64** | Contains the number of non-aggregation custom metrics. | [optional] 

## Methods

### NewUsageTimeseriesHour

`func NewUsageTimeseriesHour() *UsageTimeseriesHour`

NewUsageTimeseriesHour instantiates a new UsageTimeseriesHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageTimeseriesHourWithDefaults

`func NewUsageTimeseriesHourWithDefaults() *UsageTimeseriesHour`

NewUsageTimeseriesHourWithDefaults instantiates a new UsageTimeseriesHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHour

`func (o *UsageTimeseriesHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageTimeseriesHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageTimeseriesHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageTimeseriesHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetNumCustomInputTimeseries

`func (o *UsageTimeseriesHour) GetNumCustomInputTimeseries() int64`

GetNumCustomInputTimeseries returns the NumCustomInputTimeseries field if non-nil, zero value otherwise.

### GetNumCustomInputTimeseriesOk

`func (o *UsageTimeseriesHour) GetNumCustomInputTimeseriesOk() (*int64, bool)`

GetNumCustomInputTimeseriesOk returns a tuple with the NumCustomInputTimeseries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCustomInputTimeseries

`func (o *UsageTimeseriesHour) SetNumCustomInputTimeseries(v int64)`

SetNumCustomInputTimeseries sets NumCustomInputTimeseries field to given value.

### HasNumCustomInputTimeseries

`func (o *UsageTimeseriesHour) HasNumCustomInputTimeseries() bool`

HasNumCustomInputTimeseries returns a boolean if a field has been set.

### GetNumCustomOutputTimeseries

`func (o *UsageTimeseriesHour) GetNumCustomOutputTimeseries() int64`

GetNumCustomOutputTimeseries returns the NumCustomOutputTimeseries field if non-nil, zero value otherwise.

### GetNumCustomOutputTimeseriesOk

`func (o *UsageTimeseriesHour) GetNumCustomOutputTimeseriesOk() (*int64, bool)`

GetNumCustomOutputTimeseriesOk returns a tuple with the NumCustomOutputTimeseries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCustomOutputTimeseries

`func (o *UsageTimeseriesHour) SetNumCustomOutputTimeseries(v int64)`

SetNumCustomOutputTimeseries sets NumCustomOutputTimeseries field to given value.

### HasNumCustomOutputTimeseries

`func (o *UsageTimeseriesHour) HasNumCustomOutputTimeseries() bool`

HasNumCustomOutputTimeseries returns a boolean if a field has been set.

### GetNumCustomTimeseries

`func (o *UsageTimeseriesHour) GetNumCustomTimeseries() int64`

GetNumCustomTimeseries returns the NumCustomTimeseries field if non-nil, zero value otherwise.

### GetNumCustomTimeseriesOk

`func (o *UsageTimeseriesHour) GetNumCustomTimeseriesOk() (*int64, bool)`

GetNumCustomTimeseriesOk returns a tuple with the NumCustomTimeseries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCustomTimeseries

`func (o *UsageTimeseriesHour) SetNumCustomTimeseries(v int64)`

SetNumCustomTimeseries sets NumCustomTimeseries field to given value.

### HasNumCustomTimeseries

`func (o *UsageTimeseriesHour) HasNumCustomTimeseries() bool`

HasNumCustomTimeseries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


