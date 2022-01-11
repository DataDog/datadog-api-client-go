# FunnelWidgetDefinition

## Properties

| Name           | Type                                                            | Description                          | Notes                                          |
| -------------- | --------------------------------------------------------------- | ------------------------------------ | ---------------------------------------------- |
| **Requests**   | [**[]FunnelWidgetRequest**](FunnelWidgetRequest.md)             | Request payload used to query items. |
| **Time**       | Pointer to [**WidgetTime**](WidgetTime.md)                      |                                      | [optional]                                     |
| **Title**      | Pointer to **string**                                           | The title of the widget.             | [optional]                                     |
| **TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md)            |                                      | [optional]                                     |
| **TitleSize**  | Pointer to **string**                                           | The size of the title.               | [optional]                                     |
| **Type**       | [**FunnelWidgetDefinitionType**](FunnelWidgetDefinitionType.md) |                                      | [default to FUNNELWIDGETDEFINITIONTYPE_FUNNEL] |

## Methods

### NewFunnelWidgetDefinition

`func NewFunnelWidgetDefinition(requests []FunnelWidgetRequest, type_ FunnelWidgetDefinitionType) *FunnelWidgetDefinition`

NewFunnelWidgetDefinition instantiates a new FunnelWidgetDefinition object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewFunnelWidgetDefinitionWithDefaults

`func NewFunnelWidgetDefinitionWithDefaults() *FunnelWidgetDefinition`

NewFunnelWidgetDefinitionWithDefaults instantiates a new FunnelWidgetDefinition object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetRequests

`func (o *FunnelWidgetDefinition) GetRequests() []FunnelWidgetRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *FunnelWidgetDefinition) GetRequestsOk() (*[]FunnelWidgetRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *FunnelWidgetDefinition) SetRequests(v []FunnelWidgetRequest)`

SetRequests sets Requests field to given value.

### GetTime

`func (o *FunnelWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *FunnelWidgetDefinition) GetTimeOk() (*WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *FunnelWidgetDefinition) SetTime(v WidgetTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *FunnelWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTitle

`func (o *FunnelWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FunnelWidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FunnelWidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FunnelWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAlign

`func (o *FunnelWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *FunnelWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAlign

`func (o *FunnelWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign sets TitleAlign field to given value.

### HasTitleAlign

`func (o *FunnelWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### GetTitleSize

`func (o *FunnelWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *FunnelWidgetDefinition) GetTitleSizeOk() (*string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *FunnelWidgetDefinition) SetTitleSize(v string)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *FunnelWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetType

`func (o *FunnelWidgetDefinition) GetType() FunnelWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunnelWidgetDefinition) GetTypeOk() (*FunnelWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunnelWidgetDefinition) SetType(v FunnelWidgetDefinitionType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
