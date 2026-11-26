// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAIModelInferenceIntegrationAccountResponseAttributes Attributes of an OpenAI Model Inference integration account returned in responses.
type OpenAIModelInferenceIntegrationAccountResponseAttributes struct {
	// Authentication configured on the OpenAI Model Inference integration account.
	Authentication *OpenAIModelInferenceIntegrationAccountAuthenticationResponse `json:"authentication,omitempty"`
	// Data Datadog collects from OpenAI Model Inference, keyed by dataflow ID.
	Dataflows *OpenAIModelInferenceIntegrationDataflowsResponse `json:"dataflows,omitempty"`
	// Human-readable name of the OpenAI Model Inference integration account.
	Name string `json:"name"`
	// Settings configured on the OpenAI Model Inference integration account.
	Settings OpenAIModelInferenceIntegrationAccountSettingsResponse `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpenAIModelInferenceIntegrationAccountResponseAttributes instantiates a new OpenAIModelInferenceIntegrationAccountResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAIModelInferenceIntegrationAccountResponseAttributes(name string, settings OpenAIModelInferenceIntegrationAccountSettingsResponse) *OpenAIModelInferenceIntegrationAccountResponseAttributes {
	this := OpenAIModelInferenceIntegrationAccountResponseAttributes{}
	this.Name = name
	this.Settings = settings
	return &this
}

// NewOpenAIModelInferenceIntegrationAccountResponseAttributesWithDefaults instantiates a new OpenAIModelInferenceIntegrationAccountResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAIModelInferenceIntegrationAccountResponseAttributesWithDefaults() *OpenAIModelInferenceIntegrationAccountResponseAttributes {
	this := OpenAIModelInferenceIntegrationAccountResponseAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *OpenAIModelInferenceIntegrationAccountResponseAttributes) GetAuthentication() OpenAIModelInferenceIntegrationAccountAuthenticationResponse {
	if o == nil || o.Authentication == nil {
		var ret OpenAIModelInferenceIntegrationAccountAuthenticationResponse
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAIModelInferenceIntegrationAccountResponseAttributes) GetAuthenticationOk() (*OpenAIModelInferenceIntegrationAccountAuthenticationResponse, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *OpenAIModelInferenceIntegrationAccountResponseAttributes) HasAuthentication() bool {
	return o != nil && o.Authentication != nil
}

// SetAuthentication gets a reference to the given OpenAIModelInferenceIntegrationAccountAuthenticationResponse and assigns it to the Authentication field.
func (o *OpenAIModelInferenceIntegrationAccountResponseAttributes) SetAuthentication(v OpenAIModelInferenceIntegrationAccountAuthenticationResponse) {
	o.Authentication = &v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *OpenAIModelInferenceIntegrationAccountResponseAttributes) GetDataflows() OpenAIModelInferenceIntegrationDataflowsResponse {
	if o == nil || o.Dataflows == nil {
		var ret OpenAIModelInferenceIntegrationDataflowsResponse
		return ret
	}
	return *o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAIModelInferenceIntegrationAccountResponseAttributes) GetDataflowsOk() (*OpenAIModelInferenceIntegrationDataflowsResponse, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *OpenAIModelInferenceIntegrationAccountResponseAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given OpenAIModelInferenceIntegrationDataflowsResponse and assigns it to the Dataflows field.
func (o *OpenAIModelInferenceIntegrationAccountResponseAttributes) SetDataflows(v OpenAIModelInferenceIntegrationDataflowsResponse) {
	o.Dataflows = &v
}

// GetName returns the Name field value.
func (o *OpenAIModelInferenceIntegrationAccountResponseAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OpenAIModelInferenceIntegrationAccountResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OpenAIModelInferenceIntegrationAccountResponseAttributes) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value.
func (o *OpenAIModelInferenceIntegrationAccountResponseAttributes) GetSettings() OpenAIModelInferenceIntegrationAccountSettingsResponse {
	if o == nil {
		var ret OpenAIModelInferenceIntegrationAccountSettingsResponse
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *OpenAIModelInferenceIntegrationAccountResponseAttributes) GetSettingsOk() (*OpenAIModelInferenceIntegrationAccountSettingsResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *OpenAIModelInferenceIntegrationAccountResponseAttributes) SetSettings(v OpenAIModelInferenceIntegrationAccountSettingsResponse) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAIModelInferenceIntegrationAccountResponseAttributes) MarshalJSON() ([]byte, error) {
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
func (o *OpenAIModelInferenceIntegrationAccountResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *OpenAIModelInferenceIntegrationAccountAuthenticationResponse `json:"authentication,omitempty"`
		Dataflows      *OpenAIModelInferenceIntegrationDataflowsResponse             `json:"dataflows,omitempty"`
		Name           *string                                                       `json:"name"`
		Settings       *OpenAIModelInferenceIntegrationAccountSettingsResponse       `json:"settings"`
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
