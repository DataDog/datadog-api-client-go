# LogsIndexesOrder

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**IndexNames** | **[]string** | Array of strings identifying by their name(s) the index(es) of your organization. Logs are tested against the query filter of each index one by one, following the order of the array. Logs are eventually stored in the first matching index. | 

## Methods

### NewLogsIndexesOrder

`func NewLogsIndexesOrder(indexNames []string, ) *LogsIndexesOrder`

NewLogsIndexesOrder instantiates a new LogsIndexesOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsIndexesOrderWithDefaults

`func NewLogsIndexesOrderWithDefaults() *LogsIndexesOrder`

NewLogsIndexesOrderWithDefaults instantiates a new LogsIndexesOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexNames

`func (o *LogsIndexesOrder) GetIndexNames() []string`

GetIndexNames returns the IndexNames field if non-nil, zero value otherwise.

### GetIndexNamesOk

`func (o *LogsIndexesOrder) GetIndexNamesOk() (*[]string, bool)`

GetIndexNamesOk returns a tuple with the IndexNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNames

`func (o *LogsIndexesOrder) SetIndexNames(v []string)`

SetIndexNames sets IndexNames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


