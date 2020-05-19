# SyntheticsStepDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrowserErrors** | Pointer to [**[]SyntheticsBrowserError**](SyntheticsBrowserError.md) | Array of errors collected for a browser test. | [optional] 
**CheckType** | Pointer to [**SyntheticsCheckType**](SyntheticsCheckType.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the test. | [optional] 
**Duration** | Pointer to **float64** | Total duration in millisecond of the test. | [optional] 
**Error** | Pointer to **string** | Error returned by the test. | [optional] 
**PlayingTab** | Pointer to [**SyntheticsPlayingTab**](SyntheticsPlayingTab.md) |  | [optional] 
**Resources** | Pointer to [**[]SyntheticsResource**](SyntheticsResource.md) | Array of resources collected by the test. | [optional] 
**ScreenshotBucketKey** | Pointer to **bool** | Whether or not screenshots where collected by the test. | [optional] 
**Skipped** | Pointer to **bool** | Whether or not to skip this step. | [optional] 
**SnapshotBucketKey** | Pointer to **bool** | Whether or not snapshots where collected by the test. | [optional] 
**StepId** | Pointer to **int64** | The step ID. | [optional] 
**SubTestStepDetails** | Pointer to [**[]SyntheticsStepDetail**](SyntheticsStepDetail.md) | If this steps include a sub-test. [Subtests documentation](https://docs.datadoghq.com/synthetics/browser_tests/advanced_options/#subtests). | [optional] 
**TimeToInteractive** | Pointer to **float64** | Time before starting the step. | [optional] 
**Type** | Pointer to [**SyntheticsStepType**](SyntheticsStepType.md) |  | [optional] 
**Url** | Pointer to **string** | URL to perform the step against. | [optional] 
**Value** | Pointer to **interface{}** | Value for the step. | [optional] 
**Warnings** | Pointer to [**[]SyntheticsStepDetailWarnings**](SyntheticsStepDetail_warnings.md) | Warning collected that didn&#39;t failed the step. | [optional] 

## Methods

### NewSyntheticsStepDetail

`func NewSyntheticsStepDetail() *SyntheticsStepDetail`

NewSyntheticsStepDetail instantiates a new SyntheticsStepDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsStepDetailWithDefaults

`func NewSyntheticsStepDetailWithDefaults() *SyntheticsStepDetail`

NewSyntheticsStepDetailWithDefaults instantiates a new SyntheticsStepDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrowserErrors

`func (o *SyntheticsStepDetail) GetBrowserErrors() []SyntheticsBrowserError`

GetBrowserErrors returns the BrowserErrors field if non-nil, zero value otherwise.

### GetBrowserErrorsOk

`func (o *SyntheticsStepDetail) GetBrowserErrorsOk() (*[]SyntheticsBrowserError, bool)`

GetBrowserErrorsOk returns a tuple with the BrowserErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserErrors

`func (o *SyntheticsStepDetail) SetBrowserErrors(v []SyntheticsBrowserError)`

SetBrowserErrors sets BrowserErrors field to given value.

### HasBrowserErrors

`func (o *SyntheticsStepDetail) HasBrowserErrors() bool`

HasBrowserErrors returns a boolean if a field has been set.

### GetCheckType

`func (o *SyntheticsStepDetail) GetCheckType() SyntheticsCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *SyntheticsStepDetail) GetCheckTypeOk() (*SyntheticsCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *SyntheticsStepDetail) SetCheckType(v SyntheticsCheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *SyntheticsStepDetail) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetDescription

`func (o *SyntheticsStepDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SyntheticsStepDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SyntheticsStepDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SyntheticsStepDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuration

`func (o *SyntheticsStepDetail) GetDuration() float64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SyntheticsStepDetail) GetDurationOk() (*float64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SyntheticsStepDetail) SetDuration(v float64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SyntheticsStepDetail) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetError

`func (o *SyntheticsStepDetail) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SyntheticsStepDetail) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SyntheticsStepDetail) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SyntheticsStepDetail) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPlayingTab

`func (o *SyntheticsStepDetail) GetPlayingTab() SyntheticsPlayingTab`

GetPlayingTab returns the PlayingTab field if non-nil, zero value otherwise.

### GetPlayingTabOk

`func (o *SyntheticsStepDetail) GetPlayingTabOk() (*SyntheticsPlayingTab, bool)`

GetPlayingTabOk returns a tuple with the PlayingTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayingTab

`func (o *SyntheticsStepDetail) SetPlayingTab(v SyntheticsPlayingTab)`

SetPlayingTab sets PlayingTab field to given value.

### HasPlayingTab

`func (o *SyntheticsStepDetail) HasPlayingTab() bool`

HasPlayingTab returns a boolean if a field has been set.

### GetResources

`func (o *SyntheticsStepDetail) GetResources() []SyntheticsResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *SyntheticsStepDetail) GetResourcesOk() (*[]SyntheticsResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *SyntheticsStepDetail) SetResources(v []SyntheticsResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *SyntheticsStepDetail) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetScreenshotBucketKey

`func (o *SyntheticsStepDetail) GetScreenshotBucketKey() bool`

GetScreenshotBucketKey returns the ScreenshotBucketKey field if non-nil, zero value otherwise.

### GetScreenshotBucketKeyOk

`func (o *SyntheticsStepDetail) GetScreenshotBucketKeyOk() (*bool, bool)`

GetScreenshotBucketKeyOk returns a tuple with the ScreenshotBucketKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshotBucketKey

`func (o *SyntheticsStepDetail) SetScreenshotBucketKey(v bool)`

SetScreenshotBucketKey sets ScreenshotBucketKey field to given value.

### HasScreenshotBucketKey

`func (o *SyntheticsStepDetail) HasScreenshotBucketKey() bool`

HasScreenshotBucketKey returns a boolean if a field has been set.

### GetSkipped

`func (o *SyntheticsStepDetail) GetSkipped() bool`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *SyntheticsStepDetail) GetSkippedOk() (*bool, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *SyntheticsStepDetail) SetSkipped(v bool)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *SyntheticsStepDetail) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetSnapshotBucketKey

`func (o *SyntheticsStepDetail) GetSnapshotBucketKey() bool`

GetSnapshotBucketKey returns the SnapshotBucketKey field if non-nil, zero value otherwise.

### GetSnapshotBucketKeyOk

`func (o *SyntheticsStepDetail) GetSnapshotBucketKeyOk() (*bool, bool)`

GetSnapshotBucketKeyOk returns a tuple with the SnapshotBucketKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotBucketKey

`func (o *SyntheticsStepDetail) SetSnapshotBucketKey(v bool)`

SetSnapshotBucketKey sets SnapshotBucketKey field to given value.

### HasSnapshotBucketKey

`func (o *SyntheticsStepDetail) HasSnapshotBucketKey() bool`

HasSnapshotBucketKey returns a boolean if a field has been set.

### GetStepId

`func (o *SyntheticsStepDetail) GetStepId() int64`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *SyntheticsStepDetail) GetStepIdOk() (*int64, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *SyntheticsStepDetail) SetStepId(v int64)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *SyntheticsStepDetail) HasStepId() bool`

HasStepId returns a boolean if a field has been set.

### GetSubTestStepDetails

`func (o *SyntheticsStepDetail) GetSubTestStepDetails() []SyntheticsStepDetail`

GetSubTestStepDetails returns the SubTestStepDetails field if non-nil, zero value otherwise.

### GetSubTestStepDetailsOk

`func (o *SyntheticsStepDetail) GetSubTestStepDetailsOk() (*[]SyntheticsStepDetail, bool)`

GetSubTestStepDetailsOk returns a tuple with the SubTestStepDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTestStepDetails

`func (o *SyntheticsStepDetail) SetSubTestStepDetails(v []SyntheticsStepDetail)`

SetSubTestStepDetails sets SubTestStepDetails field to given value.

### HasSubTestStepDetails

`func (o *SyntheticsStepDetail) HasSubTestStepDetails() bool`

HasSubTestStepDetails returns a boolean if a field has been set.

### GetTimeToInteractive

`func (o *SyntheticsStepDetail) GetTimeToInteractive() float64`

GetTimeToInteractive returns the TimeToInteractive field if non-nil, zero value otherwise.

### GetTimeToInteractiveOk

`func (o *SyntheticsStepDetail) GetTimeToInteractiveOk() (*float64, bool)`

GetTimeToInteractiveOk returns a tuple with the TimeToInteractive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToInteractive

`func (o *SyntheticsStepDetail) SetTimeToInteractive(v float64)`

SetTimeToInteractive sets TimeToInteractive field to given value.

### HasTimeToInteractive

`func (o *SyntheticsStepDetail) HasTimeToInteractive() bool`

HasTimeToInteractive returns a boolean if a field has been set.

### GetType

`func (o *SyntheticsStepDetail) GetType() SyntheticsStepType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsStepDetail) GetTypeOk() (*SyntheticsStepType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticsStepDetail) SetType(v SyntheticsStepType)`

SetType sets Type field to given value.

### HasType

`func (o *SyntheticsStepDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *SyntheticsStepDetail) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SyntheticsStepDetail) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SyntheticsStepDetail) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SyntheticsStepDetail) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetValue

`func (o *SyntheticsStepDetail) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SyntheticsStepDetail) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SyntheticsStepDetail) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *SyntheticsStepDetail) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetWarnings

`func (o *SyntheticsStepDetail) GetWarnings() []SyntheticsStepDetailWarnings`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SyntheticsStepDetail) GetWarningsOk() (*[]SyntheticsStepDetailWarnings, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SyntheticsStepDetail) SetWarnings(v []SyntheticsStepDetailWarnings)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SyntheticsStepDetail) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


