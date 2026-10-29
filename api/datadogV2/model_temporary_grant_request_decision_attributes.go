// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TemporaryGrantRequestDecisionAttributes Binary decision. Overrides, reason fields, and unknown attributes are rejected, including null values.
type TemporaryGrantRequestDecisionAttributes struct {
	// Approve the exact submitted total and expiry, or deny the request.
	Decision TemporaryGrantRequestDecision `json:"decision"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTemporaryGrantRequestDecisionAttributes instantiates a new TemporaryGrantRequestDecisionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTemporaryGrantRequestDecisionAttributes(decision TemporaryGrantRequestDecision) *TemporaryGrantRequestDecisionAttributes {
	this := TemporaryGrantRequestDecisionAttributes{}
	this.Decision = decision
	return &this
}

// NewTemporaryGrantRequestDecisionAttributesWithDefaults instantiates a new TemporaryGrantRequestDecisionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTemporaryGrantRequestDecisionAttributesWithDefaults() *TemporaryGrantRequestDecisionAttributes {
	this := TemporaryGrantRequestDecisionAttributes{}
	return &this
}

// GetDecision returns the Decision field value.
func (o *TemporaryGrantRequestDecisionAttributes) GetDecision() TemporaryGrantRequestDecision {
	if o == nil {
		var ret TemporaryGrantRequestDecision
		return ret
	}
	return o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value
// and a boolean to check if the value has been set.
func (o *TemporaryGrantRequestDecisionAttributes) GetDecisionOk() (*TemporaryGrantRequestDecision, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decision, true
}

// SetDecision sets field value.
func (o *TemporaryGrantRequestDecisionAttributes) SetDecision(v TemporaryGrantRequestDecision) {
	o.Decision = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TemporaryGrantRequestDecisionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["decision"] = o.Decision
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TemporaryGrantRequestDecisionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Decision *TemporaryGrantRequestDecision `json:"decision"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Decision == nil {
		return fmt.Errorf("required field decision missing")
	}

	hasInvalidField := false
	if !all.Decision.IsValid() {
		hasInvalidField = true
	} else {
		o.Decision = *all.Decision
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
