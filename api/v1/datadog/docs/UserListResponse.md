# UserListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]User**](User.md) | TODO. | [optional] 

## Methods

### NewUserListResponse

`func NewUserListResponse() *UserListResponse`

NewUserListResponse instantiates a new UserListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserListResponseWithDefaults

`func NewUserListResponseWithDefaults() *UserListResponse`

NewUserListResponseWithDefaults instantiates a new UserListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *UserListResponse) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UserListResponse) GetUsersOk() ([]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsers

`func (o *UserListResponse) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsers

`func (o *UserListResponse) SetUsers(v []User)`

SetUsers gets a reference to the given []User and assigns it to the Users field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


