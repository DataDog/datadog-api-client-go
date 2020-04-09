# Dashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorHandle** | Pointer to **string** | TODO. | [optional] [readonly] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | TODO. | [optional] [readonly] 
**Description** | Pointer to **NullableString** | Description of the dashboard | [optional] 
**Id** | Pointer to **string** | ID of the dashboard | [optional] [readonly] 
**IsReadOnly** | Pointer to **bool** | Whether this dashboard is read-only. If True, only the author and admins can make changes to it. | [optional] [default to false]
**LayoutType** | Pointer to [**DashboardLayoutType**](DashboardLayoutType.md) |  | 
**ModifiedAt** | Pointer to [**time.Time**](time.Time.md) | TODO. | [optional] [readonly] 
**NotifyList** | Pointer to **[]string** | List of handles of users to notify when changes are made to this dashboard. | [optional] 
**TemplateVariablePresets** | Pointer to [**[]DashboardTemplateVariablePreset**](DashboardTemplateVariablePreset.md) | TODO. | [optional] 
**TemplateVariables** | Pointer to [**[]DashboardTemplateVariables**](Dashboard_template_variables.md) | TODO. | [optional] 
**Title** | Pointer to **string** | Title of the dashboard | 
**Url** | Pointer to **string** | TODO. | [optional] [readonly] 
**Widgets** | Pointer to [**[]Widget**](Widget.md) | List of widgets to display on the dashboard | 

## Methods

### NewDashboard

`func NewDashboard(layoutType DashboardLayoutType, title string, widgets []Widget, ) *Dashboard`

NewDashboard instantiates a new Dashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardWithDefaults

`func NewDashboardWithDefaults() *Dashboard`

NewDashboardWithDefaults instantiates a new Dashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorHandle

`func (o *Dashboard) GetAuthorHandle() string`

GetAuthorHandle returns the AuthorHandle field if non-nil, zero value otherwise.

### GetAuthorHandleOk

`func (o *Dashboard) GetAuthorHandleOk() (*string, bool)`

GetAuthorHandleOk returns a tuple with the AuthorHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorHandle

`func (o *Dashboard) SetAuthorHandle(v string)`

SetAuthorHandle sets AuthorHandle field to given value.

### HasAuthorHandle

`func (o *Dashboard) HasAuthorHandle() bool`

HasAuthorHandle returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Dashboard) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Dashboard) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Dashboard) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Dashboard) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *Dashboard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Dashboard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Dashboard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Dashboard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Dashboard) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Dashboard) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *Dashboard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dashboard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Dashboard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Dashboard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *Dashboard) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *Dashboard) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *Dashboard) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *Dashboard) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetLayoutType

`func (o *Dashboard) GetLayoutType() DashboardLayoutType`

GetLayoutType returns the LayoutType field if non-nil, zero value otherwise.

### GetLayoutTypeOk

`func (o *Dashboard) GetLayoutTypeOk() (*DashboardLayoutType, bool)`

GetLayoutTypeOk returns a tuple with the LayoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutType

`func (o *Dashboard) SetLayoutType(v DashboardLayoutType)`

SetLayoutType sets LayoutType field to given value.


### GetModifiedAt

`func (o *Dashboard) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Dashboard) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Dashboard) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Dashboard) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetNotifyList

`func (o *Dashboard) GetNotifyList() []string`

GetNotifyList returns the NotifyList field if non-nil, zero value otherwise.

### GetNotifyListOk

`func (o *Dashboard) GetNotifyListOk() (*[]string, bool)`

GetNotifyListOk returns a tuple with the NotifyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyList

`func (o *Dashboard) SetNotifyList(v []string)`

SetNotifyList sets NotifyList field to given value.

### HasNotifyList

`func (o *Dashboard) HasNotifyList() bool`

HasNotifyList returns a boolean if a field has been set.

### SetNotifyListNil

`func (o *Dashboard) SetNotifyListNil(b bool)`

 SetNotifyListNil sets the value for NotifyList to be an explicit nil

### UnsetNotifyList
`func (o *Dashboard) UnsetNotifyList()`

UnsetNotifyList ensures that no value is present for NotifyList, not even an explicit nil
### GetTemplateVariablePresets

`func (o *Dashboard) GetTemplateVariablePresets() []DashboardTemplateVariablePreset`

GetTemplateVariablePresets returns the TemplateVariablePresets field if non-nil, zero value otherwise.

### GetTemplateVariablePresetsOk

`func (o *Dashboard) GetTemplateVariablePresetsOk() (*[]DashboardTemplateVariablePreset, bool)`

GetTemplateVariablePresetsOk returns a tuple with the TemplateVariablePresets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVariablePresets

`func (o *Dashboard) SetTemplateVariablePresets(v []DashboardTemplateVariablePreset)`

SetTemplateVariablePresets sets TemplateVariablePresets field to given value.

### HasTemplateVariablePresets

`func (o *Dashboard) HasTemplateVariablePresets() bool`

HasTemplateVariablePresets returns a boolean if a field has been set.

### SetTemplateVariablePresetsNil

`func (o *Dashboard) SetTemplateVariablePresetsNil(b bool)`

 SetTemplateVariablePresetsNil sets the value for TemplateVariablePresets to be an explicit nil

### UnsetTemplateVariablePresets
`func (o *Dashboard) UnsetTemplateVariablePresets()`

UnsetTemplateVariablePresets ensures that no value is present for TemplateVariablePresets, not even an explicit nil
### GetTemplateVariables

`func (o *Dashboard) GetTemplateVariables() []DashboardTemplateVariables`

GetTemplateVariables returns the TemplateVariables field if non-nil, zero value otherwise.

### GetTemplateVariablesOk

`func (o *Dashboard) GetTemplateVariablesOk() (*[]DashboardTemplateVariables, bool)`

GetTemplateVariablesOk returns a tuple with the TemplateVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVariables

`func (o *Dashboard) SetTemplateVariables(v []DashboardTemplateVariables)`

SetTemplateVariables sets TemplateVariables field to given value.

### HasTemplateVariables

`func (o *Dashboard) HasTemplateVariables() bool`

HasTemplateVariables returns a boolean if a field has been set.

### SetTemplateVariablesNil

`func (o *Dashboard) SetTemplateVariablesNil(b bool)`

 SetTemplateVariablesNil sets the value for TemplateVariables to be an explicit nil

### UnsetTemplateVariables
`func (o *Dashboard) UnsetTemplateVariables()`

UnsetTemplateVariables ensures that no value is present for TemplateVariables, not even an explicit nil
### GetTitle

`func (o *Dashboard) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Dashboard) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Dashboard) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *Dashboard) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Dashboard) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Dashboard) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Dashboard) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWidgets

`func (o *Dashboard) GetWidgets() []Widget`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *Dashboard) GetWidgetsOk() (*[]Widget, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *Dashboard) SetWidgets(v []Widget)`

SetWidgets sets Widgets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


