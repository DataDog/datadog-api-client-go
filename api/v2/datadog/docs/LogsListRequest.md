# LogsListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**LogsListRequestFilter**](LogsListRequest_filter.md) |  | [optional] 
**Page** | Pointer to [**LogsListRequestPage**](LogsListRequest_page.md) |  | [optional] 
**Sort** | Pointer to [**LogsSort**](LogsSort.md) |  | [optional] 

## Methods

### NewLogsListRequest

`func NewLogsListRequest() *LogsListRequest`

NewLogsListRequest instantiates a new LogsListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsListRequestWithDefaults

`func NewLogsListRequestWithDefaults() *LogsListRequest`

NewLogsListRequestWithDefaults instantiates a new LogsListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *LogsListRequest) GetFilter() LogsListRequestFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LogsListRequest) GetFilterOk() (*LogsListRequestFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LogsListRequest) SetFilter(v LogsListRequestFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *LogsListRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPage

`func (o *LogsListRequest) GetPage() LogsListRequestPage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *LogsListRequest) GetPageOk() (*LogsListRequestPage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *LogsListRequest) SetPage(v LogsListRequestPage)`

SetPage sets Page field to given value.

### HasPage

`func (o *LogsListRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSort

`func (o *LogsListRequest) GetSort() LogsSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *LogsListRequest) GetSortOk() (*LogsSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *LogsListRequest) SetSort(v LogsSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *LogsListRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


