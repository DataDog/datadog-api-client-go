// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditPolicyVariantV2 A variant's share of traffic under a policy.
type ExperimentsBanditPolicyVariantV2 struct {
	// The variant's ID on the flag.
	FeatureFlagVariantId uuid.UUID `json:"feature_flag_variant_id"`
	// The share of traffic the variant gets, from 0 to 1.
	Probability float64 `json:"probability"`
	// The variant's key. Null when it can't be resolved.
	VariantKey datadog.NullableString `json:"variant_key"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditPolicyVariantV2 instantiates a new ExperimentsBanditPolicyVariantV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditPolicyVariantV2(featureFlagVariantId uuid.UUID, probability float64, variantKey datadog.NullableString) *ExperimentsBanditPolicyVariantV2 {
	this := ExperimentsBanditPolicyVariantV2{}
	this.FeatureFlagVariantId = featureFlagVariantId
	this.Probability = probability
	this.VariantKey = variantKey
	return &this
}

// NewExperimentsBanditPolicyVariantV2WithDefaults instantiates a new ExperimentsBanditPolicyVariantV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditPolicyVariantV2WithDefaults() *ExperimentsBanditPolicyVariantV2 {
	this := ExperimentsBanditPolicyVariantV2{}
	return &this
}

// GetFeatureFlagVariantId returns the FeatureFlagVariantId field value.
func (o *ExperimentsBanditPolicyVariantV2) GetFeatureFlagVariantId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.FeatureFlagVariantId
}

// GetFeatureFlagVariantIdOk returns a tuple with the FeatureFlagVariantId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditPolicyVariantV2) GetFeatureFlagVariantIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureFlagVariantId, true
}

// SetFeatureFlagVariantId sets field value.
func (o *ExperimentsBanditPolicyVariantV2) SetFeatureFlagVariantId(v uuid.UUID) {
	o.FeatureFlagVariantId = v
}

// GetProbability returns the Probability field value.
func (o *ExperimentsBanditPolicyVariantV2) GetProbability() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Probability
}

// GetProbabilityOk returns a tuple with the Probability field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditPolicyVariantV2) GetProbabilityOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Probability, true
}

// SetProbability sets field value.
func (o *ExperimentsBanditPolicyVariantV2) SetProbability(v float64) {
	o.Probability = v
}

// GetVariantKey returns the VariantKey field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *ExperimentsBanditPolicyVariantV2) GetVariantKey() string {
	if o == nil || o.VariantKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.VariantKey.Get()
}

// GetVariantKeyOk returns a tuple with the VariantKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditPolicyVariantV2) GetVariantKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VariantKey.Get(), o.VariantKey.IsSet()
}

// SetVariantKey sets field value.
func (o *ExperimentsBanditPolicyVariantV2) SetVariantKey(v string) {
	o.VariantKey.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditPolicyVariantV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["feature_flag_variant_id"] = o.FeatureFlagVariantId
	toSerialize["probability"] = o.Probability
	toSerialize["variant_key"] = o.VariantKey.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditPolicyVariantV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FeatureFlagVariantId *uuid.UUID             `json:"feature_flag_variant_id"`
		Probability          *float64               `json:"probability"`
		VariantKey           datadog.NullableString `json:"variant_key"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FeatureFlagVariantId == nil {
		return fmt.Errorf("required field feature_flag_variant_id missing")
	}
	if all.Probability == nil {
		return fmt.Errorf("required field probability missing")
	}
	if !all.VariantKey.IsSet() {
		return fmt.Errorf("required field variant_key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"feature_flag_variant_id", "probability", "variant_key"})
	} else {
		return err
	}
	o.FeatureFlagVariantId = *all.FeatureFlagVariantId
	o.Probability = *all.Probability
	o.VariantKey = all.VariantKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
