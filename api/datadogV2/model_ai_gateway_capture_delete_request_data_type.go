// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureDeleteRequestDataType Data capture overrides resource type.
type AiGatewayCaptureDeleteRequestDataType string

// List of AiGatewayCaptureDeleteRequestDataType.
const (
	AIGATEWAYCAPTUREDELETEREQUESTDATATYPE_DATA_CAPTURE_OVERRIDES AiGatewayCaptureDeleteRequestDataType = "data-capture-overrides"
)

var allowedAiGatewayCaptureDeleteRequestDataTypeEnumValues = []AiGatewayCaptureDeleteRequestDataType{
	AIGATEWAYCAPTUREDELETEREQUESTDATATYPE_DATA_CAPTURE_OVERRIDES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayCaptureDeleteRequestDataType) GetAllowedValues() []AiGatewayCaptureDeleteRequestDataType {
	return allowedAiGatewayCaptureDeleteRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayCaptureDeleteRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayCaptureDeleteRequestDataType(value)
	return nil
}

// NewAiGatewayCaptureDeleteRequestDataTypeFromValue returns a pointer to a valid AiGatewayCaptureDeleteRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayCaptureDeleteRequestDataTypeFromValue(v string) (*AiGatewayCaptureDeleteRequestDataType, error) {
	ev := AiGatewayCaptureDeleteRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayCaptureDeleteRequestDataType: valid values are %v", v, allowedAiGatewayCaptureDeleteRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayCaptureDeleteRequestDataType) IsValid() bool {
	for _, existing := range allowedAiGatewayCaptureDeleteRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayCaptureDeleteRequestDataType value.
func (v AiGatewayCaptureDeleteRequestDataType) Ptr() *AiGatewayCaptureDeleteRequestDataType {
	return &v
}
