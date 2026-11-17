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

// ExperimentsBanditDailyTrafficV2 An arm's observations and new subjects on one UTC day.
type ExperimentsBanditDailyTrafficV2 struct {
	// The arm's new subjects from the bandit's start through this day, including days before the requested range.
	CumulativeNewSubjects int64 `json:"cumulative_new_subjects"`
	// The arm's observations from the bandit's start through this day, including days before the requested range.
	CumulativeObservations int64 `json:"cumulative_observations"`
	// A UTC date as YYYY-MM-DD.
	Date time.Time `json:"date"`
	// The arm's variant ID on the flag.
	FeatureFlagVariantId uuid.UUID `json:"feature_flag_variant_id"`
	// Subjects whose first valid observation is on this day and arm.
	NewSubjects int64 `json:"new_subjects"`
	// The arm's valid observations on this day.
	Observations int64 `json:"observations"`
	// The arm's variant key.
	VariantKey datadog.NullableString `json:"variant_key"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditDailyTrafficV2 instantiates a new ExperimentsBanditDailyTrafficV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditDailyTrafficV2(cumulativeNewSubjects int64, cumulativeObservations int64, date time.Time, featureFlagVariantId uuid.UUID, newSubjects int64, observations int64, variantKey datadog.NullableString) *ExperimentsBanditDailyTrafficV2 {
	this := ExperimentsBanditDailyTrafficV2{}
	this.CumulativeNewSubjects = cumulativeNewSubjects
	this.CumulativeObservations = cumulativeObservations
	this.Date = date
	this.FeatureFlagVariantId = featureFlagVariantId
	this.NewSubjects = newSubjects
	this.Observations = observations
	this.VariantKey = variantKey
	return &this
}

// NewExperimentsBanditDailyTrafficV2WithDefaults instantiates a new ExperimentsBanditDailyTrafficV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditDailyTrafficV2WithDefaults() *ExperimentsBanditDailyTrafficV2 {
	this := ExperimentsBanditDailyTrafficV2{}
	return &this
}

// GetCumulativeNewSubjects returns the CumulativeNewSubjects field value.
func (o *ExperimentsBanditDailyTrafficV2) GetCumulativeNewSubjects() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CumulativeNewSubjects
}

// GetCumulativeNewSubjectsOk returns a tuple with the CumulativeNewSubjects field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditDailyTrafficV2) GetCumulativeNewSubjectsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CumulativeNewSubjects, true
}

// SetCumulativeNewSubjects sets field value.
func (o *ExperimentsBanditDailyTrafficV2) SetCumulativeNewSubjects(v int64) {
	o.CumulativeNewSubjects = v
}

// GetCumulativeObservations returns the CumulativeObservations field value.
func (o *ExperimentsBanditDailyTrafficV2) GetCumulativeObservations() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CumulativeObservations
}

// GetCumulativeObservationsOk returns a tuple with the CumulativeObservations field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditDailyTrafficV2) GetCumulativeObservationsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CumulativeObservations, true
}

// SetCumulativeObservations sets field value.
func (o *ExperimentsBanditDailyTrafficV2) SetCumulativeObservations(v int64) {
	o.CumulativeObservations = v
}

// GetDate returns the Date field value.
func (o *ExperimentsBanditDailyTrafficV2) GetDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditDailyTrafficV2) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value.
func (o *ExperimentsBanditDailyTrafficV2) SetDate(v time.Time) {
	o.Date = v
}

// GetFeatureFlagVariantId returns the FeatureFlagVariantId field value.
func (o *ExperimentsBanditDailyTrafficV2) GetFeatureFlagVariantId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.FeatureFlagVariantId
}

// GetFeatureFlagVariantIdOk returns a tuple with the FeatureFlagVariantId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditDailyTrafficV2) GetFeatureFlagVariantIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureFlagVariantId, true
}

// SetFeatureFlagVariantId sets field value.
func (o *ExperimentsBanditDailyTrafficV2) SetFeatureFlagVariantId(v uuid.UUID) {
	o.FeatureFlagVariantId = v
}

// GetNewSubjects returns the NewSubjects field value.
func (o *ExperimentsBanditDailyTrafficV2) GetNewSubjects() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.NewSubjects
}

// GetNewSubjectsOk returns a tuple with the NewSubjects field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditDailyTrafficV2) GetNewSubjectsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewSubjects, true
}

// SetNewSubjects sets field value.
func (o *ExperimentsBanditDailyTrafficV2) SetNewSubjects(v int64) {
	o.NewSubjects = v
}

// GetObservations returns the Observations field value.
func (o *ExperimentsBanditDailyTrafficV2) GetObservations() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Observations
}

// GetObservationsOk returns a tuple with the Observations field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditDailyTrafficV2) GetObservationsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Observations, true
}

// SetObservations sets field value.
func (o *ExperimentsBanditDailyTrafficV2) SetObservations(v int64) {
	o.Observations = v
}

// GetVariantKey returns the VariantKey field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *ExperimentsBanditDailyTrafficV2) GetVariantKey() string {
	if o == nil || o.VariantKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.VariantKey.Get()
}

// GetVariantKeyOk returns a tuple with the VariantKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditDailyTrafficV2) GetVariantKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VariantKey.Get(), o.VariantKey.IsSet()
}

// SetVariantKey sets field value.
func (o *ExperimentsBanditDailyTrafficV2) SetVariantKey(v string) {
	o.VariantKey.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditDailyTrafficV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cumulative_new_subjects"] = o.CumulativeNewSubjects
	toSerialize["cumulative_observations"] = o.CumulativeObservations
	toSerialize["date"] = o.Date
	toSerialize["feature_flag_variant_id"] = o.FeatureFlagVariantId
	toSerialize["new_subjects"] = o.NewSubjects
	toSerialize["observations"] = o.Observations
	toSerialize["variant_key"] = o.VariantKey.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditDailyTrafficV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CumulativeNewSubjects  *int64                 `json:"cumulative_new_subjects"`
		CumulativeObservations *int64                 `json:"cumulative_observations"`
		Date                   *time.Time             `json:"date"`
		FeatureFlagVariantId   *uuid.UUID             `json:"feature_flag_variant_id"`
		NewSubjects            *int64                 `json:"new_subjects"`
		Observations           *int64                 `json:"observations"`
		VariantKey             datadog.NullableString `json:"variant_key"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CumulativeNewSubjects == nil {
		return fmt.Errorf("required field cumulative_new_subjects missing")
	}
	if all.CumulativeObservations == nil {
		return fmt.Errorf("required field cumulative_observations missing")
	}
	if all.Date == nil {
		return fmt.Errorf("required field date missing")
	}
	if all.FeatureFlagVariantId == nil {
		return fmt.Errorf("required field feature_flag_variant_id missing")
	}
	if all.NewSubjects == nil {
		return fmt.Errorf("required field new_subjects missing")
	}
	if all.Observations == nil {
		return fmt.Errorf("required field observations missing")
	}
	if !all.VariantKey.IsSet() {
		return fmt.Errorf("required field variant_key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cumulative_new_subjects", "cumulative_observations", "date", "feature_flag_variant_id", "new_subjects", "observations", "variant_key"})
	} else {
		return err
	}
	o.CumulativeNewSubjects = *all.CumulativeNewSubjects
	o.CumulativeObservations = *all.CumulativeObservations
	o.Date = *all.Date
	o.FeatureFlagVariantId = *all.FeatureFlagVariantId
	o.NewSubjects = *all.NewSubjects
	o.Observations = *all.Observations
	o.VariantKey = all.VariantKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
