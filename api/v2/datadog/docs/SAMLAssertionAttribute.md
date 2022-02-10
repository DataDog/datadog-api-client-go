# SAMLAssertionAttribute

## Properties

| Name           | Type                                                                                   | Description                             | Notes                                                              |
| -------------- | -------------------------------------------------------------------------------------- | --------------------------------------- | ------------------------------------------------------------------ |
| **Attributes** | Pointer to [**SAMLAssertionAttributeAttributes**](SAMLAssertionAttributeAttributes.md) |                                         | [optional]                                                         |
| **Id**         | **int32**                                                                              | The ID of the SAML assertion attribute. |
| **Type**       | [**SAMLAssertionAttributesType**](SAMLAssertionAttributesType.md)                      |                                         | [default to SAMLASSERTIONATTRIBUTESTYPE_SAML_ASSERTION_ATTRIBUTES] |

## Methods

### NewSAMLAssertionAttribute

`func NewSAMLAssertionAttribute(id int32, type_ SAMLAssertionAttributesType) *SAMLAssertionAttribute`

NewSAMLAssertionAttribute instantiates a new SAMLAssertionAttribute object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSAMLAssertionAttributeWithDefaults

`func NewSAMLAssertionAttributeWithDefaults() *SAMLAssertionAttribute`

NewSAMLAssertionAttributeWithDefaults instantiates a new SAMLAssertionAttribute object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *SAMLAssertionAttribute) GetAttributes() SAMLAssertionAttributeAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SAMLAssertionAttribute) GetAttributesOk() (*SAMLAssertionAttributeAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SAMLAssertionAttribute) SetAttributes(v SAMLAssertionAttributeAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SAMLAssertionAttribute) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *SAMLAssertionAttribute) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SAMLAssertionAttribute) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SAMLAssertionAttribute) SetId(v int32)`

SetId sets Id field to given value.

### GetType

`func (o *SAMLAssertionAttribute) GetType() SAMLAssertionAttributesType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SAMLAssertionAttribute) GetTypeOk() (*SAMLAssertionAttributesType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SAMLAssertionAttribute) SetType(v SAMLAssertionAttributesType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
