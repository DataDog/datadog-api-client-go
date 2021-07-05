# LogsArchiveOrderAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**ArchiveIds** | **[]string** | An ordered array of &#x60;&lt;ARCHIVE_ID&gt;&#x60; strings, the order of archive IDs in the array define the overall archives order for Datadog. | 

## Methods

### NewLogsArchiveOrderAttributes

`func NewLogsArchiveOrderAttributes(archiveIds []string) *LogsArchiveOrderAttributes`

NewLogsArchiveOrderAttributes instantiates a new LogsArchiveOrderAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsArchiveOrderAttributesWithDefaults

`func NewLogsArchiveOrderAttributesWithDefaults() *LogsArchiveOrderAttributes`

NewLogsArchiveOrderAttributesWithDefaults instantiates a new LogsArchiveOrderAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetArchiveIds

`func (o *LogsArchiveOrderAttributes) GetArchiveIds() []string`

GetArchiveIds returns the ArchiveIds field if non-nil, zero value otherwise.

### GetArchiveIdsOk

`func (o *LogsArchiveOrderAttributes) GetArchiveIdsOk() (*[]string, bool)`

GetArchiveIdsOk returns a tuple with the ArchiveIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveIds

`func (o *LogsArchiveOrderAttributes) SetArchiveIds(v []string)`

SetArchiveIds sets ArchiveIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


