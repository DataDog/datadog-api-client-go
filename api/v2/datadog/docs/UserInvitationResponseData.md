# UserInvitationResponseData

## Properties

| Name           | Type                                                                           | Description                | Notes                                                        |
| -------------- | ------------------------------------------------------------------------------ | -------------------------- | ------------------------------------------------------------ |
| **Attributes** | Pointer to [**UserInvitationDataAttributes**](UserInvitationDataAttributes.md) |                            | [optional]                                                   |
| **Id**         | Pointer to **string**                                                          | ID of the user invitation. | [optional]                                                   |
| **Type**       | Pointer to [**UserInvitationsType**](UserInvitationsType.md)                   |                            | [optional] [default to USERINVITATIONSTYPE_USER_INVITATIONS] |

## Methods

### NewUserInvitationResponseData

`func NewUserInvitationResponseData() *UserInvitationResponseData`

NewUserInvitationResponseData instantiates a new UserInvitationResponseData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUserInvitationResponseDataWithDefaults

`func NewUserInvitationResponseDataWithDefaults() *UserInvitationResponseData`

NewUserInvitationResponseDataWithDefaults instantiates a new UserInvitationResponseData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *UserInvitationResponseData) GetAttributes() UserInvitationDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UserInvitationResponseData) GetAttributesOk() (*UserInvitationDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UserInvitationResponseData) SetAttributes(v UserInvitationDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UserInvitationResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *UserInvitationResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserInvitationResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserInvitationResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserInvitationResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserInvitationResponseData) GetType() UserInvitationsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserInvitationResponseData) GetTypeOk() (*UserInvitationsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserInvitationResponseData) SetType(v UserInvitationsType)`

SetType sets Type field to given value.

### HasType

`func (o *UserInvitationResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
