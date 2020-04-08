# ChangeWidgetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApmQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**ChangeType** | Pointer to [**WidgetChangeType**](WidgetChangeType.md) |  | [optional] 
**CompareTo** | Pointer to [**WidgetCompareTo**](WidgetCompareTo.md) |  | [optional] 
**EventQuery** | Pointer to [**EventQueryDefinition**](EventQueryDefinition.md) |  | [optional] 
**IncreaseGood** | Pointer to **bool** | Whether to show increase as good. | [optional] 
**LogQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**NetworkQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**OrderBy** | Pointer to [**WidgetOrderBy**](WidgetOrderBy.md) |  | [optional] 
**OrderDir** | Pointer to [**WidgetSort**](WidgetSort.md) |  | [optional] 
**ProcessQuery** | Pointer to [**ProcessQueryDefinition**](ProcessQueryDefinition.md) |  | [optional] 
**Q** | Pointer to **string** | TODO. | [optional] 
**RumQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**ShowPresent** | Pointer to **bool** | Whether to show the present value. | [optional] 

## Methods

### NewChangeWidgetRequest

`func NewChangeWidgetRequest() *ChangeWidgetRequest`

NewChangeWidgetRequest instantiates a new ChangeWidgetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeWidgetRequestWithDefaults

`func NewChangeWidgetRequestWithDefaults() *ChangeWidgetRequest`

NewChangeWidgetRequestWithDefaults instantiates a new ChangeWidgetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApmQuery

`func (o *ChangeWidgetRequest) GetApmQuery() LogQueryDefinition`

GetApmQuery returns the ApmQuery field if non-nil, zero value otherwise.

### GetApmQueryOk

`func (o *ChangeWidgetRequest) GetApmQueryOk() (*LogQueryDefinition, bool)`

GetApmQueryOk returns a tuple with the ApmQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmQuery

`func (o *ChangeWidgetRequest) SetApmQuery(v LogQueryDefinition)`

SetApmQuery sets ApmQuery field to given value.

### HasApmQuery

`func (o *ChangeWidgetRequest) HasApmQuery() bool`

HasApmQuery returns a boolean if a field has been set.

### GetChangeType

`func (o *ChangeWidgetRequest) GetChangeType() WidgetChangeType`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *ChangeWidgetRequest) GetChangeTypeOk() (*WidgetChangeType, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *ChangeWidgetRequest) SetChangeType(v WidgetChangeType)`

SetChangeType sets ChangeType field to given value.

### HasChangeType

`func (o *ChangeWidgetRequest) HasChangeType() bool`

HasChangeType returns a boolean if a field has been set.

### GetCompareTo

`func (o *ChangeWidgetRequest) GetCompareTo() WidgetCompareTo`

GetCompareTo returns the CompareTo field if non-nil, zero value otherwise.

### GetCompareToOk

`func (o *ChangeWidgetRequest) GetCompareToOk() (*WidgetCompareTo, bool)`

GetCompareToOk returns a tuple with the CompareTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareTo

`func (o *ChangeWidgetRequest) SetCompareTo(v WidgetCompareTo)`

SetCompareTo sets CompareTo field to given value.

### HasCompareTo

`func (o *ChangeWidgetRequest) HasCompareTo() bool`

HasCompareTo returns a boolean if a field has been set.

### GetEventQuery

`func (o *ChangeWidgetRequest) GetEventQuery() EventQueryDefinition`

GetEventQuery returns the EventQuery field if non-nil, zero value otherwise.

### GetEventQueryOk

`func (o *ChangeWidgetRequest) GetEventQueryOk() (*EventQueryDefinition, bool)`

GetEventQueryOk returns a tuple with the EventQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventQuery

`func (o *ChangeWidgetRequest) SetEventQuery(v EventQueryDefinition)`

SetEventQuery sets EventQuery field to given value.

### HasEventQuery

`func (o *ChangeWidgetRequest) HasEventQuery() bool`

HasEventQuery returns a boolean if a field has been set.

### GetIncreaseGood

`func (o *ChangeWidgetRequest) GetIncreaseGood() bool`

GetIncreaseGood returns the IncreaseGood field if non-nil, zero value otherwise.

### GetIncreaseGoodOk

`func (o *ChangeWidgetRequest) GetIncreaseGoodOk() (*bool, bool)`

GetIncreaseGoodOk returns a tuple with the IncreaseGood field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncreaseGood

`func (o *ChangeWidgetRequest) SetIncreaseGood(v bool)`

SetIncreaseGood sets IncreaseGood field to given value.

### HasIncreaseGood

`func (o *ChangeWidgetRequest) HasIncreaseGood() bool`

HasIncreaseGood returns a boolean if a field has been set.

### GetLogQuery

`func (o *ChangeWidgetRequest) GetLogQuery() LogQueryDefinition`

GetLogQuery returns the LogQuery field if non-nil, zero value otherwise.

### GetLogQueryOk

`func (o *ChangeWidgetRequest) GetLogQueryOk() (*LogQueryDefinition, bool)`

GetLogQueryOk returns a tuple with the LogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQuery

`func (o *ChangeWidgetRequest) SetLogQuery(v LogQueryDefinition)`

SetLogQuery sets LogQuery field to given value.

### HasLogQuery

`func (o *ChangeWidgetRequest) HasLogQuery() bool`

HasLogQuery returns a boolean if a field has been set.

### GetNetworkQuery

`func (o *ChangeWidgetRequest) GetNetworkQuery() LogQueryDefinition`

GetNetworkQuery returns the NetworkQuery field if non-nil, zero value otherwise.

### GetNetworkQueryOk

`func (o *ChangeWidgetRequest) GetNetworkQueryOk() (*LogQueryDefinition, bool)`

GetNetworkQueryOk returns a tuple with the NetworkQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkQuery

`func (o *ChangeWidgetRequest) SetNetworkQuery(v LogQueryDefinition)`

SetNetworkQuery sets NetworkQuery field to given value.

### HasNetworkQuery

`func (o *ChangeWidgetRequest) HasNetworkQuery() bool`

HasNetworkQuery returns a boolean if a field has been set.

### GetOrderBy

`func (o *ChangeWidgetRequest) GetOrderBy() WidgetOrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ChangeWidgetRequest) GetOrderByOk() (*WidgetOrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ChangeWidgetRequest) SetOrderBy(v WidgetOrderBy)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *ChangeWidgetRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetOrderDir

`func (o *ChangeWidgetRequest) GetOrderDir() WidgetSort`

GetOrderDir returns the OrderDir field if non-nil, zero value otherwise.

### GetOrderDirOk

`func (o *ChangeWidgetRequest) GetOrderDirOk() (*WidgetSort, bool)`

GetOrderDirOk returns a tuple with the OrderDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDir

`func (o *ChangeWidgetRequest) SetOrderDir(v WidgetSort)`

SetOrderDir sets OrderDir field to given value.

### HasOrderDir

`func (o *ChangeWidgetRequest) HasOrderDir() bool`

HasOrderDir returns a boolean if a field has been set.

### GetProcessQuery

`func (o *ChangeWidgetRequest) GetProcessQuery() ProcessQueryDefinition`

GetProcessQuery returns the ProcessQuery field if non-nil, zero value otherwise.

### GetProcessQueryOk

`func (o *ChangeWidgetRequest) GetProcessQueryOk() (*ProcessQueryDefinition, bool)`

GetProcessQueryOk returns a tuple with the ProcessQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessQuery

`func (o *ChangeWidgetRequest) SetProcessQuery(v ProcessQueryDefinition)`

SetProcessQuery sets ProcessQuery field to given value.

### HasProcessQuery

`func (o *ChangeWidgetRequest) HasProcessQuery() bool`

HasProcessQuery returns a boolean if a field has been set.

### GetQ

`func (o *ChangeWidgetRequest) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *ChangeWidgetRequest) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *ChangeWidgetRequest) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *ChangeWidgetRequest) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetRumQuery

`func (o *ChangeWidgetRequest) GetRumQuery() LogQueryDefinition`

GetRumQuery returns the RumQuery field if non-nil, zero value otherwise.

### GetRumQueryOk

`func (o *ChangeWidgetRequest) GetRumQueryOk() (*LogQueryDefinition, bool)`

GetRumQueryOk returns a tuple with the RumQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumQuery

`func (o *ChangeWidgetRequest) SetRumQuery(v LogQueryDefinition)`

SetRumQuery sets RumQuery field to given value.

### HasRumQuery

`func (o *ChangeWidgetRequest) HasRumQuery() bool`

HasRumQuery returns a boolean if a field has been set.

### GetShowPresent

`func (o *ChangeWidgetRequest) GetShowPresent() bool`

GetShowPresent returns the ShowPresent field if non-nil, zero value otherwise.

### GetShowPresentOk

`func (o *ChangeWidgetRequest) GetShowPresentOk() (*bool, bool)`

GetShowPresentOk returns a tuple with the ShowPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPresent

`func (o *ChangeWidgetRequest) SetShowPresent(v bool)`

SetShowPresent sets ShowPresent field to given value.

### HasShowPresent

`func (o *ChangeWidgetRequest) HasShowPresent() bool`

HasShowPresent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


