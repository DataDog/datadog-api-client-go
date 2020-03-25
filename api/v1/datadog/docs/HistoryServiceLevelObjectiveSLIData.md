# HistoryServiceLevelObjectiveSLIData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**History** | Pointer to [**[][]float64**](array.md) | For &#x60;monitor&#x60; based SLOs, this includes the aggregated history uptime time series. | [optional] 
**Name** | Pointer to **string** | For groups in a grouped SLO this is the group name. For monitors in a multi-monitor SLO this is the monitor name. | [optional] 
**Precision** | Pointer to **map[string]float64** | A mapping of threshold &#x60;timeframe&#x60; to number of accurate decimals, regardless of the from &amp;&amp; to timestamp. | [optional] 
**Preview** | Pointer to **bool** | For &#x60;monitor&#x60; based SLOs when &#x60;true&#x60; this indicates that a replay is in progress to give an accurate uptime calculation. | [optional] 
**SliValue** | Pointer to **float64** | The current SLI value of the SLO over the history window. | [optional] 
**SpanPrecision** | Pointer to **float64** | The amount of decimal places the SLI value is accurate to for the given from and to timestamp. | [optional] 
**Uptime** | Pointer to **float64** | Deprecated. Use &#x60;sli_value&#x60; instead. | [optional] 

## Methods

### NewHistoryServiceLevelObjectiveSLIData

`func NewHistoryServiceLevelObjectiveSLIData() *HistoryServiceLevelObjectiveSLIData`

NewHistoryServiceLevelObjectiveSLIData instantiates a new HistoryServiceLevelObjectiveSLIData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryServiceLevelObjectiveSLIDataWithDefaults

`func NewHistoryServiceLevelObjectiveSLIDataWithDefaults() *HistoryServiceLevelObjectiveSLIData`

NewHistoryServiceLevelObjectiveSLIDataWithDefaults instantiates a new HistoryServiceLevelObjectiveSLIData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistory

`func (o *HistoryServiceLevelObjectiveSLIData) GetHistory() [][]float64`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *HistoryServiceLevelObjectiveSLIData) GetHistoryOk() ([][]float64, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHistory

`func (o *HistoryServiceLevelObjectiveSLIData) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### SetHistory

`func (o *HistoryServiceLevelObjectiveSLIData) SetHistory(v [][]float64)`

SetHistory gets a reference to the given [][]float64 and assigns it to the History field.

### GetName

`func (o *HistoryServiceLevelObjectiveSLIData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HistoryServiceLevelObjectiveSLIData) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *HistoryServiceLevelObjectiveSLIData) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *HistoryServiceLevelObjectiveSLIData) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetPrecision

`func (o *HistoryServiceLevelObjectiveSLIData) GetPrecision() map[string]float64`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *HistoryServiceLevelObjectiveSLIData) GetPrecisionOk() (map[string]float64, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrecision

`func (o *HistoryServiceLevelObjectiveSLIData) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### SetPrecision

`func (o *HistoryServiceLevelObjectiveSLIData) SetPrecision(v map[string]float64)`

SetPrecision gets a reference to the given map[string]float64 and assigns it to the Precision field.

### GetPreview

`func (o *HistoryServiceLevelObjectiveSLIData) GetPreview() bool`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *HistoryServiceLevelObjectiveSLIData) GetPreviewOk() (bool, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreview

`func (o *HistoryServiceLevelObjectiveSLIData) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### SetPreview

`func (o *HistoryServiceLevelObjectiveSLIData) SetPreview(v bool)`

SetPreview gets a reference to the given bool and assigns it to the Preview field.

### GetSliValue

`func (o *HistoryServiceLevelObjectiveSLIData) GetSliValue() float64`

GetSliValue returns the SliValue field if non-nil, zero value otherwise.

### GetSliValueOk

`func (o *HistoryServiceLevelObjectiveSLIData) GetSliValueOk() (float64, bool)`

GetSliValueOk returns a tuple with the SliValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSliValue

`func (o *HistoryServiceLevelObjectiveSLIData) HasSliValue() bool`

HasSliValue returns a boolean if a field has been set.

### SetSliValue

`func (o *HistoryServiceLevelObjectiveSLIData) SetSliValue(v float64)`

SetSliValue gets a reference to the given float64 and assigns it to the SliValue field.

### GetSpanPrecision

`func (o *HistoryServiceLevelObjectiveSLIData) GetSpanPrecision() float64`

GetSpanPrecision returns the SpanPrecision field if non-nil, zero value otherwise.

### GetSpanPrecisionOk

`func (o *HistoryServiceLevelObjectiveSLIData) GetSpanPrecisionOk() (float64, bool)`

GetSpanPrecisionOk returns a tuple with the SpanPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpanPrecision

`func (o *HistoryServiceLevelObjectiveSLIData) HasSpanPrecision() bool`

HasSpanPrecision returns a boolean if a field has been set.

### SetSpanPrecision

`func (o *HistoryServiceLevelObjectiveSLIData) SetSpanPrecision(v float64)`

SetSpanPrecision gets a reference to the given float64 and assigns it to the SpanPrecision field.

### GetUptime

`func (o *HistoryServiceLevelObjectiveSLIData) GetUptime() float64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *HistoryServiceLevelObjectiveSLIData) GetUptimeOk() (float64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUptime

`func (o *HistoryServiceLevelObjectiveSLIData) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### SetUptime

`func (o *HistoryServiceLevelObjectiveSLIData) SetUptime(v float64)`

SetUptime gets a reference to the given float64 and assigns it to the Uptime field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


