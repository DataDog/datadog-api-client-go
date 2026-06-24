// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AmsIntegrationAccountAttributes Attributes for a web integration account.
type AmsIntegrationAccountAttributes struct {
	// The name of the account.
	Name string `json:"name"`
	// Integration-specific settings for the account. The structure and required fields vary by integration type.
	// Use the schema endpoint to retrieve the specific requirements for each integration.
	Settings map[string]interface{} `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAmsIntegrationAccountAttributes instantiates a new AmsIntegrationAccountAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAmsIntegrationAccountAttributes(name string, settings map[string]interface{}) *AmsIntegrationAccountAttributes {
	this := AmsIntegrationAccountAttributes{}
	this.Name = name
	this.Settings = settings
	return &this
}

// NewAmsIntegrationAccountAttributesWithDefaults instantiates a new AmsIntegrationAccountAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAmsIntegrationAccountAttributesWithDefaults() *AmsIntegrationAccountAttributes {
	this := AmsIntegrationAccountAttributes{}
	return &this
}

// GetName returns the Name field value.
func (o *AmsIntegrationAccountAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AmsIntegrationAccountAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AmsIntegrationAccountAttributes) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value.
func (o *AmsIntegrationAccountAttributes) GetSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *AmsIntegrationAccountAttributes) GetSettingsOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *AmsIntegrationAccountAttributes) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AmsIntegrationAccountAttributes) MarshalJSON() ([]byte, error) {
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
func (o *AmsIntegrationAccountAttributes) UnmarshalJSON(bytes []byte) (err error) {
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
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
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
