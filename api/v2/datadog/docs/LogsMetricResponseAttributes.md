# LogsMetricResponseAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Compute** | Pointer to [**LogsMetricResponseCompute**](LogsMetricResponseCompute.md) |  | [optional] 
**Filter** | Pointer to [**LogsMetricResponseFilter**](LogsMetricResponseFilter.md) |  | [optional] 
**GroupBy** | Pointer to [**[]LogsMetricResponseGroupBy**](LogsMetricResponseGroupBy.md) | The rules for the group by. | [optional] 

## Methods

### NewLogsMetricResponseAttributes

`func NewLogsMetricResponseAttributes() *LogsMetricResponseAttributes`

NewLogsMetricResponseAttributes instantiates a new LogsMetricResponseAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsMetricResponseAttributesWithDefaults

`func NewLogsMetricResponseAttributesWithDefaults() *LogsMetricResponseAttributes`

NewLogsMetricResponseAttributesWithDefaults instantiates a new LogsMetricResponseAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCompute

`func (o *LogsMetricResponseAttributes) GetCompute() LogsMetricResponseCompute`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *LogsMetricResponseAttributes) GetComputeOk() (*LogsMetricResponseCompute, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *LogsMetricResponseAttributes) SetCompute(v LogsMetricResponseCompute)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *LogsMetricResponseAttributes) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### GetFilter

`func (o *LogsMetricResponseAttributes) GetFilter() LogsMetricResponseFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LogsMetricResponseAttributes) GetFilterOk() (*LogsMetricResponseFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LogsMetricResponseAttributes) SetFilter(v LogsMetricResponseFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *LogsMetricResponseAttributes) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetGroupBy

`func (o *LogsMetricResponseAttributes) GetGroupBy() []LogsMetricResponseGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *LogsMetricResponseAttributes) GetGroupByOk() (*[]LogsMetricResponseGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *LogsMetricResponseAttributes) SetGroupBy(v []LogsMetricResponseGroupBy)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *LogsMetricResponseAttributes) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


