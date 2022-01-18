# SyntheticsBrowserTest

## Properties

| Name          | Type                                                                     | Description                                                                                   | Notes                                          |
| ------------- | ------------------------------------------------------------------------ | --------------------------------------------------------------------------------------------- | ---------------------------------------------- |
| **Config**    | [**SyntheticsBrowserTestConfig**](SyntheticsBrowserTestConfig.md)        |                                                                                               |
| **Locations** | **[]string**                                                             | Array of locations used to run the test.                                                      |
| **Message**   | Pointer to **string**                                                    | Notification message associated with the test. Message can either be text or an empty string. | [optional]                                     |
| **MonitorId** | Pointer to **int64**                                                     | The associated monitor ID.                                                                    | [optional] [readonly]                          |
| **Name**      | **string**                                                               | Name of the test.                                                                             |
| **Options**   | [**SyntheticsTestOptions**](SyntheticsTestOptions.md)                    |                                                                                               |
| **PublicId**  | Pointer to **string**                                                    | The public ID of the test.                                                                    | [optional] [readonly]                          |
| **Status**    | Pointer to [**SyntheticsTestPauseStatus**](SyntheticsTestPauseStatus.md) |                                                                                               | [optional]                                     |
| **Steps**     | Pointer to [**[]SyntheticsStep**](SyntheticsStep.md)                     | The steps of the test.                                                                        | [optional]                                     |
| **Tags**      | Pointer to **[]string**                                                  | Array of tags attached to the test.                                                           | [optional]                                     |
| **Type**      | [**SyntheticsBrowserTestType**](SyntheticsBrowserTestType.md)            |                                                                                               | [default to SYNTHETICSBROWSERTESTTYPE_BROWSER] |

## Methods

### NewSyntheticsBrowserTest

`func NewSyntheticsBrowserTest(config SyntheticsBrowserTestConfig, locations []string, name string, options SyntheticsTestOptions, type_ SyntheticsBrowserTestType) *SyntheticsBrowserTest`

NewSyntheticsBrowserTest instantiates a new SyntheticsBrowserTest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsBrowserTestWithDefaults

`func NewSyntheticsBrowserTestWithDefaults() *SyntheticsBrowserTest`

NewSyntheticsBrowserTestWithDefaults instantiates a new SyntheticsBrowserTest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetConfig

`func (o *SyntheticsBrowserTest) GetConfig() SyntheticsBrowserTestConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SyntheticsBrowserTest) GetConfigOk() (*SyntheticsBrowserTestConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SyntheticsBrowserTest) SetConfig(v SyntheticsBrowserTestConfig)`

SetConfig sets Config field to given value.

### GetLocations

`func (o *SyntheticsBrowserTest) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *SyntheticsBrowserTest) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *SyntheticsBrowserTest) SetLocations(v []string)`

SetLocations sets Locations field to given value.

### GetMessage

`func (o *SyntheticsBrowserTest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SyntheticsBrowserTest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SyntheticsBrowserTest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SyntheticsBrowserTest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMonitorId

`func (o *SyntheticsBrowserTest) GetMonitorId() int64`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *SyntheticsBrowserTest) GetMonitorIdOk() (*int64, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *SyntheticsBrowserTest) SetMonitorId(v int64)`

SetMonitorId sets MonitorId field to given value.

### HasMonitorId

`func (o *SyntheticsBrowserTest) HasMonitorId() bool`

HasMonitorId returns a boolean if a field has been set.

### GetName

`func (o *SyntheticsBrowserTest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticsBrowserTest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyntheticsBrowserTest) SetName(v string)`

SetName sets Name field to given value.

### GetOptions

`func (o *SyntheticsBrowserTest) GetOptions() SyntheticsTestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SyntheticsBrowserTest) GetOptionsOk() (*SyntheticsTestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SyntheticsBrowserTest) SetOptions(v SyntheticsTestOptions)`

SetOptions sets Options field to given value.

### GetPublicId

`func (o *SyntheticsBrowserTest) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *SyntheticsBrowserTest) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *SyntheticsBrowserTest) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *SyntheticsBrowserTest) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetStatus

`func (o *SyntheticsBrowserTest) GetStatus() SyntheticsTestPauseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyntheticsBrowserTest) GetStatusOk() (*SyntheticsTestPauseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyntheticsBrowserTest) SetStatus(v SyntheticsTestPauseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyntheticsBrowserTest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSteps

`func (o *SyntheticsBrowserTest) GetSteps() []SyntheticsStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *SyntheticsBrowserTest) GetStepsOk() (*[]SyntheticsStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *SyntheticsBrowserTest) SetSteps(v []SyntheticsStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *SyntheticsBrowserTest) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTags

`func (o *SyntheticsBrowserTest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SyntheticsBrowserTest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SyntheticsBrowserTest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SyntheticsBrowserTest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *SyntheticsBrowserTest) GetType() SyntheticsBrowserTestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsBrowserTest) GetTypeOk() (*SyntheticsBrowserTestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticsBrowserTest) SetType(v SyntheticsBrowserTestType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
