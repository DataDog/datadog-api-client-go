# PagerDutyServicesAndSchedules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedules** | Pointer to **[]string** | The array of your schedule URLs. | [optional] 
**Services** | Pointer to [**[]PagerDutyService**](PagerDutyService.md) | The array of PagerDuty service objects. | [optional] 

## Methods

### NewPagerDutyServicesAndSchedules

`func NewPagerDutyServicesAndSchedules() *PagerDutyServicesAndSchedules`

NewPagerDutyServicesAndSchedules instantiates a new PagerDutyServicesAndSchedules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagerDutyServicesAndSchedulesWithDefaults

`func NewPagerDutyServicesAndSchedulesWithDefaults() *PagerDutyServicesAndSchedules`

NewPagerDutyServicesAndSchedulesWithDefaults instantiates a new PagerDutyServicesAndSchedules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedules

`func (o *PagerDutyServicesAndSchedules) GetSchedules() []string`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *PagerDutyServicesAndSchedules) GetSchedulesOk() (*[]string, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *PagerDutyServicesAndSchedules) SetSchedules(v []string)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *PagerDutyServicesAndSchedules) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### GetServices

`func (o *PagerDutyServicesAndSchedules) GetServices() []PagerDutyService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *PagerDutyServicesAndSchedules) GetServicesOk() (*[]PagerDutyService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *PagerDutyServicesAndSchedules) SetServices(v []PagerDutyService)`

SetServices sets Services field to given value.

### HasServices

`func (o *PagerDutyServicesAndSchedules) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


