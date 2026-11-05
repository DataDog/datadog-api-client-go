// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureOverrideMode Capture mode for matching traffic.
type AiGatewayCaptureOverrideMode string

// List of AiGatewayCaptureOverrideMode.
const (
	AIGATEWAYCAPTUREOVERRIDEMODE_METADATA_ONLY AiGatewayCaptureOverrideMode = "metadata_only"
	AIGATEWAYCAPTUREOVERRIDEMODE_FULL_CAPTURE  AiGatewayCaptureOverrideMode = "full_capture"
)

var allowedAiGatewayCaptureOverrideModeEnumValues = []AiGatewayCaptureOverrideMode{
	AIGATEWAYCAPTUREOVERRIDEMODE_METADATA_ONLY,
	AIGATEWAYCAPTUREOVERRIDEMODE_FULL_CAPTURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayCaptureOverrideMode) GetAllowedValues() []AiGatewayCaptureOverrideMode {
	return allowedAiGatewayCaptureOverrideModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayCaptureOverrideMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayCaptureOverrideMode(value)
	return nil
}

// NewAiGatewayCaptureOverrideModeFromValue returns a pointer to a valid AiGatewayCaptureOverrideMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayCaptureOverrideModeFromValue(v string) (*AiGatewayCaptureOverrideMode, error) {
	ev := AiGatewayCaptureOverrideMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayCaptureOverrideMode: valid values are %v", v, allowedAiGatewayCaptureOverrideModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayCaptureOverrideMode) IsValid() bool {
	for _, existing := range allowedAiGatewayCaptureOverrideModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayCaptureOverrideMode value.
func (v AiGatewayCaptureOverrideMode) Ptr() *AiGatewayCaptureOverrideMode {
	return &v
}
