// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditVariantResultV2 An arm's estimates from the latest saved refresh.
type ExperimentsBanditVariantResultV2 struct {
	// A central interval around point_estimate at confidence_level. Null when the estimate is unavailable.
	ConfidenceInterval NullableExperimentsBanditNullableConfidenceIntervalV2 `json:"confidence_interval"`
	// Where the arm's reward variance came from, or why it's unavailable.
	DataStatus *ExperimentsBanditArmDataStatus `json:"data_status,omitempty"`
	// The arm's variant ID on the flag.
	FeatureFlagVariantId uuid.UUID `json:"feature_flag_variant_id"`
	// The arm's valid observations whose aggregation window has closed. The estimate uses these.
	MatureObservations int64 `json:"mature_observations"`
	// The arm's estimated mean of the metric, in the metric's units. It isn't a lift.
	PointEstimate datadog.NullableFloat64 `json:"point_estimate"`
	// Why the arm's estimate and interval are unavailable. Absent when they are available.
	PosteriorUnavailableReason *ExperimentsBanditPosteriorUnavailableReason `json:"posterior_unavailable_reason,omitempty"`
	// The probability that this arm is the best.
	ProbabilityOfBeingOptimal datadog.NullableFloat64 `json:"probability_of_being_optimal"`
	// The arm's variant key.
	VariantKey datadog.NullableString `json:"variant_key"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditVariantResultV2 instantiates a new ExperimentsBanditVariantResultV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditVariantResultV2(confidenceInterval NullableExperimentsBanditNullableConfidenceIntervalV2, featureFlagVariantId uuid.UUID, matureObservations int64, pointEstimate datadog.NullableFloat64, probabilityOfBeingOptimal datadog.NullableFloat64, variantKey datadog.NullableString) *ExperimentsBanditVariantResultV2 {
	this := ExperimentsBanditVariantResultV2{}
	this.ConfidenceInterval = confidenceInterval
	this.FeatureFlagVariantId = featureFlagVariantId
	this.MatureObservations = matureObservations
	this.PointEstimate = pointEstimate
	this.ProbabilityOfBeingOptimal = probabilityOfBeingOptimal
	this.VariantKey = variantKey
	return &this
}

// NewExperimentsBanditVariantResultV2WithDefaults instantiates a new ExperimentsBanditVariantResultV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditVariantResultV2WithDefaults() *ExperimentsBanditVariantResultV2 {
	this := ExperimentsBanditVariantResultV2{}
	return &this
}

// GetConfidenceInterval returns the ConfidenceInterval field value.
// If the value is explicit nil, the zero value for ExperimentsBanditNullableConfidenceIntervalV2 will be returned.
func (o *ExperimentsBanditVariantResultV2) GetConfidenceInterval() ExperimentsBanditNullableConfidenceIntervalV2 {
	if o == nil || o.ConfidenceInterval.Get() == nil {
		var ret ExperimentsBanditNullableConfidenceIntervalV2
		return ret
	}
	return *o.ConfidenceInterval.Get()
}

// GetConfidenceIntervalOk returns a tuple with the ConfidenceInterval field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditVariantResultV2) GetConfidenceIntervalOk() (*ExperimentsBanditNullableConfidenceIntervalV2, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfidenceInterval.Get(), o.ConfidenceInterval.IsSet()
}

// SetConfidenceInterval sets field value.
func (o *ExperimentsBanditVariantResultV2) SetConfidenceInterval(v ExperimentsBanditNullableConfidenceIntervalV2) {
	o.ConfidenceInterval.Set(&v)
}

// GetDataStatus returns the DataStatus field value if set, zero value otherwise.
func (o *ExperimentsBanditVariantResultV2) GetDataStatus() ExperimentsBanditArmDataStatus {
	if o == nil || o.DataStatus == nil {
		var ret ExperimentsBanditArmDataStatus
		return ret
	}
	return *o.DataStatus
}

// GetDataStatusOk returns a tuple with the DataStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditVariantResultV2) GetDataStatusOk() (*ExperimentsBanditArmDataStatus, bool) {
	if o == nil || o.DataStatus == nil {
		return nil, false
	}
	return o.DataStatus, true
}

// HasDataStatus returns a boolean if a field has been set.
func (o *ExperimentsBanditVariantResultV2) HasDataStatus() bool {
	return o != nil && o.DataStatus != nil
}

// SetDataStatus gets a reference to the given ExperimentsBanditArmDataStatus and assigns it to the DataStatus field.
func (o *ExperimentsBanditVariantResultV2) SetDataStatus(v ExperimentsBanditArmDataStatus) {
	o.DataStatus = &v
}

// GetFeatureFlagVariantId returns the FeatureFlagVariantId field value.
func (o *ExperimentsBanditVariantResultV2) GetFeatureFlagVariantId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.FeatureFlagVariantId
}

// GetFeatureFlagVariantIdOk returns a tuple with the FeatureFlagVariantId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditVariantResultV2) GetFeatureFlagVariantIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureFlagVariantId, true
}

// SetFeatureFlagVariantId sets field value.
func (o *ExperimentsBanditVariantResultV2) SetFeatureFlagVariantId(v uuid.UUID) {
	o.FeatureFlagVariantId = v
}

// GetMatureObservations returns the MatureObservations field value.
func (o *ExperimentsBanditVariantResultV2) GetMatureObservations() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MatureObservations
}

// GetMatureObservationsOk returns a tuple with the MatureObservations field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditVariantResultV2) GetMatureObservationsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatureObservations, true
}

// SetMatureObservations sets field value.
func (o *ExperimentsBanditVariantResultV2) SetMatureObservations(v int64) {
	o.MatureObservations = v
}

// GetPointEstimate returns the PointEstimate field value.
// If the value is explicit nil, the zero value for float64 will be returned.
func (o *ExperimentsBanditVariantResultV2) GetPointEstimate() float64 {
	if o == nil || o.PointEstimate.Get() == nil {
		var ret float64
		return ret
	}
	return *o.PointEstimate.Get()
}

// GetPointEstimateOk returns a tuple with the PointEstimate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditVariantResultV2) GetPointEstimateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PointEstimate.Get(), o.PointEstimate.IsSet()
}

// SetPointEstimate sets field value.
func (o *ExperimentsBanditVariantResultV2) SetPointEstimate(v float64) {
	o.PointEstimate.Set(&v)
}

// GetPosteriorUnavailableReason returns the PosteriorUnavailableReason field value if set, zero value otherwise.
func (o *ExperimentsBanditVariantResultV2) GetPosteriorUnavailableReason() ExperimentsBanditPosteriorUnavailableReason {
	if o == nil || o.PosteriorUnavailableReason == nil {
		var ret ExperimentsBanditPosteriorUnavailableReason
		return ret
	}
	return *o.PosteriorUnavailableReason
}

// GetPosteriorUnavailableReasonOk returns a tuple with the PosteriorUnavailableReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditVariantResultV2) GetPosteriorUnavailableReasonOk() (*ExperimentsBanditPosteriorUnavailableReason, bool) {
	if o == nil || o.PosteriorUnavailableReason == nil {
		return nil, false
	}
	return o.PosteriorUnavailableReason, true
}

// HasPosteriorUnavailableReason returns a boolean if a field has been set.
func (o *ExperimentsBanditVariantResultV2) HasPosteriorUnavailableReason() bool {
	return o != nil && o.PosteriorUnavailableReason != nil
}

// SetPosteriorUnavailableReason gets a reference to the given ExperimentsBanditPosteriorUnavailableReason and assigns it to the PosteriorUnavailableReason field.
func (o *ExperimentsBanditVariantResultV2) SetPosteriorUnavailableReason(v ExperimentsBanditPosteriorUnavailableReason) {
	o.PosteriorUnavailableReason = &v
}

// GetProbabilityOfBeingOptimal returns the ProbabilityOfBeingOptimal field value.
// If the value is explicit nil, the zero value for float64 will be returned.
func (o *ExperimentsBanditVariantResultV2) GetProbabilityOfBeingOptimal() float64 {
	if o == nil || o.ProbabilityOfBeingOptimal.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ProbabilityOfBeingOptimal.Get()
}

// GetProbabilityOfBeingOptimalOk returns a tuple with the ProbabilityOfBeingOptimal field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditVariantResultV2) GetProbabilityOfBeingOptimalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProbabilityOfBeingOptimal.Get(), o.ProbabilityOfBeingOptimal.IsSet()
}

// SetProbabilityOfBeingOptimal sets field value.
func (o *ExperimentsBanditVariantResultV2) SetProbabilityOfBeingOptimal(v float64) {
	o.ProbabilityOfBeingOptimal.Set(&v)
}

// GetVariantKey returns the VariantKey field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *ExperimentsBanditVariantResultV2) GetVariantKey() string {
	if o == nil || o.VariantKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.VariantKey.Get()
}

// GetVariantKeyOk returns a tuple with the VariantKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditVariantResultV2) GetVariantKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VariantKey.Get(), o.VariantKey.IsSet()
}

// SetVariantKey sets field value.
func (o *ExperimentsBanditVariantResultV2) SetVariantKey(v string) {
	o.VariantKey.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditVariantResultV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["confidence_interval"] = o.ConfidenceInterval.Get()
	if o.DataStatus != nil {
		toSerialize["data_status"] = o.DataStatus
	}
	toSerialize["feature_flag_variant_id"] = o.FeatureFlagVariantId
	toSerialize["mature_observations"] = o.MatureObservations
	toSerialize["point_estimate"] = o.PointEstimate.Get()
	if o.PosteriorUnavailableReason != nil {
		toSerialize["posterior_unavailable_reason"] = o.PosteriorUnavailableReason
	}
	toSerialize["probability_of_being_optimal"] = o.ProbabilityOfBeingOptimal.Get()
	toSerialize["variant_key"] = o.VariantKey.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditVariantResultV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfidenceInterval         NullableExperimentsBanditNullableConfidenceIntervalV2 `json:"confidence_interval"`
		DataStatus                 *ExperimentsBanditArmDataStatus                       `json:"data_status,omitempty"`
		FeatureFlagVariantId       *uuid.UUID                                            `json:"feature_flag_variant_id"`
		MatureObservations         *int64                                                `json:"mature_observations"`
		PointEstimate              datadog.NullableFloat64                               `json:"point_estimate"`
		PosteriorUnavailableReason *ExperimentsBanditPosteriorUnavailableReason          `json:"posterior_unavailable_reason,omitempty"`
		ProbabilityOfBeingOptimal  datadog.NullableFloat64                               `json:"probability_of_being_optimal"`
		VariantKey                 datadog.NullableString                                `json:"variant_key"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.ConfidenceInterval.IsSet() {
		return fmt.Errorf("required field confidence_interval missing")
	}
	if all.FeatureFlagVariantId == nil {
		return fmt.Errorf("required field feature_flag_variant_id missing")
	}
	if all.MatureObservations == nil {
		return fmt.Errorf("required field mature_observations missing")
	}
	if !all.PointEstimate.IsSet() {
		return fmt.Errorf("required field point_estimate missing")
	}
	if !all.ProbabilityOfBeingOptimal.IsSet() {
		return fmt.Errorf("required field probability_of_being_optimal missing")
	}
	if !all.VariantKey.IsSet() {
		return fmt.Errorf("required field variant_key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"confidence_interval", "data_status", "feature_flag_variant_id", "mature_observations", "point_estimate", "posterior_unavailable_reason", "probability_of_being_optimal", "variant_key"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ConfidenceInterval = all.ConfidenceInterval
	if all.DataStatus != nil && !all.DataStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.DataStatus = all.DataStatus
	}
	o.FeatureFlagVariantId = *all.FeatureFlagVariantId
	o.MatureObservations = *all.MatureObservations
	o.PointEstimate = all.PointEstimate
	if all.PosteriorUnavailableReason != nil && !all.PosteriorUnavailableReason.IsValid() {
		hasInvalidField = true
	} else {
		o.PosteriorUnavailableReason = all.PosteriorUnavailableReason
	}
	o.ProbabilityOfBeingOptimal = all.ProbabilityOfBeingOptimal
	o.VariantKey = all.VariantKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
