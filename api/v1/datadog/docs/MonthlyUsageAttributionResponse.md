# MonthlyUsageAttributionResponse

## Properties

| Name         | Type                                                                                 | Description                  | Notes      |
| ------------ | ------------------------------------------------------------------------------------ | ---------------------------- | ---------- |
| **Metadata** | Pointer to [**MonthlyUsageAttributionMetadata**](MonthlyUsageAttributionMetadata.md) |                              | [optional] |
| **Usage**    | Pointer to [**[]MonthlyUsageAttributionBody**](MonthlyUsageAttributionBody.md)       | Get Usage Summary by tag(s). | [optional] |

## Methods

### NewMonthlyUsageAttributionResponse

`func NewMonthlyUsageAttributionResponse() *MonthlyUsageAttributionResponse`

NewMonthlyUsageAttributionResponse instantiates a new MonthlyUsageAttributionResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMonthlyUsageAttributionResponseWithDefaults

`func NewMonthlyUsageAttributionResponseWithDefaults() *MonthlyUsageAttributionResponse`

NewMonthlyUsageAttributionResponseWithDefaults instantiates a new MonthlyUsageAttributionResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetMetadata

`func (o *MonthlyUsageAttributionResponse) GetMetadata() MonthlyUsageAttributionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MonthlyUsageAttributionResponse) GetMetadataOk() (*MonthlyUsageAttributionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MonthlyUsageAttributionResponse) SetMetadata(v MonthlyUsageAttributionMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MonthlyUsageAttributionResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUsage

`func (o *MonthlyUsageAttributionResponse) GetUsage() []MonthlyUsageAttributionBody`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *MonthlyUsageAttributionResponse) GetUsageOk() (*[]MonthlyUsageAttributionBody, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *MonthlyUsageAttributionResponse) SetUsage(v []MonthlyUsageAttributionBody)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *MonthlyUsageAttributionResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
