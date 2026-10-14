// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsReplaceExperimentProtocolRequestDataAttributes Replace a published protocol template. Omitted optional template fields are cleared. Omitted enforcement and migration_metadata are preserved. Enabled limits require their values; allocation settings require environment_id.
type ExperimentsReplaceExperimentProtocolRequestDataAttributes struct {
	// Statistical settings. Use `default_duration_days` for duration; `target_duration_days` is not accepted here.
	AnalysisPlan NullableExperimentsProtocolAnalysisPlanInput `json:"analysis_plan,omitempty"`
	// Assignment properties.
	AssignmentSourceDefaultProperties []ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems `json:"assignment_source_default_properties,omitempty"`
	// ID of the assignment source selected by the protocol.
	AssignmentSourceId datadog.NullableUUID `json:"assignment_source_id,omitempty"`
	// Default experiment duration supplied by the protocol, in days.
	DefaultDurationDays datadog.NullableInt64 `json:"default_duration_days,omitempty"`
	// Complete lock map. Omit this object to preserve current locks. Null is rejected; every field is required.
	Enforcement *ExperimentsReplaceProtocolEnforcementInput `json:"enforcement,omitempty"`
	// ID of the feature flag environment used by the experiment.
	EnvironmentId datadog.NullableUUID `json:"environment_id,omitempty"`
	// Traffic schedule. Publishing and replacement require `NO_ROLLOUT` with exactly one step, or `UNIFORM_INTERVALS` with a positive selection interval and at least one step. Requires `environment_id`.
	ExposureSchedule NullableExperimentsProtocolExposureScheduleInput `json:"exposure_schedule,omitempty"`
	// Whether the protocol requires a duration before an experiment can start.
	IsDurationRequiredToStart bool `json:"is_duration_required_to_start"`
	// Whether the protocol requires equal traffic allocation across variants.
	IsEqualSplitEnforced bool `json:"is_equal_split_enforced"`
	// Whether the protocol limits the number of metrics.
	IsMetricLimitEnabled bool `json:"is_metric_limit_enabled"`
	// Whether the protocol enforces a minimum experiment duration.
	IsMinimumDurationEnabled bool `json:"is_minimum_duration_enabled"`
	// Metric groups.
	MetricGroups []ExperimentsProtocolMetricGroupInput `json:"metric_groups"`
	// Maximum number of metrics allowed by the protocol.
	MetricLimit datadog.NullableInt64 `json:"metric_limit,omitempty"`
	// Metadata for imported resources. Must encode to at most 4096 bytes. Null clears it; omission preserves it on update.
	MigrationMetadata map[string]interface{} `json:"migration_metadata,omitempty"`
	// Minimum duration unit, DAYS or WEEKS.
	MinimumDurationUnit datadog.NullableString `json:"minimum_duration_unit,omitempty"`
	// Minimum experiment duration in the specified unit.
	MinimumDurationValue datadog.NullableInt64 `json:"minimum_duration_value,omitempty"`
	// ID of the primary metric supplied by the protocol.
	PrimaryMetricId datadog.NullableUUID `json:"primary_metric_id,omitempty"`
	// ID of the subject type used by this configuration.
	SubjectTypeId uuid.UUID `json:"subject_type_id"`
	// Targeting rules.
	TargetingRules []ExperimentsProtocolTargetingRuleInput `json:"targeting_rules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsReplaceExperimentProtocolRequestDataAttributes instantiates a new ExperimentsReplaceExperimentProtocolRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsReplaceExperimentProtocolRequestDataAttributes(isDurationRequiredToStart bool, isEqualSplitEnforced bool, isMetricLimitEnabled bool, isMinimumDurationEnabled bool, metricGroups []ExperimentsProtocolMetricGroupInput, subjectTypeId uuid.UUID) *ExperimentsReplaceExperimentProtocolRequestDataAttributes {
	this := ExperimentsReplaceExperimentProtocolRequestDataAttributes{}
	this.IsDurationRequiredToStart = isDurationRequiredToStart
	this.IsEqualSplitEnforced = isEqualSplitEnforced
	this.IsMetricLimitEnabled = isMetricLimitEnabled
	this.IsMinimumDurationEnabled = isMinimumDurationEnabled
	this.MetricGroups = metricGroups
	this.SubjectTypeId = subjectTypeId
	return &this
}

// NewExperimentsReplaceExperimentProtocolRequestDataAttributesWithDefaults instantiates a new ExperimentsReplaceExperimentProtocolRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsReplaceExperimentProtocolRequestDataAttributesWithDefaults() *ExperimentsReplaceExperimentProtocolRequestDataAttributes {
	this := ExperimentsReplaceExperimentProtocolRequestDataAttributes{}
	return &this
}

// GetAnalysisPlan returns the AnalysisPlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetAnalysisPlan() ExperimentsProtocolAnalysisPlanInput {
	if o == nil || o.AnalysisPlan.Get() == nil {
		var ret ExperimentsProtocolAnalysisPlanInput
		return ret
	}
	return *o.AnalysisPlan.Get()
}

// GetAnalysisPlanOk returns a tuple with the AnalysisPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetAnalysisPlanOk() (*ExperimentsProtocolAnalysisPlanInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnalysisPlan.Get(), o.AnalysisPlan.IsSet()
}

// HasAnalysisPlan returns a boolean if a field has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) HasAnalysisPlan() bool {
	return o != nil && o.AnalysisPlan.IsSet()
}

// SetAnalysisPlan gets a reference to the given NullableExperimentsProtocolAnalysisPlanInput and assigns it to the AnalysisPlan field.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetAnalysisPlan(v ExperimentsProtocolAnalysisPlanInput) {
	o.AnalysisPlan.Set(&v)
}

// SetAnalysisPlanNil sets the value for AnalysisPlan to be an explicit nil.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetAnalysisPlanNil() {
	o.AnalysisPlan.Set(nil)
}

// UnsetAnalysisPlan ensures that no value is present for AnalysisPlan, not even an explicit nil.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) UnsetAnalysisPlan() {
	o.AnalysisPlan.Unset()
}

// GetAssignmentSourceDefaultProperties returns the AssignmentSourceDefaultProperties field value if set, zero value otherwise.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetAssignmentSourceDefaultProperties() []ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems {
	if o == nil || o.AssignmentSourceDefaultProperties == nil {
		var ret []ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems
		return ret
	}
	return o.AssignmentSourceDefaultProperties
}

// GetAssignmentSourceDefaultPropertiesOk returns a tuple with the AssignmentSourceDefaultProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetAssignmentSourceDefaultPropertiesOk() (*[]ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems, bool) {
	if o == nil || o.AssignmentSourceDefaultProperties == nil {
		return nil, false
	}
	return &o.AssignmentSourceDefaultProperties, true
}

// HasAssignmentSourceDefaultProperties returns a boolean if a field has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) HasAssignmentSourceDefaultProperties() bool {
	return o != nil && o.AssignmentSourceDefaultProperties != nil
}

// SetAssignmentSourceDefaultProperties gets a reference to the given []ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems and assigns it to the AssignmentSourceDefaultProperties field.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetAssignmentSourceDefaultProperties(v []ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems) {
	o.AssignmentSourceDefaultProperties = v
}

// GetAssignmentSourceId returns the AssignmentSourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetAssignmentSourceId() uuid.UUID {
	if o == nil || o.AssignmentSourceId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.AssignmentSourceId.Get()
}

// GetAssignmentSourceIdOk returns a tuple with the AssignmentSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetAssignmentSourceIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignmentSourceId.Get(), o.AssignmentSourceId.IsSet()
}

// HasAssignmentSourceId returns a boolean if a field has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) HasAssignmentSourceId() bool {
	return o != nil && o.AssignmentSourceId.IsSet()
}

// SetAssignmentSourceId gets a reference to the given datadog.NullableUUID and assigns it to the AssignmentSourceId field.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetAssignmentSourceId(v uuid.UUID) {
	o.AssignmentSourceId.Set(&v)
}

// SetAssignmentSourceIdNil sets the value for AssignmentSourceId to be an explicit nil.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetAssignmentSourceIdNil() {
	o.AssignmentSourceId.Set(nil)
}

// UnsetAssignmentSourceId ensures that no value is present for AssignmentSourceId, not even an explicit nil.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) UnsetAssignmentSourceId() {
	o.AssignmentSourceId.Unset()
}

// GetDefaultDurationDays returns the DefaultDurationDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetDefaultDurationDays() int64 {
	if o == nil || o.DefaultDurationDays.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DefaultDurationDays.Get()
}

// GetDefaultDurationDaysOk returns a tuple with the DefaultDurationDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetDefaultDurationDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultDurationDays.Get(), o.DefaultDurationDays.IsSet()
}

// HasDefaultDurationDays returns a boolean if a field has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) HasDefaultDurationDays() bool {
	return o != nil && o.DefaultDurationDays.IsSet()
}

// SetDefaultDurationDays gets a reference to the given datadog.NullableInt64 and assigns it to the DefaultDurationDays field.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetDefaultDurationDays(v int64) {
	o.DefaultDurationDays.Set(&v)
}

// SetDefaultDurationDaysNil sets the value for DefaultDurationDays to be an explicit nil.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetDefaultDurationDaysNil() {
	o.DefaultDurationDays.Set(nil)
}

// UnsetDefaultDurationDays ensures that no value is present for DefaultDurationDays, not even an explicit nil.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) UnsetDefaultDurationDays() {
	o.DefaultDurationDays.Unset()
}

// GetEnforcement returns the Enforcement field value if set, zero value otherwise.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetEnforcement() ExperimentsReplaceProtocolEnforcementInput {
	if o == nil || o.Enforcement == nil {
		var ret ExperimentsReplaceProtocolEnforcementInput
		return ret
	}
	return *o.Enforcement
}

// GetEnforcementOk returns a tuple with the Enforcement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetEnforcementOk() (*ExperimentsReplaceProtocolEnforcementInput, bool) {
	if o == nil || o.Enforcement == nil {
		return nil, false
	}
	return o.Enforcement, true
}

// HasEnforcement returns a boolean if a field has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) HasEnforcement() bool {
	return o != nil && o.Enforcement != nil
}

// SetEnforcement gets a reference to the given ExperimentsReplaceProtocolEnforcementInput and assigns it to the Enforcement field.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetEnforcement(v ExperimentsReplaceProtocolEnforcementInput) {
	o.Enforcement = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetEnvironmentId() uuid.UUID {
	if o == nil || o.EnvironmentId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.EnvironmentId.Get()
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetEnvironmentIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvironmentId.Get(), o.EnvironmentId.IsSet()
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId.IsSet()
}

// SetEnvironmentId gets a reference to the given datadog.NullableUUID and assigns it to the EnvironmentId field.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetEnvironmentId(v uuid.UUID) {
	o.EnvironmentId.Set(&v)
}

// SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetEnvironmentIdNil() {
	o.EnvironmentId.Set(nil)
}

// UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) UnsetEnvironmentId() {
	o.EnvironmentId.Unset()
}

// GetExposureSchedule returns the ExposureSchedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetExposureSchedule() ExperimentsProtocolExposureScheduleInput {
	if o == nil || o.ExposureSchedule.Get() == nil {
		var ret ExperimentsProtocolExposureScheduleInput
		return ret
	}
	return *o.ExposureSchedule.Get()
}

// GetExposureScheduleOk returns a tuple with the ExposureSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetExposureScheduleOk() (*ExperimentsProtocolExposureScheduleInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExposureSchedule.Get(), o.ExposureSchedule.IsSet()
}

// HasExposureSchedule returns a boolean if a field has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) HasExposureSchedule() bool {
	return o != nil && o.ExposureSchedule.IsSet()
}

// SetExposureSchedule gets a reference to the given NullableExperimentsProtocolExposureScheduleInput and assigns it to the ExposureSchedule field.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetExposureSchedule(v ExperimentsProtocolExposureScheduleInput) {
	o.ExposureSchedule.Set(&v)
}

// SetExposureScheduleNil sets the value for ExposureSchedule to be an explicit nil.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetExposureScheduleNil() {
	o.ExposureSchedule.Set(nil)
}

// UnsetExposureSchedule ensures that no value is present for ExposureSchedule, not even an explicit nil.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) UnsetExposureSchedule() {
	o.ExposureSchedule.Unset()
}

// GetIsDurationRequiredToStart returns the IsDurationRequiredToStart field value.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetIsDurationRequiredToStart() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsDurationRequiredToStart
}

// GetIsDurationRequiredToStartOk returns a tuple with the IsDurationRequiredToStart field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetIsDurationRequiredToStartOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDurationRequiredToStart, true
}

// SetIsDurationRequiredToStart sets field value.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetIsDurationRequiredToStart(v bool) {
	o.IsDurationRequiredToStart = v
}

// GetIsEqualSplitEnforced returns the IsEqualSplitEnforced field value.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetIsEqualSplitEnforced() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsEqualSplitEnforced
}

// GetIsEqualSplitEnforcedOk returns a tuple with the IsEqualSplitEnforced field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetIsEqualSplitEnforcedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEqualSplitEnforced, true
}

// SetIsEqualSplitEnforced sets field value.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetIsEqualSplitEnforced(v bool) {
	o.IsEqualSplitEnforced = v
}

// GetIsMetricLimitEnabled returns the IsMetricLimitEnabled field value.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetIsMetricLimitEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsMetricLimitEnabled
}

// GetIsMetricLimitEnabledOk returns a tuple with the IsMetricLimitEnabled field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetIsMetricLimitEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMetricLimitEnabled, true
}

// SetIsMetricLimitEnabled sets field value.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetIsMetricLimitEnabled(v bool) {
	o.IsMetricLimitEnabled = v
}

// GetIsMinimumDurationEnabled returns the IsMinimumDurationEnabled field value.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetIsMinimumDurationEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsMinimumDurationEnabled
}

// GetIsMinimumDurationEnabledOk returns a tuple with the IsMinimumDurationEnabled field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetIsMinimumDurationEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMinimumDurationEnabled, true
}

// SetIsMinimumDurationEnabled sets field value.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetIsMinimumDurationEnabled(v bool) {
	o.IsMinimumDurationEnabled = v
}

// GetMetricGroups returns the MetricGroups field value.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetMetricGroups() []ExperimentsProtocolMetricGroupInput {
	if o == nil {
		var ret []ExperimentsProtocolMetricGroupInput
		return ret
	}
	return o.MetricGroups
}

// GetMetricGroupsOk returns a tuple with the MetricGroups field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetMetricGroupsOk() (*[]ExperimentsProtocolMetricGroupInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricGroups, true
}

// SetMetricGroups sets field value.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetMetricGroups(v []ExperimentsProtocolMetricGroupInput) {
	o.MetricGroups = v
}

// GetMetricLimit returns the MetricLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetMetricLimit() int64 {
	if o == nil || o.MetricLimit.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MetricLimit.Get()
}

// GetMetricLimitOk returns a tuple with the MetricLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetMetricLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricLimit.Get(), o.MetricLimit.IsSet()
}

// HasMetricLimit returns a boolean if a field has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) HasMetricLimit() bool {
	return o != nil && o.MetricLimit.IsSet()
}

// SetMetricLimit gets a reference to the given datadog.NullableInt64 and assigns it to the MetricLimit field.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetMetricLimit(v int64) {
	o.MetricLimit.Set(&v)
}

// SetMetricLimitNil sets the value for MetricLimit to be an explicit nil.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetMetricLimitNil() {
	o.MetricLimit.Set(nil)
}

// UnsetMetricLimit ensures that no value is present for MetricLimit, not even an explicit nil.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) UnsetMetricLimit() {
	o.MetricLimit.Unset()
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetMigrationMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetMigrationMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given map[string]interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetMigrationMetadata(v map[string]interface{}) {
	o.MigrationMetadata = v
}

// GetMinimumDurationUnit returns the MinimumDurationUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetMinimumDurationUnit() string {
	if o == nil || o.MinimumDurationUnit.Get() == nil {
		var ret string
		return ret
	}
	return *o.MinimumDurationUnit.Get()
}

// GetMinimumDurationUnitOk returns a tuple with the MinimumDurationUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetMinimumDurationUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinimumDurationUnit.Get(), o.MinimumDurationUnit.IsSet()
}

// HasMinimumDurationUnit returns a boolean if a field has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) HasMinimumDurationUnit() bool {
	return o != nil && o.MinimumDurationUnit.IsSet()
}

// SetMinimumDurationUnit gets a reference to the given datadog.NullableString and assigns it to the MinimumDurationUnit field.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetMinimumDurationUnit(v string) {
	o.MinimumDurationUnit.Set(&v)
}

// SetMinimumDurationUnitNil sets the value for MinimumDurationUnit to be an explicit nil.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetMinimumDurationUnitNil() {
	o.MinimumDurationUnit.Set(nil)
}

// UnsetMinimumDurationUnit ensures that no value is present for MinimumDurationUnit, not even an explicit nil.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) UnsetMinimumDurationUnit() {
	o.MinimumDurationUnit.Unset()
}

// GetMinimumDurationValue returns the MinimumDurationValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetMinimumDurationValue() int64 {
	if o == nil || o.MinimumDurationValue.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MinimumDurationValue.Get()
}

// GetMinimumDurationValueOk returns a tuple with the MinimumDurationValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetMinimumDurationValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinimumDurationValue.Get(), o.MinimumDurationValue.IsSet()
}

// HasMinimumDurationValue returns a boolean if a field has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) HasMinimumDurationValue() bool {
	return o != nil && o.MinimumDurationValue.IsSet()
}

// SetMinimumDurationValue gets a reference to the given datadog.NullableInt64 and assigns it to the MinimumDurationValue field.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetMinimumDurationValue(v int64) {
	o.MinimumDurationValue.Set(&v)
}

// SetMinimumDurationValueNil sets the value for MinimumDurationValue to be an explicit nil.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetMinimumDurationValueNil() {
	o.MinimumDurationValue.Set(nil)
}

// UnsetMinimumDurationValue ensures that no value is present for MinimumDurationValue, not even an explicit nil.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) UnsetMinimumDurationValue() {
	o.MinimumDurationValue.Unset()
}

// GetPrimaryMetricId returns the PrimaryMetricId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetPrimaryMetricId() uuid.UUID {
	if o == nil || o.PrimaryMetricId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.PrimaryMetricId.Get()
}

// GetPrimaryMetricIdOk returns a tuple with the PrimaryMetricId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetPrimaryMetricIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryMetricId.Get(), o.PrimaryMetricId.IsSet()
}

// HasPrimaryMetricId returns a boolean if a field has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) HasPrimaryMetricId() bool {
	return o != nil && o.PrimaryMetricId.IsSet()
}

// SetPrimaryMetricId gets a reference to the given datadog.NullableUUID and assigns it to the PrimaryMetricId field.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetPrimaryMetricId(v uuid.UUID) {
	o.PrimaryMetricId.Set(&v)
}

// SetPrimaryMetricIdNil sets the value for PrimaryMetricId to be an explicit nil.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetPrimaryMetricIdNil() {
	o.PrimaryMetricId.Set(nil)
}

// UnsetPrimaryMetricId ensures that no value is present for PrimaryMetricId, not even an explicit nil.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) UnsetPrimaryMetricId() {
	o.PrimaryMetricId.Unset()
}

// GetSubjectTypeId returns the SubjectTypeId field value.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetSubjectTypeId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.SubjectTypeId
}

// GetSubjectTypeIdOk returns a tuple with the SubjectTypeId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetSubjectTypeIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectTypeId, true
}

// SetSubjectTypeId sets field value.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetSubjectTypeId(v uuid.UUID) {
	o.SubjectTypeId = v
}

// GetTargetingRules returns the TargetingRules field value if set, zero value otherwise.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetTargetingRules() []ExperimentsProtocolTargetingRuleInput {
	if o == nil || o.TargetingRules == nil {
		var ret []ExperimentsProtocolTargetingRuleInput
		return ret
	}
	return o.TargetingRules
}

// GetTargetingRulesOk returns a tuple with the TargetingRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) GetTargetingRulesOk() (*[]ExperimentsProtocolTargetingRuleInput, bool) {
	if o == nil || o.TargetingRules == nil {
		return nil, false
	}
	return &o.TargetingRules, true
}

// HasTargetingRules returns a boolean if a field has been set.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) HasTargetingRules() bool {
	return o != nil && o.TargetingRules != nil
}

// SetTargetingRules gets a reference to the given []ExperimentsProtocolTargetingRuleInput and assigns it to the TargetingRules field.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) SetTargetingRules(v []ExperimentsProtocolTargetingRuleInput) {
	o.TargetingRules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsReplaceExperimentProtocolRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AnalysisPlan.IsSet() {
		toSerialize["analysis_plan"] = o.AnalysisPlan.Get()
	}
	if o.AssignmentSourceDefaultProperties != nil {
		toSerialize["assignment_source_default_properties"] = o.AssignmentSourceDefaultProperties
	}
	if o.AssignmentSourceId.IsSet() {
		toSerialize["assignment_source_id"] = o.AssignmentSourceId.Get()
	}
	if o.DefaultDurationDays.IsSet() {
		toSerialize["default_duration_days"] = o.DefaultDurationDays.Get()
	}
	if o.Enforcement != nil {
		toSerialize["enforcement"] = o.Enforcement
	}
	if o.EnvironmentId.IsSet() {
		toSerialize["environment_id"] = o.EnvironmentId.Get()
	}
	if o.ExposureSchedule.IsSet() {
		toSerialize["exposure_schedule"] = o.ExposureSchedule.Get()
	}
	toSerialize["is_duration_required_to_start"] = o.IsDurationRequiredToStart
	toSerialize["is_equal_split_enforced"] = o.IsEqualSplitEnforced
	toSerialize["is_metric_limit_enabled"] = o.IsMetricLimitEnabled
	toSerialize["is_minimum_duration_enabled"] = o.IsMinimumDurationEnabled
	toSerialize["metric_groups"] = o.MetricGroups
	if o.MetricLimit.IsSet() {
		toSerialize["metric_limit"] = o.MetricLimit.Get()
	}
	if o.MigrationMetadata != nil {
		toSerialize["migration_metadata"] = o.MigrationMetadata
	}
	if o.MinimumDurationUnit.IsSet() {
		toSerialize["minimum_duration_unit"] = o.MinimumDurationUnit.Get()
	}
	if o.MinimumDurationValue.IsSet() {
		toSerialize["minimum_duration_value"] = o.MinimumDurationValue.Get()
	}
	if o.PrimaryMetricId.IsSet() {
		toSerialize["primary_metric_id"] = o.PrimaryMetricId.Get()
	}
	toSerialize["subject_type_id"] = o.SubjectTypeId
	if o.TargetingRules != nil {
		toSerialize["targeting_rules"] = o.TargetingRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsReplaceExperimentProtocolRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnalysisPlan                      NullableExperimentsProtocolAnalysisPlanInput                                            `json:"analysis_plan,omitempty"`
		AssignmentSourceDefaultProperties []ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems `json:"assignment_source_default_properties,omitempty"`
		AssignmentSourceId                datadog.NullableUUID                                                                    `json:"assignment_source_id,omitempty"`
		DefaultDurationDays               datadog.NullableInt64                                                                   `json:"default_duration_days,omitempty"`
		Enforcement                       *ExperimentsReplaceProtocolEnforcementInput                                             `json:"enforcement,omitempty"`
		EnvironmentId                     datadog.NullableUUID                                                                    `json:"environment_id,omitempty"`
		ExposureSchedule                  NullableExperimentsProtocolExposureScheduleInput                                        `json:"exposure_schedule,omitempty"`
		IsDurationRequiredToStart         *bool                                                                                   `json:"is_duration_required_to_start"`
		IsEqualSplitEnforced              *bool                                                                                   `json:"is_equal_split_enforced"`
		IsMetricLimitEnabled              *bool                                                                                   `json:"is_metric_limit_enabled"`
		IsMinimumDurationEnabled          *bool                                                                                   `json:"is_minimum_duration_enabled"`
		MetricGroups                      *[]ExperimentsProtocolMetricGroupInput                                                  `json:"metric_groups"`
		MetricLimit                       datadog.NullableInt64                                                                   `json:"metric_limit,omitempty"`
		MigrationMetadata                 map[string]interface{}                                                                  `json:"migration_metadata,omitempty"`
		MinimumDurationUnit               datadog.NullableString                                                                  `json:"minimum_duration_unit,omitempty"`
		MinimumDurationValue              datadog.NullableInt64                                                                   `json:"minimum_duration_value,omitempty"`
		PrimaryMetricId                   datadog.NullableUUID                                                                    `json:"primary_metric_id,omitempty"`
		SubjectTypeId                     *uuid.UUID                                                                              `json:"subject_type_id"`
		TargetingRules                    []ExperimentsProtocolTargetingRuleInput                                                 `json:"targeting_rules,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IsDurationRequiredToStart == nil {
		return fmt.Errorf("required field is_duration_required_to_start missing")
	}
	if all.IsEqualSplitEnforced == nil {
		return fmt.Errorf("required field is_equal_split_enforced missing")
	}
	if all.IsMetricLimitEnabled == nil {
		return fmt.Errorf("required field is_metric_limit_enabled missing")
	}
	if all.IsMinimumDurationEnabled == nil {
		return fmt.Errorf("required field is_minimum_duration_enabled missing")
	}
	if all.MetricGroups == nil {
		return fmt.Errorf("required field metric_groups missing")
	}
	if all.SubjectTypeId == nil {
		return fmt.Errorf("required field subject_type_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"analysis_plan", "assignment_source_default_properties", "assignment_source_id", "default_duration_days", "enforcement", "environment_id", "exposure_schedule", "is_duration_required_to_start", "is_equal_split_enforced", "is_metric_limit_enabled", "is_minimum_duration_enabled", "metric_groups", "metric_limit", "migration_metadata", "minimum_duration_unit", "minimum_duration_value", "primary_metric_id", "subject_type_id", "targeting_rules"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AnalysisPlan = all.AnalysisPlan
	o.AssignmentSourceDefaultProperties = all.AssignmentSourceDefaultProperties
	o.AssignmentSourceId = all.AssignmentSourceId
	o.DefaultDurationDays = all.DefaultDurationDays
	if all.Enforcement != nil && all.Enforcement.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Enforcement = all.Enforcement
	o.EnvironmentId = all.EnvironmentId
	o.ExposureSchedule = all.ExposureSchedule
	o.IsDurationRequiredToStart = *all.IsDurationRequiredToStart
	o.IsEqualSplitEnforced = *all.IsEqualSplitEnforced
	o.IsMetricLimitEnabled = *all.IsMetricLimitEnabled
	o.IsMinimumDurationEnabled = *all.IsMinimumDurationEnabled
	o.MetricGroups = *all.MetricGroups
	o.MetricLimit = all.MetricLimit
	o.MigrationMetadata = all.MigrationMetadata
	o.MinimumDurationUnit = all.MinimumDurationUnit
	o.MinimumDurationValue = all.MinimumDurationValue
	o.PrimaryMetricId = all.PrimaryMetricId
	o.SubjectTypeId = *all.SubjectTypeId
	o.TargetingRules = all.TargetingRules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
