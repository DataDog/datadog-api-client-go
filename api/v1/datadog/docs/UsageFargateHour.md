# UsageFargateHour

## Properties

| Name                        | Type                     | Description                                            | Notes      |
| --------------------------- | ------------------------ | ------------------------------------------------------ | ---------- |
| **AvgProfiledFargateTasks** | Pointer to **int64**     | The average profiled task count for Fargate Profiling. | [optional] |
| **Hour**                    | Pointer to **time.Time** | The hour for the usage.                                | [optional] |
| **OrgName**                 | Pointer to **string**    | The organization name.                                 | [optional] |
| **PublicId**                | Pointer to **string**    | The organization public ID.                            | [optional] |
| **TasksCount**              | Pointer to **int64**     | The number of Fargate tasks run.                       | [optional] |

## Methods

### NewUsageFargateHour

`func NewUsageFargateHour() *UsageFargateHour`

NewUsageFargateHour instantiates a new UsageFargateHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageFargateHourWithDefaults

`func NewUsageFargateHourWithDefaults() *UsageFargateHour`

NewUsageFargateHourWithDefaults instantiates a new UsageFargateHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAvgProfiledFargateTasks

`func (o *UsageFargateHour) GetAvgProfiledFargateTasks() int64`

GetAvgProfiledFargateTasks returns the AvgProfiledFargateTasks field if non-nil, zero value otherwise.

### GetAvgProfiledFargateTasksOk

`func (o *UsageFargateHour) GetAvgProfiledFargateTasksOk() (*int64, bool)`

GetAvgProfiledFargateTasksOk returns a tuple with the AvgProfiledFargateTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgProfiledFargateTasks

`func (o *UsageFargateHour) SetAvgProfiledFargateTasks(v int64)`

SetAvgProfiledFargateTasks sets AvgProfiledFargateTasks field to given value.

### HasAvgProfiledFargateTasks

`func (o *UsageFargateHour) HasAvgProfiledFargateTasks() bool`

HasAvgProfiledFargateTasks returns a boolean if a field has been set.

### GetHour

`func (o *UsageFargateHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageFargateHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageFargateHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageFargateHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetOrgName

`func (o *UsageFargateHour) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UsageFargateHour) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UsageFargateHour) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UsageFargateHour) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPublicId

`func (o *UsageFargateHour) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UsageFargateHour) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UsageFargateHour) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UsageFargateHour) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetTasksCount

`func (o *UsageFargateHour) GetTasksCount() int64`

GetTasksCount returns the TasksCount field if non-nil, zero value otherwise.

### GetTasksCountOk

`func (o *UsageFargateHour) GetTasksCountOk() (*int64, bool)`

GetTasksCountOk returns a tuple with the TasksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksCount

`func (o *UsageFargateHour) SetTasksCount(v int64)`

SetTasksCount sets TasksCount field to given value.

### HasTasksCount

`func (o *UsageFargateHour) HasTasksCount() bool`

HasTasksCount returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
