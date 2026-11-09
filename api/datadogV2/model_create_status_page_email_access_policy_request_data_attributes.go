// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateStatusPageEmailAccessPolicyRequestDataAttributes Attributes for creating a status page email access policy.
type CreateStatusPageEmailAccessPolicyRequestDataAttributes struct {
	// The effect applied when an email address matches the policy. Deny policies take precedence over allow policies.
	Effect AccessPolicyEffect `json:"effect"`
	// The email address pattern matched by the policy. Supports an exact address (`user@example.com`), all addresses in a domain (`@example.com`), or all addresses in subdomains (`@*.example.com`).
	Pattern string `json:"pattern"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateStatusPageEmailAccessPolicyRequestDataAttributes instantiates a new CreateStatusPageEmailAccessPolicyRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateStatusPageEmailAccessPolicyRequestDataAttributes(effect AccessPolicyEffect, pattern string) *CreateStatusPageEmailAccessPolicyRequestDataAttributes {
	this := CreateStatusPageEmailAccessPolicyRequestDataAttributes{}
	this.Effect = effect
	this.Pattern = pattern
	return &this
}

// NewCreateStatusPageEmailAccessPolicyRequestDataAttributesWithDefaults instantiates a new CreateStatusPageEmailAccessPolicyRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateStatusPageEmailAccessPolicyRequestDataAttributesWithDefaults() *CreateStatusPageEmailAccessPolicyRequestDataAttributes {
	this := CreateStatusPageEmailAccessPolicyRequestDataAttributes{}
	return &this
}

// GetEffect returns the Effect field value.
func (o *CreateStatusPageEmailAccessPolicyRequestDataAttributes) GetEffect() AccessPolicyEffect {
	if o == nil {
		var ret AccessPolicyEffect
		return ret
	}
	return o.Effect
}

// GetEffectOk returns a tuple with the Effect field value
// and a boolean to check if the value has been set.
func (o *CreateStatusPageEmailAccessPolicyRequestDataAttributes) GetEffectOk() (*AccessPolicyEffect, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Effect, true
}

// SetEffect sets field value.
func (o *CreateStatusPageEmailAccessPolicyRequestDataAttributes) SetEffect(v AccessPolicyEffect) {
	o.Effect = v
}

// GetPattern returns the Pattern field value.
func (o *CreateStatusPageEmailAccessPolicyRequestDataAttributes) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *CreateStatusPageEmailAccessPolicyRequestDataAttributes) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value.
func (o *CreateStatusPageEmailAccessPolicyRequestDataAttributes) SetPattern(v string) {
	o.Pattern = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateStatusPageEmailAccessPolicyRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["effect"] = o.Effect
	toSerialize["pattern"] = o.Pattern

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateStatusPageEmailAccessPolicyRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Effect  *AccessPolicyEffect `json:"effect"`
		Pattern *string             `json:"pattern"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Effect == nil {
		return fmt.Errorf("required field effect missing")
	}
	if all.Pattern == nil {
		return fmt.Errorf("required field pattern missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"effect", "pattern"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Effect.IsValid() {
		hasInvalidField = true
	} else {
		o.Effect = *all.Effect
	}
	o.Pattern = *all.Pattern

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
