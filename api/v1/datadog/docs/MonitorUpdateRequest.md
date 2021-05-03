# MonitorUpdateRequest

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Created** | Pointer to **time.Time** | Timestamp of the monitor creation. | [optional] [readonly] 
**Creator** | Pointer to [**Creator**](Creator.md) |  | [optional] 
**Deleted** | Pointer to **NullableTime** | Whether or not the monitor is deleted. (Always &#x60;null&#x60;) | [optional] [readonly] 
**Id** | Pointer to **int64** | ID of this monitor. | [optional] [readonly] 
**Message** | Pointer to **string** | A message to include with notifications for this monitor. | [optional] 
**Modified** | Pointer to **time.Time** | Last timestamp when the monitor was edited. | [optional] [readonly] 
**Multi** | Pointer to **bool** | Whether or not the monitor is broken down on different groups. | [optional] [readonly] 
**Name** | Pointer to **string** | The monitor name. | [optional] 
**Options** | Pointer to [**MonitorOptions**](MonitorOptions.md) |  | [optional] 
**OverallState** | Pointer to [**MonitorOverallStates**](MonitorOverallStates.md) |  | [optional] 
**Priority** | Pointer to **int64** | Integer from 1 (high) to 5 (low) indicating alert severity. | [optional] 
**Query** | Pointer to **string** | The monitor query. | [optional] 
**RestrictedRoles** | Pointer to **[]string** | A list of role identifiers that can be pulled from the Roles API. Cannot be used with &#x60;locked&#x60; option. | [optional] 
**State** | Pointer to [**MonitorState**](MonitorState.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Tags associated to your monitor. | [optional] 
**Type** | Pointer to [**MonitorType**](MonitorType.md) |  | [optional] 

## Methods

### NewMonitorUpdateRequest

`func NewMonitorUpdateRequest() *MonitorUpdateRequest`

NewMonitorUpdateRequest instantiates a new MonitorUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorUpdateRequestWithDefaults

`func NewMonitorUpdateRequestWithDefaults() *MonitorUpdateRequest`

NewMonitorUpdateRequestWithDefaults instantiates a new MonitorUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *MonitorUpdateRequest) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MonitorUpdateRequest) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *MonitorUpdateRequest) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *MonitorUpdateRequest) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreator

`func (o *MonitorUpdateRequest) GetCreator() Creator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *MonitorUpdateRequest) GetCreatorOk() (*Creator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *MonitorUpdateRequest) SetCreator(v Creator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *MonitorUpdateRequest) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDeleted

`func (o *MonitorUpdateRequest) GetDeleted() time.Time`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *MonitorUpdateRequest) GetDeletedOk() (*time.Time, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *MonitorUpdateRequest) SetDeleted(v time.Time)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *MonitorUpdateRequest) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *MonitorUpdateRequest) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *MonitorUpdateRequest) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetId

`func (o *MonitorUpdateRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MonitorUpdateRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MonitorUpdateRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MonitorUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *MonitorUpdateRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MonitorUpdateRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MonitorUpdateRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MonitorUpdateRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetModified

`func (o *MonitorUpdateRequest) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *MonitorUpdateRequest) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *MonitorUpdateRequest) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *MonitorUpdateRequest) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetMulti

`func (o *MonitorUpdateRequest) GetMulti() bool`

GetMulti returns the Multi field if non-nil, zero value otherwise.

### GetMultiOk

`func (o *MonitorUpdateRequest) GetMultiOk() (*bool, bool)`

GetMultiOk returns a tuple with the Multi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulti

`func (o *MonitorUpdateRequest) SetMulti(v bool)`

SetMulti sets Multi field to given value.

### HasMulti

`func (o *MonitorUpdateRequest) HasMulti() bool`

HasMulti returns a boolean if a field has been set.

### GetName

`func (o *MonitorUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitorUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitorUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MonitorUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *MonitorUpdateRequest) GetOptions() MonitorOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MonitorUpdateRequest) GetOptionsOk() (*MonitorOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MonitorUpdateRequest) SetOptions(v MonitorOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MonitorUpdateRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOverallState

`func (o *MonitorUpdateRequest) GetOverallState() MonitorOverallStates`

GetOverallState returns the OverallState field if non-nil, zero value otherwise.

### GetOverallStateOk

`func (o *MonitorUpdateRequest) GetOverallStateOk() (*MonitorOverallStates, bool)`

GetOverallStateOk returns a tuple with the OverallState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallState

`func (o *MonitorUpdateRequest) SetOverallState(v MonitorOverallStates)`

SetOverallState sets OverallState field to given value.

### HasOverallState

`func (o *MonitorUpdateRequest) HasOverallState() bool`

HasOverallState returns a boolean if a field has been set.

### GetPriority

`func (o *MonitorUpdateRequest) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MonitorUpdateRequest) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MonitorUpdateRequest) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MonitorUpdateRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQuery

`func (o *MonitorUpdateRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MonitorUpdateRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MonitorUpdateRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *MonitorUpdateRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetRestrictedRoles

`func (o *MonitorUpdateRequest) GetRestrictedRoles() []string`

GetRestrictedRoles returns the RestrictedRoles field if non-nil, zero value otherwise.

### GetRestrictedRolesOk

`func (o *MonitorUpdateRequest) GetRestrictedRolesOk() (*[]string, bool)`

GetRestrictedRolesOk returns a tuple with the RestrictedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedRoles

`func (o *MonitorUpdateRequest) SetRestrictedRoles(v []string)`

SetRestrictedRoles sets RestrictedRoles field to given value.

### HasRestrictedRoles

`func (o *MonitorUpdateRequest) HasRestrictedRoles() bool`

HasRestrictedRoles returns a boolean if a field has been set.

### GetState

`func (o *MonitorUpdateRequest) GetState() MonitorState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MonitorUpdateRequest) GetStateOk() (*MonitorState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MonitorUpdateRequest) SetState(v MonitorState)`

SetState sets State field to given value.

### HasState

`func (o *MonitorUpdateRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *MonitorUpdateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MonitorUpdateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MonitorUpdateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MonitorUpdateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *MonitorUpdateRequest) GetType() MonitorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorUpdateRequest) GetTypeOk() (*MonitorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorUpdateRequest) SetType(v MonitorType)`

SetType sets Type field to given value.

### HasType

`func (o *MonitorUpdateRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


