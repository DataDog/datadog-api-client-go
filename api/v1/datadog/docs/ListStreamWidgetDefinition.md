# ListStreamWidgetDefinition

## Properties

| Name           | Type                                                                    | Description                                                                                                                                                        | Notes                                                   |
| -------------- | ----------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------- |
| **LegendSize** | Pointer to **string**                                                   | Available legend sizes for a widget. Should be one of \&quot;0\&quot;, \&quot;2\&quot;, \&quot;4\&quot;, \&quot;8\&quot;, \&quot;16\&quot;, or \&quot;auto\&quot;. | [optional]                                              |
| **Requests**   | [**[]ListStreamWidgetRequest**](ListStreamWidgetRequest.md)             | Request payload used to query items.                                                                                                                               |
| **ShowLegend** | Pointer to **bool**                                                     | Whether or not to display the legend on this widget.                                                                                                               | [optional]                                              |
| **Time**       | Pointer to [**WidgetTime**](WidgetTime.md)                              |                                                                                                                                                                    | [optional]                                              |
| **Title**      | Pointer to **string**                                                   | Title of the widget.                                                                                                                                               | [optional]                                              |
| **TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md)                    |                                                                                                                                                                    | [optional]                                              |
| **TitleSize**  | Pointer to **string**                                                   | Size of the title.                                                                                                                                                 | [optional]                                              |
| **Type**       | [**ListStreamWidgetDefinitionType**](ListStreamWidgetDefinitionType.md) |                                                                                                                                                                    | [default to LISTSTREAMWIDGETDEFINITIONTYPE_LIST_STREAM] |

## Methods

### NewListStreamWidgetDefinition

`func NewListStreamWidgetDefinition(requests []ListStreamWidgetRequest, type_ ListStreamWidgetDefinitionType) *ListStreamWidgetDefinition`

NewListStreamWidgetDefinition instantiates a new ListStreamWidgetDefinition object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewListStreamWidgetDefinitionWithDefaults

`func NewListStreamWidgetDefinitionWithDefaults() *ListStreamWidgetDefinition`

NewListStreamWidgetDefinitionWithDefaults instantiates a new ListStreamWidgetDefinition object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetLegendSize

`func (o *ListStreamWidgetDefinition) GetLegendSize() string`

GetLegendSize returns the LegendSize field if non-nil, zero value otherwise.

### GetLegendSizeOk

`func (o *ListStreamWidgetDefinition) GetLegendSizeOk() (*string, bool)`

GetLegendSizeOk returns a tuple with the LegendSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendSize

`func (o *ListStreamWidgetDefinition) SetLegendSize(v string)`

SetLegendSize sets LegendSize field to given value.

### HasLegendSize

`func (o *ListStreamWidgetDefinition) HasLegendSize() bool`

HasLegendSize returns a boolean if a field has been set.

### GetRequests

`func (o *ListStreamWidgetDefinition) GetRequests() []ListStreamWidgetRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *ListStreamWidgetDefinition) GetRequestsOk() (*[]ListStreamWidgetRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *ListStreamWidgetDefinition) SetRequests(v []ListStreamWidgetRequest)`

SetRequests sets Requests field to given value.

### GetShowLegend

`func (o *ListStreamWidgetDefinition) GetShowLegend() bool`

GetShowLegend returns the ShowLegend field if non-nil, zero value otherwise.

### GetShowLegendOk

`func (o *ListStreamWidgetDefinition) GetShowLegendOk() (*bool, bool)`

GetShowLegendOk returns a tuple with the ShowLegend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLegend

`func (o *ListStreamWidgetDefinition) SetShowLegend(v bool)`

SetShowLegend sets ShowLegend field to given value.

### HasShowLegend

`func (o *ListStreamWidgetDefinition) HasShowLegend() bool`

HasShowLegend returns a boolean if a field has been set.

### GetTime

`func (o *ListStreamWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ListStreamWidgetDefinition) GetTimeOk() (*WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ListStreamWidgetDefinition) SetTime(v WidgetTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *ListStreamWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTitle

`func (o *ListStreamWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListStreamWidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListStreamWidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ListStreamWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAlign

`func (o *ListStreamWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *ListStreamWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAlign

`func (o *ListStreamWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign sets TitleAlign field to given value.

### HasTitleAlign

`func (o *ListStreamWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### GetTitleSize

`func (o *ListStreamWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *ListStreamWidgetDefinition) GetTitleSizeOk() (*string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *ListStreamWidgetDefinition) SetTitleSize(v string)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *ListStreamWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetType

`func (o *ListStreamWidgetDefinition) GetType() ListStreamWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListStreamWidgetDefinition) GetTypeOk() (*ListStreamWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListStreamWidgetDefinition) SetType(v ListStreamWidgetDefinitionType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
