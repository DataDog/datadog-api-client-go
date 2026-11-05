// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayProviderIntegrationConfig Provider-specific integration configuration.
type AiGatewayProviderIntegrationConfig struct {
	// AWS region used for Bedrock model discovery, routing, and request signing.
	BedrockRegion *string `json:"bedrock_region,omitempty"`
	// Azure AI Foundry project name, used to call the project's agents.
	FoundryProject *string `json:"foundry_project,omitempty"`
	// ARM resource ID of the Azure AI Foundry resource, such as `/subscriptions/<id>/resourceGroups/<group>/providers/Microsoft.CognitiveServices/accounts/<name>`. Required for `azure_foundry` providers.
	FoundryResourceId *string `json:"foundry_resource_id,omitempty"`
	// GCP location used for Gemini model discovery and credential resolution.
	Location *string `json:"location,omitempty"`
	// GCP project ID used for Gemini model discovery and credential resolution.
	ProjectId *string `json:"project_id,omitempty"`
	// Custom provider default for discovered and new manual models. Omit to use Chat Completions. An empty list is invalid.
	SupportedOperations []AiGatewayCustomProviderSupportedOperation `json:"supported_operations,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayProviderIntegrationConfig instantiates a new AiGatewayProviderIntegrationConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayProviderIntegrationConfig() *AiGatewayProviderIntegrationConfig {
	this := AiGatewayProviderIntegrationConfig{}
	return &this
}

// NewAiGatewayProviderIntegrationConfigWithDefaults instantiates a new AiGatewayProviderIntegrationConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayProviderIntegrationConfigWithDefaults() *AiGatewayProviderIntegrationConfig {
	this := AiGatewayProviderIntegrationConfig{}
	return &this
}

// GetBedrockRegion returns the BedrockRegion field value if set, zero value otherwise.
func (o *AiGatewayProviderIntegrationConfig) GetBedrockRegion() string {
	if o == nil || o.BedrockRegion == nil {
		var ret string
		return ret
	}
	return *o.BedrockRegion
}

// GetBedrockRegionOk returns a tuple with the BedrockRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderIntegrationConfig) GetBedrockRegionOk() (*string, bool) {
	if o == nil || o.BedrockRegion == nil {
		return nil, false
	}
	return o.BedrockRegion, true
}

// HasBedrockRegion returns a boolean if a field has been set.
func (o *AiGatewayProviderIntegrationConfig) HasBedrockRegion() bool {
	return o != nil && o.BedrockRegion != nil
}

// SetBedrockRegion gets a reference to the given string and assigns it to the BedrockRegion field.
func (o *AiGatewayProviderIntegrationConfig) SetBedrockRegion(v string) {
	o.BedrockRegion = &v
}

// GetFoundryProject returns the FoundryProject field value if set, zero value otherwise.
func (o *AiGatewayProviderIntegrationConfig) GetFoundryProject() string {
	if o == nil || o.FoundryProject == nil {
		var ret string
		return ret
	}
	return *o.FoundryProject
}

// GetFoundryProjectOk returns a tuple with the FoundryProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderIntegrationConfig) GetFoundryProjectOk() (*string, bool) {
	if o == nil || o.FoundryProject == nil {
		return nil, false
	}
	return o.FoundryProject, true
}

// HasFoundryProject returns a boolean if a field has been set.
func (o *AiGatewayProviderIntegrationConfig) HasFoundryProject() bool {
	return o != nil && o.FoundryProject != nil
}

// SetFoundryProject gets a reference to the given string and assigns it to the FoundryProject field.
func (o *AiGatewayProviderIntegrationConfig) SetFoundryProject(v string) {
	o.FoundryProject = &v
}

// GetFoundryResourceId returns the FoundryResourceId field value if set, zero value otherwise.
func (o *AiGatewayProviderIntegrationConfig) GetFoundryResourceId() string {
	if o == nil || o.FoundryResourceId == nil {
		var ret string
		return ret
	}
	return *o.FoundryResourceId
}

// GetFoundryResourceIdOk returns a tuple with the FoundryResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderIntegrationConfig) GetFoundryResourceIdOk() (*string, bool) {
	if o == nil || o.FoundryResourceId == nil {
		return nil, false
	}
	return o.FoundryResourceId, true
}

// HasFoundryResourceId returns a boolean if a field has been set.
func (o *AiGatewayProviderIntegrationConfig) HasFoundryResourceId() bool {
	return o != nil && o.FoundryResourceId != nil
}

// SetFoundryResourceId gets a reference to the given string and assigns it to the FoundryResourceId field.
func (o *AiGatewayProviderIntegrationConfig) SetFoundryResourceId(v string) {
	o.FoundryResourceId = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *AiGatewayProviderIntegrationConfig) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderIntegrationConfig) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *AiGatewayProviderIntegrationConfig) HasLocation() bool {
	return o != nil && o.Location != nil
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *AiGatewayProviderIntegrationConfig) SetLocation(v string) {
	o.Location = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *AiGatewayProviderIntegrationConfig) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderIntegrationConfig) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *AiGatewayProviderIntegrationConfig) HasProjectId() bool {
	return o != nil && o.ProjectId != nil
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *AiGatewayProviderIntegrationConfig) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetSupportedOperations returns the SupportedOperations field value if set, zero value otherwise.
func (o *AiGatewayProviderIntegrationConfig) GetSupportedOperations() []AiGatewayCustomProviderSupportedOperation {
	if o == nil || o.SupportedOperations == nil {
		var ret []AiGatewayCustomProviderSupportedOperation
		return ret
	}
	return o.SupportedOperations
}

// GetSupportedOperationsOk returns a tuple with the SupportedOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderIntegrationConfig) GetSupportedOperationsOk() (*[]AiGatewayCustomProviderSupportedOperation, bool) {
	if o == nil || o.SupportedOperations == nil {
		return nil, false
	}
	return &o.SupportedOperations, true
}

// HasSupportedOperations returns a boolean if a field has been set.
func (o *AiGatewayProviderIntegrationConfig) HasSupportedOperations() bool {
	return o != nil && o.SupportedOperations != nil
}

// SetSupportedOperations gets a reference to the given []AiGatewayCustomProviderSupportedOperation and assigns it to the SupportedOperations field.
func (o *AiGatewayProviderIntegrationConfig) SetSupportedOperations(v []AiGatewayCustomProviderSupportedOperation) {
	o.SupportedOperations = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayProviderIntegrationConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BedrockRegion != nil {
		toSerialize["bedrock_region"] = o.BedrockRegion
	}
	if o.FoundryProject != nil {
		toSerialize["foundry_project"] = o.FoundryProject
	}
	if o.FoundryResourceId != nil {
		toSerialize["foundry_resource_id"] = o.FoundryResourceId
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	if o.SupportedOperations != nil {
		toSerialize["supported_operations"] = o.SupportedOperations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayProviderIntegrationConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BedrockRegion       *string                                     `json:"bedrock_region,omitempty"`
		FoundryProject      *string                                     `json:"foundry_project,omitempty"`
		FoundryResourceId   *string                                     `json:"foundry_resource_id,omitempty"`
		Location            *string                                     `json:"location,omitempty"`
		ProjectId           *string                                     `json:"project_id,omitempty"`
		SupportedOperations []AiGatewayCustomProviderSupportedOperation `json:"supported_operations,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bedrock_region", "foundry_project", "foundry_resource_id", "location", "project_id", "supported_operations"})
	} else {
		return err
	}
	o.BedrockRegion = all.BedrockRegion
	o.FoundryProject = all.FoundryProject
	o.FoundryResourceId = all.FoundryResourceId
	o.Location = all.Location
	o.ProjectId = all.ProjectId
	o.SupportedOperations = all.SupportedOperations

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
