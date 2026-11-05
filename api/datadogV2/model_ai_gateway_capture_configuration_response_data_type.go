// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureConfigurationResponseDataType Data capture configurations resource type.
type AiGatewayCaptureConfigurationResponseDataType string

// List of AiGatewayCaptureConfigurationResponseDataType.
const (
	AIGATEWAYCAPTURECONFIGURATIONRESPONSEDATATYPE_DATA_CAPTURE_CONFIGURATIONS AiGatewayCaptureConfigurationResponseDataType = "data-capture-configurations"
)

var allowedAiGatewayCaptureConfigurationResponseDataTypeEnumValues = []AiGatewayCaptureConfigurationResponseDataType{
	AIGATEWAYCAPTURECONFIGURATIONRESPONSEDATATYPE_DATA_CAPTURE_CONFIGURATIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayCaptureConfigurationResponseDataType) GetAllowedValues() []AiGatewayCaptureConfigurationResponseDataType {
	return allowedAiGatewayCaptureConfigurationResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayCaptureConfigurationResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayCaptureConfigurationResponseDataType(value)
	return nil
}

// NewAiGatewayCaptureConfigurationResponseDataTypeFromValue returns a pointer to a valid AiGatewayCaptureConfigurationResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayCaptureConfigurationResponseDataTypeFromValue(v string) (*AiGatewayCaptureConfigurationResponseDataType, error) {
	ev := AiGatewayCaptureConfigurationResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayCaptureConfigurationResponseDataType: valid values are %v", v, allowedAiGatewayCaptureConfigurationResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayCaptureConfigurationResponseDataType) IsValid() bool {
	for _, existing := range allowedAiGatewayCaptureConfigurationResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayCaptureConfigurationResponseDataType value.
func (v AiGatewayCaptureConfigurationResponseDataType) Ptr() *AiGatewayCaptureConfigurationResponseDataType {
	return &v
}
