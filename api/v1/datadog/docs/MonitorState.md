# MonitorState

## Properties

| Name       | Type                                                                | Description                                                                                                                                | Notes      |
| ---------- | ------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------ | ---------- |
| **Groups** | Pointer to [**map[string]MonitorStateGroup**](MonitorStateGroup.md) | Dictionary where the keys are groups (comma separated lists of tags) and the values are the list of groups your monitor is broken down on. | [optional] |

## Methods

### NewMonitorState

`func NewMonitorState() *MonitorState`

NewMonitorState instantiates a new MonitorState object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMonitorStateWithDefaults

`func NewMonitorStateWithDefaults() *MonitorState`

NewMonitorStateWithDefaults instantiates a new MonitorState object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetGroups

`func (o *MonitorState) GetGroups() map[string]MonitorStateGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *MonitorState) GetGroupsOk() (*map[string]MonitorStateGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *MonitorState) SetGroups(v map[string]MonitorStateGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *MonitorState) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
