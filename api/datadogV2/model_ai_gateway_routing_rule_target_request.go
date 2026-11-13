// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRoutingRuleTargetRequest Model target or reference to another routing rule used to create or replace a rule.
type AiGatewayRoutingRuleTargetRequest struct {
	// Whether this is the single default classified destination. Omission preserves a saved value; null clears it.
	Default datadog.NullableBool `json:"default,omitempty"`
	// Saved target row ID. On replacement, retain this ID only within its original route and at most once. A stale, foreign, duplicate, or destination-changing ID is rejected. Hidden route targets can be retained using only this parent-local ID and type.
	Id *int64 `json:"id,omitempty"`
	// Model name. Omit when retaining a saved target by ID or selecting a concrete target by model_id.
	Model *string `json:"model,omitempty"`
	// Opaque model catalog ID for a concrete model target. New selections require provider View; saved selections can be retained without metadata.
	ModelId *int64 `json:"model_id,omitempty"`
	// Provider slug. weighted_affinity requires an explicit provider and does not allow "*".
	Provider *string `json:"provider,omitempty"`
	// Optional reasoning effort preference for this model destination. Must be supported by the selected model and provider. For provider "*", every eligible provider must support the level. An explicit value overrides the request's reasoning setting, subject to model-group policy. Omitted or null preserves request and provider behavior; none is an explicit level, not the default. On update, omission preserves a saved target override; null explicitly clears it. Routing-rule references cannot set this field; each referenced model destination owns its override. If an override cannot be applied at runtime, the gateway keeps the destination and restores the caller's original reasoning settings, or preserves the provider default when omitted. It rechecks all applicable route and model-group policies against those settings before sending. Model-group max_reasoning_level enforces hard limits; the override is not an additional cap. Normal fallback applies if governance denies the resulting request or the destination cannot otherwise serve it. Metrics and span tags record the override outcome and incompatibility reason. When response headers have not already been committed, `X-Ai-Gateway-Reasoning-Override` reports applied or not_applied for the returned attempt, and `X-Ai-Gateway-Reasoning-Override-Reason` reports the reason when not applied.
	ReasoningLevel NullableAiGatewayRoutingReasoningLevel `json:"reasoning_level,omitempty"`
	// Destination rubric in a classified tier. Omission preserves a saved value; null clears it. Must be absent on the default destination.
	Rubric datadog.NullableString `json:"rubric,omitempty"`
	// Organization-unique routing rule slug. New references require route View. Omit to retain a saved target by its parent-local ID. Allowed only in configured tiers.
	Slug *string `json:"slug,omitempty"`
	// Type of routing target.
	Type AiGatewayRoutingRuleTargetType `json:"type"`
	// Relative weight in a weighted_affinity tier. Omission preserves a saved target weight; new targets and explicit null default to 1. Weights need not total 100. Must be omitted or null in configured tiers.
	Weight datadog.NullableInt64 `json:"weight,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayRoutingRuleTargetRequest instantiates a new AiGatewayRoutingRuleTargetRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayRoutingRuleTargetRequest(typeVar AiGatewayRoutingRuleTargetType) *AiGatewayRoutingRuleTargetRequest {
	this := AiGatewayRoutingRuleTargetRequest{}
	this.Type = typeVar
	return &this
}

// NewAiGatewayRoutingRuleTargetRequestWithDefaults instantiates a new AiGatewayRoutingRuleTargetRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayRoutingRuleTargetRequestWithDefaults() *AiGatewayRoutingRuleTargetRequest {
	this := AiGatewayRoutingRuleTargetRequest{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRoutingRuleTargetRequest) GetDefault() bool {
	if o == nil || o.Default.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Default.Get()
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRoutingRuleTargetRequest) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Default.Get(), o.Default.IsSet()
}

// HasDefault returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTargetRequest) HasDefault() bool {
	return o != nil && o.Default.IsSet()
}

// SetDefault gets a reference to the given datadog.NullableBool and assigns it to the Default field.
func (o *AiGatewayRoutingRuleTargetRequest) SetDefault(v bool) {
	o.Default.Set(&v)
}

// SetDefaultNil sets the value for Default to be an explicit nil.
func (o *AiGatewayRoutingRuleTargetRequest) SetDefaultNil() {
	o.Default.Set(nil)
}

// UnsetDefault ensures that no value is present for Default, not even an explicit nil.
func (o *AiGatewayRoutingRuleTargetRequest) UnsetDefault() {
	o.Default.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleTargetRequest) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleTargetRequest) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTargetRequest) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AiGatewayRoutingRuleTargetRequest) SetId(v int64) {
	o.Id = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleTargetRequest) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleTargetRequest) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTargetRequest) HasModel() bool {
	return o != nil && o.Model != nil
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *AiGatewayRoutingRuleTargetRequest) SetModel(v string) {
	o.Model = &v
}

// GetModelId returns the ModelId field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleTargetRequest) GetModelId() int64 {
	if o == nil || o.ModelId == nil {
		var ret int64
		return ret
	}
	return *o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleTargetRequest) GetModelIdOk() (*int64, bool) {
	if o == nil || o.ModelId == nil {
		return nil, false
	}
	return o.ModelId, true
}

// HasModelId returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTargetRequest) HasModelId() bool {
	return o != nil && o.ModelId != nil
}

// SetModelId gets a reference to the given int64 and assigns it to the ModelId field.
func (o *AiGatewayRoutingRuleTargetRequest) SetModelId(v int64) {
	o.ModelId = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleTargetRequest) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleTargetRequest) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTargetRequest) HasProvider() bool {
	return o != nil && o.Provider != nil
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *AiGatewayRoutingRuleTargetRequest) SetProvider(v string) {
	o.Provider = &v
}

// GetReasoningLevel returns the ReasoningLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRoutingRuleTargetRequest) GetReasoningLevel() AiGatewayRoutingReasoningLevel {
	if o == nil || o.ReasoningLevel.Get() == nil {
		var ret AiGatewayRoutingReasoningLevel
		return ret
	}
	return *o.ReasoningLevel.Get()
}

// GetReasoningLevelOk returns a tuple with the ReasoningLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRoutingRuleTargetRequest) GetReasoningLevelOk() (*AiGatewayRoutingReasoningLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReasoningLevel.Get(), o.ReasoningLevel.IsSet()
}

// HasReasoningLevel returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTargetRequest) HasReasoningLevel() bool {
	return o != nil && o.ReasoningLevel.IsSet()
}

// SetReasoningLevel gets a reference to the given NullableAiGatewayRoutingReasoningLevel and assigns it to the ReasoningLevel field.
func (o *AiGatewayRoutingRuleTargetRequest) SetReasoningLevel(v AiGatewayRoutingReasoningLevel) {
	o.ReasoningLevel.Set(&v)
}

// SetReasoningLevelNil sets the value for ReasoningLevel to be an explicit nil.
func (o *AiGatewayRoutingRuleTargetRequest) SetReasoningLevelNil() {
	o.ReasoningLevel.Set(nil)
}

// UnsetReasoningLevel ensures that no value is present for ReasoningLevel, not even an explicit nil.
func (o *AiGatewayRoutingRuleTargetRequest) UnsetReasoningLevel() {
	o.ReasoningLevel.Unset()
}

// GetRubric returns the Rubric field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRoutingRuleTargetRequest) GetRubric() string {
	if o == nil || o.Rubric.Get() == nil {
		var ret string
		return ret
	}
	return *o.Rubric.Get()
}

// GetRubricOk returns a tuple with the Rubric field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRoutingRuleTargetRequest) GetRubricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rubric.Get(), o.Rubric.IsSet()
}

// HasRubric returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTargetRequest) HasRubric() bool {
	return o != nil && o.Rubric.IsSet()
}

// SetRubric gets a reference to the given datadog.NullableString and assigns it to the Rubric field.
func (o *AiGatewayRoutingRuleTargetRequest) SetRubric(v string) {
	o.Rubric.Set(&v)
}

// SetRubricNil sets the value for Rubric to be an explicit nil.
func (o *AiGatewayRoutingRuleTargetRequest) SetRubricNil() {
	o.Rubric.Set(nil)
}

// UnsetRubric ensures that no value is present for Rubric, not even an explicit nil.
func (o *AiGatewayRoutingRuleTargetRequest) UnsetRubric() {
	o.Rubric.Unset()
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleTargetRequest) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleTargetRequest) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTargetRequest) HasSlug() bool {
	return o != nil && o.Slug != nil
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *AiGatewayRoutingRuleTargetRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetType returns the Type field value.
func (o *AiGatewayRoutingRuleTargetRequest) GetType() AiGatewayRoutingRuleTargetType {
	if o == nil {
		var ret AiGatewayRoutingRuleTargetType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleTargetRequest) GetTypeOk() (*AiGatewayRoutingRuleTargetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AiGatewayRoutingRuleTargetRequest) SetType(v AiGatewayRoutingRuleTargetType) {
	o.Type = v
}

// GetWeight returns the Weight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRoutingRuleTargetRequest) GetWeight() int64 {
	if o == nil || o.Weight.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Weight.Get()
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRoutingRuleTargetRequest) GetWeightOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weight.Get(), o.Weight.IsSet()
}

// HasWeight returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleTargetRequest) HasWeight() bool {
	return o != nil && o.Weight.IsSet()
}

// SetWeight gets a reference to the given datadog.NullableInt64 and assigns it to the Weight field.
func (o *AiGatewayRoutingRuleTargetRequest) SetWeight(v int64) {
	o.Weight.Set(&v)
}

// SetWeightNil sets the value for Weight to be an explicit nil.
func (o *AiGatewayRoutingRuleTargetRequest) SetWeightNil() {
	o.Weight.Set(nil)
}

// UnsetWeight ensures that no value is present for Weight, not even an explicit nil.
func (o *AiGatewayRoutingRuleTargetRequest) UnsetWeight() {
	o.Weight.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRoutingRuleTargetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Default.IsSet() {
		toSerialize["default"] = o.Default.Get()
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
	if o.Rubric.IsSet() {
		toSerialize["rubric"] = o.Rubric.Get()
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	toSerialize["type"] = o.Type
	if o.Weight.IsSet() {
		toSerialize["weight"] = o.Weight.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayRoutingRuleTargetRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Default        datadog.NullableBool                   `json:"default,omitempty"`
		Id             *int64                                 `json:"id,omitempty"`
		Model          *string                                `json:"model,omitempty"`
		ModelId        *int64                                 `json:"model_id,omitempty"`
		Provider       *string                                `json:"provider,omitempty"`
		ReasoningLevel NullableAiGatewayRoutingReasoningLevel `json:"reasoning_level,omitempty"`
		Rubric         datadog.NullableString                 `json:"rubric,omitempty"`
		Slug           *string                                `json:"slug,omitempty"`
		Type           *AiGatewayRoutingRuleTargetType        `json:"type"`
		Weight         datadog.NullableInt64                  `json:"weight,omitempty"`
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
