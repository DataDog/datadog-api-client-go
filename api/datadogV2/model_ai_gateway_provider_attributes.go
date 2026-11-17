// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayProviderAttributes AI Gateway provider configuration and metadata.
type AiGatewayProviderAttributes struct {
	// Multiplier applied to provider costs.
	CostMultiplier float64 `json:"cost_multiplier"`
	// Time when the resource was created.
	CreatedAt time.Time `json:"created_at"`
	// Whose credential the gateway sends to the provider. `integration_account` sends the integration
	// account's credential. `subscription` forwards each caller's own subscription sign-in (`ChatGPT` for
	// `openai`, `claude.ai` for `anthropic`) and needs no integration account.
	CredentialSource *AiGatewayProviderCredentialSource `json:"credential_source,omitempty"`
	// Provider integration account identifier.
	IntegrationAccountId string `json:"integration_account_id"`
	// Provider-specific integration configuration.
	IntegrationConfig *AiGatewayProviderIntegrationConfig `json:"integration_config,omitempty"`
	// Known provider implementation.
	KnownProvider AiGatewayKnownProvider `json:"known_provider"`
	// Human-readable resource name.
	Name string `json:"name"`
	// Immutable routing identifier.
	Slug string `json:"slug"`
	// Read-only. Whether the provider is working, based on the model sync and live calls. Omitted for
	// providers that are never checked, which include providers without an integration account, and
	// Azure OpenAI and Gemini providers until live traffic records a status.
	Status *AiGatewayProviderStatus `json:"status,omitempty"`
	// Time when the resource was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// Whether this provider connection is covered by a zero data retention agreement.
	ZdrEligible *bool `json:"zdr_eligible,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayProviderAttributes instantiates a new AiGatewayProviderAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayProviderAttributes(costMultiplier float64, createdAt time.Time, integrationAccountId string, knownProvider AiGatewayKnownProvider, name string, slug string, updatedAt time.Time) *AiGatewayProviderAttributes {
	this := AiGatewayProviderAttributes{}
	this.CostMultiplier = costMultiplier
	this.CreatedAt = createdAt
	this.IntegrationAccountId = integrationAccountId
	this.KnownProvider = knownProvider
	this.Name = name
	this.Slug = slug
	this.UpdatedAt = updatedAt
	return &this
}

// NewAiGatewayProviderAttributesWithDefaults instantiates a new AiGatewayProviderAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayProviderAttributesWithDefaults() *AiGatewayProviderAttributes {
	this := AiGatewayProviderAttributes{}
	return &this
}

// GetCostMultiplier returns the CostMultiplier field value.
func (o *AiGatewayProviderAttributes) GetCostMultiplier() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.CostMultiplier
}

// GetCostMultiplierOk returns a tuple with the CostMultiplier field value
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderAttributes) GetCostMultiplierOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CostMultiplier, true
}

// SetCostMultiplier sets field value.
func (o *AiGatewayProviderAttributes) SetCostMultiplier(v float64) {
	o.CostMultiplier = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AiGatewayProviderAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AiGatewayProviderAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCredentialSource returns the CredentialSource field value if set, zero value otherwise.
func (o *AiGatewayProviderAttributes) GetCredentialSource() AiGatewayProviderCredentialSource {
	if o == nil || o.CredentialSource == nil {
		var ret AiGatewayProviderCredentialSource
		return ret
	}
	return *o.CredentialSource
}

// GetCredentialSourceOk returns a tuple with the CredentialSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderAttributes) GetCredentialSourceOk() (*AiGatewayProviderCredentialSource, bool) {
	if o == nil || o.CredentialSource == nil {
		return nil, false
	}
	return o.CredentialSource, true
}

// HasCredentialSource returns a boolean if a field has been set.
func (o *AiGatewayProviderAttributes) HasCredentialSource() bool {
	return o != nil && o.CredentialSource != nil
}

// SetCredentialSource gets a reference to the given AiGatewayProviderCredentialSource and assigns it to the CredentialSource field.
func (o *AiGatewayProviderAttributes) SetCredentialSource(v AiGatewayProviderCredentialSource) {
	o.CredentialSource = &v
}

// GetIntegrationAccountId returns the IntegrationAccountId field value.
func (o *AiGatewayProviderAttributes) GetIntegrationAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IntegrationAccountId
}

// GetIntegrationAccountIdOk returns a tuple with the IntegrationAccountId field value
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderAttributes) GetIntegrationAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationAccountId, true
}

// SetIntegrationAccountId sets field value.
func (o *AiGatewayProviderAttributes) SetIntegrationAccountId(v string) {
	o.IntegrationAccountId = v
}

// GetIntegrationConfig returns the IntegrationConfig field value if set, zero value otherwise.
func (o *AiGatewayProviderAttributes) GetIntegrationConfig() AiGatewayProviderIntegrationConfig {
	if o == nil || o.IntegrationConfig == nil {
		var ret AiGatewayProviderIntegrationConfig
		return ret
	}
	return *o.IntegrationConfig
}

// GetIntegrationConfigOk returns a tuple with the IntegrationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderAttributes) GetIntegrationConfigOk() (*AiGatewayProviderIntegrationConfig, bool) {
	if o == nil || o.IntegrationConfig == nil {
		return nil, false
	}
	return o.IntegrationConfig, true
}

// HasIntegrationConfig returns a boolean if a field has been set.
func (o *AiGatewayProviderAttributes) HasIntegrationConfig() bool {
	return o != nil && o.IntegrationConfig != nil
}

// SetIntegrationConfig gets a reference to the given AiGatewayProviderIntegrationConfig and assigns it to the IntegrationConfig field.
func (o *AiGatewayProviderAttributes) SetIntegrationConfig(v AiGatewayProviderIntegrationConfig) {
	o.IntegrationConfig = &v
}

// GetKnownProvider returns the KnownProvider field value.
func (o *AiGatewayProviderAttributes) GetKnownProvider() AiGatewayKnownProvider {
	if o == nil {
		var ret AiGatewayKnownProvider
		return ret
	}
	return o.KnownProvider
}

// GetKnownProviderOk returns a tuple with the KnownProvider field value
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderAttributes) GetKnownProviderOk() (*AiGatewayKnownProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KnownProvider, true
}

// SetKnownProvider sets field value.
func (o *AiGatewayProviderAttributes) SetKnownProvider(v AiGatewayKnownProvider) {
	o.KnownProvider = v
}

// GetName returns the Name field value.
func (o *AiGatewayProviderAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AiGatewayProviderAttributes) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value.
func (o *AiGatewayProviderAttributes) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderAttributes) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value.
func (o *AiGatewayProviderAttributes) SetSlug(v string) {
	o.Slug = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AiGatewayProviderAttributes) GetStatus() AiGatewayProviderStatus {
	if o == nil || o.Status == nil {
		var ret AiGatewayProviderStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderAttributes) GetStatusOk() (*AiGatewayProviderStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiGatewayProviderAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given AiGatewayProviderStatus and assigns it to the Status field.
func (o *AiGatewayProviderAttributes) SetStatus(v AiGatewayProviderStatus) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *AiGatewayProviderAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *AiGatewayProviderAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetZdrEligible returns the ZdrEligible field value if set, zero value otherwise.
func (o *AiGatewayProviderAttributes) GetZdrEligible() bool {
	if o == nil || o.ZdrEligible == nil {
		var ret bool
		return ret
	}
	return *o.ZdrEligible
}

// GetZdrEligibleOk returns a tuple with the ZdrEligible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderAttributes) GetZdrEligibleOk() (*bool, bool) {
	if o == nil || o.ZdrEligible == nil {
		return nil, false
	}
	return o.ZdrEligible, true
}

// HasZdrEligible returns a boolean if a field has been set.
func (o *AiGatewayProviderAttributes) HasZdrEligible() bool {
	return o != nil && o.ZdrEligible != nil
}

// SetZdrEligible gets a reference to the given bool and assigns it to the ZdrEligible field.
func (o *AiGatewayProviderAttributes) SetZdrEligible(v bool) {
	o.ZdrEligible = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayProviderAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cost_multiplier"] = o.CostMultiplier
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.CredentialSource != nil {
		toSerialize["credential_source"] = o.CredentialSource
	}
	toSerialize["integration_account_id"] = o.IntegrationAccountId
	if o.IntegrationConfig != nil {
		toSerialize["integration_config"] = o.IntegrationConfig
	}
	toSerialize["known_provider"] = o.KnownProvider
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ZdrEligible != nil {
		toSerialize["zdr_eligible"] = o.ZdrEligible
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayProviderAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CostMultiplier       *float64                            `json:"cost_multiplier"`
		CreatedAt            *time.Time                          `json:"created_at"`
		CredentialSource     *AiGatewayProviderCredentialSource  `json:"credential_source,omitempty"`
		IntegrationAccountId *string                             `json:"integration_account_id"`
		IntegrationConfig    *AiGatewayProviderIntegrationConfig `json:"integration_config,omitempty"`
		KnownProvider        *AiGatewayKnownProvider             `json:"known_provider"`
		Name                 *string                             `json:"name"`
		Slug                 *string                             `json:"slug"`
		Status               *AiGatewayProviderStatus            `json:"status,omitempty"`
		UpdatedAt            *time.Time                          `json:"updated_at"`
		ZdrEligible          *bool                               `json:"zdr_eligible,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CostMultiplier == nil {
		return fmt.Errorf("required field cost_multiplier missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.IntegrationAccountId == nil {
		return fmt.Errorf("required field integration_account_id missing")
	}
	if all.KnownProvider == nil {
		return fmt.Errorf("required field known_provider missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Slug == nil {
		return fmt.Errorf("required field slug missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cost_multiplier", "created_at", "credential_source", "integration_account_id", "integration_config", "known_provider", "name", "slug", "status", "updated_at", "zdr_eligible"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CostMultiplier = *all.CostMultiplier
	o.CreatedAt = *all.CreatedAt
	if all.CredentialSource != nil && !all.CredentialSource.IsValid() {
		hasInvalidField = true
	} else {
		o.CredentialSource = all.CredentialSource
	}
	o.IntegrationAccountId = *all.IntegrationAccountId
	if all.IntegrationConfig != nil && all.IntegrationConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IntegrationConfig = all.IntegrationConfig
	if !all.KnownProvider.IsValid() {
		hasInvalidField = true
	} else {
		o.KnownProvider = *all.KnownProvider
	}
	o.Name = *all.Name
	o.Slug = *all.Slug
	if all.Status != nil && all.Status.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Status = all.Status
	o.UpdatedAt = *all.UpdatedAt
	o.ZdrEligible = all.ZdrEligible

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
