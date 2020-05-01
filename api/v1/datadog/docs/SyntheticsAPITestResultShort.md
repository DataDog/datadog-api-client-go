# SyntheticsAPITestResultShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckTime** | Pointer to **float64** | Last time the API test was performed. | [optional] 
**ProbeDc** | Pointer to **string** | Location from which the API test was performed. | [optional] 
**Result** | Pointer to [**SyntheticsAPITestResultShortResult**](SyntheticsAPITestResultShort_result.md) |  | [optional] 
**ResultId** | Pointer to **string** | ID of the API test result. | [optional] 
**Status** | Pointer to [**SyntheticsTestMonitorStatus**](SyntheticsTestMonitorStatus.md) |  | [optional] 

## Methods

### NewSyntheticsAPITestResultShort

`func NewSyntheticsAPITestResultShort() *SyntheticsAPITestResultShort`

NewSyntheticsAPITestResultShort instantiates a new SyntheticsAPITestResultShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsAPITestResultShortWithDefaults

`func NewSyntheticsAPITestResultShortWithDefaults() *SyntheticsAPITestResultShort`

NewSyntheticsAPITestResultShortWithDefaults instantiates a new SyntheticsAPITestResultShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckTime

`func (o *SyntheticsAPITestResultShort) GetCheckTime() float64`

GetCheckTime returns the CheckTime field if non-nil, zero value otherwise.

### GetCheckTimeOk

`func (o *SyntheticsAPITestResultShort) GetCheckTimeOk() (*float64, bool)`

GetCheckTimeOk returns a tuple with the CheckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTime

`func (o *SyntheticsAPITestResultShort) SetCheckTime(v float64)`

SetCheckTime sets CheckTime field to given value.

### HasCheckTime

`func (o *SyntheticsAPITestResultShort) HasCheckTime() bool`

HasCheckTime returns a boolean if a field has been set.

### GetProbeDc

`func (o *SyntheticsAPITestResultShort) GetProbeDc() string`

GetProbeDc returns the ProbeDc field if non-nil, zero value otherwise.

### GetProbeDcOk

`func (o *SyntheticsAPITestResultShort) GetProbeDcOk() (*string, bool)`

GetProbeDcOk returns a tuple with the ProbeDc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeDc

`func (o *SyntheticsAPITestResultShort) SetProbeDc(v string)`

SetProbeDc sets ProbeDc field to given value.

### HasProbeDc

`func (o *SyntheticsAPITestResultShort) HasProbeDc() bool`

HasProbeDc returns a boolean if a field has been set.

### GetResult

`func (o *SyntheticsAPITestResultShort) GetResult() SyntheticsAPITestResultShortResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SyntheticsAPITestResultShort) GetResultOk() (*SyntheticsAPITestResultShortResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SyntheticsAPITestResultShort) SetResult(v SyntheticsAPITestResultShortResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *SyntheticsAPITestResultShort) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetResultId

`func (o *SyntheticsAPITestResultShort) GetResultId() string`

GetResultId returns the ResultId field if non-nil, zero value otherwise.

### GetResultIdOk

`func (o *SyntheticsAPITestResultShort) GetResultIdOk() (*string, bool)`

GetResultIdOk returns a tuple with the ResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultId

`func (o *SyntheticsAPITestResultShort) SetResultId(v string)`

SetResultId sets ResultId field to given value.

### HasResultId

`func (o *SyntheticsAPITestResultShort) HasResultId() bool`

HasResultId returns a boolean if a field has been set.

### GetStatus

`func (o *SyntheticsAPITestResultShort) GetStatus() SyntheticsTestMonitorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyntheticsAPITestResultShort) GetStatusOk() (*SyntheticsTestMonitorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyntheticsAPITestResultShort) SetStatus(v SyntheticsTestMonitorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyntheticsAPITestResultShort) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


