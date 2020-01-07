# UsageTimeseriesHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hour** | Pointer to [**time.Time**](time.Time.md) | The hour for the usage. | [optional] 
**NumCustomTimeseries** | Pointer to **int64** | Contains the number of distinct custom metrics. | [optional] 

## Methods

### GetHour

`func (o *UsageTimeseriesHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageTimeseriesHour) GetHourOk() (time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHour

`func (o *UsageTimeseriesHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### SetHour

`func (o *UsageTimeseriesHour) SetHour(v time.Time)`

SetHour gets a reference to the given time.Time and assigns it to the Hour field.

### GetNumCustomTimeseries

`func (o *UsageTimeseriesHour) GetNumCustomTimeseries() int64`

GetNumCustomTimeseries returns the NumCustomTimeseries field if non-nil, zero value otherwise.

### GetNumCustomTimeseriesOk

`func (o *UsageTimeseriesHour) GetNumCustomTimeseriesOk() (int64, bool)`

GetNumCustomTimeseriesOk returns a tuple with the NumCustomTimeseries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNumCustomTimeseries

`func (o *UsageTimeseriesHour) HasNumCustomTimeseries() bool`

HasNumCustomTimeseries returns a boolean if a field has been set.

### SetNumCustomTimeseries

`func (o *UsageTimeseriesHour) SetNumCustomTimeseries(v int64)`

SetNumCustomTimeseries gets a reference to the given int64 and assigns it to the NumCustomTimeseries field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


