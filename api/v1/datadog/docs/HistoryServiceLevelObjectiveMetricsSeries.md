# HistoryServiceLevelObjectiveMetricsSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | Count of submitted metrics. | 
**Metadata** | Pointer to [**HistoryServiceLevelObjectiveMetricsSeriesMetadata**](HistoryServiceLevelObjectiveMetricsSeries_metadata.md) |  | 
**Sum** | Pointer to **float64** | Total Sum of the query. | 
**Values** | Pointer to **[]float64** | The query values. | 

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

`func (o *HistoryServiceLevelObjectiveMetricsSeries) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *HistoryServiceLevelObjectiveMetricsSeries) SetCount(v int64)`

SetCount sets Count field to given value.


### GetMetadata

`func (o *HistoryServiceLevelObjectiveMetricsSeries) GetMetadata() HistoryServiceLevelObjectiveMetricsSeriesMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HistoryServiceLevelObjectiveMetricsSeries) GetMetadataOk() (*HistoryServiceLevelObjectiveMetricsSeriesMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HistoryServiceLevelObjectiveMetricsSeries) SetMetadata(v HistoryServiceLevelObjectiveMetricsSeriesMetadata)`

SetMetadata sets Metadata field to given value.


### GetSum

`func (o *HistoryServiceLevelObjectiveMetricsSeries) GetSum() float64`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *HistoryServiceLevelObjectiveMetricsSeries) GetSumOk() (*float64, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *HistoryServiceLevelObjectiveMetricsSeries) SetSum(v float64)`

SetSum sets Sum field to given value.


### GetValues

`func (o *HistoryServiceLevelObjectiveMetricsSeries) GetValues() []float64`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *HistoryServiceLevelObjectiveMetricsSeries) GetValuesOk() (*[]float64, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *HistoryServiceLevelObjectiveMetricsSeries) SetValues(v []float64)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


