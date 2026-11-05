// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayPolicyRequestAttributes Attributes used to create or replace a routing policy document.
type AiGatewayPolicyRequestAttributes struct {
	// Inbound request pattern.
	Inbound string `json:"inbound"`
	// Routing priority.
	Priority *int64 `json:"priority,omitempty"`
	//
	Tiers []AiGatewayPolicyTierRequest `json:"tiers"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayPolicyRequestAttributes instantiates a new AiGatewayPolicyRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayPolicyRequestAttributes(inbound string, tiers []AiGatewayPolicyTierRequest) *AiGatewayPolicyRequestAttributes {
	this := AiGatewayPolicyRequestAttributes{}
	this.Inbound = inbound
	this.Tiers = tiers
	return &this
}

// NewAiGatewayPolicyRequestAttributesWithDefaults instantiates a new AiGatewayPolicyRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayPolicyRequestAttributesWithDefaults() *AiGatewayPolicyRequestAttributes {
	this := AiGatewayPolicyRequestAttributes{}
	return &this
}

// GetInbound returns the Inbound field value.
func (o *AiGatewayPolicyRequestAttributes) GetInbound() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Inbound
}

// GetInboundOk returns a tuple with the Inbound field value
// and a boolean to check if the value has been set.
func (o *AiGatewayPolicyRequestAttributes) GetInboundOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inbound, true
}

// SetInbound sets field value.
func (o *AiGatewayPolicyRequestAttributes) SetInbound(v string) {
	o.Inbound = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *AiGatewayPolicyRequestAttributes) GetPriority() int64 {
	if o == nil || o.Priority == nil {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayPolicyRequestAttributes) GetPriorityOk() (*int64, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *AiGatewayPolicyRequestAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *AiGatewayPolicyRequestAttributes) SetPriority(v int64) {
	o.Priority = &v
}

// GetTiers returns the Tiers field value.
func (o *AiGatewayPolicyRequestAttributes) GetTiers() []AiGatewayPolicyTierRequest {
	if o == nil {
		var ret []AiGatewayPolicyTierRequest
		return ret
	}
	return o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value
// and a boolean to check if the value has been set.
func (o *AiGatewayPolicyRequestAttributes) GetTiersOk() (*[]AiGatewayPolicyTierRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tiers, true
}

// SetTiers sets field value.
func (o *AiGatewayPolicyRequestAttributes) SetTiers(v []AiGatewayPolicyTierRequest) {
	o.Tiers = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayPolicyRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["inbound"] = o.Inbound
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	toSerialize["tiers"] = o.Tiers

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayPolicyRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Inbound  *string                       `json:"inbound"`
		Priority *int64                        `json:"priority,omitempty"`
		Tiers    *[]AiGatewayPolicyTierRequest `json:"tiers"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Inbound == nil {
		return fmt.Errorf("required field inbound missing")
	}
	if all.Tiers == nil {
		return fmt.Errorf("required field tiers missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"inbound", "priority", "tiers"})
	} else {
		return err
	}
	o.Inbound = *all.Inbound
	o.Priority = all.Priority
	o.Tiers = *all.Tiers

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
