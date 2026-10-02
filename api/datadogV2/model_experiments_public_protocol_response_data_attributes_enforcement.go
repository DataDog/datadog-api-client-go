// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPublicProtocolResponseDataAttributesEnforcement Controls that determine which protocol settings can be changed in an experiment.
type ExperimentsPublicProtocolResponseDataAttributesEnforcement struct {
	// LOCKED prevents changes to confidence interval method. EDITABLE permits changes.
	ConfidenceIntervalMethod *string `json:"confidence_interval_method,omitempty"`
	// LOCKED prevents changes to confidence level. EDITABLE permits changes.
	ConfidenceLevel *string `json:"confidence_level,omitempty"`
	// LOCKED prevents changes to CUPED variance reduction. EDITABLE permits changes.
	CupedCalculation *string `json:"cuped_calculation,omitempty"`
	// LOCKED prevents changes to default duration. EDITABLE permits changes.
	DefaultDuration *string `json:"default_duration,omitempty"`
	// LOCKED prevents changes to environment. EDITABLE permits changes.
	Environment *string `json:"environment,omitempty"`
	// LOCKED prevents changes to feature flag source. EDITABLE permits changes.
	FlagSource *string `json:"flag_source,omitempty"`
	// LOCKED prevents changes to multiple testing correction. EDITABLE permits changes.
	MultipleTestingCorrection *string `json:"multiple_testing_correction,omitempty"`
	// LOCKED prevents changes to notifications. EDITABLE permits changes.
	Notifications *string `json:"notifications,omitempty"`
	// LOCKED prevents changes to primary metric. EDITABLE permits changes.
	PrimaryMetric *string `json:"primary_metric,omitempty"`
	// LOCKED prevents changes to secondary metrics. EDITABLE permits changes.
	SecondaryMetrics *string `json:"secondary_metrics,omitempty"`
	// LOCKED prevents changes to result exploration dimensions. EDITABLE permits changes.
	SplitByExplorationDimensions *string `json:"split_by_exploration_dimensions,omitempty"`
	// LOCKED prevents changes to subject type. EDITABLE permits changes.
	SubjectType *string `json:"subject_type,omitempty"`
	// LOCKED prevents changes to targeting rules. EDITABLE permits changes.
	TargetingRules *string `json:"targeting_rules,omitempty"`
	// LOCKED prevents changes to traffic exposure. EDITABLE permits changes.
	TrafficExposure *string `json:"traffic_exposure,omitempty"`
	// LOCKED prevents changes to warehouse exposure source. EDITABLE permits changes.
	WarehouseExposureSource *string `json:"warehouse_exposure_source,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPublicProtocolResponseDataAttributesEnforcement instantiates a new ExperimentsPublicProtocolResponseDataAttributesEnforcement object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPublicProtocolResponseDataAttributesEnforcement() *ExperimentsPublicProtocolResponseDataAttributesEnforcement {
	this := ExperimentsPublicProtocolResponseDataAttributesEnforcement{}
	return &this
}

// NewExperimentsPublicProtocolResponseDataAttributesEnforcementWithDefaults instantiates a new ExperimentsPublicProtocolResponseDataAttributesEnforcement object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPublicProtocolResponseDataAttributesEnforcementWithDefaults() *ExperimentsPublicProtocolResponseDataAttributesEnforcement {
	this := ExperimentsPublicProtocolResponseDataAttributesEnforcement{}
	return &this
}

// GetConfidenceIntervalMethod returns the ConfidenceIntervalMethod field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetConfidenceIntervalMethod() string {
	if o == nil || o.ConfidenceIntervalMethod == nil {
		var ret string
		return ret
	}
	return *o.ConfidenceIntervalMethod
}

// GetConfidenceIntervalMethodOk returns a tuple with the ConfidenceIntervalMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetConfidenceIntervalMethodOk() (*string, bool) {
	if o == nil || o.ConfidenceIntervalMethod == nil {
		return nil, false
	}
	return o.ConfidenceIntervalMethod, true
}

// HasConfidenceIntervalMethod returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) HasConfidenceIntervalMethod() bool {
	return o != nil && o.ConfidenceIntervalMethod != nil
}

// SetConfidenceIntervalMethod gets a reference to the given string and assigns it to the ConfidenceIntervalMethod field.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) SetConfidenceIntervalMethod(v string) {
	o.ConfidenceIntervalMethod = &v
}

// GetConfidenceLevel returns the ConfidenceLevel field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetConfidenceLevel() string {
	if o == nil || o.ConfidenceLevel == nil {
		var ret string
		return ret
	}
	return *o.ConfidenceLevel
}

// GetConfidenceLevelOk returns a tuple with the ConfidenceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetConfidenceLevelOk() (*string, bool) {
	if o == nil || o.ConfidenceLevel == nil {
		return nil, false
	}
	return o.ConfidenceLevel, true
}

// HasConfidenceLevel returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) HasConfidenceLevel() bool {
	return o != nil && o.ConfidenceLevel != nil
}

// SetConfidenceLevel gets a reference to the given string and assigns it to the ConfidenceLevel field.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) SetConfidenceLevel(v string) {
	o.ConfidenceLevel = &v
}

// GetCupedCalculation returns the CupedCalculation field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetCupedCalculation() string {
	if o == nil || o.CupedCalculation == nil {
		var ret string
		return ret
	}
	return *o.CupedCalculation
}

// GetCupedCalculationOk returns a tuple with the CupedCalculation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetCupedCalculationOk() (*string, bool) {
	if o == nil || o.CupedCalculation == nil {
		return nil, false
	}
	return o.CupedCalculation, true
}

// HasCupedCalculation returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) HasCupedCalculation() bool {
	return o != nil && o.CupedCalculation != nil
}

// SetCupedCalculation gets a reference to the given string and assigns it to the CupedCalculation field.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) SetCupedCalculation(v string) {
	o.CupedCalculation = &v
}

// GetDefaultDuration returns the DefaultDuration field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetDefaultDuration() string {
	if o == nil || o.DefaultDuration == nil {
		var ret string
		return ret
	}
	return *o.DefaultDuration
}

// GetDefaultDurationOk returns a tuple with the DefaultDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetDefaultDurationOk() (*string, bool) {
	if o == nil || o.DefaultDuration == nil {
		return nil, false
	}
	return o.DefaultDuration, true
}

// HasDefaultDuration returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) HasDefaultDuration() bool {
	return o != nil && o.DefaultDuration != nil
}

// SetDefaultDuration gets a reference to the given string and assigns it to the DefaultDuration field.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) SetDefaultDuration(v string) {
	o.DefaultDuration = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetEnvironment() string {
	if o == nil || o.Environment == nil {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetEnvironmentOk() (*string, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) HasEnvironment() bool {
	return o != nil && o.Environment != nil
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) SetEnvironment(v string) {
	o.Environment = &v
}

// GetFlagSource returns the FlagSource field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetFlagSource() string {
	if o == nil || o.FlagSource == nil {
		var ret string
		return ret
	}
	return *o.FlagSource
}

// GetFlagSourceOk returns a tuple with the FlagSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetFlagSourceOk() (*string, bool) {
	if o == nil || o.FlagSource == nil {
		return nil, false
	}
	return o.FlagSource, true
}

// HasFlagSource returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) HasFlagSource() bool {
	return o != nil && o.FlagSource != nil
}

// SetFlagSource gets a reference to the given string and assigns it to the FlagSource field.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) SetFlagSource(v string) {
	o.FlagSource = &v
}

// GetMultipleTestingCorrection returns the MultipleTestingCorrection field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetMultipleTestingCorrection() string {
	if o == nil || o.MultipleTestingCorrection == nil {
		var ret string
		return ret
	}
	return *o.MultipleTestingCorrection
}

// GetMultipleTestingCorrectionOk returns a tuple with the MultipleTestingCorrection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetMultipleTestingCorrectionOk() (*string, bool) {
	if o == nil || o.MultipleTestingCorrection == nil {
		return nil, false
	}
	return o.MultipleTestingCorrection, true
}

// HasMultipleTestingCorrection returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) HasMultipleTestingCorrection() bool {
	return o != nil && o.MultipleTestingCorrection != nil
}

// SetMultipleTestingCorrection gets a reference to the given string and assigns it to the MultipleTestingCorrection field.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) SetMultipleTestingCorrection(v string) {
	o.MultipleTestingCorrection = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetNotifications() string {
	if o == nil || o.Notifications == nil {
		var ret string
		return ret
	}
	return *o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetNotificationsOk() (*string, bool) {
	if o == nil || o.Notifications == nil {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) HasNotifications() bool {
	return o != nil && o.Notifications != nil
}

// SetNotifications gets a reference to the given string and assigns it to the Notifications field.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) SetNotifications(v string) {
	o.Notifications = &v
}

// GetPrimaryMetric returns the PrimaryMetric field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetPrimaryMetric() string {
	if o == nil || o.PrimaryMetric == nil {
		var ret string
		return ret
	}
	return *o.PrimaryMetric
}

// GetPrimaryMetricOk returns a tuple with the PrimaryMetric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetPrimaryMetricOk() (*string, bool) {
	if o == nil || o.PrimaryMetric == nil {
		return nil, false
	}
	return o.PrimaryMetric, true
}

// HasPrimaryMetric returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) HasPrimaryMetric() bool {
	return o != nil && o.PrimaryMetric != nil
}

// SetPrimaryMetric gets a reference to the given string and assigns it to the PrimaryMetric field.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) SetPrimaryMetric(v string) {
	o.PrimaryMetric = &v
}

// GetSecondaryMetrics returns the SecondaryMetrics field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetSecondaryMetrics() string {
	if o == nil || o.SecondaryMetrics == nil {
		var ret string
		return ret
	}
	return *o.SecondaryMetrics
}

// GetSecondaryMetricsOk returns a tuple with the SecondaryMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetSecondaryMetricsOk() (*string, bool) {
	if o == nil || o.SecondaryMetrics == nil {
		return nil, false
	}
	return o.SecondaryMetrics, true
}

// HasSecondaryMetrics returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) HasSecondaryMetrics() bool {
	return o != nil && o.SecondaryMetrics != nil
}

// SetSecondaryMetrics gets a reference to the given string and assigns it to the SecondaryMetrics field.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) SetSecondaryMetrics(v string) {
	o.SecondaryMetrics = &v
}

// GetSplitByExplorationDimensions returns the SplitByExplorationDimensions field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetSplitByExplorationDimensions() string {
	if o == nil || o.SplitByExplorationDimensions == nil {
		var ret string
		return ret
	}
	return *o.SplitByExplorationDimensions
}

// GetSplitByExplorationDimensionsOk returns a tuple with the SplitByExplorationDimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetSplitByExplorationDimensionsOk() (*string, bool) {
	if o == nil || o.SplitByExplorationDimensions == nil {
		return nil, false
	}
	return o.SplitByExplorationDimensions, true
}

// HasSplitByExplorationDimensions returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) HasSplitByExplorationDimensions() bool {
	return o != nil && o.SplitByExplorationDimensions != nil
}

// SetSplitByExplorationDimensions gets a reference to the given string and assigns it to the SplitByExplorationDimensions field.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) SetSplitByExplorationDimensions(v string) {
	o.SplitByExplorationDimensions = &v
}

// GetSubjectType returns the SubjectType field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetSubjectType() string {
	if o == nil || o.SubjectType == nil {
		var ret string
		return ret
	}
	return *o.SubjectType
}

// GetSubjectTypeOk returns a tuple with the SubjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetSubjectTypeOk() (*string, bool) {
	if o == nil || o.SubjectType == nil {
		return nil, false
	}
	return o.SubjectType, true
}

// HasSubjectType returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) HasSubjectType() bool {
	return o != nil && o.SubjectType != nil
}

// SetSubjectType gets a reference to the given string and assigns it to the SubjectType field.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) SetSubjectType(v string) {
	o.SubjectType = &v
}

// GetTargetingRules returns the TargetingRules field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetTargetingRules() string {
	if o == nil || o.TargetingRules == nil {
		var ret string
		return ret
	}
	return *o.TargetingRules
}

// GetTargetingRulesOk returns a tuple with the TargetingRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetTargetingRulesOk() (*string, bool) {
	if o == nil || o.TargetingRules == nil {
		return nil, false
	}
	return o.TargetingRules, true
}

// HasTargetingRules returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) HasTargetingRules() bool {
	return o != nil && o.TargetingRules != nil
}

// SetTargetingRules gets a reference to the given string and assigns it to the TargetingRules field.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) SetTargetingRules(v string) {
	o.TargetingRules = &v
}

// GetTrafficExposure returns the TrafficExposure field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetTrafficExposure() string {
	if o == nil || o.TrafficExposure == nil {
		var ret string
		return ret
	}
	return *o.TrafficExposure
}

// GetTrafficExposureOk returns a tuple with the TrafficExposure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetTrafficExposureOk() (*string, bool) {
	if o == nil || o.TrafficExposure == nil {
		return nil, false
	}
	return o.TrafficExposure, true
}

// HasTrafficExposure returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) HasTrafficExposure() bool {
	return o != nil && o.TrafficExposure != nil
}

// SetTrafficExposure gets a reference to the given string and assigns it to the TrafficExposure field.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) SetTrafficExposure(v string) {
	o.TrafficExposure = &v
}

// GetWarehouseExposureSource returns the WarehouseExposureSource field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetWarehouseExposureSource() string {
	if o == nil || o.WarehouseExposureSource == nil {
		var ret string
		return ret
	}
	return *o.WarehouseExposureSource
}

// GetWarehouseExposureSourceOk returns a tuple with the WarehouseExposureSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) GetWarehouseExposureSourceOk() (*string, bool) {
	if o == nil || o.WarehouseExposureSource == nil {
		return nil, false
	}
	return o.WarehouseExposureSource, true
}

// HasWarehouseExposureSource returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) HasWarehouseExposureSource() bool {
	return o != nil && o.WarehouseExposureSource != nil
}

// SetWarehouseExposureSource gets a reference to the given string and assigns it to the WarehouseExposureSource field.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) SetWarehouseExposureSource(v string) {
	o.WarehouseExposureSource = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPublicProtocolResponseDataAttributesEnforcement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConfidenceIntervalMethod != nil {
		toSerialize["confidence_interval_method"] = o.ConfidenceIntervalMethod
	}
	if o.ConfidenceLevel != nil {
		toSerialize["confidence_level"] = o.ConfidenceLevel
	}
	if o.CupedCalculation != nil {
		toSerialize["cuped_calculation"] = o.CupedCalculation
	}
	if o.DefaultDuration != nil {
		toSerialize["default_duration"] = o.DefaultDuration
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.FlagSource != nil {
		toSerialize["flag_source"] = o.FlagSource
	}
	if o.MultipleTestingCorrection != nil {
		toSerialize["multiple_testing_correction"] = o.MultipleTestingCorrection
	}
	if o.Notifications != nil {
		toSerialize["notifications"] = o.Notifications
	}
	if o.PrimaryMetric != nil {
		toSerialize["primary_metric"] = o.PrimaryMetric
	}
	if o.SecondaryMetrics != nil {
		toSerialize["secondary_metrics"] = o.SecondaryMetrics
	}
	if o.SplitByExplorationDimensions != nil {
		toSerialize["split_by_exploration_dimensions"] = o.SplitByExplorationDimensions
	}
	if o.SubjectType != nil {
		toSerialize["subject_type"] = o.SubjectType
	}
	if o.TargetingRules != nil {
		toSerialize["targeting_rules"] = o.TargetingRules
	}
	if o.TrafficExposure != nil {
		toSerialize["traffic_exposure"] = o.TrafficExposure
	}
	if o.WarehouseExposureSource != nil {
		toSerialize["warehouse_exposure_source"] = o.WarehouseExposureSource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsPublicProtocolResponseDataAttributesEnforcement) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfidenceIntervalMethod     *string `json:"confidence_interval_method,omitempty"`
		ConfidenceLevel              *string `json:"confidence_level,omitempty"`
		CupedCalculation             *string `json:"cuped_calculation,omitempty"`
		DefaultDuration              *string `json:"default_duration,omitempty"`
		Environment                  *string `json:"environment,omitempty"`
		FlagSource                   *string `json:"flag_source,omitempty"`
		MultipleTestingCorrection    *string `json:"multiple_testing_correction,omitempty"`
		Notifications                *string `json:"notifications,omitempty"`
		PrimaryMetric                *string `json:"primary_metric,omitempty"`
		SecondaryMetrics             *string `json:"secondary_metrics,omitempty"`
		SplitByExplorationDimensions *string `json:"split_by_exploration_dimensions,omitempty"`
		SubjectType                  *string `json:"subject_type,omitempty"`
		TargetingRules               *string `json:"targeting_rules,omitempty"`
		TrafficExposure              *string `json:"traffic_exposure,omitempty"`
		WarehouseExposureSource      *string `json:"warehouse_exposure_source,omitempty"`
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
