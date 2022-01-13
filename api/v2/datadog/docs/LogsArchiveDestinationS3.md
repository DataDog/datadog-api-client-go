# LogsArchiveDestinationS3

## Properties

| Name            | Type                                                                | Description                                  | Notes                                        |
| --------------- | ------------------------------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| **Bucket**      | **string**                                                          | The bucket where the archive will be stored. |
| **Integration** | [**LogsArchiveIntegrationS3**](LogsArchiveIntegrationS3.md)         |                                              |
| **Path**        | Pointer to **string**                                               | The archive path.                            | [optional]                                   |
| **Type**        | [**LogsArchiveDestinationS3Type**](LogsArchiveDestinationS3Type.md) |                                              | [default to LOGSARCHIVEDESTINATIONS3TYPE_S3] |

## Methods

### NewLogsArchiveDestinationS3

`func NewLogsArchiveDestinationS3(bucket string, integration LogsArchiveIntegrationS3, type_ LogsArchiveDestinationS3Type) *LogsArchiveDestinationS3`

NewLogsArchiveDestinationS3 instantiates a new LogsArchiveDestinationS3 object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsArchiveDestinationS3WithDefaults

`func NewLogsArchiveDestinationS3WithDefaults() *LogsArchiveDestinationS3`

NewLogsArchiveDestinationS3WithDefaults instantiates a new LogsArchiveDestinationS3 object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetBucket

`func (o *LogsArchiveDestinationS3) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *LogsArchiveDestinationS3) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *LogsArchiveDestinationS3) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### GetIntegration

`func (o *LogsArchiveDestinationS3) GetIntegration() LogsArchiveIntegrationS3`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *LogsArchiveDestinationS3) GetIntegrationOk() (*LogsArchiveIntegrationS3, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *LogsArchiveDestinationS3) SetIntegration(v LogsArchiveIntegrationS3)`

SetIntegration sets Integration field to given value.

### GetPath

`func (o *LogsArchiveDestinationS3) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LogsArchiveDestinationS3) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LogsArchiveDestinationS3) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LogsArchiveDestinationS3) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetType

`func (o *LogsArchiveDestinationS3) GetType() LogsArchiveDestinationS3Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsArchiveDestinationS3) GetTypeOk() (*LogsArchiveDestinationS3Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsArchiveDestinationS3) SetType(v LogsArchiveDestinationS3Type)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
