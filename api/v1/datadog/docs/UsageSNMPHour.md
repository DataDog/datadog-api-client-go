# UsageSNMPHour

## Properties

| Name            | Type                     | Description                          | Notes      |
| --------------- | ------------------------ | ------------------------------------ | ---------- |
| **Hour**        | Pointer to **time.Time** | The hour for the usage.              | [optional] |
| **SnmpDevices** | Pointer to **int64**     | Contains the number of SNMP devices. | [optional] |

## Methods

### NewUsageSNMPHour

`func NewUsageSNMPHour() *UsageSNMPHour`

NewUsageSNMPHour instantiates a new UsageSNMPHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageSNMPHourWithDefaults

`func NewUsageSNMPHourWithDefaults() *UsageSNMPHour`

NewUsageSNMPHourWithDefaults instantiates a new UsageSNMPHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetHour

`func (o *UsageSNMPHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageSNMPHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageSNMPHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageSNMPHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetSnmpDevices

`func (o *UsageSNMPHour) GetSnmpDevices() int64`

GetSnmpDevices returns the SnmpDevices field if non-nil, zero value otherwise.

### GetSnmpDevicesOk

`func (o *UsageSNMPHour) GetSnmpDevicesOk() (*int64, bool)`

GetSnmpDevicesOk returns a tuple with the SnmpDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpDevices

`func (o *UsageSNMPHour) SetSnmpDevices(v int64)`

SetSnmpDevices sets SnmpDevices field to given value.

### HasSnmpDevices

`func (o *UsageSNMPHour) HasSnmpDevices() bool`

HasSnmpDevices returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
