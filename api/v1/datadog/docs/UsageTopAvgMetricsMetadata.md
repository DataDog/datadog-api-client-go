# UsageTopAvgMetricsMetadata

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Day** | Pointer to **time.Time** | The day value from the user request that contains the returned usage data. (If day was used the request) | [optional] 
**Month** | Pointer to **time.Time** | The month value from the user request that contains the returned usage data. (If month was used the request) | [optional] 
**Pagination** | Pointer to [**UsageAttributionPagination**](UsageAttributionPagination.md) |  | [optional] 

## Methods

### NewUsageTopAvgMetricsMetadata

`func NewUsageTopAvgMetricsMetadata() *UsageTopAvgMetricsMetadata`

NewUsageTopAvgMetricsMetadata instantiates a new UsageTopAvgMetricsMetadata object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageTopAvgMetricsMetadataWithDefaults

`func NewUsageTopAvgMetricsMetadataWithDefaults() *UsageTopAvgMetricsMetadata`

NewUsageTopAvgMetricsMetadataWithDefaults instantiates a new UsageTopAvgMetricsMetadata object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDay

`func (o *UsageTopAvgMetricsMetadata) GetDay() time.Time`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *UsageTopAvgMetricsMetadata) GetDayOk() (*time.Time, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *UsageTopAvgMetricsMetadata) SetDay(v time.Time)`

SetDay sets Day field to given value.

### HasDay

`func (o *UsageTopAvgMetricsMetadata) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetMonth

`func (o *UsageTopAvgMetricsMetadata) GetMonth() time.Time`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *UsageTopAvgMetricsMetadata) GetMonthOk() (*time.Time, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *UsageTopAvgMetricsMetadata) SetMonth(v time.Time)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *UsageTopAvgMetricsMetadata) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetPagination

`func (o *UsageTopAvgMetricsMetadata) GetPagination() UsageAttributionPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *UsageTopAvgMetricsMetadata) GetPaginationOk() (*UsageAttributionPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *UsageTopAvgMetricsMetadata) SetPagination(v UsageAttributionPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *UsageTopAvgMetricsMetadata) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


