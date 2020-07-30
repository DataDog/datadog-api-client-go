# LogsArchiveDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | **string** | The container where the archive will be stored. | 
**Integration** | [**LogsArchiveIntegrationS3**](LogsArchiveIntegrationS3.md) |  | 
**Path** | Pointer to **string** | The archive path. | [optional] 
**Region** | Pointer to **string** | The region where the archive will be stored. | [optional] 
**StorageAccount** | **string** | The associated storage account. | 
**Type** | [**LogsArchiveDestinationS3Type**](LogsArchiveDestinationS3Type.md) |  | [default to "s3"]
**Bucket** | **string** | The bucket where the archive will be stored. | 

## Methods

### NewLogsArchiveDestination

`func NewLogsArchiveDestination(container string, integration LogsArchiveIntegrationS3, storageAccount string, type_ LogsArchiveDestinationS3Type, bucket string, ) *LogsArchiveDestination`

NewLogsArchiveDestination instantiates a new LogsArchiveDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsArchiveDestinationWithDefaults

`func NewLogsArchiveDestinationWithDefaults() *LogsArchiveDestination`

NewLogsArchiveDestinationWithDefaults instantiates a new LogsArchiveDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *LogsArchiveDestination) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *LogsArchiveDestination) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *LogsArchiveDestination) SetContainer(v string)`

SetContainer sets Container field to given value.


### GetIntegration

`func (o *LogsArchiveDestination) GetIntegration() LogsArchiveIntegrationS3`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *LogsArchiveDestination) GetIntegrationOk() (*LogsArchiveIntegrationS3, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *LogsArchiveDestination) SetIntegration(v LogsArchiveIntegrationS3)`

SetIntegration sets Integration field to given value.


### GetPath

`func (o *LogsArchiveDestination) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LogsArchiveDestination) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LogsArchiveDestination) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LogsArchiveDestination) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRegion

`func (o *LogsArchiveDestination) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LogsArchiveDestination) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LogsArchiveDestination) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LogsArchiveDestination) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStorageAccount

`func (o *LogsArchiveDestination) GetStorageAccount() string`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *LogsArchiveDestination) GetStorageAccountOk() (*string, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *LogsArchiveDestination) SetStorageAccount(v string)`

SetStorageAccount sets StorageAccount field to given value.


### GetType

`func (o *LogsArchiveDestination) GetType() LogsArchiveDestinationS3Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsArchiveDestination) GetTypeOk() (*LogsArchiveDestinationS3Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsArchiveDestination) SetType(v LogsArchiveDestinationS3Type)`

SetType sets Type field to given value.


### GetBucket

`func (o *LogsArchiveDestination) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *LogsArchiveDestination) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *LogsArchiveDestination) SetBucket(v string)`

SetBucket sets Bucket field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


