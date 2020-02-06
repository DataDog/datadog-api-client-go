# SyntheticsGetApiTestLatestResultsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTimestampFetched** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]SyntheticsApiTestResultShort**](SyntheticsAPITestResultShort.md) |  | [optional] 

## Methods

### NewSyntheticsGetApiTestLatestResultsResponse

`func NewSyntheticsGetApiTestLatestResultsResponse() *SyntheticsGetApiTestLatestResultsResponse`

NewSyntheticsGetApiTestLatestResultsResponse instantiates a new SyntheticsGetApiTestLatestResultsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsGetApiTestLatestResultsResponseWithDefaults

`func NewSyntheticsGetApiTestLatestResultsResponseWithDefaults() *SyntheticsGetApiTestLatestResultsResponse`

NewSyntheticsGetApiTestLatestResultsResponseWithDefaults instantiates a new SyntheticsGetApiTestLatestResultsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTimestampFetched

`func (o *SyntheticsGetApiTestLatestResultsResponse) GetLastTimestampFetched() int64`

GetLastTimestampFetched returns the LastTimestampFetched field if non-nil, zero value otherwise.

### GetLastTimestampFetchedOk

`func (o *SyntheticsGetApiTestLatestResultsResponse) GetLastTimestampFetchedOk() (int64, bool)`

GetLastTimestampFetchedOk returns a tuple with the LastTimestampFetched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastTimestampFetched

`func (o *SyntheticsGetApiTestLatestResultsResponse) HasLastTimestampFetched() bool`

HasLastTimestampFetched returns a boolean if a field has been set.

### SetLastTimestampFetched

`func (o *SyntheticsGetApiTestLatestResultsResponse) SetLastTimestampFetched(v int64)`

SetLastTimestampFetched gets a reference to the given int64 and assigns it to the LastTimestampFetched field.

### GetResults

`func (o *SyntheticsGetApiTestLatestResultsResponse) GetResults() []SyntheticsApiTestResultShort`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SyntheticsGetApiTestLatestResultsResponse) GetResultsOk() ([]SyntheticsApiTestResultShort, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResults

`func (o *SyntheticsGetApiTestLatestResultsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResults

`func (o *SyntheticsGetApiTestLatestResultsResponse) SetResults(v []SyntheticsApiTestResultShort)`

SetResults gets a reference to the given []SyntheticsApiTestResultShort and assigns it to the Results field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


