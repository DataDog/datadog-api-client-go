// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAILLMEvaluationsIntegrationAccountResponseAttributes Attributes of an OpenAI LLM Evaluations integration account returned in responses.
type OpenAILLMEvaluationsIntegrationAccountResponseAttributes struct {
	// Authentication configured on the OpenAI LLM Evaluations integration account.
	Authentication *OpenAILLMEvaluationsIntegrationAccountAuthenticationResponse `json:"authentication,omitempty"`
	// Data Datadog collects from OpenAI LLM Evaluations, keyed by dataflow ID.
	Dataflows *OpenAILLMEvaluationsIntegrationDataflowsResponse `json:"dataflows,omitempty"`
	// Human-readable name of the OpenAI LLM Evaluations integration account.
	Name string `json:"name"`
	// Settings configured on the OpenAI LLM Evaluations integration account.
	Settings OpenAILLMEvaluationsIntegrationAccountSettingsResponse `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpenAILLMEvaluationsIntegrationAccountResponseAttributes instantiates a new OpenAILLMEvaluationsIntegrationAccountResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAILLMEvaluationsIntegrationAccountResponseAttributes(name string, settings OpenAILLMEvaluationsIntegrationAccountSettingsResponse) *OpenAILLMEvaluationsIntegrationAccountResponseAttributes {
	this := OpenAILLMEvaluationsIntegrationAccountResponseAttributes{}
	this.Name = name
	this.Settings = settings
	return &this
}

// NewOpenAILLMEvaluationsIntegrationAccountResponseAttributesWithDefaults instantiates a new OpenAILLMEvaluationsIntegrationAccountResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAILLMEvaluationsIntegrationAccountResponseAttributesWithDefaults() *OpenAILLMEvaluationsIntegrationAccountResponseAttributes {
	this := OpenAILLMEvaluationsIntegrationAccountResponseAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *OpenAILLMEvaluationsIntegrationAccountResponseAttributes) GetAuthentication() OpenAILLMEvaluationsIntegrationAccountAuthenticationResponse {
	if o == nil || o.Authentication == nil {
		var ret OpenAILLMEvaluationsIntegrationAccountAuthenticationResponse
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAILLMEvaluationsIntegrationAccountResponseAttributes) GetAuthenticationOk() (*OpenAILLMEvaluationsIntegrationAccountAuthenticationResponse, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *OpenAILLMEvaluationsIntegrationAccountResponseAttributes) HasAuthentication() bool {
	return o != nil && o.Authentication != nil
}

// SetAuthentication gets a reference to the given OpenAILLMEvaluationsIntegrationAccountAuthenticationResponse and assigns it to the Authentication field.
func (o *OpenAILLMEvaluationsIntegrationAccountResponseAttributes) SetAuthentication(v OpenAILLMEvaluationsIntegrationAccountAuthenticationResponse) {
	o.Authentication = &v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *OpenAILLMEvaluationsIntegrationAccountResponseAttributes) GetDataflows() OpenAILLMEvaluationsIntegrationDataflowsResponse {
	if o == nil || o.Dataflows == nil {
		var ret OpenAILLMEvaluationsIntegrationDataflowsResponse
		return ret
	}
	return *o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAILLMEvaluationsIntegrationAccountResponseAttributes) GetDataflowsOk() (*OpenAILLMEvaluationsIntegrationDataflowsResponse, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *OpenAILLMEvaluationsIntegrationAccountResponseAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given OpenAILLMEvaluationsIntegrationDataflowsResponse and assigns it to the Dataflows field.
func (o *OpenAILLMEvaluationsIntegrationAccountResponseAttributes) SetDataflows(v OpenAILLMEvaluationsIntegrationDataflowsResponse) {
	o.Dataflows = &v
}

// GetName returns the Name field value.
func (o *OpenAILLMEvaluationsIntegrationAccountResponseAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OpenAILLMEvaluationsIntegrationAccountResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OpenAILLMEvaluationsIntegrationAccountResponseAttributes) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value.
func (o *OpenAILLMEvaluationsIntegrationAccountResponseAttributes) GetSettings() OpenAILLMEvaluationsIntegrationAccountSettingsResponse {
	if o == nil {
		var ret OpenAILLMEvaluationsIntegrationAccountSettingsResponse
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *OpenAILLMEvaluationsIntegrationAccountResponseAttributes) GetSettingsOk() (*OpenAILLMEvaluationsIntegrationAccountSettingsResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *OpenAILLMEvaluationsIntegrationAccountResponseAttributes) SetSettings(v OpenAILLMEvaluationsIntegrationAccountSettingsResponse) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAILLMEvaluationsIntegrationAccountResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Authentication != nil {
		toSerialize["authentication"] = o.Authentication
	}
	if o.Dataflows != nil {
		toSerialize["dataflows"] = o.Dataflows
	}
	toSerialize["name"] = o.Name
	toSerialize["settings"] = o.Settings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpenAILLMEvaluationsIntegrationAccountResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *OpenAILLMEvaluationsIntegrationAccountAuthenticationResponse `json:"authentication,omitempty"`
		Dataflows      *OpenAILLMEvaluationsIntegrationDataflowsResponse             `json:"dataflows,omitempty"`
		Name           *string                                                       `json:"name"`
		Settings       *OpenAILLMEvaluationsIntegrationAccountSettingsResponse       `json:"settings"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Settings == nil {
		return fmt.Errorf("required field settings missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"authentication", "dataflows", "name", "settings"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Authentication = all.Authentication
	if all.Dataflows != nil && all.Dataflows.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Dataflows = all.Dataflows
	o.Name = *all.Name
	if all.Settings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Settings = *all.Settings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
