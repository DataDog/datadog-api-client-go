# ServiceLevelObjective

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | Creation timestamp (unix time in seconds) Always included in service level objective responses. | [optional] 
**Creator** | Pointer to [**Creator**](Creator.md) |  | [optional] 
**Description** | Pointer to **NullableString** | A user-defined description of the service level objective. Always included in service level objective responses (but may be null). Optional in create/update requests. | [optional] 
**Groups** | Pointer to **[]string** | A list of (up to 20) monitor groups (e.g. [\&quot;env:prod,role:mysql\&quot;]) that narrows the scope of a monitor service level objective. Included in service level objective responses if it is nonempty. Optional in create/update requests for monitor service level objectives, but may only be used when then length of the \&quot;monitor_ids\&quot; field is one. | [optional] 
**Id** | Pointer to **string** | A unique identifier for the service level objective object. Always included in service level objective responses. Required for update requests. | [optional] 
**ModifiedAt** | Pointer to **int64** | Modification timestamp (unix time in seconds) Always included in service level objective responses. | [optional] 
**MonitorIds** | Pointer to **[]int64** | A list of monitor ids that defines the scope of a monitor service level objective. Required if type is \&quot;monitor\&quot;. | [optional] 
**MonitorTags** | Pointer to **[]string** | The union of monitor tags for all monitors referenced by the \&quot;monitor_ids\&quot; field. Always included in service level objective responses for monitor service level objectives (but may be empty). Ignored in create/update requests. Does not affect which monitors are included in the service level objective (that is determined entirely by the monitor_ids field). | [optional] 
**Name** | Pointer to **string** | The name of the service level objective object. | 
**Query** | Pointer to [**ServiceLevelObjectiveQuery**](ServiceLevelObjective_query.md) |  | [optional] 
**Tags** | Pointer to **[]string** | A list of tags (e.g. \&quot;env:prod\&quot;) associated with this service level objective. Always included in service level objective responses (but may be empty). Optional in create/update requests. | [optional] 
**Thresholds** | Pointer to [**[]SloThreshold**](SLOThreshold.md) | The thresholds (timeframes and associated targets) for this service level objective object. | 
**Type** | Pointer to **string** | The type of the service level objective. | 
**TypeId** | Pointer to **int32** | A numeric representation of the type of the service level objective (0 for monitor, 1 for metric). Always included in service level objective responses. Ignored in create/update requests. | [optional] 

## Methods

### GetCreatedAt

`func (o *ServiceLevelObjective) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceLevelObjective) GetCreatedAtOk() (int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedAt

`func (o *ServiceLevelObjective) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAt

`func (o *ServiceLevelObjective) SetCreatedAt(v int64)`

SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.

### GetCreator

`func (o *ServiceLevelObjective) GetCreator() Creator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *ServiceLevelObjective) GetCreatorOk() (Creator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreator

`func (o *ServiceLevelObjective) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### SetCreator

`func (o *ServiceLevelObjective) SetCreator(v Creator)`

SetCreator gets a reference to the given Creator and assigns it to the Creator field.

### GetDescription

`func (o *ServiceLevelObjective) GetDescription() NullableString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceLevelObjective) GetDescriptionOk() (NullableString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *ServiceLevelObjective) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *ServiceLevelObjective) SetDescription(v NullableString)`

SetDescription gets a reference to the given NullableString and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *ServiceLevelObjective) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetGroups

`func (o *ServiceLevelObjective) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ServiceLevelObjective) GetGroupsOk() ([]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGroups

`func (o *ServiceLevelObjective) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroups

`func (o *ServiceLevelObjective) SetGroups(v []string)`

SetGroups gets a reference to the given []string and assigns it to the Groups field.

### GetId

`func (o *ServiceLevelObjective) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceLevelObjective) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *ServiceLevelObjective) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *ServiceLevelObjective) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetModifiedAt

`func (o *ServiceLevelObjective) GetModifiedAt() int64`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ServiceLevelObjective) GetModifiedAtOk() (int64, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModifiedAt

`func (o *ServiceLevelObjective) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAt

`func (o *ServiceLevelObjective) SetModifiedAt(v int64)`

SetModifiedAt gets a reference to the given int64 and assigns it to the ModifiedAt field.

### GetMonitorIds

`func (o *ServiceLevelObjective) GetMonitorIds() []int64`

GetMonitorIds returns the MonitorIds field if non-nil, zero value otherwise.

### GetMonitorIdsOk

`func (o *ServiceLevelObjective) GetMonitorIdsOk() ([]int64, bool)`

GetMonitorIdsOk returns a tuple with the MonitorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMonitorIds

`func (o *ServiceLevelObjective) HasMonitorIds() bool`

HasMonitorIds returns a boolean if a field has been set.

### SetMonitorIds

`func (o *ServiceLevelObjective) SetMonitorIds(v []int64)`

SetMonitorIds gets a reference to the given []int64 and assigns it to the MonitorIds field.

### GetMonitorTags

`func (o *ServiceLevelObjective) GetMonitorTags() []string`

GetMonitorTags returns the MonitorTags field if non-nil, zero value otherwise.

### GetMonitorTagsOk

`func (o *ServiceLevelObjective) GetMonitorTagsOk() ([]string, bool)`

GetMonitorTagsOk returns a tuple with the MonitorTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMonitorTags

`func (o *ServiceLevelObjective) HasMonitorTags() bool`

HasMonitorTags returns a boolean if a field has been set.

### SetMonitorTags

`func (o *ServiceLevelObjective) SetMonitorTags(v []string)`

SetMonitorTags gets a reference to the given []string and assigns it to the MonitorTags field.

### GetName

`func (o *ServiceLevelObjective) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceLevelObjective) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *ServiceLevelObjective) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *ServiceLevelObjective) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetQuery

`func (o *ServiceLevelObjective) GetQuery() ServiceLevelObjectiveQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ServiceLevelObjective) GetQueryOk() (ServiceLevelObjectiveQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQuery

`func (o *ServiceLevelObjective) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQuery

`func (o *ServiceLevelObjective) SetQuery(v ServiceLevelObjectiveQuery)`

SetQuery gets a reference to the given ServiceLevelObjectiveQuery and assigns it to the Query field.

### GetTags

`func (o *ServiceLevelObjective) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceLevelObjective) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *ServiceLevelObjective) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *ServiceLevelObjective) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.

### GetThresholds

`func (o *ServiceLevelObjective) GetThresholds() []SloThreshold`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *ServiceLevelObjective) GetThresholdsOk() ([]SloThreshold, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasThresholds

`func (o *ServiceLevelObjective) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### SetThresholds

`func (o *ServiceLevelObjective) SetThresholds(v []SloThreshold)`

SetThresholds gets a reference to the given []SloThreshold and assigns it to the Thresholds field.

### GetType

`func (o *ServiceLevelObjective) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceLevelObjective) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *ServiceLevelObjective) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *ServiceLevelObjective) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetTypeId

`func (o *ServiceLevelObjective) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *ServiceLevelObjective) GetTypeIdOk() (int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTypeId

`func (o *ServiceLevelObjective) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### SetTypeId

`func (o *ServiceLevelObjective) SetTypeId(v int32)`

SetTypeId gets a reference to the given int32 and assigns it to the TypeId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


