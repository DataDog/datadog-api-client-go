# UsageTraceHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hour** | Pointer to [**time.Time**](time.Time.md) | The hour for the usage. | [optional] 
**IndexedEventsCount** | Pointer to **int64** | Contains the number of Analyzed Spans indexed. | [optional] 

## Methods

### GetHour

`func (o *UsageTraceHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageTraceHour) GetHourOk() (time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHour

`func (o *UsageTraceHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### SetHour

`func (o *UsageTraceHour) SetHour(v time.Time)`

SetHour gets a reference to the given time.Time and assigns it to the Hour field.

### GetIndexedEventsCount

`func (o *UsageTraceHour) GetIndexedEventsCount() int64`

GetIndexedEventsCount returns the IndexedEventsCount field if non-nil, zero value otherwise.

### GetIndexedEventsCountOk

`func (o *UsageTraceHour) GetIndexedEventsCountOk() (int64, bool)`

GetIndexedEventsCountOk returns a tuple with the IndexedEventsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIndexedEventsCount

`func (o *UsageTraceHour) HasIndexedEventsCount() bool`

HasIndexedEventsCount returns a boolean if a field has been set.

### SetIndexedEventsCount

`func (o *UsageTraceHour) SetIndexedEventsCount(v int64)`

SetIndexedEventsCount gets a reference to the given int64 and assigns it to the IndexedEventsCount field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


