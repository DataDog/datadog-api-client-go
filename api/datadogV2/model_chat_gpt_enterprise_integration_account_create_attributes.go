// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChatGPTEnterpriseIntegrationAccountCreateAttributes Writable attributes used to create a ChatGPT Enterprise integration account.
type ChatGPTEnterpriseIntegrationAccountCreateAttributes struct {
	// Authentication for creating the ChatGPT Enterprise integration account. Exactly one method is set.
	Authentication ChatGPTEnterpriseIntegrationAccountAuthenticationRequest `json:"authentication"`
	// Data Datadog collects from ChatGPT Enterprise, keyed by dataflow id. Each dataflow turns on a distinct kind of collection: set `enabled` to start or stop it. Defaults listed on each dataflow apply when the account is created; on update, omitted fields keep their current values. Every dataflow requires a ChatGPT Enterprise workspace with access to the OpenAI Costs API; a dataflow enabled without it is stored but collects no data.
	Dataflows *ChatGPTEnterpriseIntegrationDataflowsRequest `json:"dataflows,omitempty"`
	// Human-readable name of the ChatGPT Enterprise integration account.
	Name string `json:"name"`
	// Settings for creating the ChatGPT Enterprise integration account.
	Settings ChatGPTEnterpriseIntegrationAccountSettingsRequest `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChatGPTEnterpriseIntegrationAccountCreateAttributes instantiates a new ChatGPTEnterpriseIntegrationAccountCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChatGPTEnterpriseIntegrationAccountCreateAttributes(authentication ChatGPTEnterpriseIntegrationAccountAuthenticationRequest, name string, settings ChatGPTEnterpriseIntegrationAccountSettingsRequest) *ChatGPTEnterpriseIntegrationAccountCreateAttributes {
	this := ChatGPTEnterpriseIntegrationAccountCreateAttributes{}
	this.Authentication = authentication
	this.Name = name
	this.Settings = settings
	return &this
}

// NewChatGPTEnterpriseIntegrationAccountCreateAttributesWithDefaults instantiates a new ChatGPTEnterpriseIntegrationAccountCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChatGPTEnterpriseIntegrationAccountCreateAttributesWithDefaults() *ChatGPTEnterpriseIntegrationAccountCreateAttributes {
	this := ChatGPTEnterpriseIntegrationAccountCreateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value.
func (o *ChatGPTEnterpriseIntegrationAccountCreateAttributes) GetAuthentication() ChatGPTEnterpriseIntegrationAccountAuthenticationRequest {
	if o == nil {
		var ret ChatGPTEnterpriseIntegrationAccountAuthenticationRequest
		return ret
	}
	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *ChatGPTEnterpriseIntegrationAccountCreateAttributes) GetAuthenticationOk() (*ChatGPTEnterpriseIntegrationAccountAuthenticationRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value.
func (o *ChatGPTEnterpriseIntegrationAccountCreateAttributes) SetAuthentication(v ChatGPTEnterpriseIntegrationAccountAuthenticationRequest) {
	o.Authentication = v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *ChatGPTEnterpriseIntegrationAccountCreateAttributes) GetDataflows() ChatGPTEnterpriseIntegrationDataflowsRequest {
	if o == nil || o.Dataflows == nil {
		var ret ChatGPTEnterpriseIntegrationDataflowsRequest
		return ret
	}
	return *o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatGPTEnterpriseIntegrationAccountCreateAttributes) GetDataflowsOk() (*ChatGPTEnterpriseIntegrationDataflowsRequest, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *ChatGPTEnterpriseIntegrationAccountCreateAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given ChatGPTEnterpriseIntegrationDataflowsRequest and assigns it to the Dataflows field.
func (o *ChatGPTEnterpriseIntegrationAccountCreateAttributes) SetDataflows(v ChatGPTEnterpriseIntegrationDataflowsRequest) {
	o.Dataflows = &v
}

// GetName returns the Name field value.
func (o *ChatGPTEnterpriseIntegrationAccountCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ChatGPTEnterpriseIntegrationAccountCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ChatGPTEnterpriseIntegrationAccountCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value.
func (o *ChatGPTEnterpriseIntegrationAccountCreateAttributes) GetSettings() ChatGPTEnterpriseIntegrationAccountSettingsRequest {
	if o == nil {
		var ret ChatGPTEnterpriseIntegrationAccountSettingsRequest
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *ChatGPTEnterpriseIntegrationAccountCreateAttributes) GetSettingsOk() (*ChatGPTEnterpriseIntegrationAccountSettingsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *ChatGPTEnterpriseIntegrationAccountCreateAttributes) SetSettings(v ChatGPTEnterpriseIntegrationAccountSettingsRequest) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChatGPTEnterpriseIntegrationAccountCreateAttributes) MarshalJSON() ([]byte, error) {
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
func (o *ChatGPTEnterpriseIntegrationAccountCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *ChatGPTEnterpriseIntegrationAccountAuthenticationRequest `json:"authentication"`
		Dataflows      *ChatGPTEnterpriseIntegrationDataflowsRequest             `json:"dataflows,omitempty"`
		Name           *string                                                   `json:"name"`
		Settings       *ChatGPTEnterpriseIntegrationAccountSettingsRequest       `json:"settings"`
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
