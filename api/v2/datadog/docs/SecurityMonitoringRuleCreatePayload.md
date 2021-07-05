# SecurityMonitoringRuleCreatePayload

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Cases** | [**[]SecurityMonitoringRuleCaseCreate**](SecurityMonitoringRuleCaseCreate.md) | Cases for generating signals. | 
**Filters** | Pointer to [**[]SecurityMonitoringFilter**](SecurityMonitoringFilter.md) | Additional queries to filter matched events before they are processed. | [optional] 
**HasExtendedTitle** | Pointer to **bool** | Whether the notifications include the triggering group-by values in their title. | [optional] 
**IsEnabled** | **bool** | Whether the rule is enabled. | 
**Message** | **string** | Message for generated signals. | 
**Name** | **string** | The name of the rule. | 
**Options** | [**SecurityMonitoringRuleOptions**](SecurityMonitoringRuleOptions.md) |  | 
**Queries** | [**[]SecurityMonitoringRuleQueryCreate**](SecurityMonitoringRuleQueryCreate.md) | Queries for selecting logs which are part of the rule. | 
**Tags** | Pointer to **[]string** | Tags for generated signals. | [optional] 

## Methods

### NewSecurityMonitoringRuleCreatePayload

`func NewSecurityMonitoringRuleCreatePayload(cases []SecurityMonitoringRuleCaseCreate, isEnabled bool, message string, name string, options SecurityMonitoringRuleOptions, queries []SecurityMonitoringRuleQueryCreate) *SecurityMonitoringRuleCreatePayload`

NewSecurityMonitoringRuleCreatePayload instantiates a new SecurityMonitoringRuleCreatePayload object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSecurityMonitoringRuleCreatePayloadWithDefaults

`func NewSecurityMonitoringRuleCreatePayloadWithDefaults() *SecurityMonitoringRuleCreatePayload`

NewSecurityMonitoringRuleCreatePayloadWithDefaults instantiates a new SecurityMonitoringRuleCreatePayload object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCases

`func (o *SecurityMonitoringRuleCreatePayload) GetCases() []SecurityMonitoringRuleCaseCreate`

GetCases returns the Cases field if non-nil, zero value otherwise.

### GetCasesOk

`func (o *SecurityMonitoringRuleCreatePayload) GetCasesOk() (*[]SecurityMonitoringRuleCaseCreate, bool)`

GetCasesOk returns a tuple with the Cases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCases

`func (o *SecurityMonitoringRuleCreatePayload) SetCases(v []SecurityMonitoringRuleCaseCreate)`

SetCases sets Cases field to given value.


### GetFilters

`func (o *SecurityMonitoringRuleCreatePayload) GetFilters() []SecurityMonitoringFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SecurityMonitoringRuleCreatePayload) GetFiltersOk() (*[]SecurityMonitoringFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SecurityMonitoringRuleCreatePayload) SetFilters(v []SecurityMonitoringFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *SecurityMonitoringRuleCreatePayload) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetHasExtendedTitle

`func (o *SecurityMonitoringRuleCreatePayload) GetHasExtendedTitle() bool`

GetHasExtendedTitle returns the HasExtendedTitle field if non-nil, zero value otherwise.

### GetHasExtendedTitleOk

`func (o *SecurityMonitoringRuleCreatePayload) GetHasExtendedTitleOk() (*bool, bool)`

GetHasExtendedTitleOk returns a tuple with the HasExtendedTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExtendedTitle

`func (o *SecurityMonitoringRuleCreatePayload) SetHasExtendedTitle(v bool)`

SetHasExtendedTitle sets HasExtendedTitle field to given value.

### HasHasExtendedTitle

`func (o *SecurityMonitoringRuleCreatePayload) HasHasExtendedTitle() bool`

HasHasExtendedTitle returns a boolean if a field has been set.

### GetIsEnabled

`func (o *SecurityMonitoringRuleCreatePayload) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SecurityMonitoringRuleCreatePayload) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SecurityMonitoringRuleCreatePayload) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetMessage

`func (o *SecurityMonitoringRuleCreatePayload) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SecurityMonitoringRuleCreatePayload) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SecurityMonitoringRuleCreatePayload) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetName

`func (o *SecurityMonitoringRuleCreatePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityMonitoringRuleCreatePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityMonitoringRuleCreatePayload) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *SecurityMonitoringRuleCreatePayload) GetOptions() SecurityMonitoringRuleOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SecurityMonitoringRuleCreatePayload) GetOptionsOk() (*SecurityMonitoringRuleOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SecurityMonitoringRuleCreatePayload) SetOptions(v SecurityMonitoringRuleOptions)`

SetOptions sets Options field to given value.


### GetQueries

`func (o *SecurityMonitoringRuleCreatePayload) GetQueries() []SecurityMonitoringRuleQueryCreate`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *SecurityMonitoringRuleCreatePayload) GetQueriesOk() (*[]SecurityMonitoringRuleQueryCreate, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *SecurityMonitoringRuleCreatePayload) SetQueries(v []SecurityMonitoringRuleQueryCreate)`

SetQueries sets Queries field to given value.


### GetTags

`func (o *SecurityMonitoringRuleCreatePayload) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SecurityMonitoringRuleCreatePayload) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SecurityMonitoringRuleCreatePayload) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SecurityMonitoringRuleCreatePayload) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


