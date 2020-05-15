# LogsIndexListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indexes** | Pointer to [**[]LogsIndex**](LogsIndex.md) | Array of Log index configurations. | [optional] 

## Methods

### NewLogsIndexListResponse

`func NewLogsIndexListResponse() *LogsIndexListResponse`

NewLogsIndexListResponse instantiates a new LogsIndexListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsIndexListResponseWithDefaults

`func NewLogsIndexListResponseWithDefaults() *LogsIndexListResponse`

NewLogsIndexListResponseWithDefaults instantiates a new LogsIndexListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexes

`func (o *LogsIndexListResponse) GetIndexes() []LogsIndex`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *LogsIndexListResponse) GetIndexesOk() (*[]LogsIndex, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *LogsIndexListResponse) SetIndexes(v []LogsIndex)`

SetIndexes sets Indexes field to given value.

### HasIndexes

`func (o *LogsIndexListResponse) HasIndexes() bool`

HasIndexes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


