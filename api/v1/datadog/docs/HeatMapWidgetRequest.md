# HeatMapWidgetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApmQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**EventQuery** | Pointer to [**EventQueryDefinition**](EventQueryDefinition.md) |  | [optional] 
**LogQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**NetworkQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**ProcessQuery** | Pointer to [**ProcessQueryDefinition**](ProcessQueryDefinition.md) |  | [optional] 
**Q** | Pointer to **string** |  | [optional] 
**RumQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
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

`func (o *HeatMapWidgetRequest) GetApmQueryOk() (LogQueryDefinition, bool)`

GetApmQueryOk returns a tuple with the ApmQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApmQuery

`func (o *HeatMapWidgetRequest) HasApmQuery() bool`

HasApmQuery returns a boolean if a field has been set.

### SetApmQuery

`func (o *HeatMapWidgetRequest) SetApmQuery(v LogQueryDefinition)`

SetApmQuery gets a reference to the given LogQueryDefinition and assigns it to the ApmQuery field.

### GetEventQuery

`func (o *HeatMapWidgetRequest) GetEventQuery() EventQueryDefinition`

GetEventQuery returns the EventQuery field if non-nil, zero value otherwise.

### GetEventQueryOk

`func (o *HeatMapWidgetRequest) GetEventQueryOk() (EventQueryDefinition, bool)`

GetEventQueryOk returns a tuple with the EventQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventQuery

`func (o *HeatMapWidgetRequest) HasEventQuery() bool`

HasEventQuery returns a boolean if a field has been set.

### SetEventQuery

`func (o *HeatMapWidgetRequest) SetEventQuery(v EventQueryDefinition)`

SetEventQuery gets a reference to the given EventQueryDefinition and assigns it to the EventQuery field.

### GetLogQuery

`func (o *HeatMapWidgetRequest) GetLogQuery() LogQueryDefinition`

GetLogQuery returns the LogQuery field if non-nil, zero value otherwise.

### GetLogQueryOk

`func (o *HeatMapWidgetRequest) GetLogQueryOk() (LogQueryDefinition, bool)`

GetLogQueryOk returns a tuple with the LogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLogQuery

`func (o *HeatMapWidgetRequest) HasLogQuery() bool`

HasLogQuery returns a boolean if a field has been set.

### SetLogQuery

`func (o *HeatMapWidgetRequest) SetLogQuery(v LogQueryDefinition)`

SetLogQuery gets a reference to the given LogQueryDefinition and assigns it to the LogQuery field.

### GetNetworkQuery

`func (o *HeatMapWidgetRequest) GetNetworkQuery() LogQueryDefinition`

GetNetworkQuery returns the NetworkQuery field if non-nil, zero value otherwise.

### GetNetworkQueryOk

`func (o *HeatMapWidgetRequest) GetNetworkQueryOk() (LogQueryDefinition, bool)`

GetNetworkQueryOk returns a tuple with the NetworkQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetworkQuery

`func (o *HeatMapWidgetRequest) HasNetworkQuery() bool`

HasNetworkQuery returns a boolean if a field has been set.

### SetNetworkQuery

`func (o *HeatMapWidgetRequest) SetNetworkQuery(v LogQueryDefinition)`

SetNetworkQuery gets a reference to the given LogQueryDefinition and assigns it to the NetworkQuery field.

### GetProcessQuery

`func (o *HeatMapWidgetRequest) GetProcessQuery() ProcessQueryDefinition`

GetProcessQuery returns the ProcessQuery field if non-nil, zero value otherwise.

### GetProcessQueryOk

`func (o *HeatMapWidgetRequest) GetProcessQueryOk() (ProcessQueryDefinition, bool)`

GetProcessQueryOk returns a tuple with the ProcessQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProcessQuery

`func (o *HeatMapWidgetRequest) HasProcessQuery() bool`

HasProcessQuery returns a boolean if a field has been set.

### SetProcessQuery

`func (o *HeatMapWidgetRequest) SetProcessQuery(v ProcessQueryDefinition)`

SetProcessQuery gets a reference to the given ProcessQueryDefinition and assigns it to the ProcessQuery field.

### GetQ

`func (o *HeatMapWidgetRequest) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *HeatMapWidgetRequest) GetQOk() (string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQ

`func (o *HeatMapWidgetRequest) HasQ() bool`

HasQ returns a boolean if a field has been set.

### SetQ

`func (o *HeatMapWidgetRequest) SetQ(v string)`

SetQ gets a reference to the given string and assigns it to the Q field.

### GetRumQuery

`func (o *HeatMapWidgetRequest) GetRumQuery() LogQueryDefinition`

GetRumQuery returns the RumQuery field if non-nil, zero value otherwise.

### GetRumQueryOk

`func (o *HeatMapWidgetRequest) GetRumQueryOk() (LogQueryDefinition, bool)`

GetRumQueryOk returns a tuple with the RumQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRumQuery

`func (o *HeatMapWidgetRequest) HasRumQuery() bool`

HasRumQuery returns a boolean if a field has been set.

### SetRumQuery

`func (o *HeatMapWidgetRequest) SetRumQuery(v LogQueryDefinition)`

SetRumQuery gets a reference to the given LogQueryDefinition and assigns it to the RumQuery field.

### GetStyle

`func (o *HeatMapWidgetRequest) GetStyle() WidgetStyle`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *HeatMapWidgetRequest) GetStyleOk() (WidgetStyle, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStyle

`func (o *HeatMapWidgetRequest) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### SetStyle

`func (o *HeatMapWidgetRequest) SetStyle(v WidgetStyle)`

SetStyle gets a reference to the given WidgetStyle and assigns it to the Style field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


