# LogsAggregateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | Pointer to [**[]LogsCompute**](LogsCompute.md) | The list of metrics or timeseries to compute for the retrieved buckets. | [optional] 
**Filter** | Pointer to [**LogsQueryFilter**](LogsQueryFilter.md) |  | [optional] 
**GroupBy** | Pointer to [**[]LogsGroupBy**](LogsGroupBy.md) | The rules for the group by | [optional] 
**Options** | Pointer to [**LogsQueryOptions**](LogsQueryOptions.md) |  | [optional] 
**Paging** | Pointer to [**LogsAggregateRequestPaging**](LogsAggregateRequest_paging.md) |  | [optional] 

## Methods

### NewLogsAggregateRequest

`func NewLogsAggregateRequest() *LogsAggregateRequest`

NewLogsAggregateRequest instantiates a new LogsAggregateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsAggregateRequestWithDefaults

`func NewLogsAggregateRequestWithDefaults() *LogsAggregateRequest`

NewLogsAggregateRequestWithDefaults instantiates a new LogsAggregateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *LogsAggregateRequest) GetCompute() []LogsCompute`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *LogsAggregateRequest) GetComputeOk() (*[]LogsCompute, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *LogsAggregateRequest) SetCompute(v []LogsCompute)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *LogsAggregateRequest) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### GetFilter

`func (o *LogsAggregateRequest) GetFilter() LogsQueryFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LogsAggregateRequest) GetFilterOk() (*LogsQueryFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LogsAggregateRequest) SetFilter(v LogsQueryFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *LogsAggregateRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetGroupBy

`func (o *LogsAggregateRequest) GetGroupBy() []LogsGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *LogsAggregateRequest) GetGroupByOk() (*[]LogsGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *LogsAggregateRequest) SetGroupBy(v []LogsGroupBy)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *LogsAggregateRequest) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetOptions

`func (o *LogsAggregateRequest) GetOptions() LogsQueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *LogsAggregateRequest) GetOptionsOk() (*LogsQueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *LogsAggregateRequest) SetOptions(v LogsQueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *LogsAggregateRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPaging

`func (o *LogsAggregateRequest) GetPaging() LogsAggregateRequestPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *LogsAggregateRequest) GetPagingOk() (*LogsAggregateRequestPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *LogsAggregateRequest) SetPaging(v LogsAggregateRequestPaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *LogsAggregateRequest) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


