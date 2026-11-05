// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelGroupUpdateAttributes Attributes used to update an AI Gateway model group.
type AiGatewayModelGroupUpdateAttributes struct {
	// Whether fast mode is allowed for models and routing rules in the group.
	AllowFastMode *bool `json:"allow_fast_mode,omitempty"`
	// Required when updating the default model group. Must match its current revision; a stale revision returns HTTP 409. Omit for named policies.
	ExpectedRevision *int64 `json:"expected_revision,omitempty"`
	// Maximum reasoning level permitted by a model group.
	MaxReasoningLevel NullableAiGatewayNullableMaxReasoningLevel `json:"max_reasoning_level,omitempty"`
	// Individually selected model catalog IDs. Replaces this selection only. Omit to preserve it; an empty array clears it. Null is invalid. Individual selections do not include future models. New selections require provider View; retaining or removing saved hidden selections requires only parent Edit.
	ModelIds []int64 `json:"model_ids,omitempty"`
	// Human-readable resource name.
	Name *string `json:"name,omitempty"`
	// Explicit provider selections, including current and future models. Omit to preserve selections on update; an empty array clears them. Null is invalid. New selections require provider View; saved hidden IDs can be retained or removed.
	ProviderIds []int64 `json:"provider_ids,omitempty"`
	// Organization-unique routing rule slugs. Replaces route selections only. Omit to preserve them; an empty array clears them. Null is invalid.
	RoutingRuleSlugs []string `json:"routing_rule_slugs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelGroupUpdateAttributes instantiates a new AiGatewayModelGroupUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelGroupUpdateAttributes() *AiGatewayModelGroupUpdateAttributes {
	this := AiGatewayModelGroupUpdateAttributes{}
	return &this
}

// NewAiGatewayModelGroupUpdateAttributesWithDefaults instantiates a new AiGatewayModelGroupUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelGroupUpdateAttributesWithDefaults() *AiGatewayModelGroupUpdateAttributes {
	this := AiGatewayModelGroupUpdateAttributes{}
	return &this
}

// GetAllowFastMode returns the AllowFastMode field value if set, zero value otherwise.
func (o *AiGatewayModelGroupUpdateAttributes) GetAllowFastMode() bool {
	if o == nil || o.AllowFastMode == nil {
		var ret bool
		return ret
	}
	return *o.AllowFastMode
}

// GetAllowFastModeOk returns a tuple with the AllowFastMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupUpdateAttributes) GetAllowFastModeOk() (*bool, bool) {
	if o == nil || o.AllowFastMode == nil {
		return nil, false
	}
	return o.AllowFastMode, true
}

// HasAllowFastMode returns a boolean if a field has been set.
func (o *AiGatewayModelGroupUpdateAttributes) HasAllowFastMode() bool {
	return o != nil && o.AllowFastMode != nil
}

// SetAllowFastMode gets a reference to the given bool and assigns it to the AllowFastMode field.
func (o *AiGatewayModelGroupUpdateAttributes) SetAllowFastMode(v bool) {
	o.AllowFastMode = &v
}

// GetExpectedRevision returns the ExpectedRevision field value if set, zero value otherwise.
func (o *AiGatewayModelGroupUpdateAttributes) GetExpectedRevision() int64 {
	if o == nil || o.ExpectedRevision == nil {
		var ret int64
		return ret
	}
	return *o.ExpectedRevision
}

// GetExpectedRevisionOk returns a tuple with the ExpectedRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupUpdateAttributes) GetExpectedRevisionOk() (*int64, bool) {
	if o == nil || o.ExpectedRevision == nil {
		return nil, false
	}
	return o.ExpectedRevision, true
}

// HasExpectedRevision returns a boolean if a field has been set.
func (o *AiGatewayModelGroupUpdateAttributes) HasExpectedRevision() bool {
	return o != nil && o.ExpectedRevision != nil
}

// SetExpectedRevision gets a reference to the given int64 and assigns it to the ExpectedRevision field.
func (o *AiGatewayModelGroupUpdateAttributes) SetExpectedRevision(v int64) {
	o.ExpectedRevision = &v
}

// GetMaxReasoningLevel returns the MaxReasoningLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayModelGroupUpdateAttributes) GetMaxReasoningLevel() AiGatewayNullableMaxReasoningLevel {
	if o == nil || o.MaxReasoningLevel.Get() == nil {
		var ret AiGatewayNullableMaxReasoningLevel
		return ret
	}
	return *o.MaxReasoningLevel.Get()
}

// GetMaxReasoningLevelOk returns a tuple with the MaxReasoningLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayModelGroupUpdateAttributes) GetMaxReasoningLevelOk() (*AiGatewayNullableMaxReasoningLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxReasoningLevel.Get(), o.MaxReasoningLevel.IsSet()
}

// HasMaxReasoningLevel returns a boolean if a field has been set.
func (o *AiGatewayModelGroupUpdateAttributes) HasMaxReasoningLevel() bool {
	return o != nil && o.MaxReasoningLevel.IsSet()
}

// SetMaxReasoningLevel gets a reference to the given NullableAiGatewayNullableMaxReasoningLevel and assigns it to the MaxReasoningLevel field.
func (o *AiGatewayModelGroupUpdateAttributes) SetMaxReasoningLevel(v AiGatewayNullableMaxReasoningLevel) {
	o.MaxReasoningLevel.Set(&v)
}

// SetMaxReasoningLevelNil sets the value for MaxReasoningLevel to be an explicit nil.
func (o *AiGatewayModelGroupUpdateAttributes) SetMaxReasoningLevelNil() {
	o.MaxReasoningLevel.Set(nil)
}

// UnsetMaxReasoningLevel ensures that no value is present for MaxReasoningLevel, not even an explicit nil.
func (o *AiGatewayModelGroupUpdateAttributes) UnsetMaxReasoningLevel() {
	o.MaxReasoningLevel.Unset()
}

// GetModelIds returns the ModelIds field value if set, zero value otherwise.
func (o *AiGatewayModelGroupUpdateAttributes) GetModelIds() []int64 {
	if o == nil || o.ModelIds == nil {
		var ret []int64
		return ret
	}
	return o.ModelIds
}

// GetModelIdsOk returns a tuple with the ModelIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupUpdateAttributes) GetModelIdsOk() (*[]int64, bool) {
	if o == nil || o.ModelIds == nil {
		return nil, false
	}
	return &o.ModelIds, true
}

// HasModelIds returns a boolean if a field has been set.
func (o *AiGatewayModelGroupUpdateAttributes) HasModelIds() bool {
	return o != nil && o.ModelIds != nil
}

// SetModelIds gets a reference to the given []int64 and assigns it to the ModelIds field.
func (o *AiGatewayModelGroupUpdateAttributes) SetModelIds(v []int64) {
	o.ModelIds = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AiGatewayModelGroupUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AiGatewayModelGroupUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AiGatewayModelGroupUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetProviderIds returns the ProviderIds field value if set, zero value otherwise.
func (o *AiGatewayModelGroupUpdateAttributes) GetProviderIds() []int64 {
	if o == nil || o.ProviderIds == nil {
		var ret []int64
		return ret
	}
	return o.ProviderIds
}

// GetProviderIdsOk returns a tuple with the ProviderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupUpdateAttributes) GetProviderIdsOk() (*[]int64, bool) {
	if o == nil || o.ProviderIds == nil {
		return nil, false
	}
	return &o.ProviderIds, true
}

// HasProviderIds returns a boolean if a field has been set.
func (o *AiGatewayModelGroupUpdateAttributes) HasProviderIds() bool {
	return o != nil && o.ProviderIds != nil
}

// SetProviderIds gets a reference to the given []int64 and assigns it to the ProviderIds field.
func (o *AiGatewayModelGroupUpdateAttributes) SetProviderIds(v []int64) {
	o.ProviderIds = v
}

// GetRoutingRuleSlugs returns the RoutingRuleSlugs field value if set, zero value otherwise.
func (o *AiGatewayModelGroupUpdateAttributes) GetRoutingRuleSlugs() []string {
	if o == nil || o.RoutingRuleSlugs == nil {
		var ret []string
		return ret
	}
	return o.RoutingRuleSlugs
}

// GetRoutingRuleSlugsOk returns a tuple with the RoutingRuleSlugs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupUpdateAttributes) GetRoutingRuleSlugsOk() (*[]string, bool) {
	if o == nil || o.RoutingRuleSlugs == nil {
		return nil, false
	}
	return &o.RoutingRuleSlugs, true
}

// HasRoutingRuleSlugs returns a boolean if a field has been set.
func (o *AiGatewayModelGroupUpdateAttributes) HasRoutingRuleSlugs() bool {
	return o != nil && o.RoutingRuleSlugs != nil
}

// SetRoutingRuleSlugs gets a reference to the given []string and assigns it to the RoutingRuleSlugs field.
func (o *AiGatewayModelGroupUpdateAttributes) SetRoutingRuleSlugs(v []string) {
	o.RoutingRuleSlugs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelGroupUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowFastMode != nil {
		toSerialize["allow_fast_mode"] = o.AllowFastMode
	}
	if o.ExpectedRevision != nil {
		toSerialize["expected_revision"] = o.ExpectedRevision
	}
	if o.MaxReasoningLevel.IsSet() {
		toSerialize["max_reasoning_level"] = o.MaxReasoningLevel.Get()
	}
	if o.ModelIds != nil {
		toSerialize["model_ids"] = o.ModelIds
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
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
func (o *AiGatewayModelGroupUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowFastMode     *bool                                      `json:"allow_fast_mode,omitempty"`
		ExpectedRevision  *int64                                     `json:"expected_revision,omitempty"`
		MaxReasoningLevel NullableAiGatewayNullableMaxReasoningLevel `json:"max_reasoning_level,omitempty"`
		ModelIds          []int64                                    `json:"model_ids,omitempty"`
		Name              *string                                    `json:"name,omitempty"`
		ProviderIds       []int64                                    `json:"provider_ids,omitempty"`
		RoutingRuleSlugs  []string                                   `json:"routing_rule_slugs,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allow_fast_mode", "expected_revision", "max_reasoning_level", "model_ids", "name", "provider_ids", "routing_rule_slugs"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AllowFastMode = all.AllowFastMode
	o.ExpectedRevision = all.ExpectedRevision
	if all.MaxReasoningLevel.Get() != nil && !all.MaxReasoningLevel.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.MaxReasoningLevel = all.MaxReasoningLevel
	}
	o.ModelIds = all.ModelIds
	o.Name = all.Name
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
