// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelMetadataStatus Completeness of the matched source record. Partial also identifies metadata from a base model for a dated snapshot or from a shared canonical-model fallback.
type AiGatewayModelMetadataStatus string

// List of AiGatewayModelMetadataStatus.
const (
	AIGATEWAYMODELMETADATASTATUS_AVAILABLE AiGatewayModelMetadataStatus = "available"
	AIGATEWAYMODELMETADATASTATUS_PARTIAL   AiGatewayModelMetadataStatus = "partial"
)

var allowedAiGatewayModelMetadataStatusEnumValues = []AiGatewayModelMetadataStatus{
	AIGATEWAYMODELMETADATASTATUS_AVAILABLE,
	AIGATEWAYMODELMETADATASTATUS_PARTIAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayModelMetadataStatus) GetAllowedValues() []AiGatewayModelMetadataStatus {
	return allowedAiGatewayModelMetadataStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayModelMetadataStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayModelMetadataStatus(value)
	return nil
}

// NewAiGatewayModelMetadataStatusFromValue returns a pointer to a valid AiGatewayModelMetadataStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayModelMetadataStatusFromValue(v string) (*AiGatewayModelMetadataStatus, error) {
	ev := AiGatewayModelMetadataStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayModelMetadataStatus: valid values are %v", v, allowedAiGatewayModelMetadataStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayModelMetadataStatus) IsValid() bool {
	for _, existing := range allowedAiGatewayModelMetadataStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayModelMetadataStatus value.
func (v AiGatewayModelMetadataStatus) Ptr() *AiGatewayModelMetadataStatus {
	return &v
}
