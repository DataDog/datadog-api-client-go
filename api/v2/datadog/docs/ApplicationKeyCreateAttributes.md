# ApplicationKeyCreateAttributes

## Properties

| Name       | Type                    | Description                                                                                                                                               | Notes      |
| ---------- | ----------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **Name**   | **string**              | Name of the application key.                                                                                                                              |
| **Scopes** | Pointer to **[]string** | Array of scopes to grant the application key. This feature is in private beta, please contact Datadog support to enable scopes for your application keys. | [optional] |

## Methods

### NewApplicationKeyCreateAttributes

`func NewApplicationKeyCreateAttributes(name string) *ApplicationKeyCreateAttributes`

NewApplicationKeyCreateAttributes instantiates a new ApplicationKeyCreateAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewApplicationKeyCreateAttributesWithDefaults

`func NewApplicationKeyCreateAttributesWithDefaults() *ApplicationKeyCreateAttributes`

NewApplicationKeyCreateAttributesWithDefaults instantiates a new ApplicationKeyCreateAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetName

`func (o *ApplicationKeyCreateAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationKeyCreateAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationKeyCreateAttributes) SetName(v string)`

SetName sets Name field to given value.

### GetScopes

`func (o *ApplicationKeyCreateAttributes) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApplicationKeyCreateAttributes) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApplicationKeyCreateAttributes) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ApplicationKeyCreateAttributes) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *ApplicationKeyCreateAttributes) SetScopesNil(b bool)`

SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes

`func (o *ApplicationKeyCreateAttributes) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
