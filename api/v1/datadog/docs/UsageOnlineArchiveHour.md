# UsageOnlineArchiveHour

## Properties

| Name                         | Type                     | Description                                            | Notes      |
| ---------------------------- | ------------------------ | ------------------------------------------------------ | ---------- |
| **Hour**                     | Pointer to **time.Time** | The hour for the usage.                                | [optional] |
| **OnlineArchiveEventsCount** | Pointer to **int32**     | Total count of online archived events within the hour. | [optional] |
| **OrgName**                  | Pointer to **string**    | The organization name.                                 | [optional] |
| **PublicId**                 | Pointer to **string**    | The organization public ID.                            | [optional] |

## Methods

### NewUsageOnlineArchiveHour

`func NewUsageOnlineArchiveHour() *UsageOnlineArchiveHour`

NewUsageOnlineArchiveHour instantiates a new UsageOnlineArchiveHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageOnlineArchiveHourWithDefaults

`func NewUsageOnlineArchiveHourWithDefaults() *UsageOnlineArchiveHour`

NewUsageOnlineArchiveHourWithDefaults instantiates a new UsageOnlineArchiveHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetHour

`func (o *UsageOnlineArchiveHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageOnlineArchiveHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageOnlineArchiveHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageOnlineArchiveHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetOnlineArchiveEventsCount

`func (o *UsageOnlineArchiveHour) GetOnlineArchiveEventsCount() int32`

GetOnlineArchiveEventsCount returns the OnlineArchiveEventsCount field if non-nil, zero value otherwise.

### GetOnlineArchiveEventsCountOk

`func (o *UsageOnlineArchiveHour) GetOnlineArchiveEventsCountOk() (*int32, bool)`

GetOnlineArchiveEventsCountOk returns a tuple with the OnlineArchiveEventsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineArchiveEventsCount

`func (o *UsageOnlineArchiveHour) SetOnlineArchiveEventsCount(v int32)`

SetOnlineArchiveEventsCount sets OnlineArchiveEventsCount field to given value.

### HasOnlineArchiveEventsCount

`func (o *UsageOnlineArchiveHour) HasOnlineArchiveEventsCount() bool`

HasOnlineArchiveEventsCount returns a boolean if a field has been set.

### GetOrgName

`func (o *UsageOnlineArchiveHour) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UsageOnlineArchiveHour) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UsageOnlineArchiveHour) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UsageOnlineArchiveHour) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPublicId

`func (o *UsageOnlineArchiveHour) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UsageOnlineArchiveHour) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UsageOnlineArchiveHour) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UsageOnlineArchiveHour) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
