// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode Default capture mode for the organization. Change it in **Org defaults**.
type AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode string

// List of AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode.
const (
	AIGATEWAYCAPTURECONFIGURATIONRESPONSEDATAATTRIBUTESDEFAULTMODE_METADATA_ONLY AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode = "metadata_only"
	AIGATEWAYCAPTURECONFIGURATIONRESPONSEDATAATTRIBUTESDEFAULTMODE_FULL_CAPTURE  AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode = "full_capture"
)

var allowedAiGatewayCaptureConfigurationResponseDataAttributesDefaultModeEnumValues = []AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode{
	AIGATEWAYCAPTURECONFIGURATIONRESPONSEDATAATTRIBUTESDEFAULTMODE_METADATA_ONLY,
	AIGATEWAYCAPTURECONFIGURATIONRESPONSEDATAATTRIBUTESDEFAULTMODE_FULL_CAPTURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode) GetAllowedValues() []AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode {
	return allowedAiGatewayCaptureConfigurationResponseDataAttributesDefaultModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode(value)
	return nil
}

// NewAiGatewayCaptureConfigurationResponseDataAttributesDefaultModeFromValue returns a pointer to a valid AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayCaptureConfigurationResponseDataAttributesDefaultModeFromValue(v string) (*AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode, error) {
	ev := AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode: valid values are %v", v, allowedAiGatewayCaptureConfigurationResponseDataAttributesDefaultModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode) IsValid() bool {
	for _, existing := range allowedAiGatewayCaptureConfigurationResponseDataAttributesDefaultModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode value.
func (v AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode) Ptr() *AiGatewayCaptureConfigurationResponseDataAttributesDefaultMode {
	return &v
}
