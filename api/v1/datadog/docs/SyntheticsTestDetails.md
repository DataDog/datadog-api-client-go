# SyntheticsTestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**SyntheticsTestConfig**](SyntheticsTestConfig.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | TODO. | [optional] 
**CreatedBy** | Pointer to [**SyntheticsTestAuthor**](SyntheticsTestAuthor.md) |  | [optional] 
**Locations** | Pointer to **[]string** | TODO. | [optional] 
**Message** | Pointer to **string** | TODO. | [optional] 
**ModifiedAt** | Pointer to **string** | TODO. | [optional] 
**ModifiedBy** | Pointer to [**SyntheticsTestAuthor**](SyntheticsTestAuthor.md) |  | [optional] 
**Name** | Pointer to **string** | TODO. | [optional] 
**Options** | Pointer to [**SyntheticsTestOptions**](SyntheticsTestOptions.md) |  | [optional] 
**OverallState** | Pointer to [**SyntheticsTestMonitorStatus**](SyntheticsTestMonitorStatus.md) |  | [optional] 
**PublicId** | Pointer to **string** | TODO. | [optional] 
**Status** | Pointer to [**SyntheticsTestPauseStatus**](SyntheticsTestPauseStatus.md) |  | [optional] 
**Subtype** | Pointer to [**SyntheticsTestDetailsSubType**](SyntheticsTestDetailsSubType.md) |  | [optional] 
**Tags** | Pointer to **[]string** | TODO. | [optional] 
**Type** | Pointer to [**SyntheticsTestDetailsType**](SyntheticsTestDetailsType.md) |  | [optional] 

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

`func (o *SyntheticsTestDetails) GetConfigOk() (*SyntheticsTestConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SyntheticsTestDetails) SetConfig(v SyntheticsTestConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SyntheticsTestDetails) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SyntheticsTestDetails) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SyntheticsTestDetails) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SyntheticsTestDetails) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SyntheticsTestDetails) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SyntheticsTestDetails) GetCreatedBy() SyntheticsTestAuthor`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SyntheticsTestDetails) GetCreatedByOk() (*SyntheticsTestAuthor, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SyntheticsTestDetails) SetCreatedBy(v SyntheticsTestAuthor)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SyntheticsTestDetails) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLocations

`func (o *SyntheticsTestDetails) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *SyntheticsTestDetails) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *SyntheticsTestDetails) SetLocations(v []string)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *SyntheticsTestDetails) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetMessage

`func (o *SyntheticsTestDetails) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SyntheticsTestDetails) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SyntheticsTestDetails) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SyntheticsTestDetails) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetModifiedAt

`func (o *SyntheticsTestDetails) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SyntheticsTestDetails) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SyntheticsTestDetails) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *SyntheticsTestDetails) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *SyntheticsTestDetails) GetModifiedBy() SyntheticsTestAuthor`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *SyntheticsTestDetails) GetModifiedByOk() (*SyntheticsTestAuthor, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *SyntheticsTestDetails) SetModifiedBy(v SyntheticsTestAuthor)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *SyntheticsTestDetails) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetName

`func (o *SyntheticsTestDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticsTestDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyntheticsTestDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SyntheticsTestDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *SyntheticsTestDetails) GetOptions() SyntheticsTestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SyntheticsTestDetails) GetOptionsOk() (*SyntheticsTestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SyntheticsTestDetails) SetOptions(v SyntheticsTestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SyntheticsTestDetails) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOverallState

`func (o *SyntheticsTestDetails) GetOverallState() SyntheticsTestMonitorStatus`

GetOverallState returns the OverallState field if non-nil, zero value otherwise.

### GetOverallStateOk

`func (o *SyntheticsTestDetails) GetOverallStateOk() (*SyntheticsTestMonitorStatus, bool)`

GetOverallStateOk returns a tuple with the OverallState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallState

`func (o *SyntheticsTestDetails) SetOverallState(v SyntheticsTestMonitorStatus)`

SetOverallState sets OverallState field to given value.

### HasOverallState

`func (o *SyntheticsTestDetails) HasOverallState() bool`

HasOverallState returns a boolean if a field has been set.

### GetPublicId

`func (o *SyntheticsTestDetails) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *SyntheticsTestDetails) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *SyntheticsTestDetails) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *SyntheticsTestDetails) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetStatus

`func (o *SyntheticsTestDetails) GetStatus() SyntheticsTestPauseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyntheticsTestDetails) GetStatusOk() (*SyntheticsTestPauseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyntheticsTestDetails) SetStatus(v SyntheticsTestPauseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyntheticsTestDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubtype

`func (o *SyntheticsTestDetails) GetSubtype() SyntheticsTestDetailsSubType`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *SyntheticsTestDetails) GetSubtypeOk() (*SyntheticsTestDetailsSubType, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *SyntheticsTestDetails) SetSubtype(v SyntheticsTestDetailsSubType)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *SyntheticsTestDetails) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetTags

`func (o *SyntheticsTestDetails) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SyntheticsTestDetails) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SyntheticsTestDetails) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SyntheticsTestDetails) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *SyntheticsTestDetails) GetType() SyntheticsTestDetailsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsTestDetails) GetTypeOk() (*SyntheticsTestDetailsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticsTestDetails) SetType(v SyntheticsTestDetailsType)`

SetType sets Type field to given value.

### HasType

`func (o *SyntheticsTestDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


