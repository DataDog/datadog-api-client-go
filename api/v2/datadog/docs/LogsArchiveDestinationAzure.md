# LogsArchiveDestinationAzure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | **string** | The container where the archive will be stored. | 
**Integration** | [**LogsArchiveIntegrationAzure**](LogsArchiveIntegrationAzure.md) |  | 
**Path** | Pointer to **string** | The archive path. | [optional] 
**Region** | Pointer to **string** | The region where the archive will be stored. | [optional] 
**StorageAccount** | **string** | The associated storage account. | 
**Type** | [**LogsArchiveDestinationAzureType**](LogsArchiveDestinationAzureType.md) |  | [default to "azure"]

## Methods

### NewLogsArchiveDestinationAzure

`func NewLogsArchiveDestinationAzure(container string, integration LogsArchiveIntegrationAzure, storageAccount string, type_ LogsArchiveDestinationAzureType, ) *LogsArchiveDestinationAzure`

NewLogsArchiveDestinationAzure instantiates a new LogsArchiveDestinationAzure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsArchiveDestinationAzureWithDefaults

`func NewLogsArchiveDestinationAzureWithDefaults() *LogsArchiveDestinationAzure`

NewLogsArchiveDestinationAzureWithDefaults instantiates a new LogsArchiveDestinationAzure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *LogsArchiveDestinationAzure) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *LogsArchiveDestinationAzure) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *LogsArchiveDestinationAzure) SetContainer(v string)`

SetContainer sets Container field to given value.


### GetIntegration

`func (o *LogsArchiveDestinationAzure) GetIntegration() LogsArchiveIntegrationAzure`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *LogsArchiveDestinationAzure) GetIntegrationOk() (*LogsArchiveIntegrationAzure, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *LogsArchiveDestinationAzure) SetIntegration(v LogsArchiveIntegrationAzure)`

SetIntegration sets Integration field to given value.


### GetPath

`func (o *LogsArchiveDestinationAzure) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LogsArchiveDestinationAzure) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LogsArchiveDestinationAzure) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LogsArchiveDestinationAzure) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRegion

`func (o *LogsArchiveDestinationAzure) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LogsArchiveDestinationAzure) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LogsArchiveDestinationAzure) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LogsArchiveDestinationAzure) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStorageAccount

`func (o *LogsArchiveDestinationAzure) GetStorageAccount() string`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *LogsArchiveDestinationAzure) GetStorageAccountOk() (*string, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *LogsArchiveDestinationAzure) SetStorageAccount(v string)`

SetStorageAccount sets StorageAccount field to given value.


### GetType

`func (o *LogsArchiveDestinationAzure) GetType() LogsArchiveDestinationAzureType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsArchiveDestinationAzure) GetTypeOk() (*LogsArchiveDestinationAzureType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsArchiveDestinationAzure) SetType(v LogsArchiveDestinationAzureType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


