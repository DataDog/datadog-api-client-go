# LogsArchiveCreateRequestAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | [**LogsArchiveCreateRequestDestination**](LogsArchiveCreateRequestDestination.md) |  | 
**IncludeTags** | Pointer to **bool** | To store the tags in the archive, set the value \&quot;true\&quot;. If it is set to \&quot;false\&quot;, the tags will be deleted when the logs are sent to the archive. | [optional] [default to false]
**Name** | **string** | The archive name. | 
**Query** | **string** | The archive query/filter. Logs matching this query are included in the archive. | 
**RehydrationTags** | Pointer to **[]string** | An array of tags to add to rehydrated logs from an archive. | [optional] 

## Methods

### NewLogsArchiveCreateRequestAttributes

`func NewLogsArchiveCreateRequestAttributes(destination LogsArchiveCreateRequestDestination, name string, query string, ) *LogsArchiveCreateRequestAttributes`

NewLogsArchiveCreateRequestAttributes instantiates a new LogsArchiveCreateRequestAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsArchiveCreateRequestAttributesWithDefaults

`func NewLogsArchiveCreateRequestAttributesWithDefaults() *LogsArchiveCreateRequestAttributes`

NewLogsArchiveCreateRequestAttributesWithDefaults instantiates a new LogsArchiveCreateRequestAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *LogsArchiveCreateRequestAttributes) GetDestination() LogsArchiveCreateRequestDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *LogsArchiveCreateRequestAttributes) GetDestinationOk() (*LogsArchiveCreateRequestDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *LogsArchiveCreateRequestAttributes) SetDestination(v LogsArchiveCreateRequestDestination)`

SetDestination sets Destination field to given value.


### GetIncludeTags

`func (o *LogsArchiveCreateRequestAttributes) GetIncludeTags() bool`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *LogsArchiveCreateRequestAttributes) GetIncludeTagsOk() (*bool, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *LogsArchiveCreateRequestAttributes) SetIncludeTags(v bool)`

SetIncludeTags sets IncludeTags field to given value.

### HasIncludeTags

`func (o *LogsArchiveCreateRequestAttributes) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### GetName

`func (o *LogsArchiveCreateRequestAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsArchiveCreateRequestAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsArchiveCreateRequestAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetQuery

`func (o *LogsArchiveCreateRequestAttributes) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *LogsArchiveCreateRequestAttributes) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *LogsArchiveCreateRequestAttributes) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetRehydrationTags

`func (o *LogsArchiveCreateRequestAttributes) GetRehydrationTags() []string`

GetRehydrationTags returns the RehydrationTags field if non-nil, zero value otherwise.

### GetRehydrationTagsOk

`func (o *LogsArchiveCreateRequestAttributes) GetRehydrationTagsOk() (*[]string, bool)`

GetRehydrationTagsOk returns a tuple with the RehydrationTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRehydrationTags

`func (o *LogsArchiveCreateRequestAttributes) SetRehydrationTags(v []string)`

SetRehydrationTags sets RehydrationTags field to given value.

### HasRehydrationTags

`func (o *LogsArchiveCreateRequestAttributes) HasRehydrationTags() bool`

HasRehydrationTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


