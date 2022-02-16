# AuthNMappingUpdateData

## Properties

| Name              | Type                                                                                 | Description              | Notes                                         |
| ----------------- | ------------------------------------------------------------------------------------ | ------------------------ | --------------------------------------------- |
| **Attributes**    | Pointer to [**AuthNMappingUpdateAttributes**](AuthNMappingUpdateAttributes.md)       |                          | [optional]                                    |
| **Id**            | **string**                                                                           | ID of the AuthN Mapping. |
| **Relationships** | Pointer to [**AuthNMappingUpdateRelationships**](AuthNMappingUpdateRelationships.md) |                          | [optional]                                    |
| **Type**          | [**AuthNMappingsType**](AuthNMappingsType.md)                                        |                          | [default to AUTHNMAPPINGSTYPE_AUTHN_MAPPINGS] |

## Methods

### NewAuthNMappingUpdateData

`func NewAuthNMappingUpdateData(id string, type_ AuthNMappingsType) *AuthNMappingUpdateData`

NewAuthNMappingUpdateData instantiates a new AuthNMappingUpdateData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAuthNMappingUpdateDataWithDefaults

`func NewAuthNMappingUpdateDataWithDefaults() *AuthNMappingUpdateData`

NewAuthNMappingUpdateDataWithDefaults instantiates a new AuthNMappingUpdateData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *AuthNMappingUpdateData) GetAttributes() AuthNMappingUpdateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AuthNMappingUpdateData) GetAttributesOk() (*AuthNMappingUpdateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AuthNMappingUpdateData) SetAttributes(v AuthNMappingUpdateAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AuthNMappingUpdateData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *AuthNMappingUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthNMappingUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthNMappingUpdateData) SetId(v string)`

SetId sets Id field to given value.

### GetRelationships

`func (o *AuthNMappingUpdateData) GetRelationships() AuthNMappingUpdateRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AuthNMappingUpdateData) GetRelationshipsOk() (*AuthNMappingUpdateRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AuthNMappingUpdateData) SetRelationships(v AuthNMappingUpdateRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AuthNMappingUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *AuthNMappingUpdateData) GetType() AuthNMappingsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthNMappingUpdateData) GetTypeOk() (*AuthNMappingsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthNMappingUpdateData) SetType(v AuthNMappingsType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
