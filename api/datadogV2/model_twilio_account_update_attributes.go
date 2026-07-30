// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioAccountUpdateAttributes Updatable attributes of a Twilio integration account. Every field is optional; only the fields provided are changed.
type TwilioAccountUpdateAttributes struct {
	// Authentication methods supported by the Twilio interface. Exactly one is set, selected by its `type`.
	Authentication *TwilioAuthentication `json:"authentication,omitempty"`
	// Dataflows for the Twilio interface.
	Dataflows []TwilioDataflow `json:"dataflows,omitempty"`
	// Human-readable name of the account.
	Name *string `json:"name,omitempty"`
	// Partial Twilio interface settings for updates.
	Settings *TwilioSettingsUpdate `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTwilioAccountUpdateAttributes instantiates a new TwilioAccountUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTwilioAccountUpdateAttributes() *TwilioAccountUpdateAttributes {
	this := TwilioAccountUpdateAttributes{}
	return &this
}

// NewTwilioAccountUpdateAttributesWithDefaults instantiates a new TwilioAccountUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTwilioAccountUpdateAttributesWithDefaults() *TwilioAccountUpdateAttributes {
	this := TwilioAccountUpdateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *TwilioAccountUpdateAttributes) GetAuthentication() TwilioAuthentication {
	if o == nil || o.Authentication == nil {
		var ret TwilioAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioAccountUpdateAttributes) GetAuthenticationOk() (*TwilioAuthentication, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *TwilioAccountUpdateAttributes) HasAuthentication() bool {
	return o != nil && o.Authentication != nil
}

// SetAuthentication gets a reference to the given TwilioAuthentication and assigns it to the Authentication field.
func (o *TwilioAccountUpdateAttributes) SetAuthentication(v TwilioAuthentication) {
	o.Authentication = &v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *TwilioAccountUpdateAttributes) GetDataflows() []TwilioDataflow {
	if o == nil || o.Dataflows == nil {
		var ret []TwilioDataflow
		return ret
	}
	return o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioAccountUpdateAttributes) GetDataflowsOk() (*[]TwilioDataflow, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return &o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *TwilioAccountUpdateAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given []TwilioDataflow and assigns it to the Dataflows field.
func (o *TwilioAccountUpdateAttributes) SetDataflows(v []TwilioDataflow) {
	o.Dataflows = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TwilioAccountUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioAccountUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TwilioAccountUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TwilioAccountUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *TwilioAccountUpdateAttributes) GetSettings() TwilioSettingsUpdate {
	if o == nil || o.Settings == nil {
		var ret TwilioSettingsUpdate
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioAccountUpdateAttributes) GetSettingsOk() (*TwilioSettingsUpdate, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *TwilioAccountUpdateAttributes) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given TwilioSettingsUpdate and assigns it to the Settings field.
func (o *TwilioAccountUpdateAttributes) SetSettings(v TwilioSettingsUpdate) {
	o.Settings = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TwilioAccountUpdateAttributes) MarshalJSON() ([]byte, error) {
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
func (o *TwilioAccountUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *TwilioAuthentication `json:"authentication,omitempty"`
		Dataflows      []TwilioDataflow      `json:"dataflows,omitempty"`
		Name           *string               `json:"name,omitempty"`
		Settings       *TwilioSettingsUpdate `json:"settings,omitempty"`
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
