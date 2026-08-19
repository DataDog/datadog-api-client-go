// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// McpCrossAppAccessIssuerUrlType Data type of an MCP Cross-App Access issuer URL update.
type McpCrossAppAccessIssuerUrlType string

// List of McpCrossAppAccessIssuerUrlType.
const (
	MCPCROSSAPPACCESSISSUERURLTYPE_ORG_CONFIG McpCrossAppAccessIssuerUrlType = "org_config"
)

var allowedMcpCrossAppAccessIssuerUrlTypeEnumValues = []McpCrossAppAccessIssuerUrlType{
	MCPCROSSAPPACCESSISSUERURLTYPE_ORG_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *McpCrossAppAccessIssuerUrlType) GetAllowedValues() []McpCrossAppAccessIssuerUrlType {
	return allowedMcpCrossAppAccessIssuerUrlTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *McpCrossAppAccessIssuerUrlType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = McpCrossAppAccessIssuerUrlType(value)
	return nil
}

// NewMcpCrossAppAccessIssuerUrlTypeFromValue returns a pointer to a valid McpCrossAppAccessIssuerUrlType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMcpCrossAppAccessIssuerUrlTypeFromValue(v string) (*McpCrossAppAccessIssuerUrlType, error) {
	ev := McpCrossAppAccessIssuerUrlType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for McpCrossAppAccessIssuerUrlType: valid values are %v", v, allowedMcpCrossAppAccessIssuerUrlTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v McpCrossAppAccessIssuerUrlType) IsValid() bool {
	for _, existing := range allowedMcpCrossAppAccessIssuerUrlTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to McpCrossAppAccessIssuerUrlType value.
func (v McpCrossAppAccessIssuerUrlType) Ptr() *McpCrossAppAccessIssuerUrlType {
	return &v
}
