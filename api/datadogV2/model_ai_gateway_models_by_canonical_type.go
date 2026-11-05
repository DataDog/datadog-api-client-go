// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelsByCanonicalType Canonical model grouping resource type.
type AiGatewayModelsByCanonicalType string

// List of AiGatewayModelsByCanonicalType.
const (
	AIGATEWAYMODELSBYCANONICALTYPE_MODELS_BY_CANONICAL AiGatewayModelsByCanonicalType = "models-by-canonical"
)

var allowedAiGatewayModelsByCanonicalTypeEnumValues = []AiGatewayModelsByCanonicalType{
	AIGATEWAYMODELSBYCANONICALTYPE_MODELS_BY_CANONICAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayModelsByCanonicalType) GetAllowedValues() []AiGatewayModelsByCanonicalType {
	return allowedAiGatewayModelsByCanonicalTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayModelsByCanonicalType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayModelsByCanonicalType(value)
	return nil
}

// NewAiGatewayModelsByCanonicalTypeFromValue returns a pointer to a valid AiGatewayModelsByCanonicalType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayModelsByCanonicalTypeFromValue(v string) (*AiGatewayModelsByCanonicalType, error) {
	ev := AiGatewayModelsByCanonicalType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayModelsByCanonicalType: valid values are %v", v, allowedAiGatewayModelsByCanonicalTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayModelsByCanonicalType) IsValid() bool {
	for _, existing := range allowedAiGatewayModelsByCanonicalTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayModelsByCanonicalType value.
func (v AiGatewayModelsByCanonicalType) Ptr() *AiGatewayModelsByCanonicalType {
	return &v
}
