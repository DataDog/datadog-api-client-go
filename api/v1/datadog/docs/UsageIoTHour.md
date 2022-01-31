# UsageIoTHour

## Properties

| Name               | Type                     | Description                                          | Notes      |
| ------------------ | ------------------------ | ---------------------------------------------------- | ---------- |
| **Hour**           | Pointer to **time.Time** | The hour for the usage.                              | [optional] |
| **IotDeviceCount** | Pointer to **int64**     | The total number of IoT devices during a given hour. | [optional] |
| **OrgName**        | Pointer to **string**    | The organization name.                               | [optional] |
| **PublicId**       | Pointer to **string**    | The organization public ID.                          | [optional] |

## Methods

### NewUsageIoTHour

`func NewUsageIoTHour() *UsageIoTHour`

NewUsageIoTHour instantiates a new UsageIoTHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageIoTHourWithDefaults

`func NewUsageIoTHourWithDefaults() *UsageIoTHour`

NewUsageIoTHourWithDefaults instantiates a new UsageIoTHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetHour

`func (o *UsageIoTHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageIoTHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageIoTHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageIoTHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetIotDeviceCount

`func (o *UsageIoTHour) GetIotDeviceCount() int64`

GetIotDeviceCount returns the IotDeviceCount field if non-nil, zero value otherwise.

### GetIotDeviceCountOk

`func (o *UsageIoTHour) GetIotDeviceCountOk() (*int64, bool)`

GetIotDeviceCountOk returns a tuple with the IotDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIotDeviceCount

`func (o *UsageIoTHour) SetIotDeviceCount(v int64)`

SetIotDeviceCount sets IotDeviceCount field to given value.

### HasIotDeviceCount

`func (o *UsageIoTHour) HasIotDeviceCount() bool`

HasIotDeviceCount returns a boolean if a field has been set.

### GetOrgName

`func (o *UsageIoTHour) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UsageIoTHour) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UsageIoTHour) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UsageIoTHour) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPublicId

`func (o *UsageIoTHour) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UsageIoTHour) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UsageIoTHour) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UsageIoTHour) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
