# UsageLogsHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillableIngestedBytes** | Pointer to **int64** | Contains the number of billable log bytes ingested. | [optional] 
**Hour** | Pointer to [**time.Time**](time.Time.md) | The hour for the usage. | [optional] 
**IndexedEventsCount** | Pointer to **int64** | Contains the number of log events indexed. | [optional] 
**IngestedEventsBytes** | Pointer to **int64** | Contains the number of log bytes ingested. | [optional] 

## Methods

### NewUsageLogsHour

`func NewUsageLogsHour() *UsageLogsHour`

NewUsageLogsHour instantiates a new UsageLogsHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageLogsHourWithDefaults

`func NewUsageLogsHourWithDefaults() *UsageLogsHour`

NewUsageLogsHourWithDefaults instantiates a new UsageLogsHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillableIngestedBytes

`func (o *UsageLogsHour) GetBillableIngestedBytes() int64`

GetBillableIngestedBytes returns the BillableIngestedBytes field if non-nil, zero value otherwise.

### GetBillableIngestedBytesOk

`func (o *UsageLogsHour) GetBillableIngestedBytesOk() (*int64, bool)`

GetBillableIngestedBytesOk returns a tuple with the BillableIngestedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableIngestedBytes

`func (o *UsageLogsHour) SetBillableIngestedBytes(v int64)`

SetBillableIngestedBytes sets BillableIngestedBytes field to given value.

### HasBillableIngestedBytes

`func (o *UsageLogsHour) HasBillableIngestedBytes() bool`

HasBillableIngestedBytes returns a boolean if a field has been set.

### GetHour

`func (o *UsageLogsHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageLogsHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageLogsHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageLogsHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetIndexedEventsCount

`func (o *UsageLogsHour) GetIndexedEventsCount() int64`

GetIndexedEventsCount returns the IndexedEventsCount field if non-nil, zero value otherwise.

### GetIndexedEventsCountOk

`func (o *UsageLogsHour) GetIndexedEventsCountOk() (*int64, bool)`

GetIndexedEventsCountOk returns a tuple with the IndexedEventsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedEventsCount

`func (o *UsageLogsHour) SetIndexedEventsCount(v int64)`

SetIndexedEventsCount sets IndexedEventsCount field to given value.

### HasIndexedEventsCount

`func (o *UsageLogsHour) HasIndexedEventsCount() bool`

HasIndexedEventsCount returns a boolean if a field has been set.

### GetIngestedEventsBytes

`func (o *UsageLogsHour) GetIngestedEventsBytes() int64`

GetIngestedEventsBytes returns the IngestedEventsBytes field if non-nil, zero value otherwise.

### GetIngestedEventsBytesOk

`func (o *UsageLogsHour) GetIngestedEventsBytesOk() (*int64, bool)`

GetIngestedEventsBytesOk returns a tuple with the IngestedEventsBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedEventsBytes

`func (o *UsageLogsHour) SetIngestedEventsBytes(v int64)`

SetIngestedEventsBytes sets IngestedEventsBytes field to given value.

### HasIngestedEventsBytes

`func (o *UsageLogsHour) HasIngestedEventsBytes() bool`

HasIngestedEventsBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


