# MonitorSummaryWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColorPreference** | Pointer to [**WidgetColorPreference**](WidgetColorPreference.md) |  | [optional] 
**Count** | Pointer to **int64** | The number of monitors to display. | [optional] 
**DisplayFormat** | Pointer to [**WidgetMonitorSummaryDisplayFormat**](WidgetMonitorSummaryDisplayFormat.md) |  | [optional] 
**HideZeroCounts** | Pointer to **bool** | Whether to show counts of 0 or not. | [optional] 
**Query** | **string** | Query to filter the monitors with. | 
**ShowLastTriggered** | Pointer to **bool** | Whether to show the time that has elapsed since the monitor/group triggered. | [optional] 
**Sort** | Pointer to [**WidgetMonitorSummarySort**](WidgetMonitorSummarySort.md) |  | [optional] 
**Start** | Pointer to **int64** | The start of the list. Typically 0. | [optional] 
**SummaryType** | Pointer to [**WidgetSummaryType**](WidgetSummaryType.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the widget. | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title. | [optional] 
**Type** | [**MonitorSummaryWidgetDefinitionType**](MonitorSummaryWidgetDefinitionType.md) |  | [default to MONITORSUMMARYWIDGETDEFINITIONTYPE_MANAGE_STATUS]

## Methods

### NewMonitorSummaryWidgetDefinition

`func NewMonitorSummaryWidgetDefinition(query string, type_ MonitorSummaryWidgetDefinitionType, ) *MonitorSummaryWidgetDefinition`

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

`func (o *MonitorSummaryWidgetDefinition) GetColorPreferenceOk() (*WidgetColorPreference, bool)`

GetColorPreferenceOk returns a tuple with the ColorPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorPreference

`func (o *MonitorSummaryWidgetDefinition) SetColorPreference(v WidgetColorPreference)`

SetColorPreference sets ColorPreference field to given value.

### HasColorPreference

`func (o *MonitorSummaryWidgetDefinition) HasColorPreference() bool`

HasColorPreference returns a boolean if a field has been set.

### GetCount

`func (o *MonitorSummaryWidgetDefinition) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MonitorSummaryWidgetDefinition) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MonitorSummaryWidgetDefinition) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *MonitorSummaryWidgetDefinition) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDisplayFormat

`func (o *MonitorSummaryWidgetDefinition) GetDisplayFormat() WidgetMonitorSummaryDisplayFormat`

GetDisplayFormat returns the DisplayFormat field if non-nil, zero value otherwise.

### GetDisplayFormatOk

`func (o *MonitorSummaryWidgetDefinition) GetDisplayFormatOk() (*WidgetMonitorSummaryDisplayFormat, bool)`

GetDisplayFormatOk returns a tuple with the DisplayFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayFormat

`func (o *MonitorSummaryWidgetDefinition) SetDisplayFormat(v WidgetMonitorSummaryDisplayFormat)`

SetDisplayFormat sets DisplayFormat field to given value.

### HasDisplayFormat

`func (o *MonitorSummaryWidgetDefinition) HasDisplayFormat() bool`

HasDisplayFormat returns a boolean if a field has been set.

### GetHideZeroCounts

`func (o *MonitorSummaryWidgetDefinition) GetHideZeroCounts() bool`

GetHideZeroCounts returns the HideZeroCounts field if non-nil, zero value otherwise.

### GetHideZeroCountsOk

`func (o *MonitorSummaryWidgetDefinition) GetHideZeroCountsOk() (*bool, bool)`

GetHideZeroCountsOk returns a tuple with the HideZeroCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideZeroCounts

`func (o *MonitorSummaryWidgetDefinition) SetHideZeroCounts(v bool)`

SetHideZeroCounts sets HideZeroCounts field to given value.

### HasHideZeroCounts

`func (o *MonitorSummaryWidgetDefinition) HasHideZeroCounts() bool`

HasHideZeroCounts returns a boolean if a field has been set.

### GetQuery

`func (o *MonitorSummaryWidgetDefinition) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MonitorSummaryWidgetDefinition) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MonitorSummaryWidgetDefinition) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetShowLastTriggered

`func (o *MonitorSummaryWidgetDefinition) GetShowLastTriggered() bool`

GetShowLastTriggered returns the ShowLastTriggered field if non-nil, zero value otherwise.

### GetShowLastTriggeredOk

`func (o *MonitorSummaryWidgetDefinition) GetShowLastTriggeredOk() (*bool, bool)`

GetShowLastTriggeredOk returns a tuple with the ShowLastTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLastTriggered

`func (o *MonitorSummaryWidgetDefinition) SetShowLastTriggered(v bool)`

SetShowLastTriggered sets ShowLastTriggered field to given value.

### HasShowLastTriggered

`func (o *MonitorSummaryWidgetDefinition) HasShowLastTriggered() bool`

HasShowLastTriggered returns a boolean if a field has been set.

### GetSort

`func (o *MonitorSummaryWidgetDefinition) GetSort() WidgetMonitorSummarySort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *MonitorSummaryWidgetDefinition) GetSortOk() (*WidgetMonitorSummarySort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *MonitorSummaryWidgetDefinition) SetSort(v WidgetMonitorSummarySort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *MonitorSummaryWidgetDefinition) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetStart

`func (o *MonitorSummaryWidgetDefinition) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *MonitorSummaryWidgetDefinition) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *MonitorSummaryWidgetDefinition) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *MonitorSummaryWidgetDefinition) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetSummaryType

`func (o *MonitorSummaryWidgetDefinition) GetSummaryType() WidgetSummaryType`

GetSummaryType returns the SummaryType field if non-nil, zero value otherwise.

### GetSummaryTypeOk

`func (o *MonitorSummaryWidgetDefinition) GetSummaryTypeOk() (*WidgetSummaryType, bool)`

GetSummaryTypeOk returns a tuple with the SummaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryType

`func (o *MonitorSummaryWidgetDefinition) SetSummaryType(v WidgetSummaryType)`

SetSummaryType sets SummaryType field to given value.

### HasSummaryType

`func (o *MonitorSummaryWidgetDefinition) HasSummaryType() bool`

HasSummaryType returns a boolean if a field has been set.

### GetTitle

`func (o *MonitorSummaryWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MonitorSummaryWidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MonitorSummaryWidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MonitorSummaryWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAlign

`func (o *MonitorSummaryWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *MonitorSummaryWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAlign

`func (o *MonitorSummaryWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign sets TitleAlign field to given value.

### HasTitleAlign

`func (o *MonitorSummaryWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### GetTitleSize

`func (o *MonitorSummaryWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *MonitorSummaryWidgetDefinition) GetTitleSizeOk() (*string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *MonitorSummaryWidgetDefinition) SetTitleSize(v string)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *MonitorSummaryWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetType

`func (o *MonitorSummaryWidgetDefinition) GetType() MonitorSummaryWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorSummaryWidgetDefinition) GetTypeOk() (*MonitorSummaryWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorSummaryWidgetDefinition) SetType(v MonitorSummaryWidgetDefinitionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


