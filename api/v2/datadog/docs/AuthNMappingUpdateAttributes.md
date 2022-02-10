# AuthNMappingUpdateAttributes

## Properties

| Name               | Type                  | Description                                                                         | Notes      |
| ------------------ | --------------------- | ----------------------------------------------------------------------------------- | ---------- |
| **AttributeKey**   | Pointer to **string** | Key portion of a key/value pair of the attribute sent from the Identity Provider.   | [optional] |
| **AttributeValue** | Pointer to **string** | Value portion of a key/value pair of the attribute sent from the Identity Provider. | [optional] |

## Methods

### NewAuthNMappingUpdateAttributes

`func NewAuthNMappingUpdateAttributes() *AuthNMappingUpdateAttributes`

NewAuthNMappingUpdateAttributes instantiates a new AuthNMappingUpdateAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAuthNMappingUpdateAttributesWithDefaults

`func NewAuthNMappingUpdateAttributesWithDefaults() *AuthNMappingUpdateAttributes`

NewAuthNMappingUpdateAttributesWithDefaults instantiates a new AuthNMappingUpdateAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributeKey

`func (o *AuthNMappingUpdateAttributes) GetAttributeKey() string`

GetAttributeKey returns the AttributeKey field if non-nil, zero value otherwise.

### GetAttributeKeyOk

`func (o *AuthNMappingUpdateAttributes) GetAttributeKeyOk() (*string, bool)`

GetAttributeKeyOk returns a tuple with the AttributeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeKey

`func (o *AuthNMappingUpdateAttributes) SetAttributeKey(v string)`

SetAttributeKey sets AttributeKey field to given value.

### HasAttributeKey

`func (o *AuthNMappingUpdateAttributes) HasAttributeKey() bool`

HasAttributeKey returns a boolean if a field has been set.

### GetAttributeValue

`func (o *AuthNMappingUpdateAttributes) GetAttributeValue() string`

GetAttributeValue returns the AttributeValue field if non-nil, zero value otherwise.

### GetAttributeValueOk

`func (o *AuthNMappingUpdateAttributes) GetAttributeValueOk() (*string, bool)`

GetAttributeValueOk returns a tuple with the AttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValue

`func (o *AuthNMappingUpdateAttributes) SetAttributeValue(v string)`

SetAttributeValue sets AttributeValue field to given value.

### HasAttributeValue

`func (o *AuthNMappingUpdateAttributes) HasAttributeValue() bool`

HasAttributeValue returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
