// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayDataCaptureMode Controls whether AI Gateway captures request and response content in addition to metadata.
type AiGatewayDataCaptureMode string

// List of AiGatewayDataCaptureMode.
const (
	AIGATEWAYDATACAPTUREMODE_METADATA_ONLY AiGatewayDataCaptureMode = "metadata_only"
	AIGATEWAYDATACAPTUREMODE_FULL_CAPTURE  AiGatewayDataCaptureMode = "full_capture"
)

var allowedAiGatewayDataCaptureModeEnumValues = []AiGatewayDataCaptureMode{
	AIGATEWAYDATACAPTUREMODE_METADATA_ONLY,
	AIGATEWAYDATACAPTUREMODE_FULL_CAPTURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayDataCaptureMode) GetAllowedValues() []AiGatewayDataCaptureMode {
	return allowedAiGatewayDataCaptureModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayDataCaptureMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayDataCaptureMode(value)
	return nil
}

// NewAiGatewayDataCaptureModeFromValue returns a pointer to a valid AiGatewayDataCaptureMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayDataCaptureModeFromValue(v string) (*AiGatewayDataCaptureMode, error) {
	ev := AiGatewayDataCaptureMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayDataCaptureMode: valid values are %v", v, allowedAiGatewayDataCaptureModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayDataCaptureMode) IsValid() bool {
	for _, existing := range allowedAiGatewayDataCaptureModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayDataCaptureMode value.
func (v AiGatewayDataCaptureMode) Ptr() *AiGatewayDataCaptureMode {
	return &v
}
