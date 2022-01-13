# LogsArchiveAttributes

## Properties

| Name                | Type                                                            | Description                                                                                                                                                              | Notes                         |
| ------------------- | --------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ----------------------------- |
| **Destination**     | [**NullableLogsArchiveDestination**](LogsArchiveDestination.md) |                                                                                                                                                                          |
| **IncludeTags**     | Pointer to **bool**                                             | To store the tags in the archive, set the value \&quot;true\&quot;. If it is set to \&quot;false\&quot;, the tags will be deleted when the logs are sent to the archive. | [optional] [default to false] |
| **Name**            | **string**                                                      | The archive name.                                                                                                                                                        |
| **Query**           | **string**                                                      | The archive query/filter. Logs matching this query are included in the archive.                                                                                          |
| **RehydrationTags** | Pointer to **[]string**                                         | An array of tags to add to rehydrated logs from an archive.                                                                                                              | [optional]                    |
| **State**           | Pointer to [**LogsArchiveState**](LogsArchiveState.md)          |                                                                                                                                                                          | [optional]                    |

## Methods

### NewLogsArchiveAttributes

`func NewLogsArchiveAttributes(destination NullableLogsArchiveDestination, name string, query string) *LogsArchiveAttributes`

NewLogsArchiveAttributes instantiates a new LogsArchiveAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsArchiveAttributesWithDefaults

`func NewLogsArchiveAttributesWithDefaults() *LogsArchiveAttributes`

NewLogsArchiveAttributesWithDefaults instantiates a new LogsArchiveAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDestination

`func (o *LogsArchiveAttributes) GetDestination() LogsArchiveDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *LogsArchiveAttributes) GetDestinationOk() (*LogsArchiveDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *LogsArchiveAttributes) SetDestination(v LogsArchiveDestination)`

SetDestination sets Destination field to given value.

### SetDestinationNil

`func (o *LogsArchiveAttributes) SetDestinationNil(b bool)`

SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination

`func (o *LogsArchiveAttributes) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil

### GetIncludeTags

`func (o *LogsArchiveAttributes) GetIncludeTags() bool`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *LogsArchiveAttributes) GetIncludeTagsOk() (*bool, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *LogsArchiveAttributes) SetIncludeTags(v bool)`

SetIncludeTags sets IncludeTags field to given value.

### HasIncludeTags

`func (o *LogsArchiveAttributes) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### GetName

`func (o *LogsArchiveAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsArchiveAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsArchiveAttributes) SetName(v string)`

SetName sets Name field to given value.

### GetQuery

`func (o *LogsArchiveAttributes) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *LogsArchiveAttributes) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *LogsArchiveAttributes) SetQuery(v string)`

SetQuery sets Query field to given value.

### GetRehydrationTags

`func (o *LogsArchiveAttributes) GetRehydrationTags() []string`

GetRehydrationTags returns the RehydrationTags field if non-nil, zero value otherwise.

### GetRehydrationTagsOk

`func (o *LogsArchiveAttributes) GetRehydrationTagsOk() (*[]string, bool)`

GetRehydrationTagsOk returns a tuple with the RehydrationTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRehydrationTags

`func (o *LogsArchiveAttributes) SetRehydrationTags(v []string)`

SetRehydrationTags sets RehydrationTags field to given value.

### HasRehydrationTags

`func (o *LogsArchiveAttributes) HasRehydrationTags() bool`

HasRehydrationTags returns a boolean if a field has been set.

### GetState

`func (o *LogsArchiveAttributes) GetState() LogsArchiveState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LogsArchiveAttributes) GetStateOk() (*LogsArchiveState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LogsArchiveAttributes) SetState(v LogsArchiveState)`

SetState sets State field to given value.

### HasState

`func (o *LogsArchiveAttributes) HasState() bool`

HasState returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
