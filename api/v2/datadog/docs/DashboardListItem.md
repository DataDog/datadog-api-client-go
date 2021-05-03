# DashboardListItem

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Author** | Pointer to [**Creator**](Creator.md) |  | [optional] 
**Created** | Pointer to **time.Time** | Date of creation of the dashboard. | [optional] [readonly] 
**Icon** | Pointer to **string** | URL to the icon of the dashboard. | [optional] [readonly] 
**Id** | **string** | ID of the dashboard. | 
**IsFavorite** | Pointer to **bool** | Whether or not the dashboard is in the favorites. | [optional] [readonly] 
**IsReadOnly** | Pointer to **bool** | Whether or not the dashboard is read only. | [optional] [readonly] 
**IsShared** | Pointer to **bool** | Whether the dashboard is publicly shared or not. | [optional] [readonly] 
**Modified** | Pointer to **time.Time** | Date of last edition of the dashboard. | [optional] [readonly] 
**Popularity** | Pointer to **int32** | Popularity of the dashboard. | [optional] [readonly] 
**Title** | Pointer to **string** | Title of the dashboard. | [optional] [readonly] 
**Type** | [**DashboardType**](DashboardType.md) |  | 
**Url** | Pointer to **string** | URL path to the dashboard. | [optional] [readonly] 

## Methods

### NewDashboardListItem

`func NewDashboardListItem(id string, type_ DashboardType, ) *DashboardListItem`

NewDashboardListItem instantiates a new DashboardListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardListItemWithDefaults

`func NewDashboardListItemWithDefaults() *DashboardListItem`

NewDashboardListItemWithDefaults instantiates a new DashboardListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *DashboardListItem) GetAuthor() Creator`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *DashboardListItem) GetAuthorOk() (*Creator, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *DashboardListItem) SetAuthor(v Creator)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *DashboardListItem) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCreated

`func (o *DashboardListItem) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DashboardListItem) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DashboardListItem) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DashboardListItem) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetIcon

`func (o *DashboardListItem) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *DashboardListItem) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *DashboardListItem) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *DashboardListItem) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *DashboardListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardListItem) SetId(v string)`

SetId sets Id field to given value.


### GetIsFavorite

`func (o *DashboardListItem) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *DashboardListItem) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *DashboardListItem) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *DashboardListItem) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *DashboardListItem) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *DashboardListItem) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *DashboardListItem) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *DashboardListItem) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetIsShared

`func (o *DashboardListItem) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *DashboardListItem) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *DashboardListItem) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *DashboardListItem) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetModified

`func (o *DashboardListItem) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *DashboardListItem) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *DashboardListItem) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *DashboardListItem) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetPopularity

`func (o *DashboardListItem) GetPopularity() int32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *DashboardListItem) GetPopularityOk() (*int32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *DashboardListItem) SetPopularity(v int32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *DashboardListItem) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetTitle

`func (o *DashboardListItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DashboardListItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DashboardListItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DashboardListItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *DashboardListItem) GetType() DashboardType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DashboardListItem) GetTypeOk() (*DashboardType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DashboardListItem) SetType(v DashboardType)`

SetType sets Type field to given value.


### GetUrl

`func (o *DashboardListItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DashboardListItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DashboardListItem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DashboardListItem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


