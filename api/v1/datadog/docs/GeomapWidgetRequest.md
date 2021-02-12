# GeomapWidgetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**Q** | Pointer to **string** | The widget metrics query. | [optional] 
**RumQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 

## Methods

### NewGeomapWidgetRequest

`func NewGeomapWidgetRequest() *GeomapWidgetRequest`

NewGeomapWidgetRequest instantiates a new GeomapWidgetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeomapWidgetRequestWithDefaults

`func NewGeomapWidgetRequestWithDefaults() *GeomapWidgetRequest`

NewGeomapWidgetRequestWithDefaults instantiates a new GeomapWidgetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogQuery

`func (o *GeomapWidgetRequest) GetLogQuery() LogQueryDefinition`

GetLogQuery returns the LogQuery field if non-nil, zero value otherwise.

### GetLogQueryOk

`func (o *GeomapWidgetRequest) GetLogQueryOk() (*LogQueryDefinition, bool)`

GetLogQueryOk returns a tuple with the LogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQuery

`func (o *GeomapWidgetRequest) SetLogQuery(v LogQueryDefinition)`

SetLogQuery sets LogQuery field to given value.

### HasLogQuery

`func (o *GeomapWidgetRequest) HasLogQuery() bool`

HasLogQuery returns a boolean if a field has been set.

### GetQ

`func (o *GeomapWidgetRequest) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *GeomapWidgetRequest) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *GeomapWidgetRequest) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *GeomapWidgetRequest) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetRumQuery

`func (o *GeomapWidgetRequest) GetRumQuery() LogQueryDefinition`

GetRumQuery returns the RumQuery field if non-nil, zero value otherwise.

### GetRumQueryOk

`func (o *GeomapWidgetRequest) GetRumQueryOk() (*LogQueryDefinition, bool)`

GetRumQueryOk returns a tuple with the RumQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumQuery

`func (o *GeomapWidgetRequest) SetRumQuery(v LogQueryDefinition)`

SetRumQuery sets RumQuery field to given value.

### HasRumQuery

`func (o *GeomapWidgetRequest) HasRumQuery() bool`

HasRumQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


