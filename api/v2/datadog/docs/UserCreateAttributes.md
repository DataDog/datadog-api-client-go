# UserCreateAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email of the user. | 
**Name** | Pointer to **string** | The name of the user. | [optional] 
**Title** | Pointer to **string** | The title of the user. | [optional] 

## Methods

### NewUserCreateAttributes

`func NewUserCreateAttributes(email string, ) *UserCreateAttributes`

NewUserCreateAttributes instantiates a new UserCreateAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateAttributesWithDefaults

`func NewUserCreateAttributesWithDefaults() *UserCreateAttributes`

NewUserCreateAttributesWithDefaults instantiates a new UserCreateAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserCreateAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCreateAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCreateAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *UserCreateAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserCreateAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserCreateAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserCreateAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *UserCreateAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UserCreateAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UserCreateAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UserCreateAttributes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


