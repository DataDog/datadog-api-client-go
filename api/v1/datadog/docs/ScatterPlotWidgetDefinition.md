# ScatterPlotWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColorByGroups** | Pointer to **[]string** | List of groups used for colors. | [optional] 
**Requests** | Pointer to [**ScatterPlotWidgetDefinitionRequests**](ScatterPlotWidgetDefinition_requests.md) |  | 
**Time** | Pointer to [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | Pointer to **string** | Title of your widget. | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title. | [optional] 
**Type** | Pointer to [**ScatterPlotWidgetDefinitionType**](ScatterPlotWidgetDefinitionType.md) |  | [default to "scatterplot"]
**Xaxis** | Pointer to [**WidgetAxis**](WidgetAxis.md) |  | [optional] 
**Yaxis** | Pointer to [**WidgetAxis**](WidgetAxis.md) |  | [optional] 

## Methods

### NewScatterPlotWidgetDefinition

`func NewScatterPlotWidgetDefinition(requests ScatterPlotWidgetDefinitionRequests, type_ ScatterPlotWidgetDefinitionType, ) *ScatterPlotWidgetDefinition`

NewScatterPlotWidgetDefinition instantiates a new ScatterPlotWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScatterPlotWidgetDefinitionWithDefaults

`func NewScatterPlotWidgetDefinitionWithDefaults() *ScatterPlotWidgetDefinition`

NewScatterPlotWidgetDefinitionWithDefaults instantiates a new ScatterPlotWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColorByGroups

`func (o *ScatterPlotWidgetDefinition) GetColorByGroups() []string`

GetColorByGroups returns the ColorByGroups field if non-nil, zero value otherwise.

### GetColorByGroupsOk

`func (o *ScatterPlotWidgetDefinition) GetColorByGroupsOk() (*[]string, bool)`

GetColorByGroupsOk returns a tuple with the ColorByGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorByGroups

`func (o *ScatterPlotWidgetDefinition) SetColorByGroups(v []string)`

SetColorByGroups sets ColorByGroups field to given value.

### HasColorByGroups

`func (o *ScatterPlotWidgetDefinition) HasColorByGroups() bool`

HasColorByGroups returns a boolean if a field has been set.

### GetRequests

`func (o *ScatterPlotWidgetDefinition) GetRequests() ScatterPlotWidgetDefinitionRequests`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *ScatterPlotWidgetDefinition) GetRequestsOk() (*ScatterPlotWidgetDefinitionRequests, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *ScatterPlotWidgetDefinition) SetRequests(v ScatterPlotWidgetDefinitionRequests)`

SetRequests sets Requests field to given value.


### GetTime

`func (o *ScatterPlotWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ScatterPlotWidgetDefinition) GetTimeOk() (*WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ScatterPlotWidgetDefinition) SetTime(v WidgetTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *ScatterPlotWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTitle

`func (o *ScatterPlotWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ScatterPlotWidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ScatterPlotWidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ScatterPlotWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAlign

`func (o *ScatterPlotWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *ScatterPlotWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAlign

`func (o *ScatterPlotWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign sets TitleAlign field to given value.

### HasTitleAlign

`func (o *ScatterPlotWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### GetTitleSize

`func (o *ScatterPlotWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *ScatterPlotWidgetDefinition) GetTitleSizeOk() (*string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *ScatterPlotWidgetDefinition) SetTitleSize(v string)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *ScatterPlotWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetType

`func (o *ScatterPlotWidgetDefinition) GetType() ScatterPlotWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScatterPlotWidgetDefinition) GetTypeOk() (*ScatterPlotWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScatterPlotWidgetDefinition) SetType(v ScatterPlotWidgetDefinitionType)`

SetType sets Type field to given value.


### GetXaxis

`func (o *ScatterPlotWidgetDefinition) GetXaxis() WidgetAxis`

GetXaxis returns the Xaxis field if non-nil, zero value otherwise.

### GetXaxisOk

`func (o *ScatterPlotWidgetDefinition) GetXaxisOk() (*WidgetAxis, bool)`

GetXaxisOk returns a tuple with the Xaxis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXaxis

`func (o *ScatterPlotWidgetDefinition) SetXaxis(v WidgetAxis)`

SetXaxis sets Xaxis field to given value.

### HasXaxis

`func (o *ScatterPlotWidgetDefinition) HasXaxis() bool`

HasXaxis returns a boolean if a field has been set.

### GetYaxis

`func (o *ScatterPlotWidgetDefinition) GetYaxis() WidgetAxis`

GetYaxis returns the Yaxis field if non-nil, zero value otherwise.

### GetYaxisOk

`func (o *ScatterPlotWidgetDefinition) GetYaxisOk() (*WidgetAxis, bool)`

GetYaxisOk returns a tuple with the Yaxis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaxis

`func (o *ScatterPlotWidgetDefinition) SetYaxis(v WidgetAxis)`

SetYaxis sets Yaxis field to given value.

### HasYaxis

`func (o *ScatterPlotWidgetDefinition) HasYaxis() bool`

HasYaxis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


