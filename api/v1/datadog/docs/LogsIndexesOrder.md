# LogsIndexesOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexNames** | Pointer to **[]string** | Array of strings identifying by their name(s) the index(es) of your organisation. Logs are tested against the query filter of each index one by one, following the order of the array. Logs are eventually stored in the first matching index. | 

## Methods

### GetIndexNames

`func (o *LogsIndexesOrder) GetIndexNames() []string`

GetIndexNames returns the IndexNames field if non-nil, zero value otherwise.

### GetIndexNamesOk

`func (o *LogsIndexesOrder) GetIndexNamesOk() ([]string, bool)`

GetIndexNamesOk returns a tuple with the IndexNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIndexNames

`func (o *LogsIndexesOrder) HasIndexNames() bool`

HasIndexNames returns a boolean if a field has been set.

### SetIndexNames

`func (o *LogsIndexesOrder) SetIndexNames(v []string)`

SetIndexNames gets a reference to the given []string and assigns it to the IndexNames field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


