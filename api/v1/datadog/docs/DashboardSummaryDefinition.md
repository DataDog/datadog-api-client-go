# DashboardSummaryDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorHandle** | Pointer to **string** | Identifier of the dashboard author. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation date of the dashboard. | [optional] 
**Description** | Pointer to **string** | Description of the dashboard. | [optional] 
**Id** | Pointer to **string** | Dashboard identifier. | [optional] 
**IsReadOnly** | Pointer to **bool** | Whether this dashboard is read-only. If True, only the author and admins can make changes to it. | [optional] 
**LayoutType** | Pointer to [**DashboardLayoutType**](DashboardLayoutType.md) |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** | Modification date of the dashboard. | [optional] 
**Title** | Pointer to **string** | Title of the dashboard. | [optional] 
**Url** | Pointer to **string** | URL of the dashboard. | [optional] 

## Methods

### NewDashboardSummaryDefinition

`func NewDashboardSummaryDefinition() *DashboardSummaryDefinition`

NewDashboardSummaryDefinition instantiates a new DashboardSummaryDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardSummaryDefinitionWithDefaults

`func NewDashboardSummaryDefinitionWithDefaults() *DashboardSummaryDefinition`

NewDashboardSummaryDefinitionWithDefaults instantiates a new DashboardSummaryDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorHandle

`func (o *DashboardSummaryDefinition) GetAuthorHandle() string`

GetAuthorHandle returns the AuthorHandle field if non-nil, zero value otherwise.

### GetAuthorHandleOk

`func (o *DashboardSummaryDefinition) GetAuthorHandleOk() (*string, bool)`

GetAuthorHandleOk returns a tuple with the AuthorHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorHandle

`func (o *DashboardSummaryDefinition) SetAuthorHandle(v string)`

SetAuthorHandle sets AuthorHandle field to given value.

### HasAuthorHandle

`func (o *DashboardSummaryDefinition) HasAuthorHandle() bool`

HasAuthorHandle returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DashboardSummaryDefinition) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DashboardSummaryDefinition) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DashboardSummaryDefinition) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DashboardSummaryDefinition) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DashboardSummaryDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardSummaryDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardSummaryDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DashboardSummaryDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *DashboardSummaryDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardSummaryDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardSummaryDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DashboardSummaryDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *DashboardSummaryDefinition) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *DashboardSummaryDefinition) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *DashboardSummaryDefinition) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *DashboardSummaryDefinition) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetLayoutType

`func (o *DashboardSummaryDefinition) GetLayoutType() DashboardLayoutType`

GetLayoutType returns the LayoutType field if non-nil, zero value otherwise.

### GetLayoutTypeOk

`func (o *DashboardSummaryDefinition) GetLayoutTypeOk() (*DashboardLayoutType, bool)`

GetLayoutTypeOk returns a tuple with the LayoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutType

`func (o *DashboardSummaryDefinition) SetLayoutType(v DashboardLayoutType)`

SetLayoutType sets LayoutType field to given value.

### HasLayoutType

`func (o *DashboardSummaryDefinition) HasLayoutType() bool`

HasLayoutType returns a boolean if a field has been set.

### GetModifiedAt

`func (o *DashboardSummaryDefinition) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *DashboardSummaryDefinition) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *DashboardSummaryDefinition) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *DashboardSummaryDefinition) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetTitle

`func (o *DashboardSummaryDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DashboardSummaryDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DashboardSummaryDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DashboardSummaryDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *DashboardSummaryDefinition) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DashboardSummaryDefinition) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DashboardSummaryDefinition) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DashboardSummaryDefinition) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


