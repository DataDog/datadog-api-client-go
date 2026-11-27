// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAIModelInferenceIntegrationAccountCreateAttributes Writable attributes used to create an OpenAI Model Inference integration account.
type OpenAIModelInferenceIntegrationAccountCreateAttributes struct {
	// Authentication for creating the OpenAI Model Inference integration account. Set exactly one method.
	Authentication OpenAIModelInferenceIntegrationAccountAuthenticationRequest `json:"authentication"`
	// Data Datadog collects from OpenAI Model Inference, keyed by dataflow ID. Each dataflow turns on a distinct type of collection. Set `enabled` to start or stop a dataflow. The default for each dataflow applies when the account is created. On update, omitted fields keep their current values.
	Dataflows *OpenAIModelInferenceIntegrationDataflowsRequest `json:"dataflows,omitempty"`
	// Human-readable name of the OpenAI Model Inference integration account.
	Name string `json:"name"`
	// Settings for creating the OpenAI Model Inference integration account.
	Settings OpenAIModelInferenceIntegrationAccountSettingsRequest `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpenAIModelInferenceIntegrationAccountCreateAttributes instantiates a new OpenAIModelInferenceIntegrationAccountCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAIModelInferenceIntegrationAccountCreateAttributes(authentication OpenAIModelInferenceIntegrationAccountAuthenticationRequest, name string, settings OpenAIModelInferenceIntegrationAccountSettingsRequest) *OpenAIModelInferenceIntegrationAccountCreateAttributes {
	this := OpenAIModelInferenceIntegrationAccountCreateAttributes{}
	this.Authentication = authentication
	this.Name = name
	this.Settings = settings
	return &this
}

// NewOpenAIModelInferenceIntegrationAccountCreateAttributesWithDefaults instantiates a new OpenAIModelInferenceIntegrationAccountCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAIModelInferenceIntegrationAccountCreateAttributesWithDefaults() *OpenAIModelInferenceIntegrationAccountCreateAttributes {
	this := OpenAIModelInferenceIntegrationAccountCreateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value.
func (o *OpenAIModelInferenceIntegrationAccountCreateAttributes) GetAuthentication() OpenAIModelInferenceIntegrationAccountAuthenticationRequest {
	if o == nil {
		var ret OpenAIModelInferenceIntegrationAccountAuthenticationRequest
		return ret
	}
	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *OpenAIModelInferenceIntegrationAccountCreateAttributes) GetAuthenticationOk() (*OpenAIModelInferenceIntegrationAccountAuthenticationRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value.
func (o *OpenAIModelInferenceIntegrationAccountCreateAttributes) SetAuthentication(v OpenAIModelInferenceIntegrationAccountAuthenticationRequest) {
	o.Authentication = v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *OpenAIModelInferenceIntegrationAccountCreateAttributes) GetDataflows() OpenAIModelInferenceIntegrationDataflowsRequest {
	if o == nil || o.Dataflows == nil {
		var ret OpenAIModelInferenceIntegrationDataflowsRequest
		return ret
	}
	return *o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAIModelInferenceIntegrationAccountCreateAttributes) GetDataflowsOk() (*OpenAIModelInferenceIntegrationDataflowsRequest, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *OpenAIModelInferenceIntegrationAccountCreateAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given OpenAIModelInferenceIntegrationDataflowsRequest and assigns it to the Dataflows field.
func (o *OpenAIModelInferenceIntegrationAccountCreateAttributes) SetDataflows(v OpenAIModelInferenceIntegrationDataflowsRequest) {
	o.Dataflows = &v
}

// GetName returns the Name field value.
func (o *OpenAIModelInferenceIntegrationAccountCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OpenAIModelInferenceIntegrationAccountCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OpenAIModelInferenceIntegrationAccountCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value.
func (o *OpenAIModelInferenceIntegrationAccountCreateAttributes) GetSettings() OpenAIModelInferenceIntegrationAccountSettingsRequest {
	if o == nil {
		var ret OpenAIModelInferenceIntegrationAccountSettingsRequest
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *OpenAIModelInferenceIntegrationAccountCreateAttributes) GetSettingsOk() (*OpenAIModelInferenceIntegrationAccountSettingsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *OpenAIModelInferenceIntegrationAccountCreateAttributes) SetSettings(v OpenAIModelInferenceIntegrationAccountSettingsRequest) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAIModelInferenceIntegrationAccountCreateAttributes) MarshalJSON() ([]byte, error) {
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
func (o *OpenAIModelInferenceIntegrationAccountCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *OpenAIModelInferenceIntegrationAccountAuthenticationRequest `json:"authentication"`
		Dataflows      *OpenAIModelInferenceIntegrationDataflowsRequest             `json:"dataflows,omitempty"`
		Name           *string                                                      `json:"name"`
		Settings       *OpenAIModelInferenceIntegrationAccountSettingsRequest       `json:"settings"`
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
