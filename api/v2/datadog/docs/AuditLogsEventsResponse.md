# AuditLogsEventsResponse

## Properties

| Name      | Type                                                                     | Description                           | Notes      |
| --------- | ------------------------------------------------------------------------ | ------------------------------------- | ---------- |
| **Data**  | Pointer to [**[]AuditLogsEvent**](AuditLogsEvent.md)                     | Array of events matching the request. | [optional] |
| **Links** | Pointer to [**AuditLogsResponseLinks**](AuditLogsResponseLinks.md)       |                                       | [optional] |
| **Meta**  | Pointer to [**AuditLogsResponseMetadata**](AuditLogsResponseMetadata.md) |                                       | [optional] |

## Methods

### NewAuditLogsEventsResponse

`func NewAuditLogsEventsResponse() *AuditLogsEventsResponse`

NewAuditLogsEventsResponse instantiates a new AuditLogsEventsResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAuditLogsEventsResponseWithDefaults

`func NewAuditLogsEventsResponseWithDefaults() *AuditLogsEventsResponse`

NewAuditLogsEventsResponseWithDefaults instantiates a new AuditLogsEventsResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *AuditLogsEventsResponse) GetData() []AuditLogsEvent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuditLogsEventsResponse) GetDataOk() (*[]AuditLogsEvent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuditLogsEventsResponse) SetData(v []AuditLogsEvent)`

SetData sets Data field to given value.

### HasData

`func (o *AuditLogsEventsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *AuditLogsEventsResponse) GetLinks() AuditLogsResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuditLogsEventsResponse) GetLinksOk() (*AuditLogsResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuditLogsEventsResponse) SetLinks(v AuditLogsResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuditLogsEventsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *AuditLogsEventsResponse) GetMeta() AuditLogsResponseMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AuditLogsEventsResponse) GetMetaOk() (*AuditLogsResponseMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AuditLogsEventsResponse) SetMeta(v AuditLogsResponseMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AuditLogsEventsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
