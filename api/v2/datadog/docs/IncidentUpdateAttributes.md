# IncidentUpdateAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**CustomerImpactEnd** | Pointer to **NullableTime** | Timestamp when customers were no longer impacted by the incident. | [optional] 
**CustomerImpactScope** | Pointer to **string** | A summary of the impact customers experienced during the incident. | [optional] 
**CustomerImpactStart** | Pointer to **NullableTime** | Timestamp when customers began being impacted by the incident. | [optional] 
**CustomerImpacted** | Pointer to **bool** | A flag indicating whether the incident caused customer impact. | [optional] 
**Detected** | Pointer to **NullableTime** | Timestamp when the incident was detected. | [optional] 
**Fields** | Pointer to [**map[string]IncidentFieldAttributes**](IncidentFieldAttributes.md) | A condensed view of the user-defined fields for which to update selections. | [optional] 
**NotificationHandles** | Pointer to [**[]IncidentNotificationHandle**](IncidentNotificationHandle.md) | Notification handles that will be notified of the incident during update. | [optional] 
**Resolved** | Pointer to **NullableTime** | Timestamp when the incident&#39;s state was set to resolved. | [optional] 
**Title** | Pointer to **string** | The title of the incident, which summarizes what happened. | [optional] 

## Methods

### NewIncidentUpdateAttributes

`func NewIncidentUpdateAttributes() *IncidentUpdateAttributes`

NewIncidentUpdateAttributes instantiates a new IncidentUpdateAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewIncidentUpdateAttributesWithDefaults

`func NewIncidentUpdateAttributesWithDefaults() *IncidentUpdateAttributes`

NewIncidentUpdateAttributesWithDefaults instantiates a new IncidentUpdateAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCustomerImpactEnd

`func (o *IncidentUpdateAttributes) GetCustomerImpactEnd() time.Time`

GetCustomerImpactEnd returns the CustomerImpactEnd field if non-nil, zero value otherwise.

### GetCustomerImpactEndOk

`func (o *IncidentUpdateAttributes) GetCustomerImpactEndOk() (*time.Time, bool)`

GetCustomerImpactEndOk returns a tuple with the CustomerImpactEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerImpactEnd

`func (o *IncidentUpdateAttributes) SetCustomerImpactEnd(v time.Time)`

SetCustomerImpactEnd sets CustomerImpactEnd field to given value.

### HasCustomerImpactEnd

`func (o *IncidentUpdateAttributes) HasCustomerImpactEnd() bool`

HasCustomerImpactEnd returns a boolean if a field has been set.

### SetCustomerImpactEndNil

`func (o *IncidentUpdateAttributes) SetCustomerImpactEndNil(b bool)`

 SetCustomerImpactEndNil sets the value for CustomerImpactEnd to be an explicit nil

### UnsetCustomerImpactEnd
`func (o *IncidentUpdateAttributes) UnsetCustomerImpactEnd()`

UnsetCustomerImpactEnd ensures that no value is present for CustomerImpactEnd, not even an explicit nil
### GetCustomerImpactScope

`func (o *IncidentUpdateAttributes) GetCustomerImpactScope() string`

GetCustomerImpactScope returns the CustomerImpactScope field if non-nil, zero value otherwise.

### GetCustomerImpactScopeOk

`func (o *IncidentUpdateAttributes) GetCustomerImpactScopeOk() (*string, bool)`

GetCustomerImpactScopeOk returns a tuple with the CustomerImpactScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerImpactScope

`func (o *IncidentUpdateAttributes) SetCustomerImpactScope(v string)`

SetCustomerImpactScope sets CustomerImpactScope field to given value.

### HasCustomerImpactScope

`func (o *IncidentUpdateAttributes) HasCustomerImpactScope() bool`

HasCustomerImpactScope returns a boolean if a field has been set.

### GetCustomerImpactStart

`func (o *IncidentUpdateAttributes) GetCustomerImpactStart() time.Time`

GetCustomerImpactStart returns the CustomerImpactStart field if non-nil, zero value otherwise.

### GetCustomerImpactStartOk

`func (o *IncidentUpdateAttributes) GetCustomerImpactStartOk() (*time.Time, bool)`

GetCustomerImpactStartOk returns a tuple with the CustomerImpactStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerImpactStart

`func (o *IncidentUpdateAttributes) SetCustomerImpactStart(v time.Time)`

SetCustomerImpactStart sets CustomerImpactStart field to given value.

### HasCustomerImpactStart

`func (o *IncidentUpdateAttributes) HasCustomerImpactStart() bool`

HasCustomerImpactStart returns a boolean if a field has been set.

### SetCustomerImpactStartNil

`func (o *IncidentUpdateAttributes) SetCustomerImpactStartNil(b bool)`

 SetCustomerImpactStartNil sets the value for CustomerImpactStart to be an explicit nil

### UnsetCustomerImpactStart
`func (o *IncidentUpdateAttributes) UnsetCustomerImpactStart()`

UnsetCustomerImpactStart ensures that no value is present for CustomerImpactStart, not even an explicit nil
### GetCustomerImpacted

`func (o *IncidentUpdateAttributes) GetCustomerImpacted() bool`

GetCustomerImpacted returns the CustomerImpacted field if non-nil, zero value otherwise.

### GetCustomerImpactedOk

`func (o *IncidentUpdateAttributes) GetCustomerImpactedOk() (*bool, bool)`

GetCustomerImpactedOk returns a tuple with the CustomerImpacted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerImpacted

`func (o *IncidentUpdateAttributes) SetCustomerImpacted(v bool)`

SetCustomerImpacted sets CustomerImpacted field to given value.

### HasCustomerImpacted

`func (o *IncidentUpdateAttributes) HasCustomerImpacted() bool`

HasCustomerImpacted returns a boolean if a field has been set.

### GetDetected

`func (o *IncidentUpdateAttributes) GetDetected() time.Time`

GetDetected returns the Detected field if non-nil, zero value otherwise.

### GetDetectedOk

`func (o *IncidentUpdateAttributes) GetDetectedOk() (*time.Time, bool)`

GetDetectedOk returns a tuple with the Detected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetected

`func (o *IncidentUpdateAttributes) SetDetected(v time.Time)`

SetDetected sets Detected field to given value.

### HasDetected

`func (o *IncidentUpdateAttributes) HasDetected() bool`

HasDetected returns a boolean if a field has been set.

### SetDetectedNil

`func (o *IncidentUpdateAttributes) SetDetectedNil(b bool)`

 SetDetectedNil sets the value for Detected to be an explicit nil

### UnsetDetected
`func (o *IncidentUpdateAttributes) UnsetDetected()`

UnsetDetected ensures that no value is present for Detected, not even an explicit nil
### GetFields

`func (o *IncidentUpdateAttributes) GetFields() map[string]IncidentFieldAttributes`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *IncidentUpdateAttributes) GetFieldsOk() (*map[string]IncidentFieldAttributes, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *IncidentUpdateAttributes) SetFields(v map[string]IncidentFieldAttributes)`

SetFields sets Fields field to given value.

### HasFields

`func (o *IncidentUpdateAttributes) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetNotificationHandles

`func (o *IncidentUpdateAttributes) GetNotificationHandles() []IncidentNotificationHandle`

GetNotificationHandles returns the NotificationHandles field if non-nil, zero value otherwise.

### GetNotificationHandlesOk

`func (o *IncidentUpdateAttributes) GetNotificationHandlesOk() (*[]IncidentNotificationHandle, bool)`

GetNotificationHandlesOk returns a tuple with the NotificationHandles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationHandles

`func (o *IncidentUpdateAttributes) SetNotificationHandles(v []IncidentNotificationHandle)`

SetNotificationHandles sets NotificationHandles field to given value.

### HasNotificationHandles

`func (o *IncidentUpdateAttributes) HasNotificationHandles() bool`

HasNotificationHandles returns a boolean if a field has been set.

### GetResolved

`func (o *IncidentUpdateAttributes) GetResolved() time.Time`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *IncidentUpdateAttributes) GetResolvedOk() (*time.Time, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *IncidentUpdateAttributes) SetResolved(v time.Time)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *IncidentUpdateAttributes) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### SetResolvedNil

`func (o *IncidentUpdateAttributes) SetResolvedNil(b bool)`

 SetResolvedNil sets the value for Resolved to be an explicit nil

### UnsetResolved
`func (o *IncidentUpdateAttributes) UnsetResolved()`

UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
### GetTitle

`func (o *IncidentUpdateAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IncidentUpdateAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IncidentUpdateAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IncidentUpdateAttributes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


