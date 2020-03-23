# HistoryServiceLevelObjectiveResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromTs** | Pointer to **int64** | The &#x60;from&#x60; timestamp in epoch seconds. | [optional] 
**Groups** | Pointer to [**HistoryServiceLevelObjectiveGroups**](HistoryServiceLevelObjectiveGroups.md) |  | [optional] 
**Overall** | Pointer to [**HistoryServiceLevelObjectiveOverall**](HistoryServiceLevelObjectiveOverall.md) |  | [optional] 
**Series** | Pointer to [**HistoryServiceLevelObjectiveMetrics**](HistoryServiceLevelObjectiveMetrics.md) |  | [optional] 
**Thresholds** | Pointer to [**map[string]SLOThreshold**](SLOThreshold.md) | mapping of string timeframe to the SLO threshold. | [optional] 
**ToTs** | Pointer to **int64** | The &#x60;to&#x60; timestamp in epoch seconds. | [optional] 
**Type** | Pointer to [**ServiceLevelObjectiveType**](ServiceLevelObjectiveType.md) |  | [optional] 
**TypeId** | Pointer to [**ServiceLevelObjectiveTypeNumeric**](ServiceLevelObjectiveTypeNumeric.md) |  | [optional] 

## Methods

### NewHistoryServiceLevelObjectiveResponseData

`func NewHistoryServiceLevelObjectiveResponseData() *HistoryServiceLevelObjectiveResponseData`

NewHistoryServiceLevelObjectiveResponseData instantiates a new HistoryServiceLevelObjectiveResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryServiceLevelObjectiveResponseDataWithDefaults

`func NewHistoryServiceLevelObjectiveResponseDataWithDefaults() *HistoryServiceLevelObjectiveResponseData`

NewHistoryServiceLevelObjectiveResponseDataWithDefaults instantiates a new HistoryServiceLevelObjectiveResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

`func (o *HistoryServiceLevelObjectiveResponseData) GetThresholds() map[string]SLOThreshold`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *HistoryServiceLevelObjectiveResponseData) GetThresholdsOk() (map[string]SLOThreshold, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasThresholds

`func (o *HistoryServiceLevelObjectiveResponseData) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### SetThresholds

`func (o *HistoryServiceLevelObjectiveResponseData) SetThresholds(v map[string]SLOThreshold)`

SetThresholds gets a reference to the given map[string]SLOThreshold and assigns it to the Thresholds field.

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

### GetType

`func (o *HistoryServiceLevelObjectiveResponseData) GetType() ServiceLevelObjectiveType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HistoryServiceLevelObjectiveResponseData) GetTypeOk() (ServiceLevelObjectiveType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *HistoryServiceLevelObjectiveResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *HistoryServiceLevelObjectiveResponseData) SetType(v ServiceLevelObjectiveType)`

SetType gets a reference to the given ServiceLevelObjectiveType and assigns it to the Type field.

### GetTypeId

`func (o *HistoryServiceLevelObjectiveResponseData) GetTypeId() ServiceLevelObjectiveTypeNumeric`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *HistoryServiceLevelObjectiveResponseData) GetTypeIdOk() (ServiceLevelObjectiveTypeNumeric, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTypeId

`func (o *HistoryServiceLevelObjectiveResponseData) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### SetTypeId

`func (o *HistoryServiceLevelObjectiveResponseData) SetTypeId(v ServiceLevelObjectiveTypeNumeric)`

SetTypeId gets a reference to the given ServiceLevelObjectiveTypeNumeric and assigns it to the TypeId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


