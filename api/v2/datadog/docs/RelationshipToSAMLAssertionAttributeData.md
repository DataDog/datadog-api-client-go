# RelationshipToSAMLAssertionAttributeData

## Properties

| Name     | Type                                                              | Description                             | Notes                                                              |
| -------- | ----------------------------------------------------------------- | --------------------------------------- | ------------------------------------------------------------------ |
| **Id**   | **int32**                                                         | The ID of the SAML assertion attribute. |
| **Type** | [**SAMLAssertionAttributesType**](SAMLAssertionAttributesType.md) |                                         | [default to SAMLASSERTIONATTRIBUTESTYPE_SAML_ASSERTION_ATTRIBUTES] |

## Methods

### NewRelationshipToSAMLAssertionAttributeData

`func NewRelationshipToSAMLAssertionAttributeData(id int32, type_ SAMLAssertionAttributesType) *RelationshipToSAMLAssertionAttributeData`

NewRelationshipToSAMLAssertionAttributeData instantiates a new RelationshipToSAMLAssertionAttributeData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRelationshipToSAMLAssertionAttributeDataWithDefaults

`func NewRelationshipToSAMLAssertionAttributeDataWithDefaults() *RelationshipToSAMLAssertionAttributeData`

NewRelationshipToSAMLAssertionAttributeDataWithDefaults instantiates a new RelationshipToSAMLAssertionAttributeData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetId

`func (o *RelationshipToSAMLAssertionAttributeData) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationshipToSAMLAssertionAttributeData) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationshipToSAMLAssertionAttributeData) SetId(v int32)`

SetId sets Id field to given value.

### GetType

`func (o *RelationshipToSAMLAssertionAttributeData) GetType() SAMLAssertionAttributesType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipToSAMLAssertionAttributeData) GetTypeOk() (*SAMLAssertionAttributesType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipToSAMLAssertionAttributeData) SetType(v SAMLAssertionAttributesType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
