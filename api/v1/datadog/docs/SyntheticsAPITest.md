# SyntheticsAPITest

## Properties

| Name          | Type                                                                           | Description                                    | Notes                                             |
| ------------- | ------------------------------------------------------------------------------ | ---------------------------------------------- | ------------------------------------------------- |
| **Config**    | Pointer to [**SyntheticsAPITestConfig**](SyntheticsAPITestConfig.md)           |                                                | [optional]                                        |
| **Locations** | Pointer to **[]string**                                                        | Array of locations used to run the test.       | [optional]                                        |
| **Message**   | Pointer to **string**                                                          | Notification message associated with the test. | [optional]                                        |
| **MonitorId** | Pointer to **int64**                                                           | The associated monitor ID.                     | [optional] [readonly]                             |
| **Name**      | Pointer to **string**                                                          | Name of the test.                              | [optional]                                        |
| **Options**   | Pointer to [**SyntheticsTestOptions**](SyntheticsTestOptions.md)               |                                                | [optional]                                        |
| **PublicId**  | Pointer to **string**                                                          | The public ID for the test.                    | [optional] [readonly]                             |
| **Status**    | Pointer to [**SyntheticsTestPauseStatus**](SyntheticsTestPauseStatus.md)       |                                                | [optional]                                        |
| **Subtype**   | Pointer to [**SyntheticsTestDetailsSubType**](SyntheticsTestDetailsSubType.md) |                                                | [optional]                                        |
| **Tags**      | Pointer to **[]string**                                                        | Array of tags attached to the test.            | [optional]                                        |
| **Type**      | Pointer to [**SyntheticsAPITestType**](SyntheticsAPITestType.md)               |                                                | [optional] [default to SYNTHETICSAPITESTTYPE_API] |

## Methods

### NewSyntheticsAPITest

`func NewSyntheticsAPITest() *SyntheticsAPITest`

NewSyntheticsAPITest instantiates a new SyntheticsAPITest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsAPITestWithDefaults

`func NewSyntheticsAPITestWithDefaults() *SyntheticsAPITest`

NewSyntheticsAPITestWithDefaults instantiates a new SyntheticsAPITest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetConfig

`func (o *SyntheticsAPITest) GetConfig() SyntheticsAPITestConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SyntheticsAPITest) GetConfigOk() (*SyntheticsAPITestConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SyntheticsAPITest) SetConfig(v SyntheticsAPITestConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SyntheticsAPITest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLocations

`func (o *SyntheticsAPITest) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *SyntheticsAPITest) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *SyntheticsAPITest) SetLocations(v []string)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *SyntheticsAPITest) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetMessage

`func (o *SyntheticsAPITest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SyntheticsAPITest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SyntheticsAPITest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SyntheticsAPITest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMonitorId

`func (o *SyntheticsAPITest) GetMonitorId() int64`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *SyntheticsAPITest) GetMonitorIdOk() (*int64, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *SyntheticsAPITest) SetMonitorId(v int64)`

SetMonitorId sets MonitorId field to given value.

### HasMonitorId

`func (o *SyntheticsAPITest) HasMonitorId() bool`

HasMonitorId returns a boolean if a field has been set.

### GetName

`func (o *SyntheticsAPITest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticsAPITest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyntheticsAPITest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SyntheticsAPITest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *SyntheticsAPITest) GetOptions() SyntheticsTestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SyntheticsAPITest) GetOptionsOk() (*SyntheticsTestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SyntheticsAPITest) SetOptions(v SyntheticsTestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SyntheticsAPITest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPublicId

`func (o *SyntheticsAPITest) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *SyntheticsAPITest) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *SyntheticsAPITest) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *SyntheticsAPITest) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetStatus

`func (o *SyntheticsAPITest) GetStatus() SyntheticsTestPauseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyntheticsAPITest) GetStatusOk() (*SyntheticsTestPauseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyntheticsAPITest) SetStatus(v SyntheticsTestPauseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyntheticsAPITest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubtype

`func (o *SyntheticsAPITest) GetSubtype() SyntheticsTestDetailsSubType`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *SyntheticsAPITest) GetSubtypeOk() (*SyntheticsTestDetailsSubType, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *SyntheticsAPITest) SetSubtype(v SyntheticsTestDetailsSubType)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *SyntheticsAPITest) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetTags

`func (o *SyntheticsAPITest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SyntheticsAPITest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SyntheticsAPITest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SyntheticsAPITest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *SyntheticsAPITest) GetType() SyntheticsAPITestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsAPITest) GetTypeOk() (*SyntheticsAPITestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticsAPITest) SetType(v SyntheticsAPITestType)`

SetType sets Type field to given value.

### HasType

`func (o *SyntheticsAPITest) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
