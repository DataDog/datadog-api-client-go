// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayPolicyType AI Gateway routing policy resource type.
type AiGatewayPolicyType string

// List of AiGatewayPolicyType.
const (
	AIGATEWAYPOLICYTYPE_ROUTING_DOCUMENTS AiGatewayPolicyType = "routing-documents"
)

var allowedAiGatewayPolicyTypeEnumValues = []AiGatewayPolicyType{
	AIGATEWAYPOLICYTYPE_ROUTING_DOCUMENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayPolicyType) GetAllowedValues() []AiGatewayPolicyType {
	return allowedAiGatewayPolicyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayPolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayPolicyType(value)
	return nil
}

// NewAiGatewayPolicyTypeFromValue returns a pointer to a valid AiGatewayPolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayPolicyTypeFromValue(v string) (*AiGatewayPolicyType, error) {
	ev := AiGatewayPolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayPolicyType: valid values are %v", v, allowedAiGatewayPolicyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayPolicyType) IsValid() bool {
	for _, existing := range allowedAiGatewayPolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayPolicyType value.
func (v AiGatewayPolicyType) Ptr() *AiGatewayPolicyType {
	return &v
}
