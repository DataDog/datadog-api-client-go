# UsageSyntheticsHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckCallsCount** | Pointer to **int64** | Contains the number of Synthetics API tests run. | [optional] 
**Hour** | Pointer to [**time.Time**](time.Time.md) | The hour for the usage. | [optional] 

## Methods

### GetCheckCallsCount

`func (o *UsageSyntheticsHour) GetCheckCallsCount() int64`

GetCheckCallsCount returns the CheckCallsCount field if non-nil, zero value otherwise.

### GetCheckCallsCountOk

`func (o *UsageSyntheticsHour) GetCheckCallsCountOk() (int64, bool)`

GetCheckCallsCountOk returns a tuple with the CheckCallsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCheckCallsCount

`func (o *UsageSyntheticsHour) HasCheckCallsCount() bool`

HasCheckCallsCount returns a boolean if a field has been set.

### SetCheckCallsCount

`func (o *UsageSyntheticsHour) SetCheckCallsCount(v int64)`

SetCheckCallsCount gets a reference to the given int64 and assigns it to the CheckCallsCount field.

### GetHour

`func (o *UsageSyntheticsHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageSyntheticsHour) GetHourOk() (time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHour

`func (o *UsageSyntheticsHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### SetHour

`func (o *UsageSyntheticsHour) SetHour(v time.Time)`

SetHour gets a reference to the given time.Time and assigns it to the Hour field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


