// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardConfidenceThreshold Request-level confidence threshold override. `flagging_global` is the global
// threshold applied to every threat category. Any other property is a
// per-category override keyed by an AI Guard threat tag; per-category values
// take precedence over `flagging_global`, and categories without an entry fall
// back to it. When this object omits flagging_global, the fallback is `0.5`.
// Recognized category keys include policy-evasion, indirect-prompt-injection,
// system-and-capability-discovery, credential-access, sensitive-data-collection,
// privilege-escalation, malicious-execution, private-network-target,
// destructive-or-resource-abuse, tool-abuse, poisoning-or-persistence, and
// `data-exfiltration`. Unknown category keys are ignored.
type AIGuardConfidenceThreshold struct {
	// The global confidence threshold applied to categories without a specific override.
	FlaggingGlobal *float64 `json:"flagging_global,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]float64     `json:"-"`
}

// NewAIGuardConfidenceThreshold instantiates a new AIGuardConfidenceThreshold object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAIGuardConfidenceThreshold() *AIGuardConfidenceThreshold {
	this := AIGuardConfidenceThreshold{}
	return &this
}

// NewAIGuardConfidenceThresholdWithDefaults instantiates a new AIGuardConfidenceThreshold object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAIGuardConfidenceThresholdWithDefaults() *AIGuardConfidenceThreshold {
	this := AIGuardConfidenceThreshold{}
	return &this
}

// GetFlaggingGlobal returns the FlaggingGlobal field value if set, zero value otherwise.
func (o *AIGuardConfidenceThreshold) GetFlaggingGlobal() float64 {
	if o == nil || o.FlaggingGlobal == nil {
		var ret float64
		return ret
	}
	return *o.FlaggingGlobal
}

// GetFlaggingGlobalOk returns a tuple with the FlaggingGlobal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardConfidenceThreshold) GetFlaggingGlobalOk() (*float64, bool) {
	if o == nil || o.FlaggingGlobal == nil {
		return nil, false
	}
	return o.FlaggingGlobal, true
}

// HasFlaggingGlobal returns a boolean if a field has been set.
func (o *AIGuardConfidenceThreshold) HasFlaggingGlobal() bool {
	return o != nil && o.FlaggingGlobal != nil
}

// SetFlaggingGlobal gets a reference to the given float64 and assigns it to the FlaggingGlobal field.
func (o *AIGuardConfidenceThreshold) SetFlaggingGlobal(v float64) {
	o.FlaggingGlobal = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AIGuardConfidenceThreshold) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FlaggingGlobal != nil {
		toSerialize["flagging_global"] = o.FlaggingGlobal
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AIGuardConfidenceThreshold) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FlaggingGlobal *float64 `json:"flagging_global,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]float64)
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"flagging_global"})
	} else {
		return err
	}
	o.FlaggingGlobal = all.FlaggingGlobal

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
