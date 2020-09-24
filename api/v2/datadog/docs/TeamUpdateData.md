# TeamUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**TeamUpdateAttributes**](TeamUpdateAttributes.md) |  | [optional] 
**Id** | **string** | The team&#39;s ID. | 
**Relationships** | Pointer to [**TeamRelationships**](TeamRelationships.md) |  | [optional] 
**Type** | [**TeamType**](TeamType.md) |  | [default to "teams"]

## Methods

### NewTeamUpdateData

`func NewTeamUpdateData(id string, type_ TeamType, ) *TeamUpdateData`

NewTeamUpdateData instantiates a new TeamUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamUpdateDataWithDefaults

`func NewTeamUpdateDataWithDefaults() *TeamUpdateData`

NewTeamUpdateDataWithDefaults instantiates a new TeamUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *TeamUpdateData) GetAttributes() TeamUpdateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TeamUpdateData) GetAttributesOk() (*TeamUpdateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TeamUpdateData) SetAttributes(v TeamUpdateAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TeamUpdateData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *TeamUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetRelationships

`func (o *TeamUpdateData) GetRelationships() TeamRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TeamUpdateData) GetRelationshipsOk() (*TeamRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TeamUpdateData) SetRelationships(v TeamRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TeamUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *TeamUpdateData) GetType() TeamType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TeamUpdateData) GetTypeOk() (*TeamType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TeamUpdateData) SetType(v TeamType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


