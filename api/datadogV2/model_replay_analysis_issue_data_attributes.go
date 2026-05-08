// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReplayAnalysisIssueDataAttributes Attributes of a RUM replay analysis issue.
type ReplayAnalysisIssueDataAttributes struct {
	// Up to three sample sessions affected by this issue.
	AffectedSessions []ReplayAnalysisAffectedSession `json:"affected_sessions"`
	// Unique identifier of the application where the issue was detected.
	ApplicationId string `json:"application_id"`
	// Timestamp when the issue was first detected.
	CreatedAt time.Time `json:"created_at"`
	// Human-readable description of the issue.
	Description string `json:"description"`
	// Journey query associated with the issue.
	JourneyQuery map[string]interface{} `json:"journey_query,omitempty"`
	// Name of the issue.
	Name string `json:"name"`
	// A representative session illustrating a replay analysis issue.
	RepresentativeSession ReplayAnalysisRepresentativeSession `json:"representative_session"`
	// Total number of sessions affected by this issue.
	SessionCount int64 `json:"session_count"`
	// Severity level of the issue. Valid values are `high`, `medium`, and `low`.
	Severity string `json:"severity"`
	// Timestamp when the issue was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// Validation status of the issue.
	ValidationVerdict string `json:"validation_verdict"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReplayAnalysisIssueDataAttributes instantiates a new ReplayAnalysisIssueDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReplayAnalysisIssueDataAttributes(affectedSessions []ReplayAnalysisAffectedSession, applicationId string, createdAt time.Time, description string, name string, representativeSession ReplayAnalysisRepresentativeSession, sessionCount int64, severity string, updatedAt time.Time, validationVerdict string) *ReplayAnalysisIssueDataAttributes {
	this := ReplayAnalysisIssueDataAttributes{}
	this.AffectedSessions = affectedSessions
	this.ApplicationId = applicationId
	this.CreatedAt = createdAt
	this.Description = description
	this.Name = name
	this.RepresentativeSession = representativeSession
	this.SessionCount = sessionCount
	this.Severity = severity
	this.UpdatedAt = updatedAt
	this.ValidationVerdict = validationVerdict
	return &this
}

// NewReplayAnalysisIssueDataAttributesWithDefaults instantiates a new ReplayAnalysisIssueDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReplayAnalysisIssueDataAttributesWithDefaults() *ReplayAnalysisIssueDataAttributes {
	this := ReplayAnalysisIssueDataAttributes{}
	return &this
}

// GetAffectedSessions returns the AffectedSessions field value.
func (o *ReplayAnalysisIssueDataAttributes) GetAffectedSessions() []ReplayAnalysisAffectedSession {
	if o == nil {
		var ret []ReplayAnalysisAffectedSession
		return ret
	}
	return o.AffectedSessions
}

// GetAffectedSessionsOk returns a tuple with the AffectedSessions field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisIssueDataAttributes) GetAffectedSessionsOk() (*[]ReplayAnalysisAffectedSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AffectedSessions, true
}

// SetAffectedSessions sets field value.
func (o *ReplayAnalysisIssueDataAttributes) SetAffectedSessions(v []ReplayAnalysisAffectedSession) {
	o.AffectedSessions = v
}

// GetApplicationId returns the ApplicationId field value.
func (o *ReplayAnalysisIssueDataAttributes) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisIssueDataAttributes) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value.
func (o *ReplayAnalysisIssueDataAttributes) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *ReplayAnalysisIssueDataAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisIssueDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *ReplayAnalysisIssueDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value.
func (o *ReplayAnalysisIssueDataAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisIssueDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ReplayAnalysisIssueDataAttributes) SetDescription(v string) {
	o.Description = v
}

// GetJourneyQuery returns the JourneyQuery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReplayAnalysisIssueDataAttributes) GetJourneyQuery() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.JourneyQuery
}

// GetJourneyQueryOk returns a tuple with the JourneyQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReplayAnalysisIssueDataAttributes) GetJourneyQueryOk() (*map[string]interface{}, bool) {
	if o == nil || o.JourneyQuery == nil {
		return nil, false
	}
	return &o.JourneyQuery, true
}

// HasJourneyQuery returns a boolean if a field has been set.
func (o *ReplayAnalysisIssueDataAttributes) HasJourneyQuery() bool {
	return o != nil && o.JourneyQuery != nil
}

// SetJourneyQuery gets a reference to the given map[string]interface{} and assigns it to the JourneyQuery field.
func (o *ReplayAnalysisIssueDataAttributes) SetJourneyQuery(v map[string]interface{}) {
	o.JourneyQuery = v
}

// GetName returns the Name field value.
func (o *ReplayAnalysisIssueDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisIssueDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ReplayAnalysisIssueDataAttributes) SetName(v string) {
	o.Name = v
}

// GetRepresentativeSession returns the RepresentativeSession field value.
func (o *ReplayAnalysisIssueDataAttributes) GetRepresentativeSession() ReplayAnalysisRepresentativeSession {
	if o == nil {
		var ret ReplayAnalysisRepresentativeSession
		return ret
	}
	return o.RepresentativeSession
}

// GetRepresentativeSessionOk returns a tuple with the RepresentativeSession field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisIssueDataAttributes) GetRepresentativeSessionOk() (*ReplayAnalysisRepresentativeSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepresentativeSession, true
}

// SetRepresentativeSession sets field value.
func (o *ReplayAnalysisIssueDataAttributes) SetRepresentativeSession(v ReplayAnalysisRepresentativeSession) {
	o.RepresentativeSession = v
}

// GetSessionCount returns the SessionCount field value.
func (o *ReplayAnalysisIssueDataAttributes) GetSessionCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SessionCount
}

// GetSessionCountOk returns a tuple with the SessionCount field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisIssueDataAttributes) GetSessionCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionCount, true
}

// SetSessionCount sets field value.
func (o *ReplayAnalysisIssueDataAttributes) SetSessionCount(v int64) {
	o.SessionCount = v
}

// GetSeverity returns the Severity field value.
func (o *ReplayAnalysisIssueDataAttributes) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisIssueDataAttributes) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *ReplayAnalysisIssueDataAttributes) SetSeverity(v string) {
	o.Severity = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *ReplayAnalysisIssueDataAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisIssueDataAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *ReplayAnalysisIssueDataAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetValidationVerdict returns the ValidationVerdict field value.
func (o *ReplayAnalysisIssueDataAttributes) GetValidationVerdict() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ValidationVerdict
}

// GetValidationVerdictOk returns a tuple with the ValidationVerdict field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisIssueDataAttributes) GetValidationVerdictOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidationVerdict, true
}

// SetValidationVerdict sets field value.
func (o *ReplayAnalysisIssueDataAttributes) SetValidationVerdict(v string) {
	o.ValidationVerdict = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReplayAnalysisIssueDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["affected_sessions"] = o.AffectedSessions
	toSerialize["application_id"] = o.ApplicationId
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["description"] = o.Description
	if o.JourneyQuery != nil {
		toSerialize["journey_query"] = o.JourneyQuery
	}
	toSerialize["name"] = o.Name
	toSerialize["representative_session"] = o.RepresentativeSession
	toSerialize["session_count"] = o.SessionCount
	toSerialize["severity"] = o.Severity
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["validation_verdict"] = o.ValidationVerdict

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReplayAnalysisIssueDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AffectedSessions      *[]ReplayAnalysisAffectedSession     `json:"affected_sessions"`
		ApplicationId         *string                              `json:"application_id"`
		CreatedAt             *time.Time                           `json:"created_at"`
		Description           *string                              `json:"description"`
		JourneyQuery          map[string]interface{}               `json:"journey_query,omitempty"`
		Name                  *string                              `json:"name"`
		RepresentativeSession *ReplayAnalysisRepresentativeSession `json:"representative_session"`
		SessionCount          *int64                               `json:"session_count"`
		Severity              *string                              `json:"severity"`
		UpdatedAt             *time.Time                           `json:"updated_at"`
		ValidationVerdict     *string                              `json:"validation_verdict"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AffectedSessions == nil {
		return fmt.Errorf("required field affected_sessions missing")
	}
	if all.ApplicationId == nil {
		return fmt.Errorf("required field application_id missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.RepresentativeSession == nil {
		return fmt.Errorf("required field representative_session missing")
	}
	if all.SessionCount == nil {
		return fmt.Errorf("required field session_count missing")
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	if all.ValidationVerdict == nil {
		return fmt.Errorf("required field validation_verdict missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"affected_sessions", "application_id", "created_at", "description", "journey_query", "name", "representative_session", "session_count", "severity", "updated_at", "validation_verdict"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AffectedSessions = *all.AffectedSessions
	o.ApplicationId = *all.ApplicationId
	o.CreatedAt = *all.CreatedAt
	o.Description = *all.Description
	o.JourneyQuery = all.JourneyQuery
	o.Name = *all.Name
	if all.RepresentativeSession.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RepresentativeSession = *all.RepresentativeSession
	o.SessionCount = *all.SessionCount
	o.Severity = *all.Severity
	o.UpdatedAt = *all.UpdatedAt
	o.ValidationVerdict = *all.ValidationVerdict

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
