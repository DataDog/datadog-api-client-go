// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelPricingAttributes Attributes of a model pricing resource.
type AiGatewayModelPricingAttributes struct {
	// True when automatic pricing is authoritative for this organization and provider.
	Authoritative *bool `json:"authoritative,omitempty"`
	// The automatically synced published price. Present only when it is the effective price.
	AutomaticPricing *AiGatewayModelAutomaticPricing `json:"automatic_pricing,omitempty"`
	// True when a `BLOCK` budget covering this model rejects its requests before they reach the provider.
	BlocksBudgetedRequests *bool `json:"blocks_budgeted_requests,omitempty"`
	// Effective provider multiplier applied to every reported rate.
	CostMultiplier *float64 `json:"cost_multiplier,omitempty"`
	// True when the provider bills this model per minute of audio instead of per token, as with transcription
	// models. The gateway prices only per token, so it can't charge for these models, and a per-token
	// override can't express their rate.
	DurationPriced *bool `json:"duration_priced,omitempty"`
	// Effective source of the model pricing data.
	EffectiveSource *AiGatewayModelPricingEffectiveSource `json:"effective_source,omitempty"`
	// Pricing configuration for a model.
	PricingConfig *AiGatewayModelPricingConfig `json:"pricing_config,omitempty"`
	// Whether the gateway has a price for the model.
	Status *AiGatewayModelPricingStatus `json:"status,omitempty"`
	// Human-readable sentence for `unpriced_reason`. Omitted when the model is priced or the code has no description yet.
	UnpricedDescription *string `json:"unpriced_description,omitempty"`
	// Stable code explaining why the model has no price, such as `automatic_price_absent`. Omitted when the model is priced.
	UnpricedReason *string `json:"unpriced_reason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelPricingAttributes instantiates a new AiGatewayModelPricingAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelPricingAttributes() *AiGatewayModelPricingAttributes {
	this := AiGatewayModelPricingAttributes{}
	return &this
}

// NewAiGatewayModelPricingAttributesWithDefaults instantiates a new AiGatewayModelPricingAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelPricingAttributesWithDefaults() *AiGatewayModelPricingAttributes {
	this := AiGatewayModelPricingAttributes{}
	return &this
}

// GetAuthoritative returns the Authoritative field value if set, zero value otherwise.
func (o *AiGatewayModelPricingAttributes) GetAuthoritative() bool {
	if o == nil || o.Authoritative == nil {
		var ret bool
		return ret
	}
	return *o.Authoritative
}

// GetAuthoritativeOk returns a tuple with the Authoritative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingAttributes) GetAuthoritativeOk() (*bool, bool) {
	if o == nil || o.Authoritative == nil {
		return nil, false
	}
	return o.Authoritative, true
}

// HasAuthoritative returns a boolean if a field has been set.
func (o *AiGatewayModelPricingAttributes) HasAuthoritative() bool {
	return o != nil && o.Authoritative != nil
}

// SetAuthoritative gets a reference to the given bool and assigns it to the Authoritative field.
func (o *AiGatewayModelPricingAttributes) SetAuthoritative(v bool) {
	o.Authoritative = &v
}

// GetAutomaticPricing returns the AutomaticPricing field value if set, zero value otherwise.
func (o *AiGatewayModelPricingAttributes) GetAutomaticPricing() AiGatewayModelAutomaticPricing {
	if o == nil || o.AutomaticPricing == nil {
		var ret AiGatewayModelAutomaticPricing
		return ret
	}
	return *o.AutomaticPricing
}

// GetAutomaticPricingOk returns a tuple with the AutomaticPricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingAttributes) GetAutomaticPricingOk() (*AiGatewayModelAutomaticPricing, bool) {
	if o == nil || o.AutomaticPricing == nil {
		return nil, false
	}
	return o.AutomaticPricing, true
}

// HasAutomaticPricing returns a boolean if a field has been set.
func (o *AiGatewayModelPricingAttributes) HasAutomaticPricing() bool {
	return o != nil && o.AutomaticPricing != nil
}

// SetAutomaticPricing gets a reference to the given AiGatewayModelAutomaticPricing and assigns it to the AutomaticPricing field.
func (o *AiGatewayModelPricingAttributes) SetAutomaticPricing(v AiGatewayModelAutomaticPricing) {
	o.AutomaticPricing = &v
}

// GetBlocksBudgetedRequests returns the BlocksBudgetedRequests field value if set, zero value otherwise.
func (o *AiGatewayModelPricingAttributes) GetBlocksBudgetedRequests() bool {
	if o == nil || o.BlocksBudgetedRequests == nil {
		var ret bool
		return ret
	}
	return *o.BlocksBudgetedRequests
}

// GetBlocksBudgetedRequestsOk returns a tuple with the BlocksBudgetedRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingAttributes) GetBlocksBudgetedRequestsOk() (*bool, bool) {
	if o == nil || o.BlocksBudgetedRequests == nil {
		return nil, false
	}
	return o.BlocksBudgetedRequests, true
}

// HasBlocksBudgetedRequests returns a boolean if a field has been set.
func (o *AiGatewayModelPricingAttributes) HasBlocksBudgetedRequests() bool {
	return o != nil && o.BlocksBudgetedRequests != nil
}

// SetBlocksBudgetedRequests gets a reference to the given bool and assigns it to the BlocksBudgetedRequests field.
func (o *AiGatewayModelPricingAttributes) SetBlocksBudgetedRequests(v bool) {
	o.BlocksBudgetedRequests = &v
}

// GetCostMultiplier returns the CostMultiplier field value if set, zero value otherwise.
func (o *AiGatewayModelPricingAttributes) GetCostMultiplier() float64 {
	if o == nil || o.CostMultiplier == nil {
		var ret float64
		return ret
	}
	return *o.CostMultiplier
}

// GetCostMultiplierOk returns a tuple with the CostMultiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingAttributes) GetCostMultiplierOk() (*float64, bool) {
	if o == nil || o.CostMultiplier == nil {
		return nil, false
	}
	return o.CostMultiplier, true
}

// HasCostMultiplier returns a boolean if a field has been set.
func (o *AiGatewayModelPricingAttributes) HasCostMultiplier() bool {
	return o != nil && o.CostMultiplier != nil
}

// SetCostMultiplier gets a reference to the given float64 and assigns it to the CostMultiplier field.
func (o *AiGatewayModelPricingAttributes) SetCostMultiplier(v float64) {
	o.CostMultiplier = &v
}

// GetDurationPriced returns the DurationPriced field value if set, zero value otherwise.
func (o *AiGatewayModelPricingAttributes) GetDurationPriced() bool {
	if o == nil || o.DurationPriced == nil {
		var ret bool
		return ret
	}
	return *o.DurationPriced
}

// GetDurationPricedOk returns a tuple with the DurationPriced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingAttributes) GetDurationPricedOk() (*bool, bool) {
	if o == nil || o.DurationPriced == nil {
		return nil, false
	}
	return o.DurationPriced, true
}

// HasDurationPriced returns a boolean if a field has been set.
func (o *AiGatewayModelPricingAttributes) HasDurationPriced() bool {
	return o != nil && o.DurationPriced != nil
}

// SetDurationPriced gets a reference to the given bool and assigns it to the DurationPriced field.
func (o *AiGatewayModelPricingAttributes) SetDurationPriced(v bool) {
	o.DurationPriced = &v
}

// GetEffectiveSource returns the EffectiveSource field value if set, zero value otherwise.
func (o *AiGatewayModelPricingAttributes) GetEffectiveSource() AiGatewayModelPricingEffectiveSource {
	if o == nil || o.EffectiveSource == nil {
		var ret AiGatewayModelPricingEffectiveSource
		return ret
	}
	return *o.EffectiveSource
}

// GetEffectiveSourceOk returns a tuple with the EffectiveSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingAttributes) GetEffectiveSourceOk() (*AiGatewayModelPricingEffectiveSource, bool) {
	if o == nil || o.EffectiveSource == nil {
		return nil, false
	}
	return o.EffectiveSource, true
}

// HasEffectiveSource returns a boolean if a field has been set.
func (o *AiGatewayModelPricingAttributes) HasEffectiveSource() bool {
	return o != nil && o.EffectiveSource != nil
}

// SetEffectiveSource gets a reference to the given AiGatewayModelPricingEffectiveSource and assigns it to the EffectiveSource field.
func (o *AiGatewayModelPricingAttributes) SetEffectiveSource(v AiGatewayModelPricingEffectiveSource) {
	o.EffectiveSource = &v
}

// GetPricingConfig returns the PricingConfig field value if set, zero value otherwise.
func (o *AiGatewayModelPricingAttributes) GetPricingConfig() AiGatewayModelPricingConfig {
	if o == nil || o.PricingConfig == nil {
		var ret AiGatewayModelPricingConfig
		return ret
	}
	return *o.PricingConfig
}

// GetPricingConfigOk returns a tuple with the PricingConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingAttributes) GetPricingConfigOk() (*AiGatewayModelPricingConfig, bool) {
	if o == nil || o.PricingConfig == nil {
		return nil, false
	}
	return o.PricingConfig, true
}

// HasPricingConfig returns a boolean if a field has been set.
func (o *AiGatewayModelPricingAttributes) HasPricingConfig() bool {
	return o != nil && o.PricingConfig != nil
}

// SetPricingConfig gets a reference to the given AiGatewayModelPricingConfig and assigns it to the PricingConfig field.
func (o *AiGatewayModelPricingAttributes) SetPricingConfig(v AiGatewayModelPricingConfig) {
	o.PricingConfig = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AiGatewayModelPricingAttributes) GetStatus() AiGatewayModelPricingStatus {
	if o == nil || o.Status == nil {
		var ret AiGatewayModelPricingStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingAttributes) GetStatusOk() (*AiGatewayModelPricingStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiGatewayModelPricingAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given AiGatewayModelPricingStatus and assigns it to the Status field.
func (o *AiGatewayModelPricingAttributes) SetStatus(v AiGatewayModelPricingStatus) {
	o.Status = &v
}

// GetUnpricedDescription returns the UnpricedDescription field value if set, zero value otherwise.
func (o *AiGatewayModelPricingAttributes) GetUnpricedDescription() string {
	if o == nil || o.UnpricedDescription == nil {
		var ret string
		return ret
	}
	return *o.UnpricedDescription
}

// GetUnpricedDescriptionOk returns a tuple with the UnpricedDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingAttributes) GetUnpricedDescriptionOk() (*string, bool) {
	if o == nil || o.UnpricedDescription == nil {
		return nil, false
	}
	return o.UnpricedDescription, true
}

// HasUnpricedDescription returns a boolean if a field has been set.
func (o *AiGatewayModelPricingAttributes) HasUnpricedDescription() bool {
	return o != nil && o.UnpricedDescription != nil
}

// SetUnpricedDescription gets a reference to the given string and assigns it to the UnpricedDescription field.
func (o *AiGatewayModelPricingAttributes) SetUnpricedDescription(v string) {
	o.UnpricedDescription = &v
}

// GetUnpricedReason returns the UnpricedReason field value if set, zero value otherwise.
func (o *AiGatewayModelPricingAttributes) GetUnpricedReason() string {
	if o == nil || o.UnpricedReason == nil {
		var ret string
		return ret
	}
	return *o.UnpricedReason
}

// GetUnpricedReasonOk returns a tuple with the UnpricedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingAttributes) GetUnpricedReasonOk() (*string, bool) {
	if o == nil || o.UnpricedReason == nil {
		return nil, false
	}
	return o.UnpricedReason, true
}

// HasUnpricedReason returns a boolean if a field has been set.
func (o *AiGatewayModelPricingAttributes) HasUnpricedReason() bool {
	return o != nil && o.UnpricedReason != nil
}

// SetUnpricedReason gets a reference to the given string and assigns it to the UnpricedReason field.
func (o *AiGatewayModelPricingAttributes) SetUnpricedReason(v string) {
	o.UnpricedReason = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelPricingAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Authoritative != nil {
		toSerialize["authoritative"] = o.Authoritative
	}
	if o.AutomaticPricing != nil {
		toSerialize["automatic_pricing"] = o.AutomaticPricing
	}
	if o.BlocksBudgetedRequests != nil {
		toSerialize["blocks_budgeted_requests"] = o.BlocksBudgetedRequests
	}
	if o.CostMultiplier != nil {
		toSerialize["cost_multiplier"] = o.CostMultiplier
	}
	if o.DurationPriced != nil {
		toSerialize["duration_priced"] = o.DurationPriced
	}
	if o.EffectiveSource != nil {
		toSerialize["effective_source"] = o.EffectiveSource
	}
	if o.PricingConfig != nil {
		toSerialize["pricing_config"] = o.PricingConfig
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UnpricedDescription != nil {
		toSerialize["unpriced_description"] = o.UnpricedDescription
	}
	if o.UnpricedReason != nil {
		toSerialize["unpriced_reason"] = o.UnpricedReason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelPricingAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authoritative          *bool                                 `json:"authoritative,omitempty"`
		AutomaticPricing       *AiGatewayModelAutomaticPricing       `json:"automatic_pricing,omitempty"`
		BlocksBudgetedRequests *bool                                 `json:"blocks_budgeted_requests,omitempty"`
		CostMultiplier         *float64                              `json:"cost_multiplier,omitempty"`
		DurationPriced         *bool                                 `json:"duration_priced,omitempty"`
		EffectiveSource        *AiGatewayModelPricingEffectiveSource `json:"effective_source,omitempty"`
		PricingConfig          *AiGatewayModelPricingConfig          `json:"pricing_config,omitempty"`
		Status                 *AiGatewayModelPricingStatus          `json:"status,omitempty"`
		UnpricedDescription    *string                               `json:"unpriced_description,omitempty"`
		UnpricedReason         *string                               `json:"unpriced_reason,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"authoritative", "automatic_pricing", "blocks_budgeted_requests", "cost_multiplier", "duration_priced", "effective_source", "pricing_config", "status", "unpriced_description", "unpriced_reason"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Authoritative = all.Authoritative
	if all.AutomaticPricing != nil && all.AutomaticPricing.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AutomaticPricing = all.AutomaticPricing
	o.BlocksBudgetedRequests = all.BlocksBudgetedRequests
	o.CostMultiplier = all.CostMultiplier
	o.DurationPriced = all.DurationPriced
	if all.EffectiveSource != nil && !all.EffectiveSource.IsValid() {
		hasInvalidField = true
	} else {
		o.EffectiveSource = all.EffectiveSource
	}
	if all.PricingConfig != nil && all.PricingConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PricingConfig = all.PricingConfig
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.UnpricedDescription = all.UnpricedDescription
	o.UnpricedReason = all.UnpricedReason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
