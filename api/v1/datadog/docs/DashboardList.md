# DashboardList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**Creator**](Creator.md) |  | [optional] 
**Created** | Pointer to [**time.Time**](time.Time.md) | Date of creation of the dashboard list. | [optional] [readonly] 
**DashboardCount** | Pointer to **int64** | The number of dashboards in the list. | [optional] [readonly] 
**Id** | Pointer to **int64** | The ID of the dashboard list. | [optional] [readonly] 
**IsFavorite** | Pointer to **bool** | Whether or not the list is in the favorites. | [optional] [readonly] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Date of last edition of the dashboard list. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the dashboard list. | 
**Type** | Pointer to **string** | The type of dashboard list. | [optional] [readonly] 

## Methods

### NewDashboardList

`func NewDashboardList(name string, ) *DashboardList`

NewDashboardList instantiates a new DashboardList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardListWithDefaults

`func NewDashboardListWithDefaults() *DashboardList`

NewDashboardListWithDefaults instantiates a new DashboardList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *DashboardList) GetAuthor() Creator`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *DashboardList) GetAuthorOk() (*Creator, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *DashboardList) SetAuthor(v Creator)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *DashboardList) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCreated

`func (o *DashboardList) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DashboardList) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DashboardList) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DashboardList) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDashboardCount

`func (o *DashboardList) GetDashboardCount() int64`

GetDashboardCount returns the DashboardCount field if non-nil, zero value otherwise.

### GetDashboardCountOk

`func (o *DashboardList) GetDashboardCountOk() (*int64, bool)`

GetDashboardCountOk returns a tuple with the DashboardCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardCount

`func (o *DashboardList) SetDashboardCount(v int64)`

SetDashboardCount sets DashboardCount field to given value.

### HasDashboardCount

`func (o *DashboardList) HasDashboardCount() bool`

HasDashboardCount returns a boolean if a field has been set.

### GetId

`func (o *DashboardList) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardList) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardList) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DashboardList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsFavorite

`func (o *DashboardList) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *DashboardList) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *DashboardList) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *DashboardList) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetModified

`func (o *DashboardList) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *DashboardList) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *DashboardList) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *DashboardList) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *DashboardList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardList) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *DashboardList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DashboardList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DashboardList) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DashboardList) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


