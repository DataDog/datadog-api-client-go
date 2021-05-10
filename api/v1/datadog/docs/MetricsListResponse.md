# MetricsListResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**From** | Pointer to **string** | Time when the metrics were active, seconds since the Unix epoch. | [optional] 
**Metrics** | Pointer to **[]string** | List of metric names. | [optional] 

## Methods

### NewMetricsListResponse

`func NewMetricsListResponse() *MetricsListResponse`

NewMetricsListResponse instantiates a new MetricsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsListResponseWithDefaults

`func NewMetricsListResponseWithDefaults() *MetricsListResponse`

NewMetricsListResponseWithDefaults instantiates a new MetricsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *MetricsListResponse) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MetricsListResponse) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MetricsListResponse) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MetricsListResponse) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetMetrics

`func (o *MetricsListResponse) GetMetrics() []string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MetricsListResponse) GetMetricsOk() (*[]string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MetricsListResponse) SetMetrics(v []string)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *MetricsListResponse) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


