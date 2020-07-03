# UsageBillableSummaryBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountBillableUsage** | Pointer to **int64** | The total account usage. | [optional] 
**ElapsedUsageHours** | Pointer to **int64** | Elapsed usage hours for some billable product. | [optional] 
**FirstBillableUsageHour** | Pointer to [**time.Time**](time.Time.md) | The first billable hour for the org. | [optional] 
**LastBillableUsageHour** | Pointer to [**time.Time**](time.Time.md) | The last billable hour for the org. | [optional] 
**OrgBillableUsage** | Pointer to **int64** | The number of units used within the billable timeframe. | [optional] 
**PercentageInAccount** | Pointer to **float64** | The percentage of account usage the org represents. | [optional] 
**UsageUnit** | Pointer to **string** | Units pertaining to the usage. | [optional] 

## Methods

### NewUsageBillableSummaryBody

`func NewUsageBillableSummaryBody() *UsageBillableSummaryBody`

NewUsageBillableSummaryBody instantiates a new UsageBillableSummaryBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageBillableSummaryBodyWithDefaults

`func NewUsageBillableSummaryBodyWithDefaults() *UsageBillableSummaryBody`

NewUsageBillableSummaryBodyWithDefaults instantiates a new UsageBillableSummaryBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountBillableUsage

`func (o *UsageBillableSummaryBody) GetAccountBillableUsage() int64`

GetAccountBillableUsage returns the AccountBillableUsage field if non-nil, zero value otherwise.

### GetAccountBillableUsageOk

`func (o *UsageBillableSummaryBody) GetAccountBillableUsageOk() (*int64, bool)`

GetAccountBillableUsageOk returns a tuple with the AccountBillableUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountBillableUsage

`func (o *UsageBillableSummaryBody) SetAccountBillableUsage(v int64)`

SetAccountBillableUsage sets AccountBillableUsage field to given value.

### HasAccountBillableUsage

`func (o *UsageBillableSummaryBody) HasAccountBillableUsage() bool`

HasAccountBillableUsage returns a boolean if a field has been set.

### GetElapsedUsageHours

`func (o *UsageBillableSummaryBody) GetElapsedUsageHours() int64`

GetElapsedUsageHours returns the ElapsedUsageHours field if non-nil, zero value otherwise.

### GetElapsedUsageHoursOk

`func (o *UsageBillableSummaryBody) GetElapsedUsageHoursOk() (*int64, bool)`

GetElapsedUsageHoursOk returns a tuple with the ElapsedUsageHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedUsageHours

`func (o *UsageBillableSummaryBody) SetElapsedUsageHours(v int64)`

SetElapsedUsageHours sets ElapsedUsageHours field to given value.

### HasElapsedUsageHours

`func (o *UsageBillableSummaryBody) HasElapsedUsageHours() bool`

HasElapsedUsageHours returns a boolean if a field has been set.

### GetFirstBillableUsageHour

`func (o *UsageBillableSummaryBody) GetFirstBillableUsageHour() time.Time`

GetFirstBillableUsageHour returns the FirstBillableUsageHour field if non-nil, zero value otherwise.

### GetFirstBillableUsageHourOk

`func (o *UsageBillableSummaryBody) GetFirstBillableUsageHourOk() (*time.Time, bool)`

GetFirstBillableUsageHourOk returns a tuple with the FirstBillableUsageHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstBillableUsageHour

`func (o *UsageBillableSummaryBody) SetFirstBillableUsageHour(v time.Time)`

SetFirstBillableUsageHour sets FirstBillableUsageHour field to given value.

### HasFirstBillableUsageHour

`func (o *UsageBillableSummaryBody) HasFirstBillableUsageHour() bool`

HasFirstBillableUsageHour returns a boolean if a field has been set.

### GetLastBillableUsageHour

`func (o *UsageBillableSummaryBody) GetLastBillableUsageHour() time.Time`

GetLastBillableUsageHour returns the LastBillableUsageHour field if non-nil, zero value otherwise.

### GetLastBillableUsageHourOk

`func (o *UsageBillableSummaryBody) GetLastBillableUsageHourOk() (*time.Time, bool)`

GetLastBillableUsageHourOk returns a tuple with the LastBillableUsageHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBillableUsageHour

`func (o *UsageBillableSummaryBody) SetLastBillableUsageHour(v time.Time)`

SetLastBillableUsageHour sets LastBillableUsageHour field to given value.

### HasLastBillableUsageHour

`func (o *UsageBillableSummaryBody) HasLastBillableUsageHour() bool`

HasLastBillableUsageHour returns a boolean if a field has been set.

### GetOrgBillableUsage

`func (o *UsageBillableSummaryBody) GetOrgBillableUsage() int64`

GetOrgBillableUsage returns the OrgBillableUsage field if non-nil, zero value otherwise.

### GetOrgBillableUsageOk

`func (o *UsageBillableSummaryBody) GetOrgBillableUsageOk() (*int64, bool)`

GetOrgBillableUsageOk returns a tuple with the OrgBillableUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgBillableUsage

`func (o *UsageBillableSummaryBody) SetOrgBillableUsage(v int64)`

SetOrgBillableUsage sets OrgBillableUsage field to given value.

### HasOrgBillableUsage

`func (o *UsageBillableSummaryBody) HasOrgBillableUsage() bool`

HasOrgBillableUsage returns a boolean if a field has been set.

### GetPercentageInAccount

`func (o *UsageBillableSummaryBody) GetPercentageInAccount() float64`

GetPercentageInAccount returns the PercentageInAccount field if non-nil, zero value otherwise.

### GetPercentageInAccountOk

`func (o *UsageBillableSummaryBody) GetPercentageInAccountOk() (*float64, bool)`

GetPercentageInAccountOk returns a tuple with the PercentageInAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageInAccount

`func (o *UsageBillableSummaryBody) SetPercentageInAccount(v float64)`

SetPercentageInAccount sets PercentageInAccount field to given value.

### HasPercentageInAccount

`func (o *UsageBillableSummaryBody) HasPercentageInAccount() bool`

HasPercentageInAccount returns a boolean if a field has been set.

### GetUsageUnit

`func (o *UsageBillableSummaryBody) GetUsageUnit() string`

GetUsageUnit returns the UsageUnit field if non-nil, zero value otherwise.

### GetUsageUnitOk

`func (o *UsageBillableSummaryBody) GetUsageUnitOk() (*string, bool)`

GetUsageUnitOk returns a tuple with the UsageUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageUnit

`func (o *UsageBillableSummaryBody) SetUsageUnit(v string)`

SetUsageUnit sets UsageUnit field to given value.

### HasUsageUnit

`func (o *UsageBillableSummaryBody) HasUsageUnit() bool`

HasUsageUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


