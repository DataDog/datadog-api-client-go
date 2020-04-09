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
**Q** | Pointer to **string** | TODO. | [optional] 
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

`func (o *TimeseriesWidgetRequest) GetApmQueryOk() (*LogQueryDefinition, bool)`

GetApmQueryOk returns a tuple with the ApmQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmQuery

`func (o *TimeseriesWidgetRequest) SetApmQuery(v LogQueryDefinition)`

SetApmQuery sets ApmQuery field to given value.

### HasApmQuery

`func (o *TimeseriesWidgetRequest) HasApmQuery() bool`

HasApmQuery returns a boolean if a field has been set.

### GetDisplayType

`func (o *TimeseriesWidgetRequest) GetDisplayType() WidgetDisplayType`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *TimeseriesWidgetRequest) GetDisplayTypeOk() (*WidgetDisplayType, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *TimeseriesWidgetRequest) SetDisplayType(v WidgetDisplayType)`

SetDisplayType sets DisplayType field to given value.

### HasDisplayType

`func (o *TimeseriesWidgetRequest) HasDisplayType() bool`

HasDisplayType returns a boolean if a field has been set.

### GetEventQuery

`func (o *TimeseriesWidgetRequest) GetEventQuery() EventQueryDefinition`

GetEventQuery returns the EventQuery field if non-nil, zero value otherwise.

### GetEventQueryOk

`func (o *TimeseriesWidgetRequest) GetEventQueryOk() (*EventQueryDefinition, bool)`

GetEventQueryOk returns a tuple with the EventQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventQuery

`func (o *TimeseriesWidgetRequest) SetEventQuery(v EventQueryDefinition)`

SetEventQuery sets EventQuery field to given value.

### HasEventQuery

`func (o *TimeseriesWidgetRequest) HasEventQuery() bool`

HasEventQuery returns a boolean if a field has been set.

### GetLogQuery

`func (o *TimeseriesWidgetRequest) GetLogQuery() LogQueryDefinition`

GetLogQuery returns the LogQuery field if non-nil, zero value otherwise.

### GetLogQueryOk

`func (o *TimeseriesWidgetRequest) GetLogQueryOk() (*LogQueryDefinition, bool)`

GetLogQueryOk returns a tuple with the LogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQuery

`func (o *TimeseriesWidgetRequest) SetLogQuery(v LogQueryDefinition)`

SetLogQuery sets LogQuery field to given value.

### HasLogQuery

`func (o *TimeseriesWidgetRequest) HasLogQuery() bool`

HasLogQuery returns a boolean if a field has been set.

### GetMetadata

`func (o *TimeseriesWidgetRequest) GetMetadata() []TimeseriesWidgetRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TimeseriesWidgetRequest) GetMetadataOk() (*[]TimeseriesWidgetRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TimeseriesWidgetRequest) SetMetadata(v []TimeseriesWidgetRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TimeseriesWidgetRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNetworkQuery

`func (o *TimeseriesWidgetRequest) GetNetworkQuery() LogQueryDefinition`

GetNetworkQuery returns the NetworkQuery field if non-nil, zero value otherwise.

### GetNetworkQueryOk

`func (o *TimeseriesWidgetRequest) GetNetworkQueryOk() (*LogQueryDefinition, bool)`

GetNetworkQueryOk returns a tuple with the NetworkQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkQuery

`func (o *TimeseriesWidgetRequest) SetNetworkQuery(v LogQueryDefinition)`

SetNetworkQuery sets NetworkQuery field to given value.

### HasNetworkQuery

`func (o *TimeseriesWidgetRequest) HasNetworkQuery() bool`

HasNetworkQuery returns a boolean if a field has been set.

### GetProcessQuery

`func (o *TimeseriesWidgetRequest) GetProcessQuery() ProcessQueryDefinition`

GetProcessQuery returns the ProcessQuery field if non-nil, zero value otherwise.

### GetProcessQueryOk

`func (o *TimeseriesWidgetRequest) GetProcessQueryOk() (*ProcessQueryDefinition, bool)`

GetProcessQueryOk returns a tuple with the ProcessQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessQuery

`func (o *TimeseriesWidgetRequest) SetProcessQuery(v ProcessQueryDefinition)`

SetProcessQuery sets ProcessQuery field to given value.

### HasProcessQuery

`func (o *TimeseriesWidgetRequest) HasProcessQuery() bool`

HasProcessQuery returns a boolean if a field has been set.

### GetQ

`func (o *TimeseriesWidgetRequest) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *TimeseriesWidgetRequest) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *TimeseriesWidgetRequest) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *TimeseriesWidgetRequest) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetRumQuery

`func (o *TimeseriesWidgetRequest) GetRumQuery() LogQueryDefinition`

GetRumQuery returns the RumQuery field if non-nil, zero value otherwise.

### GetRumQueryOk

`func (o *TimeseriesWidgetRequest) GetRumQueryOk() (*LogQueryDefinition, bool)`

GetRumQueryOk returns a tuple with the RumQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumQuery

`func (o *TimeseriesWidgetRequest) SetRumQuery(v LogQueryDefinition)`

SetRumQuery sets RumQuery field to given value.

### HasRumQuery

`func (o *TimeseriesWidgetRequest) HasRumQuery() bool`

HasRumQuery returns a boolean if a field has been set.

### GetStyle

`func (o *TimeseriesWidgetRequest) GetStyle() TimeseriesWidgetRequestStyle`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *TimeseriesWidgetRequest) GetStyleOk() (*TimeseriesWidgetRequestStyle, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *TimeseriesWidgetRequest) SetStyle(v TimeseriesWidgetRequestStyle)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *TimeseriesWidgetRequest) HasStyle() bool`

HasStyle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


