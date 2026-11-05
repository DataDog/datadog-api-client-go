// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureCreateRequestDataType Data capture overrides resource type.
type AiGatewayCaptureCreateRequestDataType string

// List of AiGatewayCaptureCreateRequestDataType.
const (
	AIGATEWAYCAPTURECREATEREQUESTDATATYPE_DATA_CAPTURE_OVERRIDES AiGatewayCaptureCreateRequestDataType = "data-capture-overrides"
)

var allowedAiGatewayCaptureCreateRequestDataTypeEnumValues = []AiGatewayCaptureCreateRequestDataType{
	AIGATEWAYCAPTURECREATEREQUESTDATATYPE_DATA_CAPTURE_OVERRIDES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayCaptureCreateRequestDataType) GetAllowedValues() []AiGatewayCaptureCreateRequestDataType {
	return allowedAiGatewayCaptureCreateRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayCaptureCreateRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayCaptureCreateRequestDataType(value)
	return nil
}

// NewAiGatewayCaptureCreateRequestDataTypeFromValue returns a pointer to a valid AiGatewayCaptureCreateRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayCaptureCreateRequestDataTypeFromValue(v string) (*AiGatewayCaptureCreateRequestDataType, error) {
	ev := AiGatewayCaptureCreateRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayCaptureCreateRequestDataType: valid values are %v", v, allowedAiGatewayCaptureCreateRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayCaptureCreateRequestDataType) IsValid() bool {
	for _, existing := range allowedAiGatewayCaptureCreateRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayCaptureCreateRequestDataType value.
func (v AiGatewayCaptureCreateRequestDataType) Ptr() *AiGatewayCaptureCreateRequestDataType {
	return &v
}
