# UserInvitationsResponse

## Properties

| Name     | Type                                                                         | Description                | Notes      |
| -------- | ---------------------------------------------------------------------------- | -------------------------- | ---------- |
| **Data** | Pointer to [**[]UserInvitationResponseData**](UserInvitationResponseData.md) | Array of user invitations. | [optional] |

## Methods

### NewUserInvitationsResponse

`func NewUserInvitationsResponse() *UserInvitationsResponse`

NewUserInvitationsResponse instantiates a new UserInvitationsResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUserInvitationsResponseWithDefaults

`func NewUserInvitationsResponseWithDefaults() *UserInvitationsResponse`

NewUserInvitationsResponseWithDefaults instantiates a new UserInvitationsResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *UserInvitationsResponse) GetData() []UserInvitationResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserInvitationsResponse) GetDataOk() (*[]UserInvitationResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserInvitationsResponse) SetData(v []UserInvitationResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *UserInvitationsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
