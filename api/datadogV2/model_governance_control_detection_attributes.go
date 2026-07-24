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
	AssignedTeam *string `json:"assigned_team,omitempty"`
	// The identifier of the user the detection is assigned to, if any.
	AssignedTo *string `json:"assigned_to,omitempty"`
	// How the detection's current assignment was determined. Possible values are `auto_resolved`, `manual`, `reassigned`, and `cleared`.
	AssignmentSource GovernanceControlDetectionAssignmentSource `json:"assignment_source"`
	// The unique identifier of the control that produced this detection.
	ControlId string `json:"control_id"`
	// The date and time when the detection was created.
	CreatedAt time.Time `json:"created_at"`
	// The type of detection, which determines what condition was detected.
	DetectionType string `json:"detection_type"`
	// The human-readable name of the detected resource.
	DisplayName string `json:"display_name"`
	// The date and time when the detection was marked as an exception, if applicable.
	ExceptionAt *time.Time `json:"exception_at,omitempty"`
	// The identifier of the user who marked the detection as an exception, if applicable.
	ExceptionBy *string `json:"exception_by,omitempty"`
	// Free-form metadata associated with the detection.
	Metadata interface{} `json:"metadata,omitempty"`
	// The date and time after which the detection is scheduled to be mitigated, if applicable.
	MitigateAfter *time.Time `json:"mitigate_after,omitempty"`
	// The date and time when the detection was mitigated, if applicable.
	MitigatedAt *time.Time `json:"mitigated_at,omitempty"`
	// The priority of the detection, if set.
	Priority int64 `json:"priority"`
	// The identifier of the resource the detection applies to.
	ResourceId string `json:"resource_id"`
	// A link to the detected resource in Datadog.
	ResourceUrl string `json:"resource_url"`
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
func NewGovernanceControlDetectionAttributes(assignmentSource GovernanceControlDetectionAssignmentSource, controlId string, createdAt time.Time, detectionType string, displayName string, priority int64, resourceId string, resourceUrl string, state GovernanceControlDetectionState) *GovernanceControlDetectionAttributes {
	this := GovernanceControlDetectionAttributes{}
	this.AssignmentSource = assignmentSource
	this.ControlId = controlId
	this.CreatedAt = createdAt
	this.DetectionType = detectionType
	this.DisplayName = displayName
	this.Priority = priority
	this.ResourceId = resourceId
	this.ResourceUrl = resourceUrl
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

// GetAssignedTeam returns the AssignedTeam field value if set, zero value otherwise.
func (o *GovernanceControlDetectionAttributes) GetAssignedTeam() string {
	if o == nil || o.AssignedTeam == nil {
		var ret string
		return ret
	}
	return *o.AssignedTeam
}

// GetAssignedTeamOk returns a tuple with the AssignedTeam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionAttributes) GetAssignedTeamOk() (*string, bool) {
	if o == nil || o.AssignedTeam == nil {
		return nil, false
	}
	return o.AssignedTeam, true
}

// HasAssignedTeam returns a boolean if a field has been set.
func (o *GovernanceControlDetectionAttributes) HasAssignedTeam() bool {
	return o != nil && o.AssignedTeam != nil
}

// SetAssignedTeam gets a reference to the given string and assigns it to the AssignedTeam field.
func (o *GovernanceControlDetectionAttributes) SetAssignedTeam(v string) {
	o.AssignedTeam = &v
}

// GetAssignedTo returns the AssignedTo field value if set, zero value otherwise.
func (o *GovernanceControlDetectionAttributes) GetAssignedTo() string {
	if o == nil || o.AssignedTo == nil {
		var ret string
		return ret
	}
	return *o.AssignedTo
}

// GetAssignedToOk returns a tuple with the AssignedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionAttributes) GetAssignedToOk() (*string, bool) {
	if o == nil || o.AssignedTo == nil {
		return nil, false
	}
	return o.AssignedTo, true
}

// HasAssignedTo returns a boolean if a field has been set.
func (o *GovernanceControlDetectionAttributes) HasAssignedTo() bool {
	return o != nil && o.AssignedTo != nil
}

// SetAssignedTo gets a reference to the given string and assigns it to the AssignedTo field.
func (o *GovernanceControlDetectionAttributes) SetAssignedTo(v string) {
	o.AssignedTo = &v
}

// GetAssignmentSource returns the AssignmentSource field value.
func (o *GovernanceControlDetectionAttributes) GetAssignmentSource() GovernanceControlDetectionAssignmentSource {
	if o == nil {
		var ret GovernanceControlDetectionAssignmentSource
		return ret
	}
	return o.AssignmentSource
}

// GetAssignmentSourceOk returns a tuple with the AssignmentSource field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionAttributes) GetAssignmentSourceOk() (*GovernanceControlDetectionAssignmentSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignmentSource, true
}

// SetAssignmentSource sets field value.
func (o *GovernanceControlDetectionAttributes) SetAssignmentSource(v GovernanceControlDetectionAssignmentSource) {
	o.AssignmentSource = v
}

// GetControlId returns the ControlId field value.
func (o *GovernanceControlDetectionAttributes) GetControlId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ControlId
}

// GetControlIdOk returns a tuple with the ControlId field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionAttributes) GetControlIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ControlId, true
}

// SetControlId sets field value.
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
func (o *GovernanceControlDetectionAttributes) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionAttributes) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *GovernanceControlDetectionAttributes) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetExceptionAt returns the ExceptionAt field value if set, zero value otherwise.
func (o *GovernanceControlDetectionAttributes) GetExceptionAt() time.Time {
	if o == nil || o.ExceptionAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExceptionAt
}

// GetExceptionAtOk returns a tuple with the ExceptionAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionAttributes) GetExceptionAtOk() (*time.Time, bool) {
	if o == nil || o.ExceptionAt == nil {
		return nil, false
	}
	return o.ExceptionAt, true
}

// HasExceptionAt returns a boolean if a field has been set.
func (o *GovernanceControlDetectionAttributes) HasExceptionAt() bool {
	return o != nil && o.ExceptionAt != nil
}

// SetExceptionAt gets a reference to the given time.Time and assigns it to the ExceptionAt field.
func (o *GovernanceControlDetectionAttributes) SetExceptionAt(v time.Time) {
	o.ExceptionAt = &v
}

// GetExceptionBy returns the ExceptionBy field value if set, zero value otherwise.
func (o *GovernanceControlDetectionAttributes) GetExceptionBy() string {
	if o == nil || o.ExceptionBy == nil {
		var ret string
		return ret
	}
	return *o.ExceptionBy
}

// GetExceptionByOk returns a tuple with the ExceptionBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionAttributes) GetExceptionByOk() (*string, bool) {
	if o == nil || o.ExceptionBy == nil {
		return nil, false
	}
	return o.ExceptionBy, true
}

// HasExceptionBy returns a boolean if a field has been set.
func (o *GovernanceControlDetectionAttributes) HasExceptionBy() bool {
	return o != nil && o.ExceptionBy != nil
}

// SetExceptionBy gets a reference to the given string and assigns it to the ExceptionBy field.
func (o *GovernanceControlDetectionAttributes) SetExceptionBy(v string) {
	o.ExceptionBy = &v
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

// GetMitigateAfter returns the MitigateAfter field value if set, zero value otherwise.
func (o *GovernanceControlDetectionAttributes) GetMitigateAfter() time.Time {
	if o == nil || o.MitigateAfter == nil {
		var ret time.Time
		return ret
	}
	return *o.MitigateAfter
}

// GetMitigateAfterOk returns a tuple with the MitigateAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionAttributes) GetMitigateAfterOk() (*time.Time, bool) {
	if o == nil || o.MitigateAfter == nil {
		return nil, false
	}
	return o.MitigateAfter, true
}

// HasMitigateAfter returns a boolean if a field has been set.
func (o *GovernanceControlDetectionAttributes) HasMitigateAfter() bool {
	return o != nil && o.MitigateAfter != nil
}

// SetMitigateAfter gets a reference to the given time.Time and assigns it to the MitigateAfter field.
func (o *GovernanceControlDetectionAttributes) SetMitigateAfter(v time.Time) {
	o.MitigateAfter = &v
}

// GetMitigatedAt returns the MitigatedAt field value if set, zero value otherwise.
func (o *GovernanceControlDetectionAttributes) GetMitigatedAt() time.Time {
	if o == nil || o.MitigatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.MitigatedAt
}

// GetMitigatedAtOk returns a tuple with the MitigatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionAttributes) GetMitigatedAtOk() (*time.Time, bool) {
	if o == nil || o.MitigatedAt == nil {
		return nil, false
	}
	return o.MitigatedAt, true
}

// HasMitigatedAt returns a boolean if a field has been set.
func (o *GovernanceControlDetectionAttributes) HasMitigatedAt() bool {
	return o != nil && o.MitigatedAt != nil
}

// SetMitigatedAt gets a reference to the given time.Time and assigns it to the MitigatedAt field.
func (o *GovernanceControlDetectionAttributes) SetMitigatedAt(v time.Time) {
	o.MitigatedAt = &v
}

// GetPriority returns the Priority field value.
func (o *GovernanceControlDetectionAttributes) GetPriority() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionAttributes) GetPriorityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value.
func (o *GovernanceControlDetectionAttributes) SetPriority(v int64) {
	o.Priority = v
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

// GetResourceUrl returns the ResourceUrl field value.
func (o *GovernanceControlDetectionAttributes) GetResourceUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceUrl
}

// GetResourceUrlOk returns a tuple with the ResourceUrl field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionAttributes) GetResourceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceUrl, true
}

// SetResourceUrl sets field value.
func (o *GovernanceControlDetectionAttributes) SetResourceUrl(v string) {
	o.ResourceUrl = v
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
	if o.AssignedTeam != nil {
		toSerialize["assigned_team"] = o.AssignedTeam
	}
	if o.AssignedTo != nil {
		toSerialize["assigned_to"] = o.AssignedTo
	}
	toSerialize["assignment_source"] = o.AssignmentSource
	toSerialize["control_id"] = o.ControlId
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["detection_type"] = o.DetectionType
	toSerialize["display_name"] = o.DisplayName
	if o.ExceptionAt != nil {
		if o.ExceptionAt.Nanosecond() == 0 {
			toSerialize["exception_at"] = o.ExceptionAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["exception_at"] = o.ExceptionAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ExceptionBy != nil {
		toSerialize["exception_by"] = o.ExceptionBy
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.MitigateAfter != nil {
		if o.MitigateAfter.Nanosecond() == 0 {
			toSerialize["mitigate_after"] = o.MitigateAfter.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["mitigate_after"] = o.MitigateAfter.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.MitigatedAt != nil {
		if o.MitigatedAt.Nanosecond() == 0 {
			toSerialize["mitigated_at"] = o.MitigatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["mitigated_at"] = o.MitigatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["priority"] = o.Priority
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["resource_url"] = o.ResourceUrl
	toSerialize["state"] = o.State

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceControlDetectionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssignedTeam     *string                                     `json:"assigned_team,omitempty"`
		AssignedTo       *string                                     `json:"assigned_to,omitempty"`
		AssignmentSource *GovernanceControlDetectionAssignmentSource `json:"assignment_source"`
		ControlId        *string                                     `json:"control_id"`
		CreatedAt        *time.Time                                  `json:"created_at"`
		DetectionType    *string                                     `json:"detection_type"`
		DisplayName      *string                                     `json:"display_name"`
		ExceptionAt      *time.Time                                  `json:"exception_at,omitempty"`
		ExceptionBy      *string                                     `json:"exception_by,omitempty"`
		Metadata         interface{}                                 `json:"metadata,omitempty"`
		MitigateAfter    *time.Time                                  `json:"mitigate_after,omitempty"`
		MitigatedAt      *time.Time                                  `json:"mitigated_at,omitempty"`
		Priority         *int64                                      `json:"priority"`
		ResourceId       *string                                     `json:"resource_id"`
		ResourceUrl      *string                                     `json:"resource_url"`
		State            *GovernanceControlDetectionState            `json:"state"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AssignmentSource == nil {
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
	if all.DisplayName == nil {
		return fmt.Errorf("required field display_name missing")
	}
	if all.Priority == nil {
		return fmt.Errorf("required field priority missing")
	}
	if all.ResourceId == nil {
		return fmt.Errorf("required field resource_id missing")
	}
	if all.ResourceUrl == nil {
		return fmt.Errorf("required field resource_url missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assigned_team", "assigned_to", "assignment_source", "control_id", "created_at", "detection_type", "display_name", "exception_at", "exception_by", "metadata", "mitigate_after", "mitigated_at", "priority", "resource_id", "resource_url", "state"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AssignedTeam = all.AssignedTeam
	o.AssignedTo = all.AssignedTo
	if !all.AssignmentSource.IsValid() {
		hasInvalidField = true
	} else {
		o.AssignmentSource = *all.AssignmentSource
	}
	o.ControlId = *all.ControlId
	o.CreatedAt = *all.CreatedAt
	o.DetectionType = *all.DetectionType
	o.DisplayName = *all.DisplayName
	o.ExceptionAt = all.ExceptionAt
	o.ExceptionBy = all.ExceptionBy
	o.Metadata = all.Metadata
	o.MitigateAfter = all.MitigateAfter
	o.MitigatedAt = all.MitigatedAt
	o.Priority = *all.Priority
	o.ResourceId = *all.ResourceId
	o.ResourceUrl = *all.ResourceUrl
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
