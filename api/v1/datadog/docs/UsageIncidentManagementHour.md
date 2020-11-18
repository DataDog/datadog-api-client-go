# UsageIncidentManagementHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hour** | Pointer to [**time.Time**](time.Time.md) | The hour for the usage. | [optional] 
**MonthlyActiveUsers** | Pointer to **int64** | Contains the total number monthly active users from the start of the given hour&#39;s month until the given hour. | [optional] 

## Methods

### NewUsageIncidentManagementHour

`func NewUsageIncidentManagementHour() *UsageIncidentManagementHour`

NewUsageIncidentManagementHour instantiates a new UsageIncidentManagementHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageIncidentManagementHourWithDefaults

`func NewUsageIncidentManagementHourWithDefaults() *UsageIncidentManagementHour`

NewUsageIncidentManagementHourWithDefaults instantiates a new UsageIncidentManagementHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHour

`func (o *UsageIncidentManagementHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageIncidentManagementHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageIncidentManagementHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageIncidentManagementHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMonthlyActiveUsers

`func (o *UsageIncidentManagementHour) GetMonthlyActiveUsers() int64`

GetMonthlyActiveUsers returns the MonthlyActiveUsers field if non-nil, zero value otherwise.

### GetMonthlyActiveUsersOk

`func (o *UsageIncidentManagementHour) GetMonthlyActiveUsersOk() (*int64, bool)`

GetMonthlyActiveUsersOk returns a tuple with the MonthlyActiveUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyActiveUsers

`func (o *UsageIncidentManagementHour) SetMonthlyActiveUsers(v int64)`

SetMonthlyActiveUsers sets MonthlyActiveUsers field to given value.

### HasMonthlyActiveUsers

`func (o *UsageIncidentManagementHour) HasMonthlyActiveUsers() bool`

HasMonthlyActiveUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


