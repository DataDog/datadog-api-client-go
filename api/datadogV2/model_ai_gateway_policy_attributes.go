// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayPolicyAttributes Routing policy configuration.
type AiGatewayPolicyAttributes struct {
	// Routing priority.
	Priority int64 `json:"priority"`
	//
	Tiers []AiGatewayPolicyTier `json:"tiers"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayPolicyAttributes instantiates a new AiGatewayPolicyAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayPolicyAttributes(priority int64, tiers []AiGatewayPolicyTier) *AiGatewayPolicyAttributes {
	this := AiGatewayPolicyAttributes{}
	this.Priority = priority
	this.Tiers = tiers
	return &this
}

// NewAiGatewayPolicyAttributesWithDefaults instantiates a new AiGatewayPolicyAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayPolicyAttributesWithDefaults() *AiGatewayPolicyAttributes {
	this := AiGatewayPolicyAttributes{}
	return &this
}

// GetPriority returns the Priority field value.
func (o *AiGatewayPolicyAttributes) GetPriority() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *AiGatewayPolicyAttributes) GetPriorityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value.
func (o *AiGatewayPolicyAttributes) SetPriority(v int64) {
	o.Priority = v
}

// GetTiers returns the Tiers field value.
func (o *AiGatewayPolicyAttributes) GetTiers() []AiGatewayPolicyTier {
	if o == nil {
		var ret []AiGatewayPolicyTier
		return ret
	}
	return o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value
// and a boolean to check if the value has been set.
func (o *AiGatewayPolicyAttributes) GetTiersOk() (*[]AiGatewayPolicyTier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tiers, true
}

// SetTiers sets field value.
func (o *AiGatewayPolicyAttributes) SetTiers(v []AiGatewayPolicyTier) {
	o.Tiers = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayPolicyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["priority"] = o.Priority
	toSerialize["tiers"] = o.Tiers

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayPolicyAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Priority *int64                 `json:"priority"`
		Tiers    *[]AiGatewayPolicyTier `json:"tiers"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Priority == nil {
		return fmt.Errorf("required field priority missing")
	}
	if all.Tiers == nil {
		return fmt.Errorf("required field tiers missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"priority", "tiers"})
	} else {
		return err
	}
	o.Priority = *all.Priority
	o.Tiers = *all.Tiers

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
