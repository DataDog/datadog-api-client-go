# UsageProfilingHour

## Properties

| Name                       | Type                     | Description                                                                | Notes      |
| -------------------------- | ------------------------ | -------------------------------------------------------------------------- | ---------- |
| **AvgContainerAgentCount** | Pointer to **int64**     | Get average number of container agents for that hour.                      | [optional] |
| **HostCount**              | Pointer to **int64**     | Contains the total number of profiled hosts reporting during a given hour. | [optional] |
| **Hour**                   | Pointer to **time.Time** | The hour for the usage.                                                    | [optional] |

## Methods

### NewUsageProfilingHour

`func NewUsageProfilingHour() *UsageProfilingHour`

NewUsageProfilingHour instantiates a new UsageProfilingHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageProfilingHourWithDefaults

`func NewUsageProfilingHourWithDefaults() *UsageProfilingHour`

NewUsageProfilingHourWithDefaults instantiates a new UsageProfilingHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAvgContainerAgentCount

`func (o *UsageProfilingHour) GetAvgContainerAgentCount() int64`

GetAvgContainerAgentCount returns the AvgContainerAgentCount field if non-nil, zero value otherwise.

### GetAvgContainerAgentCountOk

`func (o *UsageProfilingHour) GetAvgContainerAgentCountOk() (*int64, bool)`

GetAvgContainerAgentCountOk returns a tuple with the AvgContainerAgentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgContainerAgentCount

`func (o *UsageProfilingHour) SetAvgContainerAgentCount(v int64)`

SetAvgContainerAgentCount sets AvgContainerAgentCount field to given value.

### HasAvgContainerAgentCount

`func (o *UsageProfilingHour) HasAvgContainerAgentCount() bool`

HasAvgContainerAgentCount returns a boolean if a field has been set.

### GetHostCount

`func (o *UsageProfilingHour) GetHostCount() int64`

GetHostCount returns the HostCount field if non-nil, zero value otherwise.

### GetHostCountOk

`func (o *UsageProfilingHour) GetHostCountOk() (*int64, bool)`

GetHostCountOk returns a tuple with the HostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCount

`func (o *UsageProfilingHour) SetHostCount(v int64)`

SetHostCount sets HostCount field to given value.

### HasHostCount

`func (o *UsageProfilingHour) HasHostCount() bool`

HasHostCount returns a boolean if a field has been set.

### GetHour

`func (o *UsageProfilingHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageProfilingHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageProfilingHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageProfilingHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
