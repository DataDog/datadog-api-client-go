# SecurityFilterMeta

## Properties

| Name        | Type                  | Description        | Notes      |
| ----------- | --------------------- | ------------------ | ---------- |
| **Warning** | Pointer to **string** | A warning message. | [optional] |

## Methods

### NewSecurityFilterMeta

`func NewSecurityFilterMeta() *SecurityFilterMeta`

NewSecurityFilterMeta instantiates a new SecurityFilterMeta object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSecurityFilterMetaWithDefaults

`func NewSecurityFilterMetaWithDefaults() *SecurityFilterMeta`

NewSecurityFilterMetaWithDefaults instantiates a new SecurityFilterMeta object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetWarning

`func (o *SecurityFilterMeta) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *SecurityFilterMeta) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *SecurityFilterMeta) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *SecurityFilterMeta) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
