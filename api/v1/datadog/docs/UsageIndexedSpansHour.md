# UsageIndexedSpansHour

## Properties

| Name                   | Type                     | Description                           | Notes      |
| ---------------------- | ------------------------ | ------------------------------------- | ---------- |
| **Hour**               | Pointer to **time.Time** | The hour for the usage.               | [optional] |
| **IndexedEventsCount** | Pointer to **int64**     | Contains the number of spans indexed. | [optional] |
| **OrgName**            | Pointer to **string**    | The organization name.                | [optional] |
| **PublicId**           | Pointer to **string**    | The organization public ID.           | [optional] |

## Methods

### NewUsageIndexedSpansHour

`func NewUsageIndexedSpansHour() *UsageIndexedSpansHour`

NewUsageIndexedSpansHour instantiates a new UsageIndexedSpansHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageIndexedSpansHourWithDefaults

`func NewUsageIndexedSpansHourWithDefaults() *UsageIndexedSpansHour`

NewUsageIndexedSpansHourWithDefaults instantiates a new UsageIndexedSpansHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetHour

`func (o *UsageIndexedSpansHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageIndexedSpansHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageIndexedSpansHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageIndexedSpansHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetIndexedEventsCount

`func (o *UsageIndexedSpansHour) GetIndexedEventsCount() int64`

GetIndexedEventsCount returns the IndexedEventsCount field if non-nil, zero value otherwise.

### GetIndexedEventsCountOk

`func (o *UsageIndexedSpansHour) GetIndexedEventsCountOk() (*int64, bool)`

GetIndexedEventsCountOk returns a tuple with the IndexedEventsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedEventsCount

`func (o *UsageIndexedSpansHour) SetIndexedEventsCount(v int64)`

SetIndexedEventsCount sets IndexedEventsCount field to given value.

### HasIndexedEventsCount

`func (o *UsageIndexedSpansHour) HasIndexedEventsCount() bool`

HasIndexedEventsCount returns a boolean if a field has been set.

### GetOrgName

`func (o *UsageIndexedSpansHour) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UsageIndexedSpansHour) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UsageIndexedSpansHour) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UsageIndexedSpansHour) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPublicId

`func (o *UsageIndexedSpansHour) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UsageIndexedSpansHour) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UsageIndexedSpansHour) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UsageIndexedSpansHour) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
