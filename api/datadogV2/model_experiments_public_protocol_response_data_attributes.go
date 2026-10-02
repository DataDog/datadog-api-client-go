// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPublicProtocolResponseDataAttributes Settings and defaults supplied by the protocol.
type ExperimentsPublicProtocolResponseDataAttributes struct {
	// Default statistical settings supplied by the protocol.
	AnalysisPlan *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan `json:"analysis_plan,omitempty"`
	// Default properties supplied by the protocol's assignment source.
	AssignmentSourceDefaultProperties []ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems `json:"assignment_source_default_properties,omitempty"`
	// ID of the assignment source selected by the protocol.
	AssignmentSourceId *string `json:"assignment_source_id,omitempty"`
	// Default experiment duration supplied by the protocol, in days.
	DefaultDurationDays *int64 `json:"default_duration_days,omitempty"`
	// Text that explains the protocol.
	Description *string `json:"description,omitempty"`
	// Controls that determine which protocol settings can be changed in an experiment.
	Enforcement *ExperimentsPublicProtocolResponseDataAttributesEnforcement `json:"enforcement,omitempty"`
	// ID of the feature flag environment used by the experiment.
	EnvironmentId *string `json:"environment_id,omitempty"`
	// Schedule that controls traffic exposure for experiments created from the protocol.
	ExposureSchedule *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule `json:"exposure_schedule,omitempty"`
	// Whether the protocol requires a duration before an experiment can start.
	IsDurationRequiredToStart bool `json:"is_duration_required_to_start"`
	// Whether the protocol requires equal traffic allocation across variants.
	IsEqualSplitEnforced bool `json:"is_equal_split_enforced"`
	// Whether the protocol limits the number of metrics.
	IsMetricLimitEnabled bool `json:"is_metric_limit_enabled"`
	// Whether the protocol enforces a minimum experiment duration.
	IsMinimumDurationEnabled bool `json:"is_minimum_duration_enabled"`
	// Metric groups supplied by the protocol.
	MetricGroups []ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems `json:"metric_groups"`
	// Maximum number of metrics allowed by the protocol.
	MetricLimit *int64 `json:"metric_limit,omitempty"`
	// Metadata retained for resources imported from another system.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// Unit used to express the protocol's minimum duration.
	MinimumDurationUnit *string `json:"minimum_duration_unit,omitempty"`
	// Minimum experiment duration in the specified unit.
	MinimumDurationValue *int64 `json:"minimum_duration_value,omitempty"`
	// Display name of the protocol.
	Name string `json:"name"`
	// Subject type selected by the protocol.
	PrimaryMetric *ExperimentsPublicProtocolResponseDataAttributesSubjectType `json:"primary_metric,omitempty"`
	// ID of the primary metric supplied by the protocol.
	PrimaryMetricId *string `json:"primary_metric_id,omitempty"`
	// Time when the protocol was published.
	PublishedAt *time.Time `json:"published_at,omitempty"`
	// Publication status of the protocol.
	Status ExperimentsPublicProtocolResponseDataAttributesStatus `json:"status"`
	// Subject type selected by the protocol.
	SubjectType *ExperimentsPublicProtocolResponseDataAttributesSubjectType `json:"subject_type,omitempty"`
	// ID of the subject type used by this configuration.
	SubjectTypeId *string `json:"subject_type_id,omitempty"`
	// Rules that select subjects for the experiment.
	TargetingRules []ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems `json:"targeting_rules,omitempty"`
	// RFC3339 update time. Preserve all fractional seconds when passing this value as expected_updated_at.
	UpdatedAt string `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPublicProtocolResponseDataAttributes instantiates a new ExperimentsPublicProtocolResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPublicProtocolResponseDataAttributes(isDurationRequiredToStart bool, isEqualSplitEnforced bool, isMetricLimitEnabled bool, isMinimumDurationEnabled bool, metricGroups []ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems, name string, status ExperimentsPublicProtocolResponseDataAttributesStatus, updatedAt string) *ExperimentsPublicProtocolResponseDataAttributes {
	this := ExperimentsPublicProtocolResponseDataAttributes{}
	this.IsDurationRequiredToStart = isDurationRequiredToStart
	this.IsEqualSplitEnforced = isEqualSplitEnforced
	this.IsMetricLimitEnabled = isMetricLimitEnabled
	this.IsMinimumDurationEnabled = isMinimumDurationEnabled
	this.MetricGroups = metricGroups
	this.Name = name
	this.Status = status
	this.UpdatedAt = updatedAt
	return &this
}

// NewExperimentsPublicProtocolResponseDataAttributesWithDefaults instantiates a new ExperimentsPublicProtocolResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPublicProtocolResponseDataAttributesWithDefaults() *ExperimentsPublicProtocolResponseDataAttributes {
	this := ExperimentsPublicProtocolResponseDataAttributes{}
	return &this
}

// GetAnalysisPlan returns the AnalysisPlan field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetAnalysisPlan() ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan {
	if o == nil || o.AnalysisPlan == nil {
		var ret ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan
		return ret
	}
	return *o.AnalysisPlan
}

// GetAnalysisPlanOk returns a tuple with the AnalysisPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetAnalysisPlanOk() (*ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan, bool) {
	if o == nil || o.AnalysisPlan == nil {
		return nil, false
	}
	return o.AnalysisPlan, true
}

// HasAnalysisPlan returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) HasAnalysisPlan() bool {
	return o != nil && o.AnalysisPlan != nil
}

// SetAnalysisPlan gets a reference to the given ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan and assigns it to the AnalysisPlan field.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetAnalysisPlan(v ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan) {
	o.AnalysisPlan = &v
}

// GetAssignmentSourceDefaultProperties returns the AssignmentSourceDefaultProperties field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetAssignmentSourceDefaultProperties() []ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems {
	if o == nil || o.AssignmentSourceDefaultProperties == nil {
		var ret []ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems
		return ret
	}
	return o.AssignmentSourceDefaultProperties
}

// GetAssignmentSourceDefaultPropertiesOk returns a tuple with the AssignmentSourceDefaultProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetAssignmentSourceDefaultPropertiesOk() (*[]ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems, bool) {
	if o == nil || o.AssignmentSourceDefaultProperties == nil {
		return nil, false
	}
	return &o.AssignmentSourceDefaultProperties, true
}

// HasAssignmentSourceDefaultProperties returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) HasAssignmentSourceDefaultProperties() bool {
	return o != nil && o.AssignmentSourceDefaultProperties != nil
}

// SetAssignmentSourceDefaultProperties gets a reference to the given []ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems and assigns it to the AssignmentSourceDefaultProperties field.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetAssignmentSourceDefaultProperties(v []ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems) {
	o.AssignmentSourceDefaultProperties = v
}

// GetAssignmentSourceId returns the AssignmentSourceId field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetAssignmentSourceId() string {
	if o == nil || o.AssignmentSourceId == nil {
		var ret string
		return ret
	}
	return *o.AssignmentSourceId
}

// GetAssignmentSourceIdOk returns a tuple with the AssignmentSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetAssignmentSourceIdOk() (*string, bool) {
	if o == nil || o.AssignmentSourceId == nil {
		return nil, false
	}
	return o.AssignmentSourceId, true
}

// HasAssignmentSourceId returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) HasAssignmentSourceId() bool {
	return o != nil && o.AssignmentSourceId != nil
}

// SetAssignmentSourceId gets a reference to the given string and assigns it to the AssignmentSourceId field.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetAssignmentSourceId(v string) {
	o.AssignmentSourceId = &v
}

// GetDefaultDurationDays returns the DefaultDurationDays field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetDefaultDurationDays() int64 {
	if o == nil || o.DefaultDurationDays == nil {
		var ret int64
		return ret
	}
	return *o.DefaultDurationDays
}

// GetDefaultDurationDaysOk returns a tuple with the DefaultDurationDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetDefaultDurationDaysOk() (*int64, bool) {
	if o == nil || o.DefaultDurationDays == nil {
		return nil, false
	}
	return o.DefaultDurationDays, true
}

// HasDefaultDurationDays returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) HasDefaultDurationDays() bool {
	return o != nil && o.DefaultDurationDays != nil
}

// SetDefaultDurationDays gets a reference to the given int64 and assigns it to the DefaultDurationDays field.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetDefaultDurationDays(v int64) {
	o.DefaultDurationDays = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEnforcement returns the Enforcement field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetEnforcement() ExperimentsPublicProtocolResponseDataAttributesEnforcement {
	if o == nil || o.Enforcement == nil {
		var ret ExperimentsPublicProtocolResponseDataAttributesEnforcement
		return ret
	}
	return *o.Enforcement
}

// GetEnforcementOk returns a tuple with the Enforcement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetEnforcementOk() (*ExperimentsPublicProtocolResponseDataAttributesEnforcement, bool) {
	if o == nil || o.Enforcement == nil {
		return nil, false
	}
	return o.Enforcement, true
}

// HasEnforcement returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) HasEnforcement() bool {
	return o != nil && o.Enforcement != nil
}

// SetEnforcement gets a reference to the given ExperimentsPublicProtocolResponseDataAttributesEnforcement and assigns it to the Enforcement field.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetEnforcement(v ExperimentsPublicProtocolResponseDataAttributesEnforcement) {
	o.Enforcement = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId != nil
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetExposureSchedule returns the ExposureSchedule field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetExposureSchedule() ExperimentsPublicProtocolResponseDataAttributesExposureSchedule {
	if o == nil || o.ExposureSchedule == nil {
		var ret ExperimentsPublicProtocolResponseDataAttributesExposureSchedule
		return ret
	}
	return *o.ExposureSchedule
}

// GetExposureScheduleOk returns a tuple with the ExposureSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetExposureScheduleOk() (*ExperimentsPublicProtocolResponseDataAttributesExposureSchedule, bool) {
	if o == nil || o.ExposureSchedule == nil {
		return nil, false
	}
	return o.ExposureSchedule, true
}

// HasExposureSchedule returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) HasExposureSchedule() bool {
	return o != nil && o.ExposureSchedule != nil
}

// SetExposureSchedule gets a reference to the given ExperimentsPublicProtocolResponseDataAttributesExposureSchedule and assigns it to the ExposureSchedule field.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetExposureSchedule(v ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) {
	o.ExposureSchedule = &v
}

// GetIsDurationRequiredToStart returns the IsDurationRequiredToStart field value.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetIsDurationRequiredToStart() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsDurationRequiredToStart
}

// GetIsDurationRequiredToStartOk returns a tuple with the IsDurationRequiredToStart field value
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetIsDurationRequiredToStartOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDurationRequiredToStart, true
}

// SetIsDurationRequiredToStart sets field value.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetIsDurationRequiredToStart(v bool) {
	o.IsDurationRequiredToStart = v
}

// GetIsEqualSplitEnforced returns the IsEqualSplitEnforced field value.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetIsEqualSplitEnforced() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsEqualSplitEnforced
}

// GetIsEqualSplitEnforcedOk returns a tuple with the IsEqualSplitEnforced field value
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetIsEqualSplitEnforcedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEqualSplitEnforced, true
}

// SetIsEqualSplitEnforced sets field value.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetIsEqualSplitEnforced(v bool) {
	o.IsEqualSplitEnforced = v
}

// GetIsMetricLimitEnabled returns the IsMetricLimitEnabled field value.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetIsMetricLimitEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsMetricLimitEnabled
}

// GetIsMetricLimitEnabledOk returns a tuple with the IsMetricLimitEnabled field value
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetIsMetricLimitEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMetricLimitEnabled, true
}

// SetIsMetricLimitEnabled sets field value.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetIsMetricLimitEnabled(v bool) {
	o.IsMetricLimitEnabled = v
}

// GetIsMinimumDurationEnabled returns the IsMinimumDurationEnabled field value.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetIsMinimumDurationEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsMinimumDurationEnabled
}

// GetIsMinimumDurationEnabledOk returns a tuple with the IsMinimumDurationEnabled field value
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetIsMinimumDurationEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMinimumDurationEnabled, true
}

// SetIsMinimumDurationEnabled sets field value.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetIsMinimumDurationEnabled(v bool) {
	o.IsMinimumDurationEnabled = v
}

// GetMetricGroups returns the MetricGroups field value.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetMetricGroups() []ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems {
	if o == nil {
		var ret []ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems
		return ret
	}
	return o.MetricGroups
}

// GetMetricGroupsOk returns a tuple with the MetricGroups field value
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetMetricGroupsOk() (*[]ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricGroups, true
}

// SetMetricGroups sets field value.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetMetricGroups(v []ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems) {
	o.MetricGroups = v
}

// GetMetricLimit returns the MetricLimit field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetMetricLimit() int64 {
	if o == nil || o.MetricLimit == nil {
		var ret int64
		return ret
	}
	return *o.MetricLimit
}

// GetMetricLimitOk returns a tuple with the MetricLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetMetricLimitOk() (*int64, bool) {
	if o == nil || o.MetricLimit == nil {
		return nil, false
	}
	return o.MetricLimit, true
}

// HasMetricLimit returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) HasMetricLimit() bool {
	return o != nil && o.MetricLimit != nil
}

// SetMetricLimit gets a reference to the given int64 and assigns it to the MetricLimit field.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetMetricLimit(v int64) {
	o.MetricLimit = &v
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetMinimumDurationUnit returns the MinimumDurationUnit field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetMinimumDurationUnit() string {
	if o == nil || o.MinimumDurationUnit == nil {
		var ret string
		return ret
	}
	return *o.MinimumDurationUnit
}

// GetMinimumDurationUnitOk returns a tuple with the MinimumDurationUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetMinimumDurationUnitOk() (*string, bool) {
	if o == nil || o.MinimumDurationUnit == nil {
		return nil, false
	}
	return o.MinimumDurationUnit, true
}

// HasMinimumDurationUnit returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) HasMinimumDurationUnit() bool {
	return o != nil && o.MinimumDurationUnit != nil
}

// SetMinimumDurationUnit gets a reference to the given string and assigns it to the MinimumDurationUnit field.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetMinimumDurationUnit(v string) {
	o.MinimumDurationUnit = &v
}

// GetMinimumDurationValue returns the MinimumDurationValue field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetMinimumDurationValue() int64 {
	if o == nil || o.MinimumDurationValue == nil {
		var ret int64
		return ret
	}
	return *o.MinimumDurationValue
}

// GetMinimumDurationValueOk returns a tuple with the MinimumDurationValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetMinimumDurationValueOk() (*int64, bool) {
	if o == nil || o.MinimumDurationValue == nil {
		return nil, false
	}
	return o.MinimumDurationValue, true
}

// HasMinimumDurationValue returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) HasMinimumDurationValue() bool {
	return o != nil && o.MinimumDurationValue != nil
}

// SetMinimumDurationValue gets a reference to the given int64 and assigns it to the MinimumDurationValue field.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetMinimumDurationValue(v int64) {
	o.MinimumDurationValue = &v
}

// GetName returns the Name field value.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetName(v string) {
	o.Name = v
}

// GetPrimaryMetric returns the PrimaryMetric field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetPrimaryMetric() ExperimentsPublicProtocolResponseDataAttributesSubjectType {
	if o == nil || o.PrimaryMetric == nil {
		var ret ExperimentsPublicProtocolResponseDataAttributesSubjectType
		return ret
	}
	return *o.PrimaryMetric
}

// GetPrimaryMetricOk returns a tuple with the PrimaryMetric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetPrimaryMetricOk() (*ExperimentsPublicProtocolResponseDataAttributesSubjectType, bool) {
	if o == nil || o.PrimaryMetric == nil {
		return nil, false
	}
	return o.PrimaryMetric, true
}

// HasPrimaryMetric returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) HasPrimaryMetric() bool {
	return o != nil && o.PrimaryMetric != nil
}

// SetPrimaryMetric gets a reference to the given ExperimentsPublicProtocolResponseDataAttributesSubjectType and assigns it to the PrimaryMetric field.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetPrimaryMetric(v ExperimentsPublicProtocolResponseDataAttributesSubjectType) {
	o.PrimaryMetric = &v
}

// GetPrimaryMetricId returns the PrimaryMetricId field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetPrimaryMetricId() string {
	if o == nil || o.PrimaryMetricId == nil {
		var ret string
		return ret
	}
	return *o.PrimaryMetricId
}

// GetPrimaryMetricIdOk returns a tuple with the PrimaryMetricId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetPrimaryMetricIdOk() (*string, bool) {
	if o == nil || o.PrimaryMetricId == nil {
		return nil, false
	}
	return o.PrimaryMetricId, true
}

// HasPrimaryMetricId returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) HasPrimaryMetricId() bool {
	return o != nil && o.PrimaryMetricId != nil
}

// SetPrimaryMetricId gets a reference to the given string and assigns it to the PrimaryMetricId field.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetPrimaryMetricId(v string) {
	o.PrimaryMetricId = &v
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetPublishedAt() time.Time {
	if o == nil || o.PublishedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetPublishedAtOk() (*time.Time, bool) {
	if o == nil || o.PublishedAt == nil {
		return nil, false
	}
	return o.PublishedAt, true
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) HasPublishedAt() bool {
	return o != nil && o.PublishedAt != nil
}

// SetPublishedAt gets a reference to the given time.Time and assigns it to the PublishedAt field.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetPublishedAt(v time.Time) {
	o.PublishedAt = &v
}

// GetStatus returns the Status field value.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetStatus() ExperimentsPublicProtocolResponseDataAttributesStatus {
	if o == nil {
		var ret ExperimentsPublicProtocolResponseDataAttributesStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetStatusOk() (*ExperimentsPublicProtocolResponseDataAttributesStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetStatus(v ExperimentsPublicProtocolResponseDataAttributesStatus) {
	o.Status = v
}

// GetSubjectType returns the SubjectType field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetSubjectType() ExperimentsPublicProtocolResponseDataAttributesSubjectType {
	if o == nil || o.SubjectType == nil {
		var ret ExperimentsPublicProtocolResponseDataAttributesSubjectType
		return ret
	}
	return *o.SubjectType
}

// GetSubjectTypeOk returns a tuple with the SubjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetSubjectTypeOk() (*ExperimentsPublicProtocolResponseDataAttributesSubjectType, bool) {
	if o == nil || o.SubjectType == nil {
		return nil, false
	}
	return o.SubjectType, true
}

// HasSubjectType returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) HasSubjectType() bool {
	return o != nil && o.SubjectType != nil
}

// SetSubjectType gets a reference to the given ExperimentsPublicProtocolResponseDataAttributesSubjectType and assigns it to the SubjectType field.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetSubjectType(v ExperimentsPublicProtocolResponseDataAttributesSubjectType) {
	o.SubjectType = &v
}

// GetSubjectTypeId returns the SubjectTypeId field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetSubjectTypeId() string {
	if o == nil || o.SubjectTypeId == nil {
		var ret string
		return ret
	}
	return *o.SubjectTypeId
}

// GetSubjectTypeIdOk returns a tuple with the SubjectTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetSubjectTypeIdOk() (*string, bool) {
	if o == nil || o.SubjectTypeId == nil {
		return nil, false
	}
	return o.SubjectTypeId, true
}

// HasSubjectTypeId returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) HasSubjectTypeId() bool {
	return o != nil && o.SubjectTypeId != nil
}

// SetSubjectTypeId gets a reference to the given string and assigns it to the SubjectTypeId field.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetSubjectTypeId(v string) {
	o.SubjectTypeId = &v
}

// GetTargetingRules returns the TargetingRules field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetTargetingRules() []ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems {
	if o == nil || o.TargetingRules == nil {
		var ret []ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems
		return ret
	}
	return o.TargetingRules
}

// GetTargetingRulesOk returns a tuple with the TargetingRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetTargetingRulesOk() (*[]ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems, bool) {
	if o == nil || o.TargetingRules == nil {
		return nil, false
	}
	return &o.TargetingRules, true
}

// HasTargetingRules returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) HasTargetingRules() bool {
	return o != nil && o.TargetingRules != nil
}

// SetTargetingRules gets a reference to the given []ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems and assigns it to the TargetingRules field.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetTargetingRules(v []ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems) {
	o.TargetingRules = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *ExperimentsPublicProtocolResponseDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPublicProtocolResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AnalysisPlan != nil {
		toSerialize["analysis_plan"] = o.AnalysisPlan
	}
	if o.AssignmentSourceDefaultProperties != nil {
		toSerialize["assignment_source_default_properties"] = o.AssignmentSourceDefaultProperties
	}
	if o.AssignmentSourceId != nil {
		toSerialize["assignment_source_id"] = o.AssignmentSourceId
	}
	if o.DefaultDurationDays != nil {
		toSerialize["default_duration_days"] = o.DefaultDurationDays
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enforcement != nil {
		toSerialize["enforcement"] = o.Enforcement
	}
	if o.EnvironmentId != nil {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	if o.ExposureSchedule != nil {
		toSerialize["exposure_schedule"] = o.ExposureSchedule
	}
	toSerialize["is_duration_required_to_start"] = o.IsDurationRequiredToStart
	toSerialize["is_equal_split_enforced"] = o.IsEqualSplitEnforced
	toSerialize["is_metric_limit_enabled"] = o.IsMetricLimitEnabled
	toSerialize["is_minimum_duration_enabled"] = o.IsMinimumDurationEnabled
	toSerialize["metric_groups"] = o.MetricGroups
	if o.MetricLimit != nil {
		toSerialize["metric_limit"] = o.MetricLimit
	}
	if o.MigrationMetadata != nil {
		toSerialize["migration_metadata"] = o.MigrationMetadata
	}
	if o.MinimumDurationUnit != nil {
		toSerialize["minimum_duration_unit"] = o.MinimumDurationUnit
	}
	if o.MinimumDurationValue != nil {
		toSerialize["minimum_duration_value"] = o.MinimumDurationValue
	}
	toSerialize["name"] = o.Name
	if o.PrimaryMetric != nil {
		toSerialize["primary_metric"] = o.PrimaryMetric
	}
	if o.PrimaryMetricId != nil {
		toSerialize["primary_metric_id"] = o.PrimaryMetricId
	}
	if o.PublishedAt != nil {
		if o.PublishedAt.Nanosecond() == 0 {
			toSerialize["published_at"] = o.PublishedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["published_at"] = o.PublishedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["status"] = o.Status
	if o.SubjectType != nil {
		toSerialize["subject_type"] = o.SubjectType
	}
	if o.SubjectTypeId != nil {
		toSerialize["subject_type_id"] = o.SubjectTypeId
	}
	if o.TargetingRules != nil {
		toSerialize["targeting_rules"] = o.TargetingRules
	}
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsPublicProtocolResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnalysisPlan                      *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan                            `json:"analysis_plan,omitempty"`
		AssignmentSourceDefaultProperties []ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems `json:"assignment_source_default_properties,omitempty"`
		AssignmentSourceId                *string                                                                                 `json:"assignment_source_id,omitempty"`
		DefaultDurationDays               *int64                                                                                  `json:"default_duration_days,omitempty"`
		Description                       *string                                                                                 `json:"description,omitempty"`
		Enforcement                       *ExperimentsPublicProtocolResponseDataAttributesEnforcement                             `json:"enforcement,omitempty"`
		EnvironmentId                     *string                                                                                 `json:"environment_id,omitempty"`
		ExposureSchedule                  *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule                        `json:"exposure_schedule,omitempty"`
		IsDurationRequiredToStart         *bool                                                                                   `json:"is_duration_required_to_start"`
		IsEqualSplitEnforced              *bool                                                                                   `json:"is_equal_split_enforced"`
		IsMetricLimitEnabled              *bool                                                                                   `json:"is_metric_limit_enabled"`
		IsMinimumDurationEnabled          *bool                                                                                   `json:"is_minimum_duration_enabled"`
		MetricGroups                      *[]ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems                     `json:"metric_groups"`
		MetricLimit                       *int64                                                                                  `json:"metric_limit,omitempty"`
		MigrationMetadata                 interface{}                                                                             `json:"migration_metadata,omitempty"`
		MinimumDurationUnit               *string                                                                                 `json:"minimum_duration_unit,omitempty"`
		MinimumDurationValue              *int64                                                                                  `json:"minimum_duration_value,omitempty"`
		Name                              *string                                                                                 `json:"name"`
		PrimaryMetric                     *ExperimentsPublicProtocolResponseDataAttributesSubjectType                             `json:"primary_metric,omitempty"`
		PrimaryMetricId                   *string                                                                                 `json:"primary_metric_id,omitempty"`
		PublishedAt                       *time.Time                                                                              `json:"published_at,omitempty"`
		Status                            *ExperimentsPublicProtocolResponseDataAttributesStatus                                  `json:"status"`
		SubjectType                       *ExperimentsPublicProtocolResponseDataAttributesSubjectType                             `json:"subject_type,omitempty"`
		SubjectTypeId                     *string                                                                                 `json:"subject_type_id,omitempty"`
		TargetingRules                    []ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems                    `json:"targeting_rules,omitempty"`
		UpdatedAt                         *string                                                                                 `json:"updated_at"`
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
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"analysis_plan", "assignment_source_default_properties", "assignment_source_id", "default_duration_days", "description", "enforcement", "environment_id", "exposure_schedule", "is_duration_required_to_start", "is_equal_split_enforced", "is_metric_limit_enabled", "is_minimum_duration_enabled", "metric_groups", "metric_limit", "migration_metadata", "minimum_duration_unit", "minimum_duration_value", "name", "primary_metric", "primary_metric_id", "published_at", "status", "subject_type", "subject_type_id", "targeting_rules", "updated_at"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AnalysisPlan != nil && all.AnalysisPlan.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AnalysisPlan = all.AnalysisPlan
	o.AssignmentSourceDefaultProperties = all.AssignmentSourceDefaultProperties
	o.AssignmentSourceId = all.AssignmentSourceId
	o.DefaultDurationDays = all.DefaultDurationDays
	o.Description = all.Description
	if all.Enforcement != nil && all.Enforcement.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Enforcement = all.Enforcement
	o.EnvironmentId = all.EnvironmentId
	if all.ExposureSchedule != nil && all.ExposureSchedule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
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
	o.Name = *all.Name
	if all.PrimaryMetric != nil && all.PrimaryMetric.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PrimaryMetric = all.PrimaryMetric
	o.PrimaryMetricId = all.PrimaryMetricId
	o.PublishedAt = all.PublishedAt
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	if all.SubjectType != nil && all.SubjectType.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SubjectType = all.SubjectType
	o.SubjectTypeId = all.SubjectTypeId
	o.TargetingRules = all.TargetingRules
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
