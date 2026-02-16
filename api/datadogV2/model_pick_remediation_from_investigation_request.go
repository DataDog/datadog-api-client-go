// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PickRemediationFromInvestigationRequest
type PickRemediationFromInvestigationRequest struct {
	// The client type for action filtering.
	Client *ClientType `json:"client,omitempty"`
	// List of integrations to filter actions by.
	Integrations []string `json:"integrations,omitempty"`
	// The investigation text to extract remediation keywords from.
	Investigation string `json:"investigation"`
	// The number of keyword variants to generate.
	NumberOfKeywordVariants *int64 `json:"number_of_keyword_variants,omitempty"`
	// The number of relevant actions to return per keyword.
	NumberOfRelevantActions int64 `json:"number_of_relevant_actions"`
	// The stability level for action filtering.
	Stability *StabilityLevel `json:"stability,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPickRemediationFromInvestigationRequest instantiates a new PickRemediationFromInvestigationRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPickRemediationFromInvestigationRequest(investigation string, numberOfRelevantActions int64) *PickRemediationFromInvestigationRequest {
	this := PickRemediationFromInvestigationRequest{}
	this.Investigation = investigation
	this.NumberOfRelevantActions = numberOfRelevantActions
	return &this
}

// NewPickRemediationFromInvestigationRequestWithDefaults instantiates a new PickRemediationFromInvestigationRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPickRemediationFromInvestigationRequestWithDefaults() *PickRemediationFromInvestigationRequest {
	this := PickRemediationFromInvestigationRequest{}
	return &this
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *PickRemediationFromInvestigationRequest) GetClient() ClientType {
	if o == nil || o.Client == nil {
		var ret ClientType
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PickRemediationFromInvestigationRequest) GetClientOk() (*ClientType, bool) {
	if o == nil || o.Client == nil {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *PickRemediationFromInvestigationRequest) HasClient() bool {
	return o != nil && o.Client != nil
}

// SetClient gets a reference to the given ClientType and assigns it to the Client field.
func (o *PickRemediationFromInvestigationRequest) SetClient(v ClientType) {
	o.Client = &v
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise.
func (o *PickRemediationFromInvestigationRequest) GetIntegrations() []string {
	if o == nil || o.Integrations == nil {
		var ret []string
		return ret
	}
	return o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PickRemediationFromInvestigationRequest) GetIntegrationsOk() (*[]string, bool) {
	if o == nil || o.Integrations == nil {
		return nil, false
	}
	return &o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *PickRemediationFromInvestigationRequest) HasIntegrations() bool {
	return o != nil && o.Integrations != nil
}

// SetIntegrations gets a reference to the given []string and assigns it to the Integrations field.
func (o *PickRemediationFromInvestigationRequest) SetIntegrations(v []string) {
	o.Integrations = v
}

// GetInvestigation returns the Investigation field value.
func (o *PickRemediationFromInvestigationRequest) GetInvestigation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Investigation
}

// GetInvestigationOk returns a tuple with the Investigation field value
// and a boolean to check if the value has been set.
func (o *PickRemediationFromInvestigationRequest) GetInvestigationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Investigation, true
}

// SetInvestigation sets field value.
func (o *PickRemediationFromInvestigationRequest) SetInvestigation(v string) {
	o.Investigation = v
}

// GetNumberOfKeywordVariants returns the NumberOfKeywordVariants field value if set, zero value otherwise.
func (o *PickRemediationFromInvestigationRequest) GetNumberOfKeywordVariants() int64 {
	if o == nil || o.NumberOfKeywordVariants == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfKeywordVariants
}

// GetNumberOfKeywordVariantsOk returns a tuple with the NumberOfKeywordVariants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PickRemediationFromInvestigationRequest) GetNumberOfKeywordVariantsOk() (*int64, bool) {
	if o == nil || o.NumberOfKeywordVariants == nil {
		return nil, false
	}
	return o.NumberOfKeywordVariants, true
}

// HasNumberOfKeywordVariants returns a boolean if a field has been set.
func (o *PickRemediationFromInvestigationRequest) HasNumberOfKeywordVariants() bool {
	return o != nil && o.NumberOfKeywordVariants != nil
}

// SetNumberOfKeywordVariants gets a reference to the given int64 and assigns it to the NumberOfKeywordVariants field.
func (o *PickRemediationFromInvestigationRequest) SetNumberOfKeywordVariants(v int64) {
	o.NumberOfKeywordVariants = &v
}

// GetNumberOfRelevantActions returns the NumberOfRelevantActions field value.
func (o *PickRemediationFromInvestigationRequest) GetNumberOfRelevantActions() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.NumberOfRelevantActions
}

// GetNumberOfRelevantActionsOk returns a tuple with the NumberOfRelevantActions field value
// and a boolean to check if the value has been set.
func (o *PickRemediationFromInvestigationRequest) GetNumberOfRelevantActionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfRelevantActions, true
}

// SetNumberOfRelevantActions sets field value.
func (o *PickRemediationFromInvestigationRequest) SetNumberOfRelevantActions(v int64) {
	o.NumberOfRelevantActions = v
}

// GetStability returns the Stability field value if set, zero value otherwise.
func (o *PickRemediationFromInvestigationRequest) GetStability() StabilityLevel {
	if o == nil || o.Stability == nil {
		var ret StabilityLevel
		return ret
	}
	return *o.Stability
}

// GetStabilityOk returns a tuple with the Stability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PickRemediationFromInvestigationRequest) GetStabilityOk() (*StabilityLevel, bool) {
	if o == nil || o.Stability == nil {
		return nil, false
	}
	return o.Stability, true
}

// HasStability returns a boolean if a field has been set.
func (o *PickRemediationFromInvestigationRequest) HasStability() bool {
	return o != nil && o.Stability != nil
}

// SetStability gets a reference to the given StabilityLevel and assigns it to the Stability field.
func (o *PickRemediationFromInvestigationRequest) SetStability(v StabilityLevel) {
	o.Stability = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PickRemediationFromInvestigationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Client != nil {
		toSerialize["client"] = o.Client
	}
	if o.Integrations != nil {
		toSerialize["integrations"] = o.Integrations
	}
	toSerialize["investigation"] = o.Investigation
	if o.NumberOfKeywordVariants != nil {
		toSerialize["number_of_keyword_variants"] = o.NumberOfKeywordVariants
	}
	toSerialize["number_of_relevant_actions"] = o.NumberOfRelevantActions
	if o.Stability != nil {
		toSerialize["stability"] = o.Stability
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PickRemediationFromInvestigationRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Client                  *ClientType     `json:"client,omitempty"`
		Integrations            []string        `json:"integrations,omitempty"`
		Investigation           *string         `json:"investigation"`
		NumberOfKeywordVariants *int64          `json:"number_of_keyword_variants,omitempty"`
		NumberOfRelevantActions *int64          `json:"number_of_relevant_actions"`
		Stability               *StabilityLevel `json:"stability,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Investigation == nil {
		return fmt.Errorf("required field investigation missing")
	}
	if all.NumberOfRelevantActions == nil {
		return fmt.Errorf("required field number_of_relevant_actions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"client", "integrations", "investigation", "number_of_keyword_variants", "number_of_relevant_actions", "stability"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Client != nil && !all.Client.IsValid() {
		hasInvalidField = true
	} else {
		o.Client = all.Client
	}
	o.Integrations = all.Integrations
	o.Investigation = *all.Investigation
	o.NumberOfKeywordVariants = all.NumberOfKeywordVariants
	o.NumberOfRelevantActions = *all.NumberOfRelevantActions
	if all.Stability != nil && !all.Stability.IsValid() {
		hasInvalidField = true
	} else {
		o.Stability = all.Stability
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
