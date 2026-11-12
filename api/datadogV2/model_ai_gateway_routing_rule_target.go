// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRoutingRuleTarget Model target or reference to another routing rule.
type AiGatewayRoutingRuleTarget struct {
	// Present and true for the default classified destination.
	Default *bool `json:"default,omitempty"`
	// Saved target row ID. On replacement, retain this ID only within its original route and at most once. A stale or foreign ID is rejected.
	Id *int64 `json:"id,omitempty"`
	// Model name or pattern. Omitted when the caller cannot view its provider.
	Model *string `json:"model,omitempty"`
	// Opaque model catalog ID for a concrete model target. New selections require provider View; saved selections can be retained without metadata.
	ModelId *int64 `json:"model_id,omitempty"`
	// Provider slug. Omitted for dynamic resolution and when the caller cannot view its provider.
	Provider *string `json:"provider,omitempty"`
	// Optional reasoning effort preference for this model destination. Must be supported by the selected model and provider. For provider "*", every eligible provider must support the level. An explicit value overrides the request's reasoning setting, subject to model-group policy. Omitted or null preserves request and provider behavior; none is an explicit level, not the default. On update, omission preserves a saved target override; null explicitly clears it. Routing-rule references cannot set this field; each referenced model destination owns its override. If an override cannot be applied at runtime, the gateway keeps the destination and restores the caller's original reasoning settings, or preserves the provider default when omitted. It rechecks all applicable route and model-group policies against those settings before sending. Model-group max_reasoning_level enforces hard limits; the override is not an additional cap. Normal fallback applies if governance denies the resulting request or the destination cannot otherwise serve it. Metrics and span tags record the override outcome and incompatibility reason. When response headers have not already been committed, `X-Ai-Gateway-Reasoning-Override` reports applied or not_applied for the returned attempt, and `X-Ai-Gateway-Reasoning-Override-Reason` reports the reason when not applied.
	ReasoningLevel NullableAiGatewayRoutingReasoningLevel `json:"reasoning_level,omitempty"`
	// Destination rubric, present only for non-default targets in classified tiers.
	Rubric *string `json:"rubric,omitempty"`
	// Referenced route slug. Omitted when the caller lacks View; only the parent-local occurrence ID remains.
	Slug *string `json:"slug,omitempty"`
	// Type of routing target.
	Type AiGatewayRoutingRuleTargetType `json:"type"`
	// Effective relative weight. Returned only for targets in weighted_affinity tiers.
	Weight *int64 `json:"weight,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayRoutingRuleTarget instantiates a new AiGatewayRoutingRuleTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayRoutingRuleTarget(typeVar AiGatewayRoutingRuleTargetType) *AiGatewayRoutingRuleTarget {
	this := AiGatewayRoutingRuleTarget{}
	this.Type = typeVar
	return &this
}

// NewAiGatewayRoutingRuleTargetWithDefaults instantiates a new AiGatewayRoutingRuleTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayRoutingRuleTargetWithDefaults() *AiGatewayRoutingRuleTarget {
	this := AiGatewayRoutingRuleTarget{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleTarget) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleTarget) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTarget) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *AiGatewayRoutingRuleTarget) SetDefault(v bool) {
	o.Default = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleTarget) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleTarget) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTarget) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AiGatewayRoutingRuleTarget) SetId(v int64) {
	o.Id = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleTarget) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleTarget) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTarget) HasModel() bool {
	return o != nil && o.Model != nil
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *AiGatewayRoutingRuleTarget) SetModel(v string) {
	o.Model = &v
}

// GetModelId returns the ModelId field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleTarget) GetModelId() int64 {
	if o == nil || o.ModelId == nil {
		var ret int64
		return ret
	}
	return *o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleTarget) GetModelIdOk() (*int64, bool) {
	if o == nil || o.ModelId == nil {
		return nil, false
	}
	return o.ModelId, true
}

// HasModelId returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTarget) HasModelId() bool {
	return o != nil && o.ModelId != nil
}

// SetModelId gets a reference to the given int64 and assigns it to the ModelId field.
func (o *AiGatewayRoutingRuleTarget) SetModelId(v int64) {
	o.ModelId = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleTarget) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleTarget) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTarget) HasProvider() bool {
	return o != nil && o.Provider != nil
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *AiGatewayRoutingRuleTarget) SetProvider(v string) {
	o.Provider = &v
}

// GetReasoningLevel returns the ReasoningLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRoutingRuleTarget) GetReasoningLevel() AiGatewayRoutingReasoningLevel {
	if o == nil || o.ReasoningLevel.Get() == nil {
		var ret AiGatewayRoutingReasoningLevel
		return ret
	}
	return *o.ReasoningLevel.Get()
}

// GetReasoningLevelOk returns a tuple with the ReasoningLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRoutingRuleTarget) GetReasoningLevelOk() (*AiGatewayRoutingReasoningLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReasoningLevel.Get(), o.ReasoningLevel.IsSet()
}

// HasReasoningLevel returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTarget) HasReasoningLevel() bool {
	return o != nil && o.ReasoningLevel.IsSet()
}

// SetReasoningLevel gets a reference to the given NullableAiGatewayRoutingReasoningLevel and assigns it to the ReasoningLevel field.
func (o *AiGatewayRoutingRuleTarget) SetReasoningLevel(v AiGatewayRoutingReasoningLevel) {
	o.ReasoningLevel.Set(&v)
}

// SetReasoningLevelNil sets the value for ReasoningLevel to be an explicit nil.
func (o *AiGatewayRoutingRuleTarget) SetReasoningLevelNil() {
	o.ReasoningLevel.Set(nil)
}

// UnsetReasoningLevel ensures that no value is present for ReasoningLevel, not even an explicit nil.
func (o *AiGatewayRoutingRuleTarget) UnsetReasoningLevel() {
	o.ReasoningLevel.Unset()
}

// GetRubric returns the Rubric field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleTarget) GetRubric() string {
	if o == nil || o.Rubric == nil {
		var ret string
		return ret
	}
	return *o.Rubric
}

// GetRubricOk returns a tuple with the Rubric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleTarget) GetRubricOk() (*string, bool) {
	if o == nil || o.Rubric == nil {
		return nil, false
	}
	return o.Rubric, true
}

// HasRubric returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTarget) HasRubric() bool {
	return o != nil && o.Rubric != nil
}

// SetRubric gets a reference to the given string and assigns it to the Rubric field.
func (o *AiGatewayRoutingRuleTarget) SetRubric(v string) {
	o.Rubric = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleTarget) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleTarget) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTarget) HasSlug() bool {
	return o != nil && o.Slug != nil
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *AiGatewayRoutingRuleTarget) SetSlug(v string) {
	o.Slug = &v
}

// GetType returns the Type field value.
func (o *AiGatewayRoutingRuleTarget) GetType() AiGatewayRoutingRuleTargetType {
	if o == nil {
		var ret AiGatewayRoutingRuleTargetType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleTarget) GetTypeOk() (*AiGatewayRoutingRuleTargetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AiGatewayRoutingRuleTarget) SetType(v AiGatewayRoutingRuleTargetType) {
	o.Type = v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleTarget) GetWeight() int64 {
	if o == nil || o.Weight == nil {
		var ret int64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleTarget) GetWeightOk() (*int64, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTarget) HasWeight() bool {
	return o != nil && o.Weight != nil
}

// SetWeight gets a reference to the given int64 and assigns it to the Weight field.
func (o *AiGatewayRoutingRuleTarget) SetWeight(v int64) {
	o.Weight = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRoutingRuleTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.ModelId != nil {
		toSerialize["model_id"] = o.ModelId
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.ReasoningLevel.IsSet() {
		toSerialize["reasoning_level"] = o.ReasoningLevel.Get()
	}
	if o.Rubric != nil {
		toSerialize["rubric"] = o.Rubric
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	toSerialize["type"] = o.Type
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayRoutingRuleTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Default        *bool                                  `json:"default,omitempty"`
		Id             *int64                                 `json:"id,omitempty"`
		Model          *string                                `json:"model,omitempty"`
		ModelId        *int64                                 `json:"model_id,omitempty"`
		Provider       *string                                `json:"provider,omitempty"`
		ReasoningLevel NullableAiGatewayRoutingReasoningLevel `json:"reasoning_level,omitempty"`
		Rubric         *string                                `json:"rubric,omitempty"`
		Slug           *string                                `json:"slug,omitempty"`
		Type           *AiGatewayRoutingRuleTargetType        `json:"type"`
		Weight         *int64                                 `json:"weight,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"default", "id", "model", "model_id", "provider", "reasoning_level", "rubric", "slug", "type", "weight"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Default = all.Default
	o.Id = all.Id
	o.Model = all.Model
	o.ModelId = all.ModelId
	o.Provider = all.Provider
	if all.ReasoningLevel.Get() != nil && !all.ReasoningLevel.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.ReasoningLevel = all.ReasoningLevel
	}
	o.Rubric = all.Rubric
	o.Slug = all.Slug
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Weight = all.Weight

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
