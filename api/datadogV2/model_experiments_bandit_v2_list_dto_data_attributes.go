// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditV2ListDTODataAttributes Summary of a bandit. Use Get bandit for its setup, active policy, and latest refresh.
type ExperimentsBanditV2ListDTODataAttributes struct {
	// When the bandit was created.
	CreatedAt time.Time `json:"created_at"`
	// The bandit's description.
	Description string `json:"description"`
	// When the bandit was stopped or reached its scheduled end. Null while it hasn't ended.
	EndedAt datadog.NullableTime `json:"ended_at"`
	// The metric the bandit optimizes. Null until set on a draft.
	MetricId datadog.NullableUUID `json:"metric_id"`
	// The bandit's name.
	Name string `json:"name"`
	// The key of the arm most likely to be best, when one arm leads and the bandit is outside warmup. Null otherwise.
	OptimalVariantKey datadog.NullableString `json:"optimal_variant_key"`
	// Why no optimal variant is reported. Absent when one is, or when the cause isn't known.
	OptimalVariantUnavailableReason *ExperimentsBanditOptimalVariantUnavailableReason `json:"optimal_variant_unavailable_reason,omitempty"`
	// When the bandit started. Null until it starts.
	StartedAt datadog.NullableTime `json:"started_at"`
	// The bandit's status. STOPPED means the bandit was stopped and serves the variant chosen then. COMPLETED means it reached its scheduled end and serves its fallback.
	Status ExperimentsBanditStatus `json:"status"`
	// The subject type the bandit randomizes and measures by. Null until set on a draft.
	SubjectTypeId datadog.NullableUUID `json:"subject_type_id"`
	// The bandit's tags, without team tags.
	Tags []string `json:"tags"`
	// Handles of the teams that own the bandit.
	Teams []string `json:"teams"`
	// When the bandit was last changed.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditV2ListDTODataAttributes instantiates a new ExperimentsBanditV2ListDTODataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditV2ListDTODataAttributes(createdAt time.Time, description string, endedAt datadog.NullableTime, metricId datadog.NullableUUID, name string, optimalVariantKey datadog.NullableString, startedAt datadog.NullableTime, status ExperimentsBanditStatus, subjectTypeId datadog.NullableUUID, tags []string, teams []string, updatedAt time.Time) *ExperimentsBanditV2ListDTODataAttributes {
	this := ExperimentsBanditV2ListDTODataAttributes{}
	this.CreatedAt = createdAt
	this.Description = description
	this.EndedAt = endedAt
	this.MetricId = metricId
	this.Name = name
	this.OptimalVariantKey = optimalVariantKey
	this.StartedAt = startedAt
	this.Status = status
	this.SubjectTypeId = subjectTypeId
	this.Tags = tags
	this.Teams = teams
	this.UpdatedAt = updatedAt
	return &this
}

// NewExperimentsBanditV2ListDTODataAttributesWithDefaults instantiates a new ExperimentsBanditV2ListDTODataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditV2ListDTODataAttributesWithDefaults() *ExperimentsBanditV2ListDTODataAttributes {
	this := ExperimentsBanditV2ListDTODataAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *ExperimentsBanditV2ListDTODataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ExperimentsBanditV2ListDTODataAttributes) SetDescription(v string) {
	o.Description = v
}

// GetEndedAt returns the EndedAt field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetEndedAt() time.Time {
	if o == nil || o.EndedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndedAt.Get()
}

// GetEndedAtOk returns a tuple with the EndedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetEndedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndedAt.Get(), o.EndedAt.IsSet()
}

// SetEndedAt sets field value.
func (o *ExperimentsBanditV2ListDTODataAttributes) SetEndedAt(v time.Time) {
	o.EndedAt.Set(&v)
}

// GetMetricId returns the MetricId field value.
// If the value is explicit nil, the zero value for uuid.UUID will be returned.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetMetricId() uuid.UUID {
	if o == nil || o.MetricId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.MetricId.Get()
}

// GetMetricIdOk returns a tuple with the MetricId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetMetricIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricId.Get(), o.MetricId.IsSet()
}

// SetMetricId sets field value.
func (o *ExperimentsBanditV2ListDTODataAttributes) SetMetricId(v uuid.UUID) {
	o.MetricId.Set(&v)
}

// GetName returns the Name field value.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ExperimentsBanditV2ListDTODataAttributes) SetName(v string) {
	o.Name = v
}

// GetOptimalVariantKey returns the OptimalVariantKey field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetOptimalVariantKey() string {
	if o == nil || o.OptimalVariantKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.OptimalVariantKey.Get()
}

// GetOptimalVariantKeyOk returns a tuple with the OptimalVariantKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetOptimalVariantKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptimalVariantKey.Get(), o.OptimalVariantKey.IsSet()
}

// SetOptimalVariantKey sets field value.
func (o *ExperimentsBanditV2ListDTODataAttributes) SetOptimalVariantKey(v string) {
	o.OptimalVariantKey.Set(&v)
}

// GetOptimalVariantUnavailableReason returns the OptimalVariantUnavailableReason field value if set, zero value otherwise.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetOptimalVariantUnavailableReason() ExperimentsBanditOptimalVariantUnavailableReason {
	if o == nil || o.OptimalVariantUnavailableReason == nil {
		var ret ExperimentsBanditOptimalVariantUnavailableReason
		return ret
	}
	return *o.OptimalVariantUnavailableReason
}

// GetOptimalVariantUnavailableReasonOk returns a tuple with the OptimalVariantUnavailableReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetOptimalVariantUnavailableReasonOk() (*ExperimentsBanditOptimalVariantUnavailableReason, bool) {
	if o == nil || o.OptimalVariantUnavailableReason == nil {
		return nil, false
	}
	return o.OptimalVariantUnavailableReason, true
}

// HasOptimalVariantUnavailableReason returns a boolean if a field has been set.
func (o *ExperimentsBanditV2ListDTODataAttributes) HasOptimalVariantUnavailableReason() bool {
	return o != nil && o.OptimalVariantUnavailableReason != nil
}

// SetOptimalVariantUnavailableReason gets a reference to the given ExperimentsBanditOptimalVariantUnavailableReason and assigns it to the OptimalVariantUnavailableReason field.
func (o *ExperimentsBanditV2ListDTODataAttributes) SetOptimalVariantUnavailableReason(v ExperimentsBanditOptimalVariantUnavailableReason) {
	o.OptimalVariantUnavailableReason = &v
}

// GetStartedAt returns the StartedAt field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetStartedAt() time.Time {
	if o == nil || o.StartedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// SetStartedAt sets field value.
func (o *ExperimentsBanditV2ListDTODataAttributes) SetStartedAt(v time.Time) {
	o.StartedAt.Set(&v)
}

// GetStatus returns the Status field value.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetStatus() ExperimentsBanditStatus {
	if o == nil {
		var ret ExperimentsBanditStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetStatusOk() (*ExperimentsBanditStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ExperimentsBanditV2ListDTODataAttributes) SetStatus(v ExperimentsBanditStatus) {
	o.Status = v
}

// GetSubjectTypeId returns the SubjectTypeId field value.
// If the value is explicit nil, the zero value for uuid.UUID will be returned.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetSubjectTypeId() uuid.UUID {
	if o == nil || o.SubjectTypeId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.SubjectTypeId.Get()
}

// GetSubjectTypeIdOk returns a tuple with the SubjectTypeId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetSubjectTypeIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubjectTypeId.Get(), o.SubjectTypeId.IsSet()
}

// SetSubjectTypeId sets field value.
func (o *ExperimentsBanditV2ListDTODataAttributes) SetSubjectTypeId(v uuid.UUID) {
	o.SubjectTypeId.Set(&v)
}

// GetTags returns the Tags field value.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *ExperimentsBanditV2ListDTODataAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTeams returns the Teams field value.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetTeams() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetTeamsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Teams, true
}

// SetTeams sets field value.
func (o *ExperimentsBanditV2ListDTODataAttributes) SetTeams(v []string) {
	o.Teams = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2ListDTODataAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *ExperimentsBanditV2ListDTODataAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditV2ListDTODataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["description"] = o.Description
	toSerialize["ended_at"] = o.EndedAt.Get()
	toSerialize["metric_id"] = o.MetricId.Get()
	toSerialize["name"] = o.Name
	toSerialize["optimal_variant_key"] = o.OptimalVariantKey.Get()
	if o.OptimalVariantUnavailableReason != nil {
		toSerialize["optimal_variant_unavailable_reason"] = o.OptimalVariantUnavailableReason
	}
	toSerialize["started_at"] = o.StartedAt.Get()
	toSerialize["status"] = o.Status
	toSerialize["subject_type_id"] = o.SubjectTypeId.Get()
	toSerialize["tags"] = o.Tags
	toSerialize["teams"] = o.Teams
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditV2ListDTODataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt                       *time.Time                                        `json:"created_at"`
		Description                     *string                                           `json:"description"`
		EndedAt                         datadog.NullableTime                              `json:"ended_at"`
		MetricId                        datadog.NullableUUID                              `json:"metric_id"`
		Name                            *string                                           `json:"name"`
		OptimalVariantKey               datadog.NullableString                            `json:"optimal_variant_key"`
		OptimalVariantUnavailableReason *ExperimentsBanditOptimalVariantUnavailableReason `json:"optimal_variant_unavailable_reason,omitempty"`
		StartedAt                       datadog.NullableTime                              `json:"started_at"`
		Status                          *ExperimentsBanditStatus                          `json:"status"`
		SubjectTypeId                   datadog.NullableUUID                              `json:"subject_type_id"`
		Tags                            *[]string                                         `json:"tags"`
		Teams                           *[]string                                         `json:"teams"`
		UpdatedAt                       *time.Time                                        `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if !all.EndedAt.IsSet() {
		return fmt.Errorf("required field ended_at missing")
	}
	if !all.MetricId.IsSet() {
		return fmt.Errorf("required field metric_id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if !all.OptimalVariantKey.IsSet() {
		return fmt.Errorf("required field optimal_variant_key missing")
	}
	if !all.StartedAt.IsSet() {
		return fmt.Errorf("required field started_at missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if !all.SubjectTypeId.IsSet() {
		return fmt.Errorf("required field subject_type_id missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	if all.Teams == nil {
		return fmt.Errorf("required field teams missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "description", "ended_at", "metric_id", "name", "optimal_variant_key", "optimal_variant_unavailable_reason", "started_at", "status", "subject_type_id", "tags", "teams", "updated_at"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.Description = *all.Description
	o.EndedAt = all.EndedAt
	o.MetricId = all.MetricId
	o.Name = *all.Name
	o.OptimalVariantKey = all.OptimalVariantKey
	if all.OptimalVariantUnavailableReason != nil && !all.OptimalVariantUnavailableReason.IsValid() {
		hasInvalidField = true
	} else {
		o.OptimalVariantUnavailableReason = all.OptimalVariantUnavailableReason
	}
	o.StartedAt = all.StartedAt
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.SubjectTypeId = all.SubjectTypeId
	o.Tags = *all.Tags
	o.Teams = *all.Teams
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
