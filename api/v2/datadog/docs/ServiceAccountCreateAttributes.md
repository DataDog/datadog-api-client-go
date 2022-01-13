# ServiceAccountCreateAttributes

## Properties

| Name               | Type                  | Description                                          | Notes      |
| ------------------ | --------------------- | ---------------------------------------------------- | ---------- |
| **Email**          | **string**            | The email of the user.                               |
| **Name**           | Pointer to **string** | The name of the user.                                | [optional] |
| **ServiceAccount** | **bool**              | Whether the user is a service account. Must be true. |
| **Title**          | Pointer to **string** | The title of the user.                               | [optional] |

## Methods

### NewServiceAccountCreateAttributes

`func NewServiceAccountCreateAttributes(email string, serviceAccount bool) *ServiceAccountCreateAttributes`

NewServiceAccountCreateAttributes instantiates a new ServiceAccountCreateAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewServiceAccountCreateAttributesWithDefaults

`func NewServiceAccountCreateAttributesWithDefaults() *ServiceAccountCreateAttributes`

NewServiceAccountCreateAttributesWithDefaults instantiates a new ServiceAccountCreateAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetEmail

`func (o *ServiceAccountCreateAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ServiceAccountCreateAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ServiceAccountCreateAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### GetName

`func (o *ServiceAccountCreateAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceAccountCreateAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceAccountCreateAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceAccountCreateAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceAccount

`func (o *ServiceAccountCreateAttributes) GetServiceAccount() bool`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *ServiceAccountCreateAttributes) GetServiceAccountOk() (*bool, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *ServiceAccountCreateAttributes) SetServiceAccount(v bool)`

SetServiceAccount sets ServiceAccount field to given value.

### GetTitle

`func (o *ServiceAccountCreateAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ServiceAccountCreateAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ServiceAccountCreateAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ServiceAccountCreateAttributes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
