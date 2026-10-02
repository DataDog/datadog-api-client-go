// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsAnalysisPlanWriteV2RequestDataAttributes Statistical settings and duration targets to apply to the experiment.
type ExperimentsAnalysisPlanWriteV2RequestDataAttributes struct {
	// Parameters of the prior distribution to use for Bayesian analysis.
	BayesianPrior *ExperimentsAnalysisPlanWriteV2RequestDataAttributesBayesianPrior `json:"bayesian_prior,omitempty"`
	// Statistical method used to calculate the experiment results.
	ConfidenceIntervalMethod *ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod `json:"confidence_interval_method,omitempty"`
	// Confidence level used for statistical analysis, expressed as a fraction.
	ConfidenceLevel *float64 `json:"confidence_level,omitempty"`
	// Only a 30-day CUPED lookback is supported.
	CupedLookbackPeriodDays *int64 `json:"cuped_lookback_period_days,omitempty"`
	// Number of days configured for the experiment to end automatically.
	ExperimentAutoEndDays *int64 `json:"experiment_auto_end_days,omitempty"`
	// Minimum experiment duration in days configured in the analysis plan.
	ExperimentMinDuration *int64 `json:"experiment_min_duration,omitempty"`
	// Minimum sample size configured in the analysis plan.
	ExperimentMinSampleSize *int64 `json:"experiment_min_sample_size,omitempty"`
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

// NewExperimentsAnalysisPlanWriteV2RequestDataAttributes instantiates a new ExperimentsAnalysisPlanWriteV2RequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsAnalysisPlanWriteV2RequestDataAttributes() *ExperimentsAnalysisPlanWriteV2RequestDataAttributes {
	this := ExperimentsAnalysisPlanWriteV2RequestDataAttributes{}
	return &this
}

// NewExperimentsAnalysisPlanWriteV2RequestDataAttributesWithDefaults instantiates a new ExperimentsAnalysisPlanWriteV2RequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsAnalysisPlanWriteV2RequestDataAttributesWithDefaults() *ExperimentsAnalysisPlanWriteV2RequestDataAttributes {
	this := ExperimentsAnalysisPlanWriteV2RequestDataAttributes{}
	return &this
}

// GetBayesianPrior returns the BayesianPrior field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetBayesianPrior() ExperimentsAnalysisPlanWriteV2RequestDataAttributesBayesianPrior {
	if o == nil || o.BayesianPrior == nil {
		var ret ExperimentsAnalysisPlanWriteV2RequestDataAttributesBayesianPrior
		return ret
	}
	return *o.BayesianPrior
}

// GetBayesianPriorOk returns a tuple with the BayesianPrior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetBayesianPriorOk() (*ExperimentsAnalysisPlanWriteV2RequestDataAttributesBayesianPrior, bool) {
	if o == nil || o.BayesianPrior == nil {
		return nil, false
	}
	return o.BayesianPrior, true
}

// HasBayesianPrior returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) HasBayesianPrior() bool {
	return o != nil && o.BayesianPrior != nil
}

// SetBayesianPrior gets a reference to the given ExperimentsAnalysisPlanWriteV2RequestDataAttributesBayesianPrior and assigns it to the BayesianPrior field.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) SetBayesianPrior(v ExperimentsAnalysisPlanWriteV2RequestDataAttributesBayesianPrior) {
	o.BayesianPrior = &v
}

// GetConfidenceIntervalMethod returns the ConfidenceIntervalMethod field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetConfidenceIntervalMethod() ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod {
	if o == nil || o.ConfidenceIntervalMethod == nil {
		var ret ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod
		return ret
	}
	return *o.ConfidenceIntervalMethod
}

// GetConfidenceIntervalMethodOk returns a tuple with the ConfidenceIntervalMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetConfidenceIntervalMethodOk() (*ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod, bool) {
	if o == nil || o.ConfidenceIntervalMethod == nil {
		return nil, false
	}
	return o.ConfidenceIntervalMethod, true
}

// HasConfidenceIntervalMethod returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) HasConfidenceIntervalMethod() bool {
	return o != nil && o.ConfidenceIntervalMethod != nil
}

// SetConfidenceIntervalMethod gets a reference to the given ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod and assigns it to the ConfidenceIntervalMethod field.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) SetConfidenceIntervalMethod(v ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod) {
	o.ConfidenceIntervalMethod = &v
}

// GetConfidenceLevel returns the ConfidenceLevel field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetConfidenceLevel() float64 {
	if o == nil || o.ConfidenceLevel == nil {
		var ret float64
		return ret
	}
	return *o.ConfidenceLevel
}

// GetConfidenceLevelOk returns a tuple with the ConfidenceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetConfidenceLevelOk() (*float64, bool) {
	if o == nil || o.ConfidenceLevel == nil {
		return nil, false
	}
	return o.ConfidenceLevel, true
}

// HasConfidenceLevel returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) HasConfidenceLevel() bool {
	return o != nil && o.ConfidenceLevel != nil
}

// SetConfidenceLevel gets a reference to the given float64 and assigns it to the ConfidenceLevel field.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) SetConfidenceLevel(v float64) {
	o.ConfidenceLevel = &v
}

// GetCupedLookbackPeriodDays returns the CupedLookbackPeriodDays field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetCupedLookbackPeriodDays() int64 {
	if o == nil || o.CupedLookbackPeriodDays == nil {
		var ret int64
		return ret
	}
	return *o.CupedLookbackPeriodDays
}

// GetCupedLookbackPeriodDaysOk returns a tuple with the CupedLookbackPeriodDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetCupedLookbackPeriodDaysOk() (*int64, bool) {
	if o == nil || o.CupedLookbackPeriodDays == nil {
		return nil, false
	}
	return o.CupedLookbackPeriodDays, true
}

// HasCupedLookbackPeriodDays returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) HasCupedLookbackPeriodDays() bool {
	return o != nil && o.CupedLookbackPeriodDays != nil
}

// SetCupedLookbackPeriodDays gets a reference to the given int64 and assigns it to the CupedLookbackPeriodDays field.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) SetCupedLookbackPeriodDays(v int64) {
	o.CupedLookbackPeriodDays = &v
}

// GetExperimentAutoEndDays returns the ExperimentAutoEndDays field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetExperimentAutoEndDays() int64 {
	if o == nil || o.ExperimentAutoEndDays == nil {
		var ret int64
		return ret
	}
	return *o.ExperimentAutoEndDays
}

// GetExperimentAutoEndDaysOk returns a tuple with the ExperimentAutoEndDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetExperimentAutoEndDaysOk() (*int64, bool) {
	if o == nil || o.ExperimentAutoEndDays == nil {
		return nil, false
	}
	return o.ExperimentAutoEndDays, true
}

// HasExperimentAutoEndDays returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) HasExperimentAutoEndDays() bool {
	return o != nil && o.ExperimentAutoEndDays != nil
}

// SetExperimentAutoEndDays gets a reference to the given int64 and assigns it to the ExperimentAutoEndDays field.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) SetExperimentAutoEndDays(v int64) {
	o.ExperimentAutoEndDays = &v
}

// GetExperimentMinDuration returns the ExperimentMinDuration field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetExperimentMinDuration() int64 {
	if o == nil || o.ExperimentMinDuration == nil {
		var ret int64
		return ret
	}
	return *o.ExperimentMinDuration
}

// GetExperimentMinDurationOk returns a tuple with the ExperimentMinDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetExperimentMinDurationOk() (*int64, bool) {
	if o == nil || o.ExperimentMinDuration == nil {
		return nil, false
	}
	return o.ExperimentMinDuration, true
}

// HasExperimentMinDuration returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) HasExperimentMinDuration() bool {
	return o != nil && o.ExperimentMinDuration != nil
}

// SetExperimentMinDuration gets a reference to the given int64 and assigns it to the ExperimentMinDuration field.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) SetExperimentMinDuration(v int64) {
	o.ExperimentMinDuration = &v
}

// GetExperimentMinSampleSize returns the ExperimentMinSampleSize field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetExperimentMinSampleSize() int64 {
	if o == nil || o.ExperimentMinSampleSize == nil {
		var ret int64
		return ret
	}
	return *o.ExperimentMinSampleSize
}

// GetExperimentMinSampleSizeOk returns a tuple with the ExperimentMinSampleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetExperimentMinSampleSizeOk() (*int64, bool) {
	if o == nil || o.ExperimentMinSampleSize == nil {
		return nil, false
	}
	return o.ExperimentMinSampleSize, true
}

// HasExperimentMinSampleSize returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) HasExperimentMinSampleSize() bool {
	return o != nil && o.ExperimentMinSampleSize != nil
}

// SetExperimentMinSampleSize gets a reference to the given int64 and assigns it to the ExperimentMinSampleSize field.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) SetExperimentMinSampleSize(v int64) {
	o.ExperimentMinSampleSize = &v
}

// GetIsCupedEnabled returns the IsCupedEnabled field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetIsCupedEnabled() bool {
	if o == nil || o.IsCupedEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsCupedEnabled
}

// GetIsCupedEnabledOk returns a tuple with the IsCupedEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetIsCupedEnabledOk() (*bool, bool) {
	if o == nil || o.IsCupedEnabled == nil {
		return nil, false
	}
	return o.IsCupedEnabled, true
}

// HasIsCupedEnabled returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) HasIsCupedEnabled() bool {
	return o != nil && o.IsCupedEnabled != nil
}

// SetIsCupedEnabled gets a reference to the given bool and assigns it to the IsCupedEnabled field.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) SetIsCupedEnabled(v bool) {
	o.IsCupedEnabled = &v
}

// GetIsMultipleTestingCorrectionEnabled returns the IsMultipleTestingCorrectionEnabled field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetIsMultipleTestingCorrectionEnabled() bool {
	if o == nil || o.IsMultipleTestingCorrectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsMultipleTestingCorrectionEnabled
}

// GetIsMultipleTestingCorrectionEnabledOk returns a tuple with the IsMultipleTestingCorrectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetIsMultipleTestingCorrectionEnabledOk() (*bool, bool) {
	if o == nil || o.IsMultipleTestingCorrectionEnabled == nil {
		return nil, false
	}
	return o.IsMultipleTestingCorrectionEnabled, true
}

// HasIsMultipleTestingCorrectionEnabled returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) HasIsMultipleTestingCorrectionEnabled() bool {
	return o != nil && o.IsMultipleTestingCorrectionEnabled != nil
}

// SetIsMultipleTestingCorrectionEnabled gets a reference to the given bool and assigns it to the IsMultipleTestingCorrectionEnabled field.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) SetIsMultipleTestingCorrectionEnabled(v bool) {
	o.IsMultipleTestingCorrectionEnabled = &v
}

// GetPreferentialBonferroniPrimaryMetricWeight returns the PreferentialBonferroniPrimaryMetricWeight field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetPreferentialBonferroniPrimaryMetricWeight() float64 {
	if o == nil || o.PreferentialBonferroniPrimaryMetricWeight == nil {
		var ret float64
		return ret
	}
	return *o.PreferentialBonferroniPrimaryMetricWeight
}

// GetPreferentialBonferroniPrimaryMetricWeightOk returns a tuple with the PreferentialBonferroniPrimaryMetricWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetPreferentialBonferroniPrimaryMetricWeightOk() (*float64, bool) {
	if o == nil || o.PreferentialBonferroniPrimaryMetricWeight == nil {
		return nil, false
	}
	return o.PreferentialBonferroniPrimaryMetricWeight, true
}

// HasPreferentialBonferroniPrimaryMetricWeight returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) HasPreferentialBonferroniPrimaryMetricWeight() bool {
	return o != nil && o.PreferentialBonferroniPrimaryMetricWeight != nil
}

// SetPreferentialBonferroniPrimaryMetricWeight gets a reference to the given float64 and assigns it to the PreferentialBonferroniPrimaryMetricWeight field.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) SetPreferentialBonferroniPrimaryMetricWeight(v float64) {
	o.PreferentialBonferroniPrimaryMetricWeight = &v
}

// GetTargetDurationDays returns the TargetDurationDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetTargetDurationDays() int64 {
	if o == nil || o.TargetDurationDays.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TargetDurationDays.Get()
}

// GetTargetDurationDaysOk returns a tuple with the TargetDurationDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) GetTargetDurationDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetDurationDays.Get(), o.TargetDurationDays.IsSet()
}

// HasTargetDurationDays returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) HasTargetDurationDays() bool {
	return o != nil && o.TargetDurationDays.IsSet()
}

// SetTargetDurationDays gets a reference to the given datadog.NullableInt64 and assigns it to the TargetDurationDays field.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) SetTargetDurationDays(v int64) {
	o.TargetDurationDays.Set(&v)
}

// SetTargetDurationDaysNil sets the value for TargetDurationDays to be an explicit nil.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) SetTargetDurationDaysNil() {
	o.TargetDurationDays.Set(nil)
}

// UnsetTargetDurationDays ensures that no value is present for TargetDurationDays, not even an explicit nil.
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) UnsetTargetDurationDays() {
	o.TargetDurationDays.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsAnalysisPlanWriteV2RequestDataAttributes) MarshalJSON() ([]byte, error) {
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
func (o *ExperimentsAnalysisPlanWriteV2RequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BayesianPrior                             *ExperimentsAnalysisPlanWriteV2RequestDataAttributesBayesianPrior   `json:"bayesian_prior,omitempty"`
		ConfidenceIntervalMethod                  *ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod `json:"confidence_interval_method,omitempty"`
		ConfidenceLevel                           *float64                                                            `json:"confidence_level,omitempty"`
		CupedLookbackPeriodDays                   *int64                                                              `json:"cuped_lookback_period_days,omitempty"`
		ExperimentAutoEndDays                     *int64                                                              `json:"experiment_auto_end_days,omitempty"`
		ExperimentMinDuration                     *int64                                                              `json:"experiment_min_duration,omitempty"`
		ExperimentMinSampleSize                   *int64                                                              `json:"experiment_min_sample_size,omitempty"`
		IsCupedEnabled                            *bool                                                               `json:"is_cuped_enabled,omitempty"`
		IsMultipleTestingCorrectionEnabled        *bool                                                               `json:"is_multiple_testing_correction_enabled,omitempty"`
		PreferentialBonferroniPrimaryMetricWeight *float64                                                            `json:"preferential_bonferroni_primary_metric_weight,omitempty"`
		TargetDurationDays                        datadog.NullableInt64                                               `json:"target_duration_days,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bayesian_prior", "confidence_interval_method", "confidence_level", "cuped_lookback_period_days", "experiment_auto_end_days", "experiment_min_duration", "experiment_min_sample_size", "is_cuped_enabled", "is_multiple_testing_correction_enabled", "preferential_bonferroni_primary_metric_weight", "target_duration_days"})
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
