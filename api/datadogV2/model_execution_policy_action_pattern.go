// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionPolicyActionPattern The set of actions this policy applies to.
type ExecutionPolicyActionPattern struct {
	// The fully qualified action names this policy matches. Use `*` to match all actions
	// of the integration, or a fully qualified name prefixed with the integration's action
	// namespace (for example `com.datadoghq.script.*` for the Script integration).
	ActionFqns []string `json:"action_fqns"`
	// The integration the action pattern applies to.
	Integration ExecutionPolicyIntegration `json:"integration"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExecutionPolicyActionPattern instantiates a new ExecutionPolicyActionPattern object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExecutionPolicyActionPattern(actionFqns []string, integration ExecutionPolicyIntegration) *ExecutionPolicyActionPattern {
	this := ExecutionPolicyActionPattern{}
	this.ActionFqns = actionFqns
	this.Integration = integration
	return &this
}

// NewExecutionPolicyActionPatternWithDefaults instantiates a new ExecutionPolicyActionPattern object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExecutionPolicyActionPatternWithDefaults() *ExecutionPolicyActionPattern {
	this := ExecutionPolicyActionPattern{}
	return &this
}

// GetActionFqns returns the ActionFqns field value.
func (o *ExecutionPolicyActionPattern) GetActionFqns() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ActionFqns
}

// GetActionFqnsOk returns a tuple with the ActionFqns field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyActionPattern) GetActionFqnsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionFqns, true
}

// SetActionFqns sets field value.
func (o *ExecutionPolicyActionPattern) SetActionFqns(v []string) {
	o.ActionFqns = v
}

// GetIntegration returns the Integration field value.
func (o *ExecutionPolicyActionPattern) GetIntegration() ExecutionPolicyIntegration {
	if o == nil {
		var ret ExecutionPolicyIntegration
		return ret
	}
	return o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyActionPattern) GetIntegrationOk() (*ExecutionPolicyIntegration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Integration, true
}

// SetIntegration sets field value.
func (o *ExecutionPolicyActionPattern) SetIntegration(v ExecutionPolicyIntegration) {
	o.Integration = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExecutionPolicyActionPattern) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action_fqns"] = o.ActionFqns
	toSerialize["integration"] = o.Integration

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExecutionPolicyActionPattern) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActionFqns  *[]string                   `json:"action_fqns"`
		Integration *ExecutionPolicyIntegration `json:"integration"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ActionFqns == nil {
		return fmt.Errorf("required field action_fqns missing")
	}
	if all.Integration == nil {
		return fmt.Errorf("required field integration missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action_fqns", "integration"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ActionFqns = *all.ActionFqns
	if !all.Integration.IsValid() {
		hasInvalidField = true
	} else {
		o.Integration = *all.Integration
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
