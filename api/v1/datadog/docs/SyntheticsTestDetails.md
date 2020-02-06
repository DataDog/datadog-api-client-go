# SyntheticsTestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**SyntheticsTestConfig**](SyntheticsTestConfig.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**SyntheticsTestAuthor**](SyntheticsTestAuthor.md) |  | [optional] 
**Locations** | Pointer to **[]string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **string** |  | [optional] 
**ModifiedBy** | Pointer to [**SyntheticsTestAuthor**](SyntheticsTestAuthor.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Options** | Pointer to [**SyntheticsTestOptions**](SyntheticsTestOptions.md) |  | [optional] 
**OverallState** | Pointer to [**SyntheticsTestMonitorStatus**](SyntheticsTestMonitorStatus.md) |  | [optional] 
**PublicId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**SyntheticsTestPauseStatus**](SyntheticsTestPauseStatus.md) |  | [optional] 
**StepCount** | Pointer to **int64** |  | [optional] 
**Subtype** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewSyntheticsTestDetails

`func NewSyntheticsTestDetails() *SyntheticsTestDetails`

NewSyntheticsTestDetails instantiates a new SyntheticsTestDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsTestDetailsWithDefaults

`func NewSyntheticsTestDetailsWithDefaults() *SyntheticsTestDetails`

NewSyntheticsTestDetailsWithDefaults instantiates a new SyntheticsTestDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *SyntheticsTestDetails) GetConfig() SyntheticsTestConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SyntheticsTestDetails) GetConfigOk() (SyntheticsTestConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConfig

`func (o *SyntheticsTestDetails) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfig

`func (o *SyntheticsTestDetails) SetConfig(v SyntheticsTestConfig)`

SetConfig gets a reference to the given SyntheticsTestConfig and assigns it to the Config field.

### GetCreatedAt

`func (o *SyntheticsTestDetails) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SyntheticsTestDetails) GetCreatedAtOk() (string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedAt

`func (o *SyntheticsTestDetails) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAt

`func (o *SyntheticsTestDetails) SetCreatedAt(v string)`

SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.

### GetCreatedBy

`func (o *SyntheticsTestDetails) GetCreatedBy() SyntheticsTestAuthor`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SyntheticsTestDetails) GetCreatedByOk() (SyntheticsTestAuthor, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *SyntheticsTestDetails) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *SyntheticsTestDetails) SetCreatedBy(v SyntheticsTestAuthor)`

SetCreatedBy gets a reference to the given SyntheticsTestAuthor and assigns it to the CreatedBy field.

### GetLocations

`func (o *SyntheticsTestDetails) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *SyntheticsTestDetails) GetLocationsOk() ([]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocations

`func (o *SyntheticsTestDetails) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### SetLocations

`func (o *SyntheticsTestDetails) SetLocations(v []string)`

SetLocations gets a reference to the given []string and assigns it to the Locations field.

### GetMessage

`func (o *SyntheticsTestDetails) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SyntheticsTestDetails) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *SyntheticsTestDetails) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *SyntheticsTestDetails) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.

### GetModifiedAt

`func (o *SyntheticsTestDetails) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SyntheticsTestDetails) GetModifiedAtOk() (string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModifiedAt

`func (o *SyntheticsTestDetails) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAt

`func (o *SyntheticsTestDetails) SetModifiedAt(v string)`

SetModifiedAt gets a reference to the given string and assigns it to the ModifiedAt field.

### GetModifiedBy

`func (o *SyntheticsTestDetails) GetModifiedBy() SyntheticsTestAuthor`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *SyntheticsTestDetails) GetModifiedByOk() (SyntheticsTestAuthor, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModifiedBy

`func (o *SyntheticsTestDetails) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### SetModifiedBy

`func (o *SyntheticsTestDetails) SetModifiedBy(v SyntheticsTestAuthor)`

SetModifiedBy gets a reference to the given SyntheticsTestAuthor and assigns it to the ModifiedBy field.

### GetName

`func (o *SyntheticsTestDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticsTestDetails) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *SyntheticsTestDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *SyntheticsTestDetails) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetOptions

`func (o *SyntheticsTestDetails) GetOptions() SyntheticsTestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SyntheticsTestDetails) GetOptionsOk() (SyntheticsTestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOptions

`func (o *SyntheticsTestDetails) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptions

`func (o *SyntheticsTestDetails) SetOptions(v SyntheticsTestOptions)`

SetOptions gets a reference to the given SyntheticsTestOptions and assigns it to the Options field.

### GetOverallState

`func (o *SyntheticsTestDetails) GetOverallState() SyntheticsTestMonitorStatus`

GetOverallState returns the OverallState field if non-nil, zero value otherwise.

### GetOverallStateOk

`func (o *SyntheticsTestDetails) GetOverallStateOk() (SyntheticsTestMonitorStatus, bool)`

GetOverallStateOk returns a tuple with the OverallState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOverallState

`func (o *SyntheticsTestDetails) HasOverallState() bool`

HasOverallState returns a boolean if a field has been set.

### SetOverallState

`func (o *SyntheticsTestDetails) SetOverallState(v SyntheticsTestMonitorStatus)`

SetOverallState gets a reference to the given SyntheticsTestMonitorStatus and assigns it to the OverallState field.

### GetPublicId

`func (o *SyntheticsTestDetails) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *SyntheticsTestDetails) GetPublicIdOk() (string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublicId

`func (o *SyntheticsTestDetails) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### SetPublicId

`func (o *SyntheticsTestDetails) SetPublicId(v string)`

SetPublicId gets a reference to the given string and assigns it to the PublicId field.

### GetStatus

`func (o *SyntheticsTestDetails) GetStatus() SyntheticsTestPauseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyntheticsTestDetails) GetStatusOk() (SyntheticsTestPauseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *SyntheticsTestDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *SyntheticsTestDetails) SetStatus(v SyntheticsTestPauseStatus)`

SetStatus gets a reference to the given SyntheticsTestPauseStatus and assigns it to the Status field.

### GetStepCount

`func (o *SyntheticsTestDetails) GetStepCount() int64`

GetStepCount returns the StepCount field if non-nil, zero value otherwise.

### GetStepCountOk

`func (o *SyntheticsTestDetails) GetStepCountOk() (int64, bool)`

GetStepCountOk returns a tuple with the StepCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStepCount

`func (o *SyntheticsTestDetails) HasStepCount() bool`

HasStepCount returns a boolean if a field has been set.

### SetStepCount

`func (o *SyntheticsTestDetails) SetStepCount(v int64)`

SetStepCount gets a reference to the given int64 and assigns it to the StepCount field.

### GetSubtype

`func (o *SyntheticsTestDetails) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *SyntheticsTestDetails) GetSubtypeOk() (string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubtype

`func (o *SyntheticsTestDetails) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### SetSubtype

`func (o *SyntheticsTestDetails) SetSubtype(v string)`

SetSubtype gets a reference to the given string and assigns it to the Subtype field.

### GetTags

`func (o *SyntheticsTestDetails) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SyntheticsTestDetails) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *SyntheticsTestDetails) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *SyntheticsTestDetails) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.

### GetType

`func (o *SyntheticsTestDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsTestDetails) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *SyntheticsTestDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *SyntheticsTestDetails) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


