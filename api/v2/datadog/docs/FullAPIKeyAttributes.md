# FullAPIKeyAttributes

## Properties

| Name           | Type                  | Description                              | Notes                 |
| -------------- | --------------------- | ---------------------------------------- | --------------------- |
| **CreatedAt**  | Pointer to **string** | Creation date of the API key.            | [optional] [readonly] |
| **Key**        | Pointer to **string** | The API key.                             | [optional] [readonly] |
| **Last4**      | Pointer to **string** | The last four characters of the API key. | [optional] [readonly] |
| **ModifiedAt** | Pointer to **string** | Date the API key was last modified.      | [optional] [readonly] |
| **Name**       | Pointer to **string** | Name of the API key.                     | [optional]            |

## Methods

### NewFullAPIKeyAttributes

`func NewFullAPIKeyAttributes() *FullAPIKeyAttributes`

NewFullAPIKeyAttributes instantiates a new FullAPIKeyAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewFullAPIKeyAttributesWithDefaults

`func NewFullAPIKeyAttributesWithDefaults() *FullAPIKeyAttributes`

NewFullAPIKeyAttributesWithDefaults instantiates a new FullAPIKeyAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCreatedAt

`func (o *FullAPIKeyAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FullAPIKeyAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FullAPIKeyAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FullAPIKeyAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetKey

`func (o *FullAPIKeyAttributes) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FullAPIKeyAttributes) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FullAPIKeyAttributes) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FullAPIKeyAttributes) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLast4

`func (o *FullAPIKeyAttributes) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *FullAPIKeyAttributes) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *FullAPIKeyAttributes) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *FullAPIKeyAttributes) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetModifiedAt

`func (o *FullAPIKeyAttributes) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *FullAPIKeyAttributes) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *FullAPIKeyAttributes) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *FullAPIKeyAttributes) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *FullAPIKeyAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullAPIKeyAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullAPIKeyAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FullAPIKeyAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
