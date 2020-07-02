# LogsListRequestFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to [**time.Time**](time.Time.md) | Minimum timestamp for requested logs. | [optional] 
**Indexes** | Pointer to **[]string** | For customers with multiple indexes, the indexes to search. Defaults to &#39;*&#39; which means all indexes. | [optional] 
**Query** | Pointer to **string** | Search query following logs syntax. | [optional] 
**To** | Pointer to [**time.Time**](time.Time.md) | Maximum timestamp for requested logs. | [optional] 

## Methods

### NewLogsListRequestFilter

`func NewLogsListRequestFilter() *LogsListRequestFilter`

NewLogsListRequestFilter instantiates a new LogsListRequestFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsListRequestFilterWithDefaults

`func NewLogsListRequestFilterWithDefaults() *LogsListRequestFilter`

NewLogsListRequestFilterWithDefaults instantiates a new LogsListRequestFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *LogsListRequestFilter) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *LogsListRequestFilter) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *LogsListRequestFilter) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *LogsListRequestFilter) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetIndexes

`func (o *LogsListRequestFilter) GetIndexes() []string`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *LogsListRequestFilter) GetIndexesOk() (*[]string, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *LogsListRequestFilter) SetIndexes(v []string)`

SetIndexes sets Indexes field to given value.

### HasIndexes

`func (o *LogsListRequestFilter) HasIndexes() bool`

HasIndexes returns a boolean if a field has been set.

### GetQuery

`func (o *LogsListRequestFilter) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *LogsListRequestFilter) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *LogsListRequestFilter) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *LogsListRequestFilter) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTo

`func (o *LogsListRequestFilter) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *LogsListRequestFilter) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *LogsListRequestFilter) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *LogsListRequestFilter) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


