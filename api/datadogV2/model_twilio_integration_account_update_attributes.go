// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioIntegrationAccountUpdateAttributes Writable attributes used to update a Twilio integration account. Every field is optional; only the fields provided are changed. When `dataflows` is provided, only the dataflow ids included in the request are modified; dataflows omitted from the map keep their current configuration.
type TwilioIntegrationAccountUpdateAttributes struct {
	// Authentication for updating the Twilio integration account. Exactly one method is set.
	Authentication *TwilioIntegrationAccountAuthenticationUpdate `json:"authentication,omitempty"`
	// Dataflows to configure on the Twilio integration account, keyed by dataflow id.
	Dataflows *TwilioIntegrationDataflowsRequest `json:"dataflows,omitempty"`
	// Human-readable name of the Twilio integration account.
	Name *string `json:"name,omitempty"`
	// Settings for updating the Twilio integration account. Only the fields provided are changed.
	Settings *TwilioIntegrationAccountSettingsUpdate `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTwilioIntegrationAccountUpdateAttributes instantiates a new TwilioIntegrationAccountUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTwilioIntegrationAccountUpdateAttributes() *TwilioIntegrationAccountUpdateAttributes {
	this := TwilioIntegrationAccountUpdateAttributes{}
	return &this
}

// NewTwilioIntegrationAccountUpdateAttributesWithDefaults instantiates a new TwilioIntegrationAccountUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTwilioIntegrationAccountUpdateAttributesWithDefaults() *TwilioIntegrationAccountUpdateAttributes {
	this := TwilioIntegrationAccountUpdateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *TwilioIntegrationAccountUpdateAttributes) GetAuthentication() TwilioIntegrationAccountAuthenticationUpdate {
	if o == nil || o.Authentication == nil {
		var ret TwilioIntegrationAccountAuthenticationUpdate
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationAccountUpdateAttributes) GetAuthenticationOk() (*TwilioIntegrationAccountAuthenticationUpdate, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *TwilioIntegrationAccountUpdateAttributes) HasAuthentication() bool {
	return o != nil && o.Authentication != nil
}

// SetAuthentication gets a reference to the given TwilioIntegrationAccountAuthenticationUpdate and assigns it to the Authentication field.
func (o *TwilioIntegrationAccountUpdateAttributes) SetAuthentication(v TwilioIntegrationAccountAuthenticationUpdate) {
	o.Authentication = &v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *TwilioIntegrationAccountUpdateAttributes) GetDataflows() TwilioIntegrationDataflowsRequest {
	if o == nil || o.Dataflows == nil {
		var ret TwilioIntegrationDataflowsRequest
		return ret
	}
	return *o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationAccountUpdateAttributes) GetDataflowsOk() (*TwilioIntegrationDataflowsRequest, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *TwilioIntegrationAccountUpdateAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given TwilioIntegrationDataflowsRequest and assigns it to the Dataflows field.
func (o *TwilioIntegrationAccountUpdateAttributes) SetDataflows(v TwilioIntegrationDataflowsRequest) {
	o.Dataflows = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TwilioIntegrationAccountUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationAccountUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TwilioIntegrationAccountUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TwilioIntegrationAccountUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *TwilioIntegrationAccountUpdateAttributes) GetSettings() TwilioIntegrationAccountSettingsUpdate {
	if o == nil || o.Settings == nil {
		var ret TwilioIntegrationAccountSettingsUpdate
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationAccountUpdateAttributes) GetSettingsOk() (*TwilioIntegrationAccountSettingsUpdate, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *TwilioIntegrationAccountUpdateAttributes) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given TwilioIntegrationAccountSettingsUpdate and assigns it to the Settings field.
func (o *TwilioIntegrationAccountUpdateAttributes) SetSettings(v TwilioIntegrationAccountSettingsUpdate) {
	o.Settings = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TwilioIntegrationAccountUpdateAttributes) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TwilioIntegrationAccountUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *TwilioIntegrationAccountAuthenticationUpdate `json:"authentication,omitempty"`
		Dataflows      *TwilioIntegrationDataflowsRequest            `json:"dataflows,omitempty"`
		Name           *string                                       `json:"name,omitempty"`
		Settings       *TwilioIntegrationAccountSettingsUpdate       `json:"settings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	o.Name = all.Name
	if all.Settings != nil && all.Settings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Settings = all.Settings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
