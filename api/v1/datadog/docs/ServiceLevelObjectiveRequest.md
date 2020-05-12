# ServiceLevelObjectiveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | Creation timestamp (UNIX time in seconds)  Always included in service level objective responses. | [optional] [readonly] 
**Creator** | Pointer to [**Creator**](Creator.md) |  | [optional] 
**Description** | Pointer to **NullableString** | A user-defined description of the service level objective.  Always included in service level objective responses (but may be &#x60;null&#x60;). Optional in create/update requests. | [optional] 
**Groups** | Pointer to **[]string** | A list of (up to 20) monitor groups that narrow the scope of a monitor service level objective.  Included in service level objective responses if it is not empty. Optional in create/update requests for monitor service level objectives, but may only be used when then length of the &#x60;monitor_ids&#x60; field is one. | [optional] 
**Id** | Pointer to **string** | A unique identifier for the service level objective object.  Always included in service level objective responses. | [optional] [readonly] 
**ModifiedAt** | Pointer to **int64** | Modification timestamp (UNIX time in seconds)  Always included in service level objective responses. | [optional] [readonly] 
**MonitorIds** | Pointer to **[]int64** | A list of monitor ids that defines the scope of a monitor service level objective. **Required if type is &#x60;monitor&#x60;**. | [optional] 
**Name** | Pointer to **string** | The name of the service level objective object. | 
**Query** | Pointer to [**ServiceLevelObjectiveQuery**](ServiceLevelObjectiveQuery.md) |  | [optional] 
**Tags** | Pointer to **[]string** | A list of tags associated with this service level objective. Always included in service level objective responses (but may be empty). Optional in create/update requests. | [optional] 
**Thresholds** | Pointer to [**[]SLOThreshold**](SLOThreshold.md) | The thresholds (timeframes and associated targets) for this service level objective object. | 
**Type** | Pointer to [**SLOType**](SLOType.md) |  | 

## Methods

### NewServiceLevelObjectiveRequest

`func NewServiceLevelObjectiveRequest(name string, thresholds []SLOThreshold, type_ SLOType, ) *ServiceLevelObjectiveRequest`

NewServiceLevelObjectiveRequest instantiates a new ServiceLevelObjectiveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceLevelObjectiveRequestWithDefaults

`func NewServiceLevelObjectiveRequestWithDefaults() *ServiceLevelObjectiveRequest`

NewServiceLevelObjectiveRequestWithDefaults instantiates a new ServiceLevelObjectiveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServiceLevelObjectiveRequest) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceLevelObjectiveRequest) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceLevelObjectiveRequest) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceLevelObjectiveRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *ServiceLevelObjectiveRequest) GetCreator() Creator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *ServiceLevelObjectiveRequest) GetCreatorOk() (*Creator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *ServiceLevelObjectiveRequest) SetCreator(v Creator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *ServiceLevelObjectiveRequest) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceLevelObjectiveRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceLevelObjectiveRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceLevelObjectiveRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceLevelObjectiveRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ServiceLevelObjectiveRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ServiceLevelObjectiveRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetGroups

`func (o *ServiceLevelObjectiveRequest) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ServiceLevelObjectiveRequest) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ServiceLevelObjectiveRequest) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ServiceLevelObjectiveRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetId

`func (o *ServiceLevelObjectiveRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceLevelObjectiveRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceLevelObjectiveRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceLevelObjectiveRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *ServiceLevelObjectiveRequest) GetModifiedAt() int64`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ServiceLevelObjectiveRequest) GetModifiedAtOk() (*int64, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ServiceLevelObjectiveRequest) SetModifiedAt(v int64)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *ServiceLevelObjectiveRequest) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetMonitorIds

`func (o *ServiceLevelObjectiveRequest) GetMonitorIds() []int64`

GetMonitorIds returns the MonitorIds field if non-nil, zero value otherwise.

### GetMonitorIdsOk

`func (o *ServiceLevelObjectiveRequest) GetMonitorIdsOk() (*[]int64, bool)`

GetMonitorIdsOk returns a tuple with the MonitorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorIds

`func (o *ServiceLevelObjectiveRequest) SetMonitorIds(v []int64)`

SetMonitorIds sets MonitorIds field to given value.

### HasMonitorIds

`func (o *ServiceLevelObjectiveRequest) HasMonitorIds() bool`

HasMonitorIds returns a boolean if a field has been set.

### GetName

`func (o *ServiceLevelObjectiveRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceLevelObjectiveRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceLevelObjectiveRequest) SetName(v string)`

SetName sets Name field to given value.


### GetQuery

`func (o *ServiceLevelObjectiveRequest) GetQuery() ServiceLevelObjectiveQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ServiceLevelObjectiveRequest) GetQueryOk() (*ServiceLevelObjectiveQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ServiceLevelObjectiveRequest) SetQuery(v ServiceLevelObjectiveQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ServiceLevelObjectiveRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTags

`func (o *ServiceLevelObjectiveRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceLevelObjectiveRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceLevelObjectiveRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceLevelObjectiveRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThresholds

`func (o *ServiceLevelObjectiveRequest) GetThresholds() []SLOThreshold`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *ServiceLevelObjectiveRequest) GetThresholdsOk() (*[]SLOThreshold, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *ServiceLevelObjectiveRequest) SetThresholds(v []SLOThreshold)`

SetThresholds sets Thresholds field to given value.


### GetType

`func (o *ServiceLevelObjectiveRequest) GetType() SLOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceLevelObjectiveRequest) GetTypeOk() (*SLOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceLevelObjectiveRequest) SetType(v SLOType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


