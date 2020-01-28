# SyntheticsGetBrowserTestLatestResultsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTimestampFetched** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]SyntheticsBrowserTestResultShort**](SyntheticsBrowserTestResultShort.md) |  | [optional] 

## Methods

### GetLastTimestampFetched

`func (o *SyntheticsGetBrowserTestLatestResultsResponse) GetLastTimestampFetched() int64`

GetLastTimestampFetched returns the LastTimestampFetched field if non-nil, zero value otherwise.

### GetLastTimestampFetchedOk

`func (o *SyntheticsGetBrowserTestLatestResultsResponse) GetLastTimestampFetchedOk() (int64, bool)`

GetLastTimestampFetchedOk returns a tuple with the LastTimestampFetched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastTimestampFetched

`func (o *SyntheticsGetBrowserTestLatestResultsResponse) HasLastTimestampFetched() bool`

HasLastTimestampFetched returns a boolean if a field has been set.

### SetLastTimestampFetched

`func (o *SyntheticsGetBrowserTestLatestResultsResponse) SetLastTimestampFetched(v int64)`

SetLastTimestampFetched gets a reference to the given int64 and assigns it to the LastTimestampFetched field.

### GetResults

`func (o *SyntheticsGetBrowserTestLatestResultsResponse) GetResults() []SyntheticsBrowserTestResultShort`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SyntheticsGetBrowserTestLatestResultsResponse) GetResultsOk() ([]SyntheticsBrowserTestResultShort, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResults

`func (o *SyntheticsGetBrowserTestLatestResultsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResults

`func (o *SyntheticsGetBrowserTestLatestResultsResponse) SetResults(v []SyntheticsBrowserTestResultShort)`

SetResults gets a reference to the given []SyntheticsBrowserTestResultShort and assigns it to the Results field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


