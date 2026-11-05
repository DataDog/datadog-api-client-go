// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureOverrideResponseDataType Data capture overrides resource type.
type AiGatewayCaptureOverrideResponseDataType string

// List of AiGatewayCaptureOverrideResponseDataType.
const (
	AIGATEWAYCAPTUREOVERRIDERESPONSEDATATYPE_DATA_CAPTURE_OVERRIDES AiGatewayCaptureOverrideResponseDataType = "data-capture-overrides"
)

var allowedAiGatewayCaptureOverrideResponseDataTypeEnumValues = []AiGatewayCaptureOverrideResponseDataType{
	AIGATEWAYCAPTUREOVERRIDERESPONSEDATATYPE_DATA_CAPTURE_OVERRIDES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayCaptureOverrideResponseDataType) GetAllowedValues() []AiGatewayCaptureOverrideResponseDataType {
	return allowedAiGatewayCaptureOverrideResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayCaptureOverrideResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayCaptureOverrideResponseDataType(value)
	return nil
}

// NewAiGatewayCaptureOverrideResponseDataTypeFromValue returns a pointer to a valid AiGatewayCaptureOverrideResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayCaptureOverrideResponseDataTypeFromValue(v string) (*AiGatewayCaptureOverrideResponseDataType, error) {
	ev := AiGatewayCaptureOverrideResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayCaptureOverrideResponseDataType: valid values are %v", v, allowedAiGatewayCaptureOverrideResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayCaptureOverrideResponseDataType) IsValid() bool {
	for _, existing := range allowedAiGatewayCaptureOverrideResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayCaptureOverrideResponseDataType value.
func (v AiGatewayCaptureOverrideResponseDataType) Ptr() *AiGatewayCaptureOverrideResponseDataType {
	return &v
}
