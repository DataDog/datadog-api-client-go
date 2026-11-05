// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelReasoningControl One supported mechanism for configuring model reasoning.
type AiGatewayModelReasoningControl struct {
	// Supported effort values in provider order for an effort control.
	EffortLevels []*string `json:"effort_levels,omitempty"`
	// Maximum reasoning-token budget.
	MaxBudgetTokens *int64 `json:"max_budget_tokens,omitempty"`
	// Minimum reasoning-token budget.
	MinBudgetTokens *int64 `json:"min_budget_tokens,omitempty"`
	// Reasoning-control mechanism.
	Type AiGatewayModelReasoningControlType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelReasoningControl instantiates a new AiGatewayModelReasoningControl object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelReasoningControl(typeVar AiGatewayModelReasoningControlType) *AiGatewayModelReasoningControl {
	this := AiGatewayModelReasoningControl{}
	this.Type = typeVar
	return &this
}

// NewAiGatewayModelReasoningControlWithDefaults instantiates a new AiGatewayModelReasoningControl object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelReasoningControlWithDefaults() *AiGatewayModelReasoningControl {
	this := AiGatewayModelReasoningControl{}
	return &this
}

// GetEffortLevels returns the EffortLevels field value if set, zero value otherwise.
func (o *AiGatewayModelReasoningControl) GetEffortLevels() []*string {
	if o == nil || o.EffortLevels == nil {
		var ret []*string
		return ret
	}
	return o.EffortLevels
}

// GetEffortLevelsOk returns a tuple with the EffortLevels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelReasoningControl) GetEffortLevelsOk() (*[]*string, bool) {
	if o == nil || o.EffortLevels == nil {
		return nil, false
	}
	return &o.EffortLevels, true
}

// HasEffortLevels returns a boolean if a field has been set.
func (o *AiGatewayModelReasoningControl) HasEffortLevels() bool {
	return o != nil && o.EffortLevels != nil
}

// SetEffortLevels gets a reference to the given []*string and assigns it to the EffortLevels field.
func (o *AiGatewayModelReasoningControl) SetEffortLevels(v []*string) {
	o.EffortLevels = v
}

// GetMaxBudgetTokens returns the MaxBudgetTokens field value if set, zero value otherwise.
func (o *AiGatewayModelReasoningControl) GetMaxBudgetTokens() int64 {
	if o == nil || o.MaxBudgetTokens == nil {
		var ret int64
		return ret
	}
	return *o.MaxBudgetTokens
}

// GetMaxBudgetTokensOk returns a tuple with the MaxBudgetTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelReasoningControl) GetMaxBudgetTokensOk() (*int64, bool) {
	if o == nil || o.MaxBudgetTokens == nil {
		return nil, false
	}
	return o.MaxBudgetTokens, true
}

// HasMaxBudgetTokens returns a boolean if a field has been set.
func (o *AiGatewayModelReasoningControl) HasMaxBudgetTokens() bool {
	return o != nil && o.MaxBudgetTokens != nil
}

// SetMaxBudgetTokens gets a reference to the given int64 and assigns it to the MaxBudgetTokens field.
func (o *AiGatewayModelReasoningControl) SetMaxBudgetTokens(v int64) {
	o.MaxBudgetTokens = &v
}

// GetMinBudgetTokens returns the MinBudgetTokens field value if set, zero value otherwise.
func (o *AiGatewayModelReasoningControl) GetMinBudgetTokens() int64 {
	if o == nil || o.MinBudgetTokens == nil {
		var ret int64
		return ret
	}
	return *o.MinBudgetTokens
}

// GetMinBudgetTokensOk returns a tuple with the MinBudgetTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelReasoningControl) GetMinBudgetTokensOk() (*int64, bool) {
	if o == nil || o.MinBudgetTokens == nil {
		return nil, false
	}
	return o.MinBudgetTokens, true
}

// HasMinBudgetTokens returns a boolean if a field has been set.
func (o *AiGatewayModelReasoningControl) HasMinBudgetTokens() bool {
	return o != nil && o.MinBudgetTokens != nil
}

// SetMinBudgetTokens gets a reference to the given int64 and assigns it to the MinBudgetTokens field.
func (o *AiGatewayModelReasoningControl) SetMinBudgetTokens(v int64) {
	o.MinBudgetTokens = &v
}

// GetType returns the Type field value.
func (o *AiGatewayModelReasoningControl) GetType() AiGatewayModelReasoningControlType {
	if o == nil {
		var ret AiGatewayModelReasoningControlType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelReasoningControl) GetTypeOk() (*AiGatewayModelReasoningControlType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AiGatewayModelReasoningControl) SetType(v AiGatewayModelReasoningControlType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelReasoningControl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EffortLevels != nil {
		toSerialize["effort_levels"] = o.EffortLevels
	}
	if o.MaxBudgetTokens != nil {
		toSerialize["max_budget_tokens"] = o.MaxBudgetTokens
	}
	if o.MinBudgetTokens != nil {
		toSerialize["min_budget_tokens"] = o.MinBudgetTokens
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelReasoningControl) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EffortLevels    []*string                           `json:"effort_levels,omitempty"`
		MaxBudgetTokens *int64                              `json:"max_budget_tokens,omitempty"`
		MinBudgetTokens *int64                              `json:"min_budget_tokens,omitempty"`
		Type            *AiGatewayModelReasoningControlType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"effort_levels", "max_budget_tokens", "min_budget_tokens", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EffortLevels = all.EffortLevels
	o.MaxBudgetTokens = all.MaxBudgetTokens
	o.MinBudgetTokens = all.MinBudgetTokens
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
