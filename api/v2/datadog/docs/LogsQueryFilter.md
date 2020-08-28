# LogsQueryFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | The minimum time for the requested logs, supports date math and regular timestamps | [optional] [default to "now-15m"]
**Indexes** | Pointer to **[]string** | For customers with multiple indexes, the indexes to search. Defaults to [&#39;*&#39;] which means all indexes. | [optional] [default to ["*"]]
**Query** | Pointer to **string** | The search query - following the log search syntax. | [optional] [default to "*"]
**To** | Pointer to **string** | The maximum time for the requested logs, supports date math and regular timestamps | [optional] [default to "now"]

## Methods

### NewLogsQueryFilter

`func NewLogsQueryFilter() *LogsQueryFilter`

NewLogsQueryFilter instantiates a new LogsQueryFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsQueryFilterWithDefaults

`func NewLogsQueryFilterWithDefaults() *LogsQueryFilter`

NewLogsQueryFilterWithDefaults instantiates a new LogsQueryFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *LogsQueryFilter) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *LogsQueryFilter) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *LogsQueryFilter) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *LogsQueryFilter) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetIndexes

`func (o *LogsQueryFilter) GetIndexes() []string`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *LogsQueryFilter) GetIndexesOk() (*[]string, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *LogsQueryFilter) SetIndexes(v []string)`

SetIndexes sets Indexes field to given value.

### HasIndexes

`func (o *LogsQueryFilter) HasIndexes() bool`

HasIndexes returns a boolean if a field has been set.

### GetQuery

`func (o *LogsQueryFilter) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *LogsQueryFilter) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *LogsQueryFilter) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *LogsQueryFilter) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTo

`func (o *LogsQueryFilter) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *LogsQueryFilter) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *LogsQueryFilter) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *LogsQueryFilter) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


