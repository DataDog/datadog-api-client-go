// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureOverrideResponseDataAttributesMode Capture mode for matching traffic.
type AiGatewayCaptureOverrideResponseDataAttributesMode string

// List of AiGatewayCaptureOverrideResponseDataAttributesMode.
const (
	AIGATEWAYCAPTUREOVERRIDERESPONSEDATAATTRIBUTESMODE_METADATA_ONLY AiGatewayCaptureOverrideResponseDataAttributesMode = "metadata_only"
	AIGATEWAYCAPTUREOVERRIDERESPONSEDATAATTRIBUTESMODE_FULL_CAPTURE  AiGatewayCaptureOverrideResponseDataAttributesMode = "full_capture"
)

var allowedAiGatewayCaptureOverrideResponseDataAttributesModeEnumValues = []AiGatewayCaptureOverrideResponseDataAttributesMode{
	AIGATEWAYCAPTUREOVERRIDERESPONSEDATAATTRIBUTESMODE_METADATA_ONLY,
	AIGATEWAYCAPTUREOVERRIDERESPONSEDATAATTRIBUTESMODE_FULL_CAPTURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayCaptureOverrideResponseDataAttributesMode) GetAllowedValues() []AiGatewayCaptureOverrideResponseDataAttributesMode {
	return allowedAiGatewayCaptureOverrideResponseDataAttributesModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayCaptureOverrideResponseDataAttributesMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayCaptureOverrideResponseDataAttributesMode(value)
	return nil
}

// NewAiGatewayCaptureOverrideResponseDataAttributesModeFromValue returns a pointer to a valid AiGatewayCaptureOverrideResponseDataAttributesMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayCaptureOverrideResponseDataAttributesModeFromValue(v string) (*AiGatewayCaptureOverrideResponseDataAttributesMode, error) {
	ev := AiGatewayCaptureOverrideResponseDataAttributesMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayCaptureOverrideResponseDataAttributesMode: valid values are %v", v, allowedAiGatewayCaptureOverrideResponseDataAttributesModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayCaptureOverrideResponseDataAttributesMode) IsValid() bool {
	for _, existing := range allowedAiGatewayCaptureOverrideResponseDataAttributesModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayCaptureOverrideResponseDataAttributesMode value.
func (v AiGatewayCaptureOverrideResponseDataAttributesMode) Ptr() *AiGatewayCaptureOverrideResponseDataAttributesMode {
	return &v
}
