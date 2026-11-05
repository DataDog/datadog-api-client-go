// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayGovernanceType AI Gateway governance resource type.
type AiGatewayGovernanceType string

// List of AiGatewayGovernanceType.
const (
	AIGATEWAYGOVERNANCETYPE_GOVERNANCE AiGatewayGovernanceType = "governance"
)

var allowedAiGatewayGovernanceTypeEnumValues = []AiGatewayGovernanceType{
	AIGATEWAYGOVERNANCETYPE_GOVERNANCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayGovernanceType) GetAllowedValues() []AiGatewayGovernanceType {
	return allowedAiGatewayGovernanceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayGovernanceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayGovernanceType(value)
	return nil
}

// NewAiGatewayGovernanceTypeFromValue returns a pointer to a valid AiGatewayGovernanceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayGovernanceTypeFromValue(v string) (*AiGatewayGovernanceType, error) {
	ev := AiGatewayGovernanceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayGovernanceType: valid values are %v", v, allowedAiGatewayGovernanceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayGovernanceType) IsValid() bool {
	for _, existing := range allowedAiGatewayGovernanceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayGovernanceType value.
func (v AiGatewayGovernanceType) Ptr() *AiGatewayGovernanceType {
	return &v
}
