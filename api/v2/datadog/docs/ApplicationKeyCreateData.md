# ApplicationKeyCreateData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | [**ApplicationKeyCreateAttributes**](ApplicationKeyCreateAttributes.md) |  | 
**Type** | [**ApplicationKeysType**](ApplicationKeysType.md) |  | [default to APPLICATIONKEYSTYPE_APPLICATION_KEYS]

## Methods

### NewApplicationKeyCreateData

`func NewApplicationKeyCreateData(attributes ApplicationKeyCreateAttributes, type_ ApplicationKeysType) *ApplicationKeyCreateData`

NewApplicationKeyCreateData instantiates a new ApplicationKeyCreateData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewApplicationKeyCreateDataWithDefaults

`func NewApplicationKeyCreateDataWithDefaults() *ApplicationKeyCreateData`

NewApplicationKeyCreateDataWithDefaults instantiates a new ApplicationKeyCreateData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *ApplicationKeyCreateData) GetAttributes() ApplicationKeyCreateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationKeyCreateData) GetAttributesOk() (*ApplicationKeyCreateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationKeyCreateData) SetAttributes(v ApplicationKeyCreateAttributes)`

SetAttributes sets Attributes field to given value.


### GetType

`func (o *ApplicationKeyCreateData) GetType() ApplicationKeysType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationKeyCreateData) GetTypeOk() (*ApplicationKeysType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationKeyCreateData) SetType(v ApplicationKeysType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


