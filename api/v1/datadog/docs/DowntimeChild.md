# DowntimeChild

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Active** | Pointer to **bool** | If a scheduled downtime currently exists. | [optional] [readonly] 
**Canceled** | Pointer to **NullableInt64** | If a scheduled downtime is canceled. | [optional] [readonly] 
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
**Timezone** | Pointer to **string** | The timezone in which to display the downtime&#39;s start and end times in Datadog applications. | [optional] 
**UpdaterId** | Pointer to **NullableInt32** | ID of the last user that updated the downtime. | [optional] [readonly] 

## Methods

### NewDowntimeChild

`func NewDowntimeChild() *DowntimeChild`

NewDowntimeChild instantiates a new DowntimeChild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDowntimeChildWithDefaults

`func NewDowntimeChildWithDefaults() *DowntimeChild`

NewDowntimeChildWithDefaults instantiates a new DowntimeChild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *DowntimeChild) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DowntimeChild) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DowntimeChild) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *DowntimeChild) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCanceled

`func (o *DowntimeChild) GetCanceled() int64`

GetCanceled returns the Canceled field if non-nil, zero value otherwise.

### GetCanceledOk

`func (o *DowntimeChild) GetCanceledOk() (*int64, bool)`

GetCanceledOk returns a tuple with the Canceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceled

`func (o *DowntimeChild) SetCanceled(v int64)`

SetCanceled sets Canceled field to given value.

### HasCanceled

`func (o *DowntimeChild) HasCanceled() bool`

HasCanceled returns a boolean if a field has been set.

### SetCanceledNil

`func (o *DowntimeChild) SetCanceledNil(b bool)`

 SetCanceledNil sets the value for Canceled to be an explicit nil

### UnsetCanceled
`func (o *DowntimeChild) UnsetCanceled()`

UnsetCanceled ensures that no value is present for Canceled, not even an explicit nil
### GetCreatorId

`func (o *DowntimeChild) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *DowntimeChild) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *DowntimeChild) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *DowntimeChild) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetDisabled

`func (o *DowntimeChild) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *DowntimeChild) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *DowntimeChild) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *DowntimeChild) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDowntimeType

`func (o *DowntimeChild) GetDowntimeType() int32`

GetDowntimeType returns the DowntimeType field if non-nil, zero value otherwise.

### GetDowntimeTypeOk

`func (o *DowntimeChild) GetDowntimeTypeOk() (*int32, bool)`

GetDowntimeTypeOk returns a tuple with the DowntimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntimeType

`func (o *DowntimeChild) SetDowntimeType(v int32)`

SetDowntimeType sets DowntimeType field to given value.

### HasDowntimeType

`func (o *DowntimeChild) HasDowntimeType() bool`

HasDowntimeType returns a boolean if a field has been set.

### GetEnd

`func (o *DowntimeChild) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *DowntimeChild) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *DowntimeChild) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *DowntimeChild) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *DowntimeChild) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *DowntimeChild) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetId

`func (o *DowntimeChild) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DowntimeChild) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DowntimeChild) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DowntimeChild) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *DowntimeChild) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DowntimeChild) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DowntimeChild) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DowntimeChild) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMonitorId

`func (o *DowntimeChild) GetMonitorId() int64`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *DowntimeChild) GetMonitorIdOk() (*int64, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *DowntimeChild) SetMonitorId(v int64)`

SetMonitorId sets MonitorId field to given value.

### HasMonitorId

`func (o *DowntimeChild) HasMonitorId() bool`

HasMonitorId returns a boolean if a field has been set.

### SetMonitorIdNil

`func (o *DowntimeChild) SetMonitorIdNil(b bool)`

 SetMonitorIdNil sets the value for MonitorId to be an explicit nil

### UnsetMonitorId
`func (o *DowntimeChild) UnsetMonitorId()`

UnsetMonitorId ensures that no value is present for MonitorId, not even an explicit nil
### GetMonitorTags

`func (o *DowntimeChild) GetMonitorTags() []string`

GetMonitorTags returns the MonitorTags field if non-nil, zero value otherwise.

### GetMonitorTagsOk

`func (o *DowntimeChild) GetMonitorTagsOk() (*[]string, bool)`

GetMonitorTagsOk returns a tuple with the MonitorTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTags

`func (o *DowntimeChild) SetMonitorTags(v []string)`

SetMonitorTags sets MonitorTags field to given value.

### HasMonitorTags

`func (o *DowntimeChild) HasMonitorTags() bool`

HasMonitorTags returns a boolean if a field has been set.

### GetParentId

`func (o *DowntimeChild) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *DowntimeChild) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *DowntimeChild) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *DowntimeChild) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *DowntimeChild) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *DowntimeChild) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetRecurrence

`func (o *DowntimeChild) GetRecurrence() DowntimeRecurrence`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *DowntimeChild) GetRecurrenceOk() (*DowntimeRecurrence, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *DowntimeChild) SetRecurrence(v DowntimeRecurrence)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *DowntimeChild) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### SetRecurrenceNil

`func (o *DowntimeChild) SetRecurrenceNil(b bool)`

 SetRecurrenceNil sets the value for Recurrence to be an explicit nil

### UnsetRecurrence
`func (o *DowntimeChild) UnsetRecurrence()`

UnsetRecurrence ensures that no value is present for Recurrence, not even an explicit nil
### GetScope

`func (o *DowntimeChild) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DowntimeChild) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DowntimeChild) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *DowntimeChild) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStart

`func (o *DowntimeChild) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *DowntimeChild) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *DowntimeChild) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *DowntimeChild) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTimezone

`func (o *DowntimeChild) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *DowntimeChild) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *DowntimeChild) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *DowntimeChild) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUpdaterId

`func (o *DowntimeChild) GetUpdaterId() int32`

GetUpdaterId returns the UpdaterId field if non-nil, zero value otherwise.

### GetUpdaterIdOk

`func (o *DowntimeChild) GetUpdaterIdOk() (*int32, bool)`

GetUpdaterIdOk returns a tuple with the UpdaterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterId

`func (o *DowntimeChild) SetUpdaterId(v int32)`

SetUpdaterId sets UpdaterId field to given value.

### HasUpdaterId

`func (o *DowntimeChild) HasUpdaterId() bool`

HasUpdaterId returns a boolean if a field has been set.

### SetUpdaterIdNil

`func (o *DowntimeChild) SetUpdaterIdNil(b bool)`

 SetUpdaterIdNil sets the value for UpdaterId to be an explicit nil

### UnsetUpdaterId
`func (o *DowntimeChild) UnsetUpdaterId()`

UnsetUpdaterId ensures that no value is present for UpdaterId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


