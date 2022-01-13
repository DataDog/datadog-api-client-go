# HourlyUsageAttributionResponse

## Properties

| Name         | Type                                                                               | Description                                 | Notes      |
| ------------ | ---------------------------------------------------------------------------------- | ------------------------------------------- | ---------- |
| **Metadata** | Pointer to [**HourlyUsageAttributionMetadata**](HourlyUsageAttributionMetadata.md) |                                             | [optional] |
| **Usage**    | Pointer to [**[]HourlyUsageAttributionBody**](HourlyUsageAttributionBody.md)       | Get the hourly usage attribution by tag(s). | [optional] |

## Methods

### NewHourlyUsageAttributionResponse

`func NewHourlyUsageAttributionResponse() *HourlyUsageAttributionResponse`

NewHourlyUsageAttributionResponse instantiates a new HourlyUsageAttributionResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewHourlyUsageAttributionResponseWithDefaults

`func NewHourlyUsageAttributionResponseWithDefaults() *HourlyUsageAttributionResponse`

NewHourlyUsageAttributionResponseWithDefaults instantiates a new HourlyUsageAttributionResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetMetadata

`func (o *HourlyUsageAttributionResponse) GetMetadata() HourlyUsageAttributionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HourlyUsageAttributionResponse) GetMetadataOk() (*HourlyUsageAttributionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HourlyUsageAttributionResponse) SetMetadata(v HourlyUsageAttributionMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HourlyUsageAttributionResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUsage

`func (o *HourlyUsageAttributionResponse) GetUsage() []HourlyUsageAttributionBody`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *HourlyUsageAttributionResponse) GetUsageOk() (*[]HourlyUsageAttributionBody, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *HourlyUsageAttributionResponse) SetUsage(v []HourlyUsageAttributionBody)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *HourlyUsageAttributionResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
