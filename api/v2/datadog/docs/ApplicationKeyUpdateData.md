# ApplicationKeyUpdateData

## Properties

| Name           | Type                                                                    | Description                | Notes                                             |
| -------------- | ----------------------------------------------------------------------- | -------------------------- | ------------------------------------------------- |
| **Attributes** | [**ApplicationKeyUpdateAttributes**](ApplicationKeyUpdateAttributes.md) |                            |
| **Id**         | **string**                                                              | ID of the application key. |
| **Type**       | [**ApplicationKeysType**](ApplicationKeysType.md)                       |                            | [default to APPLICATIONKEYSTYPE_APPLICATION_KEYS] |

## Methods

### NewApplicationKeyUpdateData

`func NewApplicationKeyUpdateData(attributes ApplicationKeyUpdateAttributes, id string, type_ ApplicationKeysType) *ApplicationKeyUpdateData`

NewApplicationKeyUpdateData instantiates a new ApplicationKeyUpdateData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewApplicationKeyUpdateDataWithDefaults

`func NewApplicationKeyUpdateDataWithDefaults() *ApplicationKeyUpdateData`

NewApplicationKeyUpdateDataWithDefaults instantiates a new ApplicationKeyUpdateData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *ApplicationKeyUpdateData) GetAttributes() ApplicationKeyUpdateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationKeyUpdateData) GetAttributesOk() (*ApplicationKeyUpdateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationKeyUpdateData) SetAttributes(v ApplicationKeyUpdateAttributes)`

SetAttributes sets Attributes field to given value.

### GetId

`func (o *ApplicationKeyUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationKeyUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationKeyUpdateData) SetId(v string)`

SetId sets Id field to given value.

### GetType

`func (o *ApplicationKeyUpdateData) GetType() ApplicationKeysType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationKeyUpdateData) GetTypeOk() (*ApplicationKeysType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationKeyUpdateData) SetType(v ApplicationKeysType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
