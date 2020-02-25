# SLOWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShowErrorBudget** | Pointer to **bool** |  | [optional] 
**SloId** | Pointer to **string** |  | [optional] 
**TimeWindows** | Pointer to [**[]WidgetTimeWindows**](WidgetTimeWindows.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the widget | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title | [optional] 
**Type** | Pointer to **string** | Type of the widget | [readonly] [default to "slo"]
**ViewMode** | Pointer to [**WidgetViewMode**](WidgetViewMode.md) |  | [optional] 
**ViewType** | Pointer to **string** |  | [default to "detail"]

## Methods

### NewSLOWidgetDefinition

`func NewSLOWidgetDefinition(type_ string, viewType string, ) *SLOWidgetDefinition`

NewSLOWidgetDefinition instantiates a new SLOWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLOWidgetDefinitionWithDefaults

`func NewSLOWidgetDefinitionWithDefaults() *SLOWidgetDefinition`

NewSLOWidgetDefinitionWithDefaults instantiates a new SLOWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowErrorBudget

`func (o *SLOWidgetDefinition) GetShowErrorBudget() bool`

GetShowErrorBudget returns the ShowErrorBudget field if non-nil, zero value otherwise.

### GetShowErrorBudgetOk

`func (o *SLOWidgetDefinition) GetShowErrorBudgetOk() (bool, bool)`

GetShowErrorBudgetOk returns a tuple with the ShowErrorBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowErrorBudget

`func (o *SLOWidgetDefinition) HasShowErrorBudget() bool`

HasShowErrorBudget returns a boolean if a field has been set.

### SetShowErrorBudget

`func (o *SLOWidgetDefinition) SetShowErrorBudget(v bool)`

SetShowErrorBudget gets a reference to the given bool and assigns it to the ShowErrorBudget field.

### GetSloId

`func (o *SLOWidgetDefinition) GetSloId() string`

GetSloId returns the SloId field if non-nil, zero value otherwise.

### GetSloIdOk

`func (o *SLOWidgetDefinition) GetSloIdOk() (string, bool)`

GetSloIdOk returns a tuple with the SloId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSloId

`func (o *SLOWidgetDefinition) HasSloId() bool`

HasSloId returns a boolean if a field has been set.

### SetSloId

`func (o *SLOWidgetDefinition) SetSloId(v string)`

SetSloId gets a reference to the given string and assigns it to the SloId field.

### GetTimeWindows

`func (o *SLOWidgetDefinition) GetTimeWindows() []WidgetTimeWindows`

GetTimeWindows returns the TimeWindows field if non-nil, zero value otherwise.

### GetTimeWindowsOk

`func (o *SLOWidgetDefinition) GetTimeWindowsOk() ([]WidgetTimeWindows, bool)`

GetTimeWindowsOk returns a tuple with the TimeWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimeWindows

`func (o *SLOWidgetDefinition) HasTimeWindows() bool`

HasTimeWindows returns a boolean if a field has been set.

### SetTimeWindows

`func (o *SLOWidgetDefinition) SetTimeWindows(v []WidgetTimeWindows)`

SetTimeWindows gets a reference to the given []WidgetTimeWindows and assigns it to the TimeWindows field.

### GetTitle

`func (o *SLOWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SLOWidgetDefinition) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *SLOWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *SLOWidgetDefinition) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetTitleAlign

`func (o *SLOWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *SLOWidgetDefinition) GetTitleAlignOk() (WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitleAlign

`func (o *SLOWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### SetTitleAlign

`func (o *SLOWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.

### GetTitleSize

`func (o *SLOWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *SLOWidgetDefinition) GetTitleSizeOk() (string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitleSize

`func (o *SLOWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### SetTitleSize

`func (o *SLOWidgetDefinition) SetTitleSize(v string)`

SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.

### GetType

`func (o *SLOWidgetDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SLOWidgetDefinition) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *SLOWidgetDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *SLOWidgetDefinition) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetViewMode

`func (o *SLOWidgetDefinition) GetViewMode() WidgetViewMode`

GetViewMode returns the ViewMode field if non-nil, zero value otherwise.

### GetViewModeOk

`func (o *SLOWidgetDefinition) GetViewModeOk() (WidgetViewMode, bool)`

GetViewModeOk returns a tuple with the ViewMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasViewMode

`func (o *SLOWidgetDefinition) HasViewMode() bool`

HasViewMode returns a boolean if a field has been set.

### SetViewMode

`func (o *SLOWidgetDefinition) SetViewMode(v WidgetViewMode)`

SetViewMode gets a reference to the given WidgetViewMode and assigns it to the ViewMode field.

### GetViewType

`func (o *SLOWidgetDefinition) GetViewType() string`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *SLOWidgetDefinition) GetViewTypeOk() (string, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasViewType

`func (o *SLOWidgetDefinition) HasViewType() bool`

HasViewType returns a boolean if a field has been set.

### SetViewType

`func (o *SLOWidgetDefinition) SetViewType(v string)`

SetViewType gets a reference to the given string and assigns it to the ViewType field.


### AsWidgetDefinition

`func (s *SLOWidgetDefinition) AsWidgetDefinition() WidgetDefinition`

Convenience method to wrap this instance of SLOWidgetDefinition in WidgetDefinition

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


