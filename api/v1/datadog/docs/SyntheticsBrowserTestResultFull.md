# SyntheticsBrowserTestResultFull

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Check** | Pointer to [**SyntheticsBrowserTestResultFullCheck**](SyntheticsBrowserTestResultFullCheck.md) |  | [optional] 
**CheckTime** | Pointer to **float64** | When the browser test was conducted. | [optional] 
**CheckVersion** | Pointer to **int64** | Version of the browser test used. | [optional] 
**ProbeDc** | Pointer to **string** | Location from which the browser test was performed. | [optional] 
**Result** | Pointer to [**SyntheticsBrowserTestResultData**](SyntheticsBrowserTestResultData.md) |  | [optional] 
**ResultId** | Pointer to **string** | ID of the browser test result. | [optional] 
**Status** | Pointer to [**SyntheticsTestMonitorStatus**](SyntheticsTestMonitorStatus.md) |  | [optional] 

## Methods

### NewSyntheticsBrowserTestResultFull

`func NewSyntheticsBrowserTestResultFull() *SyntheticsBrowserTestResultFull`

NewSyntheticsBrowserTestResultFull instantiates a new SyntheticsBrowserTestResultFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsBrowserTestResultFullWithDefaults

`func NewSyntheticsBrowserTestResultFullWithDefaults() *SyntheticsBrowserTestResultFull`

NewSyntheticsBrowserTestResultFullWithDefaults instantiates a new SyntheticsBrowserTestResultFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheck

`func (o *SyntheticsBrowserTestResultFull) GetCheck() SyntheticsBrowserTestResultFullCheck`

GetCheck returns the Check field if non-nil, zero value otherwise.

### GetCheckOk

`func (o *SyntheticsBrowserTestResultFull) GetCheckOk() (*SyntheticsBrowserTestResultFullCheck, bool)`

GetCheckOk returns a tuple with the Check field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheck

`func (o *SyntheticsBrowserTestResultFull) SetCheck(v SyntheticsBrowserTestResultFullCheck)`

SetCheck sets Check field to given value.

### HasCheck

`func (o *SyntheticsBrowserTestResultFull) HasCheck() bool`

HasCheck returns a boolean if a field has been set.

### GetCheckTime

`func (o *SyntheticsBrowserTestResultFull) GetCheckTime() float64`

GetCheckTime returns the CheckTime field if non-nil, zero value otherwise.

### GetCheckTimeOk

`func (o *SyntheticsBrowserTestResultFull) GetCheckTimeOk() (*float64, bool)`

GetCheckTimeOk returns a tuple with the CheckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTime

`func (o *SyntheticsBrowserTestResultFull) SetCheckTime(v float64)`

SetCheckTime sets CheckTime field to given value.

### HasCheckTime

`func (o *SyntheticsBrowserTestResultFull) HasCheckTime() bool`

HasCheckTime returns a boolean if a field has been set.

### GetCheckVersion

`func (o *SyntheticsBrowserTestResultFull) GetCheckVersion() int64`

GetCheckVersion returns the CheckVersion field if non-nil, zero value otherwise.

### GetCheckVersionOk

`func (o *SyntheticsBrowserTestResultFull) GetCheckVersionOk() (*int64, bool)`

GetCheckVersionOk returns a tuple with the CheckVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckVersion

`func (o *SyntheticsBrowserTestResultFull) SetCheckVersion(v int64)`

SetCheckVersion sets CheckVersion field to given value.

### HasCheckVersion

`func (o *SyntheticsBrowserTestResultFull) HasCheckVersion() bool`

HasCheckVersion returns a boolean if a field has been set.

### GetProbeDc

`func (o *SyntheticsBrowserTestResultFull) GetProbeDc() string`

GetProbeDc returns the ProbeDc field if non-nil, zero value otherwise.

### GetProbeDcOk

`func (o *SyntheticsBrowserTestResultFull) GetProbeDcOk() (*string, bool)`

GetProbeDcOk returns a tuple with the ProbeDc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeDc

`func (o *SyntheticsBrowserTestResultFull) SetProbeDc(v string)`

SetProbeDc sets ProbeDc field to given value.

### HasProbeDc

`func (o *SyntheticsBrowserTestResultFull) HasProbeDc() bool`

HasProbeDc returns a boolean if a field has been set.

### GetResult

`func (o *SyntheticsBrowserTestResultFull) GetResult() SyntheticsBrowserTestResultData`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SyntheticsBrowserTestResultFull) GetResultOk() (*SyntheticsBrowserTestResultData, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SyntheticsBrowserTestResultFull) SetResult(v SyntheticsBrowserTestResultData)`

SetResult sets Result field to given value.

### HasResult

`func (o *SyntheticsBrowserTestResultFull) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetResultId

`func (o *SyntheticsBrowserTestResultFull) GetResultId() string`

GetResultId returns the ResultId field if non-nil, zero value otherwise.

### GetResultIdOk

`func (o *SyntheticsBrowserTestResultFull) GetResultIdOk() (*string, bool)`

GetResultIdOk returns a tuple with the ResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultId

`func (o *SyntheticsBrowserTestResultFull) SetResultId(v string)`

SetResultId sets ResultId field to given value.

### HasResultId

`func (o *SyntheticsBrowserTestResultFull) HasResultId() bool`

HasResultId returns a boolean if a field has been set.

### GetStatus

`func (o *SyntheticsBrowserTestResultFull) GetStatus() SyntheticsTestMonitorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyntheticsBrowserTestResultFull) GetStatusOk() (*SyntheticsTestMonitorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyntheticsBrowserTestResultFull) SetStatus(v SyntheticsTestMonitorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyntheticsBrowserTestResultFull) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


