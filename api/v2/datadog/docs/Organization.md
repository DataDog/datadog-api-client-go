# Organization

## Properties

| Name           | Type                                                               | Description             | Notes                               |
| -------------- | ------------------------------------------------------------------ | ----------------------- | ----------------------------------- |
| **Attributes** | Pointer to [**OrganizationAttributes**](OrganizationAttributes.md) |                         | [optional]                          |
| **Id**         | Pointer to **string**                                              | ID of the organization. | [optional]                          |
| **Type**       | [**OrganizationsType**](OrganizationsType.md)                      |                         | [default to ORGANIZATIONSTYPE_ORGS] |

## Methods

### NewOrganization

`func NewOrganization(type_ OrganizationsType) *Organization`

NewOrganization instantiates a new Organization object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *Organization) GetAttributes() OrganizationAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Organization) GetAttributesOk() (*OrganizationAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Organization) SetAttributes(v OrganizationAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Organization) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Organization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Organization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Organization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Organization) GetType() OrganizationsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Organization) GetTypeOk() (*OrganizationsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Organization) SetType(v OrganizationsType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
