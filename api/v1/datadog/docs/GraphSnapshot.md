# GraphSnapshot

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**GraphDef** | Pointer to **string** | A JSON document defining the graph. &#x60;graph_def&#x60; can be used instead of &#x60;metric_query&#x60;. The JSON document uses the [grammar defined here](https://docs.datadoghq.com/graphing/graphing_json/#grammar) and should be formatted to a single line then URL encoded. | [optional] 
**MetricQuery** | Pointer to **string** | The metric query. One of &#x60;metric_query&#x60; or &#x60;graph_def&#x60; is required. | [optional] 
**SnapshotUrl** | Pointer to **string** | URL of your [graph snapshot](https://docs.datadoghq.com/metrics/explorer/#snapshot). | [optional] 

## Methods

### NewGraphSnapshot

`func NewGraphSnapshot() *GraphSnapshot`

NewGraphSnapshot instantiates a new GraphSnapshot object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewGraphSnapshotWithDefaults

`func NewGraphSnapshotWithDefaults() *GraphSnapshot`

NewGraphSnapshotWithDefaults instantiates a new GraphSnapshot object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetGraphDef

`func (o *GraphSnapshot) GetGraphDef() string`

GetGraphDef returns the GraphDef field if non-nil, zero value otherwise.

### GetGraphDefOk

`func (o *GraphSnapshot) GetGraphDefOk() (*string, bool)`

GetGraphDefOk returns a tuple with the GraphDef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphDef

`func (o *GraphSnapshot) SetGraphDef(v string)`

SetGraphDef sets GraphDef field to given value.

### HasGraphDef

`func (o *GraphSnapshot) HasGraphDef() bool`

HasGraphDef returns a boolean if a field has been set.

### GetMetricQuery

`func (o *GraphSnapshot) GetMetricQuery() string`

GetMetricQuery returns the MetricQuery field if non-nil, zero value otherwise.

### GetMetricQueryOk

`func (o *GraphSnapshot) GetMetricQueryOk() (*string, bool)`

GetMetricQueryOk returns a tuple with the MetricQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricQuery

`func (o *GraphSnapshot) SetMetricQuery(v string)`

SetMetricQuery sets MetricQuery field to given value.

### HasMetricQuery

`func (o *GraphSnapshot) HasMetricQuery() bool`

HasMetricQuery returns a boolean if a field has been set.

### GetSnapshotUrl

`func (o *GraphSnapshot) GetSnapshotUrl() string`

GetSnapshotUrl returns the SnapshotUrl field if non-nil, zero value otherwise.

### GetSnapshotUrlOk

`func (o *GraphSnapshot) GetSnapshotUrlOk() (*string, bool)`

GetSnapshotUrlOk returns a tuple with the SnapshotUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotUrl

`func (o *GraphSnapshot) SetSnapshotUrl(v string)`

SetSnapshotUrl sets SnapshotUrl field to given value.

### HasSnapshotUrl

`func (o *GraphSnapshot) HasSnapshotUrl() bool`

HasSnapshotUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


