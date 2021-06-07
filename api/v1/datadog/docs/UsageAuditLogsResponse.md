# UsageAuditLogsResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Usage** | Pointer to [**[]UsageAuditLogsHour**](UsageAuditLogsHour.md) | Get hourly usage for audit logs. | [optional] 

## Methods

### NewUsageAuditLogsResponse

`func NewUsageAuditLogsResponse() *UsageAuditLogsResponse`

NewUsageAuditLogsResponse instantiates a new UsageAuditLogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageAuditLogsResponseWithDefaults

`func NewUsageAuditLogsResponseWithDefaults() *UsageAuditLogsResponse`

NewUsageAuditLogsResponseWithDefaults instantiates a new UsageAuditLogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsage

`func (o *UsageAuditLogsResponse) GetUsage() []UsageAuditLogsHour`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *UsageAuditLogsResponse) GetUsageOk() (*[]UsageAuditLogsHour, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *UsageAuditLogsResponse) SetUsage(v []UsageAuditLogsHour)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *UsageAuditLogsResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


