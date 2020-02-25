# TimeseriesWidgetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApmQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**DisplayType** | Pointer to [**WidgetDisplayType**](WidgetDisplayType.md) |  | [optional] 
**EventQuery** | Pointer to [**EventQueryDefinition**](EventQueryDefinition.md) |  | [optional] 
**LogQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**Metadata** | Pointer to [**[]TimeseriesWidgetRequestMetadata**](TimeseriesWidgetRequest_metadata.md) | Used to define expression aliases. | [optional] 
**NetworkQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**ProcessQuery** | Pointer to [**ProcessQueryDefinition**](ProcessQueryDefinition.md) |  | [optional] 
**Q** | Pointer to **string** |  | [optional] 
**RumQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**Style** | Pointer to [**TimeseriesWidgetRequestStyle**](TimeseriesWidgetRequest_style.md) |  | [optional] 

## Methods

### NewTimeseriesWidgetRequest

`func NewTimeseriesWidgetRequest() *TimeseriesWidgetRequest`

NewTimeseriesWidgetRequest instantiates a new TimeseriesWidgetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeseriesWidgetRequestWithDefaults

`func NewTimeseriesWidgetRequestWithDefaults() *TimeseriesWidgetRequest`

NewTimeseriesWidgetRequestWithDefaults instantiates a new TimeseriesWidgetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApmQuery

`func (o *TimeseriesWidgetRequest) GetApmQuery() LogQueryDefinition`

GetApmQuery returns the ApmQuery field if non-nil, zero value otherwise.

### GetApmQueryOk

`func (o *TimeseriesWidgetRequest) GetApmQueryOk() (LogQueryDefinition, bool)`

GetApmQueryOk returns a tuple with the ApmQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApmQuery

`func (o *TimeseriesWidgetRequest) HasApmQuery() bool`

HasApmQuery returns a boolean if a field has been set.

### SetApmQuery

`func (o *TimeseriesWidgetRequest) SetApmQuery(v LogQueryDefinition)`

SetApmQuery gets a reference to the given LogQueryDefinition and assigns it to the ApmQuery field.

### GetDisplayType

`func (o *TimeseriesWidgetRequest) GetDisplayType() WidgetDisplayType`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *TimeseriesWidgetRequest) GetDisplayTypeOk() (WidgetDisplayType, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayType

`func (o *TimeseriesWidgetRequest) HasDisplayType() bool`

HasDisplayType returns a boolean if a field has been set.

### SetDisplayType

`func (o *TimeseriesWidgetRequest) SetDisplayType(v WidgetDisplayType)`

SetDisplayType gets a reference to the given WidgetDisplayType and assigns it to the DisplayType field.

### GetEventQuery

`func (o *TimeseriesWidgetRequest) GetEventQuery() EventQueryDefinition`

GetEventQuery returns the EventQuery field if non-nil, zero value otherwise.

### GetEventQueryOk

`func (o *TimeseriesWidgetRequest) GetEventQueryOk() (EventQueryDefinition, bool)`

GetEventQueryOk returns a tuple with the EventQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventQuery

`func (o *TimeseriesWidgetRequest) HasEventQuery() bool`

HasEventQuery returns a boolean if a field has been set.

### SetEventQuery

`func (o *TimeseriesWidgetRequest) SetEventQuery(v EventQueryDefinition)`

SetEventQuery gets a reference to the given EventQueryDefinition and assigns it to the EventQuery field.

### GetLogQuery

`func (o *TimeseriesWidgetRequest) GetLogQuery() LogQueryDefinition`

GetLogQuery returns the LogQuery field if non-nil, zero value otherwise.

### GetLogQueryOk

`func (o *TimeseriesWidgetRequest) GetLogQueryOk() (LogQueryDefinition, bool)`

GetLogQueryOk returns a tuple with the LogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLogQuery

`func (o *TimeseriesWidgetRequest) HasLogQuery() bool`

HasLogQuery returns a boolean if a field has been set.

### SetLogQuery

`func (o *TimeseriesWidgetRequest) SetLogQuery(v LogQueryDefinition)`

SetLogQuery gets a reference to the given LogQueryDefinition and assigns it to the LogQuery field.

### GetMetadata

`func (o *TimeseriesWidgetRequest) GetMetadata() []TimeseriesWidgetRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TimeseriesWidgetRequest) GetMetadataOk() ([]TimeseriesWidgetRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMetadata

`func (o *TimeseriesWidgetRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadata

`func (o *TimeseriesWidgetRequest) SetMetadata(v []TimeseriesWidgetRequestMetadata)`

SetMetadata gets a reference to the given []TimeseriesWidgetRequestMetadata and assigns it to the Metadata field.

### GetNetworkQuery

`func (o *TimeseriesWidgetRequest) GetNetworkQuery() LogQueryDefinition`

GetNetworkQuery returns the NetworkQuery field if non-nil, zero value otherwise.

### GetNetworkQueryOk

`func (o *TimeseriesWidgetRequest) GetNetworkQueryOk() (LogQueryDefinition, bool)`

GetNetworkQueryOk returns a tuple with the NetworkQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetworkQuery

`func (o *TimeseriesWidgetRequest) HasNetworkQuery() bool`

HasNetworkQuery returns a boolean if a field has been set.

### SetNetworkQuery

`func (o *TimeseriesWidgetRequest) SetNetworkQuery(v LogQueryDefinition)`

SetNetworkQuery gets a reference to the given LogQueryDefinition and assigns it to the NetworkQuery field.

### GetProcessQuery

`func (o *TimeseriesWidgetRequest) GetProcessQuery() ProcessQueryDefinition`

GetProcessQuery returns the ProcessQuery field if non-nil, zero value otherwise.

### GetProcessQueryOk

`func (o *TimeseriesWidgetRequest) GetProcessQueryOk() (ProcessQueryDefinition, bool)`

GetProcessQueryOk returns a tuple with the ProcessQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProcessQuery

`func (o *TimeseriesWidgetRequest) HasProcessQuery() bool`

HasProcessQuery returns a boolean if a field has been set.

### SetProcessQuery

`func (o *TimeseriesWidgetRequest) SetProcessQuery(v ProcessQueryDefinition)`

SetProcessQuery gets a reference to the given ProcessQueryDefinition and assigns it to the ProcessQuery field.

### GetQ

`func (o *TimeseriesWidgetRequest) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *TimeseriesWidgetRequest) GetQOk() (string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQ

`func (o *TimeseriesWidgetRequest) HasQ() bool`

HasQ returns a boolean if a field has been set.

### SetQ

`func (o *TimeseriesWidgetRequest) SetQ(v string)`

SetQ gets a reference to the given string and assigns it to the Q field.

### GetRumQuery

`func (o *TimeseriesWidgetRequest) GetRumQuery() LogQueryDefinition`

GetRumQuery returns the RumQuery field if non-nil, zero value otherwise.

### GetRumQueryOk

`func (o *TimeseriesWidgetRequest) GetRumQueryOk() (LogQueryDefinition, bool)`

GetRumQueryOk returns a tuple with the RumQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRumQuery

`func (o *TimeseriesWidgetRequest) HasRumQuery() bool`

HasRumQuery returns a boolean if a field has been set.

### SetRumQuery

`func (o *TimeseriesWidgetRequest) SetRumQuery(v LogQueryDefinition)`

SetRumQuery gets a reference to the given LogQueryDefinition and assigns it to the RumQuery field.

### GetStyle

`func (o *TimeseriesWidgetRequest) GetStyle() TimeseriesWidgetRequestStyle`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *TimeseriesWidgetRequest) GetStyleOk() (TimeseriesWidgetRequestStyle, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStyle

`func (o *TimeseriesWidgetRequest) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### SetStyle

`func (o *TimeseriesWidgetRequest) SetStyle(v TimeseriesWidgetRequestStyle)`

SetStyle gets a reference to the given TimeseriesWidgetRequestStyle and assigns it to the Style field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


