// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditTrafficVariantV2 An arm's observations and subjects in total.
type ExperimentsBanditTrafficVariantV2 struct {
	// The arm's variant ID on the flag.
	FeatureFlagVariantId uuid.UUID `json:"feature_flag_variant_id"`
	// The arm's valid observations, mature or not.
	Observations int64 `json:"observations"`
	// Subjects whose first valid observation is on this arm. A subject counts once, on the arm its reward goes to.
	Subjects int64 `json:"subjects"`
	// The arm's variant key.
	VariantKey datadog.NullableString `json:"variant_key"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditTrafficVariantV2 instantiates a new ExperimentsBanditTrafficVariantV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditTrafficVariantV2(featureFlagVariantId uuid.UUID, observations int64, subjects int64, variantKey datadog.NullableString) *ExperimentsBanditTrafficVariantV2 {
	this := ExperimentsBanditTrafficVariantV2{}
	this.FeatureFlagVariantId = featureFlagVariantId
	this.Observations = observations
	this.Subjects = subjects
	this.VariantKey = variantKey
	return &this
}

// NewExperimentsBanditTrafficVariantV2WithDefaults instantiates a new ExperimentsBanditTrafficVariantV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditTrafficVariantV2WithDefaults() *ExperimentsBanditTrafficVariantV2 {
	this := ExperimentsBanditTrafficVariantV2{}
	return &this
}

// GetFeatureFlagVariantId returns the FeatureFlagVariantId field value.
func (o *ExperimentsBanditTrafficVariantV2) GetFeatureFlagVariantId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.FeatureFlagVariantId
}

// GetFeatureFlagVariantIdOk returns a tuple with the FeatureFlagVariantId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditTrafficVariantV2) GetFeatureFlagVariantIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureFlagVariantId, true
}

// SetFeatureFlagVariantId sets field value.
func (o *ExperimentsBanditTrafficVariantV2) SetFeatureFlagVariantId(v uuid.UUID) {
	o.FeatureFlagVariantId = v
}

// GetObservations returns the Observations field value.
func (o *ExperimentsBanditTrafficVariantV2) GetObservations() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Observations
}

// GetObservationsOk returns a tuple with the Observations field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditTrafficVariantV2) GetObservationsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Observations, true
}

// SetObservations sets field value.
func (o *ExperimentsBanditTrafficVariantV2) SetObservations(v int64) {
	o.Observations = v
}

// GetSubjects returns the Subjects field value.
func (o *ExperimentsBanditTrafficVariantV2) GetSubjects() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Subjects
}

// GetSubjectsOk returns a tuple with the Subjects field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditTrafficVariantV2) GetSubjectsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subjects, true
}

// SetSubjects sets field value.
func (o *ExperimentsBanditTrafficVariantV2) SetSubjects(v int64) {
	o.Subjects = v
}

// GetVariantKey returns the VariantKey field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *ExperimentsBanditTrafficVariantV2) GetVariantKey() string {
	if o == nil || o.VariantKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.VariantKey.Get()
}

// GetVariantKeyOk returns a tuple with the VariantKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditTrafficVariantV2) GetVariantKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VariantKey.Get(), o.VariantKey.IsSet()
}

// SetVariantKey sets field value.
func (o *ExperimentsBanditTrafficVariantV2) SetVariantKey(v string) {
	o.VariantKey.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditTrafficVariantV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["feature_flag_variant_id"] = o.FeatureFlagVariantId
	toSerialize["observations"] = o.Observations
	toSerialize["subjects"] = o.Subjects
	toSerialize["variant_key"] = o.VariantKey.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditTrafficVariantV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FeatureFlagVariantId *uuid.UUID             `json:"feature_flag_variant_id"`
		Observations         *int64                 `json:"observations"`
		Subjects             *int64                 `json:"subjects"`
		VariantKey           datadog.NullableString `json:"variant_key"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FeatureFlagVariantId == nil {
		return fmt.Errorf("required field feature_flag_variant_id missing")
	}
	if all.Observations == nil {
		return fmt.Errorf("required field observations missing")
	}
	if all.Subjects == nil {
		return fmt.Errorf("required field subjects missing")
	}
	if !all.VariantKey.IsSet() {
		return fmt.Errorf("required field variant_key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"feature_flag_variant_id", "observations", "subjects", "variant_key"})
	} else {
		return err
	}
	o.FeatureFlagVariantId = *all.FeatureFlagVariantId
	o.Observations = *all.Observations
	o.Subjects = *all.Subjects
	o.VariantKey = all.VariantKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
