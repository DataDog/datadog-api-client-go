# AuditLogsQueryFilter

## Properties

| Name      | Type                  | Description                                                                                           | Notes                             |
| --------- | --------------------- | ----------------------------------------------------------------------------------------------------- | --------------------------------- |
| **From**  | Pointer to **string** | Minimum time for the requested events. Supports date, math, and regular timestamps (in milliseconds). | [optional] [default to "now-15m"] |
| **Query** | Pointer to **string** | Search query following the Audit Logs search syntax.                                                  | [optional] [default to "*"]       |
| **To**    | Pointer to **string** | Maximum time for the requested events. Supports date, math, and regular timestamps (in milliseconds). | [optional] [default to "now"]     |

## Methods

### NewAuditLogsQueryFilter

`func NewAuditLogsQueryFilter() *AuditLogsQueryFilter`

NewAuditLogsQueryFilter instantiates a new AuditLogsQueryFilter object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAuditLogsQueryFilterWithDefaults

`func NewAuditLogsQueryFilterWithDefaults() *AuditLogsQueryFilter`

NewAuditLogsQueryFilterWithDefaults instantiates a new AuditLogsQueryFilter object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetFrom

`func (o *AuditLogsQueryFilter) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *AuditLogsQueryFilter) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *AuditLogsQueryFilter) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *AuditLogsQueryFilter) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetQuery

`func (o *AuditLogsQueryFilter) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *AuditLogsQueryFilter) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *AuditLogsQueryFilter) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *AuditLogsQueryFilter) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTo

`func (o *AuditLogsQueryFilter) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *AuditLogsQueryFilter) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *AuditLogsQueryFilter) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *AuditLogsQueryFilter) HasTo() bool`

HasTo returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
