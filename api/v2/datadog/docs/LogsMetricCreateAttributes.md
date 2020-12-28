# LogsMetricCreateAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | [**LogsMetricCompute**](LogsMetricCompute.md) |  | 
**Filter** | Pointer to [**LogsMetricFilter**](LogsMetricFilter.md) |  | [optional] 
**GroupBy** | Pointer to [**[]LogsMetricGroupBy**](LogsMetricGroupBy.md) | The rules for the group by. | [optional] 

## Methods

### NewLogsMetricCreateAttributes

`func NewLogsMetricCreateAttributes(compute LogsMetricCompute, ) *LogsMetricCreateAttributes`

NewLogsMetricCreateAttributes instantiates a new LogsMetricCreateAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsMetricCreateAttributesWithDefaults

`func NewLogsMetricCreateAttributesWithDefaults() *LogsMetricCreateAttributes`

NewLogsMetricCreateAttributesWithDefaults instantiates a new LogsMetricCreateAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *LogsMetricCreateAttributes) GetCompute() LogsMetricCompute`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *LogsMetricCreateAttributes) GetComputeOk() (*LogsMetricCompute, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *LogsMetricCreateAttributes) SetCompute(v LogsMetricCompute)`

SetCompute sets Compute field to given value.


### GetFilter

`func (o *LogsMetricCreateAttributes) GetFilter() LogsMetricFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LogsMetricCreateAttributes) GetFilterOk() (*LogsMetricFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LogsMetricCreateAttributes) SetFilter(v LogsMetricFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *LogsMetricCreateAttributes) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetGroupBy

`func (o *LogsMetricCreateAttributes) GetGroupBy() []LogsMetricGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *LogsMetricCreateAttributes) GetGroupByOk() (*[]LogsMetricGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *LogsMetricCreateAttributes) SetGroupBy(v []LogsMetricGroupBy)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *LogsMetricCreateAttributes) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


