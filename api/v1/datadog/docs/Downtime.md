# Downtime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | If a scheduled downtime currently exists. | [optional] [readonly] 
**Canceled** | Pointer to **NullableInt64** | If a scheduled downtime is cancelled. | [optional] [readonly] 
**CreatorId** | Pointer to **int32** | User ID of the downtime creator. | [optional] [readonly] 
**Disabled** | Pointer to **bool** | If a downtime has been disabled. | [optional] 
**DowntimeType** | Pointer to **int32** | &#x60;0&#x60; for a downtime applied on &#x60;*&#x60; or all, &#x60;1&#x60; when the downtime is only scoped to hosts, or &#x60;2&#x60; when the downtime is scoped to anything but hosts. | [optional] [readonly] 
**End** | Pointer to **NullableInt64** | POSIX timestamp to end the downtime. If not provided, the downtime is in effect indefinitely until you cancel it. | [optional] 
**Id** | Pointer to **int64** | The downtime ID. | [optional] [readonly] 
**Message** | Pointer to **string** | A message to include with notifications for this downtime. Email notifications can be sent to specific users by using the same &#x60;@username&#x60; notation as events. | [optional] 
**MonitorId** | Pointer to **NullableInt64** | A single monitor to which the downtime applies. If not provided, the downtime applies to all monitors. | [optional] 
**MonitorTags** | Pointer to **[]string** | A comma-separated list of monitor tags. For example, tags that are applied directly to monitors, not tags that are used in monitor queries (which are filtered by the scope parameter), to which the downtime applies. The resulting downtime applies to monitors that match ALL provided monitor tags. For example, &#x60;service:postgres&#x60; **AND** &#x60;team:frontend&#x60;. | [optional] 
**ParentId** | Pointer to **NullableInt64** | ID of the parent Downtime. | [optional] 
**Recurrence** | Pointer to [**NullableDowntimeRecurrence**](DowntimeRecurrence.md) |  | [optional] 
**Scope** | Pointer to **[]string** | The scope(s) to which the downtime applies. For example, &#x60;host:app2&#x60;. Provide multiple scopes as a comma-separated list like &#x60;env:dev,env:prod&#x60;. The resulting downtime applies to sources that matches ALL provided scopes (&#x60;env:dev&#x60; **AND** &#x60;env:prod&#x60;). | [optional] 
**Start** | Pointer to **int64** | POSIX timestamp to start the downtime. If not provided, the downtime starts the moment it is created. | [optional] 
**Timezone** | Pointer to **string** | The timezone for the downtime. | [optional] 
**UpdaterId** | Pointer to **NullableInt32** | ID of the last user that updated the downtime. | [optional] [readonly] 

## Methods

### NewDowntime

`func NewDowntime() *Downtime`

NewDowntime instantiates a new Downtime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDowntimeWithDefaults

`func NewDowntimeWithDefaults() *Downtime`

NewDowntimeWithDefaults instantiates a new Downtime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *Downtime) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Downtime) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Downtime) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Downtime) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCanceled

`func (o *Downtime) GetCanceled() int64`

GetCanceled returns the Canceled field if non-nil, zero value otherwise.

### GetCanceledOk

`func (o *Downtime) GetCanceledOk() (*int64, bool)`

GetCanceledOk returns a tuple with the Canceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceled

`func (o *Downtime) SetCanceled(v int64)`

SetCanceled sets Canceled field to given value.

### HasCanceled

`func (o *Downtime) HasCanceled() bool`

HasCanceled returns a boolean if a field has been set.

### SetCanceledNil

`func (o *Downtime) SetCanceledNil(b bool)`

 SetCanceledNil sets the value for Canceled to be an explicit nil

### UnsetCanceled
`func (o *Downtime) UnsetCanceled()`

UnsetCanceled ensures that no value is present for Canceled, not even an explicit nil
### GetCreatorId

`func (o *Downtime) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Downtime) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Downtime) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *Downtime) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetDisabled

`func (o *Downtime) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Downtime) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Downtime) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Downtime) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDowntimeType

`func (o *Downtime) GetDowntimeType() int32`

GetDowntimeType returns the DowntimeType field if non-nil, zero value otherwise.

### GetDowntimeTypeOk

`func (o *Downtime) GetDowntimeTypeOk() (*int32, bool)`

GetDowntimeTypeOk returns a tuple with the DowntimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntimeType

`func (o *Downtime) SetDowntimeType(v int32)`

SetDowntimeType sets DowntimeType field to given value.

### HasDowntimeType

`func (o *Downtime) HasDowntimeType() bool`

HasDowntimeType returns a boolean if a field has been set.

### GetEnd

`func (o *Downtime) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Downtime) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Downtime) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Downtime) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *Downtime) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *Downtime) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetId

`func (o *Downtime) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Downtime) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Downtime) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Downtime) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *Downtime) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Downtime) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Downtime) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Downtime) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMonitorId

`func (o *Downtime) GetMonitorId() int64`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *Downtime) GetMonitorIdOk() (*int64, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *Downtime) SetMonitorId(v int64)`

SetMonitorId sets MonitorId field to given value.

### HasMonitorId

`func (o *Downtime) HasMonitorId() bool`

HasMonitorId returns a boolean if a field has been set.

### SetMonitorIdNil

`func (o *Downtime) SetMonitorIdNil(b bool)`

 SetMonitorIdNil sets the value for MonitorId to be an explicit nil

### UnsetMonitorId
`func (o *Downtime) UnsetMonitorId()`

UnsetMonitorId ensures that no value is present for MonitorId, not even an explicit nil
### GetMonitorTags

`func (o *Downtime) GetMonitorTags() []string`

GetMonitorTags returns the MonitorTags field if non-nil, zero value otherwise.

### GetMonitorTagsOk

`func (o *Downtime) GetMonitorTagsOk() (*[]string, bool)`

GetMonitorTagsOk returns a tuple with the MonitorTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTags

`func (o *Downtime) SetMonitorTags(v []string)`

SetMonitorTags sets MonitorTags field to given value.

### HasMonitorTags

`func (o *Downtime) HasMonitorTags() bool`

HasMonitorTags returns a boolean if a field has been set.

### GetParentId

`func (o *Downtime) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Downtime) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Downtime) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Downtime) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *Downtime) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *Downtime) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetRecurrence

`func (o *Downtime) GetRecurrence() DowntimeRecurrence`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *Downtime) GetRecurrenceOk() (*DowntimeRecurrence, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *Downtime) SetRecurrence(v DowntimeRecurrence)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *Downtime) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### SetRecurrenceNil

`func (o *Downtime) SetRecurrenceNil(b bool)`

 SetRecurrenceNil sets the value for Recurrence to be an explicit nil

### UnsetRecurrence
`func (o *Downtime) UnsetRecurrence()`

UnsetRecurrence ensures that no value is present for Recurrence, not even an explicit nil
### GetScope

`func (o *Downtime) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Downtime) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Downtime) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Downtime) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStart

`func (o *Downtime) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Downtime) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Downtime) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *Downtime) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTimezone

`func (o *Downtime) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Downtime) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Downtime) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Downtime) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUpdaterId

`func (o *Downtime) GetUpdaterId() int32`

GetUpdaterId returns the UpdaterId field if non-nil, zero value otherwise.

### GetUpdaterIdOk

`func (o *Downtime) GetUpdaterIdOk() (*int32, bool)`

GetUpdaterIdOk returns a tuple with the UpdaterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterId

`func (o *Downtime) SetUpdaterId(v int32)`

SetUpdaterId sets UpdaterId field to given value.

### HasUpdaterId

`func (o *Downtime) HasUpdaterId() bool`

HasUpdaterId returns a boolean if a field has been set.

### SetUpdaterIdNil

`func (o *Downtime) SetUpdaterIdNil(b bool)`

 SetUpdaterIdNil sets the value for UpdaterId to be an explicit nil

### UnsetUpdaterId
`func (o *Downtime) UnsetUpdaterId()`

UnsetUpdaterId ensures that no value is present for UpdaterId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


