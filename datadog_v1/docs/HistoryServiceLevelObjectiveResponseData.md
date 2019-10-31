# HistoryServiceLevelObjectiveResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromTs** | Pointer to **int64** | the &#x60;from&#x60; timestamp in epoch seconds | [optional] 
**Groups** | Pointer to [**HistoryServiceLevelObjectiveGroups**](HistoryServiceLevelObjectiveGroups.md) |  | [optional] 
**Overall** | Pointer to [**HistoryServiceLevelObjectiveOverall**](HistoryServiceLevelObjectiveOverall.md) |  | [optional] 
**Series** | Pointer to [**HistoryServiceLevelObjectiveMetrics**](HistoryServiceLevelObjectiveMetrics.md) |  | [optional] 
**Thresholds** | Pointer to [**map[string]SloThreshold**](SloThreshold.md) | mapping of string timeframe to the SLO threshold. | [optional] 
**ToTs** | Pointer to **int64** | the &#x60;to&#x60; timestamp in epoch seconds | [optional] 

## Methods

### GetFromTs

`func (o *HistoryServiceLevelObjectiveResponseData) GetFromTs() int64`

GetFromTs returns the FromTs field if non-nil, zero value otherwise.

### GetFromTsOk

`func (o *HistoryServiceLevelObjectiveResponseData) GetFromTsOk() (int64, bool)`

GetFromTsOk returns a tuple with the FromTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFromTs

`func (o *HistoryServiceLevelObjectiveResponseData) HasFromTs() bool`

HasFromTs returns a boolean if a field has been set.

### SetFromTs

`func (o *HistoryServiceLevelObjectiveResponseData) SetFromTs(v int64)`

SetFromTs gets a reference to the given int64 and assigns it to the FromTs field.

### GetGroups

`func (o *HistoryServiceLevelObjectiveResponseData) GetGroups() HistoryServiceLevelObjectiveGroups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *HistoryServiceLevelObjectiveResponseData) GetGroupsOk() (HistoryServiceLevelObjectiveGroups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGroups

`func (o *HistoryServiceLevelObjectiveResponseData) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroups

`func (o *HistoryServiceLevelObjectiveResponseData) SetGroups(v HistoryServiceLevelObjectiveGroups)`

SetGroups gets a reference to the given HistoryServiceLevelObjectiveGroups and assigns it to the Groups field.

### GetOverall

`func (o *HistoryServiceLevelObjectiveResponseData) GetOverall() HistoryServiceLevelObjectiveOverall`

GetOverall returns the Overall field if non-nil, zero value otherwise.

### GetOverallOk

`func (o *HistoryServiceLevelObjectiveResponseData) GetOverallOk() (HistoryServiceLevelObjectiveOverall, bool)`

GetOverallOk returns a tuple with the Overall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOverall

`func (o *HistoryServiceLevelObjectiveResponseData) HasOverall() bool`

HasOverall returns a boolean if a field has been set.

### SetOverall

`func (o *HistoryServiceLevelObjectiveResponseData) SetOverall(v HistoryServiceLevelObjectiveOverall)`

SetOverall gets a reference to the given HistoryServiceLevelObjectiveOverall and assigns it to the Overall field.

### GetSeries

`func (o *HistoryServiceLevelObjectiveResponseData) GetSeries() HistoryServiceLevelObjectiveMetrics`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *HistoryServiceLevelObjectiveResponseData) GetSeriesOk() (HistoryServiceLevelObjectiveMetrics, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSeries

`func (o *HistoryServiceLevelObjectiveResponseData) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### SetSeries

`func (o *HistoryServiceLevelObjectiveResponseData) SetSeries(v HistoryServiceLevelObjectiveMetrics)`

SetSeries gets a reference to the given HistoryServiceLevelObjectiveMetrics and assigns it to the Series field.

### GetThresholds

`func (o *HistoryServiceLevelObjectiveResponseData) GetThresholds() map[string]SloThreshold`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *HistoryServiceLevelObjectiveResponseData) GetThresholdsOk() (map[string]SloThreshold, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasThresholds

`func (o *HistoryServiceLevelObjectiveResponseData) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### SetThresholds

`func (o *HistoryServiceLevelObjectiveResponseData) SetThresholds(v map[string]SloThreshold)`

SetThresholds gets a reference to the given map[string]SloThreshold and assigns it to the Thresholds field.

### GetToTs

`func (o *HistoryServiceLevelObjectiveResponseData) GetToTs() int64`

GetToTs returns the ToTs field if non-nil, zero value otherwise.

### GetToTsOk

`func (o *HistoryServiceLevelObjectiveResponseData) GetToTsOk() (int64, bool)`

GetToTsOk returns a tuple with the ToTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToTs

`func (o *HistoryServiceLevelObjectiveResponseData) HasToTs() bool`

HasToTs returns a boolean if a field has been set.

### SetToTs

`func (o *HistoryServiceLevelObjectiveResponseData) SetToTs(v int64)`

SetToTs gets a reference to the given int64 and assigns it to the ToTs field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


