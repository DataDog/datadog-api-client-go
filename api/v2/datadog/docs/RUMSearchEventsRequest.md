# RUMSearchEventsRequest

## Properties

| Name        | Type                                                         | Description | Notes      |
| ----------- | ------------------------------------------------------------ | ----------- | ---------- |
| **Filter**  | Pointer to [**RUMQueryFilter**](RUMQueryFilter.md)           |             | [optional] |
| **Options** | Pointer to [**RUMQueryOptions**](RUMQueryOptions.md)         |             | [optional] |
| **Page**    | Pointer to [**RUMQueryPageOptions**](RUMQueryPageOptions.md) |             | [optional] |
| **Sort**    | Pointer to [**RUMSort**](RUMSort.md)                         |             | [optional] |

## Methods

### NewRUMSearchEventsRequest

`func NewRUMSearchEventsRequest() *RUMSearchEventsRequest`

NewRUMSearchEventsRequest instantiates a new RUMSearchEventsRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRUMSearchEventsRequestWithDefaults

`func NewRUMSearchEventsRequestWithDefaults() *RUMSearchEventsRequest`

NewRUMSearchEventsRequestWithDefaults instantiates a new RUMSearchEventsRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetFilter

`func (o *RUMSearchEventsRequest) GetFilter() RUMQueryFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *RUMSearchEventsRequest) GetFilterOk() (*RUMQueryFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *RUMSearchEventsRequest) SetFilter(v RUMQueryFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *RUMSearchEventsRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetOptions

`func (o *RUMSearchEventsRequest) GetOptions() RUMQueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *RUMSearchEventsRequest) GetOptionsOk() (*RUMQueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *RUMSearchEventsRequest) SetOptions(v RUMQueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *RUMSearchEventsRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPage

`func (o *RUMSearchEventsRequest) GetPage() RUMQueryPageOptions`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *RUMSearchEventsRequest) GetPageOk() (*RUMQueryPageOptions, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *RUMSearchEventsRequest) SetPage(v RUMQueryPageOptions)`

SetPage sets Page field to given value.

### HasPage

`func (o *RUMSearchEventsRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSort

`func (o *RUMSearchEventsRequest) GetSort() RUMSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *RUMSearchEventsRequest) GetSortOk() (*RUMSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *RUMSearchEventsRequest) SetSort(v RUMSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *RUMSearchEventsRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
