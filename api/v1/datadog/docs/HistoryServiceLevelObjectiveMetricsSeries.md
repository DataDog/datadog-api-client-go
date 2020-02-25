# HistoryServiceLevelObjectiveMetricsSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | Count of submitted metrics | 
**Metadata** | Pointer to [**HistoryServiceLevelObjectiveMetricsSeriesMetadata**](HistoryServiceLevelObjectiveMetricsSeries_metadata.md) |  | 
**Sum** | Pointer to **float64** | Total Sum of the query | 
**Values** | Pointer to **[]float64** | The query values | 

## Methods

### NewHistoryServiceLevelObjectiveMetricsSeries

`func NewHistoryServiceLevelObjectiveMetricsSeries(count int64, metadata HistoryServiceLevelObjectiveMetricsSeriesMetadata, sum float64, values []float64, ) *HistoryServiceLevelObjectiveMetricsSeries`

NewHistoryServiceLevelObjectiveMetricsSeries instantiates a new HistoryServiceLevelObjectiveMetricsSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryServiceLevelObjectiveMetricsSeriesWithDefaults

`func NewHistoryServiceLevelObjectiveMetricsSeriesWithDefaults() *HistoryServiceLevelObjectiveMetricsSeries`

NewHistoryServiceLevelObjectiveMetricsSeriesWithDefaults instantiates a new HistoryServiceLevelObjectiveMetricsSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *HistoryServiceLevelObjectiveMetricsSeries) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *HistoryServiceLevelObjectiveMetricsSeries) GetCountOk() (int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCount

`func (o *HistoryServiceLevelObjectiveMetricsSeries) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCount

`func (o *HistoryServiceLevelObjectiveMetricsSeries) SetCount(v int64)`

SetCount gets a reference to the given int64 and assigns it to the Count field.

### GetMetadata

`func (o *HistoryServiceLevelObjectiveMetricsSeries) GetMetadata() HistoryServiceLevelObjectiveMetricsSeriesMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HistoryServiceLevelObjectiveMetricsSeries) GetMetadataOk() (HistoryServiceLevelObjectiveMetricsSeriesMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMetadata

`func (o *HistoryServiceLevelObjectiveMetricsSeries) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadata

`func (o *HistoryServiceLevelObjectiveMetricsSeries) SetMetadata(v HistoryServiceLevelObjectiveMetricsSeriesMetadata)`

SetMetadata gets a reference to the given HistoryServiceLevelObjectiveMetricsSeriesMetadata and assigns it to the Metadata field.

### GetSum

`func (o *HistoryServiceLevelObjectiveMetricsSeries) GetSum() float64`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *HistoryServiceLevelObjectiveMetricsSeries) GetSumOk() (float64, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSum

`func (o *HistoryServiceLevelObjectiveMetricsSeries) HasSum() bool`

HasSum returns a boolean if a field has been set.

### SetSum

`func (o *HistoryServiceLevelObjectiveMetricsSeries) SetSum(v float64)`

SetSum gets a reference to the given float64 and assigns it to the Sum field.

### GetValues

`func (o *HistoryServiceLevelObjectiveMetricsSeries) GetValues() []float64`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *HistoryServiceLevelObjectiveMetricsSeries) GetValuesOk() ([]float64, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValues

`func (o *HistoryServiceLevelObjectiveMetricsSeries) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValues

`func (o *HistoryServiceLevelObjectiveMetricsSeries) SetValues(v []float64)`

SetValues gets a reference to the given []float64 and assigns it to the Values field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


