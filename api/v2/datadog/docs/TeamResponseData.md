# TeamResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**TeamResponseAttributes**](TeamResponseAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | The team&#39;s ID. | [optional] 
**Relationships** | Pointer to [**TeamRelationships**](TeamRelationships.md) |  | [optional] 
**Type** | Pointer to [**TeamType**](TeamType.md) |  | [optional] [default to "teams"]

## Methods

### NewTeamResponseData

`func NewTeamResponseData() *TeamResponseData`

NewTeamResponseData instantiates a new TeamResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamResponseDataWithDefaults

`func NewTeamResponseDataWithDefaults() *TeamResponseData`

NewTeamResponseDataWithDefaults instantiates a new TeamResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *TeamResponseData) GetAttributes() TeamResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TeamResponseData) GetAttributesOk() (*TeamResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TeamResponseData) SetAttributes(v TeamResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TeamResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *TeamResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *TeamResponseData) GetRelationships() TeamRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TeamResponseData) GetRelationshipsOk() (*TeamRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TeamResponseData) SetRelationships(v TeamRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TeamResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *TeamResponseData) GetType() TeamType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TeamResponseData) GetTypeOk() (*TeamType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TeamResponseData) SetType(v TeamType)`

SetType sets Type field to given value.

### HasType

`func (o *TeamResponseData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


