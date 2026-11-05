// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelGroupCreateAttributes Attributes used to create an AI Gateway model group.
type AiGatewayModelGroupCreateAttributes struct {
	// Whether fast mode is allowed for models and routing rules in the group.
	AllowFastMode *bool `json:"allow_fast_mode,omitempty"`
	// Optional initial model-group access bindings.
	Bindings []AiGatewayModelGroupBinding `json:"bindings,omitempty"`
	// Maximum reasoning level permitted by a model group.
	MaxReasoningLevel *AiGatewayMaxReasoningLevel `json:"max_reasoning_level,omitempty"`
	// Model catalog IDs to select. New selections require provider View.
	ModelIds []int64 `json:"model_ids,omitempty"`
	// Human-readable resource name.
	Name string `json:"name"`
	// Explicit provider selections, including current and future models. Omit to preserve selections on update; an empty array clears them. Null is invalid. New selections require provider View; saved hidden IDs can be retained or removed.
	ProviderIds []int64 `json:"provider_ids,omitempty"`
	// Organization-unique routing rule slugs, such as coding-route.
	RoutingRuleSlugs []string `json:"routing_rule_slugs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelGroupCreateAttributes instantiates a new AiGatewayModelGroupCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelGroupCreateAttributes(name string) *AiGatewayModelGroupCreateAttributes {
	this := AiGatewayModelGroupCreateAttributes{}
	var allowFastMode bool = true
	this.AllowFastMode = &allowFastMode
	this.Name = name
	return &this
}

// NewAiGatewayModelGroupCreateAttributesWithDefaults instantiates a new AiGatewayModelGroupCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelGroupCreateAttributesWithDefaults() *AiGatewayModelGroupCreateAttributes {
	this := AiGatewayModelGroupCreateAttributes{}
	var allowFastMode bool = true
	this.AllowFastMode = &allowFastMode
	return &this
}

// GetAllowFastMode returns the AllowFastMode field value if set, zero value otherwise.
func (o *AiGatewayModelGroupCreateAttributes) GetAllowFastMode() bool {
	if o == nil || o.AllowFastMode == nil {
		var ret bool
		return ret
	}
	return *o.AllowFastMode
}

// GetAllowFastModeOk returns a tuple with the AllowFastMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupCreateAttributes) GetAllowFastModeOk() (*bool, bool) {
	if o == nil || o.AllowFastMode == nil {
		return nil, false
	}
	return o.AllowFastMode, true
}

// HasAllowFastMode returns a boolean if a field has been set.
func (o *AiGatewayModelGroupCreateAttributes) HasAllowFastMode() bool {
	return o != nil && o.AllowFastMode != nil
}

// SetAllowFastMode gets a reference to the given bool and assigns it to the AllowFastMode field.
func (o *AiGatewayModelGroupCreateAttributes) SetAllowFastMode(v bool) {
	o.AllowFastMode = &v
}

// GetBindings returns the Bindings field value if set, zero value otherwise.
func (o *AiGatewayModelGroupCreateAttributes) GetBindings() []AiGatewayModelGroupBinding {
	if o == nil || o.Bindings == nil {
		var ret []AiGatewayModelGroupBinding
		return ret
	}
	return o.Bindings
}

// GetBindingsOk returns a tuple with the Bindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupCreateAttributes) GetBindingsOk() (*[]AiGatewayModelGroupBinding, bool) {
	if o == nil || o.Bindings == nil {
		return nil, false
	}
	return &o.Bindings, true
}

// HasBindings returns a boolean if a field has been set.
func (o *AiGatewayModelGroupCreateAttributes) HasBindings() bool {
	return o != nil && o.Bindings != nil
}

// SetBindings gets a reference to the given []AiGatewayModelGroupBinding and assigns it to the Bindings field.
func (o *AiGatewayModelGroupCreateAttributes) SetBindings(v []AiGatewayModelGroupBinding) {
	o.Bindings = v
}

// GetMaxReasoningLevel returns the MaxReasoningLevel field value if set, zero value otherwise.
func (o *AiGatewayModelGroupCreateAttributes) GetMaxReasoningLevel() AiGatewayMaxReasoningLevel {
	if o == nil || o.MaxReasoningLevel == nil {
		var ret AiGatewayMaxReasoningLevel
		return ret
	}
	return *o.MaxReasoningLevel
}

// GetMaxReasoningLevelOk returns a tuple with the MaxReasoningLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupCreateAttributes) GetMaxReasoningLevelOk() (*AiGatewayMaxReasoningLevel, bool) {
	if o == nil || o.MaxReasoningLevel == nil {
		return nil, false
	}
	return o.MaxReasoningLevel, true
}

// HasMaxReasoningLevel returns a boolean if a field has been set.
func (o *AiGatewayModelGroupCreateAttributes) HasMaxReasoningLevel() bool {
	return o != nil && o.MaxReasoningLevel != nil
}

// SetMaxReasoningLevel gets a reference to the given AiGatewayMaxReasoningLevel and assigns it to the MaxReasoningLevel field.
func (o *AiGatewayModelGroupCreateAttributes) SetMaxReasoningLevel(v AiGatewayMaxReasoningLevel) {
	o.MaxReasoningLevel = &v
}

// GetModelIds returns the ModelIds field value if set, zero value otherwise.
func (o *AiGatewayModelGroupCreateAttributes) GetModelIds() []int64 {
	if o == nil || o.ModelIds == nil {
		var ret []int64
		return ret
	}
	return o.ModelIds
}

// GetModelIdsOk returns a tuple with the ModelIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupCreateAttributes) GetModelIdsOk() (*[]int64, bool) {
	if o == nil || o.ModelIds == nil {
		return nil, false
	}
	return &o.ModelIds, true
}

// HasModelIds returns a boolean if a field has been set.
func (o *AiGatewayModelGroupCreateAttributes) HasModelIds() bool {
	return o != nil && o.ModelIds != nil
}

// SetModelIds gets a reference to the given []int64 and assigns it to the ModelIds field.
func (o *AiGatewayModelGroupCreateAttributes) SetModelIds(v []int64) {
	o.ModelIds = v
}

// GetName returns the Name field value.
func (o *AiGatewayModelGroupCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AiGatewayModelGroupCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetProviderIds returns the ProviderIds field value if set, zero value otherwise.
func (o *AiGatewayModelGroupCreateAttributes) GetProviderIds() []int64 {
	if o == nil || o.ProviderIds == nil {
		var ret []int64
		return ret
	}
	return o.ProviderIds
}

// GetProviderIdsOk returns a tuple with the ProviderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupCreateAttributes) GetProviderIdsOk() (*[]int64, bool) {
	if o == nil || o.ProviderIds == nil {
		return nil, false
	}
	return &o.ProviderIds, true
}

// HasProviderIds returns a boolean if a field has been set.
func (o *AiGatewayModelGroupCreateAttributes) HasProviderIds() bool {
	return o != nil && o.ProviderIds != nil
}

// SetProviderIds gets a reference to the given []int64 and assigns it to the ProviderIds field.
func (o *AiGatewayModelGroupCreateAttributes) SetProviderIds(v []int64) {
	o.ProviderIds = v
}

// GetRoutingRuleSlugs returns the RoutingRuleSlugs field value if set, zero value otherwise.
func (o *AiGatewayModelGroupCreateAttributes) GetRoutingRuleSlugs() []string {
	if o == nil || o.RoutingRuleSlugs == nil {
		var ret []string
		return ret
	}
	return o.RoutingRuleSlugs
}

// GetRoutingRuleSlugsOk returns a tuple with the RoutingRuleSlugs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupCreateAttributes) GetRoutingRuleSlugsOk() (*[]string, bool) {
	if o == nil || o.RoutingRuleSlugs == nil {
		return nil, false
	}
	return &o.RoutingRuleSlugs, true
}

// HasRoutingRuleSlugs returns a boolean if a field has been set.
func (o *AiGatewayModelGroupCreateAttributes) HasRoutingRuleSlugs() bool {
	return o != nil && o.RoutingRuleSlugs != nil
}

// SetRoutingRuleSlugs gets a reference to the given []string and assigns it to the RoutingRuleSlugs field.
func (o *AiGatewayModelGroupCreateAttributes) SetRoutingRuleSlugs(v []string) {
	o.RoutingRuleSlugs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelGroupCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowFastMode != nil {
		toSerialize["allow_fast_mode"] = o.AllowFastMode
	}
	if o.Bindings != nil {
		toSerialize["bindings"] = o.Bindings
	}
	if o.MaxReasoningLevel != nil {
		toSerialize["max_reasoning_level"] = o.MaxReasoningLevel
	}
	if o.ModelIds != nil {
		toSerialize["model_ids"] = o.ModelIds
	}
	toSerialize["name"] = o.Name
	if o.ProviderIds != nil {
		toSerialize["provider_ids"] = o.ProviderIds
	}
	if o.RoutingRuleSlugs != nil {
		toSerialize["routing_rule_slugs"] = o.RoutingRuleSlugs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelGroupCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowFastMode     *bool                        `json:"allow_fast_mode,omitempty"`
		Bindings          []AiGatewayModelGroupBinding `json:"bindings,omitempty"`
		MaxReasoningLevel *AiGatewayMaxReasoningLevel  `json:"max_reasoning_level,omitempty"`
		ModelIds          []int64                      `json:"model_ids,omitempty"`
		Name              *string                      `json:"name"`
		ProviderIds       []int64                      `json:"provider_ids,omitempty"`
		RoutingRuleSlugs  []string                     `json:"routing_rule_slugs,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allow_fast_mode", "bindings", "max_reasoning_level", "model_ids", "name", "provider_ids", "routing_rule_slugs"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AllowFastMode = all.AllowFastMode
	o.Bindings = all.Bindings
	if all.MaxReasoningLevel != nil && !all.MaxReasoningLevel.IsValid() {
		hasInvalidField = true
	} else {
		o.MaxReasoningLevel = all.MaxReasoningLevel
	}
	o.ModelIds = all.ModelIds
	o.Name = *all.Name
	o.ProviderIds = all.ProviderIds
	o.RoutingRuleSlugs = all.RoutingRuleSlugs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
