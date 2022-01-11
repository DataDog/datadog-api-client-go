# UsageNetworkFlowsHour

## Properties

| Name                   | Type                     | Description                                    | Notes      |
| ---------------------- | ------------------------ | ---------------------------------------------- | ---------- |
| **Hour**               | Pointer to **time.Time** | The hour for the usage.                        | [optional] |
| **IndexedEventsCount** | Pointer to **int64**     | Contains the number of netflow events indexed. | [optional] |

## Methods

### NewUsageNetworkFlowsHour

`func NewUsageNetworkFlowsHour() *UsageNetworkFlowsHour`

NewUsageNetworkFlowsHour instantiates a new UsageNetworkFlowsHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageNetworkFlowsHourWithDefaults

`func NewUsageNetworkFlowsHourWithDefaults() *UsageNetworkFlowsHour`

NewUsageNetworkFlowsHourWithDefaults instantiates a new UsageNetworkFlowsHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetHour

`func (o *UsageNetworkFlowsHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageNetworkFlowsHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageNetworkFlowsHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageNetworkFlowsHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetIndexedEventsCount

`func (o *UsageNetworkFlowsHour) GetIndexedEventsCount() int64`

GetIndexedEventsCount returns the IndexedEventsCount field if non-nil, zero value otherwise.

### GetIndexedEventsCountOk

`func (o *UsageNetworkFlowsHour) GetIndexedEventsCountOk() (*int64, bool)`

GetIndexedEventsCountOk returns a tuple with the IndexedEventsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedEventsCount

`func (o *UsageNetworkFlowsHour) SetIndexedEventsCount(v int64)`

SetIndexedEventsCount sets IndexedEventsCount field to given value.

### HasIndexedEventsCount

`func (o *UsageNetworkFlowsHour) HasIndexedEventsCount() bool`

HasIndexedEventsCount returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
