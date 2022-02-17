# UsageLogsByRetentionHour

## Properties

| Name                             | Type                  | Description                                                                          | Notes      |
| -------------------------------- | --------------------- | ------------------------------------------------------------------------------------ | ---------- |
| **IndexedEventsCount**           | Pointer to **int64**  | Total logs indexed with this retention period during a given hour.                   | [optional] |
| **LiveIndexedEventsCount**       | Pointer to **int64**  | Live logs indexed with this retention period during a given hour.                    | [optional] |
| **OrgName**                      | Pointer to **string** | The organization name.                                                               | [optional] |
| **PublicId**                     | Pointer to **string** | The organization public ID.                                                          | [optional] |
| **RehydratedIndexedEventsCount** | Pointer to **int64**  | Rehydrated logs indexed with this retention period during a given hour.              | [optional] |
| **Retention**                    | Pointer to **string** | The retention period in days or \&quot;custom\&quot; for all custom retention usage. | [optional] |

## Methods

### NewUsageLogsByRetentionHour

`func NewUsageLogsByRetentionHour() *UsageLogsByRetentionHour`

NewUsageLogsByRetentionHour instantiates a new UsageLogsByRetentionHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageLogsByRetentionHourWithDefaults

`func NewUsageLogsByRetentionHourWithDefaults() *UsageLogsByRetentionHour`

NewUsageLogsByRetentionHourWithDefaults instantiates a new UsageLogsByRetentionHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetIndexedEventsCount

`func (o *UsageLogsByRetentionHour) GetIndexedEventsCount() int64`

GetIndexedEventsCount returns the IndexedEventsCount field if non-nil, zero value otherwise.

### GetIndexedEventsCountOk

`func (o *UsageLogsByRetentionHour) GetIndexedEventsCountOk() (*int64, bool)`

GetIndexedEventsCountOk returns a tuple with the IndexedEventsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedEventsCount

`func (o *UsageLogsByRetentionHour) SetIndexedEventsCount(v int64)`

SetIndexedEventsCount sets IndexedEventsCount field to given value.

### HasIndexedEventsCount

`func (o *UsageLogsByRetentionHour) HasIndexedEventsCount() bool`

HasIndexedEventsCount returns a boolean if a field has been set.

### GetLiveIndexedEventsCount

`func (o *UsageLogsByRetentionHour) GetLiveIndexedEventsCount() int64`

GetLiveIndexedEventsCount returns the LiveIndexedEventsCount field if non-nil, zero value otherwise.

### GetLiveIndexedEventsCountOk

`func (o *UsageLogsByRetentionHour) GetLiveIndexedEventsCountOk() (*int64, bool)`

GetLiveIndexedEventsCountOk returns a tuple with the LiveIndexedEventsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveIndexedEventsCount

`func (o *UsageLogsByRetentionHour) SetLiveIndexedEventsCount(v int64)`

SetLiveIndexedEventsCount sets LiveIndexedEventsCount field to given value.

### HasLiveIndexedEventsCount

`func (o *UsageLogsByRetentionHour) HasLiveIndexedEventsCount() bool`

HasLiveIndexedEventsCount returns a boolean if a field has been set.

### GetOrgName

`func (o *UsageLogsByRetentionHour) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UsageLogsByRetentionHour) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UsageLogsByRetentionHour) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UsageLogsByRetentionHour) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPublicId

`func (o *UsageLogsByRetentionHour) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UsageLogsByRetentionHour) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UsageLogsByRetentionHour) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UsageLogsByRetentionHour) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetRehydratedIndexedEventsCount

`func (o *UsageLogsByRetentionHour) GetRehydratedIndexedEventsCount() int64`

GetRehydratedIndexedEventsCount returns the RehydratedIndexedEventsCount field if non-nil, zero value otherwise.

### GetRehydratedIndexedEventsCountOk

`func (o *UsageLogsByRetentionHour) GetRehydratedIndexedEventsCountOk() (*int64, bool)`

GetRehydratedIndexedEventsCountOk returns a tuple with the RehydratedIndexedEventsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRehydratedIndexedEventsCount

`func (o *UsageLogsByRetentionHour) SetRehydratedIndexedEventsCount(v int64)`

SetRehydratedIndexedEventsCount sets RehydratedIndexedEventsCount field to given value.

### HasRehydratedIndexedEventsCount

`func (o *UsageLogsByRetentionHour) HasRehydratedIndexedEventsCount() bool`

HasRehydratedIndexedEventsCount returns a boolean if a field has been set.

### GetRetention

`func (o *UsageLogsByRetentionHour) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *UsageLogsByRetentionHour) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *UsageLogsByRetentionHour) SetRetention(v string)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *UsageLogsByRetentionHour) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
