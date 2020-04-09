# GroupWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LayoutType** | Pointer to [**WidgetLayoutType**](WidgetLayoutType.md) |  | 
**Title** | Pointer to **string** | Title of the widget | [optional] 
**Type** | Pointer to **string** | Type of the widget | [readonly] [default to "group"]
**Widgets** | Pointer to [**[]Widget**](Widget.md) | TODO. | 

## Methods

### NewGroupWidgetDefinition

`func NewGroupWidgetDefinition(layoutType WidgetLayoutType, type_ string, widgets []Widget, ) *GroupWidgetDefinition`

NewGroupWidgetDefinition instantiates a new GroupWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWidgetDefinitionWithDefaults

`func NewGroupWidgetDefinitionWithDefaults() *GroupWidgetDefinition`

NewGroupWidgetDefinitionWithDefaults instantiates a new GroupWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### GetType

`func (o *GroupWidgetDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupWidgetDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupWidgetDefinition) SetType(v string)`

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



### AsWidgetDefinition

`func (s *GroupWidgetDefinition) AsWidgetDefinition() WidgetDefinition`

Convenience method to wrap this instance of GroupWidgetDefinition in WidgetDefinition

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


