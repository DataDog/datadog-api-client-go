# PartialApplicationKeyAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**CreatedAt** | Pointer to **string** | Creation date of the application key. | [optional] [readonly] 
**Last4** | Pointer to **string** | The last four characters of the application key. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the application key. | [optional] 
**Scopes** | Pointer to **[]string** | Array of scopes to grant the application key. This feature is in private beta, please contact Datadog support to enable scopes for your application keys. | [optional] 

## Methods

### NewPartialApplicationKeyAttributes

`func NewPartialApplicationKeyAttributes() *PartialApplicationKeyAttributes`

NewPartialApplicationKeyAttributes instantiates a new PartialApplicationKeyAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewPartialApplicationKeyAttributesWithDefaults

`func NewPartialApplicationKeyAttributesWithDefaults() *PartialApplicationKeyAttributes`

NewPartialApplicationKeyAttributesWithDefaults instantiates a new PartialApplicationKeyAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCreatedAt

`func (o *PartialApplicationKeyAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PartialApplicationKeyAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PartialApplicationKeyAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PartialApplicationKeyAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLast4

`func (o *PartialApplicationKeyAttributes) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *PartialApplicationKeyAttributes) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *PartialApplicationKeyAttributes) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *PartialApplicationKeyAttributes) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetName

`func (o *PartialApplicationKeyAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartialApplicationKeyAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartialApplicationKeyAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartialApplicationKeyAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScopes

`func (o *PartialApplicationKeyAttributes) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PartialApplicationKeyAttributes) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PartialApplicationKeyAttributes) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PartialApplicationKeyAttributes) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *PartialApplicationKeyAttributes) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *PartialApplicationKeyAttributes) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


