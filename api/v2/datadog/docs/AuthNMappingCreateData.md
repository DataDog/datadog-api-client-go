# AuthNMappingCreateData

## Properties

| Name              | Type                                                                                 | Description | Notes                                         |
| ----------------- | ------------------------------------------------------------------------------------ | ----------- | --------------------------------------------- |
| **Attributes**    | Pointer to [**AuthNMappingCreateAttributes**](AuthNMappingCreateAttributes.md)       |             | [optional]                                    |
| **Relationships** | Pointer to [**AuthNMappingCreateRelationships**](AuthNMappingCreateRelationships.md) |             | [optional]                                    |
| **Type**          | [**AuthNMappingsType**](AuthNMappingsType.md)                                        |             | [default to AUTHNMAPPINGSTYPE_AUTHN_MAPPINGS] |

## Methods

### NewAuthNMappingCreateData

`func NewAuthNMappingCreateData(type_ AuthNMappingsType) *AuthNMappingCreateData`

NewAuthNMappingCreateData instantiates a new AuthNMappingCreateData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAuthNMappingCreateDataWithDefaults

`func NewAuthNMappingCreateDataWithDefaults() *AuthNMappingCreateData`

NewAuthNMappingCreateDataWithDefaults instantiates a new AuthNMappingCreateData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *AuthNMappingCreateData) GetAttributes() AuthNMappingCreateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AuthNMappingCreateData) GetAttributesOk() (*AuthNMappingCreateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AuthNMappingCreateData) SetAttributes(v AuthNMappingCreateAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AuthNMappingCreateData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *AuthNMappingCreateData) GetRelationships() AuthNMappingCreateRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AuthNMappingCreateData) GetRelationshipsOk() (*AuthNMappingCreateRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AuthNMappingCreateData) SetRelationships(v AuthNMappingCreateRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AuthNMappingCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *AuthNMappingCreateData) GetType() AuthNMappingsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthNMappingCreateData) GetTypeOk() (*AuthNMappingsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthNMappingCreateData) SetType(v AuthNMappingsType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
