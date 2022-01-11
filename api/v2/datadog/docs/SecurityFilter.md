# SecurityFilter

## Properties

| Name           | Type                                                                   | Description                    | Notes                                                       |
| -------------- | ---------------------------------------------------------------------- | ------------------------------ | ----------------------------------------------------------- |
| **Attributes** | Pointer to [**SecurityFilterAttributes**](SecurityFilterAttributes.md) |                                | [optional]                                                  |
| **Id**         | Pointer to **string**                                                  | The ID of the security filter. | [optional]                                                  |
| **Type**       | Pointer to [**SecurityFilterType**](SecurityFilterType.md)             |                                | [optional] [default to SECURITYFILTERTYPE_SECURITY_FILTERS] |

## Methods

### NewSecurityFilter

`func NewSecurityFilter() *SecurityFilter`

NewSecurityFilter instantiates a new SecurityFilter object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSecurityFilterWithDefaults

`func NewSecurityFilterWithDefaults() *SecurityFilter`

NewSecurityFilterWithDefaults instantiates a new SecurityFilter object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *SecurityFilter) GetAttributes() SecurityFilterAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SecurityFilter) GetAttributesOk() (*SecurityFilterAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SecurityFilter) SetAttributes(v SecurityFilterAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SecurityFilter) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *SecurityFilter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityFilter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityFilter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SecurityFilter) GetType() SecurityFilterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityFilter) GetTypeOk() (*SecurityFilterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityFilter) SetType(v SecurityFilterType)`

SetType sets Type field to given value.

### HasType

`func (o *SecurityFilter) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
