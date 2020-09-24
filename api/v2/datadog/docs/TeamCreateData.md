# TeamCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**TeamCreateAttributes**](TeamCreateAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**TeamRelationships**](TeamRelationships.md) |  | [optional] 
**Type** | [**TeamType**](TeamType.md) |  | [default to "teams"]

## Methods

### NewTeamCreateData

`func NewTeamCreateData(type_ TeamType, ) *TeamCreateData`

NewTeamCreateData instantiates a new TeamCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamCreateDataWithDefaults

`func NewTeamCreateDataWithDefaults() *TeamCreateData`

NewTeamCreateDataWithDefaults instantiates a new TeamCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *TeamCreateData) GetAttributes() TeamCreateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TeamCreateData) GetAttributesOk() (*TeamCreateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TeamCreateData) SetAttributes(v TeamCreateAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TeamCreateData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *TeamCreateData) GetRelationships() TeamRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TeamCreateData) GetRelationshipsOk() (*TeamRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TeamCreateData) SetRelationships(v TeamRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TeamCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *TeamCreateData) GetType() TeamType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TeamCreateData) GetTypeOk() (*TeamType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TeamCreateData) SetType(v TeamType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


