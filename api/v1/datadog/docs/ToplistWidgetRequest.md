# ToplistWidgetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApmQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**ConditionalFormats** | Pointer to [**[]WidgetConditionalFormat**](WidgetConditionalFormat.md) | TODO. | [optional] 
**EventQuery** | Pointer to [**EventQueryDefinition**](EventQueryDefinition.md) |  | [optional] 
**LogQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**NetworkQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**ProcessQuery** | Pointer to [**ProcessQueryDefinition**](ProcessQueryDefinition.md) |  | [optional] 
**Q** | Pointer to **string** | Widget query. | [optional] 
**RumQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**Style** | Pointer to [**TimeseriesWidgetRequestStyle**](TimeseriesWidgetRequest_style.md) |  | [optional] 

## Methods

### NewToplistWidgetRequest

`func NewToplistWidgetRequest() *ToplistWidgetRequest`

NewToplistWidgetRequest instantiates a new ToplistWidgetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToplistWidgetRequestWithDefaults

`func NewToplistWidgetRequestWithDefaults() *ToplistWidgetRequest`

NewToplistWidgetRequestWithDefaults instantiates a new ToplistWidgetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApmQuery

`func (o *ToplistWidgetRequest) GetApmQuery() LogQueryDefinition`

GetApmQuery returns the ApmQuery field if non-nil, zero value otherwise.

### GetApmQueryOk

`func (o *ToplistWidgetRequest) GetApmQueryOk() (*LogQueryDefinition, bool)`

GetApmQueryOk returns a tuple with the ApmQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmQuery

`func (o *ToplistWidgetRequest) SetApmQuery(v LogQueryDefinition)`

SetApmQuery sets ApmQuery field to given value.

### HasApmQuery

`func (o *ToplistWidgetRequest) HasApmQuery() bool`

HasApmQuery returns a boolean if a field has been set.

### GetConditionalFormats

`func (o *ToplistWidgetRequest) GetConditionalFormats() []WidgetConditionalFormat`

GetConditionalFormats returns the ConditionalFormats field if non-nil, zero value otherwise.

### GetConditionalFormatsOk

`func (o *ToplistWidgetRequest) GetConditionalFormatsOk() (*[]WidgetConditionalFormat, bool)`

GetConditionalFormatsOk returns a tuple with the ConditionalFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalFormats

`func (o *ToplistWidgetRequest) SetConditionalFormats(v []WidgetConditionalFormat)`

SetConditionalFormats sets ConditionalFormats field to given value.

### HasConditionalFormats

`func (o *ToplistWidgetRequest) HasConditionalFormats() bool`

HasConditionalFormats returns a boolean if a field has been set.

### GetEventQuery

`func (o *ToplistWidgetRequest) GetEventQuery() EventQueryDefinition`

GetEventQuery returns the EventQuery field if non-nil, zero value otherwise.

### GetEventQueryOk

`func (o *ToplistWidgetRequest) GetEventQueryOk() (*EventQueryDefinition, bool)`

GetEventQueryOk returns a tuple with the EventQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventQuery

`func (o *ToplistWidgetRequest) SetEventQuery(v EventQueryDefinition)`

SetEventQuery sets EventQuery field to given value.

### HasEventQuery

`func (o *ToplistWidgetRequest) HasEventQuery() bool`

HasEventQuery returns a boolean if a field has been set.

### GetLogQuery

`func (o *ToplistWidgetRequest) GetLogQuery() LogQueryDefinition`

GetLogQuery returns the LogQuery field if non-nil, zero value otherwise.

### GetLogQueryOk

`func (o *ToplistWidgetRequest) GetLogQueryOk() (*LogQueryDefinition, bool)`

GetLogQueryOk returns a tuple with the LogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQuery

`func (o *ToplistWidgetRequest) SetLogQuery(v LogQueryDefinition)`

SetLogQuery sets LogQuery field to given value.

### HasLogQuery

`func (o *ToplistWidgetRequest) HasLogQuery() bool`

HasLogQuery returns a boolean if a field has been set.

### GetNetworkQuery

`func (o *ToplistWidgetRequest) GetNetworkQuery() LogQueryDefinition`

GetNetworkQuery returns the NetworkQuery field if non-nil, zero value otherwise.

### GetNetworkQueryOk

`func (o *ToplistWidgetRequest) GetNetworkQueryOk() (*LogQueryDefinition, bool)`

GetNetworkQueryOk returns a tuple with the NetworkQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkQuery

`func (o *ToplistWidgetRequest) SetNetworkQuery(v LogQueryDefinition)`

SetNetworkQuery sets NetworkQuery field to given value.

### HasNetworkQuery

`func (o *ToplistWidgetRequest) HasNetworkQuery() bool`

HasNetworkQuery returns a boolean if a field has been set.

### GetProcessQuery

`func (o *ToplistWidgetRequest) GetProcessQuery() ProcessQueryDefinition`

GetProcessQuery returns the ProcessQuery field if non-nil, zero value otherwise.

### GetProcessQueryOk

`func (o *ToplistWidgetRequest) GetProcessQueryOk() (*ProcessQueryDefinition, bool)`

GetProcessQueryOk returns a tuple with the ProcessQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessQuery

`func (o *ToplistWidgetRequest) SetProcessQuery(v ProcessQueryDefinition)`

SetProcessQuery sets ProcessQuery field to given value.

### HasProcessQuery

`func (o *ToplistWidgetRequest) HasProcessQuery() bool`

HasProcessQuery returns a boolean if a field has been set.

### GetQ

`func (o *ToplistWidgetRequest) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *ToplistWidgetRequest) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *ToplistWidgetRequest) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *ToplistWidgetRequest) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetRumQuery

`func (o *ToplistWidgetRequest) GetRumQuery() LogQueryDefinition`

GetRumQuery returns the RumQuery field if non-nil, zero value otherwise.

### GetRumQueryOk

`func (o *ToplistWidgetRequest) GetRumQueryOk() (*LogQueryDefinition, bool)`

GetRumQueryOk returns a tuple with the RumQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumQuery

`func (o *ToplistWidgetRequest) SetRumQuery(v LogQueryDefinition)`

SetRumQuery sets RumQuery field to given value.

### HasRumQuery

`func (o *ToplistWidgetRequest) HasRumQuery() bool`

HasRumQuery returns a boolean if a field has been set.

### GetStyle

`func (o *ToplistWidgetRequest) GetStyle() TimeseriesWidgetRequestStyle`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ToplistWidgetRequest) GetStyleOk() (*TimeseriesWidgetRequestStyle, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ToplistWidgetRequest) SetStyle(v TimeseriesWidgetRequestStyle)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *ToplistWidgetRequest) HasStyle() bool`

HasStyle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


