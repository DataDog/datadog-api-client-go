# UsageNetworkFlowsHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hour** | Pointer to **time.Time** | The hour for the usage. | [optional] 
**IndexedEventCount** | Pointer to **int64** | Contains the number of netflow events indexed. | [optional] 

## Methods

### NewUsageNetworkFlowsHour

`func NewUsageNetworkFlowsHour() *UsageNetworkFlowsHour`

NewUsageNetworkFlowsHour instantiates a new UsageNetworkFlowsHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageNetworkFlowsHourWithDefaults

`func NewUsageNetworkFlowsHourWithDefaults() *UsageNetworkFlowsHour`

NewUsageNetworkFlowsHourWithDefaults instantiates a new UsageNetworkFlowsHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHour

`func (o *UsageNetworkFlowsHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageNetworkFlowsHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageNetworkFlowsHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageNetworkFlowsHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetIndexedEventCount

`func (o *UsageNetworkFlowsHour) GetIndexedEventCount() int64`

GetIndexedEventCount returns the IndexedEventCount field if non-nil, zero value otherwise.

### GetIndexedEventCountOk

`func (o *UsageNetworkFlowsHour) GetIndexedEventCountOk() (*int64, bool)`

GetIndexedEventCountOk returns a tuple with the IndexedEventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedEventCount

`func (o *UsageNetworkFlowsHour) SetIndexedEventCount(v int64)`

SetIndexedEventCount sets IndexedEventCount field to given value.

### HasIndexedEventCount

`func (o *UsageNetworkFlowsHour) HasIndexedEventCount() bool`

HasIndexedEventCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


