# GroupWidgetDefinition

## Properties

| Name                | Type                                                          | Description                                        | Notes                                        |
| ------------------- | ------------------------------------------------------------- | -------------------------------------------------- | -------------------------------------------- |
| **BackgroundColor** | Pointer to **string**                                         | Background color of the group title.               | [optional]                                   |
| **BannerImg**       | Pointer to **string**                                         | URL of image to display as a banner for the group. | [optional]                                   |
| **LayoutType**      | [**WidgetLayoutType**](WidgetLayoutType.md)                   |                                                    |
| **ShowTitle**       | Pointer to **bool**                                           | Whether to show the title or not.                  | [optional] [default to true]                 |
| **Title**           | Pointer to **string**                                         | Title of the widget.                               | [optional]                                   |
| **TitleAlign**      | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md)          |                                                    | [optional]                                   |
| **Type**            | [**GroupWidgetDefinitionType**](GroupWidgetDefinitionType.md) |                                                    | [default to GROUPWIDGETDEFINITIONTYPE_GROUP] |
| **Widgets**         | [**[]Widget**](Widget.md)                                     | List of widget groups.                             |

## Methods

### NewGroupWidgetDefinition

`func NewGroupWidgetDefinition(layoutType WidgetLayoutType, type_ GroupWidgetDefinitionType, widgets []Widget) *GroupWidgetDefinition`

NewGroupWidgetDefinition instantiates a new GroupWidgetDefinition object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewGroupWidgetDefinitionWithDefaults

`func NewGroupWidgetDefinitionWithDefaults() *GroupWidgetDefinition`

NewGroupWidgetDefinitionWithDefaults instantiates a new GroupWidgetDefinition object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetBackgroundColor

`func (o *GroupWidgetDefinition) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *GroupWidgetDefinition) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *GroupWidgetDefinition) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *GroupWidgetDefinition) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### GetBannerImg

`func (o *GroupWidgetDefinition) GetBannerImg() string`

GetBannerImg returns the BannerImg field if non-nil, zero value otherwise.

### GetBannerImgOk

`func (o *GroupWidgetDefinition) GetBannerImgOk() (*string, bool)`

GetBannerImgOk returns a tuple with the BannerImg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerImg

`func (o *GroupWidgetDefinition) SetBannerImg(v string)`

SetBannerImg sets BannerImg field to given value.

### HasBannerImg

`func (o *GroupWidgetDefinition) HasBannerImg() bool`

HasBannerImg returns a boolean if a field has been set.

### GetLayoutType

`func (o *GroupWidgetDefinition) GetLayoutType() WidgetLayoutType`

GetLayoutType returns the LayoutType field if non-nil, zero value otherwise.

### GetLayoutTypeOk

`func (o *GroupWidgetDefinition) GetLayoutTypeOk() (*WidgetLayoutType, bool)`

GetLayoutTypeOk returns a tuple with the LayoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutType

`func (o *GroupWidgetDefinition) SetLayoutType(v WidgetLayoutType)`

SetLayoutType sets LayoutType field to given value.

### GetShowTitle

`func (o *GroupWidgetDefinition) GetShowTitle() bool`

GetShowTitle returns the ShowTitle field if non-nil, zero value otherwise.

### GetShowTitleOk

`func (o *GroupWidgetDefinition) GetShowTitleOk() (*bool, bool)`

GetShowTitleOk returns a tuple with the ShowTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTitle

`func (o *GroupWidgetDefinition) SetShowTitle(v bool)`

SetShowTitle sets ShowTitle field to given value.

### HasShowTitle

`func (o *GroupWidgetDefinition) HasShowTitle() bool`

HasShowTitle returns a boolean if a field has been set.

### GetTitle

`func (o *GroupWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GroupWidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GroupWidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GroupWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAlign

`func (o *GroupWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *GroupWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAlign

`func (o *GroupWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign sets TitleAlign field to given value.

### HasTitleAlign

`func (o *GroupWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### GetType

`func (o *GroupWidgetDefinition) GetType() GroupWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupWidgetDefinition) GetTypeOk() (*GroupWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupWidgetDefinition) SetType(v GroupWidgetDefinitionType)`

SetType sets Type field to given value.

### GetWidgets

`func (o *GroupWidgetDefinition) GetWidgets() []Widget`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *GroupWidgetDefinition) GetWidgetsOk() (*[]Widget, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *GroupWidgetDefinition) SetWidgets(v []Widget)`

SetWidgets sets Widgets field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
