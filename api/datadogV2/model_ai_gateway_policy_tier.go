// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayPolicyTier Resolved routing policy tier.
type AiGatewayPolicyTier struct {
	//
	Targets []AiGatewayPolicyTarget `json:"targets"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayPolicyTier instantiates a new AiGatewayPolicyTier object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayPolicyTier(targets []AiGatewayPolicyTarget) *AiGatewayPolicyTier {
	this := AiGatewayPolicyTier{}
	this.Targets = targets
	return &this
}

// NewAiGatewayPolicyTierWithDefaults instantiates a new AiGatewayPolicyTier object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayPolicyTierWithDefaults() *AiGatewayPolicyTier {
	this := AiGatewayPolicyTier{}
	return &this
}

// GetTargets returns the Targets field value.
func (o *AiGatewayPolicyTier) GetTargets() []AiGatewayPolicyTarget {
	if o == nil {
		var ret []AiGatewayPolicyTarget
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *AiGatewayPolicyTier) GetTargetsOk() (*[]AiGatewayPolicyTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Targets, true
}

// SetTargets sets field value.
func (o *AiGatewayPolicyTier) SetTargets(v []AiGatewayPolicyTarget) {
	o.Targets = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayPolicyTier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["targets"] = o.Targets

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayPolicyTier) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Targets *[]AiGatewayPolicyTarget `json:"targets"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Targets == nil {
		return fmt.Errorf("required field targets missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"targets"})
	} else {
		return err
	}
	o.Targets = *all.Targets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
