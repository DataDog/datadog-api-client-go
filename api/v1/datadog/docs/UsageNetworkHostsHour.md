# UsageNetworkHostsHour

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**HostCount** | Pointer to **int64** | Contains the number of active NPM hosts. | [optional] 
**Hour** | Pointer to **time.Time** | The hour for the usage. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


