// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsReplaceProtocolEnforcementInput Complete lock map. Omit this object to preserve current locks. Null is rejected; every field is required.
type ExperimentsReplaceProtocolEnforcementInput struct {
	// Always LOCKED. The server keeps this setting locked regardless of the supplied value.
	ConfidenceIntervalMethod string `json:"confidence_interval_method"`
	// LOCKED prevents changes to confidence level. EDITABLE permits changes.
	ConfidenceLevel string `json:"confidence_level"`
	// LOCKED prevents changes to CUPED variance reduction. EDITABLE permits changes.
	CupedCalculation string `json:"cuped_calculation"`
	// LOCKED prevents changes to default duration. EDITABLE permits changes.
	DefaultDuration string `json:"default_duration"`
	// LOCKED prevents changes to environment. EDITABLE permits changes.
	Environment string `json:"environment"`
	// Always LOCKED. The server keeps this setting locked regardless of the supplied value.
	FlagSource string `json:"flag_source"`
	// LOCKED prevents changes to multiple testing correction. EDITABLE permits changes.
	MultipleTestingCorrection string `json:"multiple_testing_correction"`
	// LOCKED prevents changes to notifications. EDITABLE permits changes.
	Notifications string `json:"notifications"`
	// LOCKED prevents changes to primary metric. EDITABLE permits changes.
	PrimaryMetric string `json:"primary_metric"`
	// LOCKED prevents changes to secondary metrics. EDITABLE permits changes.
	SecondaryMetrics string `json:"secondary_metrics"`
	// LOCKED prevents changes to result exploration dimensions. EDITABLE permits changes.
	SplitByExplorationDimensions string `json:"split_by_exploration_dimensions"`
	// Always LOCKED. The server keeps this setting locked regardless of the supplied value.
	SubjectType string `json:"subject_type"`
	// LOCKED prevents changes to targeting rules. EDITABLE permits changes.
	TargetingRules string `json:"targeting_rules"`
	// LOCKED prevents changes to traffic exposure. EDITABLE permits changes.
	TrafficExposure string `json:"traffic_exposure"`
	// LOCKED prevents changes to warehouse exposure source. EDITABLE permits changes.
	WarehouseExposureSource string `json:"warehouse_exposure_source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsReplaceProtocolEnforcementInput instantiates a new ExperimentsReplaceProtocolEnforcementInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsReplaceProtocolEnforcementInput(confidenceIntervalMethod string, confidenceLevel string, cupedCalculation string, defaultDuration string, environment string, flagSource string, multipleTestingCorrection string, notifications string, primaryMetric string, secondaryMetrics string, splitByExplorationDimensions string, subjectType string, targetingRules string, trafficExposure string, warehouseExposureSource string) *ExperimentsReplaceProtocolEnforcementInput {
	this := ExperimentsReplaceProtocolEnforcementInput{}
	this.ConfidenceIntervalMethod = confidenceIntervalMethod
	this.ConfidenceLevel = confidenceLevel
	this.CupedCalculation = cupedCalculation
	this.DefaultDuration = defaultDuration
	this.Environment = environment
	this.FlagSource = flagSource
	this.MultipleTestingCorrection = multipleTestingCorrection
	this.Notifications = notifications
	this.PrimaryMetric = primaryMetric
	this.SecondaryMetrics = secondaryMetrics
	this.SplitByExplorationDimensions = splitByExplorationDimensions
	this.SubjectType = subjectType
	this.TargetingRules = targetingRules
	this.TrafficExposure = trafficExposure
	this.WarehouseExposureSource = warehouseExposureSource
	return &this
}

// NewExperimentsReplaceProtocolEnforcementInputWithDefaults instantiates a new ExperimentsReplaceProtocolEnforcementInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsReplaceProtocolEnforcementInputWithDefaults() *ExperimentsReplaceProtocolEnforcementInput {
	this := ExperimentsReplaceProtocolEnforcementInput{}
	return &this
}

// GetConfidenceIntervalMethod returns the ConfidenceIntervalMethod field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetConfidenceIntervalMethod() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfidenceIntervalMethod
}

// GetConfidenceIntervalMethodOk returns a tuple with the ConfidenceIntervalMethod field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetConfidenceIntervalMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfidenceIntervalMethod, true
}

// SetConfidenceIntervalMethod sets field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) SetConfidenceIntervalMethod(v string) {
	o.ConfidenceIntervalMethod = v
}

// GetConfidenceLevel returns the ConfidenceLevel field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetConfidenceLevel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfidenceLevel
}

// GetConfidenceLevelOk returns a tuple with the ConfidenceLevel field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetConfidenceLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfidenceLevel, true
}

// SetConfidenceLevel sets field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) SetConfidenceLevel(v string) {
	o.ConfidenceLevel = v
}

// GetCupedCalculation returns the CupedCalculation field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetCupedCalculation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CupedCalculation
}

// GetCupedCalculationOk returns a tuple with the CupedCalculation field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetCupedCalculationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CupedCalculation, true
}

// SetCupedCalculation sets field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) SetCupedCalculation(v string) {
	o.CupedCalculation = v
}

// GetDefaultDuration returns the DefaultDuration field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetDefaultDuration() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DefaultDuration
}

// GetDefaultDurationOk returns a tuple with the DefaultDuration field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetDefaultDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultDuration, true
}

// SetDefaultDuration sets field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) SetDefaultDuration(v string) {
	o.DefaultDuration = v
}

// GetEnvironment returns the Environment field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) SetEnvironment(v string) {
	o.Environment = v
}

// GetFlagSource returns the FlagSource field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetFlagSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FlagSource
}

// GetFlagSourceOk returns a tuple with the FlagSource field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetFlagSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlagSource, true
}

// SetFlagSource sets field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) SetFlagSource(v string) {
	o.FlagSource = v
}

// GetMultipleTestingCorrection returns the MultipleTestingCorrection field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetMultipleTestingCorrection() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MultipleTestingCorrection
}

// GetMultipleTestingCorrectionOk returns a tuple with the MultipleTestingCorrection field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetMultipleTestingCorrectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MultipleTestingCorrection, true
}

// SetMultipleTestingCorrection sets field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) SetMultipleTestingCorrection(v string) {
	o.MultipleTestingCorrection = v
}

// GetNotifications returns the Notifications field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetNotifications() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetNotificationsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notifications, true
}

// SetNotifications sets field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) SetNotifications(v string) {
	o.Notifications = v
}

// GetPrimaryMetric returns the PrimaryMetric field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetPrimaryMetric() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PrimaryMetric
}

// GetPrimaryMetricOk returns a tuple with the PrimaryMetric field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetPrimaryMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryMetric, true
}

// SetPrimaryMetric sets field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) SetPrimaryMetric(v string) {
	o.PrimaryMetric = v
}

// GetSecondaryMetrics returns the SecondaryMetrics field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetSecondaryMetrics() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SecondaryMetrics
}

// GetSecondaryMetricsOk returns a tuple with the SecondaryMetrics field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetSecondaryMetricsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecondaryMetrics, true
}

// SetSecondaryMetrics sets field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) SetSecondaryMetrics(v string) {
	o.SecondaryMetrics = v
}

// GetSplitByExplorationDimensions returns the SplitByExplorationDimensions field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetSplitByExplorationDimensions() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SplitByExplorationDimensions
}

// GetSplitByExplorationDimensionsOk returns a tuple with the SplitByExplorationDimensions field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetSplitByExplorationDimensionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SplitByExplorationDimensions, true
}

// SetSplitByExplorationDimensions sets field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) SetSplitByExplorationDimensions(v string) {
	o.SplitByExplorationDimensions = v
}

// GetSubjectType returns the SubjectType field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetSubjectType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SubjectType
}

// GetSubjectTypeOk returns a tuple with the SubjectType field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetSubjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectType, true
}

// SetSubjectType sets field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) SetSubjectType(v string) {
	o.SubjectType = v
}

// GetTargetingRules returns the TargetingRules field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetTargetingRules() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TargetingRules
}

// GetTargetingRulesOk returns a tuple with the TargetingRules field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetTargetingRulesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingRules, true
}

// SetTargetingRules sets field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) SetTargetingRules(v string) {
	o.TargetingRules = v
}

// GetTrafficExposure returns the TrafficExposure field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetTrafficExposure() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TrafficExposure
}

// GetTrafficExposureOk returns a tuple with the TrafficExposure field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetTrafficExposureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrafficExposure, true
}

// SetTrafficExposure sets field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) SetTrafficExposure(v string) {
	o.TrafficExposure = v
}

// GetWarehouseExposureSource returns the WarehouseExposureSource field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetWarehouseExposureSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.WarehouseExposureSource
}

// GetWarehouseExposureSourceOk returns a tuple with the WarehouseExposureSource field value
// and a boolean to check if the value has been set.
func (o *ExperimentsReplaceProtocolEnforcementInput) GetWarehouseExposureSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WarehouseExposureSource, true
}

// SetWarehouseExposureSource sets field value.
func (o *ExperimentsReplaceProtocolEnforcementInput) SetWarehouseExposureSource(v string) {
	o.WarehouseExposureSource = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsReplaceProtocolEnforcementInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["confidence_interval_method"] = o.ConfidenceIntervalMethod
	toSerialize["confidence_level"] = o.ConfidenceLevel
	toSerialize["cuped_calculation"] = o.CupedCalculation
	toSerialize["default_duration"] = o.DefaultDuration
	toSerialize["environment"] = o.Environment
	toSerialize["flag_source"] = o.FlagSource
	toSerialize["multiple_testing_correction"] = o.MultipleTestingCorrection
	toSerialize["notifications"] = o.Notifications
	toSerialize["primary_metric"] = o.PrimaryMetric
	toSerialize["secondary_metrics"] = o.SecondaryMetrics
	toSerialize["split_by_exploration_dimensions"] = o.SplitByExplorationDimensions
	toSerialize["subject_type"] = o.SubjectType
	toSerialize["targeting_rules"] = o.TargetingRules
	toSerialize["traffic_exposure"] = o.TrafficExposure
	toSerialize["warehouse_exposure_source"] = o.WarehouseExposureSource

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsReplaceProtocolEnforcementInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfidenceIntervalMethod     *string `json:"confidence_interval_method"`
		ConfidenceLevel              *string `json:"confidence_level"`
		CupedCalculation             *string `json:"cuped_calculation"`
		DefaultDuration              *string `json:"default_duration"`
		Environment                  *string `json:"environment"`
		FlagSource                   *string `json:"flag_source"`
		MultipleTestingCorrection    *string `json:"multiple_testing_correction"`
		Notifications                *string `json:"notifications"`
		PrimaryMetric                *string `json:"primary_metric"`
		SecondaryMetrics             *string `json:"secondary_metrics"`
		SplitByExplorationDimensions *string `json:"split_by_exploration_dimensions"`
		SubjectType                  *string `json:"subject_type"`
		TargetingRules               *string `json:"targeting_rules"`
		TrafficExposure              *string `json:"traffic_exposure"`
		WarehouseExposureSource      *string `json:"warehouse_exposure_source"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConfidenceIntervalMethod == nil {
		return fmt.Errorf("required field confidence_interval_method missing")
	}
	if all.ConfidenceLevel == nil {
		return fmt.Errorf("required field confidence_level missing")
	}
	if all.CupedCalculation == nil {
		return fmt.Errorf("required field cuped_calculation missing")
	}
	if all.DefaultDuration == nil {
		return fmt.Errorf("required field default_duration missing")
	}
	if all.Environment == nil {
		return fmt.Errorf("required field environment missing")
	}
	if all.FlagSource == nil {
		return fmt.Errorf("required field flag_source missing")
	}
	if all.MultipleTestingCorrection == nil {
		return fmt.Errorf("required field multiple_testing_correction missing")
	}
	if all.Notifications == nil {
		return fmt.Errorf("required field notifications missing")
	}
	if all.PrimaryMetric == nil {
		return fmt.Errorf("required field primary_metric missing")
	}
	if all.SecondaryMetrics == nil {
		return fmt.Errorf("required field secondary_metrics missing")
	}
	if all.SplitByExplorationDimensions == nil {
		return fmt.Errorf("required field split_by_exploration_dimensions missing")
	}
	if all.SubjectType == nil {
		return fmt.Errorf("required field subject_type missing")
	}
	if all.TargetingRules == nil {
		return fmt.Errorf("required field targeting_rules missing")
	}
	if all.TrafficExposure == nil {
		return fmt.Errorf("required field traffic_exposure missing")
	}
	if all.WarehouseExposureSource == nil {
		return fmt.Errorf("required field warehouse_exposure_source missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"confidence_interval_method", "confidence_level", "cuped_calculation", "default_duration", "environment", "flag_source", "multiple_testing_correction", "notifications", "primary_metric", "secondary_metrics", "split_by_exploration_dimensions", "subject_type", "targeting_rules", "traffic_exposure", "warehouse_exposure_source"})
	} else {
		return err
	}
	o.ConfidenceIntervalMethod = *all.ConfidenceIntervalMethod
	o.ConfidenceLevel = *all.ConfidenceLevel
	o.CupedCalculation = *all.CupedCalculation
	o.DefaultDuration = *all.DefaultDuration
	o.Environment = *all.Environment
	o.FlagSource = *all.FlagSource
	o.MultipleTestingCorrection = *all.MultipleTestingCorrection
	o.Notifications = *all.Notifications
	o.PrimaryMetric = *all.PrimaryMetric
	o.SecondaryMetrics = *all.SecondaryMetrics
	o.SplitByExplorationDimensions = *all.SplitByExplorationDimensions
	o.SubjectType = *all.SubjectType
	o.TargetingRules = *all.TargetingRules
	o.TrafficExposure = *all.TrafficExposure
	o.WarehouseExposureSource = *all.WarehouseExposureSource

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
