// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsAnalysisPlanV2MutationResponseDataAttributes Statistical settings and duration targets in the saved analysis plan.
type ExperimentsAnalysisPlanV2MutationResponseDataAttributes struct {
	// Parameters of the prior distribution used for Bayesian analysis.
	BayesianPrior *ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior `json:"bayesian_prior,omitempty"`
	// Statistical method used to calculate the experiment results.
	ConfidenceIntervalMethod *ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod `json:"confidence_interval_method,omitempty"`
	// Confidence level used for statistical analysis, expressed as a fraction.
	ConfidenceLevel *float64 `json:"confidence_level,omitempty"`
	// Number of days of pre-experiment data used for CUPED variance reduction.
	CupedLookbackPeriodDays *int64 `json:"cuped_lookback_period_days,omitempty"`
	// Number of days configured for the experiment to end automatically.
	ExperimentAutoEndDays *int64 `json:"experiment_auto_end_days,omitempty"`
	// Minimum experiment duration in days configured in the analysis plan.
	ExperimentMinDuration *int64 `json:"experiment_min_duration,omitempty"`
	// Minimum sample size configured in the analysis plan.
	ExperimentMinSampleSize *int64 `json:"experiment_min_sample_size,omitempty"`
	// Whether the experiment has custom analysis settings.
	HasCustomAnalysisSettings *bool `json:"has_custom_analysis_settings,omitempty"`
	// Whether CUPED uses pre-experiment data to reduce variance in the analysis.
	IsCupedEnabled *bool `json:"is_cuped_enabled,omitempty"`
	// Whether the analysis adjusts for testing multiple hypotheses.
	IsMultipleTestingCorrectionEnabled *bool `json:"is_multiple_testing_correction_enabled,omitempty"`
	// Weight assigned to the primary metric in the preferential Bonferroni correction.
	PreferentialBonferroniPrimaryMetricWeight *float64 `json:"preferential_bonferroni_primary_metric_weight,omitempty"`
	// Planned experiment duration in days.
	TargetDurationDays datadog.NullableInt64 `json:"target_duration_days,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsAnalysisPlanV2MutationResponseDataAttributes instantiates a new ExperimentsAnalysisPlanV2MutationResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsAnalysisPlanV2MutationResponseDataAttributes() *ExperimentsAnalysisPlanV2MutationResponseDataAttributes {
	this := ExperimentsAnalysisPlanV2MutationResponseDataAttributes{}
	return &this
}

// NewExperimentsAnalysisPlanV2MutationResponseDataAttributesWithDefaults instantiates a new ExperimentsAnalysisPlanV2MutationResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsAnalysisPlanV2MutationResponseDataAttributesWithDefaults() *ExperimentsAnalysisPlanV2MutationResponseDataAttributes {
	this := ExperimentsAnalysisPlanV2MutationResponseDataAttributes{}
	return &this
}

// GetBayesianPrior returns the BayesianPrior field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetBayesianPrior() ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior {
	if o == nil || o.BayesianPrior == nil {
		var ret ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior
		return ret
	}
	return *o.BayesianPrior
}

// GetBayesianPriorOk returns a tuple with the BayesianPrior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetBayesianPriorOk() (*ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior, bool) {
	if o == nil || o.BayesianPrior == nil {
		return nil, false
	}
	return o.BayesianPrior, true
}

// HasBayesianPrior returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) HasBayesianPrior() bool {
	return o != nil && o.BayesianPrior != nil
}

// SetBayesianPrior gets a reference to the given ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior and assigns it to the BayesianPrior field.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) SetBayesianPrior(v ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior) {
	o.BayesianPrior = &v
}

// GetConfidenceIntervalMethod returns the ConfidenceIntervalMethod field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetConfidenceIntervalMethod() ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod {
	if o == nil || o.ConfidenceIntervalMethod == nil {
		var ret ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod
		return ret
	}
	return *o.ConfidenceIntervalMethod
}

// GetConfidenceIntervalMethodOk returns a tuple with the ConfidenceIntervalMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetConfidenceIntervalMethodOk() (*ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod, bool) {
	if o == nil || o.ConfidenceIntervalMethod == nil {
		return nil, false
	}
	return o.ConfidenceIntervalMethod, true
}

// HasConfidenceIntervalMethod returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) HasConfidenceIntervalMethod() bool {
	return o != nil && o.ConfidenceIntervalMethod != nil
}

// SetConfidenceIntervalMethod gets a reference to the given ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod and assigns it to the ConfidenceIntervalMethod field.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) SetConfidenceIntervalMethod(v ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod) {
	o.ConfidenceIntervalMethod = &v
}

// GetConfidenceLevel returns the ConfidenceLevel field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetConfidenceLevel() float64 {
	if o == nil || o.ConfidenceLevel == nil {
		var ret float64
		return ret
	}
	return *o.ConfidenceLevel
}

// GetConfidenceLevelOk returns a tuple with the ConfidenceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetConfidenceLevelOk() (*float64, bool) {
	if o == nil || o.ConfidenceLevel == nil {
		return nil, false
	}
	return o.ConfidenceLevel, true
}

// HasConfidenceLevel returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) HasConfidenceLevel() bool {
	return o != nil && o.ConfidenceLevel != nil
}

// SetConfidenceLevel gets a reference to the given float64 and assigns it to the ConfidenceLevel field.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) SetConfidenceLevel(v float64) {
	o.ConfidenceLevel = &v
}

// GetCupedLookbackPeriodDays returns the CupedLookbackPeriodDays field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetCupedLookbackPeriodDays() int64 {
	if o == nil || o.CupedLookbackPeriodDays == nil {
		var ret int64
		return ret
	}
	return *o.CupedLookbackPeriodDays
}

// GetCupedLookbackPeriodDaysOk returns a tuple with the CupedLookbackPeriodDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetCupedLookbackPeriodDaysOk() (*int64, bool) {
	if o == nil || o.CupedLookbackPeriodDays == nil {
		return nil, false
	}
	return o.CupedLookbackPeriodDays, true
}

// HasCupedLookbackPeriodDays returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) HasCupedLookbackPeriodDays() bool {
	return o != nil && o.CupedLookbackPeriodDays != nil
}

// SetCupedLookbackPeriodDays gets a reference to the given int64 and assigns it to the CupedLookbackPeriodDays field.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) SetCupedLookbackPeriodDays(v int64) {
	o.CupedLookbackPeriodDays = &v
}

// GetExperimentAutoEndDays returns the ExperimentAutoEndDays field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetExperimentAutoEndDays() int64 {
	if o == nil || o.ExperimentAutoEndDays == nil {
		var ret int64
		return ret
	}
	return *o.ExperimentAutoEndDays
}

// GetExperimentAutoEndDaysOk returns a tuple with the ExperimentAutoEndDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetExperimentAutoEndDaysOk() (*int64, bool) {
	if o == nil || o.ExperimentAutoEndDays == nil {
		return nil, false
	}
	return o.ExperimentAutoEndDays, true
}

// HasExperimentAutoEndDays returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) HasExperimentAutoEndDays() bool {
	return o != nil && o.ExperimentAutoEndDays != nil
}

// SetExperimentAutoEndDays gets a reference to the given int64 and assigns it to the ExperimentAutoEndDays field.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) SetExperimentAutoEndDays(v int64) {
	o.ExperimentAutoEndDays = &v
}

// GetExperimentMinDuration returns the ExperimentMinDuration field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetExperimentMinDuration() int64 {
	if o == nil || o.ExperimentMinDuration == nil {
		var ret int64
		return ret
	}
	return *o.ExperimentMinDuration
}

// GetExperimentMinDurationOk returns a tuple with the ExperimentMinDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetExperimentMinDurationOk() (*int64, bool) {
	if o == nil || o.ExperimentMinDuration == nil {
		return nil, false
	}
	return o.ExperimentMinDuration, true
}

// HasExperimentMinDuration returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) HasExperimentMinDuration() bool {
	return o != nil && o.ExperimentMinDuration != nil
}

// SetExperimentMinDuration gets a reference to the given int64 and assigns it to the ExperimentMinDuration field.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) SetExperimentMinDuration(v int64) {
	o.ExperimentMinDuration = &v
}

// GetExperimentMinSampleSize returns the ExperimentMinSampleSize field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetExperimentMinSampleSize() int64 {
	if o == nil || o.ExperimentMinSampleSize == nil {
		var ret int64
		return ret
	}
	return *o.ExperimentMinSampleSize
}

// GetExperimentMinSampleSizeOk returns a tuple with the ExperimentMinSampleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetExperimentMinSampleSizeOk() (*int64, bool) {
	if o == nil || o.ExperimentMinSampleSize == nil {
		return nil, false
	}
	return o.ExperimentMinSampleSize, true
}

// HasExperimentMinSampleSize returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) HasExperimentMinSampleSize() bool {
	return o != nil && o.ExperimentMinSampleSize != nil
}

// SetExperimentMinSampleSize gets a reference to the given int64 and assigns it to the ExperimentMinSampleSize field.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) SetExperimentMinSampleSize(v int64) {
	o.ExperimentMinSampleSize = &v
}

// GetHasCustomAnalysisSettings returns the HasCustomAnalysisSettings field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetHasCustomAnalysisSettings() bool {
	if o == nil || o.HasCustomAnalysisSettings == nil {
		var ret bool
		return ret
	}
	return *o.HasCustomAnalysisSettings
}

// GetHasCustomAnalysisSettingsOk returns a tuple with the HasCustomAnalysisSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetHasCustomAnalysisSettingsOk() (*bool, bool) {
	if o == nil || o.HasCustomAnalysisSettings == nil {
		return nil, false
	}
	return o.HasCustomAnalysisSettings, true
}

// HasHasCustomAnalysisSettings returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) HasHasCustomAnalysisSettings() bool {
	return o != nil && o.HasCustomAnalysisSettings != nil
}

// SetHasCustomAnalysisSettings gets a reference to the given bool and assigns it to the HasCustomAnalysisSettings field.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) SetHasCustomAnalysisSettings(v bool) {
	o.HasCustomAnalysisSettings = &v
}

// GetIsCupedEnabled returns the IsCupedEnabled field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetIsCupedEnabled() bool {
	if o == nil || o.IsCupedEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsCupedEnabled
}

// GetIsCupedEnabledOk returns a tuple with the IsCupedEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetIsCupedEnabledOk() (*bool, bool) {
	if o == nil || o.IsCupedEnabled == nil {
		return nil, false
	}
	return o.IsCupedEnabled, true
}

// HasIsCupedEnabled returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) HasIsCupedEnabled() bool {
	return o != nil && o.IsCupedEnabled != nil
}

// SetIsCupedEnabled gets a reference to the given bool and assigns it to the IsCupedEnabled field.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) SetIsCupedEnabled(v bool) {
	o.IsCupedEnabled = &v
}

// GetIsMultipleTestingCorrectionEnabled returns the IsMultipleTestingCorrectionEnabled field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetIsMultipleTestingCorrectionEnabled() bool {
	if o == nil || o.IsMultipleTestingCorrectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsMultipleTestingCorrectionEnabled
}

// GetIsMultipleTestingCorrectionEnabledOk returns a tuple with the IsMultipleTestingCorrectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetIsMultipleTestingCorrectionEnabledOk() (*bool, bool) {
	if o == nil || o.IsMultipleTestingCorrectionEnabled == nil {
		return nil, false
	}
	return o.IsMultipleTestingCorrectionEnabled, true
}

// HasIsMultipleTestingCorrectionEnabled returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) HasIsMultipleTestingCorrectionEnabled() bool {
	return o != nil && o.IsMultipleTestingCorrectionEnabled != nil
}

// SetIsMultipleTestingCorrectionEnabled gets a reference to the given bool and assigns it to the IsMultipleTestingCorrectionEnabled field.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) SetIsMultipleTestingCorrectionEnabled(v bool) {
	o.IsMultipleTestingCorrectionEnabled = &v
}

// GetPreferentialBonferroniPrimaryMetricWeight returns the PreferentialBonferroniPrimaryMetricWeight field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetPreferentialBonferroniPrimaryMetricWeight() float64 {
	if o == nil || o.PreferentialBonferroniPrimaryMetricWeight == nil {
		var ret float64
		return ret
	}
	return *o.PreferentialBonferroniPrimaryMetricWeight
}

// GetPreferentialBonferroniPrimaryMetricWeightOk returns a tuple with the PreferentialBonferroniPrimaryMetricWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetPreferentialBonferroniPrimaryMetricWeightOk() (*float64, bool) {
	if o == nil || o.PreferentialBonferroniPrimaryMetricWeight == nil {
		return nil, false
	}
	return o.PreferentialBonferroniPrimaryMetricWeight, true
}

// HasPreferentialBonferroniPrimaryMetricWeight returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) HasPreferentialBonferroniPrimaryMetricWeight() bool {
	return o != nil && o.PreferentialBonferroniPrimaryMetricWeight != nil
}

// SetPreferentialBonferroniPrimaryMetricWeight gets a reference to the given float64 and assigns it to the PreferentialBonferroniPrimaryMetricWeight field.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) SetPreferentialBonferroniPrimaryMetricWeight(v float64) {
	o.PreferentialBonferroniPrimaryMetricWeight = &v
}

// GetTargetDurationDays returns the TargetDurationDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetTargetDurationDays() int64 {
	if o == nil || o.TargetDurationDays.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TargetDurationDays.Get()
}

// GetTargetDurationDaysOk returns a tuple with the TargetDurationDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) GetTargetDurationDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetDurationDays.Get(), o.TargetDurationDays.IsSet()
}

// HasTargetDurationDays returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) HasTargetDurationDays() bool {
	return o != nil && o.TargetDurationDays.IsSet()
}

// SetTargetDurationDays gets a reference to the given datadog.NullableInt64 and assigns it to the TargetDurationDays field.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) SetTargetDurationDays(v int64) {
	o.TargetDurationDays.Set(&v)
}

// SetTargetDurationDaysNil sets the value for TargetDurationDays to be an explicit nil.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) SetTargetDurationDaysNil() {
	o.TargetDurationDays.Set(nil)
}

// UnsetTargetDurationDays ensures that no value is present for TargetDurationDays, not even an explicit nil.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) UnsetTargetDurationDays() {
	o.TargetDurationDays.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsAnalysisPlanV2MutationResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BayesianPrior != nil {
		toSerialize["bayesian_prior"] = o.BayesianPrior
	}
	if o.ConfidenceIntervalMethod != nil {
		toSerialize["confidence_interval_method"] = o.ConfidenceIntervalMethod
	}
	if o.ConfidenceLevel != nil {
		toSerialize["confidence_level"] = o.ConfidenceLevel
	}
	if o.CupedLookbackPeriodDays != nil {
		toSerialize["cuped_lookback_period_days"] = o.CupedLookbackPeriodDays
	}
	if o.ExperimentAutoEndDays != nil {
		toSerialize["experiment_auto_end_days"] = o.ExperimentAutoEndDays
	}
	if o.ExperimentMinDuration != nil {
		toSerialize["experiment_min_duration"] = o.ExperimentMinDuration
	}
	if o.ExperimentMinSampleSize != nil {
		toSerialize["experiment_min_sample_size"] = o.ExperimentMinSampleSize
	}
	if o.HasCustomAnalysisSettings != nil {
		toSerialize["has_custom_analysis_settings"] = o.HasCustomAnalysisSettings
	}
	if o.IsCupedEnabled != nil {
		toSerialize["is_cuped_enabled"] = o.IsCupedEnabled
	}
	if o.IsMultipleTestingCorrectionEnabled != nil {
		toSerialize["is_multiple_testing_correction_enabled"] = o.IsMultipleTestingCorrectionEnabled
	}
	if o.PreferentialBonferroniPrimaryMetricWeight != nil {
		toSerialize["preferential_bonferroni_primary_metric_weight"] = o.PreferentialBonferroniPrimaryMetricWeight
	}
	if o.TargetDurationDays.IsSet() {
		toSerialize["target_duration_days"] = o.TargetDurationDays.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BayesianPrior                             *ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior `json:"bayesian_prior,omitempty"`
		ConfidenceIntervalMethod                  *ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod   `json:"confidence_interval_method,omitempty"`
		ConfidenceLevel                           *float64                                                              `json:"confidence_level,omitempty"`
		CupedLookbackPeriodDays                   *int64                                                                `json:"cuped_lookback_period_days,omitempty"`
		ExperimentAutoEndDays                     *int64                                                                `json:"experiment_auto_end_days,omitempty"`
		ExperimentMinDuration                     *int64                                                                `json:"experiment_min_duration,omitempty"`
		ExperimentMinSampleSize                   *int64                                                                `json:"experiment_min_sample_size,omitempty"`
		HasCustomAnalysisSettings                 *bool                                                                 `json:"has_custom_analysis_settings,omitempty"`
		IsCupedEnabled                            *bool                                                                 `json:"is_cuped_enabled,omitempty"`
		IsMultipleTestingCorrectionEnabled        *bool                                                                 `json:"is_multiple_testing_correction_enabled,omitempty"`
		PreferentialBonferroniPrimaryMetricWeight *float64                                                              `json:"preferential_bonferroni_primary_metric_weight,omitempty"`
		TargetDurationDays                        datadog.NullableInt64                                                 `json:"target_duration_days,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bayesian_prior", "confidence_interval_method", "confidence_level", "cuped_lookback_period_days", "experiment_auto_end_days", "experiment_min_duration", "experiment_min_sample_size", "has_custom_analysis_settings", "is_cuped_enabled", "is_multiple_testing_correction_enabled", "preferential_bonferroni_primary_metric_weight", "target_duration_days"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.BayesianPrior != nil && all.BayesianPrior.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BayesianPrior = all.BayesianPrior
	if all.ConfidenceIntervalMethod != nil && !all.ConfidenceIntervalMethod.IsValid() {
		hasInvalidField = true
	} else {
		o.ConfidenceIntervalMethod = all.ConfidenceIntervalMethod
	}
	o.ConfidenceLevel = all.ConfidenceLevel
	o.CupedLookbackPeriodDays = all.CupedLookbackPeriodDays
	o.ExperimentAutoEndDays = all.ExperimentAutoEndDays
	o.ExperimentMinDuration = all.ExperimentMinDuration
	o.ExperimentMinSampleSize = all.ExperimentMinSampleSize
	o.HasCustomAnalysisSettings = all.HasCustomAnalysisSettings
	o.IsCupedEnabled = all.IsCupedEnabled
	o.IsMultipleTestingCorrectionEnabled = all.IsMultipleTestingCorrectionEnabled
	o.PreferentialBonferroniPrimaryMetricWeight = all.PreferentialBonferroniPrimaryMetricWeight
	o.TargetDurationDays = all.TargetDurationDays

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
