# DowntimeRecurrence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UntilDate** | Pointer to **NullableInt64** |  | [optional] 
**UntilOccurrences** | Pointer to **NullableInt32** |  | [optional] 
**WeekDays** | Pointer to **[]string** |  | [optional] 

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

`func (o *DowntimeRecurrence) GetPeriodOk() (int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriod

`func (o *DowntimeRecurrence) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriod

`func (o *DowntimeRecurrence) SetPeriod(v int32)`

SetPeriod gets a reference to the given int32 and assigns it to the Period field.

### GetType

`func (o *DowntimeRecurrence) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DowntimeRecurrence) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *DowntimeRecurrence) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *DowntimeRecurrence) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetUntilDate

`func (o *DowntimeRecurrence) GetUntilDate() NullableInt64`

GetUntilDate returns the UntilDate field if non-nil, zero value otherwise.

### GetUntilDateOk

`func (o *DowntimeRecurrence) GetUntilDateOk() (NullableInt64, bool)`

GetUntilDateOk returns a tuple with the UntilDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUntilDate

`func (o *DowntimeRecurrence) HasUntilDate() bool`

HasUntilDate returns a boolean if a field has been set.

### SetUntilDate

`func (o *DowntimeRecurrence) SetUntilDate(v NullableInt64)`

SetUntilDate gets a reference to the given NullableInt64 and assigns it to the UntilDate field.

### SetUntilDateExplicitNull

`func (o *DowntimeRecurrence) SetUntilDateExplicitNull(b bool)`

SetUntilDateExplicitNull (un)sets UntilDate to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UntilDate value is set to nil even if false is passed
### GetUntilOccurrences

`func (o *DowntimeRecurrence) GetUntilOccurrences() NullableInt32`

GetUntilOccurrences returns the UntilOccurrences field if non-nil, zero value otherwise.

### GetUntilOccurrencesOk

`func (o *DowntimeRecurrence) GetUntilOccurrencesOk() (NullableInt32, bool)`

GetUntilOccurrencesOk returns a tuple with the UntilOccurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUntilOccurrences

`func (o *DowntimeRecurrence) HasUntilOccurrences() bool`

HasUntilOccurrences returns a boolean if a field has been set.

### SetUntilOccurrences

`func (o *DowntimeRecurrence) SetUntilOccurrences(v NullableInt32)`

SetUntilOccurrences gets a reference to the given NullableInt32 and assigns it to the UntilOccurrences field.

### SetUntilOccurrencesExplicitNull

`func (o *DowntimeRecurrence) SetUntilOccurrencesExplicitNull(b bool)`

SetUntilOccurrencesExplicitNull (un)sets UntilOccurrences to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UntilOccurrences value is set to nil even if false is passed
### GetWeekDays

`func (o *DowntimeRecurrence) GetWeekDays() []string`

GetWeekDays returns the WeekDays field if non-nil, zero value otherwise.

### GetWeekDaysOk

`func (o *DowntimeRecurrence) GetWeekDaysOk() ([]string, bool)`

GetWeekDaysOk returns a tuple with the WeekDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWeekDays

`func (o *DowntimeRecurrence) HasWeekDays() bool`

HasWeekDays returns a boolean if a field has been set.

### SetWeekDays

`func (o *DowntimeRecurrence) SetWeekDays(v []string)`

SetWeekDays gets a reference to the given []string and assigns it to the WeekDays field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


