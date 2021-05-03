# SecurityMonitoringRuleResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Cases** | Pointer to [**[]SecurityMonitoringRuleCase**](SecurityMonitoringRuleCase.md) | Cases for generating signals. | [optional] 
**CreatedAt** | Pointer to **int64** | When the rule was created, timestamp in milliseconds. | [optional] 
**CreationAuthorId** | Pointer to **int64** | User ID of the user who created the rule. | [optional] 
**Filters** | Pointer to [**[]SecurityMonitoringFilter**](SecurityMonitoringFilter.md) | Additional queries to filter matched events before they are processed. | [optional] 
**Id** | Pointer to **string** | The ID of the rule. | [optional] 
**IsDefault** | Pointer to **bool** | Whether the rule is included by default. | [optional] 
**IsDeleted** | Pointer to **bool** | Whether the rule has been deleted. | [optional] 
**IsEnabled** | Pointer to **bool** | Whether the rule is enabled. | [optional] 
**Message** | Pointer to **string** | Message for generated signals. | [optional] 
**Name** | Pointer to **string** | The name of the rule. | [optional] 
**Options** | Pointer to [**SecurityMonitoringRuleOptions**](SecurityMonitoringRuleOptions.md) |  | [optional] 
**Queries** | Pointer to [**[]SecurityMonitoringRuleQuery**](SecurityMonitoringRuleQuery.md) | Queries for selecting logs which are part of the rule. | [optional] 
**Tags** | Pointer to **[]string** | Tags for generated signals. | [optional] 
**Version** | Pointer to **int64** | The version of the rule. | [optional] 

## Methods

### NewSecurityMonitoringRuleResponse

`func NewSecurityMonitoringRuleResponse() *SecurityMonitoringRuleResponse`

NewSecurityMonitoringRuleResponse instantiates a new SecurityMonitoringRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityMonitoringRuleResponseWithDefaults

`func NewSecurityMonitoringRuleResponseWithDefaults() *SecurityMonitoringRuleResponse`

NewSecurityMonitoringRuleResponseWithDefaults instantiates a new SecurityMonitoringRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCases

`func (o *SecurityMonitoringRuleResponse) GetCases() []SecurityMonitoringRuleCase`

GetCases returns the Cases field if non-nil, zero value otherwise.

### GetCasesOk

`func (o *SecurityMonitoringRuleResponse) GetCasesOk() (*[]SecurityMonitoringRuleCase, bool)`

GetCasesOk returns a tuple with the Cases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCases

`func (o *SecurityMonitoringRuleResponse) SetCases(v []SecurityMonitoringRuleCase)`

SetCases sets Cases field to given value.

### HasCases

`func (o *SecurityMonitoringRuleResponse) HasCases() bool`

HasCases returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SecurityMonitoringRuleResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecurityMonitoringRuleResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecurityMonitoringRuleResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SecurityMonitoringRuleResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreationAuthorId

`func (o *SecurityMonitoringRuleResponse) GetCreationAuthorId() int64`

GetCreationAuthorId returns the CreationAuthorId field if non-nil, zero value otherwise.

### GetCreationAuthorIdOk

`func (o *SecurityMonitoringRuleResponse) GetCreationAuthorIdOk() (*int64, bool)`

GetCreationAuthorIdOk returns a tuple with the CreationAuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationAuthorId

`func (o *SecurityMonitoringRuleResponse) SetCreationAuthorId(v int64)`

SetCreationAuthorId sets CreationAuthorId field to given value.

### HasCreationAuthorId

`func (o *SecurityMonitoringRuleResponse) HasCreationAuthorId() bool`

HasCreationAuthorId returns a boolean if a field has been set.

### GetFilters

`func (o *SecurityMonitoringRuleResponse) GetFilters() []SecurityMonitoringFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SecurityMonitoringRuleResponse) GetFiltersOk() (*[]SecurityMonitoringFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SecurityMonitoringRuleResponse) SetFilters(v []SecurityMonitoringFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *SecurityMonitoringRuleResponse) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetId

`func (o *SecurityMonitoringRuleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityMonitoringRuleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityMonitoringRuleResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityMonitoringRuleResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *SecurityMonitoringRuleResponse) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *SecurityMonitoringRuleResponse) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *SecurityMonitoringRuleResponse) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *SecurityMonitoringRuleResponse) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsDeleted

`func (o *SecurityMonitoringRuleResponse) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *SecurityMonitoringRuleResponse) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *SecurityMonitoringRuleResponse) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *SecurityMonitoringRuleResponse) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsEnabled

`func (o *SecurityMonitoringRuleResponse) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SecurityMonitoringRuleResponse) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SecurityMonitoringRuleResponse) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *SecurityMonitoringRuleResponse) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetMessage

`func (o *SecurityMonitoringRuleResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SecurityMonitoringRuleResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SecurityMonitoringRuleResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SecurityMonitoringRuleResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *SecurityMonitoringRuleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityMonitoringRuleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityMonitoringRuleResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityMonitoringRuleResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *SecurityMonitoringRuleResponse) GetOptions() SecurityMonitoringRuleOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SecurityMonitoringRuleResponse) GetOptionsOk() (*SecurityMonitoringRuleOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SecurityMonitoringRuleResponse) SetOptions(v SecurityMonitoringRuleOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SecurityMonitoringRuleResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetQueries

`func (o *SecurityMonitoringRuleResponse) GetQueries() []SecurityMonitoringRuleQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *SecurityMonitoringRuleResponse) GetQueriesOk() (*[]SecurityMonitoringRuleQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *SecurityMonitoringRuleResponse) SetQueries(v []SecurityMonitoringRuleQuery)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *SecurityMonitoringRuleResponse) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetTags

`func (o *SecurityMonitoringRuleResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SecurityMonitoringRuleResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SecurityMonitoringRuleResponse) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SecurityMonitoringRuleResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *SecurityMonitoringRuleResponse) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SecurityMonitoringRuleResponse) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SecurityMonitoringRuleResponse) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SecurityMonitoringRuleResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


