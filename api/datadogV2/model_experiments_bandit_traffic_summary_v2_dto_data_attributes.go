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

// ExperimentsBanditTrafficSummaryV2DTODataAttributes What the bandit has observed, as of its latest saved refresh.
type ExperimentsBanditTrafficSummaryV2DTODataAttributes struct {
	// Valid observations per arm on the UTC day of their anchor time, for each day from the bandit's start to the data cutoff that overlaps the requested range. Days without any have zero counts. Empty before the first saved refresh.
	Daily []ExperimentsBanditDailyTrafficV2 `json:"daily"`
	// The latest warehouse data the latest saved refresh read.
	DataAsOf datadog.NullableTime `json:"data_as_of"`
	// How many observations the latest saved refresh saw, whatever the range. Null when its counts can't be read.
	ObservationCounts NullableExperimentsBanditObservationCountsV2 `json:"observation_counts"`
	// The latest refresh that saved results. The counts come from it.
	RefreshId datadog.NullableUUID `json:"refresh_id"`
	// Distinct subjects with a valid observation, whatever the range.
	TotalSubjects int64 `json:"total_subjects"`
	// One row per arm, whatever the range.
	Variants []ExperimentsBanditTrafficVariantV2 `json:"variants"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditTrafficSummaryV2DTODataAttributes instantiates a new ExperimentsBanditTrafficSummaryV2DTODataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditTrafficSummaryV2DTODataAttributes(daily []ExperimentsBanditDailyTrafficV2, dataAsOf datadog.NullableTime, observationCounts NullableExperimentsBanditObservationCountsV2, refreshId datadog.NullableUUID, totalSubjects int64, variants []ExperimentsBanditTrafficVariantV2) *ExperimentsBanditTrafficSummaryV2DTODataAttributes {
	this := ExperimentsBanditTrafficSummaryV2DTODataAttributes{}
	this.Daily = daily
	this.DataAsOf = dataAsOf
	this.ObservationCounts = observationCounts
	this.RefreshId = refreshId
	this.TotalSubjects = totalSubjects
	this.Variants = variants
	return &this
}

// NewExperimentsBanditTrafficSummaryV2DTODataAttributesWithDefaults instantiates a new ExperimentsBanditTrafficSummaryV2DTODataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditTrafficSummaryV2DTODataAttributesWithDefaults() *ExperimentsBanditTrafficSummaryV2DTODataAttributes {
	this := ExperimentsBanditTrafficSummaryV2DTODataAttributes{}
	return &this
}

// GetDaily returns the Daily field value.
func (o *ExperimentsBanditTrafficSummaryV2DTODataAttributes) GetDaily() []ExperimentsBanditDailyTrafficV2 {
	if o == nil {
		var ret []ExperimentsBanditDailyTrafficV2
		return ret
	}
	return o.Daily
}

// GetDailyOk returns a tuple with the Daily field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditTrafficSummaryV2DTODataAttributes) GetDailyOk() (*[]ExperimentsBanditDailyTrafficV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Daily, true
}

// SetDaily sets field value.
func (o *ExperimentsBanditTrafficSummaryV2DTODataAttributes) SetDaily(v []ExperimentsBanditDailyTrafficV2) {
	o.Daily = v
}

// GetDataAsOf returns the DataAsOf field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *ExperimentsBanditTrafficSummaryV2DTODataAttributes) GetDataAsOf() time.Time {
	if o == nil || o.DataAsOf.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DataAsOf.Get()
}

// GetDataAsOfOk returns a tuple with the DataAsOf field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditTrafficSummaryV2DTODataAttributes) GetDataAsOfOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataAsOf.Get(), o.DataAsOf.IsSet()
}

// SetDataAsOf sets field value.
func (o *ExperimentsBanditTrafficSummaryV2DTODataAttributes) SetDataAsOf(v time.Time) {
	o.DataAsOf.Set(&v)
}

// GetObservationCounts returns the ObservationCounts field value.
// If the value is explicit nil, the zero value for ExperimentsBanditObservationCountsV2 will be returned.
func (o *ExperimentsBanditTrafficSummaryV2DTODataAttributes) GetObservationCounts() ExperimentsBanditObservationCountsV2 {
	if o == nil || o.ObservationCounts.Get() == nil {
		var ret ExperimentsBanditObservationCountsV2
		return ret
	}
	return *o.ObservationCounts.Get()
}

// GetObservationCountsOk returns a tuple with the ObservationCounts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditTrafficSummaryV2DTODataAttributes) GetObservationCountsOk() (*ExperimentsBanditObservationCountsV2, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObservationCounts.Get(), o.ObservationCounts.IsSet()
}

// SetObservationCounts sets field value.
func (o *ExperimentsBanditTrafficSummaryV2DTODataAttributes) SetObservationCounts(v ExperimentsBanditObservationCountsV2) {
	o.ObservationCounts.Set(&v)
}

// GetRefreshId returns the RefreshId field value.
// If the value is explicit nil, the zero value for uuid.UUID will be returned.
func (o *ExperimentsBanditTrafficSummaryV2DTODataAttributes) GetRefreshId() uuid.UUID {
	if o == nil || o.RefreshId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.RefreshId.Get()
}

// GetRefreshIdOk returns a tuple with the RefreshId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditTrafficSummaryV2DTODataAttributes) GetRefreshIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshId.Get(), o.RefreshId.IsSet()
}

// SetRefreshId sets field value.
func (o *ExperimentsBanditTrafficSummaryV2DTODataAttributes) SetRefreshId(v uuid.UUID) {
	o.RefreshId.Set(&v)
}

// GetTotalSubjects returns the TotalSubjects field value.
func (o *ExperimentsBanditTrafficSummaryV2DTODataAttributes) GetTotalSubjects() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalSubjects
}

// GetTotalSubjectsOk returns a tuple with the TotalSubjects field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditTrafficSummaryV2DTODataAttributes) GetTotalSubjectsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSubjects, true
}

// SetTotalSubjects sets field value.
func (o *ExperimentsBanditTrafficSummaryV2DTODataAttributes) SetTotalSubjects(v int64) {
	o.TotalSubjects = v
}

// GetVariants returns the Variants field value.
func (o *ExperimentsBanditTrafficSummaryV2DTODataAttributes) GetVariants() []ExperimentsBanditTrafficVariantV2 {
	if o == nil {
		var ret []ExperimentsBanditTrafficVariantV2
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditTrafficSummaryV2DTODataAttributes) GetVariantsOk() (*[]ExperimentsBanditTrafficVariantV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variants, true
}

// SetVariants sets field value.
func (o *ExperimentsBanditTrafficSummaryV2DTODataAttributes) SetVariants(v []ExperimentsBanditTrafficVariantV2) {
	o.Variants = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditTrafficSummaryV2DTODataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["daily"] = o.Daily
	toSerialize["data_as_of"] = o.DataAsOf.Get()
	toSerialize["observation_counts"] = o.ObservationCounts.Get()
	toSerialize["refresh_id"] = o.RefreshId.Get()
	toSerialize["total_subjects"] = o.TotalSubjects
	toSerialize["variants"] = o.Variants

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditTrafficSummaryV2DTODataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Daily             *[]ExperimentsBanditDailyTrafficV2           `json:"daily"`
		DataAsOf          datadog.NullableTime                         `json:"data_as_of"`
		ObservationCounts NullableExperimentsBanditObservationCountsV2 `json:"observation_counts"`
		RefreshId         datadog.NullableUUID                         `json:"refresh_id"`
		TotalSubjects     *int64                                       `json:"total_subjects"`
		Variants          *[]ExperimentsBanditTrafficVariantV2         `json:"variants"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Daily == nil {
		return fmt.Errorf("required field daily missing")
	}
	if !all.DataAsOf.IsSet() {
		return fmt.Errorf("required field data_as_of missing")
	}
	if !all.ObservationCounts.IsSet() {
		return fmt.Errorf("required field observation_counts missing")
	}
	if !all.RefreshId.IsSet() {
		return fmt.Errorf("required field refresh_id missing")
	}
	if all.TotalSubjects == nil {
		return fmt.Errorf("required field total_subjects missing")
	}
	if all.Variants == nil {
		return fmt.Errorf("required field variants missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"daily", "data_as_of", "observation_counts", "refresh_id", "total_subjects", "variants"})
	} else {
		return err
	}
	o.Daily = *all.Daily
	o.DataAsOf = all.DataAsOf
	o.ObservationCounts = all.ObservationCounts
	o.RefreshId = all.RefreshId
	o.TotalSubjects = *all.TotalSubjects
	o.Variants = *all.Variants

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
