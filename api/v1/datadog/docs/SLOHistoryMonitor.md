# SLOHistoryMonitor

## Properties

| Name                     | Type                                                                                   | Description                                                                                                                                                                                                                                           | Notes      |
| ------------------------ | -------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **ErrorBudgetRemaining** | Pointer to **map[string]float64**                                                      | A mapping of threshold &#x60;timeframe&#x60; to the remaining error budget.                                                                                                                                                                           | [optional] |
| **Errors**               | Pointer to [**[]SLOHistoryResponseErrorWithType**](SLOHistoryResponseErrorWithType.md) | An array of error objects returned while querying the history data for the service level objective.                                                                                                                                                   | [optional] |
| **Group**                | Pointer to **string**                                                                  | For groups in a grouped SLO, this is the group name.                                                                                                                                                                                                  | [optional] |
| **History**              | Pointer to **[][]float64**                                                             | For &#x60;monitor&#x60; based SLOs, this includes the aggregated history as arrays that include time series and uptime data where &#x60;0&#x3D;monitor&#x60; is in &#x60;OK&#x60; state and &#x60;1&#x3D;monitor&#x60; is in &#x60;alert&#x60; state. | [optional] |
| **MonitorModified**      | Pointer to **int64**                                                                   | For &#x60;monitor&#x60; based SLOs, this is the last modified timestamp in epoch seconds of the monitor.                                                                                                                                              | [optional] |
| **MonitorType**          | Pointer to **string**                                                                  | For &#x60;monitor&#x60; based SLOs, this describes the type of monitor.                                                                                                                                                                               | [optional] |
| **Name**                 | Pointer to **string**                                                                  | For groups in a grouped SLO, this is the group name. For monitors in a multi-monitor SLO, this is the monitor name.                                                                                                                                   | [optional] |
| **Precision**            | Pointer to **float64**                                                                 | The amount of decimal places the SLI value is accurate to for the given from &#x60;&amp;&amp;&#x60; to timestamp. Use &#x60;span_precision&#x60; instead.                                                                                             | [optional] |
| **Preview**              | Pointer to **bool**                                                                    | For &#x60;monitor&#x60; based SLOs, when &#x60;true&#x60; this indicates that a replay is in progress to give an accurate uptime calculation.                                                                                                         | [optional] |
| **SliValue**             | Pointer to **float64**                                                                 | The current SLI value of the SLO over the history window.                                                                                                                                                                                             | [optional] |
| **SpanPrecision**        | Pointer to **float64**                                                                 | The amount of decimal places the SLI value is accurate to for the given from &#x60;&amp;&amp;&#x60; to timestamp.                                                                                                                                     | [optional] |
| **Uptime**               | Pointer to **float64**                                                                 | Use &#x60;sli_value&#x60; instead.                                                                                                                                                                                                                    | [optional] |

## Methods

### NewSLOHistoryMonitor

`func NewSLOHistoryMonitor() *SLOHistoryMonitor`

NewSLOHistoryMonitor instantiates a new SLOHistoryMonitor object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSLOHistoryMonitorWithDefaults

`func NewSLOHistoryMonitorWithDefaults() *SLOHistoryMonitor`

NewSLOHistoryMonitorWithDefaults instantiates a new SLOHistoryMonitor object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetErrorBudgetRemaining

`func (o *SLOHistoryMonitor) GetErrorBudgetRemaining() map[string]float64`

GetErrorBudgetRemaining returns the ErrorBudgetRemaining field if non-nil, zero value otherwise.

### GetErrorBudgetRemainingOk

`func (o *SLOHistoryMonitor) GetErrorBudgetRemainingOk() (*map[string]float64, bool)`

GetErrorBudgetRemainingOk returns a tuple with the ErrorBudgetRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorBudgetRemaining

`func (o *SLOHistoryMonitor) SetErrorBudgetRemaining(v map[string]float64)`

SetErrorBudgetRemaining sets ErrorBudgetRemaining field to given value.

### HasErrorBudgetRemaining

`func (o *SLOHistoryMonitor) HasErrorBudgetRemaining() bool`

HasErrorBudgetRemaining returns a boolean if a field has been set.

### GetErrors

`func (o *SLOHistoryMonitor) GetErrors() []SLOHistoryResponseErrorWithType`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SLOHistoryMonitor) GetErrorsOk() (*[]SLOHistoryResponseErrorWithType, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SLOHistoryMonitor) SetErrors(v []SLOHistoryResponseErrorWithType)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SLOHistoryMonitor) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetGroup

`func (o *SLOHistoryMonitor) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SLOHistoryMonitor) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SLOHistoryMonitor) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *SLOHistoryMonitor) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHistory

`func (o *SLOHistoryMonitor) GetHistory() [][]float64`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *SLOHistoryMonitor) GetHistoryOk() (*[][]float64, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *SLOHistoryMonitor) SetHistory(v [][]float64)`

SetHistory sets History field to given value.

### HasHistory

`func (o *SLOHistoryMonitor) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetMonitorModified

`func (o *SLOHistoryMonitor) GetMonitorModified() int64`

GetMonitorModified returns the MonitorModified field if non-nil, zero value otherwise.

### GetMonitorModifiedOk

`func (o *SLOHistoryMonitor) GetMonitorModifiedOk() (*int64, bool)`

GetMonitorModifiedOk returns a tuple with the MonitorModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorModified

`func (o *SLOHistoryMonitor) SetMonitorModified(v int64)`

SetMonitorModified sets MonitorModified field to given value.

### HasMonitorModified

`func (o *SLOHistoryMonitor) HasMonitorModified() bool`

HasMonitorModified returns a boolean if a field has been set.

### GetMonitorType

`func (o *SLOHistoryMonitor) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *SLOHistoryMonitor) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *SLOHistoryMonitor) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.

### HasMonitorType

`func (o *SLOHistoryMonitor) HasMonitorType() bool`

HasMonitorType returns a boolean if a field has been set.

### GetName

`func (o *SLOHistoryMonitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SLOHistoryMonitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SLOHistoryMonitor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SLOHistoryMonitor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrecision

`func (o *SLOHistoryMonitor) GetPrecision() float64`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *SLOHistoryMonitor) GetPrecisionOk() (*float64, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *SLOHistoryMonitor) SetPrecision(v float64)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *SLOHistoryMonitor) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetPreview

`func (o *SLOHistoryMonitor) GetPreview() bool`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *SLOHistoryMonitor) GetPreviewOk() (*bool, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *SLOHistoryMonitor) SetPreview(v bool)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *SLOHistoryMonitor) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetSliValue

`func (o *SLOHistoryMonitor) GetSliValue() float64`

GetSliValue returns the SliValue field if non-nil, zero value otherwise.

### GetSliValueOk

`func (o *SLOHistoryMonitor) GetSliValueOk() (*float64, bool)`

GetSliValueOk returns a tuple with the SliValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliValue

`func (o *SLOHistoryMonitor) SetSliValue(v float64)`

SetSliValue sets SliValue field to given value.

### HasSliValue

`func (o *SLOHistoryMonitor) HasSliValue() bool`

HasSliValue returns a boolean if a field has been set.

### GetSpanPrecision

`func (o *SLOHistoryMonitor) GetSpanPrecision() float64`

GetSpanPrecision returns the SpanPrecision field if non-nil, zero value otherwise.

### GetSpanPrecisionOk

`func (o *SLOHistoryMonitor) GetSpanPrecisionOk() (*float64, bool)`

GetSpanPrecisionOk returns a tuple with the SpanPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanPrecision

`func (o *SLOHistoryMonitor) SetSpanPrecision(v float64)`

SetSpanPrecision sets SpanPrecision field to given value.

### HasSpanPrecision

`func (o *SLOHistoryMonitor) HasSpanPrecision() bool`

HasSpanPrecision returns a boolean if a field has been set.

### GetUptime

`func (o *SLOHistoryMonitor) GetUptime() float64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *SLOHistoryMonitor) GetUptimeOk() (*float64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *SLOHistoryMonitor) SetUptime(v float64)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *SLOHistoryMonitor) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
