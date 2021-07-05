# UsageTraceResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Usage** | Pointer to [**[]UsageTraceHour**](UsageTraceHour.md) | Array with the number of hourly traces indexed for a given organization. | [optional] 

## Methods

### NewUsageTraceResponse

`func NewUsageTraceResponse() *UsageTraceResponse`

NewUsageTraceResponse instantiates a new UsageTraceResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageTraceResponseWithDefaults

`func NewUsageTraceResponseWithDefaults() *UsageTraceResponse`

NewUsageTraceResponseWithDefaults instantiates a new UsageTraceResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetUsage

`func (o *UsageTraceResponse) GetUsage() []UsageTraceHour`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *UsageTraceResponse) GetUsageOk() (*[]UsageTraceHour, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *UsageTraceResponse) SetUsage(v []UsageTraceHour)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *UsageTraceResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


