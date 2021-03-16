# UsageTopAvgMetricsMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | Pointer to **interface{}** | The day value from the user request that contains the returned usage data. (If day was used the request) | [optional] 
**Month** | Pointer to **interface{}** | The month value from the user request that contains the returned usage data. (If month was used the request) | [optional] 
**Pagination** | Pointer to [**UsageAttributionPagination**](UsageAttributionPagination.md) |  | [optional] 

## Methods

### NewUsageTopAvgMetricsMetadata

`func NewUsageTopAvgMetricsMetadata() *UsageTopAvgMetricsMetadata`

NewUsageTopAvgMetricsMetadata instantiates a new UsageTopAvgMetricsMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageTopAvgMetricsMetadataWithDefaults

`func NewUsageTopAvgMetricsMetadataWithDefaults() *UsageTopAvgMetricsMetadata`

NewUsageTopAvgMetricsMetadataWithDefaults instantiates a new UsageTopAvgMetricsMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *UsageTopAvgMetricsMetadata) GetDay() interface{}`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *UsageTopAvgMetricsMetadata) GetDayOk() (*interface{}, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *UsageTopAvgMetricsMetadata) SetDay(v interface{})`

SetDay sets Day field to given value.

### HasDay

`func (o *UsageTopAvgMetricsMetadata) HasDay() bool`

HasDay returns a boolean if a field has been set.

### SetDayNil

`func (o *UsageTopAvgMetricsMetadata) SetDayNil(b bool)`

 SetDayNil sets the value for Day to be an explicit nil

### UnsetDay
`func (o *UsageTopAvgMetricsMetadata) UnsetDay()`

UnsetDay ensures that no value is present for Day, not even an explicit nil
### GetMonth

`func (o *UsageTopAvgMetricsMetadata) GetMonth() interface{}`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *UsageTopAvgMetricsMetadata) GetMonthOk() (*interface{}, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *UsageTopAvgMetricsMetadata) SetMonth(v interface{})`

SetMonth sets Month field to given value.

### HasMonth

`func (o *UsageTopAvgMetricsMetadata) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### SetMonthNil

`func (o *UsageTopAvgMetricsMetadata) SetMonthNil(b bool)`

 SetMonthNil sets the value for Month to be an explicit nil

### UnsetMonth
`func (o *UsageTopAvgMetricsMetadata) UnsetMonth()`

UnsetMonth ensures that no value is present for Month, not even an explicit nil
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


