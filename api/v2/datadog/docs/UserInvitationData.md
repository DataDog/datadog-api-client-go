# UserInvitationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Relationships** | [**UserInvitationRelationships**](UserInvitationRelationships.md) |  | 
**Type** | [**UserInvitationsType**](UserInvitationsType.md) |  | [default to "user_invitations"]

## Methods

### NewUserInvitationData

`func NewUserInvitationData(relationships UserInvitationRelationships, type_ UserInvitationsType, ) *UserInvitationData`

NewUserInvitationData instantiates a new UserInvitationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInvitationDataWithDefaults

`func NewUserInvitationDataWithDefaults() *UserInvitationData`

NewUserInvitationDataWithDefaults instantiates a new UserInvitationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelationships

`func (o *UserInvitationData) GetRelationships() UserInvitationRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *UserInvitationData) GetRelationshipsOk() (*UserInvitationRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *UserInvitationData) SetRelationships(v UserInvitationRelationships)`

SetRelationships sets Relationships field to given value.


### GetType

`func (o *UserInvitationData) GetType() UserInvitationsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserInvitationData) GetTypeOk() (*UserInvitationsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserInvitationData) SetType(v UserInvitationsType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


