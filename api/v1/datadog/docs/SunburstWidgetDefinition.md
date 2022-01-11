# SunburstWidgetDefinition

## Properties

| Name            | Type                                                                | Description                          | Notes                                              |
| --------------- | ------------------------------------------------------------------- | ------------------------------------ | -------------------------------------------------- |
| **CustomLinks** | Pointer to [**[]WidgetCustomLink**](WidgetCustomLink.md)            | List of custom links.                | [optional]                                         |
| **HideTotal**   | Pointer to **bool**                                                 | Show the total value in this widget. | [optional]                                         |
| **Legend**      | Pointer to [**SunburstWidgetLegend**](SunburstWidgetLegend.md)      |                                      | [optional]                                         |
| **Requests**    | [**[]SunburstWidgetRequest**](SunburstWidgetRequest.md)             | List of sunburst widget requests.    |
| **Time**        | Pointer to [**WidgetTime**](WidgetTime.md)                          |                                      | [optional]                                         |
| **Title**       | Pointer to **string**                                               | Title of your widget.                | [optional]                                         |
| **TitleAlign**  | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md)                |                                      | [optional]                                         |
| **TitleSize**   | Pointer to **string**                                               | Size of the title.                   | [optional]                                         |
| **Type**        | [**SunburstWidgetDefinitionType**](SunburstWidgetDefinitionType.md) |                                      | [default to SUNBURSTWIDGETDEFINITIONTYPE_SUNBURST] |

## Methods

### NewSunburstWidgetDefinition

`func NewSunburstWidgetDefinition(requests []SunburstWidgetRequest, type_ SunburstWidgetDefinitionType) *SunburstWidgetDefinition`

NewSunburstWidgetDefinition instantiates a new SunburstWidgetDefinition object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSunburstWidgetDefinitionWithDefaults

`func NewSunburstWidgetDefinitionWithDefaults() *SunburstWidgetDefinition`

NewSunburstWidgetDefinitionWithDefaults instantiates a new SunburstWidgetDefinition object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCustomLinks

`func (o *SunburstWidgetDefinition) GetCustomLinks() []WidgetCustomLink`

GetCustomLinks returns the CustomLinks field if non-nil, zero value otherwise.

### GetCustomLinksOk

`func (o *SunburstWidgetDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool)`

GetCustomLinksOk returns a tuple with the CustomLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLinks

`func (o *SunburstWidgetDefinition) SetCustomLinks(v []WidgetCustomLink)`

SetCustomLinks sets CustomLinks field to given value.

### HasCustomLinks

`func (o *SunburstWidgetDefinition) HasCustomLinks() bool`

HasCustomLinks returns a boolean if a field has been set.

### GetHideTotal

`func (o *SunburstWidgetDefinition) GetHideTotal() bool`

GetHideTotal returns the HideTotal field if non-nil, zero value otherwise.

### GetHideTotalOk

`func (o *SunburstWidgetDefinition) GetHideTotalOk() (*bool, bool)`

GetHideTotalOk returns a tuple with the HideTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideTotal

`func (o *SunburstWidgetDefinition) SetHideTotal(v bool)`

SetHideTotal sets HideTotal field to given value.

### HasHideTotal

`func (o *SunburstWidgetDefinition) HasHideTotal() bool`

HasHideTotal returns a boolean if a field has been set.

### GetLegend

`func (o *SunburstWidgetDefinition) GetLegend() SunburstWidgetLegend`

GetLegend returns the Legend field if non-nil, zero value otherwise.

### GetLegendOk

`func (o *SunburstWidgetDefinition) GetLegendOk() (*SunburstWidgetLegend, bool)`

GetLegendOk returns a tuple with the Legend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegend

`func (o *SunburstWidgetDefinition) SetLegend(v SunburstWidgetLegend)`

SetLegend sets Legend field to given value.

### HasLegend

`func (o *SunburstWidgetDefinition) HasLegend() bool`

HasLegend returns a boolean if a field has been set.

### GetRequests

`func (o *SunburstWidgetDefinition) GetRequests() []SunburstWidgetRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *SunburstWidgetDefinition) GetRequestsOk() (*[]SunburstWidgetRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *SunburstWidgetDefinition) SetRequests(v []SunburstWidgetRequest)`

SetRequests sets Requests field to given value.

### GetTime

`func (o *SunburstWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *SunburstWidgetDefinition) GetTimeOk() (*WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *SunburstWidgetDefinition) SetTime(v WidgetTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *SunburstWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTitle

`func (o *SunburstWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SunburstWidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SunburstWidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SunburstWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAlign

`func (o *SunburstWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *SunburstWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAlign

`func (o *SunburstWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign sets TitleAlign field to given value.

### HasTitleAlign

`func (o *SunburstWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### GetTitleSize

`func (o *SunburstWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *SunburstWidgetDefinition) GetTitleSizeOk() (*string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *SunburstWidgetDefinition) SetTitleSize(v string)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *SunburstWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetType

`func (o *SunburstWidgetDefinition) GetType() SunburstWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SunburstWidgetDefinition) GetTypeOk() (*SunburstWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SunburstWidgetDefinition) SetType(v SunburstWidgetDefinitionType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
