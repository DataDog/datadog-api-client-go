# UsageRumSessionsHour

## Properties

| Name                    | Type                     | Description                                                                                        | Notes      |
| ----------------------- | ------------------------ | -------------------------------------------------------------------------------------------------- | ---------- |
| **Hour**                | Pointer to **time.Time** | The hour for the usage.                                                                            | [optional] |
| **OrgName**             | Pointer to **string**    | The organization name.                                                                             | [optional] |
| **PublicId**            | Pointer to **string**    | The organization public ID.                                                                        | [optional] |
| **ReplaySessionCount**  | Pointer to **int64**     | Contains the number of RUM Replay Sessions (data available beginning November 1, 2021).            | [optional] |
| **SessionCount**        | Pointer to **int64**     | Contains the number of browser RUM Lite Sessions.                                                  | [optional] |
| **SessionCountAndroid** | Pointer to **int64**     | Contains the number of mobile RUM Sessions on Android (data available beginning December 1, 2020). | [optional] |
| **SessionCountIos**     | Pointer to **int64**     | Contains the number of mobile RUM Sessions on iOS (data available beginning December 1, 2020).     | [optional] |

## Methods

### NewUsageRumSessionsHour

`func NewUsageRumSessionsHour() *UsageRumSessionsHour`

NewUsageRumSessionsHour instantiates a new UsageRumSessionsHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageRumSessionsHourWithDefaults

`func NewUsageRumSessionsHourWithDefaults() *UsageRumSessionsHour`

NewUsageRumSessionsHourWithDefaults instantiates a new UsageRumSessionsHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetHour

`func (o *UsageRumSessionsHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageRumSessionsHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageRumSessionsHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageRumSessionsHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetOrgName

`func (o *UsageRumSessionsHour) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UsageRumSessionsHour) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UsageRumSessionsHour) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UsageRumSessionsHour) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPublicId

`func (o *UsageRumSessionsHour) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UsageRumSessionsHour) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UsageRumSessionsHour) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UsageRumSessionsHour) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetReplaySessionCount

`func (o *UsageRumSessionsHour) GetReplaySessionCount() int64`

GetReplaySessionCount returns the ReplaySessionCount field if non-nil, zero value otherwise.

### GetReplaySessionCountOk

`func (o *UsageRumSessionsHour) GetReplaySessionCountOk() (*int64, bool)`

GetReplaySessionCountOk returns a tuple with the ReplaySessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaySessionCount

`func (o *UsageRumSessionsHour) SetReplaySessionCount(v int64)`

SetReplaySessionCount sets ReplaySessionCount field to given value.

### HasReplaySessionCount

`func (o *UsageRumSessionsHour) HasReplaySessionCount() bool`

HasReplaySessionCount returns a boolean if a field has been set.

### GetSessionCount

`func (o *UsageRumSessionsHour) GetSessionCount() int64`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *UsageRumSessionsHour) GetSessionCountOk() (*int64, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *UsageRumSessionsHour) SetSessionCount(v int64)`

SetSessionCount sets SessionCount field to given value.

### HasSessionCount

`func (o *UsageRumSessionsHour) HasSessionCount() bool`

HasSessionCount returns a boolean if a field has been set.

### GetSessionCountAndroid

`func (o *UsageRumSessionsHour) GetSessionCountAndroid() int64`

GetSessionCountAndroid returns the SessionCountAndroid field if non-nil, zero value otherwise.

### GetSessionCountAndroidOk

`func (o *UsageRumSessionsHour) GetSessionCountAndroidOk() (*int64, bool)`

GetSessionCountAndroidOk returns a tuple with the SessionCountAndroid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCountAndroid

`func (o *UsageRumSessionsHour) SetSessionCountAndroid(v int64)`

SetSessionCountAndroid sets SessionCountAndroid field to given value.

### HasSessionCountAndroid

`func (o *UsageRumSessionsHour) HasSessionCountAndroid() bool`

HasSessionCountAndroid returns a boolean if a field has been set.

### GetSessionCountIos

`func (o *UsageRumSessionsHour) GetSessionCountIos() int64`

GetSessionCountIos returns the SessionCountIos field if non-nil, zero value otherwise.

### GetSessionCountIosOk

`func (o *UsageRumSessionsHour) GetSessionCountIosOk() (*int64, bool)`

GetSessionCountIosOk returns a tuple with the SessionCountIos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCountIos

`func (o *UsageRumSessionsHour) SetSessionCountIos(v int64)`

SetSessionCountIos sets SessionCountIos field to given value.

### HasSessionCountIos

`func (o *UsageRumSessionsHour) HasSessionCountIos() bool`

HasSessionCountIos returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
