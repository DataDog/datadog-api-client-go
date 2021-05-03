# UsageTopAvgMetricsResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Metadata** | Pointer to [**UsageTopAvgMetricsMetadata**](UsageTopAvgMetricsMetadata.md) |  | [optional] 
**Usage** | Pointer to [**[]UsageTopAvgMetricsHour**](UsageTopAvgMetricsHour.md) | Number of hourly recorded custom metrics for a given organization. | [optional] 

## Methods

### NewUsageTopAvgMetricsResponse

`func NewUsageTopAvgMetricsResponse() *UsageTopAvgMetricsResponse`

NewUsageTopAvgMetricsResponse instantiates a new UsageTopAvgMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageTopAvgMetricsResponseWithDefaults

`func NewUsageTopAvgMetricsResponseWithDefaults() *UsageTopAvgMetricsResponse`

NewUsageTopAvgMetricsResponseWithDefaults instantiates a new UsageTopAvgMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *UsageTopAvgMetricsResponse) GetMetadata() UsageTopAvgMetricsMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UsageTopAvgMetricsResponse) GetMetadataOk() (*UsageTopAvgMetricsMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UsageTopAvgMetricsResponse) SetMetadata(v UsageTopAvgMetricsMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UsageTopAvgMetricsResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUsage

`func (o *UsageTopAvgMetricsResponse) GetUsage() []UsageTopAvgMetricsHour`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *UsageTopAvgMetricsResponse) GetUsageOk() (*[]UsageTopAvgMetricsHour, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *UsageTopAvgMetricsResponse) SetUsage(v []UsageTopAvgMetricsHour)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *UsageTopAvgMetricsResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


