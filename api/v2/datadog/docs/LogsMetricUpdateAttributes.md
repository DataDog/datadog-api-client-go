# LogsMetricUpdateAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Filter** | Pointer to [**LogsMetricFilter**](LogsMetricFilter.md) |  | [optional] 
**GroupBy** | Pointer to [**[]LogsMetricGroupBy**](LogsMetricGroupBy.md) | The rules for the group by. | [optional] 

## Methods

### NewLogsMetricUpdateAttributes

`func NewLogsMetricUpdateAttributes() *LogsMetricUpdateAttributes`

NewLogsMetricUpdateAttributes instantiates a new LogsMetricUpdateAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsMetricUpdateAttributesWithDefaults

`func NewLogsMetricUpdateAttributesWithDefaults() *LogsMetricUpdateAttributes`

NewLogsMetricUpdateAttributesWithDefaults instantiates a new LogsMetricUpdateAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *LogsMetricUpdateAttributes) GetFilter() LogsMetricFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LogsMetricUpdateAttributes) GetFilterOk() (*LogsMetricFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LogsMetricUpdateAttributes) SetFilter(v LogsMetricFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *LogsMetricUpdateAttributes) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetGroupBy

`func (o *LogsMetricUpdateAttributes) GetGroupBy() []LogsMetricGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *LogsMetricUpdateAttributes) GetGroupByOk() (*[]LogsMetricGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *LogsMetricUpdateAttributes) SetGroupBy(v []LogsMetricGroupBy)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *LogsMetricUpdateAttributes) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


