# DistributionWidgetDefinition

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**LegendSize** | Pointer to **string** | (Deprecated) The widget legend was replaced by a tooltip and sidebar. | [optional] 
**Markers** | Pointer to [**[]WidgetMarker**](WidgetMarker.md) | List of markers. | [optional] 
**Requests** | [**[]DistributionWidgetRequest**](DistributionWidgetRequest.md) | Array of one request object to display in the widget.  See the dedicated [Request JSON schema documentation](https://docs.datadoghq.com/dashboards/graphing_json/request_json)  to learn how to build the &#x60;REQUEST_SCHEMA&#x60;. | 
**ShowLegend** | Pointer to **bool** | (Deprecated) The widget legend was replaced by a tooltip and sidebar. | [optional] 
**Time** | Pointer to [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the widget. | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title. | [optional] 
**Type** | [**DistributionWidgetDefinitionType**](DistributionWidgetDefinitionType.md) |  | [default to DISTRIBUTIONWIDGETDEFINITIONTYPE_DISTRIBUTION]
**Xaxis** | Pointer to [**DistributionWidgetXAxis**](DistributionWidgetXAxis.md) |  | [optional] 
**Yaxis** | Pointer to [**DistributionWidgetYAxis**](DistributionWidgetYAxis.md) |  | [optional] 

## Methods

### NewDistributionWidgetDefinition

`func NewDistributionWidgetDefinition(requests []DistributionWidgetRequest, type_ DistributionWidgetDefinitionType) *DistributionWidgetDefinition`

NewDistributionWidgetDefinition instantiates a new DistributionWidgetDefinition object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewDistributionWidgetDefinitionWithDefaults

`func NewDistributionWidgetDefinitionWithDefaults() *DistributionWidgetDefinition`

NewDistributionWidgetDefinitionWithDefaults instantiates a new DistributionWidgetDefinition object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetLegendSize

`func (o *DistributionWidgetDefinition) GetLegendSize() string`

GetLegendSize returns the LegendSize field if non-nil, zero value otherwise.

### GetLegendSizeOk

`func (o *DistributionWidgetDefinition) GetLegendSizeOk() (*string, bool)`

GetLegendSizeOk returns a tuple with the LegendSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendSize

`func (o *DistributionWidgetDefinition) SetLegendSize(v string)`

SetLegendSize sets LegendSize field to given value.

### HasLegendSize

`func (o *DistributionWidgetDefinition) HasLegendSize() bool`

HasLegendSize returns a boolean if a field has been set.

### GetMarkers

`func (o *DistributionWidgetDefinition) GetMarkers() []WidgetMarker`

GetMarkers returns the Markers field if non-nil, zero value otherwise.

### GetMarkersOk

`func (o *DistributionWidgetDefinition) GetMarkersOk() (*[]WidgetMarker, bool)`

GetMarkersOk returns a tuple with the Markers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkers

`func (o *DistributionWidgetDefinition) SetMarkers(v []WidgetMarker)`

SetMarkers sets Markers field to given value.

### HasMarkers

`func (o *DistributionWidgetDefinition) HasMarkers() bool`

HasMarkers returns a boolean if a field has been set.

### GetRequests

`func (o *DistributionWidgetDefinition) GetRequests() []DistributionWidgetRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *DistributionWidgetDefinition) GetRequestsOk() (*[]DistributionWidgetRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *DistributionWidgetDefinition) SetRequests(v []DistributionWidgetRequest)`

SetRequests sets Requests field to given value.


### GetShowLegend

`func (o *DistributionWidgetDefinition) GetShowLegend() bool`

GetShowLegend returns the ShowLegend field if non-nil, zero value otherwise.

### GetShowLegendOk

`func (o *DistributionWidgetDefinition) GetShowLegendOk() (*bool, bool)`

GetShowLegendOk returns a tuple with the ShowLegend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLegend

`func (o *DistributionWidgetDefinition) SetShowLegend(v bool)`

SetShowLegend sets ShowLegend field to given value.

### HasShowLegend

`func (o *DistributionWidgetDefinition) HasShowLegend() bool`

HasShowLegend returns a boolean if a field has been set.

### GetTime

`func (o *DistributionWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DistributionWidgetDefinition) GetTimeOk() (*WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DistributionWidgetDefinition) SetTime(v WidgetTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *DistributionWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTitle

`func (o *DistributionWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DistributionWidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DistributionWidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DistributionWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAlign

`func (o *DistributionWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *DistributionWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAlign

`func (o *DistributionWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign sets TitleAlign field to given value.

### HasTitleAlign

`func (o *DistributionWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### GetTitleSize

`func (o *DistributionWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *DistributionWidgetDefinition) GetTitleSizeOk() (*string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *DistributionWidgetDefinition) SetTitleSize(v string)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *DistributionWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetType

`func (o *DistributionWidgetDefinition) GetType() DistributionWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DistributionWidgetDefinition) GetTypeOk() (*DistributionWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DistributionWidgetDefinition) SetType(v DistributionWidgetDefinitionType)`

SetType sets Type field to given value.


### GetXaxis

`func (o *DistributionWidgetDefinition) GetXaxis() DistributionWidgetXAxis`

GetXaxis returns the Xaxis field if non-nil, zero value otherwise.

### GetXaxisOk

`func (o *DistributionWidgetDefinition) GetXaxisOk() (*DistributionWidgetXAxis, bool)`

GetXaxisOk returns a tuple with the Xaxis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXaxis

`func (o *DistributionWidgetDefinition) SetXaxis(v DistributionWidgetXAxis)`

SetXaxis sets Xaxis field to given value.

### HasXaxis

`func (o *DistributionWidgetDefinition) HasXaxis() bool`

HasXaxis returns a boolean if a field has been set.

### GetYaxis

`func (o *DistributionWidgetDefinition) GetYaxis() DistributionWidgetYAxis`

GetYaxis returns the Yaxis field if non-nil, zero value otherwise.

### GetYaxisOk

`func (o *DistributionWidgetDefinition) GetYaxisOk() (*DistributionWidgetYAxis, bool)`

GetYaxisOk returns a tuple with the Yaxis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaxis

`func (o *DistributionWidgetDefinition) SetYaxis(v DistributionWidgetYAxis)`

SetYaxis sets Yaxis field to given value.

### HasYaxis

`func (o *DistributionWidgetDefinition) HasYaxis() bool`

HasYaxis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


