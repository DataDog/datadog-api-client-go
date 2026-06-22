// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationGuardrailDecision The guardrail decision applied to a plan or investigation.
type RemediationGuardrailDecision struct {
	// The verdict a guardrail applied to a plan or investigation.
	Decision *RemediationGuardrailVerdict `json:"decision,omitempty"`
	// ID of the matching guardrail. Set when the decision is not denied; may be empty when denied because no rule matched.
	GuardrailId *string `json:"guardrail_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationGuardrailDecision instantiates a new RemediationGuardrailDecision object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationGuardrailDecision() *RemediationGuardrailDecision {
	this := RemediationGuardrailDecision{}
	return &this
}

// NewRemediationGuardrailDecisionWithDefaults instantiates a new RemediationGuardrailDecision object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationGuardrailDecisionWithDefaults() *RemediationGuardrailDecision {
	this := RemediationGuardrailDecision{}
	return &this
}

// GetDecision returns the Decision field value if set, zero value otherwise.
func (o *RemediationGuardrailDecision) GetDecision() RemediationGuardrailVerdict {
	if o == nil || o.Decision == nil {
		var ret RemediationGuardrailVerdict
		return ret
	}
	return *o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationGuardrailDecision) GetDecisionOk() (*RemediationGuardrailVerdict, bool) {
	if o == nil || o.Decision == nil {
		return nil, false
	}
	return o.Decision, true
}

// HasDecision returns a boolean if a field has been set.
func (o *RemediationGuardrailDecision) HasDecision() bool {
	return o != nil && o.Decision != nil
}

// SetDecision gets a reference to the given RemediationGuardrailVerdict and assigns it to the Decision field.
func (o *RemediationGuardrailDecision) SetDecision(v RemediationGuardrailVerdict) {
	o.Decision = &v
}

// GetGuardrailId returns the GuardrailId field value if set, zero value otherwise.
func (o *RemediationGuardrailDecision) GetGuardrailId() string {
	if o == nil || o.GuardrailId == nil {
		var ret string
		return ret
	}
	return *o.GuardrailId
}

// GetGuardrailIdOk returns a tuple with the GuardrailId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationGuardrailDecision) GetGuardrailIdOk() (*string, bool) {
	if o == nil || o.GuardrailId == nil {
		return nil, false
	}
	return o.GuardrailId, true
}

// HasGuardrailId returns a boolean if a field has been set.
func (o *RemediationGuardrailDecision) HasGuardrailId() bool {
	return o != nil && o.GuardrailId != nil
}

// SetGuardrailId gets a reference to the given string and assigns it to the GuardrailId field.
func (o *RemediationGuardrailDecision) SetGuardrailId(v string) {
	o.GuardrailId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationGuardrailDecision) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Decision != nil {
		toSerialize["decision"] = o.Decision
	}
	if o.GuardrailId != nil {
		toSerialize["guardrail_id"] = o.GuardrailId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationGuardrailDecision) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Decision    *RemediationGuardrailVerdict `json:"decision,omitempty"`
		GuardrailId *string                      `json:"guardrail_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"decision", "guardrail_id"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Decision != nil && !all.Decision.IsValid() {
		hasInvalidField = true
	} else {
		o.Decision = all.Decision
	}
	o.GuardrailId = all.GuardrailId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
