# FullApplicationKeyAttributes

## Properties

| Name          | Type                    | Description                                                                                                                                               | Notes                 |
| ------------- | ----------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------- |
| **CreatedAt** | Pointer to **string**   | Creation date of the application key.                                                                                                                     | [optional] [readonly] |
| **Key**       | Pointer to **string**   | The application key.                                                                                                                                      | [optional] [readonly] |
| **Last4**     | Pointer to **string**   | The last four characters of the application key.                                                                                                          | [optional] [readonly] |
| **Name**      | Pointer to **string**   | Name of the application key.                                                                                                                              | [optional]            |
| **Scopes**    | Pointer to **[]string** | Array of scopes to grant the application key. This feature is in private beta, please contact Datadog support to enable scopes for your application keys. | [optional]            |

## Methods

### NewFullApplicationKeyAttributes

`func NewFullApplicationKeyAttributes() *FullApplicationKeyAttributes`

NewFullApplicationKeyAttributes instantiates a new FullApplicationKeyAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewFullApplicationKeyAttributesWithDefaults

`func NewFullApplicationKeyAttributesWithDefaults() *FullApplicationKeyAttributes`

NewFullApplicationKeyAttributesWithDefaults instantiates a new FullApplicationKeyAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCreatedAt

`func (o *FullApplicationKeyAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FullApplicationKeyAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FullApplicationKeyAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FullApplicationKeyAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetKey

`func (o *FullApplicationKeyAttributes) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FullApplicationKeyAttributes) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FullApplicationKeyAttributes) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FullApplicationKeyAttributes) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLast4

`func (o *FullApplicationKeyAttributes) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *FullApplicationKeyAttributes) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *FullApplicationKeyAttributes) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *FullApplicationKeyAttributes) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetName

`func (o *FullApplicationKeyAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullApplicationKeyAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullApplicationKeyAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FullApplicationKeyAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScopes

`func (o *FullApplicationKeyAttributes) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *FullApplicationKeyAttributes) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *FullApplicationKeyAttributes) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *FullApplicationKeyAttributes) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *FullApplicationKeyAttributes) SetScopesNil(b bool)`

SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes

`func (o *FullApplicationKeyAttributes) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
