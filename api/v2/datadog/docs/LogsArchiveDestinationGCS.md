# LogsArchiveDestinationGCS

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Bucket** | **string** | The bucket where the archive will be stored. | 
**Integration** | [**LogsArchiveIntegrationGCS**](LogsArchiveIntegrationGCS.md) |  | 
**Path** | Pointer to **string** | The archive path. | [optional] 
**Type** | [**LogsArchiveDestinationGCSType**](LogsArchiveDestinationGCSType.md) |  | [default to LOGSARCHIVEDESTINATIONGCSTYPE_GCS]

## Methods

### NewLogsArchiveDestinationGCS

`func NewLogsArchiveDestinationGCS(bucket string, integration LogsArchiveIntegrationGCS, type_ LogsArchiveDestinationGCSType) *LogsArchiveDestinationGCS`

NewLogsArchiveDestinationGCS instantiates a new LogsArchiveDestinationGCS object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsArchiveDestinationGCSWithDefaults

`func NewLogsArchiveDestinationGCSWithDefaults() *LogsArchiveDestinationGCS`

NewLogsArchiveDestinationGCSWithDefaults instantiates a new LogsArchiveDestinationGCS object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetBucket

`func (o *LogsArchiveDestinationGCS) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *LogsArchiveDestinationGCS) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *LogsArchiveDestinationGCS) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetIntegration

`func (o *LogsArchiveDestinationGCS) GetIntegration() LogsArchiveIntegrationGCS`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *LogsArchiveDestinationGCS) GetIntegrationOk() (*LogsArchiveIntegrationGCS, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *LogsArchiveDestinationGCS) SetIntegration(v LogsArchiveIntegrationGCS)`

SetIntegration sets Integration field to given value.


### GetPath

`func (o *LogsArchiveDestinationGCS) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LogsArchiveDestinationGCS) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LogsArchiveDestinationGCS) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LogsArchiveDestinationGCS) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetType

`func (o *LogsArchiveDestinationGCS) GetType() LogsArchiveDestinationGCSType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsArchiveDestinationGCS) GetTypeOk() (*LogsArchiveDestinationGCSType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsArchiveDestinationGCS) SetType(v LogsArchiveDestinationGCSType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


