# LogQueryDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | Pointer to [**LogsQueryCompute**](LogsQueryCompute.md) |  | [optional] 
**GroupBy** | Pointer to [**[]LogQueryDefinitionGroupBy**](LogQueryDefinition_group_by.md) | TODO. | [optional] 
**Index** | Pointer to **string** | A coma separated-list of index names. Use \&quot;*\&quot; query all indexes at once. [Multiple Indexes](https://docs.datadoghq.com/logs/indexes/#multiple-indexes) | [optional] 
**MultiCompute** | Pointer to [**[]LogsQueryCompute**](LogsQueryCompute.md) | This field is mutually exclusive with &#x60;compute&#x60;. | [optional] 
**Search** | Pointer to [**LogQueryDefinitionSearch**](LogQueryDefinition_search.md) |  | [optional] 

## Methods

### NewLogQueryDefinition

`func NewLogQueryDefinition() *LogQueryDefinition`

NewLogQueryDefinition instantiates a new LogQueryDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogQueryDefinitionWithDefaults

`func NewLogQueryDefinitionWithDefaults() *LogQueryDefinition`

NewLogQueryDefinitionWithDefaults instantiates a new LogQueryDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *LogQueryDefinition) GetCompute() LogsQueryCompute`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *LogQueryDefinition) GetComputeOk() (*LogsQueryCompute, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *LogQueryDefinition) SetCompute(v LogsQueryCompute)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *LogQueryDefinition) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### GetGroupBy

`func (o *LogQueryDefinition) GetGroupBy() []LogQueryDefinitionGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *LogQueryDefinition) GetGroupByOk() (*[]LogQueryDefinitionGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *LogQueryDefinition) SetGroupBy(v []LogQueryDefinitionGroupBy)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *LogQueryDefinition) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetIndex

`func (o *LogQueryDefinition) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *LogQueryDefinition) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *LogQueryDefinition) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *LogQueryDefinition) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetMultiCompute

`func (o *LogQueryDefinition) GetMultiCompute() []LogsQueryCompute`

GetMultiCompute returns the MultiCompute field if non-nil, zero value otherwise.

### GetMultiComputeOk

`func (o *LogQueryDefinition) GetMultiComputeOk() (*[]LogsQueryCompute, bool)`

GetMultiComputeOk returns a tuple with the MultiCompute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiCompute

`func (o *LogQueryDefinition) SetMultiCompute(v []LogsQueryCompute)`

SetMultiCompute sets MultiCompute field to given value.

### HasMultiCompute

`func (o *LogQueryDefinition) HasMultiCompute() bool`

HasMultiCompute returns a boolean if a field has been set.

### GetSearch

`func (o *LogQueryDefinition) GetSearch() LogQueryDefinitionSearch`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *LogQueryDefinition) GetSearchOk() (*LogQueryDefinitionSearch, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *LogQueryDefinition) SetSearch(v LogQueryDefinitionSearch)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *LogQueryDefinition) HasSearch() bool`

HasSearch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


