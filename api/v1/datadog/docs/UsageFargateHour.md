# UsageFargateHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hour** | Pointer to [**time.Time**](time.Time.md) | The hour for the usage. | [optional] 
**TasksCount** | Pointer to **int64** | Contains the number of Fargate tasks run. | [optional] 

## Methods

### NewUsageFargateHour

`func NewUsageFargateHour() *UsageFargateHour`

NewUsageFargateHour instantiates a new UsageFargateHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageFargateHourWithDefaults

`func NewUsageFargateHourWithDefaults() *UsageFargateHour`

NewUsageFargateHourWithDefaults instantiates a new UsageFargateHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHour

`func (o *UsageFargateHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageFargateHour) GetHourOk() (time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHour

`func (o *UsageFargateHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### SetHour

`func (o *UsageFargateHour) SetHour(v time.Time)`

SetHour gets a reference to the given time.Time and assigns it to the Hour field.

### GetTasksCount

`func (o *UsageFargateHour) GetTasksCount() int64`

GetTasksCount returns the TasksCount field if non-nil, zero value otherwise.

### GetTasksCountOk

`func (o *UsageFargateHour) GetTasksCountOk() (int64, bool)`

GetTasksCountOk returns a tuple with the TasksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTasksCount

`func (o *UsageFargateHour) HasTasksCount() bool`

HasTasksCount returns a boolean if a field has been set.

### SetTasksCount

`func (o *UsageFargateHour) SetTasksCount(v int64)`

SetTasksCount gets a reference to the given int64 and assigns it to the TasksCount field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


