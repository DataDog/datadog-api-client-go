# SecurityFilterCreateAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**ExclusionFilters** | [**[]SecurityFilterExclusionFilter**](SecurityFilterExclusionFilter.md) | Exclusion filters to exclude some logs from the security filter. | 
**FilteredDataType** | [**SecurityFilterFilteredDataType**](SecurityFilterFilteredDataType.md) |  | 
**IsEnabled** | **bool** | Whether the security filter is enabled. | 
**Name** | **string** | The name of the security filter. | 
**Query** | **string** | The query of the security filter. | 

## Methods

### NewSecurityFilterCreateAttributes

`func NewSecurityFilterCreateAttributes(exclusionFilters []SecurityFilterExclusionFilter, filteredDataType SecurityFilterFilteredDataType, isEnabled bool, name string, query string, ) *SecurityFilterCreateAttributes`

NewSecurityFilterCreateAttributes instantiates a new SecurityFilterCreateAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityFilterCreateAttributesWithDefaults

`func NewSecurityFilterCreateAttributesWithDefaults() *SecurityFilterCreateAttributes`

NewSecurityFilterCreateAttributesWithDefaults instantiates a new SecurityFilterCreateAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclusionFilters

`func (o *SecurityFilterCreateAttributes) GetExclusionFilters() []SecurityFilterExclusionFilter`

GetExclusionFilters returns the ExclusionFilters field if non-nil, zero value otherwise.

### GetExclusionFiltersOk

`func (o *SecurityFilterCreateAttributes) GetExclusionFiltersOk() (*[]SecurityFilterExclusionFilter, bool)`

GetExclusionFiltersOk returns a tuple with the ExclusionFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionFilters

`func (o *SecurityFilterCreateAttributes) SetExclusionFilters(v []SecurityFilterExclusionFilter)`

SetExclusionFilters sets ExclusionFilters field to given value.


### GetFilteredDataType

`func (o *SecurityFilterCreateAttributes) GetFilteredDataType() SecurityFilterFilteredDataType`

GetFilteredDataType returns the FilteredDataType field if non-nil, zero value otherwise.

### GetFilteredDataTypeOk

`func (o *SecurityFilterCreateAttributes) GetFilteredDataTypeOk() (*SecurityFilterFilteredDataType, bool)`

GetFilteredDataTypeOk returns a tuple with the FilteredDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredDataType

`func (o *SecurityFilterCreateAttributes) SetFilteredDataType(v SecurityFilterFilteredDataType)`

SetFilteredDataType sets FilteredDataType field to given value.


### GetIsEnabled

`func (o *SecurityFilterCreateAttributes) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SecurityFilterCreateAttributes) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SecurityFilterCreateAttributes) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetName

`func (o *SecurityFilterCreateAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityFilterCreateAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityFilterCreateAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetQuery

`func (o *SecurityFilterCreateAttributes) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SecurityFilterCreateAttributes) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SecurityFilterCreateAttributes) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


