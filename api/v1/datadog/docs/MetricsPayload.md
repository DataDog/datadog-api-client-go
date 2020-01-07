# MetricsPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Series** | Pointer to [**[]Series**](Series.md) | A list of time series to submit to Datadog | [optional] 

## Methods

### GetSeries

`func (o *MetricsPayload) GetSeries() []Series`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *MetricsPayload) GetSeriesOk() ([]Series, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSeries

`func (o *MetricsPayload) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### SetSeries

`func (o *MetricsPayload) SetSeries(v []Series)`

SetSeries gets a reference to the given []Series and assigns it to the Series field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


