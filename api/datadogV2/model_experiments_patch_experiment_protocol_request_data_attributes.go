// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPatchExperimentProtocolRequestDataAttributes Partial protocol update. Omitted attributes remain unchanged. Published protocols accept only name, description, and migration_metadata; archived protocols reject updates.
type ExperimentsPatchExperimentProtocolRequestDataAttributes struct {
	// Partial statistical settings. Omitted fields remain unchanged. Null clears the whole plan. Null clears nullable fields; null or an empty string clears confidence_interval_method.
	AnalysisPlan NullableExperimentsPatchProtocolAnalysisPlanInput `json:"analysis_plan,omitempty"`
	// Assignment properties. Omission or null preserves them; an array replaces them.
	AssignmentSourceDefaultProperties []ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems `json:"assignment_source_default_properties,omitempty"`
	// ID of the assignment source selected by the protocol. Omission or null keeps the current value. An empty string clears it.
	AssignmentSourceId datadog.NullableString `json:"assignment_source_id,omitempty"`
	// Default experiment duration supplied by the protocol, in days. Omission preserves it; null clears it. Sets analysis_plan.target_duration_days on experiments created from this protocol.
	DefaultDurationDays datadog.NullableInt64 `json:"default_duration_days,omitempty"`
	// Description. Empty string clears it; omission or null preserves it.
	Description datadog.NullableString `json:"description,omitempty"`
	// Partial lock map. Omitted fields remain unchanged. Null clears the whole map. Empty or null field values become `EDITABLE`, except `subject_type`, `flag_source`, and `confidence_interval_method`, which remain `LOCKED`.
	Enforcement NullableExperimentsPatchProtocolEnforcementInput `json:"enforcement,omitempty"`
	// ID of the feature flag environment used by the experiment. Omission or null keeps the current value. An empty string clears it.
	EnvironmentId datadog.NullableString `json:"environment_id,omitempty"`
	// Partial traffic schedule. Omitted fields remain unchanged. Null clears the whole schedule. Null clears nullable fields and rollout_steps; null sets strategy to an empty string and `autostart` to false.
	ExposureSchedule NullableExperimentsPatchProtocolExposureScheduleInput `json:"exposure_schedule,omitempty"`
	// Whether the protocol requires a duration before an experiment can start. Omission or null preserves it.
	IsDurationRequiredToStart datadog.NullableBool `json:"is_duration_required_to_start,omitempty"`
	// Whether the protocol requires equal traffic allocation across variants. Omission or null preserves it.
	IsEqualSplitEnforced datadog.NullableBool `json:"is_equal_split_enforced,omitempty"`
	// Whether the protocol limits the number of metrics. Omission or null preserves it.
	IsMetricLimitEnabled datadog.NullableBool `json:"is_metric_limit_enabled,omitempty"`
	// Whether the protocol enforces a minimum experiment duration. Omission or null preserves it.
	IsMinimumDurationEnabled datadog.NullableBool `json:"is_minimum_duration_enabled,omitempty"`
	// Metric groups. Omission or null preserves them; an array replaces them.
	MetricGroups []ExperimentsProtocolMetricGroupInput `json:"metric_groups,omitempty"`
	// Maximum number of metrics allowed by the protocol. Omission preserves it; null clears it.
	MetricLimit datadog.NullableInt64 `json:"metric_limit,omitempty"`
	// Metadata for imported resources. Must encode to at most 4096 bytes. Null clears it; omission preserves it on update.
	MigrationMetadata map[string]interface{} `json:"migration_metadata,omitempty"`
	// Minimum duration unit, `DAYS` or `WEEKS`. Empty string clears it. Omission or null preserves it.
	MinimumDurationUnit datadog.NullableString `json:"minimum_duration_unit,omitempty"`
	// Minimum experiment duration in the specified unit. Omission preserves it; null clears it.
	MinimumDurationValue datadog.NullableInt64 `json:"minimum_duration_value,omitempty"`
	// Protocol name. Leading and trailing spaces are removed. The trimmed name must contain 1 to 255 characters.
	Name datadog.NullableString `json:"name,omitempty"`
	// ID of the primary metric supplied by the protocol. Omission or null keeps the current value. An empty string clears it.
	PrimaryMetricId datadog.NullableString `json:"primary_metric_id,omitempty"`
	// ID of the subject type used by this configuration. Omission or null keeps the current value. An empty string clears it.
	SubjectTypeId datadog.NullableString `json:"subject_type_id,omitempty"`
	// Targeting rules. Omission or null preserves them; an array replaces them.
	TargetingRules []ExperimentsProtocolTargetingRuleInput `json:"targeting_rules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPatchExperimentProtocolRequestDataAttributes instantiates a new ExperimentsPatchExperimentProtocolRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPatchExperimentProtocolRequestDataAttributes() *ExperimentsPatchExperimentProtocolRequestDataAttributes {
	this := ExperimentsPatchExperimentProtocolRequestDataAttributes{}
	return &this
}

// NewExperimentsPatchExperimentProtocolRequestDataAttributesWithDefaults instantiates a new ExperimentsPatchExperimentProtocolRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPatchExperimentProtocolRequestDataAttributesWithDefaults() *ExperimentsPatchExperimentProtocolRequestDataAttributes {
	this := ExperimentsPatchExperimentProtocolRequestDataAttributes{}
	return &this
}

// GetAnalysisPlan returns the AnalysisPlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetAnalysisPlan() ExperimentsPatchProtocolAnalysisPlanInput {
	if o == nil || o.AnalysisPlan.Get() == nil {
		var ret ExperimentsPatchProtocolAnalysisPlanInput
		return ret
	}
	return *o.AnalysisPlan.Get()
}

// GetAnalysisPlanOk returns a tuple with the AnalysisPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetAnalysisPlanOk() (*ExperimentsPatchProtocolAnalysisPlanInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnalysisPlan.Get(), o.AnalysisPlan.IsSet()
}

// HasAnalysisPlan returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasAnalysisPlan() bool {
	return o != nil && o.AnalysisPlan.IsSet()
}

// SetAnalysisPlan gets a reference to the given NullableExperimentsPatchProtocolAnalysisPlanInput and assigns it to the AnalysisPlan field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetAnalysisPlan(v ExperimentsPatchProtocolAnalysisPlanInput) {
	o.AnalysisPlan.Set(&v)
}

// SetAnalysisPlanNil sets the value for AnalysisPlan to be an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetAnalysisPlanNil() {
	o.AnalysisPlan.Set(nil)
}

// UnsetAnalysisPlan ensures that no value is present for AnalysisPlan, not even an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) UnsetAnalysisPlan() {
	o.AnalysisPlan.Unset()
}

// GetAssignmentSourceDefaultProperties returns the AssignmentSourceDefaultProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetAssignmentSourceDefaultProperties() []ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems {
	if o == nil {
		var ret []ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems
		return ret
	}
	return o.AssignmentSourceDefaultProperties
}

// GetAssignmentSourceDefaultPropertiesOk returns a tuple with the AssignmentSourceDefaultProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetAssignmentSourceDefaultPropertiesOk() (*[]ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems, bool) {
	if o == nil || o.AssignmentSourceDefaultProperties == nil {
		return nil, false
	}
	return &o.AssignmentSourceDefaultProperties, true
}

// HasAssignmentSourceDefaultProperties returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasAssignmentSourceDefaultProperties() bool {
	return o != nil && o.AssignmentSourceDefaultProperties != nil
}

// SetAssignmentSourceDefaultProperties gets a reference to the given []ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems and assigns it to the AssignmentSourceDefaultProperties field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetAssignmentSourceDefaultProperties(v []ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems) {
	o.AssignmentSourceDefaultProperties = v
}

// GetAssignmentSourceId returns the AssignmentSourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetAssignmentSourceId() string {
	if o == nil || o.AssignmentSourceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AssignmentSourceId.Get()
}

// GetAssignmentSourceIdOk returns a tuple with the AssignmentSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetAssignmentSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignmentSourceId.Get(), o.AssignmentSourceId.IsSet()
}

// HasAssignmentSourceId returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasAssignmentSourceId() bool {
	return o != nil && o.AssignmentSourceId.IsSet()
}

// SetAssignmentSourceId gets a reference to the given datadog.NullableString and assigns it to the AssignmentSourceId field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetAssignmentSourceId(v string) {
	o.AssignmentSourceId.Set(&v)
}

// SetAssignmentSourceIdNil sets the value for AssignmentSourceId to be an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetAssignmentSourceIdNil() {
	o.AssignmentSourceId.Set(nil)
}

// UnsetAssignmentSourceId ensures that no value is present for AssignmentSourceId, not even an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) UnsetAssignmentSourceId() {
	o.AssignmentSourceId.Unset()
}

// GetDefaultDurationDays returns the DefaultDurationDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetDefaultDurationDays() int64 {
	if o == nil || o.DefaultDurationDays.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DefaultDurationDays.Get()
}

// GetDefaultDurationDaysOk returns a tuple with the DefaultDurationDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetDefaultDurationDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultDurationDays.Get(), o.DefaultDurationDays.IsSet()
}

// HasDefaultDurationDays returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasDefaultDurationDays() bool {
	return o != nil && o.DefaultDurationDays.IsSet()
}

// SetDefaultDurationDays gets a reference to the given datadog.NullableInt64 and assigns it to the DefaultDurationDays field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetDefaultDurationDays(v int64) {
	o.DefaultDurationDays.Set(&v)
}

// SetDefaultDurationDaysNil sets the value for DefaultDurationDays to be an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetDefaultDurationDaysNil() {
	o.DefaultDurationDays.Set(nil)
}

// UnsetDefaultDurationDays ensures that no value is present for DefaultDurationDays, not even an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) UnsetDefaultDurationDays() {
	o.DefaultDurationDays.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetEnforcement returns the Enforcement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetEnforcement() ExperimentsPatchProtocolEnforcementInput {
	if o == nil || o.Enforcement.Get() == nil {
		var ret ExperimentsPatchProtocolEnforcementInput
		return ret
	}
	return *o.Enforcement.Get()
}

// GetEnforcementOk returns a tuple with the Enforcement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetEnforcementOk() (*ExperimentsPatchProtocolEnforcementInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enforcement.Get(), o.Enforcement.IsSet()
}

// HasEnforcement returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasEnforcement() bool {
	return o != nil && o.Enforcement.IsSet()
}

// SetEnforcement gets a reference to the given NullableExperimentsPatchProtocolEnforcementInput and assigns it to the Enforcement field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetEnforcement(v ExperimentsPatchProtocolEnforcementInput) {
	o.Enforcement.Set(&v)
}

// SetEnforcementNil sets the value for Enforcement to be an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetEnforcementNil() {
	o.Enforcement.Set(nil)
}

// UnsetEnforcement ensures that no value is present for Enforcement, not even an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) UnsetEnforcement() {
	o.Enforcement.Unset()
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId.Get()
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvironmentId.Get(), o.EnvironmentId.IsSet()
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId.IsSet()
}

// SetEnvironmentId gets a reference to the given datadog.NullableString and assigns it to the EnvironmentId field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetEnvironmentId(v string) {
	o.EnvironmentId.Set(&v)
}

// SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetEnvironmentIdNil() {
	o.EnvironmentId.Set(nil)
}

// UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) UnsetEnvironmentId() {
	o.EnvironmentId.Unset()
}

// GetExposureSchedule returns the ExposureSchedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetExposureSchedule() ExperimentsPatchProtocolExposureScheduleInput {
	if o == nil || o.ExposureSchedule.Get() == nil {
		var ret ExperimentsPatchProtocolExposureScheduleInput
		return ret
	}
	return *o.ExposureSchedule.Get()
}

// GetExposureScheduleOk returns a tuple with the ExposureSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetExposureScheduleOk() (*ExperimentsPatchProtocolExposureScheduleInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExposureSchedule.Get(), o.ExposureSchedule.IsSet()
}

// HasExposureSchedule returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasExposureSchedule() bool {
	return o != nil && o.ExposureSchedule.IsSet()
}

// SetExposureSchedule gets a reference to the given NullableExperimentsPatchProtocolExposureScheduleInput and assigns it to the ExposureSchedule field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetExposureSchedule(v ExperimentsPatchProtocolExposureScheduleInput) {
	o.ExposureSchedule.Set(&v)
}

// SetExposureScheduleNil sets the value for ExposureSchedule to be an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetExposureScheduleNil() {
	o.ExposureSchedule.Set(nil)
}

// UnsetExposureSchedule ensures that no value is present for ExposureSchedule, not even an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) UnsetExposureSchedule() {
	o.ExposureSchedule.Unset()
}

// GetIsDurationRequiredToStart returns the IsDurationRequiredToStart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetIsDurationRequiredToStart() bool {
	if o == nil || o.IsDurationRequiredToStart.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsDurationRequiredToStart.Get()
}

// GetIsDurationRequiredToStartOk returns a tuple with the IsDurationRequiredToStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetIsDurationRequiredToStartOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDurationRequiredToStart.Get(), o.IsDurationRequiredToStart.IsSet()
}

// HasIsDurationRequiredToStart returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasIsDurationRequiredToStart() bool {
	return o != nil && o.IsDurationRequiredToStart.IsSet()
}

// SetIsDurationRequiredToStart gets a reference to the given datadog.NullableBool and assigns it to the IsDurationRequiredToStart field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetIsDurationRequiredToStart(v bool) {
	o.IsDurationRequiredToStart.Set(&v)
}

// SetIsDurationRequiredToStartNil sets the value for IsDurationRequiredToStart to be an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetIsDurationRequiredToStartNil() {
	o.IsDurationRequiredToStart.Set(nil)
}

// UnsetIsDurationRequiredToStart ensures that no value is present for IsDurationRequiredToStart, not even an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) UnsetIsDurationRequiredToStart() {
	o.IsDurationRequiredToStart.Unset()
}

// GetIsEqualSplitEnforced returns the IsEqualSplitEnforced field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetIsEqualSplitEnforced() bool {
	if o == nil || o.IsEqualSplitEnforced.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsEqualSplitEnforced.Get()
}

// GetIsEqualSplitEnforcedOk returns a tuple with the IsEqualSplitEnforced field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetIsEqualSplitEnforcedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsEqualSplitEnforced.Get(), o.IsEqualSplitEnforced.IsSet()
}

// HasIsEqualSplitEnforced returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasIsEqualSplitEnforced() bool {
	return o != nil && o.IsEqualSplitEnforced.IsSet()
}

// SetIsEqualSplitEnforced gets a reference to the given datadog.NullableBool and assigns it to the IsEqualSplitEnforced field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetIsEqualSplitEnforced(v bool) {
	o.IsEqualSplitEnforced.Set(&v)
}

// SetIsEqualSplitEnforcedNil sets the value for IsEqualSplitEnforced to be an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetIsEqualSplitEnforcedNil() {
	o.IsEqualSplitEnforced.Set(nil)
}

// UnsetIsEqualSplitEnforced ensures that no value is present for IsEqualSplitEnforced, not even an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) UnsetIsEqualSplitEnforced() {
	o.IsEqualSplitEnforced.Unset()
}

// GetIsMetricLimitEnabled returns the IsMetricLimitEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetIsMetricLimitEnabled() bool {
	if o == nil || o.IsMetricLimitEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsMetricLimitEnabled.Get()
}

// GetIsMetricLimitEnabledOk returns a tuple with the IsMetricLimitEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetIsMetricLimitEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsMetricLimitEnabled.Get(), o.IsMetricLimitEnabled.IsSet()
}

// HasIsMetricLimitEnabled returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasIsMetricLimitEnabled() bool {
	return o != nil && o.IsMetricLimitEnabled.IsSet()
}

// SetIsMetricLimitEnabled gets a reference to the given datadog.NullableBool and assigns it to the IsMetricLimitEnabled field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetIsMetricLimitEnabled(v bool) {
	o.IsMetricLimitEnabled.Set(&v)
}

// SetIsMetricLimitEnabledNil sets the value for IsMetricLimitEnabled to be an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetIsMetricLimitEnabledNil() {
	o.IsMetricLimitEnabled.Set(nil)
}

// UnsetIsMetricLimitEnabled ensures that no value is present for IsMetricLimitEnabled, not even an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) UnsetIsMetricLimitEnabled() {
	o.IsMetricLimitEnabled.Unset()
}

// GetIsMinimumDurationEnabled returns the IsMinimumDurationEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetIsMinimumDurationEnabled() bool {
	if o == nil || o.IsMinimumDurationEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsMinimumDurationEnabled.Get()
}

// GetIsMinimumDurationEnabledOk returns a tuple with the IsMinimumDurationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetIsMinimumDurationEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsMinimumDurationEnabled.Get(), o.IsMinimumDurationEnabled.IsSet()
}

// HasIsMinimumDurationEnabled returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasIsMinimumDurationEnabled() bool {
	return o != nil && o.IsMinimumDurationEnabled.IsSet()
}

// SetIsMinimumDurationEnabled gets a reference to the given datadog.NullableBool and assigns it to the IsMinimumDurationEnabled field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetIsMinimumDurationEnabled(v bool) {
	o.IsMinimumDurationEnabled.Set(&v)
}

// SetIsMinimumDurationEnabledNil sets the value for IsMinimumDurationEnabled to be an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetIsMinimumDurationEnabledNil() {
	o.IsMinimumDurationEnabled.Set(nil)
}

// UnsetIsMinimumDurationEnabled ensures that no value is present for IsMinimumDurationEnabled, not even an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) UnsetIsMinimumDurationEnabled() {
	o.IsMinimumDurationEnabled.Unset()
}

// GetMetricGroups returns the MetricGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetMetricGroups() []ExperimentsProtocolMetricGroupInput {
	if o == nil {
		var ret []ExperimentsProtocolMetricGroupInput
		return ret
	}
	return o.MetricGroups
}

// GetMetricGroupsOk returns a tuple with the MetricGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetMetricGroupsOk() (*[]ExperimentsProtocolMetricGroupInput, bool) {
	if o == nil || o.MetricGroups == nil {
		return nil, false
	}
	return &o.MetricGroups, true
}

// HasMetricGroups returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasMetricGroups() bool {
	return o != nil && o.MetricGroups != nil
}

// SetMetricGroups gets a reference to the given []ExperimentsProtocolMetricGroupInput and assigns it to the MetricGroups field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetMetricGroups(v []ExperimentsProtocolMetricGroupInput) {
	o.MetricGroups = v
}

// GetMetricLimit returns the MetricLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetMetricLimit() int64 {
	if o == nil || o.MetricLimit.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MetricLimit.Get()
}

// GetMetricLimitOk returns a tuple with the MetricLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetMetricLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricLimit.Get(), o.MetricLimit.IsSet()
}

// HasMetricLimit returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasMetricLimit() bool {
	return o != nil && o.MetricLimit.IsSet()
}

// SetMetricLimit gets a reference to the given datadog.NullableInt64 and assigns it to the MetricLimit field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetMetricLimit(v int64) {
	o.MetricLimit.Set(&v)
}

// SetMetricLimitNil sets the value for MetricLimit to be an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetMetricLimitNil() {
	o.MetricLimit.Set(nil)
}

// UnsetMetricLimit ensures that no value is present for MetricLimit, not even an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) UnsetMetricLimit() {
	o.MetricLimit.Unset()
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetMigrationMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetMigrationMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given map[string]interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetMigrationMetadata(v map[string]interface{}) {
	o.MigrationMetadata = v
}

// GetMinimumDurationUnit returns the MinimumDurationUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetMinimumDurationUnit() string {
	if o == nil || o.MinimumDurationUnit.Get() == nil {
		var ret string
		return ret
	}
	return *o.MinimumDurationUnit.Get()
}

// GetMinimumDurationUnitOk returns a tuple with the MinimumDurationUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetMinimumDurationUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinimumDurationUnit.Get(), o.MinimumDurationUnit.IsSet()
}

// HasMinimumDurationUnit returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasMinimumDurationUnit() bool {
	return o != nil && o.MinimumDurationUnit.IsSet()
}

// SetMinimumDurationUnit gets a reference to the given datadog.NullableString and assigns it to the MinimumDurationUnit field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetMinimumDurationUnit(v string) {
	o.MinimumDurationUnit.Set(&v)
}

// SetMinimumDurationUnitNil sets the value for MinimumDurationUnit to be an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetMinimumDurationUnitNil() {
	o.MinimumDurationUnit.Set(nil)
}

// UnsetMinimumDurationUnit ensures that no value is present for MinimumDurationUnit, not even an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) UnsetMinimumDurationUnit() {
	o.MinimumDurationUnit.Unset()
}

// GetMinimumDurationValue returns the MinimumDurationValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetMinimumDurationValue() int64 {
	if o == nil || o.MinimumDurationValue.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MinimumDurationValue.Get()
}

// GetMinimumDurationValueOk returns a tuple with the MinimumDurationValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetMinimumDurationValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinimumDurationValue.Get(), o.MinimumDurationValue.IsSet()
}

// HasMinimumDurationValue returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasMinimumDurationValue() bool {
	return o != nil && o.MinimumDurationValue.IsSet()
}

// SetMinimumDurationValue gets a reference to the given datadog.NullableInt64 and assigns it to the MinimumDurationValue field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetMinimumDurationValue(v int64) {
	o.MinimumDurationValue.Set(&v)
}

// SetMinimumDurationValueNil sets the value for MinimumDurationValue to be an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetMinimumDurationValueNil() {
	o.MinimumDurationValue.Set(nil)
}

// UnsetMinimumDurationValue ensures that no value is present for MinimumDurationValue, not even an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) UnsetMinimumDurationValue() {
	o.MinimumDurationValue.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given datadog.NullableString and assigns it to the Name field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) UnsetName() {
	o.Name.Unset()
}

// GetPrimaryMetricId returns the PrimaryMetricId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetPrimaryMetricId() string {
	if o == nil || o.PrimaryMetricId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrimaryMetricId.Get()
}

// GetPrimaryMetricIdOk returns a tuple with the PrimaryMetricId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetPrimaryMetricIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryMetricId.Get(), o.PrimaryMetricId.IsSet()
}

// HasPrimaryMetricId returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasPrimaryMetricId() bool {
	return o != nil && o.PrimaryMetricId.IsSet()
}

// SetPrimaryMetricId gets a reference to the given datadog.NullableString and assigns it to the PrimaryMetricId field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetPrimaryMetricId(v string) {
	o.PrimaryMetricId.Set(&v)
}

// SetPrimaryMetricIdNil sets the value for PrimaryMetricId to be an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetPrimaryMetricIdNil() {
	o.PrimaryMetricId.Set(nil)
}

// UnsetPrimaryMetricId ensures that no value is present for PrimaryMetricId, not even an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) UnsetPrimaryMetricId() {
	o.PrimaryMetricId.Unset()
}

// GetSubjectTypeId returns the SubjectTypeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetSubjectTypeId() string {
	if o == nil || o.SubjectTypeId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubjectTypeId.Get()
}

// GetSubjectTypeIdOk returns a tuple with the SubjectTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetSubjectTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubjectTypeId.Get(), o.SubjectTypeId.IsSet()
}

// HasSubjectTypeId returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasSubjectTypeId() bool {
	return o != nil && o.SubjectTypeId.IsSet()
}

// SetSubjectTypeId gets a reference to the given datadog.NullableString and assigns it to the SubjectTypeId field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetSubjectTypeId(v string) {
	o.SubjectTypeId.Set(&v)
}

// SetSubjectTypeIdNil sets the value for SubjectTypeId to be an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetSubjectTypeIdNil() {
	o.SubjectTypeId.Set(nil)
}

// UnsetSubjectTypeId ensures that no value is present for SubjectTypeId, not even an explicit nil.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) UnsetSubjectTypeId() {
	o.SubjectTypeId.Unset()
}

// GetTargetingRules returns the TargetingRules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetTargetingRules() []ExperimentsProtocolTargetingRuleInput {
	if o == nil {
		var ret []ExperimentsProtocolTargetingRuleInput
		return ret
	}
	return o.TargetingRules
}

// GetTargetingRulesOk returns a tuple with the TargetingRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) GetTargetingRulesOk() (*[]ExperimentsProtocolTargetingRuleInput, bool) {
	if o == nil || o.TargetingRules == nil {
		return nil, false
	}
	return &o.TargetingRules, true
}

// HasTargetingRules returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) HasTargetingRules() bool {
	return o != nil && o.TargetingRules != nil
}

// SetTargetingRules gets a reference to the given []ExperimentsProtocolTargetingRuleInput and assigns it to the TargetingRules field.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) SetTargetingRules(v []ExperimentsProtocolTargetingRuleInput) {
	o.TargetingRules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPatchExperimentProtocolRequestDataAttributes) MarshalJSON() ([]byte, error) {
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
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Enforcement.IsSet() {
		toSerialize["enforcement"] = o.Enforcement.Get()
	}
	if o.EnvironmentId.IsSet() {
		toSerialize["environment_id"] = o.EnvironmentId.Get()
	}
	if o.ExposureSchedule.IsSet() {
		toSerialize["exposure_schedule"] = o.ExposureSchedule.Get()
	}
	if o.IsDurationRequiredToStart.IsSet() {
		toSerialize["is_duration_required_to_start"] = o.IsDurationRequiredToStart.Get()
	}
	if o.IsEqualSplitEnforced.IsSet() {
		toSerialize["is_equal_split_enforced"] = o.IsEqualSplitEnforced.Get()
	}
	if o.IsMetricLimitEnabled.IsSet() {
		toSerialize["is_metric_limit_enabled"] = o.IsMetricLimitEnabled.Get()
	}
	if o.IsMinimumDurationEnabled.IsSet() {
		toSerialize["is_minimum_duration_enabled"] = o.IsMinimumDurationEnabled.Get()
	}
	if o.MetricGroups != nil {
		toSerialize["metric_groups"] = o.MetricGroups
	}
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
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.PrimaryMetricId.IsSet() {
		toSerialize["primary_metric_id"] = o.PrimaryMetricId.Get()
	}
	if o.SubjectTypeId.IsSet() {
		toSerialize["subject_type_id"] = o.SubjectTypeId.Get()
	}
	if o.TargetingRules != nil {
		toSerialize["targeting_rules"] = o.TargetingRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsPatchExperimentProtocolRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnalysisPlan                      NullableExperimentsPatchProtocolAnalysisPlanInput                                       `json:"analysis_plan,omitempty"`
		AssignmentSourceDefaultProperties []ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems `json:"assignment_source_default_properties,omitempty"`
		AssignmentSourceId                datadog.NullableString                                                                  `json:"assignment_source_id,omitempty"`
		DefaultDurationDays               datadog.NullableInt64                                                                   `json:"default_duration_days,omitempty"`
		Description                       datadog.NullableString                                                                  `json:"description,omitempty"`
		Enforcement                       NullableExperimentsPatchProtocolEnforcementInput                                        `json:"enforcement,omitempty"`
		EnvironmentId                     datadog.NullableString                                                                  `json:"environment_id,omitempty"`
		ExposureSchedule                  NullableExperimentsPatchProtocolExposureScheduleInput                                   `json:"exposure_schedule,omitempty"`
		IsDurationRequiredToStart         datadog.NullableBool                                                                    `json:"is_duration_required_to_start,omitempty"`
		IsEqualSplitEnforced              datadog.NullableBool                                                                    `json:"is_equal_split_enforced,omitempty"`
		IsMetricLimitEnabled              datadog.NullableBool                                                                    `json:"is_metric_limit_enabled,omitempty"`
		IsMinimumDurationEnabled          datadog.NullableBool                                                                    `json:"is_minimum_duration_enabled,omitempty"`
		MetricGroups                      []ExperimentsProtocolMetricGroupInput                                                   `json:"metric_groups,omitempty"`
		MetricLimit                       datadog.NullableInt64                                                                   `json:"metric_limit,omitempty"`
		MigrationMetadata                 map[string]interface{}                                                                  `json:"migration_metadata,omitempty"`
		MinimumDurationUnit               datadog.NullableString                                                                  `json:"minimum_duration_unit,omitempty"`
		MinimumDurationValue              datadog.NullableInt64                                                                   `json:"minimum_duration_value,omitempty"`
		Name                              datadog.NullableString                                                                  `json:"name,omitempty"`
		PrimaryMetricId                   datadog.NullableString                                                                  `json:"primary_metric_id,omitempty"`
		SubjectTypeId                     datadog.NullableString                                                                  `json:"subject_type_id,omitempty"`
		TargetingRules                    []ExperimentsProtocolTargetingRuleInput                                                 `json:"targeting_rules,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"analysis_plan", "assignment_source_default_properties", "assignment_source_id", "default_duration_days", "description", "enforcement", "environment_id", "exposure_schedule", "is_duration_required_to_start", "is_equal_split_enforced", "is_metric_limit_enabled", "is_minimum_duration_enabled", "metric_groups", "metric_limit", "migration_metadata", "minimum_duration_unit", "minimum_duration_value", "name", "primary_metric_id", "subject_type_id", "targeting_rules"})
	} else {
		return err
	}
	o.AnalysisPlan = all.AnalysisPlan
	o.AssignmentSourceDefaultProperties = all.AssignmentSourceDefaultProperties
	o.AssignmentSourceId = all.AssignmentSourceId
	o.DefaultDurationDays = all.DefaultDurationDays
	o.Description = all.Description
	o.Enforcement = all.Enforcement
	o.EnvironmentId = all.EnvironmentId
	o.ExposureSchedule = all.ExposureSchedule
	o.IsDurationRequiredToStart = all.IsDurationRequiredToStart
	o.IsEqualSplitEnforced = all.IsEqualSplitEnforced
	o.IsMetricLimitEnabled = all.IsMetricLimitEnabled
	o.IsMinimumDurationEnabled = all.IsMinimumDurationEnabled
	o.MetricGroups = all.MetricGroups
	o.MetricLimit = all.MetricLimit
	o.MigrationMetadata = all.MigrationMetadata
	o.MinimumDurationUnit = all.MinimumDurationUnit
	o.MinimumDurationValue = all.MinimumDurationValue
	o.Name = all.Name
	o.PrimaryMetricId = all.PrimaryMetricId
	o.SubjectTypeId = all.SubjectTypeId
	o.TargetingRules = all.TargetingRules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
