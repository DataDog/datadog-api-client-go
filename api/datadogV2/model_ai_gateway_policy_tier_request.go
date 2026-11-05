// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayPolicyTierRequest Routing policy tier supplied in a request.
type AiGatewayPolicyTierRequest struct {
	// Routing priority.
	Priority *int64 `json:"priority,omitempty"`
	//
	Targets []AiGatewayPolicyTargetRequest `json:"targets"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayPolicyTierRequest instantiates a new AiGatewayPolicyTierRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayPolicyTierRequest(targets []AiGatewayPolicyTargetRequest) *AiGatewayPolicyTierRequest {
	this := AiGatewayPolicyTierRequest{}
	this.Targets = targets
	return &this
}

// NewAiGatewayPolicyTierRequestWithDefaults instantiates a new AiGatewayPolicyTierRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayPolicyTierRequestWithDefaults() *AiGatewayPolicyTierRequest {
	this := AiGatewayPolicyTierRequest{}
	return &this
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *AiGatewayPolicyTierRequest) GetPriority() int64 {
	if o == nil || o.Priority == nil {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayPolicyTierRequest) GetPriorityOk() (*int64, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *AiGatewayPolicyTierRequest) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *AiGatewayPolicyTierRequest) SetPriority(v int64) {
	o.Priority = &v
}

// GetTargets returns the Targets field value.
func (o *AiGatewayPolicyTierRequest) GetTargets() []AiGatewayPolicyTargetRequest {
	if o == nil {
		var ret []AiGatewayPolicyTargetRequest
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *AiGatewayPolicyTierRequest) GetTargetsOk() (*[]AiGatewayPolicyTargetRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Targets, true
}

// SetTargets sets field value.
func (o *AiGatewayPolicyTierRequest) SetTargets(v []AiGatewayPolicyTargetRequest) {
	o.Targets = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayPolicyTierRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	toSerialize["targets"] = o.Targets

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayPolicyTierRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Priority *int64                          `json:"priority,omitempty"`
		Targets  *[]AiGatewayPolicyTargetRequest `json:"targets"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Targets == nil {
		return fmt.Errorf("required field targets missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"priority", "targets"})
	} else {
		return err
	}
	o.Priority = all.Priority
	o.Targets = *all.Targets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
