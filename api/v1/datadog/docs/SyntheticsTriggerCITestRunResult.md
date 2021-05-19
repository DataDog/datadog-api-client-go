# SyntheticsTriggerCITestRunResult

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Device** | Pointer to [**SyntheticsDeviceID**](SyntheticsDeviceID.md) |  | [optional] 
**Location** | Pointer to **int64** | The location ID of the test run. | [optional] 
**PublicId** | Pointer to **string** | The public ID of the Synthetics test. | [optional] 
**ResultId** | Pointer to **string** | ID of the result. | [optional] 

## Methods

### NewSyntheticsTriggerCITestRunResult

`func NewSyntheticsTriggerCITestRunResult() *SyntheticsTriggerCITestRunResult`

NewSyntheticsTriggerCITestRunResult instantiates a new SyntheticsTriggerCITestRunResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsTriggerCITestRunResultWithDefaults

`func NewSyntheticsTriggerCITestRunResultWithDefaults() *SyntheticsTriggerCITestRunResult`

NewSyntheticsTriggerCITestRunResultWithDefaults instantiates a new SyntheticsTriggerCITestRunResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *SyntheticsTriggerCITestRunResult) GetDevice() SyntheticsDeviceID`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *SyntheticsTriggerCITestRunResult) GetDeviceOk() (*SyntheticsDeviceID, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *SyntheticsTriggerCITestRunResult) SetDevice(v SyntheticsDeviceID)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *SyntheticsTriggerCITestRunResult) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetLocation

`func (o *SyntheticsTriggerCITestRunResult) GetLocation() int64`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SyntheticsTriggerCITestRunResult) GetLocationOk() (*int64, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SyntheticsTriggerCITestRunResult) SetLocation(v int64)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SyntheticsTriggerCITestRunResult) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPublicId

`func (o *SyntheticsTriggerCITestRunResult) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *SyntheticsTriggerCITestRunResult) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *SyntheticsTriggerCITestRunResult) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *SyntheticsTriggerCITestRunResult) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetResultId

`func (o *SyntheticsTriggerCITestRunResult) GetResultId() string`

GetResultId returns the ResultId field if non-nil, zero value otherwise.

### GetResultIdOk

`func (o *SyntheticsTriggerCITestRunResult) GetResultIdOk() (*string, bool)`

GetResultIdOk returns a tuple with the ResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultId

`func (o *SyntheticsTriggerCITestRunResult) SetResultId(v string)`

SetResultId sets ResultId field to given value.

### HasResultId

`func (o *SyntheticsTriggerCITestRunResult) HasResultId() bool`

HasResultId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


