// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureUpdateRequestDataType Data capture overrides resource type.
type AiGatewayCaptureUpdateRequestDataType string

// List of AiGatewayCaptureUpdateRequestDataType.
const (
	AIGATEWAYCAPTUREUPDATEREQUESTDATATYPE_DATA_CAPTURE_OVERRIDES AiGatewayCaptureUpdateRequestDataType = "data-capture-overrides"
)

var allowedAiGatewayCaptureUpdateRequestDataTypeEnumValues = []AiGatewayCaptureUpdateRequestDataType{
	AIGATEWAYCAPTUREUPDATEREQUESTDATATYPE_DATA_CAPTURE_OVERRIDES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayCaptureUpdateRequestDataType) GetAllowedValues() []AiGatewayCaptureUpdateRequestDataType {
	return allowedAiGatewayCaptureUpdateRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayCaptureUpdateRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayCaptureUpdateRequestDataType(value)
	return nil
}

// NewAiGatewayCaptureUpdateRequestDataTypeFromValue returns a pointer to a valid AiGatewayCaptureUpdateRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayCaptureUpdateRequestDataTypeFromValue(v string) (*AiGatewayCaptureUpdateRequestDataType, error) {
	ev := AiGatewayCaptureUpdateRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayCaptureUpdateRequestDataType: valid values are %v", v, allowedAiGatewayCaptureUpdateRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayCaptureUpdateRequestDataType) IsValid() bool {
	for _, existing := range allowedAiGatewayCaptureUpdateRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayCaptureUpdateRequestDataType value.
func (v AiGatewayCaptureUpdateRequestDataType) Ptr() *AiGatewayCaptureUpdateRequestDataType {
	return &v
}
