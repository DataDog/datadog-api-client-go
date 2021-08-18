# IncidentResponseAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Created** | Pointer to **time.Time** | Timestamp when the incident was created. | [optional] [readonly] 
**CustomerImpactDuration** | Pointer to **int64** | Length of the incident&#39;s customer impact in seconds. Equals the difference between &#x60;customer_impact_start&#x60; and &#x60;customer_impact_end&#x60;. | [optional] [readonly] 
**CustomerImpactEnd** | Pointer to **NullableTime** | Timestamp when customers were no longer impacted by the incident. | [optional] 
**CustomerImpactScope** | Pointer to **NullableString** | A summary of the impact customers experienced during the incident. | [optional] 
**CustomerImpactStart** | Pointer to **NullableTime** | Timestamp when customers began being impacted by the incident. | [optional] 
**CustomerImpacted** | Pointer to **bool** | A flag indicating whether the incident caused customer impact. | [optional] 
**Detected** | Pointer to **NullableTime** | Timestamp when the incident was detected. | [optional] 
**Fields** | Pointer to [**map[string]IncidentFieldAttributes**](IncidentFieldAttributes.md) | A condensed view of the user-defined fields attached to incidents. | [optional] 
**Modified** | Pointer to **time.Time** | Timestamp when the incident was last modified. | [optional] [readonly] 
**NotificationHandles** | Pointer to [**[]IncidentNotificationHandle**](IncidentNotificationHandle.md) | Notification handles that will be notified of the incident during update. | [optional] 
**PostmortemId** | Pointer to **string** | The UUID of the postmortem object attached to the incident. | [optional] 
**PublicId** | Pointer to **int64** | The monotonically increasing integer ID for the incident. | [optional] 
**Resolved** | Pointer to **NullableTime** | Timestamp when the incident&#39;s state was set to resolved. | [optional] 
**TimeToDetect** | Pointer to **int64** | The amount of time in seconds to detect the incident. Equals the difference between &#x60;customer_impact_start&#x60; and &#x60;detected&#x60;. | [optional] [readonly] 
**TimeToInternalResponse** | Pointer to **int64** | The amount of time in seconds to call incident after detection. Equals the difference of &#x60;detected&#x60; and &#x60;created&#x60;. | [optional] [readonly] 
**TimeToRepair** | Pointer to **int64** | The amount of time in seconds to resolve customer impact after detecting the issue. Equals the difference between &#x60;customer_impact_end&#x60; and &#x60;detected&#x60;. | [optional] [readonly] 
**TimeToResolve** | Pointer to **int64** | The amount of time in seconds to resolve the incident after it was created. Equals the difference between &#x60;created&#x60; and &#x60;resolved&#x60;. | [optional] [readonly] 
**Title** | **string** | The title of the incident, which summarizes what happened. | 

## Methods

### NewIncidentResponseAttributes

`func NewIncidentResponseAttributes(title string) *IncidentResponseAttributes`

NewIncidentResponseAttributes instantiates a new IncidentResponseAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewIncidentResponseAttributesWithDefaults

`func NewIncidentResponseAttributesWithDefaults() *IncidentResponseAttributes`

NewIncidentResponseAttributesWithDefaults instantiates a new IncidentResponseAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCreated

`func (o *IncidentResponseAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IncidentResponseAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IncidentResponseAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IncidentResponseAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCustomerImpactDuration

`func (o *IncidentResponseAttributes) GetCustomerImpactDuration() int64`

GetCustomerImpactDuration returns the CustomerImpactDuration field if non-nil, zero value otherwise.

### GetCustomerImpactDurationOk

`func (o *IncidentResponseAttributes) GetCustomerImpactDurationOk() (*int64, bool)`

GetCustomerImpactDurationOk returns a tuple with the CustomerImpactDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerImpactDuration

`func (o *IncidentResponseAttributes) SetCustomerImpactDuration(v int64)`

SetCustomerImpactDuration sets CustomerImpactDuration field to given value.

### HasCustomerImpactDuration

`func (o *IncidentResponseAttributes) HasCustomerImpactDuration() bool`

HasCustomerImpactDuration returns a boolean if a field has been set.

### GetCustomerImpactEnd

`func (o *IncidentResponseAttributes) GetCustomerImpactEnd() time.Time`

GetCustomerImpactEnd returns the CustomerImpactEnd field if non-nil, zero value otherwise.

### GetCustomerImpactEndOk

`func (o *IncidentResponseAttributes) GetCustomerImpactEndOk() (*time.Time, bool)`

GetCustomerImpactEndOk returns a tuple with the CustomerImpactEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerImpactEnd

`func (o *IncidentResponseAttributes) SetCustomerImpactEnd(v time.Time)`

SetCustomerImpactEnd sets CustomerImpactEnd field to given value.

### HasCustomerImpactEnd

`func (o *IncidentResponseAttributes) HasCustomerImpactEnd() bool`

HasCustomerImpactEnd returns a boolean if a field has been set.

### SetCustomerImpactEndNil

`func (o *IncidentResponseAttributes) SetCustomerImpactEndNil(b bool)`

 SetCustomerImpactEndNil sets the value for CustomerImpactEnd to be an explicit nil

### UnsetCustomerImpactEnd
`func (o *IncidentResponseAttributes) UnsetCustomerImpactEnd()`

UnsetCustomerImpactEnd ensures that no value is present for CustomerImpactEnd, not even an explicit nil
### GetCustomerImpactScope

`func (o *IncidentResponseAttributes) GetCustomerImpactScope() string`

GetCustomerImpactScope returns the CustomerImpactScope field if non-nil, zero value otherwise.

### GetCustomerImpactScopeOk

`func (o *IncidentResponseAttributes) GetCustomerImpactScopeOk() (*string, bool)`

GetCustomerImpactScopeOk returns a tuple with the CustomerImpactScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerImpactScope

`func (o *IncidentResponseAttributes) SetCustomerImpactScope(v string)`

SetCustomerImpactScope sets CustomerImpactScope field to given value.

### HasCustomerImpactScope

`func (o *IncidentResponseAttributes) HasCustomerImpactScope() bool`

HasCustomerImpactScope returns a boolean if a field has been set.

### SetCustomerImpactScopeNil

`func (o *IncidentResponseAttributes) SetCustomerImpactScopeNil(b bool)`

 SetCustomerImpactScopeNil sets the value for CustomerImpactScope to be an explicit nil

### UnsetCustomerImpactScope
`func (o *IncidentResponseAttributes) UnsetCustomerImpactScope()`

UnsetCustomerImpactScope ensures that no value is present for CustomerImpactScope, not even an explicit nil
### GetCustomerImpactStart

`func (o *IncidentResponseAttributes) GetCustomerImpactStart() time.Time`

GetCustomerImpactStart returns the CustomerImpactStart field if non-nil, zero value otherwise.

### GetCustomerImpactStartOk

`func (o *IncidentResponseAttributes) GetCustomerImpactStartOk() (*time.Time, bool)`

GetCustomerImpactStartOk returns a tuple with the CustomerImpactStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerImpactStart

`func (o *IncidentResponseAttributes) SetCustomerImpactStart(v time.Time)`

SetCustomerImpactStart sets CustomerImpactStart field to given value.

### HasCustomerImpactStart

`func (o *IncidentResponseAttributes) HasCustomerImpactStart() bool`

HasCustomerImpactStart returns a boolean if a field has been set.

### SetCustomerImpactStartNil

`func (o *IncidentResponseAttributes) SetCustomerImpactStartNil(b bool)`

 SetCustomerImpactStartNil sets the value for CustomerImpactStart to be an explicit nil

### UnsetCustomerImpactStart
`func (o *IncidentResponseAttributes) UnsetCustomerImpactStart()`

UnsetCustomerImpactStart ensures that no value is present for CustomerImpactStart, not even an explicit nil
### GetCustomerImpacted

`func (o *IncidentResponseAttributes) GetCustomerImpacted() bool`

GetCustomerImpacted returns the CustomerImpacted field if non-nil, zero value otherwise.

### GetCustomerImpactedOk

`func (o *IncidentResponseAttributes) GetCustomerImpactedOk() (*bool, bool)`

GetCustomerImpactedOk returns a tuple with the CustomerImpacted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerImpacted

`func (o *IncidentResponseAttributes) SetCustomerImpacted(v bool)`

SetCustomerImpacted sets CustomerImpacted field to given value.

### HasCustomerImpacted

`func (o *IncidentResponseAttributes) HasCustomerImpacted() bool`

HasCustomerImpacted returns a boolean if a field has been set.

### GetDetected

`func (o *IncidentResponseAttributes) GetDetected() time.Time`

GetDetected returns the Detected field if non-nil, zero value otherwise.

### GetDetectedOk

`func (o *IncidentResponseAttributes) GetDetectedOk() (*time.Time, bool)`

GetDetectedOk returns a tuple with the Detected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetected

`func (o *IncidentResponseAttributes) SetDetected(v time.Time)`

SetDetected sets Detected field to given value.

### HasDetected

`func (o *IncidentResponseAttributes) HasDetected() bool`

HasDetected returns a boolean if a field has been set.

### SetDetectedNil

`func (o *IncidentResponseAttributes) SetDetectedNil(b bool)`

 SetDetectedNil sets the value for Detected to be an explicit nil

### UnsetDetected
`func (o *IncidentResponseAttributes) UnsetDetected()`

UnsetDetected ensures that no value is present for Detected, not even an explicit nil
### GetFields

`func (o *IncidentResponseAttributes) GetFields() map[string]IncidentFieldAttributes`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *IncidentResponseAttributes) GetFieldsOk() (*map[string]IncidentFieldAttributes, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *IncidentResponseAttributes) SetFields(v map[string]IncidentFieldAttributes)`

SetFields sets Fields field to given value.

### HasFields

`func (o *IncidentResponseAttributes) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetModified

`func (o *IncidentResponseAttributes) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *IncidentResponseAttributes) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *IncidentResponseAttributes) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *IncidentResponseAttributes) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetNotificationHandles

`func (o *IncidentResponseAttributes) GetNotificationHandles() []IncidentNotificationHandle`

GetNotificationHandles returns the NotificationHandles field if non-nil, zero value otherwise.

### GetNotificationHandlesOk

`func (o *IncidentResponseAttributes) GetNotificationHandlesOk() (*[]IncidentNotificationHandle, bool)`

GetNotificationHandlesOk returns a tuple with the NotificationHandles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationHandles

`func (o *IncidentResponseAttributes) SetNotificationHandles(v []IncidentNotificationHandle)`

SetNotificationHandles sets NotificationHandles field to given value.

### HasNotificationHandles

`func (o *IncidentResponseAttributes) HasNotificationHandles() bool`

HasNotificationHandles returns a boolean if a field has been set.

### GetPostmortemId

`func (o *IncidentResponseAttributes) GetPostmortemId() string`

GetPostmortemId returns the PostmortemId field if non-nil, zero value otherwise.

### GetPostmortemIdOk

`func (o *IncidentResponseAttributes) GetPostmortemIdOk() (*string, bool)`

GetPostmortemIdOk returns a tuple with the PostmortemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostmortemId

`func (o *IncidentResponseAttributes) SetPostmortemId(v string)`

SetPostmortemId sets PostmortemId field to given value.

### HasPostmortemId

`func (o *IncidentResponseAttributes) HasPostmortemId() bool`

HasPostmortemId returns a boolean if a field has been set.

### GetPublicId

`func (o *IncidentResponseAttributes) GetPublicId() int64`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *IncidentResponseAttributes) GetPublicIdOk() (*int64, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *IncidentResponseAttributes) SetPublicId(v int64)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *IncidentResponseAttributes) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetResolved

`func (o *IncidentResponseAttributes) GetResolved() time.Time`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *IncidentResponseAttributes) GetResolvedOk() (*time.Time, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *IncidentResponseAttributes) SetResolved(v time.Time)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *IncidentResponseAttributes) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### SetResolvedNil

`func (o *IncidentResponseAttributes) SetResolvedNil(b bool)`

 SetResolvedNil sets the value for Resolved to be an explicit nil

### UnsetResolved
`func (o *IncidentResponseAttributes) UnsetResolved()`

UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
### GetTimeToDetect

`func (o *IncidentResponseAttributes) GetTimeToDetect() int64`

GetTimeToDetect returns the TimeToDetect field if non-nil, zero value otherwise.

### GetTimeToDetectOk

`func (o *IncidentResponseAttributes) GetTimeToDetectOk() (*int64, bool)`

GetTimeToDetectOk returns a tuple with the TimeToDetect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToDetect

`func (o *IncidentResponseAttributes) SetTimeToDetect(v int64)`

SetTimeToDetect sets TimeToDetect field to given value.

### HasTimeToDetect

`func (o *IncidentResponseAttributes) HasTimeToDetect() bool`

HasTimeToDetect returns a boolean if a field has been set.

### GetTimeToInternalResponse

`func (o *IncidentResponseAttributes) GetTimeToInternalResponse() int64`

GetTimeToInternalResponse returns the TimeToInternalResponse field if non-nil, zero value otherwise.

### GetTimeToInternalResponseOk

`func (o *IncidentResponseAttributes) GetTimeToInternalResponseOk() (*int64, bool)`

GetTimeToInternalResponseOk returns a tuple with the TimeToInternalResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToInternalResponse

`func (o *IncidentResponseAttributes) SetTimeToInternalResponse(v int64)`

SetTimeToInternalResponse sets TimeToInternalResponse field to given value.

### HasTimeToInternalResponse

`func (o *IncidentResponseAttributes) HasTimeToInternalResponse() bool`

HasTimeToInternalResponse returns a boolean if a field has been set.

### GetTimeToRepair

`func (o *IncidentResponseAttributes) GetTimeToRepair() int64`

GetTimeToRepair returns the TimeToRepair field if non-nil, zero value otherwise.

### GetTimeToRepairOk

`func (o *IncidentResponseAttributes) GetTimeToRepairOk() (*int64, bool)`

GetTimeToRepairOk returns a tuple with the TimeToRepair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToRepair

`func (o *IncidentResponseAttributes) SetTimeToRepair(v int64)`

SetTimeToRepair sets TimeToRepair field to given value.

### HasTimeToRepair

`func (o *IncidentResponseAttributes) HasTimeToRepair() bool`

HasTimeToRepair returns a boolean if a field has been set.

### GetTimeToResolve

`func (o *IncidentResponseAttributes) GetTimeToResolve() int64`

GetTimeToResolve returns the TimeToResolve field if non-nil, zero value otherwise.

### GetTimeToResolveOk

`func (o *IncidentResponseAttributes) GetTimeToResolveOk() (*int64, bool)`

GetTimeToResolveOk returns a tuple with the TimeToResolve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToResolve

`func (o *IncidentResponseAttributes) SetTimeToResolve(v int64)`

SetTimeToResolve sets TimeToResolve field to given value.

### HasTimeToResolve

`func (o *IncidentResponseAttributes) HasTimeToResolve() bool`

HasTimeToResolve returns a boolean if a field has been set.

### GetTitle

`func (o *IncidentResponseAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IncidentResponseAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IncidentResponseAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


