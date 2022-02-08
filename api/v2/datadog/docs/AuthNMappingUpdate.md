# AuthNMappingUpdate

## Properties

| Name              | Type                                                                                 | Description              | Notes                                         |
| ----------------- | ------------------------------------------------------------------------------------ | ------------------------ | --------------------------------------------- |
| **Attributes**    | Pointer to [**AuthNMappingUpdateAttributes**](AuthNMappingUpdateAttributes.md)       |                          | [optional]                                    |
| **Id**            | **string**                                                                           | ID of the AuthN Mapping. |
| **Relationships** | Pointer to [**AuthNMappingUpdateRelationships**](AuthNMappingUpdateRelationships.md) |                          | [optional]                                    |
| **Type**          | [**AuthNMappingsType**](AuthNMappingsType.md)                                        |                          | [default to AUTHNMAPPINGSTYPE_AUTHN_MAPPINGS] |

## Methods

### NewAuthNMappingUpdate

`func NewAuthNMappingUpdate(id string, type_ AuthNMappingsType) *AuthNMappingUpdate`

NewAuthNMappingUpdate instantiates a new AuthNMappingUpdate object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAuthNMappingUpdateWithDefaults

`func NewAuthNMappingUpdateWithDefaults() *AuthNMappingUpdate`

NewAuthNMappingUpdateWithDefaults instantiates a new AuthNMappingUpdate object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *AuthNMappingUpdate) GetAttributes() AuthNMappingUpdateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AuthNMappingUpdate) GetAttributesOk() (*AuthNMappingUpdateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AuthNMappingUpdate) SetAttributes(v AuthNMappingUpdateAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AuthNMappingUpdate) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *AuthNMappingUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthNMappingUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthNMappingUpdate) SetId(v string)`

SetId sets Id field to given value.

### GetRelationships

`func (o *AuthNMappingUpdate) GetRelationships() AuthNMappingUpdateRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AuthNMappingUpdate) GetRelationshipsOk() (*AuthNMappingUpdateRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AuthNMappingUpdate) SetRelationships(v AuthNMappingUpdateRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AuthNMappingUpdate) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *AuthNMappingUpdate) GetType() AuthNMappingsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthNMappingUpdate) GetTypeOk() (*AuthNMappingsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthNMappingUpdate) SetType(v AuthNMappingsType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
