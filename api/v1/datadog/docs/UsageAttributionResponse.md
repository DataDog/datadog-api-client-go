# UsageAttributionResponse

## Properties

| Name         | Type                                                                   | Description                  | Notes      |
| ------------ | ---------------------------------------------------------------------- | ---------------------------- | ---------- |
| **Metadata** | Pointer to [**UsageAttributionMetadata**](UsageAttributionMetadata.md) |                              | [optional] |
| **Usage**    | Pointer to [**[]UsageAttributionBody**](UsageAttributionBody.md)       | Get Usage Summary by tag(s). | [optional] |

## Methods

### NewUsageAttributionResponse

`func NewUsageAttributionResponse() *UsageAttributionResponse`

NewUsageAttributionResponse instantiates a new UsageAttributionResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageAttributionResponseWithDefaults

`func NewUsageAttributionResponseWithDefaults() *UsageAttributionResponse`

NewUsageAttributionResponseWithDefaults instantiates a new UsageAttributionResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetMetadata

`func (o *UsageAttributionResponse) GetMetadata() UsageAttributionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UsageAttributionResponse) GetMetadataOk() (*UsageAttributionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UsageAttributionResponse) SetMetadata(v UsageAttributionMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UsageAttributionResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUsage

`func (o *UsageAttributionResponse) GetUsage() []UsageAttributionBody`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *UsageAttributionResponse) GetUsageOk() (*[]UsageAttributionBody, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *UsageAttributionResponse) SetUsage(v []UsageAttributionBody)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *UsageAttributionResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
