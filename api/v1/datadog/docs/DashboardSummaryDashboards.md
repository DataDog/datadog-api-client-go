# DashboardSummaryDashboards

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorHandle** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsReadOnly** | Pointer to **bool** |  | [optional] 
**LayoutType** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewDashboardSummaryDashboards

`func NewDashboardSummaryDashboards() *DashboardSummaryDashboards`

NewDashboardSummaryDashboards instantiates a new DashboardSummaryDashboards object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardSummaryDashboardsWithDefaults

`func NewDashboardSummaryDashboardsWithDefaults() *DashboardSummaryDashboards`

NewDashboardSummaryDashboardsWithDefaults instantiates a new DashboardSummaryDashboards object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorHandle

`func (o *DashboardSummaryDashboards) GetAuthorHandle() string`

GetAuthorHandle returns the AuthorHandle field if non-nil, zero value otherwise.

### GetAuthorHandleOk

`func (o *DashboardSummaryDashboards) GetAuthorHandleOk() (string, bool)`

GetAuthorHandleOk returns a tuple with the AuthorHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAuthorHandle

`func (o *DashboardSummaryDashboards) HasAuthorHandle() bool`

HasAuthorHandle returns a boolean if a field has been set.

### SetAuthorHandle

`func (o *DashboardSummaryDashboards) SetAuthorHandle(v string)`

SetAuthorHandle gets a reference to the given string and assigns it to the AuthorHandle field.

### GetCreatedAt

`func (o *DashboardSummaryDashboards) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DashboardSummaryDashboards) GetCreatedAtOk() (time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedAt

`func (o *DashboardSummaryDashboards) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAt

`func (o *DashboardSummaryDashboards) SetCreatedAt(v time.Time)`

SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.

### GetDescription

`func (o *DashboardSummaryDashboards) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardSummaryDashboards) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *DashboardSummaryDashboards) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *DashboardSummaryDashboards) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetId

`func (o *DashboardSummaryDashboards) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardSummaryDashboards) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *DashboardSummaryDashboards) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *DashboardSummaryDashboards) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetIsReadOnly

`func (o *DashboardSummaryDashboards) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *DashboardSummaryDashboards) GetIsReadOnlyOk() (bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsReadOnly

`func (o *DashboardSummaryDashboards) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### SetIsReadOnly

`func (o *DashboardSummaryDashboards) SetIsReadOnly(v bool)`

SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.

### GetLayoutType

`func (o *DashboardSummaryDashboards) GetLayoutType() string`

GetLayoutType returns the LayoutType field if non-nil, zero value otherwise.

### GetLayoutTypeOk

`func (o *DashboardSummaryDashboards) GetLayoutTypeOk() (string, bool)`

GetLayoutTypeOk returns a tuple with the LayoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLayoutType

`func (o *DashboardSummaryDashboards) HasLayoutType() bool`

HasLayoutType returns a boolean if a field has been set.

### SetLayoutType

`func (o *DashboardSummaryDashboards) SetLayoutType(v string)`

SetLayoutType gets a reference to the given string and assigns it to the LayoutType field.

### GetModifiedAt

`func (o *DashboardSummaryDashboards) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *DashboardSummaryDashboards) GetModifiedAtOk() (time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModifiedAt

`func (o *DashboardSummaryDashboards) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAt

`func (o *DashboardSummaryDashboards) SetModifiedAt(v time.Time)`

SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.

### GetTitle

`func (o *DashboardSummaryDashboards) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DashboardSummaryDashboards) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *DashboardSummaryDashboards) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *DashboardSummaryDashboards) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetUrl

`func (o *DashboardSummaryDashboards) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DashboardSummaryDashboards) GetUrlOk() (string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUrl

`func (o *DashboardSummaryDashboards) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrl

`func (o *DashboardSummaryDashboards) SetUrl(v string)`

SetUrl gets a reference to the given string and assigns it to the Url field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


