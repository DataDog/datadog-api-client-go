// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayProviderCreateAttributes Attributes used to create an AI Gateway provider.
type AiGatewayProviderCreateAttributes struct {
	// Multiplier applied to provider costs. Must be greater than 0 and at most 1000000.
	CostMultiplier *float64 `json:"cost_multiplier,omitempty"`
	// Whose credential the gateway sends to the provider. Defaults to `integration_account`, which sends
	// the integration account's credential. `subscription` forwards each caller's own subscription sign-in
	// (`ChatGPT` for `openai`, `claude.ai` for `anthropic`) and needs no integration account; it is available
	// only to organizations with subscription providers enabled. A subscription provider's models are
	// optional: callers can use any model their subscription includes, and a model you add sets its
	// pricing. Immutable.
	CredentialSource *AiGatewayProviderCreateCredentialSource `json:"credential_source,omitempty"`
	// Reserved for initial View/Edit permissions. Supplying this attribute, including an empty array or null, returns HTTP 503 before discovery or persistence. Omit it for normal creation.
	InitialBindings []AiGatewayProviderInitialBinding `json:"initial_bindings,omitempty"`
	// Provider integration account identifier.
	IntegrationAccountId *string `json:"integration_account_id,omitempty"`
	// Provider-specific integration configuration.
	IntegrationConfig *AiGatewayProviderIntegrationConfig `json:"integration_config,omitempty"`
	// Known provider implementation.
	KnownProvider *AiGatewayKnownProvider `json:"known_provider,omitempty"`
	// Manually managed models to add alongside models discovered from the provider integration.
	Models []AiGatewayModelInput `json:"models,omitempty"`
	// Human-readable resource name.
	Name string `json:"name"`
	// Immutable routing identifier.
	Slug string `json:"slug"`
	// Whether this provider connection is covered by a zero data retention agreement. Defaults to `false`.
	// For Bedrock, zero data retention is granted per region, so configure one provider per region and set
	// this only on the covered ones.
	ZdrEligible *bool `json:"zdr_eligible,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayProviderCreateAttributes instantiates a new AiGatewayProviderCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayProviderCreateAttributes(name string, slug string) *AiGatewayProviderCreateAttributes {
	this := AiGatewayProviderCreateAttributes{}
	var costMultiplier float64 = 1.0
	this.CostMultiplier = &costMultiplier
	var credentialSource AiGatewayProviderCreateCredentialSource = AIGATEWAYPROVIDERCREATECREDENTIALSOURCE_INTEGRATION_ACCOUNT
	this.CredentialSource = &credentialSource
	this.Name = name
	this.Slug = slug
	return &this
}

// NewAiGatewayProviderCreateAttributesWithDefaults instantiates a new AiGatewayProviderCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayProviderCreateAttributesWithDefaults() *AiGatewayProviderCreateAttributes {
	this := AiGatewayProviderCreateAttributes{}
	var costMultiplier float64 = 1.0
	this.CostMultiplier = &costMultiplier
	var credentialSource AiGatewayProviderCreateCredentialSource = AIGATEWAYPROVIDERCREATECREDENTIALSOURCE_INTEGRATION_ACCOUNT
	this.CredentialSource = &credentialSource
	return &this
}

// GetCostMultiplier returns the CostMultiplier field value if set, zero value otherwise.
func (o *AiGatewayProviderCreateAttributes) GetCostMultiplier() float64 {
	if o == nil || o.CostMultiplier == nil {
		var ret float64
		return ret
	}
	return *o.CostMultiplier
}

// GetCostMultiplierOk returns a tuple with the CostMultiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderCreateAttributes) GetCostMultiplierOk() (*float64, bool) {
	if o == nil || o.CostMultiplier == nil {
		return nil, false
	}
	return o.CostMultiplier, true
}

// HasCostMultiplier returns a boolean if a field has been set.
func (o *AiGatewayProviderCreateAttributes) HasCostMultiplier() bool {
	return o != nil && o.CostMultiplier != nil
}

// SetCostMultiplier gets a reference to the given float64 and assigns it to the CostMultiplier field.
func (o *AiGatewayProviderCreateAttributes) SetCostMultiplier(v float64) {
	o.CostMultiplier = &v
}

// GetCredentialSource returns the CredentialSource field value if set, zero value otherwise.
func (o *AiGatewayProviderCreateAttributes) GetCredentialSource() AiGatewayProviderCreateCredentialSource {
	if o == nil || o.CredentialSource == nil {
		var ret AiGatewayProviderCreateCredentialSource
		return ret
	}
	return *o.CredentialSource
}

// GetCredentialSourceOk returns a tuple with the CredentialSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderCreateAttributes) GetCredentialSourceOk() (*AiGatewayProviderCreateCredentialSource, bool) {
	if o == nil || o.CredentialSource == nil {
		return nil, false
	}
	return o.CredentialSource, true
}

// HasCredentialSource returns a boolean if a field has been set.
func (o *AiGatewayProviderCreateAttributes) HasCredentialSource() bool {
	return o != nil && o.CredentialSource != nil
}

// SetCredentialSource gets a reference to the given AiGatewayProviderCreateCredentialSource and assigns it to the CredentialSource field.
func (o *AiGatewayProviderCreateAttributes) SetCredentialSource(v AiGatewayProviderCreateCredentialSource) {
	o.CredentialSource = &v
}

// GetInitialBindings returns the InitialBindings field value if set, zero value otherwise.
func (o *AiGatewayProviderCreateAttributes) GetInitialBindings() []AiGatewayProviderInitialBinding {
	if o == nil || o.InitialBindings == nil {
		var ret []AiGatewayProviderInitialBinding
		return ret
	}
	return o.InitialBindings
}

// GetInitialBindingsOk returns a tuple with the InitialBindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderCreateAttributes) GetInitialBindingsOk() (*[]AiGatewayProviderInitialBinding, bool) {
	if o == nil || o.InitialBindings == nil {
		return nil, false
	}
	return &o.InitialBindings, true
}

// HasInitialBindings returns a boolean if a field has been set.
func (o *AiGatewayProviderCreateAttributes) HasInitialBindings() bool {
	return o != nil && o.InitialBindings != nil
}

// SetInitialBindings gets a reference to the given []AiGatewayProviderInitialBinding and assigns it to the InitialBindings field.
func (o *AiGatewayProviderCreateAttributes) SetInitialBindings(v []AiGatewayProviderInitialBinding) {
	o.InitialBindings = v
}

// GetIntegrationAccountId returns the IntegrationAccountId field value if set, zero value otherwise.
func (o *AiGatewayProviderCreateAttributes) GetIntegrationAccountId() string {
	if o == nil || o.IntegrationAccountId == nil {
		var ret string
		return ret
	}
	return *o.IntegrationAccountId
}

// GetIntegrationAccountIdOk returns a tuple with the IntegrationAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderCreateAttributes) GetIntegrationAccountIdOk() (*string, bool) {
	if o == nil || o.IntegrationAccountId == nil {
		return nil, false
	}
	return o.IntegrationAccountId, true
}

// HasIntegrationAccountId returns a boolean if a field has been set.
func (o *AiGatewayProviderCreateAttributes) HasIntegrationAccountId() bool {
	return o != nil && o.IntegrationAccountId != nil
}

// SetIntegrationAccountId gets a reference to the given string and assigns it to the IntegrationAccountId field.
func (o *AiGatewayProviderCreateAttributes) SetIntegrationAccountId(v string) {
	o.IntegrationAccountId = &v
}

// GetIntegrationConfig returns the IntegrationConfig field value if set, zero value otherwise.
func (o *AiGatewayProviderCreateAttributes) GetIntegrationConfig() AiGatewayProviderIntegrationConfig {
	if o == nil || o.IntegrationConfig == nil {
		var ret AiGatewayProviderIntegrationConfig
		return ret
	}
	return *o.IntegrationConfig
}

// GetIntegrationConfigOk returns a tuple with the IntegrationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderCreateAttributes) GetIntegrationConfigOk() (*AiGatewayProviderIntegrationConfig, bool) {
	if o == nil || o.IntegrationConfig == nil {
		return nil, false
	}
	return o.IntegrationConfig, true
}

// HasIntegrationConfig returns a boolean if a field has been set.
func (o *AiGatewayProviderCreateAttributes) HasIntegrationConfig() bool {
	return o != nil && o.IntegrationConfig != nil
}

// SetIntegrationConfig gets a reference to the given AiGatewayProviderIntegrationConfig and assigns it to the IntegrationConfig field.
func (o *AiGatewayProviderCreateAttributes) SetIntegrationConfig(v AiGatewayProviderIntegrationConfig) {
	o.IntegrationConfig = &v
}

// GetKnownProvider returns the KnownProvider field value if set, zero value otherwise.
func (o *AiGatewayProviderCreateAttributes) GetKnownProvider() AiGatewayKnownProvider {
	if o == nil || o.KnownProvider == nil {
		var ret AiGatewayKnownProvider
		return ret
	}
	return *o.KnownProvider
}

// GetKnownProviderOk returns a tuple with the KnownProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderCreateAttributes) GetKnownProviderOk() (*AiGatewayKnownProvider, bool) {
	if o == nil || o.KnownProvider == nil {
		return nil, false
	}
	return o.KnownProvider, true
}

// HasKnownProvider returns a boolean if a field has been set.
func (o *AiGatewayProviderCreateAttributes) HasKnownProvider() bool {
	return o != nil && o.KnownProvider != nil
}

// SetKnownProvider gets a reference to the given AiGatewayKnownProvider and assigns it to the KnownProvider field.
func (o *AiGatewayProviderCreateAttributes) SetKnownProvider(v AiGatewayKnownProvider) {
	o.KnownProvider = &v
}

// GetModels returns the Models field value if set, zero value otherwise.
func (o *AiGatewayProviderCreateAttributes) GetModels() []AiGatewayModelInput {
	if o == nil || o.Models == nil {
		var ret []AiGatewayModelInput
		return ret
	}
	return o.Models
}

// GetModelsOk returns a tuple with the Models field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderCreateAttributes) GetModelsOk() (*[]AiGatewayModelInput, bool) {
	if o == nil || o.Models == nil {
		return nil, false
	}
	return &o.Models, true
}

// HasModels returns a boolean if a field has been set.
func (o *AiGatewayProviderCreateAttributes) HasModels() bool {
	return o != nil && o.Models != nil
}

// SetModels gets a reference to the given []AiGatewayModelInput and assigns it to the Models field.
func (o *AiGatewayProviderCreateAttributes) SetModels(v []AiGatewayModelInput) {
	o.Models = v
}

// GetName returns the Name field value.
func (o *AiGatewayProviderCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AiGatewayProviderCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value.
func (o *AiGatewayProviderCreateAttributes) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderCreateAttributes) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value.
func (o *AiGatewayProviderCreateAttributes) SetSlug(v string) {
	o.Slug = v
}

// GetZdrEligible returns the ZdrEligible field value if set, zero value otherwise.
func (o *AiGatewayProviderCreateAttributes) GetZdrEligible() bool {
	if o == nil || o.ZdrEligible == nil {
		var ret bool
		return ret
	}
	return *o.ZdrEligible
}

// GetZdrEligibleOk returns a tuple with the ZdrEligible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderCreateAttributes) GetZdrEligibleOk() (*bool, bool) {
	if o == nil || o.ZdrEligible == nil {
		return nil, false
	}
	return o.ZdrEligible, true
}

// HasZdrEligible returns a boolean if a field has been set.
func (o *AiGatewayProviderCreateAttributes) HasZdrEligible() bool {
	return o != nil && o.ZdrEligible != nil
}

// SetZdrEligible gets a reference to the given bool and assigns it to the ZdrEligible field.
func (o *AiGatewayProviderCreateAttributes) SetZdrEligible(v bool) {
	o.ZdrEligible = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayProviderCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CostMultiplier != nil {
		toSerialize["cost_multiplier"] = o.CostMultiplier
	}
	if o.CredentialSource != nil {
		toSerialize["credential_source"] = o.CredentialSource
	}
	if o.InitialBindings != nil {
		toSerialize["initial_bindings"] = o.InitialBindings
	}
	if o.IntegrationAccountId != nil {
		toSerialize["integration_account_id"] = o.IntegrationAccountId
	}
	if o.IntegrationConfig != nil {
		toSerialize["integration_config"] = o.IntegrationConfig
	}
	if o.KnownProvider != nil {
		toSerialize["known_provider"] = o.KnownProvider
	}
	if o.Models != nil {
		toSerialize["models"] = o.Models
	}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if o.ZdrEligible != nil {
		toSerialize["zdr_eligible"] = o.ZdrEligible
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayProviderCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CostMultiplier       *float64                                 `json:"cost_multiplier,omitempty"`
		CredentialSource     *AiGatewayProviderCreateCredentialSource `json:"credential_source,omitempty"`
		InitialBindings      []AiGatewayProviderInitialBinding        `json:"initial_bindings,omitempty"`
		IntegrationAccountId *string                                  `json:"integration_account_id,omitempty"`
		IntegrationConfig    *AiGatewayProviderIntegrationConfig      `json:"integration_config,omitempty"`
		KnownProvider        *AiGatewayKnownProvider                  `json:"known_provider,omitempty"`
		Models               []AiGatewayModelInput                    `json:"models,omitempty"`
		Name                 *string                                  `json:"name"`
		Slug                 *string                                  `json:"slug"`
		ZdrEligible          *bool                                    `json:"zdr_eligible,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Slug == nil {
		return fmt.Errorf("required field slug missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cost_multiplier", "credential_source", "initial_bindings", "integration_account_id", "integration_config", "known_provider", "models", "name", "slug", "zdr_eligible"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CostMultiplier = all.CostMultiplier
	if all.CredentialSource != nil && !all.CredentialSource.IsValid() {
		hasInvalidField = true
	} else {
		o.CredentialSource = all.CredentialSource
	}
	o.InitialBindings = all.InitialBindings
	o.IntegrationAccountId = all.IntegrationAccountId
	if all.IntegrationConfig != nil && all.IntegrationConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IntegrationConfig = all.IntegrationConfig
	if all.KnownProvider != nil && !all.KnownProvider.IsValid() {
		hasInvalidField = true
	} else {
		o.KnownProvider = all.KnownProvider
	}
	o.Models = all.Models
	o.Name = *all.Name
	o.Slug = *all.Slug
	o.ZdrEligible = all.ZdrEligible

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
