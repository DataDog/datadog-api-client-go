// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationAccountOAuthAuthType The authentication method type.
type DatabricksIntegrationAccountOAuthAuthType string

// List of DatabricksIntegrationAccountOAuthAuthType.
const (
	DATABRICKSINTEGRATIONACCOUNTOAUTHAUTHTYPE_DATABRICKS_OAUTH DatabricksIntegrationAccountOAuthAuthType = "databricks-oauth"
)

var allowedDatabricksIntegrationAccountOAuthAuthTypeEnumValues = []DatabricksIntegrationAccountOAuthAuthType{
	DATABRICKSINTEGRATIONACCOUNTOAUTHAUTHTYPE_DATABRICKS_OAUTH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatabricksIntegrationAccountOAuthAuthType) GetAllowedValues() []DatabricksIntegrationAccountOAuthAuthType {
	return allowedDatabricksIntegrationAccountOAuthAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatabricksIntegrationAccountOAuthAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatabricksIntegrationAccountOAuthAuthType(value)
	return nil
}

// NewDatabricksIntegrationAccountOAuthAuthTypeFromValue returns a pointer to a valid DatabricksIntegrationAccountOAuthAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatabricksIntegrationAccountOAuthAuthTypeFromValue(v string) (*DatabricksIntegrationAccountOAuthAuthType, error) {
	ev := DatabricksIntegrationAccountOAuthAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatabricksIntegrationAccountOAuthAuthType: valid values are %v", v, allowedDatabricksIntegrationAccountOAuthAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatabricksIntegrationAccountOAuthAuthType) IsValid() bool {
	for _, existing := range allowedDatabricksIntegrationAccountOAuthAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatabricksIntegrationAccountOAuthAuthType value.
func (v DatabricksIntegrationAccountOAuthAuthType) Ptr() *DatabricksIntegrationAccountOAuthAuthType {
	return &v
}
