# SLOHistorySLIData

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

### NewSLOHistorySLIData

`func NewSLOHistorySLIData() *SLOHistorySLIData`

NewSLOHistorySLIData instantiates a new SLOHistorySLIData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLOHistorySLIDataWithDefaults

`func NewSLOHistorySLIDataWithDefaults() *SLOHistorySLIData`

NewSLOHistorySLIDataWithDefaults instantiates a new SLOHistorySLIData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistory

`func (o *SLOHistorySLIData) GetHistory() [][]float64`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *SLOHistorySLIData) GetHistoryOk() (*[][]float64, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *SLOHistorySLIData) SetHistory(v [][]float64)`

SetHistory sets History field to given value.

### HasHistory

`func (o *SLOHistorySLIData) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetName

`func (o *SLOHistorySLIData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SLOHistorySLIData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SLOHistorySLIData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SLOHistorySLIData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrecision

`func (o *SLOHistorySLIData) GetPrecision() map[string]float64`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *SLOHistorySLIData) GetPrecisionOk() (*map[string]float64, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *SLOHistorySLIData) SetPrecision(v map[string]float64)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *SLOHistorySLIData) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetPreview

`func (o *SLOHistorySLIData) GetPreview() bool`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *SLOHistorySLIData) GetPreviewOk() (*bool, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *SLOHistorySLIData) SetPreview(v bool)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *SLOHistorySLIData) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetSliValue

`func (o *SLOHistorySLIData) GetSliValue() float64`

GetSliValue returns the SliValue field if non-nil, zero value otherwise.

### GetSliValueOk

`func (o *SLOHistorySLIData) GetSliValueOk() (*float64, bool)`

GetSliValueOk returns a tuple with the SliValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliValue

`func (o *SLOHistorySLIData) SetSliValue(v float64)`

SetSliValue sets SliValue field to given value.

### HasSliValue

`func (o *SLOHistorySLIData) HasSliValue() bool`

HasSliValue returns a boolean if a field has been set.

### GetSpanPrecision

`func (o *SLOHistorySLIData) GetSpanPrecision() float64`

GetSpanPrecision returns the SpanPrecision field if non-nil, zero value otherwise.

### GetSpanPrecisionOk

`func (o *SLOHistorySLIData) GetSpanPrecisionOk() (*float64, bool)`

GetSpanPrecisionOk returns a tuple with the SpanPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanPrecision

`func (o *SLOHistorySLIData) SetSpanPrecision(v float64)`

SetSpanPrecision sets SpanPrecision field to given value.

### HasSpanPrecision

`func (o *SLOHistorySLIData) HasSpanPrecision() bool`

HasSpanPrecision returns a boolean if a field has been set.

### GetUptime

`func (o *SLOHistorySLIData) GetUptime() float64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *SLOHistorySLIData) GetUptimeOk() (*float64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *SLOHistorySLIData) SetUptime(v float64)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *SLOHistorySLIData) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


