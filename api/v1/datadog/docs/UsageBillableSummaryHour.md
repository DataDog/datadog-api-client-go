# UsageBillableSummaryHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingPlan** | Pointer to **string** | The billing plan. | [optional] 
**EndDate** | Pointer to [**time.Time**](time.Time.md) | Shows the last date of usage. | [optional] 
**NumOrgs** | Pointer to **int64** | The number of organizations. | [optional] 
**OrgName** | Pointer to **string** | The organization name. | [optional] 
**PublicId** | Pointer to **string** | The organization public ID. | [optional] 
**RatioInMonth** | Pointer to **int64** | Shows usage aggregation for a billing period. | [optional] 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Shows the first date of usage. | [optional] 
**Usage** | Pointer to [**UsageBillableSummaryKeys**](UsageBillableSummaryKeys.md) |  | [optional] 

## Methods

### NewUsageBillableSummaryHour

`func NewUsageBillableSummaryHour() *UsageBillableSummaryHour`

NewUsageBillableSummaryHour instantiates a new UsageBillableSummaryHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageBillableSummaryHourWithDefaults

`func NewUsageBillableSummaryHourWithDefaults() *UsageBillableSummaryHour`

NewUsageBillableSummaryHourWithDefaults instantiates a new UsageBillableSummaryHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingPlan

`func (o *UsageBillableSummaryHour) GetBillingPlan() string`

GetBillingPlan returns the BillingPlan field if non-nil, zero value otherwise.

### GetBillingPlanOk

`func (o *UsageBillableSummaryHour) GetBillingPlanOk() (*string, bool)`

GetBillingPlanOk returns a tuple with the BillingPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPlan

`func (o *UsageBillableSummaryHour) SetBillingPlan(v string)`

SetBillingPlan sets BillingPlan field to given value.

### HasBillingPlan

`func (o *UsageBillableSummaryHour) HasBillingPlan() bool`

HasBillingPlan returns a boolean if a field has been set.

### GetEndDate

`func (o *UsageBillableSummaryHour) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UsageBillableSummaryHour) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UsageBillableSummaryHour) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UsageBillableSummaryHour) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetNumOrgs

`func (o *UsageBillableSummaryHour) GetNumOrgs() int64`

GetNumOrgs returns the NumOrgs field if non-nil, zero value otherwise.

### GetNumOrgsOk

`func (o *UsageBillableSummaryHour) GetNumOrgsOk() (*int64, bool)`

GetNumOrgsOk returns a tuple with the NumOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOrgs

`func (o *UsageBillableSummaryHour) SetNumOrgs(v int64)`

SetNumOrgs sets NumOrgs field to given value.

### HasNumOrgs

`func (o *UsageBillableSummaryHour) HasNumOrgs() bool`

HasNumOrgs returns a boolean if a field has been set.

### GetOrgName

`func (o *UsageBillableSummaryHour) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UsageBillableSummaryHour) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UsageBillableSummaryHour) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UsageBillableSummaryHour) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPublicId

`func (o *UsageBillableSummaryHour) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UsageBillableSummaryHour) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UsageBillableSummaryHour) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UsageBillableSummaryHour) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetRatioInMonth

`func (o *UsageBillableSummaryHour) GetRatioInMonth() int64`

GetRatioInMonth returns the RatioInMonth field if non-nil, zero value otherwise.

### GetRatioInMonthOk

`func (o *UsageBillableSummaryHour) GetRatioInMonthOk() (*int64, bool)`

GetRatioInMonthOk returns a tuple with the RatioInMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatioInMonth

`func (o *UsageBillableSummaryHour) SetRatioInMonth(v int64)`

SetRatioInMonth sets RatioInMonth field to given value.

### HasRatioInMonth

`func (o *UsageBillableSummaryHour) HasRatioInMonth() bool`

HasRatioInMonth returns a boolean if a field has been set.

### GetStartDate

`func (o *UsageBillableSummaryHour) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UsageBillableSummaryHour) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UsageBillableSummaryHour) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UsageBillableSummaryHour) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetUsage

`func (o *UsageBillableSummaryHour) GetUsage() UsageBillableSummaryKeys`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *UsageBillableSummaryHour) GetUsageOk() (*UsageBillableSummaryKeys, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *UsageBillableSummaryHour) SetUsage(v UsageBillableSummaryKeys)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *UsageBillableSummaryHour) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


