# GeomapWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomLinks** | Pointer to [**[]WidgetCustomLink**](WidgetCustomLink.md) | A list of custom links. | [optional] 
**Requests** | [**[]GeomapWidgetRequest**](GeomapWidgetRequest.md) | Array of one request object to display in the widget. The request must contain a &#x60;group-by&#x60; tag whose value is a country ISO code.  See the [Request JSON schema documentation](https://docs.datadoghq.com/dashboards/graphing_json/request_json) for information about building the &#x60;REQUEST_SCHEMA&#x60;. | 
**Style** | [**GeomapWidgetDefinitionStyle**](GeomapWidgetDefinitionStyle.md) |  | 
**Time** | Pointer to [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | Pointer to **string** | The title of your widget. | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | The size of the title. | [optional] 
**Type** | [**GeomapWidgetDefinitionType**](GeomapWidgetDefinitionType.md) |  | [default to GEOMAPWIDGETDEFINITIONTYPE_GEOMAP]
**View** | [**GeomapWidgetDefinitionView**](GeomapWidgetDefinitionView.md) |  | 

## Methods

### NewGeomapWidgetDefinition

`func NewGeomapWidgetDefinition(requests []GeomapWidgetRequest, style GeomapWidgetDefinitionStyle, type_ GeomapWidgetDefinitionType, view GeomapWidgetDefinitionView, ) *GeomapWidgetDefinition`

NewGeomapWidgetDefinition instantiates a new GeomapWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeomapWidgetDefinitionWithDefaults

`func NewGeomapWidgetDefinitionWithDefaults() *GeomapWidgetDefinition`

NewGeomapWidgetDefinitionWithDefaults instantiates a new GeomapWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomLinks

`func (o *GeomapWidgetDefinition) GetCustomLinks() []WidgetCustomLink`

GetCustomLinks returns the CustomLinks field if non-nil, zero value otherwise.

### GetCustomLinksOk

`func (o *GeomapWidgetDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool)`

GetCustomLinksOk returns a tuple with the CustomLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLinks

`func (o *GeomapWidgetDefinition) SetCustomLinks(v []WidgetCustomLink)`

SetCustomLinks sets CustomLinks field to given value.

### HasCustomLinks

`func (o *GeomapWidgetDefinition) HasCustomLinks() bool`

HasCustomLinks returns a boolean if a field has been set.

### GetRequests

`func (o *GeomapWidgetDefinition) GetRequests() []GeomapWidgetRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *GeomapWidgetDefinition) GetRequestsOk() (*[]GeomapWidgetRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *GeomapWidgetDefinition) SetRequests(v []GeomapWidgetRequest)`

SetRequests sets Requests field to given value.


### GetStyle

`func (o *GeomapWidgetDefinition) GetStyle() GeomapWidgetDefinitionStyle`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *GeomapWidgetDefinition) GetStyleOk() (*GeomapWidgetDefinitionStyle, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *GeomapWidgetDefinition) SetStyle(v GeomapWidgetDefinitionStyle)`

SetStyle sets Style field to given value.


### GetTime

`func (o *GeomapWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GeomapWidgetDefinition) GetTimeOk() (*WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GeomapWidgetDefinition) SetTime(v WidgetTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *GeomapWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTitle

`func (o *GeomapWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GeomapWidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GeomapWidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GeomapWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAlign

`func (o *GeomapWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *GeomapWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAlign

`func (o *GeomapWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign sets TitleAlign field to given value.

### HasTitleAlign

`func (o *GeomapWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### GetTitleSize

`func (o *GeomapWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *GeomapWidgetDefinition) GetTitleSizeOk() (*string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *GeomapWidgetDefinition) SetTitleSize(v string)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *GeomapWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetType

`func (o *GeomapWidgetDefinition) GetType() GeomapWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GeomapWidgetDefinition) GetTypeOk() (*GeomapWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GeomapWidgetDefinition) SetType(v GeomapWidgetDefinitionType)`

SetType sets Type field to given value.


### GetView

`func (o *GeomapWidgetDefinition) GetView() GeomapWidgetDefinitionView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GeomapWidgetDefinition) GetViewOk() (*GeomapWidgetDefinitionView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GeomapWidgetDefinition) SetView(v GeomapWidgetDefinitionView)`

SetView sets View field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


