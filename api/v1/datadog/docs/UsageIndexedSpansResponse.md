# UsageIndexedSpansResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Usage** | Pointer to [**[]UsageIndexedSpansHour**](UsageIndexedSpansHour.md) | Array with the number of hourly traces indexed for a given organization. | [optional] 

## Methods

### NewUsageIndexedSpansResponse

`func NewUsageIndexedSpansResponse() *UsageIndexedSpansResponse`

NewUsageIndexedSpansResponse instantiates a new UsageIndexedSpansResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageIndexedSpansResponseWithDefaults

`func NewUsageIndexedSpansResponseWithDefaults() *UsageIndexedSpansResponse`

NewUsageIndexedSpansResponseWithDefaults instantiates a new UsageIndexedSpansResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetUsage

`func (o *UsageIndexedSpansResponse) GetUsage() []UsageIndexedSpansHour`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *UsageIndexedSpansResponse) GetUsageOk() (*[]UsageIndexedSpansHour, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *UsageIndexedSpansResponse) SetUsage(v []UsageIndexedSpansHour)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *UsageIndexedSpansResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


