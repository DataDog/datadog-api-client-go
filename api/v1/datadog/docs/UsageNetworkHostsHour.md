# UsageNetworkHostsHour

## Properties

| Name          | Type                     | Description                              | Notes      |
| ------------- | ------------------------ | ---------------------------------------- | ---------- |
| **HostCount** | Pointer to **int64**     | Contains the number of active NPM hosts. | [optional] |
| **Hour**      | Pointer to **time.Time** | The hour for the usage.                  | [optional] |
| **OrgName**   | Pointer to **string**    | The organization name.                   | [optional] |
| **PublicId**  | Pointer to **string**    | The organization public ID.              | [optional] |

## Methods

### NewUsageNetworkHostsHour

`func NewUsageNetworkHostsHour() *UsageNetworkHostsHour`

NewUsageNetworkHostsHour instantiates a new UsageNetworkHostsHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageNetworkHostsHourWithDefaults

`func NewUsageNetworkHostsHourWithDefaults() *UsageNetworkHostsHour`

NewUsageNetworkHostsHourWithDefaults instantiates a new UsageNetworkHostsHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetHostCount

`func (o *UsageNetworkHostsHour) GetHostCount() int64`

GetHostCount returns the HostCount field if non-nil, zero value otherwise.

### GetHostCountOk

`func (o *UsageNetworkHostsHour) GetHostCountOk() (*int64, bool)`

GetHostCountOk returns a tuple with the HostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCount

`func (o *UsageNetworkHostsHour) SetHostCount(v int64)`

SetHostCount sets HostCount field to given value.

### HasHostCount

`func (o *UsageNetworkHostsHour) HasHostCount() bool`

HasHostCount returns a boolean if a field has been set.

### GetHour

`func (o *UsageNetworkHostsHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageNetworkHostsHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageNetworkHostsHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageNetworkHostsHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetOrgName

`func (o *UsageNetworkHostsHour) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UsageNetworkHostsHour) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UsageNetworkHostsHour) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UsageNetworkHostsHour) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPublicId

`func (o *UsageNetworkHostsHour) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UsageNetworkHostsHour) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UsageNetworkHostsHour) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UsageNetworkHostsHour) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
