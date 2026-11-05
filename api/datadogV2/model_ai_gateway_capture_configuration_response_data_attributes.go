// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureConfigurationResponseDataAttributes Organization default mode, scoped overrides, and current configuration revision.
type AiGatewayCaptureConfigurationResponseDataAttributes struct {
	// Opaque revision of the organization default and all overrides. Writes require the revision returned by a current read or preview.
	ConfigurationRevision string `json:"configuration_revision"`
	// Items Metadata-only capture adds, each listed once. An item requires the items whose information it already reveals: `tool_output_size`, `tool_names` and `conversation_shape` require `tool_usage`, and `input_size_by_role` requires `tool_output_size`. A set that leaves out a required item is rejected. When several Metadata-only overrides match a request, it gets only the items all of them select.
	DefaultMetadataCapture []AiGatewayMetadataCaptureItem `json:"default_metadata_capture,omitempty"`
	// Default capture mode for the organization. Change it in **Org defaults**.
	DefaultMode AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode `json:"default_mode"`
	// Scoped overrides currently configured for the organization.
	Overrides []AiGatewayCaptureOverride `json:"overrides"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayCaptureConfigurationResponseDataAttributes instantiates a new AiGatewayCaptureConfigurationResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayCaptureConfigurationResponseDataAttributes(configurationRevision string, defaultMode AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode, overrides []AiGatewayCaptureOverride) *AiGatewayCaptureConfigurationResponseDataAttributes {
	this := AiGatewayCaptureConfigurationResponseDataAttributes{}
	this.ConfigurationRevision = configurationRevision
	this.DefaultMode = defaultMode
	this.Overrides = overrides
	return &this
}

// NewAiGatewayCaptureConfigurationResponseDataAttributesWithDefaults instantiates a new AiGatewayCaptureConfigurationResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayCaptureConfigurationResponseDataAttributesWithDefaults() *AiGatewayCaptureConfigurationResponseDataAttributes {
	this := AiGatewayCaptureConfigurationResponseDataAttributes{}
	return &this
}

// GetConfigurationRevision returns the ConfigurationRevision field value.
func (o *AiGatewayCaptureConfigurationResponseDataAttributes) GetConfigurationRevision() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigurationRevision
}

// GetConfigurationRevisionOk returns a tuple with the ConfigurationRevision field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureConfigurationResponseDataAttributes) GetConfigurationRevisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationRevision, true
}

// SetConfigurationRevision sets field value.
func (o *AiGatewayCaptureConfigurationResponseDataAttributes) SetConfigurationRevision(v string) {
	o.ConfigurationRevision = v
}

// GetDefaultMetadataCapture returns the DefaultMetadataCapture field value if set, zero value otherwise.
func (o *AiGatewayCaptureConfigurationResponseDataAttributes) GetDefaultMetadataCapture() []AiGatewayMetadataCaptureItem {
	if o == nil || o.DefaultMetadataCapture == nil {
		var ret []AiGatewayMetadataCaptureItem
		return ret
	}
	return o.DefaultMetadataCapture
}

// GetDefaultMetadataCaptureOk returns a tuple with the DefaultMetadataCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureConfigurationResponseDataAttributes) GetDefaultMetadataCaptureOk() (*[]AiGatewayMetadataCaptureItem, bool) {
	if o == nil || o.DefaultMetadataCapture == nil {
		return nil, false
	}
	return &o.DefaultMetadataCapture, true
}

// HasDefaultMetadataCapture returns a boolean if a field has been set.
func (o *AiGatewayCaptureConfigurationResponseDataAttributes) HasDefaultMetadataCapture() bool {
	return o != nil && o.DefaultMetadataCapture != nil
}

// SetDefaultMetadataCapture gets a reference to the given []AiGatewayMetadataCaptureItem and assigns it to the DefaultMetadataCapture field.
func (o *AiGatewayCaptureConfigurationResponseDataAttributes) SetDefaultMetadataCapture(v []AiGatewayMetadataCaptureItem) {
	o.DefaultMetadataCapture = v
}

// GetDefaultMode returns the DefaultMode field value.
func (o *AiGatewayCaptureConfigurationResponseDataAttributes) GetDefaultMode() AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode {
	if o == nil {
		var ret AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode
		return ret
	}
	return o.DefaultMode
}

// GetDefaultModeOk returns a tuple with the DefaultMode field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureConfigurationResponseDataAttributes) GetDefaultModeOk() (*AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultMode, true
}

// SetDefaultMode sets field value.
func (o *AiGatewayCaptureConfigurationResponseDataAttributes) SetDefaultMode(v AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode) {
	o.DefaultMode = v
}

// GetOverrides returns the Overrides field value.
func (o *AiGatewayCaptureConfigurationResponseDataAttributes) GetOverrides() []AiGatewayCaptureOverride {
	if o == nil {
		var ret []AiGatewayCaptureOverride
		return ret
	}
	return o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureConfigurationResponseDataAttributes) GetOverridesOk() (*[]AiGatewayCaptureOverride, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Overrides, true
}

// SetOverrides sets field value.
func (o *AiGatewayCaptureConfigurationResponseDataAttributes) SetOverrides(v []AiGatewayCaptureOverride) {
	o.Overrides = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayCaptureConfigurationResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["configuration_revision"] = o.ConfigurationRevision
	if o.DefaultMetadataCapture != nil {
		toSerialize["default_metadata_capture"] = o.DefaultMetadataCapture
	}
	toSerialize["default_mode"] = o.DefaultMode
	toSerialize["overrides"] = o.Overrides

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayCaptureConfigurationResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigurationRevision  *string                                                         `json:"configuration_revision"`
		DefaultMetadataCapture []AiGatewayMetadataCaptureItem                                  `json:"default_metadata_capture,omitempty"`
		DefaultMode            *AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode `json:"default_mode"`
		Overrides              *[]AiGatewayCaptureOverride                                     `json:"overrides"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConfigurationRevision == nil {
		return fmt.Errorf("required field configuration_revision missing")
	}
	if all.DefaultMode == nil {
		return fmt.Errorf("required field default_mode missing")
	}
	if all.Overrides == nil {
		return fmt.Errorf("required field overrides missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"configuration_revision", "default_metadata_capture", "default_mode", "overrides"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ConfigurationRevision = *all.ConfigurationRevision
	o.DefaultMetadataCapture = all.DefaultMetadataCapture
	if !all.DefaultMode.IsValid() {
		hasInvalidField = true
	} else {
		o.DefaultMode = *all.DefaultMode
	}
	o.Overrides = *all.Overrides

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
