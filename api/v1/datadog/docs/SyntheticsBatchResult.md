# SyntheticsBatchResult

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Device** | Pointer to [**SyntheticsDeviceID**](SyntheticsDeviceID.md) |  | [optional] 
**Duration** | Pointer to **float64** | Total duration in millisecond of the test. | [optional] 
**ExecutionRule** | Pointer to [**SyntheticsTestExecutionRule**](SyntheticsTestExecutionRule.md) |  | [optional] 
**Location** | Pointer to **string** | Name of the location. | [optional] 
**ResultId** | Pointer to **string** | The ID of the result to get. | [optional] 
**Retries** | Pointer to **float64** | Total duration in millisecond of the test. | [optional] 
**Status** | Pointer to [**SyntheticsStatus**](SyntheticsStatus.md) |  | [optional] 
**TestName** | Pointer to **string** | Name of the test. | [optional] 
**TestPublicId** | Pointer to **string** | The public ID of the Synthetic test. | [optional] 
**TestType** | Pointer to [**SyntheticsTestDetailsType**](SyntheticsTestDetailsType.md) |  | [optional] 

## Methods

### NewSyntheticsBatchResult

`func NewSyntheticsBatchResult() *SyntheticsBatchResult`

NewSyntheticsBatchResult instantiates a new SyntheticsBatchResult object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsBatchResultWithDefaults

`func NewSyntheticsBatchResultWithDefaults() *SyntheticsBatchResult`

NewSyntheticsBatchResultWithDefaults instantiates a new SyntheticsBatchResult object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDevice

`func (o *SyntheticsBatchResult) GetDevice() SyntheticsDeviceID`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *SyntheticsBatchResult) GetDeviceOk() (*SyntheticsDeviceID, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *SyntheticsBatchResult) SetDevice(v SyntheticsDeviceID)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *SyntheticsBatchResult) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDuration

`func (o *SyntheticsBatchResult) GetDuration() float64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SyntheticsBatchResult) GetDurationOk() (*float64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SyntheticsBatchResult) SetDuration(v float64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SyntheticsBatchResult) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetExecutionRule

`func (o *SyntheticsBatchResult) GetExecutionRule() SyntheticsTestExecutionRule`

GetExecutionRule returns the ExecutionRule field if non-nil, zero value otherwise.

### GetExecutionRuleOk

`func (o *SyntheticsBatchResult) GetExecutionRuleOk() (*SyntheticsTestExecutionRule, bool)`

GetExecutionRuleOk returns a tuple with the ExecutionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionRule

`func (o *SyntheticsBatchResult) SetExecutionRule(v SyntheticsTestExecutionRule)`

SetExecutionRule sets ExecutionRule field to given value.

### HasExecutionRule

`func (o *SyntheticsBatchResult) HasExecutionRule() bool`

HasExecutionRule returns a boolean if a field has been set.

### GetLocation

`func (o *SyntheticsBatchResult) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SyntheticsBatchResult) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SyntheticsBatchResult) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SyntheticsBatchResult) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetResultId

`func (o *SyntheticsBatchResult) GetResultId() string`

GetResultId returns the ResultId field if non-nil, zero value otherwise.

### GetResultIdOk

`func (o *SyntheticsBatchResult) GetResultIdOk() (*string, bool)`

GetResultIdOk returns a tuple with the ResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultId

`func (o *SyntheticsBatchResult) SetResultId(v string)`

SetResultId sets ResultId field to given value.

### HasResultId

`func (o *SyntheticsBatchResult) HasResultId() bool`

HasResultId returns a boolean if a field has been set.

### GetRetries

`func (o *SyntheticsBatchResult) GetRetries() float64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *SyntheticsBatchResult) GetRetriesOk() (*float64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *SyntheticsBatchResult) SetRetries(v float64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *SyntheticsBatchResult) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetStatus

`func (o *SyntheticsBatchResult) GetStatus() SyntheticsStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyntheticsBatchResult) GetStatusOk() (*SyntheticsStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyntheticsBatchResult) SetStatus(v SyntheticsStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyntheticsBatchResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTestName

`func (o *SyntheticsBatchResult) GetTestName() string`

GetTestName returns the TestName field if non-nil, zero value otherwise.

### GetTestNameOk

`func (o *SyntheticsBatchResult) GetTestNameOk() (*string, bool)`

GetTestNameOk returns a tuple with the TestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestName

`func (o *SyntheticsBatchResult) SetTestName(v string)`

SetTestName sets TestName field to given value.

### HasTestName

`func (o *SyntheticsBatchResult) HasTestName() bool`

HasTestName returns a boolean if a field has been set.

### GetTestPublicId

`func (o *SyntheticsBatchResult) GetTestPublicId() string`

GetTestPublicId returns the TestPublicId field if non-nil, zero value otherwise.

### GetTestPublicIdOk

`func (o *SyntheticsBatchResult) GetTestPublicIdOk() (*string, bool)`

GetTestPublicIdOk returns a tuple with the TestPublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPublicId

`func (o *SyntheticsBatchResult) SetTestPublicId(v string)`

SetTestPublicId sets TestPublicId field to given value.

### HasTestPublicId

`func (o *SyntheticsBatchResult) HasTestPublicId() bool`

HasTestPublicId returns a boolean if a field has been set.

### GetTestType

`func (o *SyntheticsBatchResult) GetTestType() SyntheticsTestDetailsType`

GetTestType returns the TestType field if non-nil, zero value otherwise.

### GetTestTypeOk

`func (o *SyntheticsBatchResult) GetTestTypeOk() (*SyntheticsTestDetailsType, bool)`

GetTestTypeOk returns a tuple with the TestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestType

`func (o *SyntheticsBatchResult) SetTestType(v SyntheticsTestDetailsType)`

SetTestType sets TestType field to given value.

### HasTestType

`func (o *SyntheticsBatchResult) HasTestType() bool`

HasTestType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


