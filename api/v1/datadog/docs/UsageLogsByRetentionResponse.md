# UsageLogsByRetentionResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Usage** | Pointer to [**[]UsageLogsByRetentionHour**](UsageLogsByRetentionHour.md) | Get hourly usage for indexed logs by retention period. | [optional] 

## Methods

### NewUsageLogsByRetentionResponse

`func NewUsageLogsByRetentionResponse() *UsageLogsByRetentionResponse`

NewUsageLogsByRetentionResponse instantiates a new UsageLogsByRetentionResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageLogsByRetentionResponseWithDefaults

`func NewUsageLogsByRetentionResponseWithDefaults() *UsageLogsByRetentionResponse`

NewUsageLogsByRetentionResponseWithDefaults instantiates a new UsageLogsByRetentionResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetUsage

`func (o *UsageLogsByRetentionResponse) GetUsage() []UsageLogsByRetentionHour`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *UsageLogsByRetentionResponse) GetUsageOk() (*[]UsageLogsByRetentionHour, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *UsageLogsByRetentionResponse) SetUsage(v []UsageLogsByRetentionHour)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *UsageLogsByRetentionResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


