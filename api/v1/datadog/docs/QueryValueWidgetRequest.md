# QueryValueWidgetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregator** | Pointer to [**WidgetAggregator**](WidgetAggregator.md) |  | [optional] 
**ApmQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**ConditionalFormats** | Pointer to [**[]WidgetConditionalFormat**](WidgetConditionalFormat.md) |  | [optional] 
**EventQuery** | Pointer to [**EventQueryDefinition**](EventQueryDefinition.md) |  | [optional] 
**LogQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**NetworkQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**ProcessQuery** | Pointer to [**ProcessQueryDefinition**](ProcessQueryDefinition.md) |  | [optional] 
**Q** | Pointer to **string** |  | [optional] 
**RumQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 

## Methods

### NewQueryValueWidgetRequest

`func NewQueryValueWidgetRequest() *QueryValueWidgetRequest`

NewQueryValueWidgetRequest instantiates a new QueryValueWidgetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryValueWidgetRequestWithDefaults

`func NewQueryValueWidgetRequestWithDefaults() *QueryValueWidgetRequest`

NewQueryValueWidgetRequestWithDefaults instantiates a new QueryValueWidgetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregator

`func (o *QueryValueWidgetRequest) GetAggregator() WidgetAggregator`

GetAggregator returns the Aggregator field if non-nil, zero value otherwise.

### GetAggregatorOk

`func (o *QueryValueWidgetRequest) GetAggregatorOk() (WidgetAggregator, bool)`

GetAggregatorOk returns a tuple with the Aggregator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAggregator

`func (o *QueryValueWidgetRequest) HasAggregator() bool`

HasAggregator returns a boolean if a field has been set.

### SetAggregator

`func (o *QueryValueWidgetRequest) SetAggregator(v WidgetAggregator)`

SetAggregator gets a reference to the given WidgetAggregator and assigns it to the Aggregator field.

### GetApmQuery

`func (o *QueryValueWidgetRequest) GetApmQuery() LogQueryDefinition`

GetApmQuery returns the ApmQuery field if non-nil, zero value otherwise.

### GetApmQueryOk

`func (o *QueryValueWidgetRequest) GetApmQueryOk() (LogQueryDefinition, bool)`

GetApmQueryOk returns a tuple with the ApmQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApmQuery

`func (o *QueryValueWidgetRequest) HasApmQuery() bool`

HasApmQuery returns a boolean if a field has been set.

### SetApmQuery

`func (o *QueryValueWidgetRequest) SetApmQuery(v LogQueryDefinition)`

SetApmQuery gets a reference to the given LogQueryDefinition and assigns it to the ApmQuery field.

### GetConditionalFormats

`func (o *QueryValueWidgetRequest) GetConditionalFormats() []WidgetConditionalFormat`

GetConditionalFormats returns the ConditionalFormats field if non-nil, zero value otherwise.

### GetConditionalFormatsOk

`func (o *QueryValueWidgetRequest) GetConditionalFormatsOk() ([]WidgetConditionalFormat, bool)`

GetConditionalFormatsOk returns a tuple with the ConditionalFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConditionalFormats

`func (o *QueryValueWidgetRequest) HasConditionalFormats() bool`

HasConditionalFormats returns a boolean if a field has been set.

### SetConditionalFormats

`func (o *QueryValueWidgetRequest) SetConditionalFormats(v []WidgetConditionalFormat)`

SetConditionalFormats gets a reference to the given []WidgetConditionalFormat and assigns it to the ConditionalFormats field.

### GetEventQuery

`func (o *QueryValueWidgetRequest) GetEventQuery() EventQueryDefinition`

GetEventQuery returns the EventQuery field if non-nil, zero value otherwise.

### GetEventQueryOk

`func (o *QueryValueWidgetRequest) GetEventQueryOk() (EventQueryDefinition, bool)`

GetEventQueryOk returns a tuple with the EventQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventQuery

`func (o *QueryValueWidgetRequest) HasEventQuery() bool`

HasEventQuery returns a boolean if a field has been set.

### SetEventQuery

`func (o *QueryValueWidgetRequest) SetEventQuery(v EventQueryDefinition)`

SetEventQuery gets a reference to the given EventQueryDefinition and assigns it to the EventQuery field.

### GetLogQuery

`func (o *QueryValueWidgetRequest) GetLogQuery() LogQueryDefinition`

GetLogQuery returns the LogQuery field if non-nil, zero value otherwise.

### GetLogQueryOk

`func (o *QueryValueWidgetRequest) GetLogQueryOk() (LogQueryDefinition, bool)`

GetLogQueryOk returns a tuple with the LogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLogQuery

`func (o *QueryValueWidgetRequest) HasLogQuery() bool`

HasLogQuery returns a boolean if a field has been set.

### SetLogQuery

`func (o *QueryValueWidgetRequest) SetLogQuery(v LogQueryDefinition)`

SetLogQuery gets a reference to the given LogQueryDefinition and assigns it to the LogQuery field.

### GetNetworkQuery

`func (o *QueryValueWidgetRequest) GetNetworkQuery() LogQueryDefinition`

GetNetworkQuery returns the NetworkQuery field if non-nil, zero value otherwise.

### GetNetworkQueryOk

`func (o *QueryValueWidgetRequest) GetNetworkQueryOk() (LogQueryDefinition, bool)`

GetNetworkQueryOk returns a tuple with the NetworkQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetworkQuery

`func (o *QueryValueWidgetRequest) HasNetworkQuery() bool`

HasNetworkQuery returns a boolean if a field has been set.

### SetNetworkQuery

`func (o *QueryValueWidgetRequest) SetNetworkQuery(v LogQueryDefinition)`

SetNetworkQuery gets a reference to the given LogQueryDefinition and assigns it to the NetworkQuery field.

### GetProcessQuery

`func (o *QueryValueWidgetRequest) GetProcessQuery() ProcessQueryDefinition`

GetProcessQuery returns the ProcessQuery field if non-nil, zero value otherwise.

### GetProcessQueryOk

`func (o *QueryValueWidgetRequest) GetProcessQueryOk() (ProcessQueryDefinition, bool)`

GetProcessQueryOk returns a tuple with the ProcessQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProcessQuery

`func (o *QueryValueWidgetRequest) HasProcessQuery() bool`

HasProcessQuery returns a boolean if a field has been set.

### SetProcessQuery

`func (o *QueryValueWidgetRequest) SetProcessQuery(v ProcessQueryDefinition)`

SetProcessQuery gets a reference to the given ProcessQueryDefinition and assigns it to the ProcessQuery field.

### GetQ

`func (o *QueryValueWidgetRequest) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *QueryValueWidgetRequest) GetQOk() (string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQ

`func (o *QueryValueWidgetRequest) HasQ() bool`

HasQ returns a boolean if a field has been set.

### SetQ

`func (o *QueryValueWidgetRequest) SetQ(v string)`

SetQ gets a reference to the given string and assigns it to the Q field.

### GetRumQuery

`func (o *QueryValueWidgetRequest) GetRumQuery() LogQueryDefinition`

GetRumQuery returns the RumQuery field if non-nil, zero value otherwise.

### GetRumQueryOk

`func (o *QueryValueWidgetRequest) GetRumQueryOk() (LogQueryDefinition, bool)`

GetRumQueryOk returns a tuple with the RumQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRumQuery

`func (o *QueryValueWidgetRequest) HasRumQuery() bool`

HasRumQuery returns a boolean if a field has been set.

### SetRumQuery

`func (o *QueryValueWidgetRequest) SetRumQuery(v LogQueryDefinition)`

SetRumQuery gets a reference to the given LogQueryDefinition and assigns it to the RumQuery field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


