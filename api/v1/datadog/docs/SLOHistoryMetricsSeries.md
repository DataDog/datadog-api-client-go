# SLOHistoryMetricsSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | Count of submitted metrics. | 
**Metadata** | Pointer to [**SLOHistoryMetricsSeriesMetadata**](SLOHistoryMetricsSeries_metadata.md) |  | 
**Sum** | Pointer to **float64** | Total sum of the query. | 
**Values** | Pointer to **[]float64** | The query values. | 

## Methods

### NewSLOHistoryMetricsSeries

`func NewSLOHistoryMetricsSeries(count int64, metadata SLOHistoryMetricsSeriesMetadata, sum float64, values []float64, ) *SLOHistoryMetricsSeries`

NewSLOHistoryMetricsSeries instantiates a new SLOHistoryMetricsSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLOHistoryMetricsSeriesWithDefaults

`func NewSLOHistoryMetricsSeriesWithDefaults() *SLOHistoryMetricsSeries`

NewSLOHistoryMetricsSeriesWithDefaults instantiates a new SLOHistoryMetricsSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *SLOHistoryMetricsSeries) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SLOHistoryMetricsSeries) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SLOHistoryMetricsSeries) SetCount(v int64)`

SetCount sets Count field to given value.


### GetMetadata

`func (o *SLOHistoryMetricsSeries) GetMetadata() SLOHistoryMetricsSeriesMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SLOHistoryMetricsSeries) GetMetadataOk() (*SLOHistoryMetricsSeriesMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SLOHistoryMetricsSeries) SetMetadata(v SLOHistoryMetricsSeriesMetadata)`

SetMetadata sets Metadata field to given value.


### GetSum

`func (o *SLOHistoryMetricsSeries) GetSum() float64`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *SLOHistoryMetricsSeries) GetSumOk() (*float64, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *SLOHistoryMetricsSeries) SetSum(v float64)`

SetSum sets Sum field to given value.


### GetValues

`func (o *SLOHistoryMetricsSeries) GetValues() []float64`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SLOHistoryMetricsSeries) GetValuesOk() (*[]float64, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SLOHistoryMetricsSeries) SetValues(v []float64)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


