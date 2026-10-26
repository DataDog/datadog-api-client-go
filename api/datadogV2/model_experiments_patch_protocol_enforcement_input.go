// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPatchProtocolEnforcementInput Partial lock map. Omitted fields remain unchanged. Null clears the whole map. Empty or null field values become `EDITABLE`, except `subject_type`, `flag_source`, and `confidence_interval_method`, which remain `LOCKED`.
type ExperimentsPatchProtocolEnforcementInput struct {
	// Always `LOCKED`. The server keeps this setting locked regardless of the supplied value.
	ConfidenceIntervalMethod datadog.NullableString `json:"confidence_interval_method,omitempty"`
	// LOCKED prevents changes to confidence level. EDITABLE permits changes.
	ConfidenceLevel datadog.NullableString `json:"confidence_level,omitempty"`
	// LOCKED prevents changes to CUPED variance reduction. EDITABLE permits changes.
	CupedCalculation datadog.NullableString `json:"cuped_calculation,omitempty"`
	// LOCKED prevents changes to default duration. EDITABLE permits changes.
	DefaultDuration datadog.NullableString `json:"default_duration,omitempty"`
	// LOCKED prevents changes to environment. EDITABLE permits changes.
	Environment datadog.NullableString `json:"environment,omitempty"`
	// Always LOCKED. The server keeps this setting locked regardless of the supplied value.
	FlagSource datadog.NullableString `json:"flag_source,omitempty"`
	// LOCKED prevents changes to multiple testing correction. EDITABLE permits changes.
	MultipleTestingCorrection datadog.NullableString `json:"multiple_testing_correction,omitempty"`
	// LOCKED prevents changes to notifications. EDITABLE permits changes.
	Notifications datadog.NullableString `json:"notifications,omitempty"`
	// LOCKED prevents changes to primary metric. EDITABLE permits changes.
	PrimaryMetric datadog.NullableString `json:"primary_metric,omitempty"`
	// LOCKED prevents changes to secondary metrics. EDITABLE permits changes.
	SecondaryMetrics datadog.NullableString `json:"secondary_metrics,omitempty"`
	// LOCKED prevents changes to result exploration dimensions. EDITABLE permits changes.
	SplitByExplorationDimensions datadog.NullableString `json:"split_by_exploration_dimensions,omitempty"`
	// Always LOCKED. The server keeps this setting locked regardless of the supplied value.
	SubjectType datadog.NullableString `json:"subject_type,omitempty"`
	// LOCKED prevents changes to targeting rules. EDITABLE permits changes.
	TargetingRules datadog.NullableString `json:"targeting_rules,omitempty"`
	// LOCKED prevents changes to traffic exposure. EDITABLE permits changes.
	TrafficExposure datadog.NullableString `json:"traffic_exposure,omitempty"`
	// LOCKED prevents changes to warehouse exposure source. EDITABLE permits changes.
	WarehouseExposureSource datadog.NullableString `json:"warehouse_exposure_source,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPatchProtocolEnforcementInput instantiates a new ExperimentsPatchProtocolEnforcementInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPatchProtocolEnforcementInput() *ExperimentsPatchProtocolEnforcementInput {
	this := ExperimentsPatchProtocolEnforcementInput{}
	return &this
}

// NewExperimentsPatchProtocolEnforcementInputWithDefaults instantiates a new ExperimentsPatchProtocolEnforcementInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPatchProtocolEnforcementInputWithDefaults() *ExperimentsPatchProtocolEnforcementInput {
	this := ExperimentsPatchProtocolEnforcementInput{}
	return &this
}

// GetConfidenceIntervalMethod returns the ConfidenceIntervalMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolEnforcementInput) GetConfidenceIntervalMethod() string {
	if o == nil || o.ConfidenceIntervalMethod.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfidenceIntervalMethod.Get()
}

// GetConfidenceIntervalMethodOk returns a tuple with the ConfidenceIntervalMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolEnforcementInput) GetConfidenceIntervalMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfidenceIntervalMethod.Get(), o.ConfidenceIntervalMethod.IsSet()
}

// HasConfidenceIntervalMethod returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolEnforcementInput) HasConfidenceIntervalMethod() bool {
	return o != nil && o.ConfidenceIntervalMethod.IsSet()
}

// SetConfidenceIntervalMethod gets a reference to the given datadog.NullableString and assigns it to the ConfidenceIntervalMethod field.
func (o *ExperimentsPatchProtocolEnforcementInput) SetConfidenceIntervalMethod(v string) {
	o.ConfidenceIntervalMethod.Set(&v)
}

// SetConfidenceIntervalMethodNil sets the value for ConfidenceIntervalMethod to be an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) SetConfidenceIntervalMethodNil() {
	o.ConfidenceIntervalMethod.Set(nil)
}

// UnsetConfidenceIntervalMethod ensures that no value is present for ConfidenceIntervalMethod, not even an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) UnsetConfidenceIntervalMethod() {
	o.ConfidenceIntervalMethod.Unset()
}

// GetConfidenceLevel returns the ConfidenceLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolEnforcementInput) GetConfidenceLevel() string {
	if o == nil || o.ConfidenceLevel.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfidenceLevel.Get()
}

// GetConfidenceLevelOk returns a tuple with the ConfidenceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolEnforcementInput) GetConfidenceLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfidenceLevel.Get(), o.ConfidenceLevel.IsSet()
}

// HasConfidenceLevel returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolEnforcementInput) HasConfidenceLevel() bool {
	return o != nil && o.ConfidenceLevel.IsSet()
}

// SetConfidenceLevel gets a reference to the given datadog.NullableString and assigns it to the ConfidenceLevel field.
func (o *ExperimentsPatchProtocolEnforcementInput) SetConfidenceLevel(v string) {
	o.ConfidenceLevel.Set(&v)
}

// SetConfidenceLevelNil sets the value for ConfidenceLevel to be an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) SetConfidenceLevelNil() {
	o.ConfidenceLevel.Set(nil)
}

// UnsetConfidenceLevel ensures that no value is present for ConfidenceLevel, not even an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) UnsetConfidenceLevel() {
	o.ConfidenceLevel.Unset()
}

// GetCupedCalculation returns the CupedCalculation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolEnforcementInput) GetCupedCalculation() string {
	if o == nil || o.CupedCalculation.Get() == nil {
		var ret string
		return ret
	}
	return *o.CupedCalculation.Get()
}

// GetCupedCalculationOk returns a tuple with the CupedCalculation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolEnforcementInput) GetCupedCalculationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CupedCalculation.Get(), o.CupedCalculation.IsSet()
}

// HasCupedCalculation returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolEnforcementInput) HasCupedCalculation() bool {
	return o != nil && o.CupedCalculation.IsSet()
}

// SetCupedCalculation gets a reference to the given datadog.NullableString and assigns it to the CupedCalculation field.
func (o *ExperimentsPatchProtocolEnforcementInput) SetCupedCalculation(v string) {
	o.CupedCalculation.Set(&v)
}

// SetCupedCalculationNil sets the value for CupedCalculation to be an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) SetCupedCalculationNil() {
	o.CupedCalculation.Set(nil)
}

// UnsetCupedCalculation ensures that no value is present for CupedCalculation, not even an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) UnsetCupedCalculation() {
	o.CupedCalculation.Unset()
}

// GetDefaultDuration returns the DefaultDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolEnforcementInput) GetDefaultDuration() string {
	if o == nil || o.DefaultDuration.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultDuration.Get()
}

// GetDefaultDurationOk returns a tuple with the DefaultDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolEnforcementInput) GetDefaultDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultDuration.Get(), o.DefaultDuration.IsSet()
}

// HasDefaultDuration returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolEnforcementInput) HasDefaultDuration() bool {
	return o != nil && o.DefaultDuration.IsSet()
}

// SetDefaultDuration gets a reference to the given datadog.NullableString and assigns it to the DefaultDuration field.
func (o *ExperimentsPatchProtocolEnforcementInput) SetDefaultDuration(v string) {
	o.DefaultDuration.Set(&v)
}

// SetDefaultDurationNil sets the value for DefaultDuration to be an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) SetDefaultDurationNil() {
	o.DefaultDuration.Set(nil)
}

// UnsetDefaultDuration ensures that no value is present for DefaultDuration, not even an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) UnsetDefaultDuration() {
	o.DefaultDuration.Unset()
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolEnforcementInput) GetEnvironment() string {
	if o == nil || o.Environment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolEnforcementInput) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolEnforcementInput) HasEnvironment() bool {
	return o != nil && o.Environment.IsSet()
}

// SetEnvironment gets a reference to the given datadog.NullableString and assigns it to the Environment field.
func (o *ExperimentsPatchProtocolEnforcementInput) SetEnvironment(v string) {
	o.Environment.Set(&v)
}

// SetEnvironmentNil sets the value for Environment to be an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) UnsetEnvironment() {
	o.Environment.Unset()
}

// GetFlagSource returns the FlagSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolEnforcementInput) GetFlagSource() string {
	if o == nil || o.FlagSource.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlagSource.Get()
}

// GetFlagSourceOk returns a tuple with the FlagSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolEnforcementInput) GetFlagSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlagSource.Get(), o.FlagSource.IsSet()
}

// HasFlagSource returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolEnforcementInput) HasFlagSource() bool {
	return o != nil && o.FlagSource.IsSet()
}

// SetFlagSource gets a reference to the given datadog.NullableString and assigns it to the FlagSource field.
func (o *ExperimentsPatchProtocolEnforcementInput) SetFlagSource(v string) {
	o.FlagSource.Set(&v)
}

// SetFlagSourceNil sets the value for FlagSource to be an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) SetFlagSourceNil() {
	o.FlagSource.Set(nil)
}

// UnsetFlagSource ensures that no value is present for FlagSource, not even an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) UnsetFlagSource() {
	o.FlagSource.Unset()
}

// GetMultipleTestingCorrection returns the MultipleTestingCorrection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolEnforcementInput) GetMultipleTestingCorrection() string {
	if o == nil || o.MultipleTestingCorrection.Get() == nil {
		var ret string
		return ret
	}
	return *o.MultipleTestingCorrection.Get()
}

// GetMultipleTestingCorrectionOk returns a tuple with the MultipleTestingCorrection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolEnforcementInput) GetMultipleTestingCorrectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MultipleTestingCorrection.Get(), o.MultipleTestingCorrection.IsSet()
}

// HasMultipleTestingCorrection returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolEnforcementInput) HasMultipleTestingCorrection() bool {
	return o != nil && o.MultipleTestingCorrection.IsSet()
}

// SetMultipleTestingCorrection gets a reference to the given datadog.NullableString and assigns it to the MultipleTestingCorrection field.
func (o *ExperimentsPatchProtocolEnforcementInput) SetMultipleTestingCorrection(v string) {
	o.MultipleTestingCorrection.Set(&v)
}

// SetMultipleTestingCorrectionNil sets the value for MultipleTestingCorrection to be an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) SetMultipleTestingCorrectionNil() {
	o.MultipleTestingCorrection.Set(nil)
}

// UnsetMultipleTestingCorrection ensures that no value is present for MultipleTestingCorrection, not even an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) UnsetMultipleTestingCorrection() {
	o.MultipleTestingCorrection.Unset()
}

// GetNotifications returns the Notifications field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolEnforcementInput) GetNotifications() string {
	if o == nil || o.Notifications.Get() == nil {
		var ret string
		return ret
	}
	return *o.Notifications.Get()
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolEnforcementInput) GetNotificationsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notifications.Get(), o.Notifications.IsSet()
}

// HasNotifications returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolEnforcementInput) HasNotifications() bool {
	return o != nil && o.Notifications.IsSet()
}

// SetNotifications gets a reference to the given datadog.NullableString and assigns it to the Notifications field.
func (o *ExperimentsPatchProtocolEnforcementInput) SetNotifications(v string) {
	o.Notifications.Set(&v)
}

// SetNotificationsNil sets the value for Notifications to be an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) SetNotificationsNil() {
	o.Notifications.Set(nil)
}

// UnsetNotifications ensures that no value is present for Notifications, not even an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) UnsetNotifications() {
	o.Notifications.Unset()
}

// GetPrimaryMetric returns the PrimaryMetric field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolEnforcementInput) GetPrimaryMetric() string {
	if o == nil || o.PrimaryMetric.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrimaryMetric.Get()
}

// GetPrimaryMetricOk returns a tuple with the PrimaryMetric field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolEnforcementInput) GetPrimaryMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryMetric.Get(), o.PrimaryMetric.IsSet()
}

// HasPrimaryMetric returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolEnforcementInput) HasPrimaryMetric() bool {
	return o != nil && o.PrimaryMetric.IsSet()
}

// SetPrimaryMetric gets a reference to the given datadog.NullableString and assigns it to the PrimaryMetric field.
func (o *ExperimentsPatchProtocolEnforcementInput) SetPrimaryMetric(v string) {
	o.PrimaryMetric.Set(&v)
}

// SetPrimaryMetricNil sets the value for PrimaryMetric to be an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) SetPrimaryMetricNil() {
	o.PrimaryMetric.Set(nil)
}

// UnsetPrimaryMetric ensures that no value is present for PrimaryMetric, not even an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) UnsetPrimaryMetric() {
	o.PrimaryMetric.Unset()
}

// GetSecondaryMetrics returns the SecondaryMetrics field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolEnforcementInput) GetSecondaryMetrics() string {
	if o == nil || o.SecondaryMetrics.Get() == nil {
		var ret string
		return ret
	}
	return *o.SecondaryMetrics.Get()
}

// GetSecondaryMetricsOk returns a tuple with the SecondaryMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolEnforcementInput) GetSecondaryMetricsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecondaryMetrics.Get(), o.SecondaryMetrics.IsSet()
}

// HasSecondaryMetrics returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolEnforcementInput) HasSecondaryMetrics() bool {
	return o != nil && o.SecondaryMetrics.IsSet()
}

// SetSecondaryMetrics gets a reference to the given datadog.NullableString and assigns it to the SecondaryMetrics field.
func (o *ExperimentsPatchProtocolEnforcementInput) SetSecondaryMetrics(v string) {
	o.SecondaryMetrics.Set(&v)
}

// SetSecondaryMetricsNil sets the value for SecondaryMetrics to be an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) SetSecondaryMetricsNil() {
	o.SecondaryMetrics.Set(nil)
}

// UnsetSecondaryMetrics ensures that no value is present for SecondaryMetrics, not even an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) UnsetSecondaryMetrics() {
	o.SecondaryMetrics.Unset()
}

// GetSplitByExplorationDimensions returns the SplitByExplorationDimensions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolEnforcementInput) GetSplitByExplorationDimensions() string {
	if o == nil || o.SplitByExplorationDimensions.Get() == nil {
		var ret string
		return ret
	}
	return *o.SplitByExplorationDimensions.Get()
}

// GetSplitByExplorationDimensionsOk returns a tuple with the SplitByExplorationDimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolEnforcementInput) GetSplitByExplorationDimensionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SplitByExplorationDimensions.Get(), o.SplitByExplorationDimensions.IsSet()
}

// HasSplitByExplorationDimensions returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolEnforcementInput) HasSplitByExplorationDimensions() bool {
	return o != nil && o.SplitByExplorationDimensions.IsSet()
}

// SetSplitByExplorationDimensions gets a reference to the given datadog.NullableString and assigns it to the SplitByExplorationDimensions field.
func (o *ExperimentsPatchProtocolEnforcementInput) SetSplitByExplorationDimensions(v string) {
	o.SplitByExplorationDimensions.Set(&v)
}

// SetSplitByExplorationDimensionsNil sets the value for SplitByExplorationDimensions to be an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) SetSplitByExplorationDimensionsNil() {
	o.SplitByExplorationDimensions.Set(nil)
}

// UnsetSplitByExplorationDimensions ensures that no value is present for SplitByExplorationDimensions, not even an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) UnsetSplitByExplorationDimensions() {
	o.SplitByExplorationDimensions.Unset()
}

// GetSubjectType returns the SubjectType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolEnforcementInput) GetSubjectType() string {
	if o == nil || o.SubjectType.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubjectType.Get()
}

// GetSubjectTypeOk returns a tuple with the SubjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolEnforcementInput) GetSubjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubjectType.Get(), o.SubjectType.IsSet()
}

// HasSubjectType returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolEnforcementInput) HasSubjectType() bool {
	return o != nil && o.SubjectType.IsSet()
}

// SetSubjectType gets a reference to the given datadog.NullableString and assigns it to the SubjectType field.
func (o *ExperimentsPatchProtocolEnforcementInput) SetSubjectType(v string) {
	o.SubjectType.Set(&v)
}

// SetSubjectTypeNil sets the value for SubjectType to be an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) SetSubjectTypeNil() {
	o.SubjectType.Set(nil)
}

// UnsetSubjectType ensures that no value is present for SubjectType, not even an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) UnsetSubjectType() {
	o.SubjectType.Unset()
}

// GetTargetingRules returns the TargetingRules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolEnforcementInput) GetTargetingRules() string {
	if o == nil || o.TargetingRules.Get() == nil {
		var ret string
		return ret
	}
	return *o.TargetingRules.Get()
}

// GetTargetingRulesOk returns a tuple with the TargetingRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolEnforcementInput) GetTargetingRulesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetingRules.Get(), o.TargetingRules.IsSet()
}

// HasTargetingRules returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolEnforcementInput) HasTargetingRules() bool {
	return o != nil && o.TargetingRules.IsSet()
}

// SetTargetingRules gets a reference to the given datadog.NullableString and assigns it to the TargetingRules field.
func (o *ExperimentsPatchProtocolEnforcementInput) SetTargetingRules(v string) {
	o.TargetingRules.Set(&v)
}

// SetTargetingRulesNil sets the value for TargetingRules to be an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) SetTargetingRulesNil() {
	o.TargetingRules.Set(nil)
}

// UnsetTargetingRules ensures that no value is present for TargetingRules, not even an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) UnsetTargetingRules() {
	o.TargetingRules.Unset()
}

// GetTrafficExposure returns the TrafficExposure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolEnforcementInput) GetTrafficExposure() string {
	if o == nil || o.TrafficExposure.Get() == nil {
		var ret string
		return ret
	}
	return *o.TrafficExposure.Get()
}

// GetTrafficExposureOk returns a tuple with the TrafficExposure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolEnforcementInput) GetTrafficExposureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrafficExposure.Get(), o.TrafficExposure.IsSet()
}

// HasTrafficExposure returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolEnforcementInput) HasTrafficExposure() bool {
	return o != nil && o.TrafficExposure.IsSet()
}

// SetTrafficExposure gets a reference to the given datadog.NullableString and assigns it to the TrafficExposure field.
func (o *ExperimentsPatchProtocolEnforcementInput) SetTrafficExposure(v string) {
	o.TrafficExposure.Set(&v)
}

// SetTrafficExposureNil sets the value for TrafficExposure to be an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) SetTrafficExposureNil() {
	o.TrafficExposure.Set(nil)
}

// UnsetTrafficExposure ensures that no value is present for TrafficExposure, not even an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) UnsetTrafficExposure() {
	o.TrafficExposure.Unset()
}

// GetWarehouseExposureSource returns the WarehouseExposureSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolEnforcementInput) GetWarehouseExposureSource() string {
	if o == nil || o.WarehouseExposureSource.Get() == nil {
		var ret string
		return ret
	}
	return *o.WarehouseExposureSource.Get()
}

// GetWarehouseExposureSourceOk returns a tuple with the WarehouseExposureSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolEnforcementInput) GetWarehouseExposureSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WarehouseExposureSource.Get(), o.WarehouseExposureSource.IsSet()
}

// HasWarehouseExposureSource returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolEnforcementInput) HasWarehouseExposureSource() bool {
	return o != nil && o.WarehouseExposureSource.IsSet()
}

// SetWarehouseExposureSource gets a reference to the given datadog.NullableString and assigns it to the WarehouseExposureSource field.
func (o *ExperimentsPatchProtocolEnforcementInput) SetWarehouseExposureSource(v string) {
	o.WarehouseExposureSource.Set(&v)
}

// SetWarehouseExposureSourceNil sets the value for WarehouseExposureSource to be an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) SetWarehouseExposureSourceNil() {
	o.WarehouseExposureSource.Set(nil)
}

// UnsetWarehouseExposureSource ensures that no value is present for WarehouseExposureSource, not even an explicit nil.
func (o *ExperimentsPatchProtocolEnforcementInput) UnsetWarehouseExposureSource() {
	o.WarehouseExposureSource.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPatchProtocolEnforcementInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConfidenceIntervalMethod.IsSet() {
		toSerialize["confidence_interval_method"] = o.ConfidenceIntervalMethod.Get()
	}
	if o.ConfidenceLevel.IsSet() {
		toSerialize["confidence_level"] = o.ConfidenceLevel.Get()
	}
	if o.CupedCalculation.IsSet() {
		toSerialize["cuped_calculation"] = o.CupedCalculation.Get()
	}
	if o.DefaultDuration.IsSet() {
		toSerialize["default_duration"] = o.DefaultDuration.Get()
	}
	if o.Environment.IsSet() {
		toSerialize["environment"] = o.Environment.Get()
	}
	if o.FlagSource.IsSet() {
		toSerialize["flag_source"] = o.FlagSource.Get()
	}
	if o.MultipleTestingCorrection.IsSet() {
		toSerialize["multiple_testing_correction"] = o.MultipleTestingCorrection.Get()
	}
	if o.Notifications.IsSet() {
		toSerialize["notifications"] = o.Notifications.Get()
	}
	if o.PrimaryMetric.IsSet() {
		toSerialize["primary_metric"] = o.PrimaryMetric.Get()
	}
	if o.SecondaryMetrics.IsSet() {
		toSerialize["secondary_metrics"] = o.SecondaryMetrics.Get()
	}
	if o.SplitByExplorationDimensions.IsSet() {
		toSerialize["split_by_exploration_dimensions"] = o.SplitByExplorationDimensions.Get()
	}
	if o.SubjectType.IsSet() {
		toSerialize["subject_type"] = o.SubjectType.Get()
	}
	if o.TargetingRules.IsSet() {
		toSerialize["targeting_rules"] = o.TargetingRules.Get()
	}
	if o.TrafficExposure.IsSet() {
		toSerialize["traffic_exposure"] = o.TrafficExposure.Get()
	}
	if o.WarehouseExposureSource.IsSet() {
		toSerialize["warehouse_exposure_source"] = o.WarehouseExposureSource.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsPatchProtocolEnforcementInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfidenceIntervalMethod     datadog.NullableString `json:"confidence_interval_method,omitempty"`
		ConfidenceLevel              datadog.NullableString `json:"confidence_level,omitempty"`
		CupedCalculation             datadog.NullableString `json:"cuped_calculation,omitempty"`
		DefaultDuration              datadog.NullableString `json:"default_duration,omitempty"`
		Environment                  datadog.NullableString `json:"environment,omitempty"`
		FlagSource                   datadog.NullableString `json:"flag_source,omitempty"`
		MultipleTestingCorrection    datadog.NullableString `json:"multiple_testing_correction,omitempty"`
		Notifications                datadog.NullableString `json:"notifications,omitempty"`
		PrimaryMetric                datadog.NullableString `json:"primary_metric,omitempty"`
		SecondaryMetrics             datadog.NullableString `json:"secondary_metrics,omitempty"`
		SplitByExplorationDimensions datadog.NullableString `json:"split_by_exploration_dimensions,omitempty"`
		SubjectType                  datadog.NullableString `json:"subject_type,omitempty"`
		TargetingRules               datadog.NullableString `json:"targeting_rules,omitempty"`
		TrafficExposure              datadog.NullableString `json:"traffic_exposure,omitempty"`
		WarehouseExposureSource      datadog.NullableString `json:"warehouse_exposure_source,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"confidence_interval_method", "confidence_level", "cuped_calculation", "default_duration", "environment", "flag_source", "multiple_testing_correction", "notifications", "primary_metric", "secondary_metrics", "split_by_exploration_dimensions", "subject_type", "targeting_rules", "traffic_exposure", "warehouse_exposure_source"})
	} else {
		return err
	}
	o.ConfidenceIntervalMethod = all.ConfidenceIntervalMethod
	o.ConfidenceLevel = all.ConfidenceLevel
	o.CupedCalculation = all.CupedCalculation
	o.DefaultDuration = all.DefaultDuration
	o.Environment = all.Environment
	o.FlagSource = all.FlagSource
	o.MultipleTestingCorrection = all.MultipleTestingCorrection
	o.Notifications = all.Notifications
	o.PrimaryMetric = all.PrimaryMetric
	o.SecondaryMetrics = all.SecondaryMetrics
	o.SplitByExplorationDimensions = all.SplitByExplorationDimensions
	o.SubjectType = all.SubjectType
	o.TargetingRules = all.TargetingRules
	o.TrafficExposure = all.TrafficExposure
	o.WarehouseExposureSource = all.WarehouseExposureSource

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableExperimentsPatchProtocolEnforcementInput handles when a null is used for ExperimentsPatchProtocolEnforcementInput.
type NullableExperimentsPatchProtocolEnforcementInput struct {
	value *ExperimentsPatchProtocolEnforcementInput
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsPatchProtocolEnforcementInput) Get() *ExperimentsPatchProtocolEnforcementInput {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsPatchProtocolEnforcementInput) Set(val *ExperimentsPatchProtocolEnforcementInput) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsPatchProtocolEnforcementInput) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableExperimentsPatchProtocolEnforcementInput) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsPatchProtocolEnforcementInput initializes the struct as if Set has been called.
func NewNullableExperimentsPatchProtocolEnforcementInput(val *ExperimentsPatchProtocolEnforcementInput) *NullableExperimentsPatchProtocolEnforcementInput {
	return &NullableExperimentsPatchProtocolEnforcementInput{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsPatchProtocolEnforcementInput) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsPatchProtocolEnforcementInput) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
