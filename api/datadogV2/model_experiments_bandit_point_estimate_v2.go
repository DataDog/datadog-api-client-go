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

// ExperimentsBanditPointEstimateV2 An arm's point estimate from one saved refresh.
type ExperimentsBanditPointEstimateV2 struct {
	// A central interval around point_estimate at confidence_level.
	ConfidenceInterval ExperimentsBanditConfidenceIntervalV2 `json:"confidence_interval"`
	// The latest warehouse data the refresh read.
	DataAsOf datadog.NullableTime `json:"data_as_of"`
	// The arm's variant ID on the flag.
	FeatureFlagVariantId uuid.UUID `json:"feature_flag_variant_id"`
	// The arm's estimated mean of the metric, in the metric's units.
	PointEstimate float64 `json:"point_estimate"`
	// The refresh the estimate comes from.
	RefreshId uuid.UUID `json:"refresh_id"`
	// When the refresh saved its results.
	ResultsLastUpdated time.Time `json:"results_last_updated"`
	// The arm's variant key.
	VariantKey datadog.NullableString `json:"variant_key"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditPointEstimateV2 instantiates a new ExperimentsBanditPointEstimateV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditPointEstimateV2(confidenceInterval ExperimentsBanditConfidenceIntervalV2, dataAsOf datadog.NullableTime, featureFlagVariantId uuid.UUID, pointEstimate float64, refreshId uuid.UUID, resultsLastUpdated time.Time, variantKey datadog.NullableString) *ExperimentsBanditPointEstimateV2 {
	this := ExperimentsBanditPointEstimateV2{}
	this.ConfidenceInterval = confidenceInterval
	this.DataAsOf = dataAsOf
	this.FeatureFlagVariantId = featureFlagVariantId
	this.PointEstimate = pointEstimate
	this.RefreshId = refreshId
	this.ResultsLastUpdated = resultsLastUpdated
	this.VariantKey = variantKey
	return &this
}

// NewExperimentsBanditPointEstimateV2WithDefaults instantiates a new ExperimentsBanditPointEstimateV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditPointEstimateV2WithDefaults() *ExperimentsBanditPointEstimateV2 {
	this := ExperimentsBanditPointEstimateV2{}
	return &this
}

// GetConfidenceInterval returns the ConfidenceInterval field value.
func (o *ExperimentsBanditPointEstimateV2) GetConfidenceInterval() ExperimentsBanditConfidenceIntervalV2 {
	if o == nil {
		var ret ExperimentsBanditConfidenceIntervalV2
		return ret
	}
	return o.ConfidenceInterval
}

// GetConfidenceIntervalOk returns a tuple with the ConfidenceInterval field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditPointEstimateV2) GetConfidenceIntervalOk() (*ExperimentsBanditConfidenceIntervalV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfidenceInterval, true
}

// SetConfidenceInterval sets field value.
func (o *ExperimentsBanditPointEstimateV2) SetConfidenceInterval(v ExperimentsBanditConfidenceIntervalV2) {
	o.ConfidenceInterval = v
}

// GetDataAsOf returns the DataAsOf field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *ExperimentsBanditPointEstimateV2) GetDataAsOf() time.Time {
	if o == nil || o.DataAsOf.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DataAsOf.Get()
}

// GetDataAsOfOk returns a tuple with the DataAsOf field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditPointEstimateV2) GetDataAsOfOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataAsOf.Get(), o.DataAsOf.IsSet()
}

// SetDataAsOf sets field value.
func (o *ExperimentsBanditPointEstimateV2) SetDataAsOf(v time.Time) {
	o.DataAsOf.Set(&v)
}

// GetFeatureFlagVariantId returns the FeatureFlagVariantId field value.
func (o *ExperimentsBanditPointEstimateV2) GetFeatureFlagVariantId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.FeatureFlagVariantId
}

// GetFeatureFlagVariantIdOk returns a tuple with the FeatureFlagVariantId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditPointEstimateV2) GetFeatureFlagVariantIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureFlagVariantId, true
}

// SetFeatureFlagVariantId sets field value.
func (o *ExperimentsBanditPointEstimateV2) SetFeatureFlagVariantId(v uuid.UUID) {
	o.FeatureFlagVariantId = v
}

// GetPointEstimate returns the PointEstimate field value.
func (o *ExperimentsBanditPointEstimateV2) GetPointEstimate() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.PointEstimate
}

// GetPointEstimateOk returns a tuple with the PointEstimate field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditPointEstimateV2) GetPointEstimateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PointEstimate, true
}

// SetPointEstimate sets field value.
func (o *ExperimentsBanditPointEstimateV2) SetPointEstimate(v float64) {
	o.PointEstimate = v
}

// GetRefreshId returns the RefreshId field value.
func (o *ExperimentsBanditPointEstimateV2) GetRefreshId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.RefreshId
}

// GetRefreshIdOk returns a tuple with the RefreshId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditPointEstimateV2) GetRefreshIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefreshId, true
}

// SetRefreshId sets field value.
func (o *ExperimentsBanditPointEstimateV2) SetRefreshId(v uuid.UUID) {
	o.RefreshId = v
}

// GetResultsLastUpdated returns the ResultsLastUpdated field value.
func (o *ExperimentsBanditPointEstimateV2) GetResultsLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ResultsLastUpdated
}

// GetResultsLastUpdatedOk returns a tuple with the ResultsLastUpdated field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditPointEstimateV2) GetResultsLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultsLastUpdated, true
}

// SetResultsLastUpdated sets field value.
func (o *ExperimentsBanditPointEstimateV2) SetResultsLastUpdated(v time.Time) {
	o.ResultsLastUpdated = v
}

// GetVariantKey returns the VariantKey field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *ExperimentsBanditPointEstimateV2) GetVariantKey() string {
	if o == nil || o.VariantKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.VariantKey.Get()
}

// GetVariantKeyOk returns a tuple with the VariantKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditPointEstimateV2) GetVariantKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VariantKey.Get(), o.VariantKey.IsSet()
}

// SetVariantKey sets field value.
func (o *ExperimentsBanditPointEstimateV2) SetVariantKey(v string) {
	o.VariantKey.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditPointEstimateV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["confidence_interval"] = o.ConfidenceInterval
	toSerialize["data_as_of"] = o.DataAsOf.Get()
	toSerialize["feature_flag_variant_id"] = o.FeatureFlagVariantId
	toSerialize["point_estimate"] = o.PointEstimate
	toSerialize["refresh_id"] = o.RefreshId
	if o.ResultsLastUpdated.Nanosecond() == 0 {
		toSerialize["results_last_updated"] = o.ResultsLastUpdated.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["results_last_updated"] = o.ResultsLastUpdated.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["variant_key"] = o.VariantKey.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditPointEstimateV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfidenceInterval   *ExperimentsBanditConfidenceIntervalV2 `json:"confidence_interval"`
		DataAsOf             datadog.NullableTime                   `json:"data_as_of"`
		FeatureFlagVariantId *uuid.UUID                             `json:"feature_flag_variant_id"`
		PointEstimate        *float64                               `json:"point_estimate"`
		RefreshId            *uuid.UUID                             `json:"refresh_id"`
		ResultsLastUpdated   *time.Time                             `json:"results_last_updated"`
		VariantKey           datadog.NullableString                 `json:"variant_key"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConfidenceInterval == nil {
		return fmt.Errorf("required field confidence_interval missing")
	}
	if !all.DataAsOf.IsSet() {
		return fmt.Errorf("required field data_as_of missing")
	}
	if all.FeatureFlagVariantId == nil {
		return fmt.Errorf("required field feature_flag_variant_id missing")
	}
	if all.PointEstimate == nil {
		return fmt.Errorf("required field point_estimate missing")
	}
	if all.RefreshId == nil {
		return fmt.Errorf("required field refresh_id missing")
	}
	if all.ResultsLastUpdated == nil {
		return fmt.Errorf("required field results_last_updated missing")
	}
	if !all.VariantKey.IsSet() {
		return fmt.Errorf("required field variant_key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"confidence_interval", "data_as_of", "feature_flag_variant_id", "point_estimate", "refresh_id", "results_last_updated", "variant_key"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ConfidenceInterval.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ConfidenceInterval = *all.ConfidenceInterval
	o.DataAsOf = all.DataAsOf
	o.FeatureFlagVariantId = *all.FeatureFlagVariantId
	o.PointEstimate = *all.PointEstimate
	o.RefreshId = *all.RefreshId
	o.ResultsLastUpdated = *all.ResultsLastUpdated
	o.VariantKey = all.VariantKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
