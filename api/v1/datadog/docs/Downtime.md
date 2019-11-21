# Downtime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] [readonly] 
**Canceled** | Pointer to **NullableInt64** |  | [optional] [readonly] 
**CreatorId** | Pointer to **int32** |  | [optional] [readonly] 
**Disabled** | Pointer to **bool** |  | [optional] 
**DowntimeType** | Pointer to **int32** |  | [optional] [readonly] 
**End** | Pointer to **NullableInt64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**Message** | Pointer to **string** |  | [optional] 
**MonitorId** | Pointer to **NullableInt64** |  | [optional] 
**MonitorTags** | Pointer to **[]string** |  | [optional] 
**ParentId** | Pointer to **NullableInt64** |  | [optional] 
**Recurrence** | Pointer to [**NullableDowntimeRecurrence**](DowntimeRecurrence.md) |  | [optional] 
**Scope** | Pointer to **[]string** |  | [optional] 
**Start** | Pointer to **int64** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**UpdaterId** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### GetActive

`func (o *Downtime) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Downtime) GetActiveOk() (bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActive

`func (o *Downtime) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActive

`func (o *Downtime) SetActive(v bool)`

SetActive gets a reference to the given bool and assigns it to the Active field.

### GetCanceled

`func (o *Downtime) GetCanceled() NullableInt64`

GetCanceled returns the Canceled field if non-nil, zero value otherwise.

### GetCanceledOk

`func (o *Downtime) GetCanceledOk() (NullableInt64, bool)`

GetCanceledOk returns a tuple with the Canceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCanceled

`func (o *Downtime) HasCanceled() bool`

HasCanceled returns a boolean if a field has been set.

### SetCanceled

`func (o *Downtime) SetCanceled(v NullableInt64)`

SetCanceled gets a reference to the given NullableInt64 and assigns it to the Canceled field.

### SetCanceledExplicitNull

`func (o *Downtime) SetCanceledExplicitNull(b bool)`

SetCanceledExplicitNull (un)sets Canceled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Canceled value is set to nil even if false is passed
### GetCreatorId

`func (o *Downtime) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Downtime) GetCreatorIdOk() (int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatorId

`func (o *Downtime) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorId

`func (o *Downtime) SetCreatorId(v int32)`

SetCreatorId gets a reference to the given int32 and assigns it to the CreatorId field.

### GetDisabled

`func (o *Downtime) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Downtime) GetDisabledOk() (bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisabled

`func (o *Downtime) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabled

`func (o *Downtime) SetDisabled(v bool)`

SetDisabled gets a reference to the given bool and assigns it to the Disabled field.

### GetDowntimeType

`func (o *Downtime) GetDowntimeType() int32`

GetDowntimeType returns the DowntimeType field if non-nil, zero value otherwise.

### GetDowntimeTypeOk

`func (o *Downtime) GetDowntimeTypeOk() (int32, bool)`

GetDowntimeTypeOk returns a tuple with the DowntimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDowntimeType

`func (o *Downtime) HasDowntimeType() bool`

HasDowntimeType returns a boolean if a field has been set.

### SetDowntimeType

`func (o *Downtime) SetDowntimeType(v int32)`

SetDowntimeType gets a reference to the given int32 and assigns it to the DowntimeType field.

### GetEnd

`func (o *Downtime) GetEnd() NullableInt64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Downtime) GetEndOk() (NullableInt64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnd

`func (o *Downtime) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEnd

`func (o *Downtime) SetEnd(v NullableInt64)`

SetEnd gets a reference to the given NullableInt64 and assigns it to the End field.

### SetEndExplicitNull

`func (o *Downtime) SetEndExplicitNull(b bool)`

SetEndExplicitNull (un)sets End to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The End value is set to nil even if false is passed
### GetId

`func (o *Downtime) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Downtime) GetIdOk() (int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Downtime) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Downtime) SetId(v int64)`

SetId gets a reference to the given int64 and assigns it to the Id field.

### GetMessage

`func (o *Downtime) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Downtime) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *Downtime) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *Downtime) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.

### GetMonitorId

`func (o *Downtime) GetMonitorId() NullableInt64`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *Downtime) GetMonitorIdOk() (NullableInt64, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMonitorId

`func (o *Downtime) HasMonitorId() bool`

HasMonitorId returns a boolean if a field has been set.

### SetMonitorId

`func (o *Downtime) SetMonitorId(v NullableInt64)`

SetMonitorId gets a reference to the given NullableInt64 and assigns it to the MonitorId field.

### SetMonitorIdExplicitNull

`func (o *Downtime) SetMonitorIdExplicitNull(b bool)`

SetMonitorIdExplicitNull (un)sets MonitorId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MonitorId value is set to nil even if false is passed
### GetMonitorTags

`func (o *Downtime) GetMonitorTags() []string`

GetMonitorTags returns the MonitorTags field if non-nil, zero value otherwise.

### GetMonitorTagsOk

`func (o *Downtime) GetMonitorTagsOk() ([]string, bool)`

GetMonitorTagsOk returns a tuple with the MonitorTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMonitorTags

`func (o *Downtime) HasMonitorTags() bool`

HasMonitorTags returns a boolean if a field has been set.

### SetMonitorTags

`func (o *Downtime) SetMonitorTags(v []string)`

SetMonitorTags gets a reference to the given []string and assigns it to the MonitorTags field.

### GetParentId

`func (o *Downtime) GetParentId() NullableInt64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Downtime) GetParentIdOk() (NullableInt64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentId

`func (o *Downtime) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentId

`func (o *Downtime) SetParentId(v NullableInt64)`

SetParentId gets a reference to the given NullableInt64 and assigns it to the ParentId field.

### SetParentIdExplicitNull

`func (o *Downtime) SetParentIdExplicitNull(b bool)`

SetParentIdExplicitNull (un)sets ParentId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentId value is set to nil even if false is passed
### GetRecurrence

`func (o *Downtime) GetRecurrence() NullableDowntimeRecurrence`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *Downtime) GetRecurrenceOk() (NullableDowntimeRecurrence, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecurrence

`func (o *Downtime) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### SetRecurrence

`func (o *Downtime) SetRecurrence(v NullableDowntimeRecurrence)`

SetRecurrence gets a reference to the given NullableDowntimeRecurrence and assigns it to the Recurrence field.

### SetRecurrenceExplicitNull

`func (o *Downtime) SetRecurrenceExplicitNull(b bool)`

SetRecurrenceExplicitNull (un)sets Recurrence to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Recurrence value is set to nil even if false is passed
### GetScope

`func (o *Downtime) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Downtime) GetScopeOk() ([]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScope

`func (o *Downtime) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScope

`func (o *Downtime) SetScope(v []string)`

SetScope gets a reference to the given []string and assigns it to the Scope field.

### GetStart

`func (o *Downtime) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Downtime) GetStartOk() (int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStart

`func (o *Downtime) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStart

`func (o *Downtime) SetStart(v int64)`

SetStart gets a reference to the given int64 and assigns it to the Start field.

### GetTimezone

`func (o *Downtime) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Downtime) GetTimezoneOk() (string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimezone

`func (o *Downtime) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezone

`func (o *Downtime) SetTimezone(v string)`

SetTimezone gets a reference to the given string and assigns it to the Timezone field.

### GetUpdaterId

`func (o *Downtime) GetUpdaterId() int32`

GetUpdaterId returns the UpdaterId field if non-nil, zero value otherwise.

### GetUpdaterIdOk

`func (o *Downtime) GetUpdaterIdOk() (int32, bool)`

GetUpdaterIdOk returns a tuple with the UpdaterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdaterId

`func (o *Downtime) HasUpdaterId() bool`

HasUpdaterId returns a boolean if a field has been set.

### SetUpdaterId

`func (o *Downtime) SetUpdaterId(v int32)`

SetUpdaterId gets a reference to the given int32 and assigns it to the UpdaterId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


