# SyntheticsBrowserTestResultData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrowserType** | Pointer to **string** |  | [optional] 
**BrowserVersion** | Pointer to **string** |  | [optional] 
**Device** | Pointer to [**SyntheticsDevice**](SyntheticsDevice.md) |  | [optional] 
**Duration** | Pointer to **float64** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Passed** | Pointer to **bool** |  | [optional] 
**ReceivedEmailCount** | Pointer to **int64** |  | [optional] 
**StartUrl** | Pointer to **string** |  | [optional] 
**StepDetails** | Pointer to [**[]SyntheticsStepDetail**](SyntheticsStepDetail.md) |  | [optional] 
**ThumbnailsBucketKey** | Pointer to **bool** |  | [optional] 
**TimeToInteractive** | Pointer to **float64** |  | [optional] 

## Methods

### GetBrowserType

`func (o *SyntheticsBrowserTestResultData) GetBrowserType() string`

GetBrowserType returns the BrowserType field if non-nil, zero value otherwise.

### GetBrowserTypeOk

`func (o *SyntheticsBrowserTestResultData) GetBrowserTypeOk() (string, bool)`

GetBrowserTypeOk returns a tuple with the BrowserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserType

`func (o *SyntheticsBrowserTestResultData) HasBrowserType() bool`

HasBrowserType returns a boolean if a field has been set.

### SetBrowserType

`func (o *SyntheticsBrowserTestResultData) SetBrowserType(v string)`

SetBrowserType gets a reference to the given string and assigns it to the BrowserType field.

### GetBrowserVersion

`func (o *SyntheticsBrowserTestResultData) GetBrowserVersion() string`

GetBrowserVersion returns the BrowserVersion field if non-nil, zero value otherwise.

### GetBrowserVersionOk

`func (o *SyntheticsBrowserTestResultData) GetBrowserVersionOk() (string, bool)`

GetBrowserVersionOk returns a tuple with the BrowserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserVersion

`func (o *SyntheticsBrowserTestResultData) HasBrowserVersion() bool`

HasBrowserVersion returns a boolean if a field has been set.

### SetBrowserVersion

`func (o *SyntheticsBrowserTestResultData) SetBrowserVersion(v string)`

SetBrowserVersion gets a reference to the given string and assigns it to the BrowserVersion field.

### GetDevice

`func (o *SyntheticsBrowserTestResultData) GetDevice() SyntheticsDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *SyntheticsBrowserTestResultData) GetDeviceOk() (SyntheticsDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDevice

`func (o *SyntheticsBrowserTestResultData) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDevice

`func (o *SyntheticsBrowserTestResultData) SetDevice(v SyntheticsDevice)`

SetDevice gets a reference to the given SyntheticsDevice and assigns it to the Device field.

### GetDuration

`func (o *SyntheticsBrowserTestResultData) GetDuration() float64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SyntheticsBrowserTestResultData) GetDurationOk() (float64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDuration

`func (o *SyntheticsBrowserTestResultData) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDuration

`func (o *SyntheticsBrowserTestResultData) SetDuration(v float64)`

SetDuration gets a reference to the given float64 and assigns it to the Duration field.

### GetError

`func (o *SyntheticsBrowserTestResultData) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SyntheticsBrowserTestResultData) GetErrorOk() (string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasError

`func (o *SyntheticsBrowserTestResultData) HasError() bool`

HasError returns a boolean if a field has been set.

### SetError

`func (o *SyntheticsBrowserTestResultData) SetError(v string)`

SetError gets a reference to the given string and assigns it to the Error field.

### GetPassed

`func (o *SyntheticsBrowserTestResultData) GetPassed() bool`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *SyntheticsBrowserTestResultData) GetPassedOk() (bool, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPassed

`func (o *SyntheticsBrowserTestResultData) HasPassed() bool`

HasPassed returns a boolean if a field has been set.

### SetPassed

`func (o *SyntheticsBrowserTestResultData) SetPassed(v bool)`

SetPassed gets a reference to the given bool and assigns it to the Passed field.

### GetReceivedEmailCount

`func (o *SyntheticsBrowserTestResultData) GetReceivedEmailCount() int64`

GetReceivedEmailCount returns the ReceivedEmailCount field if non-nil, zero value otherwise.

### GetReceivedEmailCountOk

`func (o *SyntheticsBrowserTestResultData) GetReceivedEmailCountOk() (int64, bool)`

GetReceivedEmailCountOk returns a tuple with the ReceivedEmailCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReceivedEmailCount

`func (o *SyntheticsBrowserTestResultData) HasReceivedEmailCount() bool`

HasReceivedEmailCount returns a boolean if a field has been set.

### SetReceivedEmailCount

`func (o *SyntheticsBrowserTestResultData) SetReceivedEmailCount(v int64)`

SetReceivedEmailCount gets a reference to the given int64 and assigns it to the ReceivedEmailCount field.

### GetStartUrl

`func (o *SyntheticsBrowserTestResultData) GetStartUrl() string`

GetStartUrl returns the StartUrl field if non-nil, zero value otherwise.

### GetStartUrlOk

`func (o *SyntheticsBrowserTestResultData) GetStartUrlOk() (string, bool)`

GetStartUrlOk returns a tuple with the StartUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartUrl

`func (o *SyntheticsBrowserTestResultData) HasStartUrl() bool`

HasStartUrl returns a boolean if a field has been set.

### SetStartUrl

`func (o *SyntheticsBrowserTestResultData) SetStartUrl(v string)`

SetStartUrl gets a reference to the given string and assigns it to the StartUrl field.

### GetStepDetails

`func (o *SyntheticsBrowserTestResultData) GetStepDetails() []SyntheticsStepDetail`

GetStepDetails returns the StepDetails field if non-nil, zero value otherwise.

### GetStepDetailsOk

`func (o *SyntheticsBrowserTestResultData) GetStepDetailsOk() ([]SyntheticsStepDetail, bool)`

GetStepDetailsOk returns a tuple with the StepDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStepDetails

`func (o *SyntheticsBrowserTestResultData) HasStepDetails() bool`

HasStepDetails returns a boolean if a field has been set.

### SetStepDetails

`func (o *SyntheticsBrowserTestResultData) SetStepDetails(v []SyntheticsStepDetail)`

SetStepDetails gets a reference to the given []SyntheticsStepDetail and assigns it to the StepDetails field.

### GetThumbnailsBucketKey

`func (o *SyntheticsBrowserTestResultData) GetThumbnailsBucketKey() bool`

GetThumbnailsBucketKey returns the ThumbnailsBucketKey field if non-nil, zero value otherwise.

### GetThumbnailsBucketKeyOk

`func (o *SyntheticsBrowserTestResultData) GetThumbnailsBucketKeyOk() (bool, bool)`

GetThumbnailsBucketKeyOk returns a tuple with the ThumbnailsBucketKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasThumbnailsBucketKey

`func (o *SyntheticsBrowserTestResultData) HasThumbnailsBucketKey() bool`

HasThumbnailsBucketKey returns a boolean if a field has been set.

### SetThumbnailsBucketKey

`func (o *SyntheticsBrowserTestResultData) SetThumbnailsBucketKey(v bool)`

SetThumbnailsBucketKey gets a reference to the given bool and assigns it to the ThumbnailsBucketKey field.

### GetTimeToInteractive

`func (o *SyntheticsBrowserTestResultData) GetTimeToInteractive() float64`

GetTimeToInteractive returns the TimeToInteractive field if non-nil, zero value otherwise.

### GetTimeToInteractiveOk

`func (o *SyntheticsBrowserTestResultData) GetTimeToInteractiveOk() (float64, bool)`

GetTimeToInteractiveOk returns a tuple with the TimeToInteractive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimeToInteractive

`func (o *SyntheticsBrowserTestResultData) HasTimeToInteractive() bool`

HasTimeToInteractive returns a boolean if a field has been set.

### SetTimeToInteractive

`func (o *SyntheticsBrowserTestResultData) SetTimeToInteractive(v float64)`

SetTimeToInteractive gets a reference to the given float64 and assigns it to the TimeToInteractive field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


