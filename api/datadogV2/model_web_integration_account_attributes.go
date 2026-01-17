// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WebIntegrationAccountAttributes Attributes for a web integration account.
type WebIntegrationAccountAttributes struct {
	// The name of the account.
	Name string `json:"name"`
	// Integration-specific settings for the account. The structure and required fields vary by integration type.
	// Use the schema endpoint to retrieve the specific requirements for each integration.
	Settings map[string]interface{} `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWebIntegrationAccountAttributes instantiates a new WebIntegrationAccountAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWebIntegrationAccountAttributes(name string, settings map[string]interface{}) *WebIntegrationAccountAttributes {
	this := WebIntegrationAccountAttributes{}
	this.Name = name
	this.Settings = settings
	return &this
}

// NewWebIntegrationAccountAttributesWithDefaults instantiates a new WebIntegrationAccountAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWebIntegrationAccountAttributesWithDefaults() *WebIntegrationAccountAttributes {
	this := WebIntegrationAccountAttributes{}
	return &this
}

// GetName returns the Name field value.
func (o *WebIntegrationAccountAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WebIntegrationAccountAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *WebIntegrationAccountAttributes) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value.
func (o *WebIntegrationAccountAttributes) GetSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *WebIntegrationAccountAttributes) GetSettingsOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *WebIntegrationAccountAttributes) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WebIntegrationAccountAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["settings"] = o.Settings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WebIntegrationAccountAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name     *string                 `json:"name"`
		Settings *map[string]interface{} `json:"settings"`
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "settings"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Settings = *all.Settings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
