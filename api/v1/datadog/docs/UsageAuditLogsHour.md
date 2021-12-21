# UsageAuditLogsHour

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Hour** | Pointer to **time.Time** | The hour for the usage. | [optional] 
**LinesIndexed** | Pointer to **int64** | The total number of audit logs lines indexed during a given hour. | [optional] 
**OrgName** | Pointer to **string** | The organization name. | [optional] 
**PublicId** | Pointer to **string** | The organization public ID. | [optional] 

## Methods

### NewUsageAuditLogsHour

`func NewUsageAuditLogsHour() *UsageAuditLogsHour`

NewUsageAuditLogsHour instantiates a new UsageAuditLogsHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageAuditLogsHourWithDefaults

`func NewUsageAuditLogsHourWithDefaults() *UsageAuditLogsHour`

NewUsageAuditLogsHourWithDefaults instantiates a new UsageAuditLogsHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetHour

`func (o *UsageAuditLogsHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageAuditLogsHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageAuditLogsHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageAuditLogsHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetLinesIndexed

`func (o *UsageAuditLogsHour) GetLinesIndexed() int64`

GetLinesIndexed returns the LinesIndexed field if non-nil, zero value otherwise.

### GetLinesIndexedOk

`func (o *UsageAuditLogsHour) GetLinesIndexedOk() (*int64, bool)`

GetLinesIndexedOk returns a tuple with the LinesIndexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesIndexed

`func (o *UsageAuditLogsHour) SetLinesIndexed(v int64)`

SetLinesIndexed sets LinesIndexed field to given value.

### HasLinesIndexed

`func (o *UsageAuditLogsHour) HasLinesIndexed() bool`

HasLinesIndexed returns a boolean if a field has been set.

### GetOrgName

`func (o *UsageAuditLogsHour) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UsageAuditLogsHour) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UsageAuditLogsHour) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UsageAuditLogsHour) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPublicId

`func (o *UsageAuditLogsHour) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UsageAuditLogsHour) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UsageAuditLogsHour) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UsageAuditLogsHour) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


