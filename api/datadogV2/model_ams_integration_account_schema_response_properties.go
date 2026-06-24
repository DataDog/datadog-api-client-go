// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AmsIntegrationAccountSchemaResponseProperties The properties object containing settings and secrets schema definitions.
// Both are always present in every integration schema, but the fields within each
// vary depending on the specific integration.
type AmsIntegrationAccountSchemaResponseProperties struct {
	// JSON Schema definition for the secrets object.
	// Contains sensitive credentials required for the integration such as API keys,
	// tokens, and passwords. These values are write-only and never returned in responses.
	Secrets AmsIntegrationAccountSchemaResponseSecretsObject `json:"secrets"`
	// JSON Schema definition for the settings object.
	// Contains integration-specific configuration fields such as account identifiers,
	// feature toggles, and non-sensitive configuration options.
	Settings AmsIntegrationAccountSchemaResponseSettingsObject `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAmsIntegrationAccountSchemaResponseProperties instantiates a new AmsIntegrationAccountSchemaResponseProperties object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAmsIntegrationAccountSchemaResponseProperties(secrets AmsIntegrationAccountSchemaResponseSecretsObject, settings AmsIntegrationAccountSchemaResponseSettingsObject) *AmsIntegrationAccountSchemaResponseProperties {
	this := AmsIntegrationAccountSchemaResponseProperties{}
	this.Secrets = secrets
	this.Settings = settings
	return &this
}

// NewAmsIntegrationAccountSchemaResponsePropertiesWithDefaults instantiates a new AmsIntegrationAccountSchemaResponseProperties object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAmsIntegrationAccountSchemaResponsePropertiesWithDefaults() *AmsIntegrationAccountSchemaResponseProperties {
	this := AmsIntegrationAccountSchemaResponseProperties{}
	return &this
}

// GetSecrets returns the Secrets field value.
func (o *AmsIntegrationAccountSchemaResponseProperties) GetSecrets() AmsIntegrationAccountSchemaResponseSecretsObject {
	if o == nil {
		var ret AmsIntegrationAccountSchemaResponseSecretsObject
		return ret
	}
	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value
// and a boolean to check if the value has been set.
func (o *AmsIntegrationAccountSchemaResponseProperties) GetSecretsOk() (*AmsIntegrationAccountSchemaResponseSecretsObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secrets, true
}

// SetSecrets sets field value.
func (o *AmsIntegrationAccountSchemaResponseProperties) SetSecrets(v AmsIntegrationAccountSchemaResponseSecretsObject) {
	o.Secrets = v
}

// GetSettings returns the Settings field value.
func (o *AmsIntegrationAccountSchemaResponseProperties) GetSettings() AmsIntegrationAccountSchemaResponseSettingsObject {
	if o == nil {
		var ret AmsIntegrationAccountSchemaResponseSettingsObject
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *AmsIntegrationAccountSchemaResponseProperties) GetSettingsOk() (*AmsIntegrationAccountSchemaResponseSettingsObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *AmsIntegrationAccountSchemaResponseProperties) SetSettings(v AmsIntegrationAccountSchemaResponseSettingsObject) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AmsIntegrationAccountSchemaResponseProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["secrets"] = o.Secrets
	toSerialize["settings"] = o.Settings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AmsIntegrationAccountSchemaResponseProperties) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Secrets  *AmsIntegrationAccountSchemaResponseSecretsObject  `json:"secrets"`
		Settings *AmsIntegrationAccountSchemaResponseSettingsObject `json:"settings"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Secrets == nil {
		return fmt.Errorf("required field secrets missing")
	}
	if all.Settings == nil {
		return fmt.Errorf("required field settings missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"secrets", "settings"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Secrets.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Secrets = *all.Secrets
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
