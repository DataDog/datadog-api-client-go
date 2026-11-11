// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChatGPTEnterpriseIntegrationAccountResponseAttributes Attributes of a ChatGPT Enterprise integration account returned in responses.
type ChatGPTEnterpriseIntegrationAccountResponseAttributes struct {
	// Authentication configured on the ChatGPT Enterprise integration account.
	Authentication *ChatGPTEnterpriseIntegrationAccountAuthenticationResponse `json:"authentication,omitempty"`
	// Data Datadog collects from ChatGPT Enterprise, keyed by dataflow id.
	Dataflows *ChatGPTEnterpriseIntegrationDataflowsResponse `json:"dataflows,omitempty"`
	// Human-readable name of the ChatGPT Enterprise integration account.
	Name string `json:"name"`
	// Settings configured on the ChatGPT Enterprise integration account.
	Settings ChatGPTEnterpriseIntegrationAccountSettingsResponse `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChatGPTEnterpriseIntegrationAccountResponseAttributes instantiates a new ChatGPTEnterpriseIntegrationAccountResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChatGPTEnterpriseIntegrationAccountResponseAttributes(name string, settings ChatGPTEnterpriseIntegrationAccountSettingsResponse) *ChatGPTEnterpriseIntegrationAccountResponseAttributes {
	this := ChatGPTEnterpriseIntegrationAccountResponseAttributes{}
	this.Name = name
	this.Settings = settings
	return &this
}

// NewChatGPTEnterpriseIntegrationAccountResponseAttributesWithDefaults instantiates a new ChatGPTEnterpriseIntegrationAccountResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChatGPTEnterpriseIntegrationAccountResponseAttributesWithDefaults() *ChatGPTEnterpriseIntegrationAccountResponseAttributes {
	this := ChatGPTEnterpriseIntegrationAccountResponseAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *ChatGPTEnterpriseIntegrationAccountResponseAttributes) GetAuthentication() ChatGPTEnterpriseIntegrationAccountAuthenticationResponse {
	if o == nil || o.Authentication == nil {
		var ret ChatGPTEnterpriseIntegrationAccountAuthenticationResponse
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatGPTEnterpriseIntegrationAccountResponseAttributes) GetAuthenticationOk() (*ChatGPTEnterpriseIntegrationAccountAuthenticationResponse, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *ChatGPTEnterpriseIntegrationAccountResponseAttributes) HasAuthentication() bool {
	return o != nil && o.Authentication != nil
}

// SetAuthentication gets a reference to the given ChatGPTEnterpriseIntegrationAccountAuthenticationResponse and assigns it to the Authentication field.
func (o *ChatGPTEnterpriseIntegrationAccountResponseAttributes) SetAuthentication(v ChatGPTEnterpriseIntegrationAccountAuthenticationResponse) {
	o.Authentication = &v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *ChatGPTEnterpriseIntegrationAccountResponseAttributes) GetDataflows() ChatGPTEnterpriseIntegrationDataflowsResponse {
	if o == nil || o.Dataflows == nil {
		var ret ChatGPTEnterpriseIntegrationDataflowsResponse
		return ret
	}
	return *o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatGPTEnterpriseIntegrationAccountResponseAttributes) GetDataflowsOk() (*ChatGPTEnterpriseIntegrationDataflowsResponse, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *ChatGPTEnterpriseIntegrationAccountResponseAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given ChatGPTEnterpriseIntegrationDataflowsResponse and assigns it to the Dataflows field.
func (o *ChatGPTEnterpriseIntegrationAccountResponseAttributes) SetDataflows(v ChatGPTEnterpriseIntegrationDataflowsResponse) {
	o.Dataflows = &v
}

// GetName returns the Name field value.
func (o *ChatGPTEnterpriseIntegrationAccountResponseAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ChatGPTEnterpriseIntegrationAccountResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ChatGPTEnterpriseIntegrationAccountResponseAttributes) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value.
func (o *ChatGPTEnterpriseIntegrationAccountResponseAttributes) GetSettings() ChatGPTEnterpriseIntegrationAccountSettingsResponse {
	if o == nil {
		var ret ChatGPTEnterpriseIntegrationAccountSettingsResponse
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *ChatGPTEnterpriseIntegrationAccountResponseAttributes) GetSettingsOk() (*ChatGPTEnterpriseIntegrationAccountSettingsResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *ChatGPTEnterpriseIntegrationAccountResponseAttributes) SetSettings(v ChatGPTEnterpriseIntegrationAccountSettingsResponse) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChatGPTEnterpriseIntegrationAccountResponseAttributes) MarshalJSON() ([]byte, error) {
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
func (o *ChatGPTEnterpriseIntegrationAccountResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *ChatGPTEnterpriseIntegrationAccountAuthenticationResponse `json:"authentication,omitempty"`
		Dataflows      *ChatGPTEnterpriseIntegrationDataflowsResponse             `json:"dataflows,omitempty"`
		Name           *string                                                    `json:"name"`
		Settings       *ChatGPTEnterpriseIntegrationAccountSettingsResponse       `json:"settings"`
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
