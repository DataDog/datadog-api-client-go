# SyntheticsTriggerTest

## Properties

| Name         | Type                                                                     | Description                                      | Notes      |
| ------------ | ------------------------------------------------------------------------ | ------------------------------------------------ | ---------- |
| **Metadata** | Pointer to [**SyntheticsCIBatchMetadata**](SyntheticsCIBatchMetadata.md) |                                                  | [optional] |
| **PublicId** | **string**                                                               | The public ID of the Synthetics test to trigger. |

## Methods

### NewSyntheticsTriggerTest

`func NewSyntheticsTriggerTest(publicId string) *SyntheticsTriggerTest`

NewSyntheticsTriggerTest instantiates a new SyntheticsTriggerTest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsTriggerTestWithDefaults

`func NewSyntheticsTriggerTestWithDefaults() *SyntheticsTriggerTest`

NewSyntheticsTriggerTestWithDefaults instantiates a new SyntheticsTriggerTest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetMetadata

`func (o *SyntheticsTriggerTest) GetMetadata() SyntheticsCIBatchMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SyntheticsTriggerTest) GetMetadataOk() (*SyntheticsCIBatchMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SyntheticsTriggerTest) SetMetadata(v SyntheticsCIBatchMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SyntheticsTriggerTest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPublicId

`func (o *SyntheticsTriggerTest) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *SyntheticsTriggerTest) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *SyntheticsTriggerTest) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
