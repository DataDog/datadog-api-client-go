# DowntimeRecurrence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | Pointer to **int32** | How often to repeat as an integer. For example, to repeat every 3 days, select a type of &#x60;days&#x60; and a period of &#x60;3&#x60;. | [optional] 
**Type** | Pointer to **string** | The type of recurrence. Choose from &#x60;days&#x60;, &#x60;weeks&#x60;, &#x60;months&#x60;, &#x60;years&#x60;. | [optional] 
**UntilDate** | Pointer to **NullableInt64** | The date at which the recurrence should end as a POSIX timestamp. &#x60;until_occurences&#x60; and &#x60;until_date&#x60; are mutually exclusive. | [optional] 
**UntilOccurrences** | Pointer to **NullableInt32** | How many times the downtime is rescheduled. &#x60;until_occurences&#x60; and &#x60;until_date&#x60; are mutually exclusive. | [optional] 
**WeekDays** | Pointer to **[]string** | A list of week days to repeat on. Choose from &#x60;Mon&#x60;, &#x60;Tue&#x60;, &#x60;Wed&#x60;, &#x60;Thu&#x60;, &#x60;Fri&#x60;, &#x60;Sat&#x60; or &#x60;Sun&#x60;. Only applicable when type is weeks. First letter must be capitalized. | [optional] 

## Methods

### NewDowntimeRecurrence

`func NewDowntimeRecurrence() *DowntimeRecurrence`

NewDowntimeRecurrence instantiates a new DowntimeRecurrence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDowntimeRecurrenceWithDefaults

`func NewDowntimeRecurrenceWithDefaults() *DowntimeRecurrence`

NewDowntimeRecurrenceWithDefaults instantiates a new DowntimeRecurrence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *DowntimeRecurrence) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DowntimeRecurrence) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DowntimeRecurrence) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *DowntimeRecurrence) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetType

`func (o *DowntimeRecurrence) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DowntimeRecurrence) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DowntimeRecurrence) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DowntimeRecurrence) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUntilDate

`func (o *DowntimeRecurrence) GetUntilDate() int64`

GetUntilDate returns the UntilDate field if non-nil, zero value otherwise.

### GetUntilDateOk

`func (o *DowntimeRecurrence) GetUntilDateOk() (*int64, bool)`

GetUntilDateOk returns a tuple with the UntilDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntilDate

`func (o *DowntimeRecurrence) SetUntilDate(v int64)`

SetUntilDate sets UntilDate field to given value.

### HasUntilDate

`func (o *DowntimeRecurrence) HasUntilDate() bool`

HasUntilDate returns a boolean if a field has been set.

### SetUntilDateNil

`func (o *DowntimeRecurrence) SetUntilDateNil(b bool)`

 SetUntilDateNil sets the value for UntilDate to be an explicit nil

### UnsetUntilDate
`func (o *DowntimeRecurrence) UnsetUntilDate()`

UnsetUntilDate ensures that no value is present for UntilDate, not even an explicit nil
### GetUntilOccurrences

`func (o *DowntimeRecurrence) GetUntilOccurrences() int32`

GetUntilOccurrences returns the UntilOccurrences field if non-nil, zero value otherwise.

### GetUntilOccurrencesOk

`func (o *DowntimeRecurrence) GetUntilOccurrencesOk() (*int32, bool)`

GetUntilOccurrencesOk returns a tuple with the UntilOccurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntilOccurrences

`func (o *DowntimeRecurrence) SetUntilOccurrences(v int32)`

SetUntilOccurrences sets UntilOccurrences field to given value.

### HasUntilOccurrences

`func (o *DowntimeRecurrence) HasUntilOccurrences() bool`

HasUntilOccurrences returns a boolean if a field has been set.

### SetUntilOccurrencesNil

`func (o *DowntimeRecurrence) SetUntilOccurrencesNil(b bool)`

 SetUntilOccurrencesNil sets the value for UntilOccurrences to be an explicit nil

### UnsetUntilOccurrences
`func (o *DowntimeRecurrence) UnsetUntilOccurrences()`

UnsetUntilOccurrences ensures that no value is present for UntilOccurrences, not even an explicit nil
### GetWeekDays

`func (o *DowntimeRecurrence) GetWeekDays() []string`

GetWeekDays returns the WeekDays field if non-nil, zero value otherwise.

### GetWeekDaysOk

`func (o *DowntimeRecurrence) GetWeekDaysOk() (*[]string, bool)`

GetWeekDaysOk returns a tuple with the WeekDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekDays

`func (o *DowntimeRecurrence) SetWeekDays(v []string)`

SetWeekDays sets WeekDays field to given value.

### HasWeekDays

`func (o *DowntimeRecurrence) HasWeekDays() bool`

HasWeekDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


