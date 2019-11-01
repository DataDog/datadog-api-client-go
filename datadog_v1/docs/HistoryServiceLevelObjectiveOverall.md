# HistoryServiceLevelObjectiveOverall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**History** | Pointer to [**[][]float32**](array.md) | For &#x60;monitor&#x60; based SLOs, this includes the aggregated history uptime time series. | [optional] 
**Name** | Pointer to **string** | For &#x60;monitor&#x60; based SLOs this represents the overall group. | [optional] 
**Precision** | Pointer to **map[string]float32** | A mapping of threshold &#x60;timeframe&#x60; to number of accurate decimals, regardless of the from &amp;&amp; to timestamp. | [optional] 
**Preview** | Pointer to **bool** | For &#x60;monitor&#x60; based SLOs when &#x60;true&#x60; this indicates that a replay is in progress to give an accurate uptime calculation. | [optional] 
**SpanPrecision** | Pointer to **float32** | The amount of decimal places the uptime value is accurate to for the given from and to timestamp. | [optional] 
**Uptime** | Pointer to **float32** | The uptime value of the SLO history window. | [optional] 

## Methods

### GetHistory

`func (o *HistoryServiceLevelObjectiveOverall) GetHistory() [][]float32`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *HistoryServiceLevelObjectiveOverall) GetHistoryOk() ([][]float32, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHistory

`func (o *HistoryServiceLevelObjectiveOverall) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### SetHistory

`func (o *HistoryServiceLevelObjectiveOverall) SetHistory(v [][]float32)`

SetHistory gets a reference to the given [][]float32 and assigns it to the History field.

### GetName

`func (o *HistoryServiceLevelObjectiveOverall) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HistoryServiceLevelObjectiveOverall) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *HistoryServiceLevelObjectiveOverall) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *HistoryServiceLevelObjectiveOverall) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetPrecision

`func (o *HistoryServiceLevelObjectiveOverall) GetPrecision() map[string]float32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *HistoryServiceLevelObjectiveOverall) GetPrecisionOk() (map[string]float32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrecision

`func (o *HistoryServiceLevelObjectiveOverall) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### SetPrecision

`func (o *HistoryServiceLevelObjectiveOverall) SetPrecision(v map[string]float32)`

SetPrecision gets a reference to the given map[string]float32 and assigns it to the Precision field.

### GetPreview

`func (o *HistoryServiceLevelObjectiveOverall) GetPreview() bool`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *HistoryServiceLevelObjectiveOverall) GetPreviewOk() (bool, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreview

`func (o *HistoryServiceLevelObjectiveOverall) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### SetPreview

`func (o *HistoryServiceLevelObjectiveOverall) SetPreview(v bool)`

SetPreview gets a reference to the given bool and assigns it to the Preview field.

### GetSpanPrecision

`func (o *HistoryServiceLevelObjectiveOverall) GetSpanPrecision() float32`

GetSpanPrecision returns the SpanPrecision field if non-nil, zero value otherwise.

### GetSpanPrecisionOk

`func (o *HistoryServiceLevelObjectiveOverall) GetSpanPrecisionOk() (float32, bool)`

GetSpanPrecisionOk returns a tuple with the SpanPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpanPrecision

`func (o *HistoryServiceLevelObjectiveOverall) HasSpanPrecision() bool`

HasSpanPrecision returns a boolean if a field has been set.

### SetSpanPrecision

`func (o *HistoryServiceLevelObjectiveOverall) SetSpanPrecision(v float32)`

SetSpanPrecision gets a reference to the given float32 and assigns it to the SpanPrecision field.

### GetUptime

`func (o *HistoryServiceLevelObjectiveOverall) GetUptime() float32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *HistoryServiceLevelObjectiveOverall) GetUptimeOk() (float32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUptime

`func (o *HistoryServiceLevelObjectiveOverall) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### SetUptime

`func (o *HistoryServiceLevelObjectiveOverall) SetUptime(v float32)`

SetUptime gets a reference to the given float32 and assigns it to the Uptime field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


