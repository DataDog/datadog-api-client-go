# DistributionWidgetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApmQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**EventQuery** | Pointer to [**EventQueryDefinition**](EventQueryDefinition.md) |  | [optional] 
**LogQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**NetworkQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**ProcessQuery** | Pointer to [**ProcessQueryDefinition**](ProcessQueryDefinition.md) |  | [optional] 
**Q** | Pointer to **string** | Widget query. | [optional] 
**RumQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**Style** | Pointer to [**WidgetStyle**](WidgetStyle.md) |  | [optional] 

## Methods

### NewDistributionWidgetRequest

`func NewDistributionWidgetRequest() *DistributionWidgetRequest`

NewDistributionWidgetRequest instantiates a new DistributionWidgetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributionWidgetRequestWithDefaults

`func NewDistributionWidgetRequestWithDefaults() *DistributionWidgetRequest`

NewDistributionWidgetRequestWithDefaults instantiates a new DistributionWidgetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApmQuery

`func (o *DistributionWidgetRequest) GetApmQuery() LogQueryDefinition`

GetApmQuery returns the ApmQuery field if non-nil, zero value otherwise.

### GetApmQueryOk

`func (o *DistributionWidgetRequest) GetApmQueryOk() (*LogQueryDefinition, bool)`

GetApmQueryOk returns a tuple with the ApmQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmQuery

`func (o *DistributionWidgetRequest) SetApmQuery(v LogQueryDefinition)`

SetApmQuery sets ApmQuery field to given value.

### HasApmQuery

`func (o *DistributionWidgetRequest) HasApmQuery() bool`

HasApmQuery returns a boolean if a field has been set.

### GetEventQuery

`func (o *DistributionWidgetRequest) GetEventQuery() EventQueryDefinition`

GetEventQuery returns the EventQuery field if non-nil, zero value otherwise.

### GetEventQueryOk

`func (o *DistributionWidgetRequest) GetEventQueryOk() (*EventQueryDefinition, bool)`

GetEventQueryOk returns a tuple with the EventQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventQuery

`func (o *DistributionWidgetRequest) SetEventQuery(v EventQueryDefinition)`

SetEventQuery sets EventQuery field to given value.

### HasEventQuery

`func (o *DistributionWidgetRequest) HasEventQuery() bool`

HasEventQuery returns a boolean if a field has been set.

### GetLogQuery

`func (o *DistributionWidgetRequest) GetLogQuery() LogQueryDefinition`

GetLogQuery returns the LogQuery field if non-nil, zero value otherwise.

### GetLogQueryOk

`func (o *DistributionWidgetRequest) GetLogQueryOk() (*LogQueryDefinition, bool)`

GetLogQueryOk returns a tuple with the LogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQuery

`func (o *DistributionWidgetRequest) SetLogQuery(v LogQueryDefinition)`

SetLogQuery sets LogQuery field to given value.

### HasLogQuery

`func (o *DistributionWidgetRequest) HasLogQuery() bool`

HasLogQuery returns a boolean if a field has been set.

### GetNetworkQuery

`func (o *DistributionWidgetRequest) GetNetworkQuery() LogQueryDefinition`

GetNetworkQuery returns the NetworkQuery field if non-nil, zero value otherwise.

### GetNetworkQueryOk

`func (o *DistributionWidgetRequest) GetNetworkQueryOk() (*LogQueryDefinition, bool)`

GetNetworkQueryOk returns a tuple with the NetworkQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkQuery

`func (o *DistributionWidgetRequest) SetNetworkQuery(v LogQueryDefinition)`

SetNetworkQuery sets NetworkQuery field to given value.

### HasNetworkQuery

`func (o *DistributionWidgetRequest) HasNetworkQuery() bool`

HasNetworkQuery returns a boolean if a field has been set.

### GetProcessQuery

`func (o *DistributionWidgetRequest) GetProcessQuery() ProcessQueryDefinition`

GetProcessQuery returns the ProcessQuery field if non-nil, zero value otherwise.

### GetProcessQueryOk

`func (o *DistributionWidgetRequest) GetProcessQueryOk() (*ProcessQueryDefinition, bool)`

GetProcessQueryOk returns a tuple with the ProcessQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessQuery

`func (o *DistributionWidgetRequest) SetProcessQuery(v ProcessQueryDefinition)`

SetProcessQuery sets ProcessQuery field to given value.

### HasProcessQuery

`func (o *DistributionWidgetRequest) HasProcessQuery() bool`

HasProcessQuery returns a boolean if a field has been set.

### GetQ

`func (o *DistributionWidgetRequest) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *DistributionWidgetRequest) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *DistributionWidgetRequest) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *DistributionWidgetRequest) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetRumQuery

`func (o *DistributionWidgetRequest) GetRumQuery() LogQueryDefinition`

GetRumQuery returns the RumQuery field if non-nil, zero value otherwise.

### GetRumQueryOk

`func (o *DistributionWidgetRequest) GetRumQueryOk() (*LogQueryDefinition, bool)`

GetRumQueryOk returns a tuple with the RumQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumQuery

`func (o *DistributionWidgetRequest) SetRumQuery(v LogQueryDefinition)`

SetRumQuery sets RumQuery field to given value.

### HasRumQuery

`func (o *DistributionWidgetRequest) HasRumQuery() bool`

HasRumQuery returns a boolean if a field has been set.

### GetStyle

`func (o *DistributionWidgetRequest) GetStyle() WidgetStyle`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *DistributionWidgetRequest) GetStyleOk() (*WidgetStyle, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *DistributionWidgetRequest) SetStyle(v WidgetStyle)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *DistributionWidgetRequest) HasStyle() bool`

HasStyle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


