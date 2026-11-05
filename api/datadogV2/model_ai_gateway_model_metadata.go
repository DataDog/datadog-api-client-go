// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelMetadata Optional, typed enrichment for a model catalog entry. Missing properties remain unknown.
type AiGatewayModelMetadata struct {
	// Model features reported by the enrichment source. Omitted properties are unknown.
	Capabilities AiGatewayModelCapabilities `json:"capabilities"`
	// Human-readable model description.
	Description *string `json:"description,omitempty"`
	// Human-readable model name.
	DisplayName *string `json:"display_name,omitempty"`
	// Model family identifier.
	Family *string `json:"family,omitempty"`
	// Primary model kind. Decision models return a choice or score instead of generated text.
	Kind *AiGatewayModelMetadataKind `json:"kind,omitempty"`
	// Knowledge cutoff reported by the source.
	KnowledgeCutoff *string `json:"knowledge_cutoff,omitempty"`
	// Date the source record was last updated.
	LastUpdated *string `json:"last_updated,omitempty"`
	// Model lifecycle state.
	Lifecycle *AiGatewayModelMetadataLifecycle `json:"lifecycle,omitempty"`
	// Token limits reported by the enrichment source.
	Limits AiGatewayModelLimits `json:"limits"`
	// Supported model input and output modalities.
	Modalities AiGatewayModelModalities `json:"modalities"`
	// Whether model weights are openly available.
	OpenWeights *bool `json:"open_weights,omitempty"`
	// Supported reasoning controls. Absence means controls are unknown or unavailable.
	Reasoning *AiGatewayModelReasoning `json:"reasoning,omitempty"`
	// Provider-specific reference prices from the enrichment source. Omitted for shared canonical-model fallback matches. These do not determine gateway charges.
	ReferencePricing []AiGatewayModelReferencePrice `json:"reference_pricing,omitempty"`
	// Model release date.
	ReleaseDate *string `json:"release_date,omitempty"`
	// Provenance for model metadata enrichment.
	Source AiGatewayModelMetadataSource `json:"source"`
	// Completeness of the matched source record. Partial also identifies metadata from a base model for a dated snapshot or from a shared canonical-model fallback.
	Status AiGatewayModelMetadataStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelMetadata instantiates a new AiGatewayModelMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelMetadata(capabilities AiGatewayModelCapabilities, limits AiGatewayModelLimits, modalities AiGatewayModelModalities, source AiGatewayModelMetadataSource, status AiGatewayModelMetadataStatus) *AiGatewayModelMetadata {
	this := AiGatewayModelMetadata{}
	this.Capabilities = capabilities
	this.Limits = limits
	this.Modalities = modalities
	this.Source = source
	this.Status = status
	return &this
}

// NewAiGatewayModelMetadataWithDefaults instantiates a new AiGatewayModelMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelMetadataWithDefaults() *AiGatewayModelMetadata {
	this := AiGatewayModelMetadata{}
	return &this
}

// GetCapabilities returns the Capabilities field value.
func (o *AiGatewayModelMetadata) GetCapabilities() AiGatewayModelCapabilities {
	if o == nil {
		var ret AiGatewayModelCapabilities
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadata) GetCapabilitiesOk() (*AiGatewayModelCapabilities, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capabilities, true
}

// SetCapabilities sets field value.
func (o *AiGatewayModelMetadata) SetCapabilities(v AiGatewayModelCapabilities) {
	o.Capabilities = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AiGatewayModelMetadata) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadata) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AiGatewayModelMetadata) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AiGatewayModelMetadata) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AiGatewayModelMetadata) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadata) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AiGatewayModelMetadata) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AiGatewayModelMetadata) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *AiGatewayModelMetadata) GetFamily() string {
	if o == nil || o.Family == nil {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadata) GetFamilyOk() (*string, bool) {
	if o == nil || o.Family == nil {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *AiGatewayModelMetadata) HasFamily() bool {
	return o != nil && o.Family != nil
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *AiGatewayModelMetadata) SetFamily(v string) {
	o.Family = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *AiGatewayModelMetadata) GetKind() AiGatewayModelMetadataKind {
	if o == nil || o.Kind == nil {
		var ret AiGatewayModelMetadataKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadata) GetKindOk() (*AiGatewayModelMetadataKind, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *AiGatewayModelMetadata) HasKind() bool {
	return o != nil && o.Kind != nil
}

// SetKind gets a reference to the given AiGatewayModelMetadataKind and assigns it to the Kind field.
func (o *AiGatewayModelMetadata) SetKind(v AiGatewayModelMetadataKind) {
	o.Kind = &v
}

// GetKnowledgeCutoff returns the KnowledgeCutoff field value if set, zero value otherwise.
func (o *AiGatewayModelMetadata) GetKnowledgeCutoff() string {
	if o == nil || o.KnowledgeCutoff == nil {
		var ret string
		return ret
	}
	return *o.KnowledgeCutoff
}

// GetKnowledgeCutoffOk returns a tuple with the KnowledgeCutoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadata) GetKnowledgeCutoffOk() (*string, bool) {
	if o == nil || o.KnowledgeCutoff == nil {
		return nil, false
	}
	return o.KnowledgeCutoff, true
}

// HasKnowledgeCutoff returns a boolean if a field has been set.
func (o *AiGatewayModelMetadata) HasKnowledgeCutoff() bool {
	return o != nil && o.KnowledgeCutoff != nil
}

// SetKnowledgeCutoff gets a reference to the given string and assigns it to the KnowledgeCutoff field.
func (o *AiGatewayModelMetadata) SetKnowledgeCutoff(v string) {
	o.KnowledgeCutoff = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AiGatewayModelMetadata) GetLastUpdated() string {
	if o == nil || o.LastUpdated == nil {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadata) GetLastUpdatedOk() (*string, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AiGatewayModelMetadata) HasLastUpdated() bool {
	return o != nil && o.LastUpdated != nil
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *AiGatewayModelMetadata) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *AiGatewayModelMetadata) GetLifecycle() AiGatewayModelMetadataLifecycle {
	if o == nil || o.Lifecycle == nil {
		var ret AiGatewayModelMetadataLifecycle
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadata) GetLifecycleOk() (*AiGatewayModelMetadataLifecycle, bool) {
	if o == nil || o.Lifecycle == nil {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *AiGatewayModelMetadata) HasLifecycle() bool {
	return o != nil && o.Lifecycle != nil
}

// SetLifecycle gets a reference to the given AiGatewayModelMetadataLifecycle and assigns it to the Lifecycle field.
func (o *AiGatewayModelMetadata) SetLifecycle(v AiGatewayModelMetadataLifecycle) {
	o.Lifecycle = &v
}

// GetLimits returns the Limits field value.
func (o *AiGatewayModelMetadata) GetLimits() AiGatewayModelLimits {
	if o == nil {
		var ret AiGatewayModelLimits
		return ret
	}
	return o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadata) GetLimitsOk() (*AiGatewayModelLimits, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limits, true
}

// SetLimits sets field value.
func (o *AiGatewayModelMetadata) SetLimits(v AiGatewayModelLimits) {
	o.Limits = v
}

// GetModalities returns the Modalities field value.
func (o *AiGatewayModelMetadata) GetModalities() AiGatewayModelModalities {
	if o == nil {
		var ret AiGatewayModelModalities
		return ret
	}
	return o.Modalities
}

// GetModalitiesOk returns a tuple with the Modalities field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadata) GetModalitiesOk() (*AiGatewayModelModalities, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modalities, true
}

// SetModalities sets field value.
func (o *AiGatewayModelMetadata) SetModalities(v AiGatewayModelModalities) {
	o.Modalities = v
}

// GetOpenWeights returns the OpenWeights field value if set, zero value otherwise.
func (o *AiGatewayModelMetadata) GetOpenWeights() bool {
	if o == nil || o.OpenWeights == nil {
		var ret bool
		return ret
	}
	return *o.OpenWeights
}

// GetOpenWeightsOk returns a tuple with the OpenWeights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadata) GetOpenWeightsOk() (*bool, bool) {
	if o == nil || o.OpenWeights == nil {
		return nil, false
	}
	return o.OpenWeights, true
}

// HasOpenWeights returns a boolean if a field has been set.
func (o *AiGatewayModelMetadata) HasOpenWeights() bool {
	return o != nil && o.OpenWeights != nil
}

// SetOpenWeights gets a reference to the given bool and assigns it to the OpenWeights field.
func (o *AiGatewayModelMetadata) SetOpenWeights(v bool) {
	o.OpenWeights = &v
}

// GetReasoning returns the Reasoning field value if set, zero value otherwise.
func (o *AiGatewayModelMetadata) GetReasoning() AiGatewayModelReasoning {
	if o == nil || o.Reasoning == nil {
		var ret AiGatewayModelReasoning
		return ret
	}
	return *o.Reasoning
}

// GetReasoningOk returns a tuple with the Reasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadata) GetReasoningOk() (*AiGatewayModelReasoning, bool) {
	if o == nil || o.Reasoning == nil {
		return nil, false
	}
	return o.Reasoning, true
}

// HasReasoning returns a boolean if a field has been set.
func (o *AiGatewayModelMetadata) HasReasoning() bool {
	return o != nil && o.Reasoning != nil
}

// SetReasoning gets a reference to the given AiGatewayModelReasoning and assigns it to the Reasoning field.
func (o *AiGatewayModelMetadata) SetReasoning(v AiGatewayModelReasoning) {
	o.Reasoning = &v
}

// GetReferencePricing returns the ReferencePricing field value if set, zero value otherwise.
func (o *AiGatewayModelMetadata) GetReferencePricing() []AiGatewayModelReferencePrice {
	if o == nil || o.ReferencePricing == nil {
		var ret []AiGatewayModelReferencePrice
		return ret
	}
	return o.ReferencePricing
}

// GetReferencePricingOk returns a tuple with the ReferencePricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadata) GetReferencePricingOk() (*[]AiGatewayModelReferencePrice, bool) {
	if o == nil || o.ReferencePricing == nil {
		return nil, false
	}
	return &o.ReferencePricing, true
}

// HasReferencePricing returns a boolean if a field has been set.
func (o *AiGatewayModelMetadata) HasReferencePricing() bool {
	return o != nil && o.ReferencePricing != nil
}

// SetReferencePricing gets a reference to the given []AiGatewayModelReferencePrice and assigns it to the ReferencePricing field.
func (o *AiGatewayModelMetadata) SetReferencePricing(v []AiGatewayModelReferencePrice) {
	o.ReferencePricing = v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *AiGatewayModelMetadata) GetReleaseDate() string {
	if o == nil || o.ReleaseDate == nil {
		var ret string
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadata) GetReleaseDateOk() (*string, bool) {
	if o == nil || o.ReleaseDate == nil {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *AiGatewayModelMetadata) HasReleaseDate() bool {
	return o != nil && o.ReleaseDate != nil
}

// SetReleaseDate gets a reference to the given string and assigns it to the ReleaseDate field.
func (o *AiGatewayModelMetadata) SetReleaseDate(v string) {
	o.ReleaseDate = &v
}

// GetSource returns the Source field value.
func (o *AiGatewayModelMetadata) GetSource() AiGatewayModelMetadataSource {
	if o == nil {
		var ret AiGatewayModelMetadataSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadata) GetSourceOk() (*AiGatewayModelMetadataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *AiGatewayModelMetadata) SetSource(v AiGatewayModelMetadataSource) {
	o.Source = v
}

// GetStatus returns the Status field value.
func (o *AiGatewayModelMetadata) GetStatus() AiGatewayModelMetadataStatus {
	if o == nil {
		var ret AiGatewayModelMetadataStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadata) GetStatusOk() (*AiGatewayModelMetadataStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *AiGatewayModelMetadata) SetStatus(v AiGatewayModelMetadataStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["capabilities"] = o.Capabilities
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Family != nil {
		toSerialize["family"] = o.Family
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.KnowledgeCutoff != nil {
		toSerialize["knowledge_cutoff"] = o.KnowledgeCutoff
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.Lifecycle != nil {
		toSerialize["lifecycle"] = o.Lifecycle
	}
	toSerialize["limits"] = o.Limits
	toSerialize["modalities"] = o.Modalities
	if o.OpenWeights != nil {
		toSerialize["open_weights"] = o.OpenWeights
	}
	if o.Reasoning != nil {
		toSerialize["reasoning"] = o.Reasoning
	}
	if o.ReferencePricing != nil {
		toSerialize["reference_pricing"] = o.ReferencePricing
	}
	if o.ReleaseDate != nil {
		toSerialize["release_date"] = o.ReleaseDate
	}
	toSerialize["source"] = o.Source
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Capabilities     *AiGatewayModelCapabilities      `json:"capabilities"`
		Description      *string                          `json:"description,omitempty"`
		DisplayName      *string                          `json:"display_name,omitempty"`
		Family           *string                          `json:"family,omitempty"`
		Kind             *AiGatewayModelMetadataKind      `json:"kind,omitempty"`
		KnowledgeCutoff  *string                          `json:"knowledge_cutoff,omitempty"`
		LastUpdated      *string                          `json:"last_updated,omitempty"`
		Lifecycle        *AiGatewayModelMetadataLifecycle `json:"lifecycle,omitempty"`
		Limits           *AiGatewayModelLimits            `json:"limits"`
		Modalities       *AiGatewayModelModalities        `json:"modalities"`
		OpenWeights      *bool                            `json:"open_weights,omitempty"`
		Reasoning        *AiGatewayModelReasoning         `json:"reasoning,omitempty"`
		ReferencePricing []AiGatewayModelReferencePrice   `json:"reference_pricing,omitempty"`
		ReleaseDate      *string                          `json:"release_date,omitempty"`
		Source           *AiGatewayModelMetadataSource    `json:"source"`
		Status           *AiGatewayModelMetadataStatus    `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Capabilities == nil {
		return fmt.Errorf("required field capabilities missing")
	}
	if all.Limits == nil {
		return fmt.Errorf("required field limits missing")
	}
	if all.Modalities == nil {
		return fmt.Errorf("required field modalities missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"capabilities", "description", "display_name", "family", "kind", "knowledge_cutoff", "last_updated", "lifecycle", "limits", "modalities", "open_weights", "reasoning", "reference_pricing", "release_date", "source", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Capabilities.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Capabilities = *all.Capabilities
	o.Description = all.Description
	o.DisplayName = all.DisplayName
	o.Family = all.Family
	if all.Kind != nil && !all.Kind.IsValid() {
		hasInvalidField = true
	} else {
		o.Kind = all.Kind
	}
	o.KnowledgeCutoff = all.KnowledgeCutoff
	o.LastUpdated = all.LastUpdated
	if all.Lifecycle != nil && !all.Lifecycle.IsValid() {
		hasInvalidField = true
	} else {
		o.Lifecycle = all.Lifecycle
	}
	if all.Limits.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Limits = *all.Limits
	if all.Modalities.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Modalities = *all.Modalities
	o.OpenWeights = all.OpenWeights
	if all.Reasoning != nil && all.Reasoning.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Reasoning = all.Reasoning
	o.ReferencePricing = all.ReferencePricing
	o.ReleaseDate = all.ReleaseDate
	if all.Source.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Source = *all.Source
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
