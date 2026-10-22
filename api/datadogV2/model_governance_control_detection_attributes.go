// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceControlDetectionAttributes The attributes of a governance control detection.
type GovernanceControlDetectionAttributes struct {
	// The identifier of the team the detection is assigned to, if any.
	AssignedTeam datadog.NullableString `json:"assigned_team,omitempty"`
	// The identifier of the user the detection is assigned to, if any.
	AssignedTo datadog.NullableString `json:"assigned_to,omitempty"`
	// How the detection's current assignment was determined. Possible values are `auto_resolved`, `manual`, `reassigned`, and `cleared`. `null` when not set.
	AssignmentSource NullableGovernanceControlDetectionAssignmentSource `json:"assignment_source"`
	// DEPRECATED: mirrors `detection_type` for backward compatibility; use `detection_type`
	// instead.
	// Deprecated
	ControlId string `json:"control_id"`
	// The date and time when the detection was created.
	CreatedAt time.Time `json:"created_at"`
	// The type of detection, which determines what condition was detected.
	DetectionType string `json:"detection_type"`
	// The human-readable name of the detected resource.
	DisplayName datadog.NullableString `json:"display_name"`
	// The date and time when the detection was marked as an exception, if applicable.
	ExceptionAt datadog.NullableTime `json:"exception_at,omitempty"`
	// The identifier of the user who marked the detection as an exception, if applicable.
	ExceptionBy datadog.NullableString `json:"exception_by,omitempty"`
	// Free-form metadata associated with the detection.
	Metadata interface{} `json:"metadata,omitempty"`
	// The date and time after which the detection is scheduled to be mitigated, if applicable.
	MitigateAfter datadog.NullableTime `json:"mitigate_after,omitempty"`
	// The date and time when the detection was mitigated, if applicable.
	MitigatedAt datadog.NullableTime `json:"mitigated_at,omitempty"`
	// The priority of the detection, if set.
	Priority datadog.NullableInt64 `json:"priority"`
	// The identifier of the resource the detection applies to.
	ResourceId string `json:"resource_id"`
	// The type of resource the detection applies to, for example `api_key` or `dashboard`.
	ResourceType datadog.NullableString `json:"resource_type"`
	// The current state of the detection. Possible values are `active`, `exception`, `mitigated`, `inactive`, `obsolete`, `resolved_externally`, and `mitigation_in_progress`.
	State GovernanceControlDetectionState `json:"state"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceControlDetectionAttributes instantiates a new GovernanceControlDetectionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceControlDetectionAttributes(assignmentSource NullableGovernanceControlDetectionAssignmentSource, controlId string, createdAt time.Time, detectionType string, displayName datadog.NullableString, priority datadog.NullableInt64, resourceId string, resourceType datadog.NullableString, state GovernanceControlDetectionState) *GovernanceControlDetectionAttributes {
	this := GovernanceControlDetectionAttributes{}
	this.AssignmentSource = assignmentSource
	this.ControlId = controlId
	this.CreatedAt = createdAt
	this.DetectionType = detectionType
	this.DisplayName = displayName
	this.Priority = priority
	this.ResourceId = resourceId
	this.ResourceType = resourceType
	this.State = state
	return &this
}

// NewGovernanceControlDetectionAttributesWithDefaults instantiates a new GovernanceControlDetectionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceControlDetectionAttributesWithDefaults() *GovernanceControlDetectionAttributes {
	this := GovernanceControlDetectionAttributes{}
	return &this
}

// GetAssignedTeam returns the AssignedTeam field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GovernanceControlDetectionAttributes) GetAssignedTeam() string {
	if o == nil || o.AssignedTeam.Get() == nil {
		var ret string
		return ret
	}
	return *o.AssignedTeam.Get()
}

// GetAssignedTeamOk returns a tuple with the AssignedTeam field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GovernanceControlDetectionAttributes) GetAssignedTeamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedTeam.Get(), o.AssignedTeam.IsSet()
}

// HasAssignedTeam returns a boolean if a field has been set.
func (o *GovernanceControlDetectionAttributes) HasAssignedTeam() bool {
	return o != nil && o.AssignedTeam.IsSet()
}

// SetAssignedTeam gets a reference to the given datadog.NullableString and assigns it to the AssignedTeam field.
func (o *GovernanceControlDetectionAttributes) SetAssignedTeam(v string) {
	o.AssignedTeam.Set(&v)
}

// SetAssignedTeamNil sets the value for AssignedTeam to be an explicit nil.
func (o *GovernanceControlDetectionAttributes) SetAssignedTeamNil() {
	o.AssignedTeam.Set(nil)
}

// UnsetAssignedTeam ensures that no value is present for AssignedTeam, not even an explicit nil.
func (o *GovernanceControlDetectionAttributes) UnsetAssignedTeam() {
	o.AssignedTeam.Unset()
}

// GetAssignedTo returns the AssignedTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GovernanceControlDetectionAttributes) GetAssignedTo() string {
	if o == nil || o.AssignedTo.Get() == nil {
		var ret string
		return ret
	}
	return *o.AssignedTo.Get()
}

// GetAssignedToOk returns a tuple with the AssignedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GovernanceControlDetectionAttributes) GetAssignedToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedTo.Get(), o.AssignedTo.IsSet()
}

// HasAssignedTo returns a boolean if a field has been set.
func (o *GovernanceControlDetectionAttributes) HasAssignedTo() bool {
	return o != nil && o.AssignedTo.IsSet()
}

// SetAssignedTo gets a reference to the given datadog.NullableString and assigns it to the AssignedTo field.
func (o *GovernanceControlDetectionAttributes) SetAssignedTo(v string) {
	o.AssignedTo.Set(&v)
}

// SetAssignedToNil sets the value for AssignedTo to be an explicit nil.
func (o *GovernanceControlDetectionAttributes) SetAssignedToNil() {
	o.AssignedTo.Set(nil)
}

// UnsetAssignedTo ensures that no value is present for AssignedTo, not even an explicit nil.
func (o *GovernanceControlDetectionAttributes) UnsetAssignedTo() {
	o.AssignedTo.Unset()
}

// GetAssignmentSource returns the AssignmentSource field value.
// If the value is explicit nil, the zero value for GovernanceControlDetectionAssignmentSource will be returned.
func (o *GovernanceControlDetectionAttributes) GetAssignmentSource() GovernanceControlDetectionAssignmentSource {
	if o == nil || o.AssignmentSource.Get() == nil {
		var ret GovernanceControlDetectionAssignmentSource
		return ret
	}
	return *o.AssignmentSource.Get()
}

// GetAssignmentSourceOk returns a tuple with the AssignmentSource field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GovernanceControlDetectionAttributes) GetAssignmentSourceOk() (*GovernanceControlDetectionAssignmentSource, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignmentSource.Get(), o.AssignmentSource.IsSet()
}

// SetAssignmentSource sets field value.
func (o *GovernanceControlDetectionAttributes) SetAssignmentSource(v GovernanceControlDetectionAssignmentSource) {
	o.AssignmentSource.Set(&v)
}

// GetControlId returns the ControlId field value.
// Deprecated
func (o *GovernanceControlDetectionAttributes) GetControlId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ControlId
}

// GetControlIdOk returns a tuple with the ControlId field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *GovernanceControlDetectionAttributes) GetControlIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ControlId, true
}

// SetControlId sets field value.
// Deprecated
func (o *GovernanceControlDetectionAttributes) SetControlId(v string) {
	o.ControlId = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *GovernanceControlDetectionAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *GovernanceControlDetectionAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDetectionType returns the DetectionType field value.
func (o *GovernanceControlDetectionAttributes) GetDetectionType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DetectionType
}

// GetDetectionTypeOk returns a tuple with the DetectionType field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionAttributes) GetDetectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DetectionType, true
}

// SetDetectionType sets field value.
func (o *GovernanceControlDetectionAttributes) SetDetectionType(v string) {
	o.DetectionType = v
}

// GetDisplayName returns the DisplayName field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *GovernanceControlDetectionAttributes) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GovernanceControlDetectionAttributes) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// SetDisplayName sets field value.
func (o *GovernanceControlDetectionAttributes) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}

// GetExceptionAt returns the ExceptionAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GovernanceControlDetectionAttributes) GetExceptionAt() time.Time {
	if o == nil || o.ExceptionAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExceptionAt.Get()
}

// GetExceptionAtOk returns a tuple with the ExceptionAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GovernanceControlDetectionAttributes) GetExceptionAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExceptionAt.Get(), o.ExceptionAt.IsSet()
}

// HasExceptionAt returns a boolean if a field has been set.
func (o *GovernanceControlDetectionAttributes) HasExceptionAt() bool {
	return o != nil && o.ExceptionAt.IsSet()
}

// SetExceptionAt gets a reference to the given datadog.NullableTime and assigns it to the ExceptionAt field.
func (o *GovernanceControlDetectionAttributes) SetExceptionAt(v time.Time) {
	o.ExceptionAt.Set(&v)
}

// SetExceptionAtNil sets the value for ExceptionAt to be an explicit nil.
func (o *GovernanceControlDetectionAttributes) SetExceptionAtNil() {
	o.ExceptionAt.Set(nil)
}

// UnsetExceptionAt ensures that no value is present for ExceptionAt, not even an explicit nil.
func (o *GovernanceControlDetectionAttributes) UnsetExceptionAt() {
	o.ExceptionAt.Unset()
}

// GetExceptionBy returns the ExceptionBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GovernanceControlDetectionAttributes) GetExceptionBy() string {
	if o == nil || o.ExceptionBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExceptionBy.Get()
}

// GetExceptionByOk returns a tuple with the ExceptionBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GovernanceControlDetectionAttributes) GetExceptionByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExceptionBy.Get(), o.ExceptionBy.IsSet()
}

// HasExceptionBy returns a boolean if a field has been set.
func (o *GovernanceControlDetectionAttributes) HasExceptionBy() bool {
	return o != nil && o.ExceptionBy.IsSet()
}

// SetExceptionBy gets a reference to the given datadog.NullableString and assigns it to the ExceptionBy field.
func (o *GovernanceControlDetectionAttributes) SetExceptionBy(v string) {
	o.ExceptionBy.Set(&v)
}

// SetExceptionByNil sets the value for ExceptionBy to be an explicit nil.
func (o *GovernanceControlDetectionAttributes) SetExceptionByNil() {
	o.ExceptionBy.Set(nil)
}

// UnsetExceptionBy ensures that no value is present for ExceptionBy, not even an explicit nil.
func (o *GovernanceControlDetectionAttributes) UnsetExceptionBy() {
	o.ExceptionBy.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GovernanceControlDetectionAttributes) GetMetadata() interface{} {
	if o == nil || o.Metadata == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GovernanceControlDetectionAttributes) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GovernanceControlDetectionAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetMitigateAfter returns the MitigateAfter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GovernanceControlDetectionAttributes) GetMitigateAfter() time.Time {
	if o == nil || o.MitigateAfter.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.MitigateAfter.Get()
}

// GetMitigateAfterOk returns a tuple with the MitigateAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GovernanceControlDetectionAttributes) GetMitigateAfterOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.MitigateAfter.Get(), o.MitigateAfter.IsSet()
}

// HasMitigateAfter returns a boolean if a field has been set.
func (o *GovernanceControlDetectionAttributes) HasMitigateAfter() bool {
	return o != nil && o.MitigateAfter.IsSet()
}

// SetMitigateAfter gets a reference to the given datadog.NullableTime and assigns it to the MitigateAfter field.
func (o *GovernanceControlDetectionAttributes) SetMitigateAfter(v time.Time) {
	o.MitigateAfter.Set(&v)
}

// SetMitigateAfterNil sets the value for MitigateAfter to be an explicit nil.
func (o *GovernanceControlDetectionAttributes) SetMitigateAfterNil() {
	o.MitigateAfter.Set(nil)
}

// UnsetMitigateAfter ensures that no value is present for MitigateAfter, not even an explicit nil.
func (o *GovernanceControlDetectionAttributes) UnsetMitigateAfter() {
	o.MitigateAfter.Unset()
}

// GetMitigatedAt returns the MitigatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GovernanceControlDetectionAttributes) GetMitigatedAt() time.Time {
	if o == nil || o.MitigatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.MitigatedAt.Get()
}

// GetMitigatedAtOk returns a tuple with the MitigatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GovernanceControlDetectionAttributes) GetMitigatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.MitigatedAt.Get(), o.MitigatedAt.IsSet()
}

// HasMitigatedAt returns a boolean if a field has been set.
func (o *GovernanceControlDetectionAttributes) HasMitigatedAt() bool {
	return o != nil && o.MitigatedAt.IsSet()
}

// SetMitigatedAt gets a reference to the given datadog.NullableTime and assigns it to the MitigatedAt field.
func (o *GovernanceControlDetectionAttributes) SetMitigatedAt(v time.Time) {
	o.MitigatedAt.Set(&v)
}

// SetMitigatedAtNil sets the value for MitigatedAt to be an explicit nil.
func (o *GovernanceControlDetectionAttributes) SetMitigatedAtNil() {
	o.MitigatedAt.Set(nil)
}

// UnsetMitigatedAt ensures that no value is present for MitigatedAt, not even an explicit nil.
func (o *GovernanceControlDetectionAttributes) UnsetMitigatedAt() {
	o.MitigatedAt.Unset()
}

// GetPriority returns the Priority field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *GovernanceControlDetectionAttributes) GetPriority() int64 {
	if o == nil || o.Priority.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GovernanceControlDetectionAttributes) GetPriorityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// SetPriority sets field value.
func (o *GovernanceControlDetectionAttributes) SetPriority(v int64) {
	o.Priority.Set(&v)
}

// GetResourceId returns the ResourceId field value.
func (o *GovernanceControlDetectionAttributes) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionAttributes) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value.
func (o *GovernanceControlDetectionAttributes) SetResourceId(v string) {
	o.ResourceId = v
}

// GetResourceType returns the ResourceType field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *GovernanceControlDetectionAttributes) GetResourceType() string {
	if o == nil || o.ResourceType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResourceType.Get()
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GovernanceControlDetectionAttributes) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceType.Get(), o.ResourceType.IsSet()
}

// SetResourceType sets field value.
func (o *GovernanceControlDetectionAttributes) SetResourceType(v string) {
	o.ResourceType.Set(&v)
}

// GetState returns the State field value.
func (o *GovernanceControlDetectionAttributes) GetState() GovernanceControlDetectionState {
	if o == nil {
		var ret GovernanceControlDetectionState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionAttributes) GetStateOk() (*GovernanceControlDetectionState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *GovernanceControlDetectionAttributes) SetState(v GovernanceControlDetectionState) {
	o.State = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceControlDetectionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssignedTeam.IsSet() {
		toSerialize["assigned_team"] = o.AssignedTeam.Get()
	}
	if o.AssignedTo.IsSet() {
		toSerialize["assigned_to"] = o.AssignedTo.Get()
	}
	toSerialize["assignment_source"] = o.AssignmentSource.Get()
	toSerialize["control_id"] = o.ControlId
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["detection_type"] = o.DetectionType
	toSerialize["display_name"] = o.DisplayName.Get()
	if o.ExceptionAt.IsSet() {
		toSerialize["exception_at"] = o.ExceptionAt.Get()
	}
	if o.ExceptionBy.IsSet() {
		toSerialize["exception_by"] = o.ExceptionBy.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.MitigateAfter.IsSet() {
		toSerialize["mitigate_after"] = o.MitigateAfter.Get()
	}
	if o.MitigatedAt.IsSet() {
		toSerialize["mitigated_at"] = o.MitigatedAt.Get()
	}
	toSerialize["priority"] = o.Priority.Get()
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["resource_type"] = o.ResourceType.Get()
	toSerialize["state"] = o.State

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceControlDetectionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssignedTeam     datadog.NullableString                             `json:"assigned_team,omitempty"`
		AssignedTo       datadog.NullableString                             `json:"assigned_to,omitempty"`
		AssignmentSource NullableGovernanceControlDetectionAssignmentSource `json:"assignment_source"`
		ControlId        *string                                            `json:"control_id"`
		CreatedAt        *time.Time                                         `json:"created_at"`
		DetectionType    *string                                            `json:"detection_type"`
		DisplayName      datadog.NullableString                             `json:"display_name"`
		ExceptionAt      datadog.NullableTime                               `json:"exception_at,omitempty"`
		ExceptionBy      datadog.NullableString                             `json:"exception_by,omitempty"`
		Metadata         interface{}                                        `json:"metadata,omitempty"`
		MitigateAfter    datadog.NullableTime                               `json:"mitigate_after,omitempty"`
		MitigatedAt      datadog.NullableTime                               `json:"mitigated_at,omitempty"`
		Priority         datadog.NullableInt64                              `json:"priority"`
		ResourceId       *string                                            `json:"resource_id"`
		ResourceType     datadog.NullableString                             `json:"resource_type"`
		State            *GovernanceControlDetectionState                   `json:"state"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.AssignmentSource.IsSet() {
		return fmt.Errorf("required field assignment_source missing")
	}
	if all.ControlId == nil {
		return fmt.Errorf("required field control_id missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.DetectionType == nil {
		return fmt.Errorf("required field detection_type missing")
	}
	if !all.DisplayName.IsSet() {
		return fmt.Errorf("required field display_name missing")
	}
	if !all.Priority.IsSet() {
		return fmt.Errorf("required field priority missing")
	}
	if all.ResourceId == nil {
		return fmt.Errorf("required field resource_id missing")
	}
	if !all.ResourceType.IsSet() {
		return fmt.Errorf("required field resource_type missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assigned_team", "assigned_to", "assignment_source", "control_id", "created_at", "detection_type", "display_name", "exception_at", "exception_by", "metadata", "mitigate_after", "mitigated_at", "priority", "resource_id", "resource_type", "state"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AssignedTeam = all.AssignedTeam
	o.AssignedTo = all.AssignedTo
	if all.AssignmentSource.Get() != nil && !all.AssignmentSource.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.AssignmentSource = all.AssignmentSource
	}
	o.ControlId = *all.ControlId
	o.CreatedAt = *all.CreatedAt
	o.DetectionType = *all.DetectionType
	o.DisplayName = all.DisplayName
	o.ExceptionAt = all.ExceptionAt
	o.ExceptionBy = all.ExceptionBy
	o.Metadata = all.Metadata
	o.MitigateAfter = all.MitigateAfter
	o.MitigatedAt = all.MitigatedAt
	o.Priority = all.Priority
	o.ResourceId = *all.ResourceId
	o.ResourceType = all.ResourceType
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
