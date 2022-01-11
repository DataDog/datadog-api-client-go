# SyntheticsGetAPITestLatestResultsResponse

## Properties

| Name                     | Type                                                                             | Description                           | Notes      |
| ------------------------ | -------------------------------------------------------------------------------- | ------------------------------------- | ---------- |
| **LastTimestampFetched** | Pointer to **int64**                                                             | Timestamp of the latest API test run. | [optional] |
| **Results**              | Pointer to [**[]SyntheticsAPITestResultShort**](SyntheticsAPITestResultShort.md) | Result of the latest API test run.    | [optional] |

## Methods

### NewSyntheticsGetAPITestLatestResultsResponse

`func NewSyntheticsGetAPITestLatestResultsResponse() *SyntheticsGetAPITestLatestResultsResponse`

NewSyntheticsGetAPITestLatestResultsResponse instantiates a new SyntheticsGetAPITestLatestResultsResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsGetAPITestLatestResultsResponseWithDefaults

`func NewSyntheticsGetAPITestLatestResultsResponseWithDefaults() *SyntheticsGetAPITestLatestResultsResponse`

NewSyntheticsGetAPITestLatestResultsResponseWithDefaults instantiates a new SyntheticsGetAPITestLatestResultsResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetLastTimestampFetched

`func (o *SyntheticsGetAPITestLatestResultsResponse) GetLastTimestampFetched() int64`

GetLastTimestampFetched returns the LastTimestampFetched field if non-nil, zero value otherwise.

### GetLastTimestampFetchedOk

`func (o *SyntheticsGetAPITestLatestResultsResponse) GetLastTimestampFetchedOk() (*int64, bool)`

GetLastTimestampFetchedOk returns a tuple with the LastTimestampFetched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimestampFetched

`func (o *SyntheticsGetAPITestLatestResultsResponse) SetLastTimestampFetched(v int64)`

SetLastTimestampFetched sets LastTimestampFetched field to given value.

### HasLastTimestampFetched

`func (o *SyntheticsGetAPITestLatestResultsResponse) HasLastTimestampFetched() bool`

HasLastTimestampFetched returns a boolean if a field has been set.

### GetResults

`func (o *SyntheticsGetAPITestLatestResultsResponse) GetResults() []SyntheticsAPITestResultShort`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SyntheticsGetAPITestLatestResultsResponse) GetResultsOk() (*[]SyntheticsAPITestResultShort, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SyntheticsGetAPITestLatestResultsResponse) SetResults(v []SyntheticsAPITestResultShort)`

SetResults sets Results field to given value.

### HasResults

`func (o *SyntheticsGetAPITestLatestResultsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
