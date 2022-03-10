# AuditLogsSearchEventsRequest

## Properties

| Name        | Type                                                                     | Description | Notes      |
| ----------- | ------------------------------------------------------------------------ | ----------- | ---------- |
| **Filter**  | Pointer to [**AuditLogsQueryFilter**](AuditLogsQueryFilter.md)           |             | [optional] |
| **Options** | Pointer to [**AuditLogsQueryOptions**](AuditLogsQueryOptions.md)         |             | [optional] |
| **Page**    | Pointer to [**AuditLogsQueryPageOptions**](AuditLogsQueryPageOptions.md) |             | [optional] |
| **Sort**    | Pointer to [**AuditLogsSort**](AuditLogsSort.md)                         |             | [optional] |

## Methods

### NewAuditLogsSearchEventsRequest

`func NewAuditLogsSearchEventsRequest() *AuditLogsSearchEventsRequest`

NewAuditLogsSearchEventsRequest instantiates a new AuditLogsSearchEventsRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAuditLogsSearchEventsRequestWithDefaults

`func NewAuditLogsSearchEventsRequestWithDefaults() *AuditLogsSearchEventsRequest`

NewAuditLogsSearchEventsRequestWithDefaults instantiates a new AuditLogsSearchEventsRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetFilter

`func (o *AuditLogsSearchEventsRequest) GetFilter() AuditLogsQueryFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AuditLogsSearchEventsRequest) GetFilterOk() (*AuditLogsQueryFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AuditLogsSearchEventsRequest) SetFilter(v AuditLogsQueryFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AuditLogsSearchEventsRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetOptions

`func (o *AuditLogsSearchEventsRequest) GetOptions() AuditLogsQueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AuditLogsSearchEventsRequest) GetOptionsOk() (*AuditLogsQueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AuditLogsSearchEventsRequest) SetOptions(v AuditLogsQueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AuditLogsSearchEventsRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPage

`func (o *AuditLogsSearchEventsRequest) GetPage() AuditLogsQueryPageOptions`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AuditLogsSearchEventsRequest) GetPageOk() (*AuditLogsQueryPageOptions, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AuditLogsSearchEventsRequest) SetPage(v AuditLogsQueryPageOptions)`

SetPage sets Page field to given value.

### HasPage

`func (o *AuditLogsSearchEventsRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSort

`func (o *AuditLogsSearchEventsRequest) GetSort() AuditLogsSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *AuditLogsSearchEventsRequest) GetSortOk() (*AuditLogsSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *AuditLogsSearchEventsRequest) SetSort(v AuditLogsSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *AuditLogsSearchEventsRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
