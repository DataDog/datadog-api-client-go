# SyntheticsAPITestResultFull

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Check** | Pointer to [**SyntheticsAPITestResultFullCheck**](SyntheticsAPITestResultFullCheck.md) |  | [optional] 
**CheckTime** | Pointer to **float64** | When the API test was conducted. | [optional] 
**CheckVersion** | Pointer to **int64** | Version of the API test used. | [optional] 
**ProbeDc** | Pointer to **string** | Locations for which to query the API test results. | [optional] 
**Result** | Pointer to [**SyntheticsAPITestResultData**](SyntheticsAPITestResultData.md) |  | [optional] 
**ResultId** | Pointer to **string** | ID of the API test result. | [optional] 
**Status** | Pointer to [**SyntheticsTestMonitorStatus**](SyntheticsTestMonitorStatus.md) |  | [optional] 

## Methods

### NewSyntheticsAPITestResultFull

`func NewSyntheticsAPITestResultFull() *SyntheticsAPITestResultFull`

NewSyntheticsAPITestResultFull instantiates a new SyntheticsAPITestResultFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsAPITestResultFullWithDefaults

`func NewSyntheticsAPITestResultFullWithDefaults() *SyntheticsAPITestResultFull`

NewSyntheticsAPITestResultFullWithDefaults instantiates a new SyntheticsAPITestResultFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheck

`func (o *SyntheticsAPITestResultFull) GetCheck() SyntheticsAPITestResultFullCheck`

GetCheck returns the Check field if non-nil, zero value otherwise.

### GetCheckOk

`func (o *SyntheticsAPITestResultFull) GetCheckOk() (*SyntheticsAPITestResultFullCheck, bool)`

GetCheckOk returns a tuple with the Check field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheck

`func (o *SyntheticsAPITestResultFull) SetCheck(v SyntheticsAPITestResultFullCheck)`

SetCheck sets Check field to given value.

### HasCheck

`func (o *SyntheticsAPITestResultFull) HasCheck() bool`

HasCheck returns a boolean if a field has been set.

### GetCheckTime

`func (o *SyntheticsAPITestResultFull) GetCheckTime() float64`

GetCheckTime returns the CheckTime field if non-nil, zero value otherwise.

### GetCheckTimeOk

`func (o *SyntheticsAPITestResultFull) GetCheckTimeOk() (*float64, bool)`

GetCheckTimeOk returns a tuple with the CheckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTime

`func (o *SyntheticsAPITestResultFull) SetCheckTime(v float64)`

SetCheckTime sets CheckTime field to given value.

### HasCheckTime

`func (o *SyntheticsAPITestResultFull) HasCheckTime() bool`

HasCheckTime returns a boolean if a field has been set.

### GetCheckVersion

`func (o *SyntheticsAPITestResultFull) GetCheckVersion() int64`

GetCheckVersion returns the CheckVersion field if non-nil, zero value otherwise.

### GetCheckVersionOk

`func (o *SyntheticsAPITestResultFull) GetCheckVersionOk() (*int64, bool)`

GetCheckVersionOk returns a tuple with the CheckVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckVersion

`func (o *SyntheticsAPITestResultFull) SetCheckVersion(v int64)`

SetCheckVersion sets CheckVersion field to given value.

### HasCheckVersion

`func (o *SyntheticsAPITestResultFull) HasCheckVersion() bool`

HasCheckVersion returns a boolean if a field has been set.

### GetProbeDc

`func (o *SyntheticsAPITestResultFull) GetProbeDc() string`

GetProbeDc returns the ProbeDc field if non-nil, zero value otherwise.

### GetProbeDcOk

`func (o *SyntheticsAPITestResultFull) GetProbeDcOk() (*string, bool)`

GetProbeDcOk returns a tuple with the ProbeDc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeDc

`func (o *SyntheticsAPITestResultFull) SetProbeDc(v string)`

SetProbeDc sets ProbeDc field to given value.

### HasProbeDc

`func (o *SyntheticsAPITestResultFull) HasProbeDc() bool`

HasProbeDc returns a boolean if a field has been set.

### GetResult

`func (o *SyntheticsAPITestResultFull) GetResult() SyntheticsAPITestResultData`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SyntheticsAPITestResultFull) GetResultOk() (*SyntheticsAPITestResultData, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SyntheticsAPITestResultFull) SetResult(v SyntheticsAPITestResultData)`

SetResult sets Result field to given value.

### HasResult

`func (o *SyntheticsAPITestResultFull) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetResultId

`func (o *SyntheticsAPITestResultFull) GetResultId() string`

GetResultId returns the ResultId field if non-nil, zero value otherwise.

### GetResultIdOk

`func (o *SyntheticsAPITestResultFull) GetResultIdOk() (*string, bool)`

GetResultIdOk returns a tuple with the ResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultId

`func (o *SyntheticsAPITestResultFull) SetResultId(v string)`

SetResultId sets ResultId field to given value.

### HasResultId

`func (o *SyntheticsAPITestResultFull) HasResultId() bool`

HasResultId returns a boolean if a field has been set.

### GetStatus

`func (o *SyntheticsAPITestResultFull) GetStatus() SyntheticsTestMonitorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyntheticsAPITestResultFull) GetStatusOk() (*SyntheticsTestMonitorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyntheticsAPITestResultFull) SetStatus(v SyntheticsTestMonitorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyntheticsAPITestResultFull) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


