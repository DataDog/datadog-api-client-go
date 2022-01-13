# ApplicationKeyUpdateAttributes

## Properties

| Name       | Type                    | Description                                                                                                                                               | Notes      |
| ---------- | ----------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **Name**   | Pointer to **string**   | Name of the application key.                                                                                                                              | [optional] |
| **Scopes** | Pointer to **[]string** | Array of scopes to grant the application key. This feature is in private beta, please contact Datadog support to enable scopes for your application keys. | [optional] |

## Methods

### NewApplicationKeyUpdateAttributes

`func NewApplicationKeyUpdateAttributes() *ApplicationKeyUpdateAttributes`

NewApplicationKeyUpdateAttributes instantiates a new ApplicationKeyUpdateAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewApplicationKeyUpdateAttributesWithDefaults

`func NewApplicationKeyUpdateAttributesWithDefaults() *ApplicationKeyUpdateAttributes`

NewApplicationKeyUpdateAttributesWithDefaults instantiates a new ApplicationKeyUpdateAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetName

`func (o *ApplicationKeyUpdateAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationKeyUpdateAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationKeyUpdateAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationKeyUpdateAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScopes

`func (o *ApplicationKeyUpdateAttributes) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApplicationKeyUpdateAttributes) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApplicationKeyUpdateAttributes) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ApplicationKeyUpdateAttributes) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *ApplicationKeyUpdateAttributes) SetScopesNil(b bool)`

SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes

`func (o *ApplicationKeyUpdateAttributes) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
