// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioIntegrationAccountCreateAttributes Writable attributes used to create a Twilio integration account.
type TwilioIntegrationAccountCreateAttributes struct {
	// Authentication for creating the Twilio integration account. Exactly one method is set.
	Authentication TwilioIntegrationAccountAuthenticationRequest `json:"authentication"`
	// Dataflows to configure on the Twilio integration account, keyed by dataflow id.
	Dataflows *TwilioIntegrationDataflowsRequest `json:"dataflows,omitempty"`
	// Human-readable name of the Twilio integration account.
	Name string `json:"name"`
	// Settings for creating the Twilio integration account.
	Settings TwilioIntegrationAccountSettingsRequest `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTwilioIntegrationAccountCreateAttributes instantiates a new TwilioIntegrationAccountCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTwilioIntegrationAccountCreateAttributes(authentication TwilioIntegrationAccountAuthenticationRequest, name string, settings TwilioIntegrationAccountSettingsRequest) *TwilioIntegrationAccountCreateAttributes {
	this := TwilioIntegrationAccountCreateAttributes{}
	this.Authentication = authentication
	this.Name = name
	this.Settings = settings
	return &this
}

// NewTwilioIntegrationAccountCreateAttributesWithDefaults instantiates a new TwilioIntegrationAccountCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTwilioIntegrationAccountCreateAttributesWithDefaults() *TwilioIntegrationAccountCreateAttributes {
	this := TwilioIntegrationAccountCreateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value.
func (o *TwilioIntegrationAccountCreateAttributes) GetAuthentication() TwilioIntegrationAccountAuthenticationRequest {
	if o == nil {
		var ret TwilioIntegrationAccountAuthenticationRequest
		return ret
	}
	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationAccountCreateAttributes) GetAuthenticationOk() (*TwilioIntegrationAccountAuthenticationRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value.
func (o *TwilioIntegrationAccountCreateAttributes) SetAuthentication(v TwilioIntegrationAccountAuthenticationRequest) {
	o.Authentication = v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *TwilioIntegrationAccountCreateAttributes) GetDataflows() TwilioIntegrationDataflowsRequest {
	if o == nil || o.Dataflows == nil {
		var ret TwilioIntegrationDataflowsRequest
		return ret
	}
	return *o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationAccountCreateAttributes) GetDataflowsOk() (*TwilioIntegrationDataflowsRequest, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *TwilioIntegrationAccountCreateAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given TwilioIntegrationDataflowsRequest and assigns it to the Dataflows field.
func (o *TwilioIntegrationAccountCreateAttributes) SetDataflows(v TwilioIntegrationDataflowsRequest) {
	o.Dataflows = &v
}

// GetName returns the Name field value.
func (o *TwilioIntegrationAccountCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationAccountCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *TwilioIntegrationAccountCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value.
func (o *TwilioIntegrationAccountCreateAttributes) GetSettings() TwilioIntegrationAccountSettingsRequest {
	if o == nil {
		var ret TwilioIntegrationAccountSettingsRequest
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationAccountCreateAttributes) GetSettingsOk() (*TwilioIntegrationAccountSettingsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *TwilioIntegrationAccountCreateAttributes) SetSettings(v TwilioIntegrationAccountSettingsRequest) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TwilioIntegrationAccountCreateAttributes) MarshalJSON() ([]byte, error) {
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
func (o *TwilioIntegrationAccountCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *TwilioIntegrationAccountAuthenticationRequest `json:"authentication"`
		Dataflows      *TwilioIntegrationDataflowsRequest             `json:"dataflows,omitempty"`
		Name           *string                                        `json:"name"`
		Settings       *TwilioIntegrationAccountSettingsRequest       `json:"settings"`
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
