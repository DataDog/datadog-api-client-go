# SyntheticsAPITestResultShortResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passed** | Pointer to **bool** | Describes if the test run has passed or failed. | [optional] 
**Timings** | Pointer to [**SyntheticsTiming**](SyntheticsTiming.md) |  | [optional] 

## Methods

### NewSyntheticsAPITestResultShortResult

`func NewSyntheticsAPITestResultShortResult() *SyntheticsAPITestResultShortResult`

NewSyntheticsAPITestResultShortResult instantiates a new SyntheticsAPITestResultShortResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsAPITestResultShortResultWithDefaults

`func NewSyntheticsAPITestResultShortResultWithDefaults() *SyntheticsAPITestResultShortResult`

NewSyntheticsAPITestResultShortResultWithDefaults instantiates a new SyntheticsAPITestResultShortResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassed

`func (o *SyntheticsAPITestResultShortResult) GetPassed() bool`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *SyntheticsAPITestResultShortResult) GetPassedOk() (*bool, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *SyntheticsAPITestResultShortResult) SetPassed(v bool)`

SetPassed sets Passed field to given value.

### HasPassed

`func (o *SyntheticsAPITestResultShortResult) HasPassed() bool`

HasPassed returns a boolean if a field has been set.

### GetTimings

`func (o *SyntheticsAPITestResultShortResult) GetTimings() SyntheticsTiming`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *SyntheticsAPITestResultShortResult) GetTimingsOk() (*SyntheticsTiming, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimings

`func (o *SyntheticsAPITestResultShortResult) SetTimings(v SyntheticsTiming)`

SetTimings sets Timings field to given value.

### HasTimings

`func (o *SyntheticsAPITestResultShortResult) HasTimings() bool`

HasTimings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


