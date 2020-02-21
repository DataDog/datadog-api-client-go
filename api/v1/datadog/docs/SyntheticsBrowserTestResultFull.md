# SyntheticsBrowserTestResultFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Check** | Pointer to [**SyntheticsAPITestResultFullCheck**](SyntheticsAPITestResultFull_check.md) |  | [optional] 
**CheckTime** | Pointer to **float64** |  | [optional] 
**CheckVersion** | Pointer to **int64** |  | [optional] 
**ProbeDc** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**SyntheticsBrowserTestResultData**](SyntheticsBrowserTestResultData.md) |  | [optional] 
**ResultId** | Pointer to **string** |  | [optional] 
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

`func (o *SyntheticsBrowserTestResultFull) GetCheck() SyntheticsAPITestResultFullCheck`

GetCheck returns the Check field if non-nil, zero value otherwise.

### GetCheckOk

`func (o *SyntheticsBrowserTestResultFull) GetCheckOk() (SyntheticsAPITestResultFullCheck, bool)`

GetCheckOk returns a tuple with the Check field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCheck

`func (o *SyntheticsBrowserTestResultFull) HasCheck() bool`

HasCheck returns a boolean if a field has been set.

### SetCheck

`func (o *SyntheticsBrowserTestResultFull) SetCheck(v SyntheticsAPITestResultFullCheck)`

SetCheck gets a reference to the given SyntheticsAPITestResultFullCheck and assigns it to the Check field.

### GetCheckTime

`func (o *SyntheticsBrowserTestResultFull) GetCheckTime() float64`

GetCheckTime returns the CheckTime field if non-nil, zero value otherwise.

### GetCheckTimeOk

`func (o *SyntheticsBrowserTestResultFull) GetCheckTimeOk() (float64, bool)`

GetCheckTimeOk returns a tuple with the CheckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCheckTime

`func (o *SyntheticsBrowserTestResultFull) HasCheckTime() bool`

HasCheckTime returns a boolean if a field has been set.

### SetCheckTime

`func (o *SyntheticsBrowserTestResultFull) SetCheckTime(v float64)`

SetCheckTime gets a reference to the given float64 and assigns it to the CheckTime field.

### GetCheckVersion

`func (o *SyntheticsBrowserTestResultFull) GetCheckVersion() int64`

GetCheckVersion returns the CheckVersion field if non-nil, zero value otherwise.

### GetCheckVersionOk

`func (o *SyntheticsBrowserTestResultFull) GetCheckVersionOk() (int64, bool)`

GetCheckVersionOk returns a tuple with the CheckVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCheckVersion

`func (o *SyntheticsBrowserTestResultFull) HasCheckVersion() bool`

HasCheckVersion returns a boolean if a field has been set.

### SetCheckVersion

`func (o *SyntheticsBrowserTestResultFull) SetCheckVersion(v int64)`

SetCheckVersion gets a reference to the given int64 and assigns it to the CheckVersion field.

### GetProbeDc

`func (o *SyntheticsBrowserTestResultFull) GetProbeDc() string`

GetProbeDc returns the ProbeDc field if non-nil, zero value otherwise.

### GetProbeDcOk

`func (o *SyntheticsBrowserTestResultFull) GetProbeDcOk() (string, bool)`

GetProbeDcOk returns a tuple with the ProbeDc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProbeDc

`func (o *SyntheticsBrowserTestResultFull) HasProbeDc() bool`

HasProbeDc returns a boolean if a field has been set.

### SetProbeDc

`func (o *SyntheticsBrowserTestResultFull) SetProbeDc(v string)`

SetProbeDc gets a reference to the given string and assigns it to the ProbeDc field.

### GetResult

`func (o *SyntheticsBrowserTestResultFull) GetResult() SyntheticsBrowserTestResultData`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SyntheticsBrowserTestResultFull) GetResultOk() (SyntheticsBrowserTestResultData, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResult

`func (o *SyntheticsBrowserTestResultFull) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResult

`func (o *SyntheticsBrowserTestResultFull) SetResult(v SyntheticsBrowserTestResultData)`

SetResult gets a reference to the given SyntheticsBrowserTestResultData and assigns it to the Result field.

### GetResultId

`func (o *SyntheticsBrowserTestResultFull) GetResultId() string`

GetResultId returns the ResultId field if non-nil, zero value otherwise.

### GetResultIdOk

`func (o *SyntheticsBrowserTestResultFull) GetResultIdOk() (string, bool)`

GetResultIdOk returns a tuple with the ResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResultId

`func (o *SyntheticsBrowserTestResultFull) HasResultId() bool`

HasResultId returns a boolean if a field has been set.

### SetResultId

`func (o *SyntheticsBrowserTestResultFull) SetResultId(v string)`

SetResultId gets a reference to the given string and assigns it to the ResultId field.

### GetStatus

`func (o *SyntheticsBrowserTestResultFull) GetStatus() SyntheticsTestMonitorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyntheticsBrowserTestResultFull) GetStatusOk() (SyntheticsTestMonitorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *SyntheticsBrowserTestResultFull) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *SyntheticsBrowserTestResultFull) SetStatus(v SyntheticsTestMonitorStatus)`

SetStatus gets a reference to the given SyntheticsTestMonitorStatus and assigns it to the Status field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


