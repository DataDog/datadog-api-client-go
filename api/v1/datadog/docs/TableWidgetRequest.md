# TableWidgetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregator** | Pointer to [**WidgetAggregator**](WidgetAggregator.md) |  | [optional] 
**Alias** | Pointer to **string** | The column name (defaults to the metric name) | [optional] 
**ApmQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**ConditionalFormats** | Pointer to [**[]WidgetConditionalFormat**](WidgetConditionalFormat.md) | TODO. | [optional] 
**EventQuery** | Pointer to [**EventQueryDefinition**](EventQueryDefinition.md) |  | [optional] 
**Limit** | Pointer to **int64** | For metric queries, the number of lines to show in the table. Only one request should have this property. | [optional] 
**LogQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**NetworkQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**Order** | Pointer to [**WidgetSort**](WidgetSort.md) |  | [optional] 
**ProcessQuery** | Pointer to [**ProcessQueryDefinition**](ProcessQueryDefinition.md) |  | [optional] 
**Q** | Pointer to **string** | TODO. | [optional] 
**RumQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 

## Methods

### NewTableWidgetRequest

`func NewTableWidgetRequest() *TableWidgetRequest`

NewTableWidgetRequest instantiates a new TableWidgetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableWidgetRequestWithDefaults

`func NewTableWidgetRequestWithDefaults() *TableWidgetRequest`

NewTableWidgetRequestWithDefaults instantiates a new TableWidgetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregator

`func (o *TableWidgetRequest) GetAggregator() WidgetAggregator`

GetAggregator returns the Aggregator field if non-nil, zero value otherwise.

### GetAggregatorOk

`func (o *TableWidgetRequest) GetAggregatorOk() (*WidgetAggregator, bool)`

GetAggregatorOk returns a tuple with the Aggregator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregator

`func (o *TableWidgetRequest) SetAggregator(v WidgetAggregator)`

SetAggregator sets Aggregator field to given value.

### HasAggregator

`func (o *TableWidgetRequest) HasAggregator() bool`

HasAggregator returns a boolean if a field has been set.

### GetAlias

`func (o *TableWidgetRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *TableWidgetRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *TableWidgetRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *TableWidgetRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetApmQuery

`func (o *TableWidgetRequest) GetApmQuery() LogQueryDefinition`

GetApmQuery returns the ApmQuery field if non-nil, zero value otherwise.

### GetApmQueryOk

`func (o *TableWidgetRequest) GetApmQueryOk() (*LogQueryDefinition, bool)`

GetApmQueryOk returns a tuple with the ApmQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmQuery

`func (o *TableWidgetRequest) SetApmQuery(v LogQueryDefinition)`

SetApmQuery sets ApmQuery field to given value.

### HasApmQuery

`func (o *TableWidgetRequest) HasApmQuery() bool`

HasApmQuery returns a boolean if a field has been set.

### GetConditionalFormats

`func (o *TableWidgetRequest) GetConditionalFormats() []WidgetConditionalFormat`

GetConditionalFormats returns the ConditionalFormats field if non-nil, zero value otherwise.

### GetConditionalFormatsOk

`func (o *TableWidgetRequest) GetConditionalFormatsOk() (*[]WidgetConditionalFormat, bool)`

GetConditionalFormatsOk returns a tuple with the ConditionalFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalFormats

`func (o *TableWidgetRequest) SetConditionalFormats(v []WidgetConditionalFormat)`

SetConditionalFormats sets ConditionalFormats field to given value.

### HasConditionalFormats

`func (o *TableWidgetRequest) HasConditionalFormats() bool`

HasConditionalFormats returns a boolean if a field has been set.

### GetEventQuery

`func (o *TableWidgetRequest) GetEventQuery() EventQueryDefinition`

GetEventQuery returns the EventQuery field if non-nil, zero value otherwise.

### GetEventQueryOk

`func (o *TableWidgetRequest) GetEventQueryOk() (*EventQueryDefinition, bool)`

GetEventQueryOk returns a tuple with the EventQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventQuery

`func (o *TableWidgetRequest) SetEventQuery(v EventQueryDefinition)`

SetEventQuery sets EventQuery field to given value.

### HasEventQuery

`func (o *TableWidgetRequest) HasEventQuery() bool`

HasEventQuery returns a boolean if a field has been set.

### GetLimit

`func (o *TableWidgetRequest) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TableWidgetRequest) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TableWidgetRequest) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TableWidgetRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLogQuery

`func (o *TableWidgetRequest) GetLogQuery() LogQueryDefinition`

GetLogQuery returns the LogQuery field if non-nil, zero value otherwise.

### GetLogQueryOk

`func (o *TableWidgetRequest) GetLogQueryOk() (*LogQueryDefinition, bool)`

GetLogQueryOk returns a tuple with the LogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQuery

`func (o *TableWidgetRequest) SetLogQuery(v LogQueryDefinition)`

SetLogQuery sets LogQuery field to given value.

### HasLogQuery

`func (o *TableWidgetRequest) HasLogQuery() bool`

HasLogQuery returns a boolean if a field has been set.

### GetNetworkQuery

`func (o *TableWidgetRequest) GetNetworkQuery() LogQueryDefinition`

GetNetworkQuery returns the NetworkQuery field if non-nil, zero value otherwise.

### GetNetworkQueryOk

`func (o *TableWidgetRequest) GetNetworkQueryOk() (*LogQueryDefinition, bool)`

GetNetworkQueryOk returns a tuple with the NetworkQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkQuery

`func (o *TableWidgetRequest) SetNetworkQuery(v LogQueryDefinition)`

SetNetworkQuery sets NetworkQuery field to given value.

### HasNetworkQuery

`func (o *TableWidgetRequest) HasNetworkQuery() bool`

HasNetworkQuery returns a boolean if a field has been set.

### GetOrder

`func (o *TableWidgetRequest) GetOrder() WidgetSort`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TableWidgetRequest) GetOrderOk() (*WidgetSort, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TableWidgetRequest) SetOrder(v WidgetSort)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TableWidgetRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetProcessQuery

`func (o *TableWidgetRequest) GetProcessQuery() ProcessQueryDefinition`

GetProcessQuery returns the ProcessQuery field if non-nil, zero value otherwise.

### GetProcessQueryOk

`func (o *TableWidgetRequest) GetProcessQueryOk() (*ProcessQueryDefinition, bool)`

GetProcessQueryOk returns a tuple with the ProcessQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessQuery

`func (o *TableWidgetRequest) SetProcessQuery(v ProcessQueryDefinition)`

SetProcessQuery sets ProcessQuery field to given value.

### HasProcessQuery

`func (o *TableWidgetRequest) HasProcessQuery() bool`

HasProcessQuery returns a boolean if a field has been set.

### GetQ

`func (o *TableWidgetRequest) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *TableWidgetRequest) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *TableWidgetRequest) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *TableWidgetRequest) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetRumQuery

`func (o *TableWidgetRequest) GetRumQuery() LogQueryDefinition`

GetRumQuery returns the RumQuery field if non-nil, zero value otherwise.

### GetRumQueryOk

`func (o *TableWidgetRequest) GetRumQueryOk() (*LogQueryDefinition, bool)`

GetRumQueryOk returns a tuple with the RumQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumQuery

`func (o *TableWidgetRequest) SetRumQuery(v LogQueryDefinition)`

SetRumQuery sets RumQuery field to given value.

### HasRumQuery

`func (o *TableWidgetRequest) HasRumQuery() bool`

HasRumQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


