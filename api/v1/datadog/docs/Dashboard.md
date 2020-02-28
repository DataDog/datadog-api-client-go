# Dashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorHandle** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**Description** | Pointer to **NullableString** | Description of the dashboard | [optional] 
**Id** | Pointer to **string** | ID of the dashboard | [optional] [readonly] 
**IsReadOnly** | Pointer to **bool** | Whether this dashboard is read-only. If True, only the author and admins can make changes to it. | [optional] [default to false]
**LayoutType** | Pointer to [**DashboardLayoutType**](DashboardLayoutType.md) |  | 
**ModifiedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**NotifyList** | Pointer to **[]string** | List of handles of users to notify when changes are made to this dashboard. | [optional] 
**TemplateVariablePresets** | Pointer to [**[]DashboardTemplateVariablePreset**](DashboardTemplateVariablePreset.md) |  | [optional] 
**TemplateVariables** | Pointer to [**[]DashboardTemplateVariables**](Dashboard_template_variables.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the dashboard | 
**Url** | Pointer to **string** |  | [optional] [readonly] 
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

`func (o *Dashboard) GetAuthorHandleOk() (string, bool)`

GetAuthorHandleOk returns a tuple with the AuthorHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAuthorHandle

`func (o *Dashboard) HasAuthorHandle() bool`

HasAuthorHandle returns a boolean if a field has been set.

### SetAuthorHandle

`func (o *Dashboard) SetAuthorHandle(v string)`

SetAuthorHandle gets a reference to the given string and assigns it to the AuthorHandle field.

### GetCreatedAt

`func (o *Dashboard) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Dashboard) GetCreatedAtOk() (time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedAt

`func (o *Dashboard) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAt

`func (o *Dashboard) SetCreatedAt(v time.Time)`

SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.

### GetDescription

`func (o *Dashboard) GetDescription() NullableString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Dashboard) GetDescriptionOk() (NullableString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Dashboard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Dashboard) SetDescription(v NullableString)`

SetDescription gets a reference to the given NullableString and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *Dashboard) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetId

`func (o *Dashboard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dashboard) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Dashboard) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Dashboard) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetIsReadOnly

`func (o *Dashboard) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *Dashboard) GetIsReadOnlyOk() (bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsReadOnly

`func (o *Dashboard) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### SetIsReadOnly

`func (o *Dashboard) SetIsReadOnly(v bool)`

SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.

### GetLayoutType

`func (o *Dashboard) GetLayoutType() DashboardLayoutType`

GetLayoutType returns the LayoutType field if non-nil, zero value otherwise.

### GetLayoutTypeOk

`func (o *Dashboard) GetLayoutTypeOk() (DashboardLayoutType, bool)`

GetLayoutTypeOk returns a tuple with the LayoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLayoutType

`func (o *Dashboard) HasLayoutType() bool`

HasLayoutType returns a boolean if a field has been set.

### SetLayoutType

`func (o *Dashboard) SetLayoutType(v DashboardLayoutType)`

SetLayoutType gets a reference to the given DashboardLayoutType and assigns it to the LayoutType field.

### GetModifiedAt

`func (o *Dashboard) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Dashboard) GetModifiedAtOk() (time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModifiedAt

`func (o *Dashboard) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAt

`func (o *Dashboard) SetModifiedAt(v time.Time)`

SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.

### GetNotifyList

`func (o *Dashboard) GetNotifyList() []string`

GetNotifyList returns the NotifyList field if non-nil, zero value otherwise.

### GetNotifyListOk

`func (o *Dashboard) GetNotifyListOk() ([]string, bool)`

GetNotifyListOk returns a tuple with the NotifyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotifyList

`func (o *Dashboard) HasNotifyList() bool`

HasNotifyList returns a boolean if a field has been set.

### SetNotifyList

`func (o *Dashboard) SetNotifyList(v []string)`

SetNotifyList gets a reference to the given []string and assigns it to the NotifyList field.

### SetNotifyListExplicitNull

`func (o *Dashboard) SetNotifyListExplicitNull(b bool)`

SetNotifyListExplicitNull (un)sets NotifyList to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NotifyList value is set to nil even if false is passed
### GetTemplateVariablePresets

`func (o *Dashboard) GetTemplateVariablePresets() []DashboardTemplateVariablePreset`

GetTemplateVariablePresets returns the TemplateVariablePresets field if non-nil, zero value otherwise.

### GetTemplateVariablePresetsOk

`func (o *Dashboard) GetTemplateVariablePresetsOk() ([]DashboardTemplateVariablePreset, bool)`

GetTemplateVariablePresetsOk returns a tuple with the TemplateVariablePresets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTemplateVariablePresets

`func (o *Dashboard) HasTemplateVariablePresets() bool`

HasTemplateVariablePresets returns a boolean if a field has been set.

### SetTemplateVariablePresets

`func (o *Dashboard) SetTemplateVariablePresets(v []DashboardTemplateVariablePreset)`

SetTemplateVariablePresets gets a reference to the given []DashboardTemplateVariablePreset and assigns it to the TemplateVariablePresets field.

### SetTemplateVariablePresetsExplicitNull

`func (o *Dashboard) SetTemplateVariablePresetsExplicitNull(b bool)`

SetTemplateVariablePresetsExplicitNull (un)sets TemplateVariablePresets to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TemplateVariablePresets value is set to nil even if false is passed
### GetTemplateVariables

`func (o *Dashboard) GetTemplateVariables() []DashboardTemplateVariables`

GetTemplateVariables returns the TemplateVariables field if non-nil, zero value otherwise.

### GetTemplateVariablesOk

`func (o *Dashboard) GetTemplateVariablesOk() ([]DashboardTemplateVariables, bool)`

GetTemplateVariablesOk returns a tuple with the TemplateVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTemplateVariables

`func (o *Dashboard) HasTemplateVariables() bool`

HasTemplateVariables returns a boolean if a field has been set.

### SetTemplateVariables

`func (o *Dashboard) SetTemplateVariables(v []DashboardTemplateVariables)`

SetTemplateVariables gets a reference to the given []DashboardTemplateVariables and assigns it to the TemplateVariables field.

### SetTemplateVariablesExplicitNull

`func (o *Dashboard) SetTemplateVariablesExplicitNull(b bool)`

SetTemplateVariablesExplicitNull (un)sets TemplateVariables to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TemplateVariables value is set to nil even if false is passed
### GetTitle

`func (o *Dashboard) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Dashboard) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *Dashboard) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *Dashboard) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetUrl

`func (o *Dashboard) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Dashboard) GetUrlOk() (string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUrl

`func (o *Dashboard) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrl

`func (o *Dashboard) SetUrl(v string)`

SetUrl gets a reference to the given string and assigns it to the Url field.

### GetWidgets

`func (o *Dashboard) GetWidgets() []Widget`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *Dashboard) GetWidgetsOk() ([]Widget, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWidgets

`func (o *Dashboard) HasWidgets() bool`

HasWidgets returns a boolean if a field has been set.

### SetWidgets

`func (o *Dashboard) SetWidgets(v []Widget)`

SetWidgets gets a reference to the given []Widget and assigns it to the Widgets field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


