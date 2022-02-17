# UsageNetworkFlowsHour

## Properties

| Name                   | Type                     | Description                                    | Notes      |
| ---------------------- | ------------------------ | ---------------------------------------------- | ---------- |
| **Hour**               | Pointer to **time.Time** | The hour for the usage.                        | [optional] |
| **IndexedEventsCount** | Pointer to **int64**     | Contains the number of netflow events indexed. | [optional] |
| **OrgName**            | Pointer to **string**    | The organization name.                         | [optional] |
| **PublicId**           | Pointer to **string**    | The organization public ID.                    | [optional] |

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

### GetOrgName

`func (o *UsageNetworkFlowsHour) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UsageNetworkFlowsHour) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UsageNetworkFlowsHour) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UsageNetworkFlowsHour) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPublicId

`func (o *UsageNetworkFlowsHour) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UsageNetworkFlowsHour) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UsageNetworkFlowsHour) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UsageNetworkFlowsHour) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
