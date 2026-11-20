// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAILLMEvaluationsIntegrationAccountCreateAttributes Writable attributes used to create an OpenAI LLM Evaluations integration account.
type OpenAILLMEvaluationsIntegrationAccountCreateAttributes struct {
	// Authentication for creating the OpenAI LLM Evaluations integration account. Exactly one method is set.
	Authentication OpenAILLMEvaluationsIntegrationAccountAuthenticationRequest `json:"authentication"`
	// Data Datadog collects from OpenAI LLM Evaluations, keyed by dataflow ID. Each dataflow turns on a distinct type of collection. Set `enabled` to start or stop a dataflow. The default for each dataflow applies when the account is created. On update, omitted fields keep their current values.
	Dataflows *OpenAILLMEvaluationsIntegrationDataflowsRequest `json:"dataflows,omitempty"`
	// Human-readable name of the OpenAI LLM Evaluations integration account.
	Name string `json:"name"`
	// Settings for creating the OpenAI LLM Evaluations integration account.
	Settings OpenAILLMEvaluationsIntegrationAccountSettingsRequest `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpenAILLMEvaluationsIntegrationAccountCreateAttributes instantiates a new OpenAILLMEvaluationsIntegrationAccountCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAILLMEvaluationsIntegrationAccountCreateAttributes(authentication OpenAILLMEvaluationsIntegrationAccountAuthenticationRequest, name string, settings OpenAILLMEvaluationsIntegrationAccountSettingsRequest) *OpenAILLMEvaluationsIntegrationAccountCreateAttributes {
	this := OpenAILLMEvaluationsIntegrationAccountCreateAttributes{}
	this.Authentication = authentication
	this.Name = name
	this.Settings = settings
	return &this
}

// NewOpenAILLMEvaluationsIntegrationAccountCreateAttributesWithDefaults instantiates a new OpenAILLMEvaluationsIntegrationAccountCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAILLMEvaluationsIntegrationAccountCreateAttributesWithDefaults() *OpenAILLMEvaluationsIntegrationAccountCreateAttributes {
	this := OpenAILLMEvaluationsIntegrationAccountCreateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value.
func (o *OpenAILLMEvaluationsIntegrationAccountCreateAttributes) GetAuthentication() OpenAILLMEvaluationsIntegrationAccountAuthenticationRequest {
	if o == nil {
		var ret OpenAILLMEvaluationsIntegrationAccountAuthenticationRequest
		return ret
	}
	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *OpenAILLMEvaluationsIntegrationAccountCreateAttributes) GetAuthenticationOk() (*OpenAILLMEvaluationsIntegrationAccountAuthenticationRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value.
func (o *OpenAILLMEvaluationsIntegrationAccountCreateAttributes) SetAuthentication(v OpenAILLMEvaluationsIntegrationAccountAuthenticationRequest) {
	o.Authentication = v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *OpenAILLMEvaluationsIntegrationAccountCreateAttributes) GetDataflows() OpenAILLMEvaluationsIntegrationDataflowsRequest {
	if o == nil || o.Dataflows == nil {
		var ret OpenAILLMEvaluationsIntegrationDataflowsRequest
		return ret
	}
	return *o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAILLMEvaluationsIntegrationAccountCreateAttributes) GetDataflowsOk() (*OpenAILLMEvaluationsIntegrationDataflowsRequest, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *OpenAILLMEvaluationsIntegrationAccountCreateAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given OpenAILLMEvaluationsIntegrationDataflowsRequest and assigns it to the Dataflows field.
func (o *OpenAILLMEvaluationsIntegrationAccountCreateAttributes) SetDataflows(v OpenAILLMEvaluationsIntegrationDataflowsRequest) {
	o.Dataflows = &v
}

// GetName returns the Name field value.
func (o *OpenAILLMEvaluationsIntegrationAccountCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OpenAILLMEvaluationsIntegrationAccountCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OpenAILLMEvaluationsIntegrationAccountCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value.
func (o *OpenAILLMEvaluationsIntegrationAccountCreateAttributes) GetSettings() OpenAILLMEvaluationsIntegrationAccountSettingsRequest {
	if o == nil {
		var ret OpenAILLMEvaluationsIntegrationAccountSettingsRequest
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *OpenAILLMEvaluationsIntegrationAccountCreateAttributes) GetSettingsOk() (*OpenAILLMEvaluationsIntegrationAccountSettingsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *OpenAILLMEvaluationsIntegrationAccountCreateAttributes) SetSettings(v OpenAILLMEvaluationsIntegrationAccountSettingsRequest) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAILLMEvaluationsIntegrationAccountCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["authentication"] = o.Authentication
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
func (o *OpenAILLMEvaluationsIntegrationAccountCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *OpenAILLMEvaluationsIntegrationAccountAuthenticationRequest `json:"authentication"`
		Dataflows      *OpenAILLMEvaluationsIntegrationDataflowsRequest             `json:"dataflows,omitempty"`
		Name           *string                                                      `json:"name"`
		Settings       *OpenAILLMEvaluationsIntegrationAccountSettingsRequest       `json:"settings"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Authentication == nil {
		return fmt.Errorf("required field authentication missing")
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
	o.Authentication = *all.Authentication
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
