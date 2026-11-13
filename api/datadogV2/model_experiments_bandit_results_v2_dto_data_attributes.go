// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditResultsV2DTODataAttributes What the bandit's refreshes have learned.
type ExperimentsBanditResultsV2DTODataAttributes struct {
	// The probability mass of each confidence interval.
	ConfidenceLevel float64 `json:"confidence_level"`
	// The configuration version the latest saved refresh used. It can be older than the bandit's, because changes such as a new end date don't invalidate results.
	ConfigVersion datadog.NullableInt64 `json:"config_version"`
	// The latest warehouse data the latest saved refresh read.
	DataAsOf datadog.NullableTime `json:"data_as_of"`
	// The direction the bandit optimizes the metric in. It's fixed when the bandit starts. Null before the first saved refresh.
	DesiredChange NullableExperimentsBanditDesiredChange `json:"desired_change"`
	// An estimate of when the next scheduled refresh starts. Null when the bandit isn't running or the time can't be estimated. Only daily and hourly schedules get an estimate.
	NextRefreshAt datadog.NullableTime `json:"next_refresh_at"`
	// The key of the arm most likely to be best, when one arm leads and the bandit is outside warmup. Null otherwise.
	OptimalVariantKey datadog.NullableString `json:"optimal_variant_key"`
	// Why no optimal variant is reported. Absent when one is, or when the cause isn't known.
	OptimalVariantUnavailableReason *ExperimentsBanditOptimalVariantUnavailableReason `json:"optimal_variant_unavailable_reason,omitempty"`
	// Each arm's point estimate from up to the latest 400 saved refreshes in the requested range, oldest first.
	PointEstimates []ExperimentsBanditPointEstimateV2 `json:"point_estimates"`
	// The latest refresh that saved results. The estimates come from it.
	RefreshId datadog.NullableUUID `json:"refresh_id"`
	// When the latest saved refresh saved its results.
	ResultsLastUpdated datadog.NullableTime `json:"results_last_updated"`
	// WAITING_FOR_FIRST_REFRESH until a refresh saves results. STATS_UNAVAILABLE when the latest saved refresh has no complete probabilities of being optimal. ENDED once the bandit is stopped or completed; its last estimates stay. Otherwise WARMUP or OPTIMIZING says which policy the bandit serves. A paused bandit keeps its state.
	State ExperimentsBanditResultsState `json:"state"`
	// Whether the latest saved refresh calculated its statistics. Absent for older refreshes.
	StatisticsStatus *ExperimentsBanditStatisticsStatus `json:"statistics_status,omitempty"`
	// One row per arm, in arm order.
	Variants []ExperimentsBanditVariantResultV2 `json:"variants"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditResultsV2DTODataAttributes instantiates a new ExperimentsBanditResultsV2DTODataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditResultsV2DTODataAttributes(confidenceLevel float64, configVersion datadog.NullableInt64, dataAsOf datadog.NullableTime, desiredChange NullableExperimentsBanditDesiredChange, nextRefreshAt datadog.NullableTime, optimalVariantKey datadog.NullableString, pointEstimates []ExperimentsBanditPointEstimateV2, refreshId datadog.NullableUUID, resultsLastUpdated datadog.NullableTime, state ExperimentsBanditResultsState, variants []ExperimentsBanditVariantResultV2) *ExperimentsBanditResultsV2DTODataAttributes {
	this := ExperimentsBanditResultsV2DTODataAttributes{}
	this.ConfidenceLevel = confidenceLevel
	this.ConfigVersion = configVersion
	this.DataAsOf = dataAsOf
	this.DesiredChange = desiredChange
	this.NextRefreshAt = nextRefreshAt
	this.OptimalVariantKey = optimalVariantKey
	this.PointEstimates = pointEstimates
	this.RefreshId = refreshId
	this.ResultsLastUpdated = resultsLastUpdated
	this.State = state
	this.Variants = variants
	return &this
}

// NewExperimentsBanditResultsV2DTODataAttributesWithDefaults instantiates a new ExperimentsBanditResultsV2DTODataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditResultsV2DTODataAttributesWithDefaults() *ExperimentsBanditResultsV2DTODataAttributes {
	this := ExperimentsBanditResultsV2DTODataAttributes{}
	return &this
}

// GetConfidenceLevel returns the ConfidenceLevel field value.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetConfidenceLevel() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.ConfidenceLevel
}

// GetConfidenceLevelOk returns a tuple with the ConfidenceLevel field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetConfidenceLevelOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfidenceLevel, true
}

// SetConfidenceLevel sets field value.
func (o *ExperimentsBanditResultsV2DTODataAttributes) SetConfidenceLevel(v float64) {
	o.ConfidenceLevel = v
}

// GetConfigVersion returns the ConfigVersion field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetConfigVersion() int64 {
	if o == nil || o.ConfigVersion.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ConfigVersion.Get()
}

// GetConfigVersionOk returns a tuple with the ConfigVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetConfigVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigVersion.Get(), o.ConfigVersion.IsSet()
}

// SetConfigVersion sets field value.
func (o *ExperimentsBanditResultsV2DTODataAttributes) SetConfigVersion(v int64) {
	o.ConfigVersion.Set(&v)
}

// GetDataAsOf returns the DataAsOf field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetDataAsOf() time.Time {
	if o == nil || o.DataAsOf.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DataAsOf.Get()
}

// GetDataAsOfOk returns a tuple with the DataAsOf field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetDataAsOfOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataAsOf.Get(), o.DataAsOf.IsSet()
}

// SetDataAsOf sets field value.
func (o *ExperimentsBanditResultsV2DTODataAttributes) SetDataAsOf(v time.Time) {
	o.DataAsOf.Set(&v)
}

// GetDesiredChange returns the DesiredChange field value.
// If the value is explicit nil, the zero value for ExperimentsBanditDesiredChange will be returned.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetDesiredChange() ExperimentsBanditDesiredChange {
	if o == nil || o.DesiredChange.Get() == nil {
		var ret ExperimentsBanditDesiredChange
		return ret
	}
	return *o.DesiredChange.Get()
}

// GetDesiredChangeOk returns a tuple with the DesiredChange field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetDesiredChangeOk() (*ExperimentsBanditDesiredChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.DesiredChange.Get(), o.DesiredChange.IsSet()
}

// SetDesiredChange sets field value.
func (o *ExperimentsBanditResultsV2DTODataAttributes) SetDesiredChange(v ExperimentsBanditDesiredChange) {
	o.DesiredChange.Set(&v)
}

// GetNextRefreshAt returns the NextRefreshAt field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetNextRefreshAt() time.Time {
	if o == nil || o.NextRefreshAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.NextRefreshAt.Get()
}

// GetNextRefreshAtOk returns a tuple with the NextRefreshAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetNextRefreshAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextRefreshAt.Get(), o.NextRefreshAt.IsSet()
}

// SetNextRefreshAt sets field value.
func (o *ExperimentsBanditResultsV2DTODataAttributes) SetNextRefreshAt(v time.Time) {
	o.NextRefreshAt.Set(&v)
}

// GetOptimalVariantKey returns the OptimalVariantKey field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetOptimalVariantKey() string {
	if o == nil || o.OptimalVariantKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.OptimalVariantKey.Get()
}

// GetOptimalVariantKeyOk returns a tuple with the OptimalVariantKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetOptimalVariantKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptimalVariantKey.Get(), o.OptimalVariantKey.IsSet()
}

// SetOptimalVariantKey sets field value.
func (o *ExperimentsBanditResultsV2DTODataAttributes) SetOptimalVariantKey(v string) {
	o.OptimalVariantKey.Set(&v)
}

// GetOptimalVariantUnavailableReason returns the OptimalVariantUnavailableReason field value if set, zero value otherwise.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetOptimalVariantUnavailableReason() ExperimentsBanditOptimalVariantUnavailableReason {
	if o == nil || o.OptimalVariantUnavailableReason == nil {
		var ret ExperimentsBanditOptimalVariantUnavailableReason
		return ret
	}
	return *o.OptimalVariantUnavailableReason
}

// GetOptimalVariantUnavailableReasonOk returns a tuple with the OptimalVariantUnavailableReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetOptimalVariantUnavailableReasonOk() (*ExperimentsBanditOptimalVariantUnavailableReason, bool) {
	if o == nil || o.OptimalVariantUnavailableReason == nil {
		return nil, false
	}
	return o.OptimalVariantUnavailableReason, true
}

// HasOptimalVariantUnavailableReason returns a boolean if a field has been set.
func (o *ExperimentsBanditResultsV2DTODataAttributes) HasOptimalVariantUnavailableReason() bool {
	return o != nil && o.OptimalVariantUnavailableReason != nil
}

// SetOptimalVariantUnavailableReason gets a reference to the given ExperimentsBanditOptimalVariantUnavailableReason and assigns it to the OptimalVariantUnavailableReason field.
func (o *ExperimentsBanditResultsV2DTODataAttributes) SetOptimalVariantUnavailableReason(v ExperimentsBanditOptimalVariantUnavailableReason) {
	o.OptimalVariantUnavailableReason = &v
}

// GetPointEstimates returns the PointEstimates field value.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetPointEstimates() []ExperimentsBanditPointEstimateV2 {
	if o == nil {
		var ret []ExperimentsBanditPointEstimateV2
		return ret
	}
	return o.PointEstimates
}

// GetPointEstimatesOk returns a tuple with the PointEstimates field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetPointEstimatesOk() (*[]ExperimentsBanditPointEstimateV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PointEstimates, true
}

// SetPointEstimates sets field value.
func (o *ExperimentsBanditResultsV2DTODataAttributes) SetPointEstimates(v []ExperimentsBanditPointEstimateV2) {
	o.PointEstimates = v
}

// GetRefreshId returns the RefreshId field value.
// If the value is explicit nil, the zero value for uuid.UUID will be returned.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetRefreshId() uuid.UUID {
	if o == nil || o.RefreshId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.RefreshId.Get()
}

// GetRefreshIdOk returns a tuple with the RefreshId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetRefreshIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshId.Get(), o.RefreshId.IsSet()
}

// SetRefreshId sets field value.
func (o *ExperimentsBanditResultsV2DTODataAttributes) SetRefreshId(v uuid.UUID) {
	o.RefreshId.Set(&v)
}

// GetResultsLastUpdated returns the ResultsLastUpdated field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetResultsLastUpdated() time.Time {
	if o == nil || o.ResultsLastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ResultsLastUpdated.Get()
}

// GetResultsLastUpdatedOk returns a tuple with the ResultsLastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetResultsLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResultsLastUpdated.Get(), o.ResultsLastUpdated.IsSet()
}

// SetResultsLastUpdated sets field value.
func (o *ExperimentsBanditResultsV2DTODataAttributes) SetResultsLastUpdated(v time.Time) {
	o.ResultsLastUpdated.Set(&v)
}

// GetState returns the State field value.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetState() ExperimentsBanditResultsState {
	if o == nil {
		var ret ExperimentsBanditResultsState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetStateOk() (*ExperimentsBanditResultsState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *ExperimentsBanditResultsV2DTODataAttributes) SetState(v ExperimentsBanditResultsState) {
	o.State = v
}

// GetStatisticsStatus returns the StatisticsStatus field value if set, zero value otherwise.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetStatisticsStatus() ExperimentsBanditStatisticsStatus {
	if o == nil || o.StatisticsStatus == nil {
		var ret ExperimentsBanditStatisticsStatus
		return ret
	}
	return *o.StatisticsStatus
}

// GetStatisticsStatusOk returns a tuple with the StatisticsStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetStatisticsStatusOk() (*ExperimentsBanditStatisticsStatus, bool) {
	if o == nil || o.StatisticsStatus == nil {
		return nil, false
	}
	return o.StatisticsStatus, true
}

// HasStatisticsStatus returns a boolean if a field has been set.
func (o *ExperimentsBanditResultsV2DTODataAttributes) HasStatisticsStatus() bool {
	return o != nil && o.StatisticsStatus != nil
}

// SetStatisticsStatus gets a reference to the given ExperimentsBanditStatisticsStatus and assigns it to the StatisticsStatus field.
func (o *ExperimentsBanditResultsV2DTODataAttributes) SetStatisticsStatus(v ExperimentsBanditStatisticsStatus) {
	o.StatisticsStatus = &v
}

// GetVariants returns the Variants field value.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetVariants() []ExperimentsBanditVariantResultV2 {
	if o == nil {
		var ret []ExperimentsBanditVariantResultV2
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditResultsV2DTODataAttributes) GetVariantsOk() (*[]ExperimentsBanditVariantResultV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variants, true
}

// SetVariants sets field value.
func (o *ExperimentsBanditResultsV2DTODataAttributes) SetVariants(v []ExperimentsBanditVariantResultV2) {
	o.Variants = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditResultsV2DTODataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["confidence_level"] = o.ConfidenceLevel
	toSerialize["config_version"] = o.ConfigVersion.Get()
	toSerialize["data_as_of"] = o.DataAsOf.Get()
	toSerialize["desired_change"] = o.DesiredChange.Get()
	toSerialize["next_refresh_at"] = o.NextRefreshAt.Get()
	toSerialize["optimal_variant_key"] = o.OptimalVariantKey.Get()
	if o.OptimalVariantUnavailableReason != nil {
		toSerialize["optimal_variant_unavailable_reason"] = o.OptimalVariantUnavailableReason
	}
	toSerialize["point_estimates"] = o.PointEstimates
	toSerialize["refresh_id"] = o.RefreshId.Get()
	toSerialize["results_last_updated"] = o.ResultsLastUpdated.Get()
	toSerialize["state"] = o.State
	if o.StatisticsStatus != nil {
		toSerialize["statistics_status"] = o.StatisticsStatus
	}
	toSerialize["variants"] = o.Variants

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditResultsV2DTODataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfidenceLevel                 *float64                                          `json:"confidence_level"`
		ConfigVersion                   datadog.NullableInt64                             `json:"config_version"`
		DataAsOf                        datadog.NullableTime                              `json:"data_as_of"`
		DesiredChange                   NullableExperimentsBanditDesiredChange            `json:"desired_change"`
		NextRefreshAt                   datadog.NullableTime                              `json:"next_refresh_at"`
		OptimalVariantKey               datadog.NullableString                            `json:"optimal_variant_key"`
		OptimalVariantUnavailableReason *ExperimentsBanditOptimalVariantUnavailableReason `json:"optimal_variant_unavailable_reason,omitempty"`
		PointEstimates                  *[]ExperimentsBanditPointEstimateV2               `json:"point_estimates"`
		RefreshId                       datadog.NullableUUID                              `json:"refresh_id"`
		ResultsLastUpdated              datadog.NullableTime                              `json:"results_last_updated"`
		State                           *ExperimentsBanditResultsState                    `json:"state"`
		StatisticsStatus                *ExperimentsBanditStatisticsStatus                `json:"statistics_status,omitempty"`
		Variants                        *[]ExperimentsBanditVariantResultV2               `json:"variants"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConfidenceLevel == nil {
		return fmt.Errorf("required field confidence_level missing")
	}
	if !all.ConfigVersion.IsSet() {
		return fmt.Errorf("required field config_version missing")
	}
	if !all.DataAsOf.IsSet() {
		return fmt.Errorf("required field data_as_of missing")
	}
	if !all.DesiredChange.IsSet() {
		return fmt.Errorf("required field desired_change missing")
	}
	if !all.NextRefreshAt.IsSet() {
		return fmt.Errorf("required field next_refresh_at missing")
	}
	if !all.OptimalVariantKey.IsSet() {
		return fmt.Errorf("required field optimal_variant_key missing")
	}
	if all.PointEstimates == nil {
		return fmt.Errorf("required field point_estimates missing")
	}
	if !all.RefreshId.IsSet() {
		return fmt.Errorf("required field refresh_id missing")
	}
	if !all.ResultsLastUpdated.IsSet() {
		return fmt.Errorf("required field results_last_updated missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	if all.Variants == nil {
		return fmt.Errorf("required field variants missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"confidence_level", "config_version", "data_as_of", "desired_change", "next_refresh_at", "optimal_variant_key", "optimal_variant_unavailable_reason", "point_estimates", "refresh_id", "results_last_updated", "state", "statistics_status", "variants"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ConfidenceLevel = *all.ConfidenceLevel
	o.ConfigVersion = all.ConfigVersion
	o.DataAsOf = all.DataAsOf
	if all.DesiredChange.Get() != nil && !all.DesiredChange.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.DesiredChange = all.DesiredChange
	}
	o.NextRefreshAt = all.NextRefreshAt
	o.OptimalVariantKey = all.OptimalVariantKey
	if all.OptimalVariantUnavailableReason != nil && !all.OptimalVariantUnavailableReason.IsValid() {
		hasInvalidField = true
	} else {
		o.OptimalVariantUnavailableReason = all.OptimalVariantUnavailableReason
	}
	o.PointEstimates = *all.PointEstimates
	o.RefreshId = all.RefreshId
	o.ResultsLastUpdated = all.ResultsLastUpdated
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}
	if all.StatisticsStatus != nil && !all.StatisticsStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.StatisticsStatus = all.StatisticsStatus
	}
	o.Variants = *all.Variants

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
