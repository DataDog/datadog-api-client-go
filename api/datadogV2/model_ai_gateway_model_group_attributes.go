// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelGroupAttributes Complete AI Gateway model group configuration.
type AiGatewayModelGroupAttributes struct {
	// Whether fast mode is allowed for models and routing rules in the group.
	AllowFastMode bool `json:"allow_fast_mode"`
	// Time when the resource was created.
	CreatedAt time.Time `json:"created_at"`
	// Maximum reasoning level permitted by a model group.
	MaxReasoningLevel NullableAiGatewayNullableMaxReasoningLevel `json:"max_reasoning_level"`
	// All saved model catalog IDs, including selections whose provider metadata is hidden. Preserve these IDs independently of the visible catalog.
	ModelIds []int64 `json:"model_ids"`
	// Metadata for saved model selections whose providers the caller can view. Missing metadata does not remove a saved ID.
	Models []AiGatewayModelGroupModel `json:"models"`
	// Human-readable resource name.
	Name string `json:"name"`
	// All explicitly selected provider IDs, including hidden providers. Embedded provider metadata includes only viewable providers.
	ProviderIds []int64 `json:"provider_ids"`
	// Viewable explicitly selected providers, separate from individually selected models. Use provider_ids to retain all selections.
	Providers []AiGatewayModelGroupProvider `json:"providers"`
	// Revision of the organization-wide policy. Present for the default model group and absent for named policies.
	Revision *int64 `json:"revision,omitempty"`
	//
	RoutingRules []AiGatewayModelGroupRoutingRule `json:"routing_rules"`
	// Time when the resource was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelGroupAttributes instantiates a new AiGatewayModelGroupAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelGroupAttributes(allowFastMode bool, createdAt time.Time, maxReasoningLevel NullableAiGatewayNullableMaxReasoningLevel, modelIds []int64, models []AiGatewayModelGroupModel, name string, providerIds []int64, providers []AiGatewayModelGroupProvider, routingRules []AiGatewayModelGroupRoutingRule, updatedAt time.Time) *AiGatewayModelGroupAttributes {
	this := AiGatewayModelGroupAttributes{}
	this.AllowFastMode = allowFastMode
	this.CreatedAt = createdAt
	this.MaxReasoningLevel = maxReasoningLevel
	this.ModelIds = modelIds
	this.Models = models
	this.Name = name
	this.ProviderIds = providerIds
	this.Providers = providers
	this.RoutingRules = routingRules
	this.UpdatedAt = updatedAt
	return &this
}

// NewAiGatewayModelGroupAttributesWithDefaults instantiates a new AiGatewayModelGroupAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelGroupAttributesWithDefaults() *AiGatewayModelGroupAttributes {
	this := AiGatewayModelGroupAttributes{}
	return &this
}

// GetAllowFastMode returns the AllowFastMode field value.
func (o *AiGatewayModelGroupAttributes) GetAllowFastMode() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.AllowFastMode
}

// GetAllowFastModeOk returns a tuple with the AllowFastMode field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupAttributes) GetAllowFastModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowFastMode, true
}

// SetAllowFastMode sets field value.
func (o *AiGatewayModelGroupAttributes) SetAllowFastMode(v bool) {
	o.AllowFastMode = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AiGatewayModelGroupAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AiGatewayModelGroupAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetMaxReasoningLevel returns the MaxReasoningLevel field value.
// If the value is explicit nil, the zero value for AiGatewayNullableMaxReasoningLevel will be returned.
func (o *AiGatewayModelGroupAttributes) GetMaxReasoningLevel() AiGatewayNullableMaxReasoningLevel {
	if o == nil || o.MaxReasoningLevel.Get() == nil {
		var ret AiGatewayNullableMaxReasoningLevel
		return ret
	}
	return *o.MaxReasoningLevel.Get()
}

// GetMaxReasoningLevelOk returns a tuple with the MaxReasoningLevel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayModelGroupAttributes) GetMaxReasoningLevelOk() (*AiGatewayNullableMaxReasoningLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxReasoningLevel.Get(), o.MaxReasoningLevel.IsSet()
}

// SetMaxReasoningLevel sets field value.
func (o *AiGatewayModelGroupAttributes) SetMaxReasoningLevel(v AiGatewayNullableMaxReasoningLevel) {
	o.MaxReasoningLevel.Set(&v)
}

// GetModelIds returns the ModelIds field value.
func (o *AiGatewayModelGroupAttributes) GetModelIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.ModelIds
}

// GetModelIdsOk returns a tuple with the ModelIds field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupAttributes) GetModelIdsOk() (*[]int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelIds, true
}

// SetModelIds sets field value.
func (o *AiGatewayModelGroupAttributes) SetModelIds(v []int64) {
	o.ModelIds = v
}

// GetModels returns the Models field value.
func (o *AiGatewayModelGroupAttributes) GetModels() []AiGatewayModelGroupModel {
	if o == nil {
		var ret []AiGatewayModelGroupModel
		return ret
	}
	return o.Models
}

// GetModelsOk returns a tuple with the Models field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupAttributes) GetModelsOk() (*[]AiGatewayModelGroupModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Models, true
}

// SetModels sets field value.
func (o *AiGatewayModelGroupAttributes) SetModels(v []AiGatewayModelGroupModel) {
	o.Models = v
}

// GetName returns the Name field value.
func (o *AiGatewayModelGroupAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AiGatewayModelGroupAttributes) SetName(v string) {
	o.Name = v
}

// GetProviderIds returns the ProviderIds field value.
func (o *AiGatewayModelGroupAttributes) GetProviderIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.ProviderIds
}

// GetProviderIdsOk returns a tuple with the ProviderIds field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupAttributes) GetProviderIdsOk() (*[]int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderIds, true
}

// SetProviderIds sets field value.
func (o *AiGatewayModelGroupAttributes) SetProviderIds(v []int64) {
	o.ProviderIds = v
}

// GetProviders returns the Providers field value.
func (o *AiGatewayModelGroupAttributes) GetProviders() []AiGatewayModelGroupProvider {
	if o == nil {
		var ret []AiGatewayModelGroupProvider
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupAttributes) GetProvidersOk() (*[]AiGatewayModelGroupProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Providers, true
}

// SetProviders sets field value.
func (o *AiGatewayModelGroupAttributes) SetProviders(v []AiGatewayModelGroupProvider) {
	o.Providers = v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *AiGatewayModelGroupAttributes) GetRevision() int64 {
	if o == nil || o.Revision == nil {
		var ret int64
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupAttributes) GetRevisionOk() (*int64, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *AiGatewayModelGroupAttributes) HasRevision() bool {
	return o != nil && o.Revision != nil
}

// SetRevision gets a reference to the given int64 and assigns it to the Revision field.
func (o *AiGatewayModelGroupAttributes) SetRevision(v int64) {
	o.Revision = &v
}

// GetRoutingRules returns the RoutingRules field value.
func (o *AiGatewayModelGroupAttributes) GetRoutingRules() []AiGatewayModelGroupRoutingRule {
	if o == nil {
		var ret []AiGatewayModelGroupRoutingRule
		return ret
	}
	return o.RoutingRules
}

// GetRoutingRulesOk returns a tuple with the RoutingRules field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupAttributes) GetRoutingRulesOk() (*[]AiGatewayModelGroupRoutingRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoutingRules, true
}

// SetRoutingRules sets field value.
func (o *AiGatewayModelGroupAttributes) SetRoutingRules(v []AiGatewayModelGroupRoutingRule) {
	o.RoutingRules = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *AiGatewayModelGroupAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *AiGatewayModelGroupAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelGroupAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["allow_fast_mode"] = o.AllowFastMode
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["max_reasoning_level"] = o.MaxReasoningLevel.Get()
	toSerialize["model_ids"] = o.ModelIds
	toSerialize["models"] = o.Models
	toSerialize["name"] = o.Name
	toSerialize["provider_ids"] = o.ProviderIds
	toSerialize["providers"] = o.Providers
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	toSerialize["routing_rules"] = o.RoutingRules
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelGroupAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowFastMode     *bool                                      `json:"allow_fast_mode"`
		CreatedAt         *time.Time                                 `json:"created_at"`
		MaxReasoningLevel NullableAiGatewayNullableMaxReasoningLevel `json:"max_reasoning_level"`
		ModelIds          *[]int64                                   `json:"model_ids"`
		Models            *[]AiGatewayModelGroupModel                `json:"models"`
		Name              *string                                    `json:"name"`
		ProviderIds       *[]int64                                   `json:"provider_ids"`
		Providers         *[]AiGatewayModelGroupProvider             `json:"providers"`
		Revision          *int64                                     `json:"revision,omitempty"`
		RoutingRules      *[]AiGatewayModelGroupRoutingRule          `json:"routing_rules"`
		UpdatedAt         *time.Time                                 `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AllowFastMode == nil {
		return fmt.Errorf("required field allow_fast_mode missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if !all.MaxReasoningLevel.IsSet() {
		return fmt.Errorf("required field max_reasoning_level missing")
	}
	if all.ModelIds == nil {
		return fmt.Errorf("required field model_ids missing")
	}
	if all.Models == nil {
		return fmt.Errorf("required field models missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ProviderIds == nil {
		return fmt.Errorf("required field provider_ids missing")
	}
	if all.Providers == nil {
		return fmt.Errorf("required field providers missing")
	}
	if all.RoutingRules == nil {
		return fmt.Errorf("required field routing_rules missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allow_fast_mode", "created_at", "max_reasoning_level", "model_ids", "models", "name", "provider_ids", "providers", "revision", "routing_rules", "updated_at"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AllowFastMode = *all.AllowFastMode
	o.CreatedAt = *all.CreatedAt
	if all.MaxReasoningLevel.Get() != nil && !all.MaxReasoningLevel.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.MaxReasoningLevel = all.MaxReasoningLevel
	}
	o.ModelIds = *all.ModelIds
	o.Models = *all.Models
	o.Name = *all.Name
	o.ProviderIds = *all.ProviderIds
	o.Providers = *all.Providers
	o.Revision = all.Revision
	o.RoutingRules = *all.RoutingRules
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
