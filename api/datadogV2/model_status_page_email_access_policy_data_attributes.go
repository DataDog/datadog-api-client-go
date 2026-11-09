// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPageEmailAccessPolicyDataAttributes Attributes of a status page email access policy.
type StatusPageEmailAccessPolicyDataAttributes struct {
	// Timestamp of when the email access policy was created.
	CreatedAt time.Time `json:"created_at"`
	// The effect applied when an email address matches the policy. Deny policies take precedence over allow policies.
	Effect AccessPolicyEffect `json:"effect"`
	// The email address pattern matched by the policy.
	Pattern string `json:"pattern"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStatusPageEmailAccessPolicyDataAttributes instantiates a new StatusPageEmailAccessPolicyDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStatusPageEmailAccessPolicyDataAttributes(createdAt time.Time, effect AccessPolicyEffect, pattern string) *StatusPageEmailAccessPolicyDataAttributes {
	this := StatusPageEmailAccessPolicyDataAttributes{}
	this.CreatedAt = createdAt
	this.Effect = effect
	this.Pattern = pattern
	return &this
}

// NewStatusPageEmailAccessPolicyDataAttributesWithDefaults instantiates a new StatusPageEmailAccessPolicyDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStatusPageEmailAccessPolicyDataAttributesWithDefaults() *StatusPageEmailAccessPolicyDataAttributes {
	this := StatusPageEmailAccessPolicyDataAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *StatusPageEmailAccessPolicyDataAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *StatusPageEmailAccessPolicyDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *StatusPageEmailAccessPolicyDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEffect returns the Effect field value.
func (o *StatusPageEmailAccessPolicyDataAttributes) GetEffect() AccessPolicyEffect {
	if o == nil {
		var ret AccessPolicyEffect
		return ret
	}
	return o.Effect
}

// GetEffectOk returns a tuple with the Effect field value
// and a boolean to check if the value has been set.
func (o *StatusPageEmailAccessPolicyDataAttributes) GetEffectOk() (*AccessPolicyEffect, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Effect, true
}

// SetEffect sets field value.
func (o *StatusPageEmailAccessPolicyDataAttributes) SetEffect(v AccessPolicyEffect) {
	o.Effect = v
}

// GetPattern returns the Pattern field value.
func (o *StatusPageEmailAccessPolicyDataAttributes) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *StatusPageEmailAccessPolicyDataAttributes) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value.
func (o *StatusPageEmailAccessPolicyDataAttributes) SetPattern(v string) {
	o.Pattern = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StatusPageEmailAccessPolicyDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["effect"] = o.Effect
	toSerialize["pattern"] = o.Pattern

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StatusPageEmailAccessPolicyDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt *time.Time          `json:"created_at"`
		Effect    *AccessPolicyEffect `json:"effect"`
		Pattern   *string             `json:"pattern"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Effect == nil {
		return fmt.Errorf("required field effect missing")
	}
	if all.Pattern == nil {
		return fmt.Errorf("required field pattern missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "effect", "pattern"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
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
