# SyntheticsBatchDetailsData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Metadata** | Pointer to [**SyntheticsCIBatchMetadata**](SyntheticsCIBatchMetadata.md) |  | [optional] 
**Results** | Pointer to [**[]SyntheticsBatchResult**](SyntheticsBatchResult.md) | List of results for the batch. | [optional] 
**Status** | Pointer to [**SyntheticsStatus**](SyntheticsStatus.md) |  | [optional] 

## Methods

### NewSyntheticsBatchDetailsData

`func NewSyntheticsBatchDetailsData() *SyntheticsBatchDetailsData`

NewSyntheticsBatchDetailsData instantiates a new SyntheticsBatchDetailsData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsBatchDetailsDataWithDefaults

`func NewSyntheticsBatchDetailsDataWithDefaults() *SyntheticsBatchDetailsData`

NewSyntheticsBatchDetailsDataWithDefaults instantiates a new SyntheticsBatchDetailsData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetMetadata

`func (o *SyntheticsBatchDetailsData) GetMetadata() SyntheticsCIBatchMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SyntheticsBatchDetailsData) GetMetadataOk() (*SyntheticsCIBatchMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SyntheticsBatchDetailsData) SetMetadata(v SyntheticsCIBatchMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SyntheticsBatchDetailsData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResults

`func (o *SyntheticsBatchDetailsData) GetResults() []SyntheticsBatchResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SyntheticsBatchDetailsData) GetResultsOk() (*[]SyntheticsBatchResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SyntheticsBatchDetailsData) SetResults(v []SyntheticsBatchResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *SyntheticsBatchDetailsData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetStatus

`func (o *SyntheticsBatchDetailsData) GetStatus() SyntheticsStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyntheticsBatchDetailsData) GetStatusOk() (*SyntheticsStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyntheticsBatchDetailsData) SetStatus(v SyntheticsStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyntheticsBatchDetailsData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


