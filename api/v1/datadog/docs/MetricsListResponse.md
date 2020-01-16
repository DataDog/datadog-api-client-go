# MetricsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | Time when the metrics where active, seconds since the Unix epoch. | [optional] 
**Metrics** | Pointer to **[]string** | List of metric names | [optional] 

## Methods

### GetFrom

`func (o *MetricsListResponse) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MetricsListResponse) GetFromOk() (string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFrom

`func (o *MetricsListResponse) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFrom

`func (o *MetricsListResponse) SetFrom(v string)`

SetFrom gets a reference to the given string and assigns it to the From field.

### GetMetrics

`func (o *MetricsListResponse) GetMetrics() []string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MetricsListResponse) GetMetricsOk() ([]string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMetrics

`func (o *MetricsListResponse) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### SetMetrics

`func (o *MetricsListResponse) SetMetrics(v []string)`

SetMetrics gets a reference to the given []string and assigns it to the Metrics field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


