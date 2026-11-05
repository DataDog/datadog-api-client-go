// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureCreateRequestDataAttributesMode Capture mode for matching traffic.
type AiGatewayCaptureCreateRequestDataAttributesMode string

// List of AiGatewayCaptureCreateRequestDataAttributesMode.
const (
	AIGATEWAYCAPTURECREATEREQUESTDATAATTRIBUTESMODE_METADATA_ONLY AiGatewayCaptureCreateRequestDataAttributesMode = "metadata_only"
	AIGATEWAYCAPTURECREATEREQUESTDATAATTRIBUTESMODE_FULL_CAPTURE  AiGatewayCaptureCreateRequestDataAttributesMode = "full_capture"
)

var allowedAiGatewayCaptureCreateRequestDataAttributesModeEnumValues = []AiGatewayCaptureCreateRequestDataAttributesMode{
	AIGATEWAYCAPTURECREATEREQUESTDATAATTRIBUTESMODE_METADATA_ONLY,
	AIGATEWAYCAPTURECREATEREQUESTDATAATTRIBUTESMODE_FULL_CAPTURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayCaptureCreateRequestDataAttributesMode) GetAllowedValues() []AiGatewayCaptureCreateRequestDataAttributesMode {
	return allowedAiGatewayCaptureCreateRequestDataAttributesModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayCaptureCreateRequestDataAttributesMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayCaptureCreateRequestDataAttributesMode(value)
	return nil
}

// NewAiGatewayCaptureCreateRequestDataAttributesModeFromValue returns a pointer to a valid AiGatewayCaptureCreateRequestDataAttributesMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayCaptureCreateRequestDataAttributesModeFromValue(v string) (*AiGatewayCaptureCreateRequestDataAttributesMode, error) {
	ev := AiGatewayCaptureCreateRequestDataAttributesMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayCaptureCreateRequestDataAttributesMode: valid values are %v", v, allowedAiGatewayCaptureCreateRequestDataAttributesModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayCaptureCreateRequestDataAttributesMode) IsValid() bool {
	for _, existing := range allowedAiGatewayCaptureCreateRequestDataAttributesModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayCaptureCreateRequestDataAttributesMode value.
func (v AiGatewayCaptureCreateRequestDataAttributesMode) Ptr() *AiGatewayCaptureCreateRequestDataAttributesMode {
	return &v
}
