# MonitorSummaryWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColorPreference** | Pointer to [**WidgetColorPreference**](WidgetColorPreference.md) |  | [optional] 
**DisplayFormat** | Pointer to [**WidgetMonitorSummaryDisplayFormat**](WidgetMonitorSummaryDisplayFormat.md) |  | [optional] 
**HideZeroCounts** | Pointer to **bool** | Whether to show counts of 0 or not | [optional] 
**Query** | Pointer to **string** | Query to filter the monitors with | 
**ShowLastTriggered** | Pointer to **bool** | Whether to show the time that has elapsed since the monitor/group triggered | [optional] 
**Sort** | Pointer to [**WidgetSort**](WidgetSort.md) |  | [optional] 
**SummaryType** | Pointer to [**WidgetSummaryType**](WidgetSummaryType.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the widget | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title | [optional] 
**Type** | Pointer to **string** | Type of the widget | [readonly] [default to "manage_status"]

## Methods

### NewMonitorSummaryWidgetDefinition

`func NewMonitorSummaryWidgetDefinition(query string, type_ string, ) *MonitorSummaryWidgetDefinition`

NewMonitorSummaryWidgetDefinition instantiates a new MonitorSummaryWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorSummaryWidgetDefinitionWithDefaults

`func NewMonitorSummaryWidgetDefinitionWithDefaults() *MonitorSummaryWidgetDefinition`

NewMonitorSummaryWidgetDefinitionWithDefaults instantiates a new MonitorSummaryWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColorPreference

`func (o *MonitorSummaryWidgetDefinition) GetColorPreference() WidgetColorPreference`

GetColorPreference returns the ColorPreference field if non-nil, zero value otherwise.

### GetColorPreferenceOk

`func (o *MonitorSummaryWidgetDefinition) GetColorPreferenceOk() (WidgetColorPreference, bool)`

GetColorPreferenceOk returns a tuple with the ColorPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColorPreference

`func (o *MonitorSummaryWidgetDefinition) HasColorPreference() bool`

HasColorPreference returns a boolean if a field has been set.

### SetColorPreference

`func (o *MonitorSummaryWidgetDefinition) SetColorPreference(v WidgetColorPreference)`

SetColorPreference gets a reference to the given WidgetColorPreference and assigns it to the ColorPreference field.

### GetDisplayFormat

`func (o *MonitorSummaryWidgetDefinition) GetDisplayFormat() WidgetMonitorSummaryDisplayFormat`

GetDisplayFormat returns the DisplayFormat field if non-nil, zero value otherwise.

### GetDisplayFormatOk

`func (o *MonitorSummaryWidgetDefinition) GetDisplayFormatOk() (WidgetMonitorSummaryDisplayFormat, bool)`

GetDisplayFormatOk returns a tuple with the DisplayFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayFormat

`func (o *MonitorSummaryWidgetDefinition) HasDisplayFormat() bool`

HasDisplayFormat returns a boolean if a field has been set.

### SetDisplayFormat

`func (o *MonitorSummaryWidgetDefinition) SetDisplayFormat(v WidgetMonitorSummaryDisplayFormat)`

SetDisplayFormat gets a reference to the given WidgetMonitorSummaryDisplayFormat and assigns it to the DisplayFormat field.

### GetHideZeroCounts

`func (o *MonitorSummaryWidgetDefinition) GetHideZeroCounts() bool`

GetHideZeroCounts returns the HideZeroCounts field if non-nil, zero value otherwise.

### GetHideZeroCountsOk

`func (o *MonitorSummaryWidgetDefinition) GetHideZeroCountsOk() (bool, bool)`

GetHideZeroCountsOk returns a tuple with the HideZeroCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHideZeroCounts

`func (o *MonitorSummaryWidgetDefinition) HasHideZeroCounts() bool`

HasHideZeroCounts returns a boolean if a field has been set.

### SetHideZeroCounts

`func (o *MonitorSummaryWidgetDefinition) SetHideZeroCounts(v bool)`

SetHideZeroCounts gets a reference to the given bool and assigns it to the HideZeroCounts field.

### GetQuery

`func (o *MonitorSummaryWidgetDefinition) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MonitorSummaryWidgetDefinition) GetQueryOk() (string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQuery

`func (o *MonitorSummaryWidgetDefinition) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQuery

`func (o *MonitorSummaryWidgetDefinition) SetQuery(v string)`

SetQuery gets a reference to the given string and assigns it to the Query field.

### GetShowLastTriggered

`func (o *MonitorSummaryWidgetDefinition) GetShowLastTriggered() bool`

GetShowLastTriggered returns the ShowLastTriggered field if non-nil, zero value otherwise.

### GetShowLastTriggeredOk

`func (o *MonitorSummaryWidgetDefinition) GetShowLastTriggeredOk() (bool, bool)`

GetShowLastTriggeredOk returns a tuple with the ShowLastTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowLastTriggered

`func (o *MonitorSummaryWidgetDefinition) HasShowLastTriggered() bool`

HasShowLastTriggered returns a boolean if a field has been set.

### SetShowLastTriggered

`func (o *MonitorSummaryWidgetDefinition) SetShowLastTriggered(v bool)`

SetShowLastTriggered gets a reference to the given bool and assigns it to the ShowLastTriggered field.

### GetSort

`func (o *MonitorSummaryWidgetDefinition) GetSort() WidgetSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *MonitorSummaryWidgetDefinition) GetSortOk() (WidgetSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSort

`func (o *MonitorSummaryWidgetDefinition) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSort

`func (o *MonitorSummaryWidgetDefinition) SetSort(v WidgetSort)`

SetSort gets a reference to the given WidgetSort and assigns it to the Sort field.

### GetSummaryType

`func (o *MonitorSummaryWidgetDefinition) GetSummaryType() WidgetSummaryType`

GetSummaryType returns the SummaryType field if non-nil, zero value otherwise.

### GetSummaryTypeOk

`func (o *MonitorSummaryWidgetDefinition) GetSummaryTypeOk() (WidgetSummaryType, bool)`

GetSummaryTypeOk returns a tuple with the SummaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSummaryType

`func (o *MonitorSummaryWidgetDefinition) HasSummaryType() bool`

HasSummaryType returns a boolean if a field has been set.

### SetSummaryType

`func (o *MonitorSummaryWidgetDefinition) SetSummaryType(v WidgetSummaryType)`

SetSummaryType gets a reference to the given WidgetSummaryType and assigns it to the SummaryType field.

### GetTitle

`func (o *MonitorSummaryWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MonitorSummaryWidgetDefinition) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *MonitorSummaryWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *MonitorSummaryWidgetDefinition) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetTitleAlign

`func (o *MonitorSummaryWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *MonitorSummaryWidgetDefinition) GetTitleAlignOk() (WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitleAlign

`func (o *MonitorSummaryWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### SetTitleAlign

`func (o *MonitorSummaryWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.

### GetTitleSize

`func (o *MonitorSummaryWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *MonitorSummaryWidgetDefinition) GetTitleSizeOk() (string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitleSize

`func (o *MonitorSummaryWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### SetTitleSize

`func (o *MonitorSummaryWidgetDefinition) SetTitleSize(v string)`

SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.

### GetType

`func (o *MonitorSummaryWidgetDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorSummaryWidgetDefinition) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *MonitorSummaryWidgetDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *MonitorSummaryWidgetDefinition) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.


### AsWidgetDefinition

`func (s *MonitorSummaryWidgetDefinition) AsWidgetDefinition() WidgetDefinition`

Convenience method to wrap this instance of MonitorSummaryWidgetDefinition in WidgetDefinition

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


