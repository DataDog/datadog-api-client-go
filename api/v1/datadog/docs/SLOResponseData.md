# SLOResponseData

## Properties

| Name                   | Type                                                                       | Description                                                                                                                                                                                                                                                                                                                                                                                 | Notes                 |
| ---------------------- | -------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------- |
| **ConfiguredAlertIds** | Pointer to **[]int64**                                                     | A list of SLO monitors IDs that reference this SLO. This field is returned only when &#x60;with_configured_alert_ids&#x60; parameter is true in query.                                                                                                                                                                                                                                      | [optional]            |
| **CreatedAt**          | Pointer to **int64**                                                       | Creation timestamp (UNIX time in seconds) Always included in service level objective responses.                                                                                                                                                                                                                                                                                             | [optional] [readonly] |
| **Creator**            | Pointer to [**Creator**](Creator.md)                                       |                                                                                                                                                                                                                                                                                                                                                                                             | [optional]            |
| **Description**        | Pointer to **NullableString**                                              | A user-defined description of the service level objective. Always included in service level objective responses (but may be &#x60;null&#x60;). Optional in create/update requests.                                                                                                                                                                                                          | [optional]            |
| **Groups**             | Pointer to **[]string**                                                    | A list of (up to 20) monitor groups that narrow the scope of a monitor service level objective. Included in service level objective responses if it is not empty. Optional in create/update requests for monitor service level objectives, but may only be used when then length of the &#x60;monitor_ids&#x60; field is one.                                                               | [optional]            |
| **Id**                 | Pointer to **string**                                                      | A unique identifier for the service level objective object. Always included in service level objective responses.                                                                                                                                                                                                                                                                           | [optional] [readonly] |
| **ModifiedAt**         | Pointer to **int64**                                                       | Modification timestamp (UNIX time in seconds) Always included in service level objective responses.                                                                                                                                                                                                                                                                                         | [optional] [readonly] |
| **MonitorIds**         | Pointer to **[]int64**                                                     | A list of monitor ids that defines the scope of a monitor service level objective. **Required if type is &#x60;monitor&#x60;**.                                                                                                                                                                                                                                                             | [optional]            |
| **MonitorTags**        | Pointer to **[]string**                                                    | The union of monitor tags for all monitors referenced by the &#x60;monitor_ids&#x60; field. Always included in service level objective responses for monitor service level objectives (but may be empty). Ignored in create/update requests. Does not affect which monitors are included in the service level objective (that is determined entirely by the &#x60;monitor_ids&#x60; field). | [optional]            |
| **Name**               | Pointer to **string**                                                      | The name of the service level objective object.                                                                                                                                                                                                                                                                                                                                             | [optional]            |
| **Query**              | Pointer to [**ServiceLevelObjectiveQuery**](ServiceLevelObjectiveQuery.md) |                                                                                                                                                                                                                                                                                                                                                                                             | [optional]            |
| **Tags**               | Pointer to **[]string**                                                    | A list of tags associated with this service level objective. Always included in service level objective responses (but may be empty). Optional in create/update requests.                                                                                                                                                                                                                   | [optional]            |
| **Thresholds**         | Pointer to [**[]SLOThreshold**](SLOThreshold.md)                           | The thresholds (timeframes and associated targets) for this service level objective object.                                                                                                                                                                                                                                                                                                 | [optional]            |
| **Type**               | Pointer to [**SLOType**](SLOType.md)                                       |                                                                                                                                                                                                                                                                                                                                                                                             | [optional]            |

## Methods

### NewSLOResponseData

`func NewSLOResponseData() *SLOResponseData`

NewSLOResponseData instantiates a new SLOResponseData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSLOResponseDataWithDefaults

`func NewSLOResponseDataWithDefaults() *SLOResponseData`

NewSLOResponseDataWithDefaults instantiates a new SLOResponseData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetConfiguredAlertIds

`func (o *SLOResponseData) GetConfiguredAlertIds() []int64`

GetConfiguredAlertIds returns the ConfiguredAlertIds field if non-nil, zero value otherwise.

### GetConfiguredAlertIdsOk

`func (o *SLOResponseData) GetConfiguredAlertIdsOk() (*[]int64, bool)`

GetConfiguredAlertIdsOk returns a tuple with the ConfiguredAlertIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredAlertIds

`func (o *SLOResponseData) SetConfiguredAlertIds(v []int64)`

SetConfiguredAlertIds sets ConfiguredAlertIds field to given value.

### HasConfiguredAlertIds

`func (o *SLOResponseData) HasConfiguredAlertIds() bool`

HasConfiguredAlertIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SLOResponseData) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SLOResponseData) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SLOResponseData) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SLOResponseData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *SLOResponseData) GetCreator() Creator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *SLOResponseData) GetCreatorOk() (*Creator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *SLOResponseData) SetCreator(v Creator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *SLOResponseData) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDescription

`func (o *SLOResponseData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SLOResponseData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SLOResponseData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SLOResponseData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SLOResponseData) SetDescriptionNil(b bool)`

SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription

`func (o *SLOResponseData) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

### GetGroups

`func (o *SLOResponseData) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *SLOResponseData) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *SLOResponseData) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *SLOResponseData) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetId

`func (o *SLOResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SLOResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SLOResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SLOResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *SLOResponseData) GetModifiedAt() int64`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SLOResponseData) GetModifiedAtOk() (*int64, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SLOResponseData) SetModifiedAt(v int64)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *SLOResponseData) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetMonitorIds

`func (o *SLOResponseData) GetMonitorIds() []int64`

GetMonitorIds returns the MonitorIds field if non-nil, zero value otherwise.

### GetMonitorIdsOk

`func (o *SLOResponseData) GetMonitorIdsOk() (*[]int64, bool)`

GetMonitorIdsOk returns a tuple with the MonitorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorIds

`func (o *SLOResponseData) SetMonitorIds(v []int64)`

SetMonitorIds sets MonitorIds field to given value.

### HasMonitorIds

`func (o *SLOResponseData) HasMonitorIds() bool`

HasMonitorIds returns a boolean if a field has been set.

### GetMonitorTags

`func (o *SLOResponseData) GetMonitorTags() []string`

GetMonitorTags returns the MonitorTags field if non-nil, zero value otherwise.

### GetMonitorTagsOk

`func (o *SLOResponseData) GetMonitorTagsOk() (*[]string, bool)`

GetMonitorTagsOk returns a tuple with the MonitorTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTags

`func (o *SLOResponseData) SetMonitorTags(v []string)`

SetMonitorTags sets MonitorTags field to given value.

### HasMonitorTags

`func (o *SLOResponseData) HasMonitorTags() bool`

HasMonitorTags returns a boolean if a field has been set.

### GetName

`func (o *SLOResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SLOResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SLOResponseData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SLOResponseData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuery

`func (o *SLOResponseData) GetQuery() ServiceLevelObjectiveQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SLOResponseData) GetQueryOk() (*ServiceLevelObjectiveQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SLOResponseData) SetQuery(v ServiceLevelObjectiveQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SLOResponseData) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTags

`func (o *SLOResponseData) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SLOResponseData) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SLOResponseData) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SLOResponseData) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThresholds

`func (o *SLOResponseData) GetThresholds() []SLOThreshold`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *SLOResponseData) GetThresholdsOk() (*[]SLOThreshold, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *SLOResponseData) SetThresholds(v []SLOThreshold)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *SLOResponseData) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### GetType

`func (o *SLOResponseData) GetType() SLOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SLOResponseData) GetTypeOk() (*SLOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SLOResponseData) SetType(v SLOType)`

SetType sets Type field to given value.

### HasType

`func (o *SLOResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
