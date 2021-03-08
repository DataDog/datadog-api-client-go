# CheckStatusWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Check** | **string** | Name of the check to use in the widget. | 
**Group** | Pointer to **string** | Group reporting a single check. | [optional] 
**GroupBy** | Pointer to **[]string** | List of tag prefixes to group by in the case of a cluster check. | [optional] 
**Grouping** | [**WidgetGrouping**](WidgetGrouping.md) |  | 
**Tags** | Pointer to **[]string** | List of tags used to filter the groups reporting a cluster check. | [optional] 
**Time** | Pointer to [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the widget. | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title. | [optional] 
**Type** | [**CheckStatusWidgetDefinitionType**](CheckStatusWidgetDefinitionType.md) |  | [default to CHECKSTATUSWIDGETDEFINITIONTYPE_CHECK_STATUS]

## Methods

### NewCheckStatusWidgetDefinition

`func NewCheckStatusWidgetDefinition(check string, grouping WidgetGrouping, type_ CheckStatusWidgetDefinitionType, ) *CheckStatusWidgetDefinition`

NewCheckStatusWidgetDefinition instantiates a new CheckStatusWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckStatusWidgetDefinitionWithDefaults

`func NewCheckStatusWidgetDefinitionWithDefaults() *CheckStatusWidgetDefinition`

NewCheckStatusWidgetDefinitionWithDefaults instantiates a new CheckStatusWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheck

`func (o *CheckStatusWidgetDefinition) GetCheck() string`

GetCheck returns the Check field if non-nil, zero value otherwise.

### GetCheckOk

`func (o *CheckStatusWidgetDefinition) GetCheckOk() (*string, bool)`

GetCheckOk returns a tuple with the Check field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheck

`func (o *CheckStatusWidgetDefinition) SetCheck(v string)`

SetCheck sets Check field to given value.


### GetGroup

`func (o *CheckStatusWidgetDefinition) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CheckStatusWidgetDefinition) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CheckStatusWidgetDefinition) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CheckStatusWidgetDefinition) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupBy

`func (o *CheckStatusWidgetDefinition) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *CheckStatusWidgetDefinition) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *CheckStatusWidgetDefinition) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *CheckStatusWidgetDefinition) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetGrouping

`func (o *CheckStatusWidgetDefinition) GetGrouping() WidgetGrouping`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *CheckStatusWidgetDefinition) GetGroupingOk() (*WidgetGrouping, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *CheckStatusWidgetDefinition) SetGrouping(v WidgetGrouping)`

SetGrouping sets Grouping field to given value.


### GetTags

`func (o *CheckStatusWidgetDefinition) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CheckStatusWidgetDefinition) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CheckStatusWidgetDefinition) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CheckStatusWidgetDefinition) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTime

`func (o *CheckStatusWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CheckStatusWidgetDefinition) GetTimeOk() (*WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CheckStatusWidgetDefinition) SetTime(v WidgetTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *CheckStatusWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTitle

`func (o *CheckStatusWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CheckStatusWidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CheckStatusWidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CheckStatusWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAlign

`func (o *CheckStatusWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *CheckStatusWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAlign

`func (o *CheckStatusWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign sets TitleAlign field to given value.

### HasTitleAlign

`func (o *CheckStatusWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### GetTitleSize

`func (o *CheckStatusWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *CheckStatusWidgetDefinition) GetTitleSizeOk() (*string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *CheckStatusWidgetDefinition) SetTitleSize(v string)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *CheckStatusWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetType

`func (o *CheckStatusWidgetDefinition) GetType() CheckStatusWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckStatusWidgetDefinition) GetTypeOk() (*CheckStatusWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckStatusWidgetDefinition) SetType(v CheckStatusWidgetDefinitionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


