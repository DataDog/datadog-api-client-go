# UsageCWSHour

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**CwsContainerCount** | Pointer to **int64** | The total number of Cloud Workload Security container hours from the start of the given hour’s month until the given hour. | [optional] 
**CwsHostCount** | Pointer to **int64** | The total number of Cloud Workload Security host hours from the start of the given hour’s month until the given hour. | [optional] 
**Hour** | Pointer to **time.Time** | The hour for the usage. | [optional] 

## Methods

### NewUsageCWSHour

`func NewUsageCWSHour() *UsageCWSHour`

NewUsageCWSHour instantiates a new UsageCWSHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageCWSHourWithDefaults

`func NewUsageCWSHourWithDefaults() *UsageCWSHour`

NewUsageCWSHourWithDefaults instantiates a new UsageCWSHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCwsContainerCount

`func (o *UsageCWSHour) GetCwsContainerCount() int64`

GetCwsContainerCount returns the CwsContainerCount field if non-nil, zero value otherwise.

### GetCwsContainerCountOk

`func (o *UsageCWSHour) GetCwsContainerCountOk() (*int64, bool)`

GetCwsContainerCountOk returns a tuple with the CwsContainerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwsContainerCount

`func (o *UsageCWSHour) SetCwsContainerCount(v int64)`

SetCwsContainerCount sets CwsContainerCount field to given value.

### HasCwsContainerCount

`func (o *UsageCWSHour) HasCwsContainerCount() bool`

HasCwsContainerCount returns a boolean if a field has been set.

### GetCwsHostCount

`func (o *UsageCWSHour) GetCwsHostCount() int64`

GetCwsHostCount returns the CwsHostCount field if non-nil, zero value otherwise.

### GetCwsHostCountOk

`func (o *UsageCWSHour) GetCwsHostCountOk() (*int64, bool)`

GetCwsHostCountOk returns a tuple with the CwsHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwsHostCount

`func (o *UsageCWSHour) SetCwsHostCount(v int64)`

SetCwsHostCount sets CwsHostCount field to given value.

### HasCwsHostCount

`func (o *UsageCWSHour) HasCwsHostCount() bool`

HasCwsHostCount returns a boolean if a field has been set.

### GetHour

`func (o *UsageCWSHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageCWSHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageCWSHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageCWSHour) HasHour() bool`

HasHour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


