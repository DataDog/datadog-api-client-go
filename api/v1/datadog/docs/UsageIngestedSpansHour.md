# UsageIngestedSpansHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hour** | Pointer to **time.Time** | The hour for the usage. | [optional] 
**IngestedEventsBytes** | Pointer to **int64** | Contains the total number of bytes ingested during a given hour. | [optional] 

## Methods

### NewUsageIngestedSpansHour

`func NewUsageIngestedSpansHour() *UsageIngestedSpansHour`

NewUsageIngestedSpansHour instantiates a new UsageIngestedSpansHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageIngestedSpansHourWithDefaults

`func NewUsageIngestedSpansHourWithDefaults() *UsageIngestedSpansHour`

NewUsageIngestedSpansHourWithDefaults instantiates a new UsageIngestedSpansHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHour

`func (o *UsageIngestedSpansHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageIngestedSpansHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageIngestedSpansHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageIngestedSpansHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetIngestedEventsBytes

`func (o *UsageIngestedSpansHour) GetIngestedEventsBytes() int64`

GetIngestedEventsBytes returns the IngestedEventsBytes field if non-nil, zero value otherwise.

### GetIngestedEventsBytesOk

`func (o *UsageIngestedSpansHour) GetIngestedEventsBytesOk() (*int64, bool)`

GetIngestedEventsBytesOk returns a tuple with the IngestedEventsBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedEventsBytes

`func (o *UsageIngestedSpansHour) SetIngestedEventsBytes(v int64)`

SetIngestedEventsBytes sets IngestedEventsBytes field to given value.

### HasIngestedEventsBytes

`func (o *UsageIngestedSpansHour) HasIngestedEventsBytes() bool`

HasIngestedEventsBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


