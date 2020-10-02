# TableWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomLinks** | Pointer to [**[]WidgetCustomLink**](WidgetCustomLink.md) | List of custom links. | [optional] 
**HasSearchBar** | Pointer to [**TableWidgetHasSearchBar**](TableWidgetHasSearchBar.md) |  | [optional] [default to "auto"]
**Requests** | [**[]TableWidgetRequest**](TableWidgetRequest.md) | Widget definition. | 
**Time** | Pointer to [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | Pointer to **string** | Title of your widget. | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title. | [optional] 
**Type** | [**TableWidgetDefinitionType**](TableWidgetDefinitionType.md) |  | [default to "query_table"]

## Methods

### NewTableWidgetDefinition

`func NewTableWidgetDefinition(requests []TableWidgetRequest, type_ TableWidgetDefinitionType, ) *TableWidgetDefinition`

NewTableWidgetDefinition instantiates a new TableWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableWidgetDefinitionWithDefaults

`func NewTableWidgetDefinitionWithDefaults() *TableWidgetDefinition`

NewTableWidgetDefinitionWithDefaults instantiates a new TableWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomLinks

`func (o *TableWidgetDefinition) GetCustomLinks() []WidgetCustomLink`

GetCustomLinks returns the CustomLinks field if non-nil, zero value otherwise.

### GetCustomLinksOk

`func (o *TableWidgetDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool)`

GetCustomLinksOk returns a tuple with the CustomLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLinks

`func (o *TableWidgetDefinition) SetCustomLinks(v []WidgetCustomLink)`

SetCustomLinks sets CustomLinks field to given value.

### HasCustomLinks

`func (o *TableWidgetDefinition) HasCustomLinks() bool`

HasCustomLinks returns a boolean if a field has been set.

### GetHasSearchBar

`func (o *TableWidgetDefinition) GetHasSearchBar() TableWidgetHasSearchBar`

GetHasSearchBar returns the HasSearchBar field if non-nil, zero value otherwise.

### GetHasSearchBarOk

`func (o *TableWidgetDefinition) GetHasSearchBarOk() (*TableWidgetHasSearchBar, bool)`

GetHasSearchBarOk returns a tuple with the HasSearchBar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSearchBar

`func (o *TableWidgetDefinition) SetHasSearchBar(v TableWidgetHasSearchBar)`

SetHasSearchBar sets HasSearchBar field to given value.

### HasHasSearchBar

`func (o *TableWidgetDefinition) HasHasSearchBar() bool`

HasHasSearchBar returns a boolean if a field has been set.

### GetRequests

`func (o *TableWidgetDefinition) GetRequests() []TableWidgetRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *TableWidgetDefinition) GetRequestsOk() (*[]TableWidgetRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *TableWidgetDefinition) SetRequests(v []TableWidgetRequest)`

SetRequests sets Requests field to given value.


### GetTime

`func (o *TableWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *TableWidgetDefinition) GetTimeOk() (*WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *TableWidgetDefinition) SetTime(v WidgetTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *TableWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTitle

`func (o *TableWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TableWidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TableWidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TableWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAlign

`func (o *TableWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *TableWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAlign

`func (o *TableWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign sets TitleAlign field to given value.

### HasTitleAlign

`func (o *TableWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### GetTitleSize

`func (o *TableWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *TableWidgetDefinition) GetTitleSizeOk() (*string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *TableWidgetDefinition) SetTitleSize(v string)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *TableWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetType

`func (o *TableWidgetDefinition) GetType() TableWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TableWidgetDefinition) GetTypeOk() (*TableWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TableWidgetDefinition) SetType(v TableWidgetDefinitionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


