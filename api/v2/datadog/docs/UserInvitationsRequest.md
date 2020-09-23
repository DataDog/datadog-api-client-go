# UserInvitationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]UserInvitationData**](UserInvitationData.md) | List of user invitations. | 

## Methods

### NewUserInvitationsRequest

`func NewUserInvitationsRequest(data []UserInvitationData, ) *UserInvitationsRequest`

NewUserInvitationsRequest instantiates a new UserInvitationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInvitationsRequestWithDefaults

`func NewUserInvitationsRequestWithDefaults() *UserInvitationsRequest`

NewUserInvitationsRequestWithDefaults instantiates a new UserInvitationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UserInvitationsRequest) GetData() []UserInvitationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserInvitationsRequest) GetDataOk() (*[]UserInvitationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserInvitationsRequest) SetData(v []UserInvitationData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


