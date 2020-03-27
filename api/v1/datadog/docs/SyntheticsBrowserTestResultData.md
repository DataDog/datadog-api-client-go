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

### NewSyntheticsBrowserTestResultData

`func NewSyntheticsBrowserTestResultData() *SyntheticsBrowserTestResultData`

NewSyntheticsBrowserTestResultData instantiates a new SyntheticsBrowserTestResultData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsBrowserTestResultDataWithDefaults

`func NewSyntheticsBrowserTestResultDataWithDefaults() *SyntheticsBrowserTestResultData`

NewSyntheticsBrowserTestResultDataWithDefaults instantiates a new SyntheticsBrowserTestResultData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrowserType

`func (o *SyntheticsBrowserTestResultData) GetBrowserType() string`

GetBrowserType returns the BrowserType field if non-nil, zero value otherwise.

### GetBrowserTypeOk

`func (o *SyntheticsBrowserTestResultData) GetBrowserTypeOk() (*string, bool)`

GetBrowserTypeOk returns a tuple with the BrowserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserType

`func (o *SyntheticsBrowserTestResultData) SetBrowserType(v string)`

SetBrowserType sets BrowserType field to given value.

### HasBrowserType

`func (o *SyntheticsBrowserTestResultData) HasBrowserType() bool`

HasBrowserType returns a boolean if a field has been set.

### GetBrowserVersion

`func (o *SyntheticsBrowserTestResultData) GetBrowserVersion() string`

GetBrowserVersion returns the BrowserVersion field if non-nil, zero value otherwise.

### GetBrowserVersionOk

`func (o *SyntheticsBrowserTestResultData) GetBrowserVersionOk() (*string, bool)`

GetBrowserVersionOk returns a tuple with the BrowserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserVersion

`func (o *SyntheticsBrowserTestResultData) SetBrowserVersion(v string)`

SetBrowserVersion sets BrowserVersion field to given value.

### HasBrowserVersion

`func (o *SyntheticsBrowserTestResultData) HasBrowserVersion() bool`

HasBrowserVersion returns a boolean if a field has been set.

### GetDevice

`func (o *SyntheticsBrowserTestResultData) GetDevice() SyntheticsDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *SyntheticsBrowserTestResultData) GetDeviceOk() (*SyntheticsDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *SyntheticsBrowserTestResultData) SetDevice(v SyntheticsDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *SyntheticsBrowserTestResultData) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDuration

`func (o *SyntheticsBrowserTestResultData) GetDuration() float64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SyntheticsBrowserTestResultData) GetDurationOk() (*float64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SyntheticsBrowserTestResultData) SetDuration(v float64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SyntheticsBrowserTestResultData) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetError

`func (o *SyntheticsBrowserTestResultData) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SyntheticsBrowserTestResultData) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SyntheticsBrowserTestResultData) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SyntheticsBrowserTestResultData) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPassed

`func (o *SyntheticsBrowserTestResultData) GetPassed() bool`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *SyntheticsBrowserTestResultData) GetPassedOk() (*bool, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *SyntheticsBrowserTestResultData) SetPassed(v bool)`

SetPassed sets Passed field to given value.

### HasPassed

`func (o *SyntheticsBrowserTestResultData) HasPassed() bool`

HasPassed returns a boolean if a field has been set.

### GetReceivedEmailCount

`func (o *SyntheticsBrowserTestResultData) GetReceivedEmailCount() int64`

GetReceivedEmailCount returns the ReceivedEmailCount field if non-nil, zero value otherwise.

### GetReceivedEmailCountOk

`func (o *SyntheticsBrowserTestResultData) GetReceivedEmailCountOk() (*int64, bool)`

GetReceivedEmailCountOk returns a tuple with the ReceivedEmailCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedEmailCount

`func (o *SyntheticsBrowserTestResultData) SetReceivedEmailCount(v int64)`

SetReceivedEmailCount sets ReceivedEmailCount field to given value.

### HasReceivedEmailCount

`func (o *SyntheticsBrowserTestResultData) HasReceivedEmailCount() bool`

HasReceivedEmailCount returns a boolean if a field has been set.

### GetStartUrl

`func (o *SyntheticsBrowserTestResultData) GetStartUrl() string`

GetStartUrl returns the StartUrl field if non-nil, zero value otherwise.

### GetStartUrlOk

`func (o *SyntheticsBrowserTestResultData) GetStartUrlOk() (*string, bool)`

GetStartUrlOk returns a tuple with the StartUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartUrl

`func (o *SyntheticsBrowserTestResultData) SetStartUrl(v string)`

SetStartUrl sets StartUrl field to given value.

### HasStartUrl

`func (o *SyntheticsBrowserTestResultData) HasStartUrl() bool`

HasStartUrl returns a boolean if a field has been set.

### GetStepDetails

`func (o *SyntheticsBrowserTestResultData) GetStepDetails() []SyntheticsStepDetail`

GetStepDetails returns the StepDetails field if non-nil, zero value otherwise.

### GetStepDetailsOk

`func (o *SyntheticsBrowserTestResultData) GetStepDetailsOk() (*[]SyntheticsStepDetail, bool)`

GetStepDetailsOk returns a tuple with the StepDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepDetails

`func (o *SyntheticsBrowserTestResultData) SetStepDetails(v []SyntheticsStepDetail)`

SetStepDetails sets StepDetails field to given value.

### HasStepDetails

`func (o *SyntheticsBrowserTestResultData) HasStepDetails() bool`

HasStepDetails returns a boolean if a field has been set.

### GetThumbnailsBucketKey

`func (o *SyntheticsBrowserTestResultData) GetThumbnailsBucketKey() bool`

GetThumbnailsBucketKey returns the ThumbnailsBucketKey field if non-nil, zero value otherwise.

### GetThumbnailsBucketKeyOk

`func (o *SyntheticsBrowserTestResultData) GetThumbnailsBucketKeyOk() (*bool, bool)`

GetThumbnailsBucketKeyOk returns a tuple with the ThumbnailsBucketKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailsBucketKey

`func (o *SyntheticsBrowserTestResultData) SetThumbnailsBucketKey(v bool)`

SetThumbnailsBucketKey sets ThumbnailsBucketKey field to given value.

### HasThumbnailsBucketKey

`func (o *SyntheticsBrowserTestResultData) HasThumbnailsBucketKey() bool`

HasThumbnailsBucketKey returns a boolean if a field has been set.

### GetTimeToInteractive

`func (o *SyntheticsBrowserTestResultData) GetTimeToInteractive() float64`

GetTimeToInteractive returns the TimeToInteractive field if non-nil, zero value otherwise.

### GetTimeToInteractiveOk

`func (o *SyntheticsBrowserTestResultData) GetTimeToInteractiveOk() (*float64, bool)`

GetTimeToInteractiveOk returns a tuple with the TimeToInteractive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToInteractive

`func (o *SyntheticsBrowserTestResultData) SetTimeToInteractive(v float64)`

SetTimeToInteractive sets TimeToInteractive field to given value.

### HasTimeToInteractive

`func (o *SyntheticsBrowserTestResultData) HasTimeToInteractive() bool`

HasTimeToInteractive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


