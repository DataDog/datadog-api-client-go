# LogsArchiveCreateRequestDestination

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

### NewLogsArchiveCreateRequestDestination

`func NewLogsArchiveCreateRequestDestination(container string, integration LogsArchiveIntegrationS3, storageAccount string, type_ LogsArchiveDestinationS3Type, bucket string, ) *LogsArchiveCreateRequestDestination`

NewLogsArchiveCreateRequestDestination instantiates a new LogsArchiveCreateRequestDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsArchiveCreateRequestDestinationWithDefaults

`func NewLogsArchiveCreateRequestDestinationWithDefaults() *LogsArchiveCreateRequestDestination`

NewLogsArchiveCreateRequestDestinationWithDefaults instantiates a new LogsArchiveCreateRequestDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *LogsArchiveCreateRequestDestination) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *LogsArchiveCreateRequestDestination) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *LogsArchiveCreateRequestDestination) SetContainer(v string)`

SetContainer sets Container field to given value.


### GetIntegration

`func (o *LogsArchiveCreateRequestDestination) GetIntegration() LogsArchiveIntegrationS3`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *LogsArchiveCreateRequestDestination) GetIntegrationOk() (*LogsArchiveIntegrationS3, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *LogsArchiveCreateRequestDestination) SetIntegration(v LogsArchiveIntegrationS3)`

SetIntegration sets Integration field to given value.


### GetPath

`func (o *LogsArchiveCreateRequestDestination) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LogsArchiveCreateRequestDestination) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LogsArchiveCreateRequestDestination) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LogsArchiveCreateRequestDestination) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRegion

`func (o *LogsArchiveCreateRequestDestination) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LogsArchiveCreateRequestDestination) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LogsArchiveCreateRequestDestination) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LogsArchiveCreateRequestDestination) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStorageAccount

`func (o *LogsArchiveCreateRequestDestination) GetStorageAccount() string`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *LogsArchiveCreateRequestDestination) GetStorageAccountOk() (*string, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *LogsArchiveCreateRequestDestination) SetStorageAccount(v string)`

SetStorageAccount sets StorageAccount field to given value.


### GetType

`func (o *LogsArchiveCreateRequestDestination) GetType() LogsArchiveDestinationS3Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsArchiveCreateRequestDestination) GetTypeOk() (*LogsArchiveDestinationS3Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsArchiveCreateRequestDestination) SetType(v LogsArchiveDestinationS3Type)`

SetType sets Type field to given value.


### GetBucket

`func (o *LogsArchiveCreateRequestDestination) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *LogsArchiveCreateRequestDestination) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *LogsArchiveCreateRequestDestination) SetBucket(v string)`

SetBucket sets Bucket field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


