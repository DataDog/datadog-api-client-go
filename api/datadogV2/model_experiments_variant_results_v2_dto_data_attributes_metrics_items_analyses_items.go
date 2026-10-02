// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems One statistical comparison for a metric and variant.
type ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems struct {
	// Lower and upper bounds of the reported statistical interval.
	ConfidenceInterval *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval `json:"confidence_interval,omitempty"`
	// Configured nominal confidence level. Interval bounds can use an adjusted level for multiple testing or hybrid methods.
	ConfidenceLevel *float64 `json:"confidence_level,omitempty"`
	// Expected reduction in minimum expected regret from collecting more sample data.
	Evsi *float64 `json:"evsi,omitempty"`
	// Expected effect when the effect is positive.
	ExpectationAboveZero *float64 `json:"expectation_above_zero,omitempty"`
	// Expected effect when the effect is negative.
	ExpectationBelowZero *float64 `json:"expectation_below_zero,omitempty"`
	// Estimated lift across the population. Calculated as metric coverage multiplied by the experiment lift.
	GlobalLift *float64 `json:"global_lift,omitempty"`
	// Lower bound of the estimated lift across the population.
	GlobalLiftLowerBound *float64 `json:"global_lift_lower_bound,omitempty"`
	// Upper bound of the estimated lift across the population.
	GlobalLiftUpperBound *float64 `json:"global_lift_upper_bound,omitempty"`
	// Whether CUPED used pre-experiment data to reduce variance in this result.
	IsCupedAdjusted *bool `json:"is_cuped_adjusted,omitempty"`
	// Whether the statistical result is marked as unreliable.
	IsUnreliable *bool `json:"is_unreliable,omitempty"`
	// Whether the reported lift is relative or absolute.
	LiftType *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType `json:"lift_type,omitempty"`
	// Statistical method used to calculate this result.
	Method *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod `json:"method,omitempty"`
	// The smaller expected opportunity cost of choosing treatment or control.
	MinimumExpectedRegret *float64 `json:"minimum_expected_regret,omitempty"`
	// Probability, under the no-effect hypothesis, of a result at least as extreme as the observed result.
	PValue *float64 `json:"p_value,omitempty"`
	// Estimated difference between the variant and control for this metric.
	PointEstimate *float64 `json:"point_estimate,omitempty"`
	// Estimated probability that the effect is greater than zero.
	ProbabilityAboveZero *float64 `json:"probability_above_zero,omitempty"`
	// Estimated probability that the effect is less than zero.
	ProbabilityBelowZero *float64 `json:"probability_below_zero,omitempty"`
	// Estimated uncertainty in the effect estimate.
	StandardError *float64 `json:"standard_error,omitempty"`
	// Reason that the statistical result is marked as unreliable.
	UnreliableReason *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason `json:"unreliable_reason,omitempty"`
	// Metric value calculated for this variant.
	VariantMetricValue datadog.NullableFloat64 `json:"variant_metric_value,omitempty"`
	// Standardized statistic used to compare the observed effect with zero.
	ZScore *float64 `json:"z_score,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems instantiates a new ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems() *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems {
	this := ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems{}
	return &this
}

// NewExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsWithDefaults instantiates a new ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsWithDefaults() *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems {
	this := ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems{}
	return &this
}

// GetConfidenceInterval returns the ConfidenceInterval field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetConfidenceInterval() ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval {
	if o == nil || o.ConfidenceInterval == nil {
		var ret ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval
		return ret
	}
	return *o.ConfidenceInterval
}

// GetConfidenceIntervalOk returns a tuple with the ConfidenceInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetConfidenceIntervalOk() (*ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval, bool) {
	if o == nil || o.ConfidenceInterval == nil {
		return nil, false
	}
	return o.ConfidenceInterval, true
}

// HasConfidenceInterval returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasConfidenceInterval() bool {
	return o != nil && o.ConfidenceInterval != nil
}

// SetConfidenceInterval gets a reference to the given ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval and assigns it to the ConfidenceInterval field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetConfidenceInterval(v ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval) {
	o.ConfidenceInterval = &v
}

// GetConfidenceLevel returns the ConfidenceLevel field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetConfidenceLevel() float64 {
	if o == nil || o.ConfidenceLevel == nil {
		var ret float64
		return ret
	}
	return *o.ConfidenceLevel
}

// GetConfidenceLevelOk returns a tuple with the ConfidenceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetConfidenceLevelOk() (*float64, bool) {
	if o == nil || o.ConfidenceLevel == nil {
		return nil, false
	}
	return o.ConfidenceLevel, true
}

// HasConfidenceLevel returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasConfidenceLevel() bool {
	return o != nil && o.ConfidenceLevel != nil
}

// SetConfidenceLevel gets a reference to the given float64 and assigns it to the ConfidenceLevel field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetConfidenceLevel(v float64) {
	o.ConfidenceLevel = &v
}

// GetEvsi returns the Evsi field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetEvsi() float64 {
	if o == nil || o.Evsi == nil {
		var ret float64
		return ret
	}
	return *o.Evsi
}

// GetEvsiOk returns a tuple with the Evsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetEvsiOk() (*float64, bool) {
	if o == nil || o.Evsi == nil {
		return nil, false
	}
	return o.Evsi, true
}

// HasEvsi returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasEvsi() bool {
	return o != nil && o.Evsi != nil
}

// SetEvsi gets a reference to the given float64 and assigns it to the Evsi field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetEvsi(v float64) {
	o.Evsi = &v
}

// GetExpectationAboveZero returns the ExpectationAboveZero field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetExpectationAboveZero() float64 {
	if o == nil || o.ExpectationAboveZero == nil {
		var ret float64
		return ret
	}
	return *o.ExpectationAboveZero
}

// GetExpectationAboveZeroOk returns a tuple with the ExpectationAboveZero field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetExpectationAboveZeroOk() (*float64, bool) {
	if o == nil || o.ExpectationAboveZero == nil {
		return nil, false
	}
	return o.ExpectationAboveZero, true
}

// HasExpectationAboveZero returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasExpectationAboveZero() bool {
	return o != nil && o.ExpectationAboveZero != nil
}

// SetExpectationAboveZero gets a reference to the given float64 and assigns it to the ExpectationAboveZero field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetExpectationAboveZero(v float64) {
	o.ExpectationAboveZero = &v
}

// GetExpectationBelowZero returns the ExpectationBelowZero field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetExpectationBelowZero() float64 {
	if o == nil || o.ExpectationBelowZero == nil {
		var ret float64
		return ret
	}
	return *o.ExpectationBelowZero
}

// GetExpectationBelowZeroOk returns a tuple with the ExpectationBelowZero field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetExpectationBelowZeroOk() (*float64, bool) {
	if o == nil || o.ExpectationBelowZero == nil {
		return nil, false
	}
	return o.ExpectationBelowZero, true
}

// HasExpectationBelowZero returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasExpectationBelowZero() bool {
	return o != nil && o.ExpectationBelowZero != nil
}

// SetExpectationBelowZero gets a reference to the given float64 and assigns it to the ExpectationBelowZero field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetExpectationBelowZero(v float64) {
	o.ExpectationBelowZero = &v
}

// GetGlobalLift returns the GlobalLift field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetGlobalLift() float64 {
	if o == nil || o.GlobalLift == nil {
		var ret float64
		return ret
	}
	return *o.GlobalLift
}

// GetGlobalLiftOk returns a tuple with the GlobalLift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetGlobalLiftOk() (*float64, bool) {
	if o == nil || o.GlobalLift == nil {
		return nil, false
	}
	return o.GlobalLift, true
}

// HasGlobalLift returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasGlobalLift() bool {
	return o != nil && o.GlobalLift != nil
}

// SetGlobalLift gets a reference to the given float64 and assigns it to the GlobalLift field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetGlobalLift(v float64) {
	o.GlobalLift = &v
}

// GetGlobalLiftLowerBound returns the GlobalLiftLowerBound field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetGlobalLiftLowerBound() float64 {
	if o == nil || o.GlobalLiftLowerBound == nil {
		var ret float64
		return ret
	}
	return *o.GlobalLiftLowerBound
}

// GetGlobalLiftLowerBoundOk returns a tuple with the GlobalLiftLowerBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetGlobalLiftLowerBoundOk() (*float64, bool) {
	if o == nil || o.GlobalLiftLowerBound == nil {
		return nil, false
	}
	return o.GlobalLiftLowerBound, true
}

// HasGlobalLiftLowerBound returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasGlobalLiftLowerBound() bool {
	return o != nil && o.GlobalLiftLowerBound != nil
}

// SetGlobalLiftLowerBound gets a reference to the given float64 and assigns it to the GlobalLiftLowerBound field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetGlobalLiftLowerBound(v float64) {
	o.GlobalLiftLowerBound = &v
}

// GetGlobalLiftUpperBound returns the GlobalLiftUpperBound field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetGlobalLiftUpperBound() float64 {
	if o == nil || o.GlobalLiftUpperBound == nil {
		var ret float64
		return ret
	}
	return *o.GlobalLiftUpperBound
}

// GetGlobalLiftUpperBoundOk returns a tuple with the GlobalLiftUpperBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetGlobalLiftUpperBoundOk() (*float64, bool) {
	if o == nil || o.GlobalLiftUpperBound == nil {
		return nil, false
	}
	return o.GlobalLiftUpperBound, true
}

// HasGlobalLiftUpperBound returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasGlobalLiftUpperBound() bool {
	return o != nil && o.GlobalLiftUpperBound != nil
}

// SetGlobalLiftUpperBound gets a reference to the given float64 and assigns it to the GlobalLiftUpperBound field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetGlobalLiftUpperBound(v float64) {
	o.GlobalLiftUpperBound = &v
}

// GetIsCupedAdjusted returns the IsCupedAdjusted field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetIsCupedAdjusted() bool {
	if o == nil || o.IsCupedAdjusted == nil {
		var ret bool
		return ret
	}
	return *o.IsCupedAdjusted
}

// GetIsCupedAdjustedOk returns a tuple with the IsCupedAdjusted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetIsCupedAdjustedOk() (*bool, bool) {
	if o == nil || o.IsCupedAdjusted == nil {
		return nil, false
	}
	return o.IsCupedAdjusted, true
}

// HasIsCupedAdjusted returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasIsCupedAdjusted() bool {
	return o != nil && o.IsCupedAdjusted != nil
}

// SetIsCupedAdjusted gets a reference to the given bool and assigns it to the IsCupedAdjusted field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetIsCupedAdjusted(v bool) {
	o.IsCupedAdjusted = &v
}

// GetIsUnreliable returns the IsUnreliable field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetIsUnreliable() bool {
	if o == nil || o.IsUnreliable == nil {
		var ret bool
		return ret
	}
	return *o.IsUnreliable
}

// GetIsUnreliableOk returns a tuple with the IsUnreliable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetIsUnreliableOk() (*bool, bool) {
	if o == nil || o.IsUnreliable == nil {
		return nil, false
	}
	return o.IsUnreliable, true
}

// HasIsUnreliable returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasIsUnreliable() bool {
	return o != nil && o.IsUnreliable != nil
}

// SetIsUnreliable gets a reference to the given bool and assigns it to the IsUnreliable field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetIsUnreliable(v bool) {
	o.IsUnreliable = &v
}

// GetLiftType returns the LiftType field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetLiftType() ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType {
	if o == nil || o.LiftType == nil {
		var ret ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType
		return ret
	}
	return *o.LiftType
}

// GetLiftTypeOk returns a tuple with the LiftType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetLiftTypeOk() (*ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType, bool) {
	if o == nil || o.LiftType == nil {
		return nil, false
	}
	return o.LiftType, true
}

// HasLiftType returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasLiftType() bool {
	return o != nil && o.LiftType != nil
}

// SetLiftType gets a reference to the given ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType and assigns it to the LiftType field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetLiftType(v ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType) {
	o.LiftType = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetMethod() ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod {
	if o == nil || o.Method == nil {
		var ret ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetMethodOk() (*ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasMethod() bool {
	return o != nil && o.Method != nil
}

// SetMethod gets a reference to the given ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod and assigns it to the Method field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetMethod(v ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod) {
	o.Method = &v
}

// GetMinimumExpectedRegret returns the MinimumExpectedRegret field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetMinimumExpectedRegret() float64 {
	if o == nil || o.MinimumExpectedRegret == nil {
		var ret float64
		return ret
	}
	return *o.MinimumExpectedRegret
}

// GetMinimumExpectedRegretOk returns a tuple with the MinimumExpectedRegret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetMinimumExpectedRegretOk() (*float64, bool) {
	if o == nil || o.MinimumExpectedRegret == nil {
		return nil, false
	}
	return o.MinimumExpectedRegret, true
}

// HasMinimumExpectedRegret returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasMinimumExpectedRegret() bool {
	return o != nil && o.MinimumExpectedRegret != nil
}

// SetMinimumExpectedRegret gets a reference to the given float64 and assigns it to the MinimumExpectedRegret field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetMinimumExpectedRegret(v float64) {
	o.MinimumExpectedRegret = &v
}

// GetPValue returns the PValue field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetPValue() float64 {
	if o == nil || o.PValue == nil {
		var ret float64
		return ret
	}
	return *o.PValue
}

// GetPValueOk returns a tuple with the PValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetPValueOk() (*float64, bool) {
	if o == nil || o.PValue == nil {
		return nil, false
	}
	return o.PValue, true
}

// HasPValue returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasPValue() bool {
	return o != nil && o.PValue != nil
}

// SetPValue gets a reference to the given float64 and assigns it to the PValue field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetPValue(v float64) {
	o.PValue = &v
}

// GetPointEstimate returns the PointEstimate field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetPointEstimate() float64 {
	if o == nil || o.PointEstimate == nil {
		var ret float64
		return ret
	}
	return *o.PointEstimate
}

// GetPointEstimateOk returns a tuple with the PointEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetPointEstimateOk() (*float64, bool) {
	if o == nil || o.PointEstimate == nil {
		return nil, false
	}
	return o.PointEstimate, true
}

// HasPointEstimate returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasPointEstimate() bool {
	return o != nil && o.PointEstimate != nil
}

// SetPointEstimate gets a reference to the given float64 and assigns it to the PointEstimate field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetPointEstimate(v float64) {
	o.PointEstimate = &v
}

// GetProbabilityAboveZero returns the ProbabilityAboveZero field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetProbabilityAboveZero() float64 {
	if o == nil || o.ProbabilityAboveZero == nil {
		var ret float64
		return ret
	}
	return *o.ProbabilityAboveZero
}

// GetProbabilityAboveZeroOk returns a tuple with the ProbabilityAboveZero field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetProbabilityAboveZeroOk() (*float64, bool) {
	if o == nil || o.ProbabilityAboveZero == nil {
		return nil, false
	}
	return o.ProbabilityAboveZero, true
}

// HasProbabilityAboveZero returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasProbabilityAboveZero() bool {
	return o != nil && o.ProbabilityAboveZero != nil
}

// SetProbabilityAboveZero gets a reference to the given float64 and assigns it to the ProbabilityAboveZero field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetProbabilityAboveZero(v float64) {
	o.ProbabilityAboveZero = &v
}

// GetProbabilityBelowZero returns the ProbabilityBelowZero field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetProbabilityBelowZero() float64 {
	if o == nil || o.ProbabilityBelowZero == nil {
		var ret float64
		return ret
	}
	return *o.ProbabilityBelowZero
}

// GetProbabilityBelowZeroOk returns a tuple with the ProbabilityBelowZero field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetProbabilityBelowZeroOk() (*float64, bool) {
	if o == nil || o.ProbabilityBelowZero == nil {
		return nil, false
	}
	return o.ProbabilityBelowZero, true
}

// HasProbabilityBelowZero returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasProbabilityBelowZero() bool {
	return o != nil && o.ProbabilityBelowZero != nil
}

// SetProbabilityBelowZero gets a reference to the given float64 and assigns it to the ProbabilityBelowZero field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetProbabilityBelowZero(v float64) {
	o.ProbabilityBelowZero = &v
}

// GetStandardError returns the StandardError field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetStandardError() float64 {
	if o == nil || o.StandardError == nil {
		var ret float64
		return ret
	}
	return *o.StandardError
}

// GetStandardErrorOk returns a tuple with the StandardError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetStandardErrorOk() (*float64, bool) {
	if o == nil || o.StandardError == nil {
		return nil, false
	}
	return o.StandardError, true
}

// HasStandardError returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasStandardError() bool {
	return o != nil && o.StandardError != nil
}

// SetStandardError gets a reference to the given float64 and assigns it to the StandardError field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetStandardError(v float64) {
	o.StandardError = &v
}

// GetUnreliableReason returns the UnreliableReason field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetUnreliableReason() ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason {
	if o == nil || o.UnreliableReason == nil {
		var ret ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason
		return ret
	}
	return *o.UnreliableReason
}

// GetUnreliableReasonOk returns a tuple with the UnreliableReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetUnreliableReasonOk() (*ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason, bool) {
	if o == nil || o.UnreliableReason == nil {
		return nil, false
	}
	return o.UnreliableReason, true
}

// HasUnreliableReason returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasUnreliableReason() bool {
	return o != nil && o.UnreliableReason != nil
}

// SetUnreliableReason gets a reference to the given ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason and assigns it to the UnreliableReason field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetUnreliableReason(v ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason) {
	o.UnreliableReason = &v
}

// GetVariantMetricValue returns the VariantMetricValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetVariantMetricValue() float64 {
	if o == nil || o.VariantMetricValue.Get() == nil {
		var ret float64
		return ret
	}
	return *o.VariantMetricValue.Get()
}

// GetVariantMetricValueOk returns a tuple with the VariantMetricValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetVariantMetricValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.VariantMetricValue.Get(), o.VariantMetricValue.IsSet()
}

// HasVariantMetricValue returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasVariantMetricValue() bool {
	return o != nil && o.VariantMetricValue.IsSet()
}

// SetVariantMetricValue gets a reference to the given datadog.NullableFloat64 and assigns it to the VariantMetricValue field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetVariantMetricValue(v float64) {
	o.VariantMetricValue.Set(&v)
}

// SetVariantMetricValueNil sets the value for VariantMetricValue to be an explicit nil.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetVariantMetricValueNil() {
	o.VariantMetricValue.Set(nil)
}

// UnsetVariantMetricValue ensures that no value is present for VariantMetricValue, not even an explicit nil.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) UnsetVariantMetricValue() {
	o.VariantMetricValue.Unset()
}

// GetZScore returns the ZScore field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetZScore() float64 {
	if o == nil || o.ZScore == nil {
		var ret float64
		return ret
	}
	return *o.ZScore
}

// GetZScoreOk returns a tuple with the ZScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) GetZScoreOk() (*float64, bool) {
	if o == nil || o.ZScore == nil {
		return nil, false
	}
	return o.ZScore, true
}

// HasZScore returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) HasZScore() bool {
	return o != nil && o.ZScore != nil
}

// SetZScore gets a reference to the given float64 and assigns it to the ZScore field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) SetZScore(v float64) {
	o.ZScore = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConfidenceInterval != nil {
		toSerialize["confidence_interval"] = o.ConfidenceInterval
	}
	if o.ConfidenceLevel != nil {
		toSerialize["confidence_level"] = o.ConfidenceLevel
	}
	if o.Evsi != nil {
		toSerialize["evsi"] = o.Evsi
	}
	if o.ExpectationAboveZero != nil {
		toSerialize["expectation_above_zero"] = o.ExpectationAboveZero
	}
	if o.ExpectationBelowZero != nil {
		toSerialize["expectation_below_zero"] = o.ExpectationBelowZero
	}
	if o.GlobalLift != nil {
		toSerialize["global_lift"] = o.GlobalLift
	}
	if o.GlobalLiftLowerBound != nil {
		toSerialize["global_lift_lower_bound"] = o.GlobalLiftLowerBound
	}
	if o.GlobalLiftUpperBound != nil {
		toSerialize["global_lift_upper_bound"] = o.GlobalLiftUpperBound
	}
	if o.IsCupedAdjusted != nil {
		toSerialize["is_cuped_adjusted"] = o.IsCupedAdjusted
	}
	if o.IsUnreliable != nil {
		toSerialize["is_unreliable"] = o.IsUnreliable
	}
	if o.LiftType != nil {
		toSerialize["lift_type"] = o.LiftType
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.MinimumExpectedRegret != nil {
		toSerialize["minimum_expected_regret"] = o.MinimumExpectedRegret
	}
	if o.PValue != nil {
		toSerialize["p_value"] = o.PValue
	}
	if o.PointEstimate != nil {
		toSerialize["point_estimate"] = o.PointEstimate
	}
	if o.ProbabilityAboveZero != nil {
		toSerialize["probability_above_zero"] = o.ProbabilityAboveZero
	}
	if o.ProbabilityBelowZero != nil {
		toSerialize["probability_below_zero"] = o.ProbabilityBelowZero
	}
	if o.StandardError != nil {
		toSerialize["standard_error"] = o.StandardError
	}
	if o.UnreliableReason != nil {
		toSerialize["unreliable_reason"] = o.UnreliableReason
	}
	if o.VariantMetricValue.IsSet() {
		toSerialize["variant_metric_value"] = o.VariantMetricValue.Get()
	}
	if o.ZScore != nil {
		toSerialize["z_score"] = o.ZScore
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfidenceInterval    *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval `json:"confidence_interval,omitempty"`
		ConfidenceLevel       *float64                                                                                 `json:"confidence_level,omitempty"`
		Evsi                  *float64                                                                                 `json:"evsi,omitempty"`
		ExpectationAboveZero  *float64                                                                                 `json:"expectation_above_zero,omitempty"`
		ExpectationBelowZero  *float64                                                                                 `json:"expectation_below_zero,omitempty"`
		GlobalLift            *float64                                                                                 `json:"global_lift,omitempty"`
		GlobalLiftLowerBound  *float64                                                                                 `json:"global_lift_lower_bound,omitempty"`
		GlobalLiftUpperBound  *float64                                                                                 `json:"global_lift_upper_bound,omitempty"`
		IsCupedAdjusted       *bool                                                                                    `json:"is_cuped_adjusted,omitempty"`
		IsUnreliable          *bool                                                                                    `json:"is_unreliable,omitempty"`
		LiftType              *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType           `json:"lift_type,omitempty"`
		Method                *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod             `json:"method,omitempty"`
		MinimumExpectedRegret *float64                                                                                 `json:"minimum_expected_regret,omitempty"`
		PValue                *float64                                                                                 `json:"p_value,omitempty"`
		PointEstimate         *float64                                                                                 `json:"point_estimate,omitempty"`
		ProbabilityAboveZero  *float64                                                                                 `json:"probability_above_zero,omitempty"`
		ProbabilityBelowZero  *float64                                                                                 `json:"probability_below_zero,omitempty"`
		StandardError         *float64                                                                                 `json:"standard_error,omitempty"`
		UnreliableReason      *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason   `json:"unreliable_reason,omitempty"`
		VariantMetricValue    datadog.NullableFloat64                                                                  `json:"variant_metric_value,omitempty"`
		ZScore                *float64                                                                                 `json:"z_score,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"confidence_interval", "confidence_level", "evsi", "expectation_above_zero", "expectation_below_zero", "global_lift", "global_lift_lower_bound", "global_lift_upper_bound", "is_cuped_adjusted", "is_unreliable", "lift_type", "method", "minimum_expected_regret", "p_value", "point_estimate", "probability_above_zero", "probability_below_zero", "standard_error", "unreliable_reason", "variant_metric_value", "z_score"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ConfidenceInterval != nil && all.ConfidenceInterval.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ConfidenceInterval = all.ConfidenceInterval
	o.ConfidenceLevel = all.ConfidenceLevel
	o.Evsi = all.Evsi
	o.ExpectationAboveZero = all.ExpectationAboveZero
	o.ExpectationBelowZero = all.ExpectationBelowZero
	o.GlobalLift = all.GlobalLift
	o.GlobalLiftLowerBound = all.GlobalLiftLowerBound
	o.GlobalLiftUpperBound = all.GlobalLiftUpperBound
	o.IsCupedAdjusted = all.IsCupedAdjusted
	o.IsUnreliable = all.IsUnreliable
	if all.LiftType != nil && !all.LiftType.IsValid() {
		hasInvalidField = true
	} else {
		o.LiftType = all.LiftType
	}
	if all.Method != nil && !all.Method.IsValid() {
		hasInvalidField = true
	} else {
		o.Method = all.Method
	}
	o.MinimumExpectedRegret = all.MinimumExpectedRegret
	o.PValue = all.PValue
	o.PointEstimate = all.PointEstimate
	o.ProbabilityAboveZero = all.ProbabilityAboveZero
	o.ProbabilityBelowZero = all.ProbabilityBelowZero
	o.StandardError = all.StandardError
	if all.UnreliableReason != nil && !all.UnreliableReason.IsValid() {
		hasInvalidField = true
	} else {
		o.UnreliableReason = all.UnreliableReason
	}
	o.VariantMetricValue = all.VariantMetricValue
	o.ZScore = all.ZScore

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
