# UsageSDSHour

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Hour** | Pointer to **time.Time** | The hour for the usage. | [optional] 
**LogsScannedBytes** | Pointer to **int64** | The total number of bytes scanned of logs usage by the Sensitive Data Scanner from the start of the given hour’s month until the given hour. | [optional] 
**TotalScannedBytes** | Pointer to **int64** | The total number of bytes scanned across all usage types by the Sensitive Data Scanner from the start of the given hour’s month until the given hour. | [optional] 

## Methods

### NewUsageSDSHour

`func NewUsageSDSHour() *UsageSDSHour`

NewUsageSDSHour instantiates a new UsageSDSHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageSDSHourWithDefaults

`func NewUsageSDSHourWithDefaults() *UsageSDSHour`

NewUsageSDSHourWithDefaults instantiates a new UsageSDSHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetHour

`func (o *UsageSDSHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageSDSHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageSDSHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageSDSHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetLogsScannedBytes

`func (o *UsageSDSHour) GetLogsScannedBytes() int64`

GetLogsScannedBytes returns the LogsScannedBytes field if non-nil, zero value otherwise.

### GetLogsScannedBytesOk

`func (o *UsageSDSHour) GetLogsScannedBytesOk() (*int64, bool)`

GetLogsScannedBytesOk returns a tuple with the LogsScannedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsScannedBytes

`func (o *UsageSDSHour) SetLogsScannedBytes(v int64)`

SetLogsScannedBytes sets LogsScannedBytes field to given value.

### HasLogsScannedBytes

`func (o *UsageSDSHour) HasLogsScannedBytes() bool`

HasLogsScannedBytes returns a boolean if a field has been set.

### GetTotalScannedBytes

`func (o *UsageSDSHour) GetTotalScannedBytes() int64`

GetTotalScannedBytes returns the TotalScannedBytes field if non-nil, zero value otherwise.

### GetTotalScannedBytesOk

`func (o *UsageSDSHour) GetTotalScannedBytesOk() (*int64, bool)`

GetTotalScannedBytesOk returns a tuple with the TotalScannedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalScannedBytes

`func (o *UsageSDSHour) SetTotalScannedBytes(v int64)`

SetTotalScannedBytes sets TotalScannedBytes field to given value.

### HasTotalScannedBytes

`func (o *UsageSDSHour) HasTotalScannedBytes() bool`

HasTotalScannedBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


