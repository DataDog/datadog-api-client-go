# SecurityFilterAttributes

## Properties

| Name                 | Type                                                                                               | Description                                                                             | Notes      |
| -------------------- | -------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | ---------- |
| **ExclusionFilters** | Pointer to [**[]SecurityFilterExclusionFilterResponse**](SecurityFilterExclusionFilterResponse.md) | The list of exclusion filters applied in this security filter.                          | [optional] |
| **FilteredDataType** | Pointer to [**SecurityFilterFilteredDataType**](SecurityFilterFilteredDataType.md)                 |                                                                                         | [optional] |
| **IsBuiltin**        | Pointer to **bool**                                                                                | Whether the security filter is the built-in filter.                                     | [optional] |
| **IsEnabled**        | Pointer to **bool**                                                                                | Whether the security filter is enabled.                                                 | [optional] |
| **Name**             | Pointer to **string**                                                                              | The security filter name.                                                               | [optional] |
| **Query**            | Pointer to **string**                                                                              | The security filter query. Logs accepted by this query will be accepted by this filter. | [optional] |
| **Version**          | Pointer to **int32**                                                                               | The version of the security filter.                                                     | [optional] |

## Methods

### NewSecurityFilterAttributes

`func NewSecurityFilterAttributes() *SecurityFilterAttributes`

NewSecurityFilterAttributes instantiates a new SecurityFilterAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSecurityFilterAttributesWithDefaults

`func NewSecurityFilterAttributesWithDefaults() *SecurityFilterAttributes`

NewSecurityFilterAttributesWithDefaults instantiates a new SecurityFilterAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetExclusionFilters

`func (o *SecurityFilterAttributes) GetExclusionFilters() []SecurityFilterExclusionFilterResponse`

GetExclusionFilters returns the ExclusionFilters field if non-nil, zero value otherwise.

### GetExclusionFiltersOk

`func (o *SecurityFilterAttributes) GetExclusionFiltersOk() (*[]SecurityFilterExclusionFilterResponse, bool)`

GetExclusionFiltersOk returns a tuple with the ExclusionFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionFilters

`func (o *SecurityFilterAttributes) SetExclusionFilters(v []SecurityFilterExclusionFilterResponse)`

SetExclusionFilters sets ExclusionFilters field to given value.

### HasExclusionFilters

`func (o *SecurityFilterAttributes) HasExclusionFilters() bool`

HasExclusionFilters returns a boolean if a field has been set.

### GetFilteredDataType

`func (o *SecurityFilterAttributes) GetFilteredDataType() SecurityFilterFilteredDataType`

GetFilteredDataType returns the FilteredDataType field if non-nil, zero value otherwise.

### GetFilteredDataTypeOk

`func (o *SecurityFilterAttributes) GetFilteredDataTypeOk() (*SecurityFilterFilteredDataType, bool)`

GetFilteredDataTypeOk returns a tuple with the FilteredDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredDataType

`func (o *SecurityFilterAttributes) SetFilteredDataType(v SecurityFilterFilteredDataType)`

SetFilteredDataType sets FilteredDataType field to given value.

### HasFilteredDataType

`func (o *SecurityFilterAttributes) HasFilteredDataType() bool`

HasFilteredDataType returns a boolean if a field has been set.

### GetIsBuiltin

`func (o *SecurityFilterAttributes) GetIsBuiltin() bool`

GetIsBuiltin returns the IsBuiltin field if non-nil, zero value otherwise.

### GetIsBuiltinOk

`func (o *SecurityFilterAttributes) GetIsBuiltinOk() (*bool, bool)`

GetIsBuiltinOk returns a tuple with the IsBuiltin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltin

`func (o *SecurityFilterAttributes) SetIsBuiltin(v bool)`

SetIsBuiltin sets IsBuiltin field to given value.

### HasIsBuiltin

`func (o *SecurityFilterAttributes) HasIsBuiltin() bool`

HasIsBuiltin returns a boolean if a field has been set.

### GetIsEnabled

`func (o *SecurityFilterAttributes) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SecurityFilterAttributes) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SecurityFilterAttributes) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *SecurityFilterAttributes) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *SecurityFilterAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityFilterAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityFilterAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityFilterAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuery

`func (o *SecurityFilterAttributes) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SecurityFilterAttributes) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SecurityFilterAttributes) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SecurityFilterAttributes) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetVersion

`func (o *SecurityFilterAttributes) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SecurityFilterAttributes) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SecurityFilterAttributes) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SecurityFilterAttributes) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
