# AuthNMappingAttributes

## Properties

| Name                         | Type                     | Description                                                                         | Notes                 |
| ---------------------------- | ------------------------ | ----------------------------------------------------------------------------------- | --------------------- |
| **AttributeKey**             | Pointer to **string**    | Key portion of a key/value pair of the attribute sent from the Identity Provider.   | [optional]            |
| **AttributeValue**           | Pointer to **string**    | Value portion of a key/value pair of the attribute sent from the Identity Provider. | [optional]            |
| **CreatedAt**                | Pointer to **time.Time** | Creation time of the AuthN Mapping.                                                 | [optional] [readonly] |
| **ModifiedAt**               | Pointer to **time.Time** | Time of last AuthN Mapping modification.                                            | [optional] [readonly] |
| **SamlAssertionAttributeId** | Pointer to **int32**     | The ID of the SAML assertion attribute.                                             | [optional]            |

## Methods

### NewAuthNMappingAttributes

`func NewAuthNMappingAttributes() *AuthNMappingAttributes`

NewAuthNMappingAttributes instantiates a new AuthNMappingAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAuthNMappingAttributesWithDefaults

`func NewAuthNMappingAttributesWithDefaults() *AuthNMappingAttributes`

NewAuthNMappingAttributesWithDefaults instantiates a new AuthNMappingAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributeKey

`func (o *AuthNMappingAttributes) GetAttributeKey() string`

GetAttributeKey returns the AttributeKey field if non-nil, zero value otherwise.

### GetAttributeKeyOk

`func (o *AuthNMappingAttributes) GetAttributeKeyOk() (*string, bool)`

GetAttributeKeyOk returns a tuple with the AttributeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeKey

`func (o *AuthNMappingAttributes) SetAttributeKey(v string)`

SetAttributeKey sets AttributeKey field to given value.

### HasAttributeKey

`func (o *AuthNMappingAttributes) HasAttributeKey() bool`

HasAttributeKey returns a boolean if a field has been set.

### GetAttributeValue

`func (o *AuthNMappingAttributes) GetAttributeValue() string`

GetAttributeValue returns the AttributeValue field if non-nil, zero value otherwise.

### GetAttributeValueOk

`func (o *AuthNMappingAttributes) GetAttributeValueOk() (*string, bool)`

GetAttributeValueOk returns a tuple with the AttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValue

`func (o *AuthNMappingAttributes) SetAttributeValue(v string)`

SetAttributeValue sets AttributeValue field to given value.

### HasAttributeValue

`func (o *AuthNMappingAttributes) HasAttributeValue() bool`

HasAttributeValue returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuthNMappingAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthNMappingAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthNMappingAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthNMappingAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *AuthNMappingAttributes) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AuthNMappingAttributes) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AuthNMappingAttributes) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *AuthNMappingAttributes) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetSamlAssertionAttributeId

`func (o *AuthNMappingAttributes) GetSamlAssertionAttributeId() int32`

GetSamlAssertionAttributeId returns the SamlAssertionAttributeId field if non-nil, zero value otherwise.

### GetSamlAssertionAttributeIdOk

`func (o *AuthNMappingAttributes) GetSamlAssertionAttributeIdOk() (*int32, bool)`

GetSamlAssertionAttributeIdOk returns a tuple with the SamlAssertionAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAssertionAttributeId

`func (o *AuthNMappingAttributes) SetSamlAssertionAttributeId(v int32)`

SetSamlAssertionAttributeId sets SamlAssertionAttributeId field to given value.

### HasSamlAssertionAttributeId

`func (o *AuthNMappingAttributes) HasSamlAssertionAttributeId() bool`

HasSamlAssertionAttributeId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
