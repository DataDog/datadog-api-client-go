// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureDeleteRequestDataAttributesMode Capture mode for matching traffic.
type AiGatewayCaptureDeleteRequestDataAttributesMode string

// List of AiGatewayCaptureDeleteRequestDataAttributesMode.
const (
	AIGATEWAYCAPTUREDELETEREQUESTDATAATTRIBUTESMODE_METADATA_ONLY AiGatewayCaptureDeleteRequestDataAttributesMode = "metadata_only"
	AIGATEWAYCAPTUREDELETEREQUESTDATAATTRIBUTESMODE_FULL_CAPTURE  AiGatewayCaptureDeleteRequestDataAttributesMode = "full_capture"
)

var allowedAiGatewayCaptureDeleteRequestDataAttributesModeEnumValues = []AiGatewayCaptureDeleteRequestDataAttributesMode{
	AIGATEWAYCAPTUREDELETEREQUESTDATAATTRIBUTESMODE_METADATA_ONLY,
	AIGATEWAYCAPTUREDELETEREQUESTDATAATTRIBUTESMODE_FULL_CAPTURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayCaptureDeleteRequestDataAttributesMode) GetAllowedValues() []AiGatewayCaptureDeleteRequestDataAttributesMode {
	return allowedAiGatewayCaptureDeleteRequestDataAttributesModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayCaptureDeleteRequestDataAttributesMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayCaptureDeleteRequestDataAttributesMode(value)
	return nil
}

// NewAiGatewayCaptureDeleteRequestDataAttributesModeFromValue returns a pointer to a valid AiGatewayCaptureDeleteRequestDataAttributesMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayCaptureDeleteRequestDataAttributesModeFromValue(v string) (*AiGatewayCaptureDeleteRequestDataAttributesMode, error) {
	ev := AiGatewayCaptureDeleteRequestDataAttributesMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayCaptureDeleteRequestDataAttributesMode: valid values are %v", v, allowedAiGatewayCaptureDeleteRequestDataAttributesModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayCaptureDeleteRequestDataAttributesMode) IsValid() bool {
	for _, existing := range allowedAiGatewayCaptureDeleteRequestDataAttributesModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayCaptureDeleteRequestDataAttributesMode value.
func (v AiGatewayCaptureDeleteRequestDataAttributesMode) Ptr() *AiGatewayCaptureDeleteRequestDataAttributesMode {
	return &v
}
