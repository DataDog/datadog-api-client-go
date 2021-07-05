# MetricsPayload

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Series** | [**[]Series**](Series.md) | A list of time series to submit to Datadog. | 

## Methods

### NewMetricsPayload

`func NewMetricsPayload(series []Series) *MetricsPayload`

NewMetricsPayload instantiates a new MetricsPayload object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMetricsPayloadWithDefaults

`func NewMetricsPayloadWithDefaults() *MetricsPayload`

NewMetricsPayloadWithDefaults instantiates a new MetricsPayload object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetSeries

`func (o *MetricsPayload) GetSeries() []Series`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *MetricsPayload) GetSeriesOk() (*[]Series, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *MetricsPayload) SetSeries(v []Series)`

SetSeries sets Series field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


