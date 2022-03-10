# AuditLogsQueryOptions

## Properties

| Name           | Type                  | Description                                                                         | Notes                         |
| -------------- | --------------------- | ----------------------------------------------------------------------------------- | ----------------------------- |
| **TimeOffset** | Pointer to **int64**  | Time offset (in seconds) to apply to the query.                                     | [optional]                    |
| **Timezone**   | Pointer to **string** | Timezone code. Can be specified as an offset, for example: \&quot;UTC+03:00\&quot;. | [optional] [default to "UTC"] |

## Methods

### NewAuditLogsQueryOptions

`func NewAuditLogsQueryOptions() *AuditLogsQueryOptions`

NewAuditLogsQueryOptions instantiates a new AuditLogsQueryOptions object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAuditLogsQueryOptionsWithDefaults

`func NewAuditLogsQueryOptionsWithDefaults() *AuditLogsQueryOptions`

NewAuditLogsQueryOptionsWithDefaults instantiates a new AuditLogsQueryOptions object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetTimeOffset

`func (o *AuditLogsQueryOptions) GetTimeOffset() int64`

GetTimeOffset returns the TimeOffset field if non-nil, zero value otherwise.

### GetTimeOffsetOk

`func (o *AuditLogsQueryOptions) GetTimeOffsetOk() (*int64, bool)`

GetTimeOffsetOk returns a tuple with the TimeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffset

`func (o *AuditLogsQueryOptions) SetTimeOffset(v int64)`

SetTimeOffset sets TimeOffset field to given value.

### HasTimeOffset

`func (o *AuditLogsQueryOptions) HasTimeOffset() bool`

HasTimeOffset returns a boolean if a field has been set.

### GetTimezone

`func (o *AuditLogsQueryOptions) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AuditLogsQueryOptions) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AuditLogsQueryOptions) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *AuditLogsQueryOptions) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
