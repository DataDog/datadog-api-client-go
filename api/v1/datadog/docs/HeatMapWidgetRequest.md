# HeatMapWidgetRequest

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**ApmQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**EventQuery** | Pointer to [**EventQueryDefinition**](EventQueryDefinition.md) |  | [optional] 
**LogQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**NetworkQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**ProcessQuery** | Pointer to [**ProcessQueryDefinition**](ProcessQueryDefinition.md) |  | [optional] 
**ProfileMetricsQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**Q** | Pointer to **string** | Widget query. | [optional] 
**RumQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**SecurityQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**Style** | Pointer to [**WidgetStyle**](WidgetStyle.md) |  | [optional] 

## Methods

### NewHeatMapWidgetRequest

`func NewHeatMapWidgetRequest() *HeatMapWidgetRequest`

NewHeatMapWidgetRequest instantiates a new HeatMapWidgetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeatMapWidgetRequestWithDefaults

`func NewHeatMapWidgetRequestWithDefaults() *HeatMapWidgetRequest`

NewHeatMapWidgetRequestWithDefaults instantiates a new HeatMapWidgetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApmQuery

`func (o *HeatMapWidgetRequest) GetApmQuery() LogQueryDefinition`

GetApmQuery returns the ApmQuery field if non-nil, zero value otherwise.

### GetApmQueryOk

`func (o *HeatMapWidgetRequest) GetApmQueryOk() (*LogQueryDefinition, bool)`

GetApmQueryOk returns a tuple with the ApmQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmQuery

`func (o *HeatMapWidgetRequest) SetApmQuery(v LogQueryDefinition)`

SetApmQuery sets ApmQuery field to given value.

### HasApmQuery

`func (o *HeatMapWidgetRequest) HasApmQuery() bool`

HasApmQuery returns a boolean if a field has been set.

### GetEventQuery

`func (o *HeatMapWidgetRequest) GetEventQuery() EventQueryDefinition`

GetEventQuery returns the EventQuery field if non-nil, zero value otherwise.

### GetEventQueryOk

`func (o *HeatMapWidgetRequest) GetEventQueryOk() (*EventQueryDefinition, bool)`

GetEventQueryOk returns a tuple with the EventQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventQuery

`func (o *HeatMapWidgetRequest) SetEventQuery(v EventQueryDefinition)`

SetEventQuery sets EventQuery field to given value.

### HasEventQuery

`func (o *HeatMapWidgetRequest) HasEventQuery() bool`

HasEventQuery returns a boolean if a field has been set.

### GetLogQuery

`func (o *HeatMapWidgetRequest) GetLogQuery() LogQueryDefinition`

GetLogQuery returns the LogQuery field if non-nil, zero value otherwise.

### GetLogQueryOk

`func (o *HeatMapWidgetRequest) GetLogQueryOk() (*LogQueryDefinition, bool)`

GetLogQueryOk returns a tuple with the LogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQuery

`func (o *HeatMapWidgetRequest) SetLogQuery(v LogQueryDefinition)`

SetLogQuery sets LogQuery field to given value.

### HasLogQuery

`func (o *HeatMapWidgetRequest) HasLogQuery() bool`

HasLogQuery returns a boolean if a field has been set.

### GetNetworkQuery

`func (o *HeatMapWidgetRequest) GetNetworkQuery() LogQueryDefinition`

GetNetworkQuery returns the NetworkQuery field if non-nil, zero value otherwise.

### GetNetworkQueryOk

`func (o *HeatMapWidgetRequest) GetNetworkQueryOk() (*LogQueryDefinition, bool)`

GetNetworkQueryOk returns a tuple with the NetworkQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkQuery

`func (o *HeatMapWidgetRequest) SetNetworkQuery(v LogQueryDefinition)`

SetNetworkQuery sets NetworkQuery field to given value.

### HasNetworkQuery

`func (o *HeatMapWidgetRequest) HasNetworkQuery() bool`

HasNetworkQuery returns a boolean if a field has been set.

### GetProcessQuery

`func (o *HeatMapWidgetRequest) GetProcessQuery() ProcessQueryDefinition`

GetProcessQuery returns the ProcessQuery field if non-nil, zero value otherwise.

### GetProcessQueryOk

`func (o *HeatMapWidgetRequest) GetProcessQueryOk() (*ProcessQueryDefinition, bool)`

GetProcessQueryOk returns a tuple with the ProcessQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessQuery

`func (o *HeatMapWidgetRequest) SetProcessQuery(v ProcessQueryDefinition)`

SetProcessQuery sets ProcessQuery field to given value.

### HasProcessQuery

`func (o *HeatMapWidgetRequest) HasProcessQuery() bool`

HasProcessQuery returns a boolean if a field has been set.

### GetProfileMetricsQuery

`func (o *HeatMapWidgetRequest) GetProfileMetricsQuery() LogQueryDefinition`

GetProfileMetricsQuery returns the ProfileMetricsQuery field if non-nil, zero value otherwise.

### GetProfileMetricsQueryOk

`func (o *HeatMapWidgetRequest) GetProfileMetricsQueryOk() (*LogQueryDefinition, bool)`

GetProfileMetricsQueryOk returns a tuple with the ProfileMetricsQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMetricsQuery

`func (o *HeatMapWidgetRequest) SetProfileMetricsQuery(v LogQueryDefinition)`

SetProfileMetricsQuery sets ProfileMetricsQuery field to given value.

### HasProfileMetricsQuery

`func (o *HeatMapWidgetRequest) HasProfileMetricsQuery() bool`

HasProfileMetricsQuery returns a boolean if a field has been set.

### GetQ

`func (o *HeatMapWidgetRequest) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *HeatMapWidgetRequest) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *HeatMapWidgetRequest) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *HeatMapWidgetRequest) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetRumQuery

`func (o *HeatMapWidgetRequest) GetRumQuery() LogQueryDefinition`

GetRumQuery returns the RumQuery field if non-nil, zero value otherwise.

### GetRumQueryOk

`func (o *HeatMapWidgetRequest) GetRumQueryOk() (*LogQueryDefinition, bool)`

GetRumQueryOk returns a tuple with the RumQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumQuery

`func (o *HeatMapWidgetRequest) SetRumQuery(v LogQueryDefinition)`

SetRumQuery sets RumQuery field to given value.

### HasRumQuery

`func (o *HeatMapWidgetRequest) HasRumQuery() bool`

HasRumQuery returns a boolean if a field has been set.

### GetSecurityQuery

`func (o *HeatMapWidgetRequest) GetSecurityQuery() LogQueryDefinition`

GetSecurityQuery returns the SecurityQuery field if non-nil, zero value otherwise.

### GetSecurityQueryOk

`func (o *HeatMapWidgetRequest) GetSecurityQueryOk() (*LogQueryDefinition, bool)`

GetSecurityQueryOk returns a tuple with the SecurityQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityQuery

`func (o *HeatMapWidgetRequest) SetSecurityQuery(v LogQueryDefinition)`

SetSecurityQuery sets SecurityQuery field to given value.

### HasSecurityQuery

`func (o *HeatMapWidgetRequest) HasSecurityQuery() bool`

HasSecurityQuery returns a boolean if a field has been set.

### GetStyle

`func (o *HeatMapWidgetRequest) GetStyle() WidgetStyle`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *HeatMapWidgetRequest) GetStyleOk() (*WidgetStyle, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *HeatMapWidgetRequest) SetStyle(v WidgetStyle)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *HeatMapWidgetRequest) HasStyle() bool`

HasStyle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


