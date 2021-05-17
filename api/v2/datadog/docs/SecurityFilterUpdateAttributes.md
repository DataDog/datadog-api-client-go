# SecurityFilterUpdateAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**ExclusionFilters** | Pointer to [**[]SecurityFilterExclusionFilter**](SecurityFilterExclusionFilter.md) | Exclusion filters to exclude some logs from the security filter. | [optional] 
**FilteredDataType** | Pointer to [**SecurityFilterFilteredDataType**](SecurityFilterFilteredDataType.md) |  | [optional] 
**IsEnabled** | Pointer to **bool** | Whether the security filter is enabled. | [optional] 
**Name** | Pointer to **string** | The name of the security filter. | [optional] 
**Query** | Pointer to **string** | The query of the security filter. | [optional] 
**Version** | Pointer to **int32** | The version of the security filter to update. | [optional] 

## Methods

### NewSecurityFilterUpdateAttributes

`func NewSecurityFilterUpdateAttributes() *SecurityFilterUpdateAttributes`

NewSecurityFilterUpdateAttributes instantiates a new SecurityFilterUpdateAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityFilterUpdateAttributesWithDefaults

`func NewSecurityFilterUpdateAttributesWithDefaults() *SecurityFilterUpdateAttributes`

NewSecurityFilterUpdateAttributesWithDefaults instantiates a new SecurityFilterUpdateAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclusionFilters

`func (o *SecurityFilterUpdateAttributes) GetExclusionFilters() []SecurityFilterExclusionFilter`

GetExclusionFilters returns the ExclusionFilters field if non-nil, zero value otherwise.

### GetExclusionFiltersOk

`func (o *SecurityFilterUpdateAttributes) GetExclusionFiltersOk() (*[]SecurityFilterExclusionFilter, bool)`

GetExclusionFiltersOk returns a tuple with the ExclusionFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionFilters

`func (o *SecurityFilterUpdateAttributes) SetExclusionFilters(v []SecurityFilterExclusionFilter)`

SetExclusionFilters sets ExclusionFilters field to given value.

### HasExclusionFilters

`func (o *SecurityFilterUpdateAttributes) HasExclusionFilters() bool`

HasExclusionFilters returns a boolean if a field has been set.

### GetFilteredDataType

`func (o *SecurityFilterUpdateAttributes) GetFilteredDataType() SecurityFilterFilteredDataType`

GetFilteredDataType returns the FilteredDataType field if non-nil, zero value otherwise.

### GetFilteredDataTypeOk

`func (o *SecurityFilterUpdateAttributes) GetFilteredDataTypeOk() (*SecurityFilterFilteredDataType, bool)`

GetFilteredDataTypeOk returns a tuple with the FilteredDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredDataType

`func (o *SecurityFilterUpdateAttributes) SetFilteredDataType(v SecurityFilterFilteredDataType)`

SetFilteredDataType sets FilteredDataType field to given value.

### HasFilteredDataType

`func (o *SecurityFilterUpdateAttributes) HasFilteredDataType() bool`

HasFilteredDataType returns a boolean if a field has been set.

### GetIsEnabled

`func (o *SecurityFilterUpdateAttributes) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SecurityFilterUpdateAttributes) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SecurityFilterUpdateAttributes) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *SecurityFilterUpdateAttributes) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *SecurityFilterUpdateAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityFilterUpdateAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityFilterUpdateAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityFilterUpdateAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuery

`func (o *SecurityFilterUpdateAttributes) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SecurityFilterUpdateAttributes) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SecurityFilterUpdateAttributes) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SecurityFilterUpdateAttributes) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetVersion

`func (o *SecurityFilterUpdateAttributes) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SecurityFilterUpdateAttributes) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SecurityFilterUpdateAttributes) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SecurityFilterUpdateAttributes) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


