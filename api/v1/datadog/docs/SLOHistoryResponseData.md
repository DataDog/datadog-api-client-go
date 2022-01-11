# SLOHistoryResponseData

## Properties

| Name           | Type                                                       | Description                                                                                                                                                                                          | Notes      |
| -------------- | ---------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **FromTs**     | Pointer to **int64**                                       | The &#x60;from&#x60; timestamp in epoch seconds.                                                                                                                                                     | [optional] |
| **GroupBy**    | Pointer to **[]string**                                    | For &#x60;metric&#x60; based SLOs where the query includes a group-by clause, this represents the list of grouping parameters. This is not included in responses for &#x60;monitor&#x60; based SLOs. | [optional] |
| **Groups**     | Pointer to [**[]SLOHistoryMonitor**](SLOHistoryMonitor.md) | For grouped SLOs, this represents SLI data for specific groups. This is not included in the responses for &#x60;metric&#x60; based SLOs.                                                             | [optional] |
| **Monitors**   | Pointer to [**[]SLOHistoryMonitor**](SLOHistoryMonitor.md) | For multi-monitor SLOs, this represents SLI data for specific monitors. This is not included in the responses for &#x60;metric&#x60; based SLOs.                                                     | [optional] |
| **Overall**    | Pointer to [**SLOHistorySLIData**](SLOHistorySLIData.md)   |                                                                                                                                                                                                      | [optional] |
| **Series**     | Pointer to [**SLOHistoryMetrics**](SLOHistoryMetrics.md)   |                                                                                                                                                                                                      | [optional] |
| **Thresholds** | Pointer to [**map[string]SLOThreshold**](SLOThreshold.md)  | mapping of string timeframe to the SLO threshold.                                                                                                                                                    | [optional] |
| **ToTs**       | Pointer to **int64**                                       | The &#x60;to&#x60; timestamp in epoch seconds.                                                                                                                                                       | [optional] |
| **Type**       | Pointer to [**SLOType**](SLOType.md)                       |                                                                                                                                                                                                      | [optional] |
| **TypeId**     | Pointer to [**SLOTypeNumeric**](SLOTypeNumeric.md)         |                                                                                                                                                                                                      | [optional] |

## Methods

### NewSLOHistoryResponseData

`func NewSLOHistoryResponseData() *SLOHistoryResponseData`

NewSLOHistoryResponseData instantiates a new SLOHistoryResponseData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSLOHistoryResponseDataWithDefaults

`func NewSLOHistoryResponseDataWithDefaults() *SLOHistoryResponseData`

NewSLOHistoryResponseDataWithDefaults instantiates a new SLOHistoryResponseData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetFromTs

`func (o *SLOHistoryResponseData) GetFromTs() int64`

GetFromTs returns the FromTs field if non-nil, zero value otherwise.

### GetFromTsOk

`func (o *SLOHistoryResponseData) GetFromTsOk() (*int64, bool)`

GetFromTsOk returns a tuple with the FromTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTs

`func (o *SLOHistoryResponseData) SetFromTs(v int64)`

SetFromTs sets FromTs field to given value.

### HasFromTs

`func (o *SLOHistoryResponseData) HasFromTs() bool`

HasFromTs returns a boolean if a field has been set.

### GetGroupBy

`func (o *SLOHistoryResponseData) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *SLOHistoryResponseData) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *SLOHistoryResponseData) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *SLOHistoryResponseData) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetGroups

`func (o *SLOHistoryResponseData) GetGroups() []SLOHistoryMonitor`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *SLOHistoryResponseData) GetGroupsOk() (*[]SLOHistoryMonitor, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *SLOHistoryResponseData) SetGroups(v []SLOHistoryMonitor)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *SLOHistoryResponseData) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetMonitors

`func (o *SLOHistoryResponseData) GetMonitors() []SLOHistoryMonitor`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *SLOHistoryResponseData) GetMonitorsOk() (*[]SLOHistoryMonitor, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *SLOHistoryResponseData) SetMonitors(v []SLOHistoryMonitor)`

SetMonitors sets Monitors field to given value.

### HasMonitors

`func (o *SLOHistoryResponseData) HasMonitors() bool`

HasMonitors returns a boolean if a field has been set.

### GetOverall

`func (o *SLOHistoryResponseData) GetOverall() SLOHistorySLIData`

GetOverall returns the Overall field if non-nil, zero value otherwise.

### GetOverallOk

`func (o *SLOHistoryResponseData) GetOverallOk() (*SLOHistorySLIData, bool)`

GetOverallOk returns a tuple with the Overall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverall

`func (o *SLOHistoryResponseData) SetOverall(v SLOHistorySLIData)`

SetOverall sets Overall field to given value.

### HasOverall

`func (o *SLOHistoryResponseData) HasOverall() bool`

HasOverall returns a boolean if a field has been set.

### GetSeries

`func (o *SLOHistoryResponseData) GetSeries() SLOHistoryMetrics`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *SLOHistoryResponseData) GetSeriesOk() (*SLOHistoryMetrics, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *SLOHistoryResponseData) SetSeries(v SLOHistoryMetrics)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *SLOHistoryResponseData) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetThresholds

`func (o *SLOHistoryResponseData) GetThresholds() map[string]SLOThreshold`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *SLOHistoryResponseData) GetThresholdsOk() (*map[string]SLOThreshold, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *SLOHistoryResponseData) SetThresholds(v map[string]SLOThreshold)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *SLOHistoryResponseData) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### GetToTs

`func (o *SLOHistoryResponseData) GetToTs() int64`

GetToTs returns the ToTs field if non-nil, zero value otherwise.

### GetToTsOk

`func (o *SLOHistoryResponseData) GetToTsOk() (*int64, bool)`

GetToTsOk returns a tuple with the ToTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTs

`func (o *SLOHistoryResponseData) SetToTs(v int64)`

SetToTs sets ToTs field to given value.

### HasToTs

`func (o *SLOHistoryResponseData) HasToTs() bool`

HasToTs returns a boolean if a field has been set.

### GetType

`func (o *SLOHistoryResponseData) GetType() SLOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SLOHistoryResponseData) GetTypeOk() (*SLOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SLOHistoryResponseData) SetType(v SLOType)`

SetType sets Type field to given value.

### HasType

`func (o *SLOHistoryResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeId

`func (o *SLOHistoryResponseData) GetTypeId() SLOTypeNumeric`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *SLOHistoryResponseData) GetTypeIdOk() (*SLOTypeNumeric, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *SLOHistoryResponseData) SetTypeId(v SLOTypeNumeric)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *SLOHistoryResponseData) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
