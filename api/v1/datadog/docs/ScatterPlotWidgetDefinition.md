# ScatterPlotWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColorByGroups** | Pointer to **[]string** | List of groups used for colors. | [optional] 
**Requests** | Pointer to [**ScatterPlotWidgetDefinitionRequests**](ScatterPlotWidgetDefinition_requests.md) |  | 
**Time** | Pointer to [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | Pointer to **string** | Title of your widget. | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title | [optional] 
**Type** | Pointer to **string** | Type of the widget | [readonly] [default to "scatterplot"]
**Xaxis** | Pointer to [**WidgetAxis**](WidgetAxis.md) |  | [optional] 
**Yaxis** | Pointer to [**WidgetAxis**](WidgetAxis.md) |  | [optional] 

## Methods

### NewScatterPlotWidgetDefinition

`func NewScatterPlotWidgetDefinition(requests ScatterPlotWidgetDefinitionRequests, type_ string, ) *ScatterPlotWidgetDefinition`

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

`func (o *ScatterPlotWidgetDefinition) GetColorByGroupsOk() ([]string, bool)`

GetColorByGroupsOk returns a tuple with the ColorByGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColorByGroups

`func (o *ScatterPlotWidgetDefinition) HasColorByGroups() bool`

HasColorByGroups returns a boolean if a field has been set.

### SetColorByGroups

`func (o *ScatterPlotWidgetDefinition) SetColorByGroups(v []string)`

SetColorByGroups gets a reference to the given []string and assigns it to the ColorByGroups field.

### GetRequests

`func (o *ScatterPlotWidgetDefinition) GetRequests() ScatterPlotWidgetDefinitionRequests`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *ScatterPlotWidgetDefinition) GetRequestsOk() (ScatterPlotWidgetDefinitionRequests, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequests

`func (o *ScatterPlotWidgetDefinition) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### SetRequests

`func (o *ScatterPlotWidgetDefinition) SetRequests(v ScatterPlotWidgetDefinitionRequests)`

SetRequests gets a reference to the given ScatterPlotWidgetDefinitionRequests and assigns it to the Requests field.

### GetTime

`func (o *ScatterPlotWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ScatterPlotWidgetDefinition) GetTimeOk() (WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTime

`func (o *ScatterPlotWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTime

`func (o *ScatterPlotWidgetDefinition) SetTime(v WidgetTime)`

SetTime gets a reference to the given WidgetTime and assigns it to the Time field.

### GetTitle

`func (o *ScatterPlotWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ScatterPlotWidgetDefinition) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *ScatterPlotWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *ScatterPlotWidgetDefinition) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetTitleAlign

`func (o *ScatterPlotWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *ScatterPlotWidgetDefinition) GetTitleAlignOk() (WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitleAlign

`func (o *ScatterPlotWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### SetTitleAlign

`func (o *ScatterPlotWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.

### GetTitleSize

`func (o *ScatterPlotWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *ScatterPlotWidgetDefinition) GetTitleSizeOk() (string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitleSize

`func (o *ScatterPlotWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### SetTitleSize

`func (o *ScatterPlotWidgetDefinition) SetTitleSize(v string)`

SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.

### GetType

`func (o *ScatterPlotWidgetDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScatterPlotWidgetDefinition) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *ScatterPlotWidgetDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *ScatterPlotWidgetDefinition) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetXaxis

`func (o *ScatterPlotWidgetDefinition) GetXaxis() WidgetAxis`

GetXaxis returns the Xaxis field if non-nil, zero value otherwise.

### GetXaxisOk

`func (o *ScatterPlotWidgetDefinition) GetXaxisOk() (WidgetAxis, bool)`

GetXaxisOk returns a tuple with the Xaxis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasXaxis

`func (o *ScatterPlotWidgetDefinition) HasXaxis() bool`

HasXaxis returns a boolean if a field has been set.

### SetXaxis

`func (o *ScatterPlotWidgetDefinition) SetXaxis(v WidgetAxis)`

SetXaxis gets a reference to the given WidgetAxis and assigns it to the Xaxis field.

### GetYaxis

`func (o *ScatterPlotWidgetDefinition) GetYaxis() WidgetAxis`

GetYaxis returns the Yaxis field if non-nil, zero value otherwise.

### GetYaxisOk

`func (o *ScatterPlotWidgetDefinition) GetYaxisOk() (WidgetAxis, bool)`

GetYaxisOk returns a tuple with the Yaxis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasYaxis

`func (o *ScatterPlotWidgetDefinition) HasYaxis() bool`

HasYaxis returns a boolean if a field has been set.

### SetYaxis

`func (o *ScatterPlotWidgetDefinition) SetYaxis(v WidgetAxis)`

SetYaxis gets a reference to the given WidgetAxis and assigns it to the Yaxis field.


### AsWidgetDefinition

`func (s *ScatterPlotWidgetDefinition) AsWidgetDefinition() WidgetDefinition`

Convenience method to wrap this instance of ScatterPlotWidgetDefinition in WidgetDefinition

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


