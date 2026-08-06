// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioAccountAttributes Attributes of a Twilio integration account. The Twilio configuration is hoisted directly onto the attributes; there is no interface wrapper because the `twilio` interface is fixed by the endpoint path.
type TwilioAccountAttributes struct {
	// Authentication methods supported by the Twilio interface. Exactly one is set, selected by its `type`.
	Authentication TwilioAuthentication `json:"authentication"`
	// Dataflows for the Twilio interface.
	Dataflows []TwilioDataflow `json:"dataflows,omitempty"`
	// Human-readable name of the account.
	Name string `json:"name"`
	// Read-only permission information for the account, derived from its restriction policy.
	Permissions *IntegrationAccountPermissions `json:"permissions,omitempty"`
	// Twilio interface settings.
	Settings *TwilioSettings `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTwilioAccountAttributes instantiates a new TwilioAccountAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTwilioAccountAttributes(authentication TwilioAuthentication, name string) *TwilioAccountAttributes {
	this := TwilioAccountAttributes{}
	this.Authentication = authentication
	this.Name = name
	return &this
}

// NewTwilioAccountAttributesWithDefaults instantiates a new TwilioAccountAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTwilioAccountAttributesWithDefaults() *TwilioAccountAttributes {
	this := TwilioAccountAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value.
func (o *TwilioAccountAttributes) GetAuthentication() TwilioAuthentication {
	if o == nil {
		var ret TwilioAuthentication
		return ret
	}
	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *TwilioAccountAttributes) GetAuthenticationOk() (*TwilioAuthentication, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value.
func (o *TwilioAccountAttributes) SetAuthentication(v TwilioAuthentication) {
	o.Authentication = v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *TwilioAccountAttributes) GetDataflows() []TwilioDataflow {
	if o == nil || o.Dataflows == nil {
		var ret []TwilioDataflow
		return ret
	}
	return o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioAccountAttributes) GetDataflowsOk() (*[]TwilioDataflow, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return &o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *TwilioAccountAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given []TwilioDataflow and assigns it to the Dataflows field.
func (o *TwilioAccountAttributes) SetDataflows(v []TwilioDataflow) {
	o.Dataflows = v
}

// GetName returns the Name field value.
func (o *TwilioAccountAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TwilioAccountAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *TwilioAccountAttributes) SetName(v string) {
	o.Name = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *TwilioAccountAttributes) GetPermissions() IntegrationAccountPermissions {
	if o == nil || o.Permissions == nil {
		var ret IntegrationAccountPermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioAccountAttributes) GetPermissionsOk() (*IntegrationAccountPermissions, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *TwilioAccountAttributes) HasPermissions() bool {
	return o != nil && o.Permissions != nil
}

// SetPermissions gets a reference to the given IntegrationAccountPermissions and assigns it to the Permissions field.
func (o *TwilioAccountAttributes) SetPermissions(v IntegrationAccountPermissions) {
	o.Permissions = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *TwilioAccountAttributes) GetSettings() TwilioSettings {
	if o == nil || o.Settings == nil {
		var ret TwilioSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioAccountAttributes) GetSettingsOk() (*TwilioSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *TwilioAccountAttributes) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given TwilioSettings and assigns it to the Settings field.
func (o *TwilioAccountAttributes) SetSettings(v TwilioSettings) {
	o.Settings = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TwilioAccountAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["authentication"] = o.Authentication
	if o.Dataflows != nil {
		toSerialize["dataflows"] = o.Dataflows
	}
	toSerialize["name"] = o.Name
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
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
func (o *TwilioAccountAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *TwilioAuthentication          `json:"authentication"`
		Dataflows      []TwilioDataflow               `json:"dataflows,omitempty"`
		Name           *string                        `json:"name"`
		Permissions    *IntegrationAccountPermissions `json:"permissions,omitempty"`
		Settings       *TwilioSettings                `json:"settings,omitempty"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"authentication", "dataflows", "name", "permissions", "settings"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Authentication = *all.Authentication
	o.Dataflows = all.Dataflows
	o.Name = *all.Name
	if all.Permissions != nil && all.Permissions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Permissions = all.Permissions
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
