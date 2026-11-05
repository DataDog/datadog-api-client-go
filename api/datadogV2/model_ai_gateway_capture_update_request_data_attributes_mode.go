// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureUpdateRequestDataAttributesMode Capture mode for matching traffic.
type AiGatewayCaptureUpdateRequestDataAttributesMode string

// List of AiGatewayCaptureUpdateRequestDataAttributesMode.
const (
	AIGATEWAYCAPTUREUPDATEREQUESTDATAATTRIBUTESMODE_METADATA_ONLY AiGatewayCaptureUpdateRequestDataAttributesMode = "metadata_only"
	AIGATEWAYCAPTUREUPDATEREQUESTDATAATTRIBUTESMODE_FULL_CAPTURE  AiGatewayCaptureUpdateRequestDataAttributesMode = "full_capture"
)

var allowedAiGatewayCaptureUpdateRequestDataAttributesModeEnumValues = []AiGatewayCaptureUpdateRequestDataAttributesMode{
	AIGATEWAYCAPTUREUPDATEREQUESTDATAATTRIBUTESMODE_METADATA_ONLY,
	AIGATEWAYCAPTUREUPDATEREQUESTDATAATTRIBUTESMODE_FULL_CAPTURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayCaptureUpdateRequestDataAttributesMode) GetAllowedValues() []AiGatewayCaptureUpdateRequestDataAttributesMode {
	return allowedAiGatewayCaptureUpdateRequestDataAttributesModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayCaptureUpdateRequestDataAttributesMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayCaptureUpdateRequestDataAttributesMode(value)
	return nil
}

// NewAiGatewayCaptureUpdateRequestDataAttributesModeFromValue returns a pointer to a valid AiGatewayCaptureUpdateRequestDataAttributesMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayCaptureUpdateRequestDataAttributesModeFromValue(v string) (*AiGatewayCaptureUpdateRequestDataAttributesMode, error) {
	ev := AiGatewayCaptureUpdateRequestDataAttributesMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayCaptureUpdateRequestDataAttributesMode: valid values are %v", v, allowedAiGatewayCaptureUpdateRequestDataAttributesModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayCaptureUpdateRequestDataAttributesMode) IsValid() bool {
	for _, existing := range allowedAiGatewayCaptureUpdateRequestDataAttributesModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayCaptureUpdateRequestDataAttributesMode value.
func (v AiGatewayCaptureUpdateRequestDataAttributesMode) Ptr() *AiGatewayCaptureUpdateRequestDataAttributesMode {
	return &v
}
