// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationAccountPatAuthType The authentication method type.
type DatabricksIntegrationAccountPatAuthType string

// List of DatabricksIntegrationAccountPatAuthType.
const (
	DATABRICKSINTEGRATIONACCOUNTPATAUTHTYPE_PAT DatabricksIntegrationAccountPatAuthType = "pat"
)

var allowedDatabricksIntegrationAccountPatAuthTypeEnumValues = []DatabricksIntegrationAccountPatAuthType{
	DATABRICKSINTEGRATIONACCOUNTPATAUTHTYPE_PAT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatabricksIntegrationAccountPatAuthType) GetAllowedValues() []DatabricksIntegrationAccountPatAuthType {
	return allowedDatabricksIntegrationAccountPatAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatabricksIntegrationAccountPatAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatabricksIntegrationAccountPatAuthType(value)
	return nil
}

// NewDatabricksIntegrationAccountPatAuthTypeFromValue returns a pointer to a valid DatabricksIntegrationAccountPatAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatabricksIntegrationAccountPatAuthTypeFromValue(v string) (*DatabricksIntegrationAccountPatAuthType, error) {
	ev := DatabricksIntegrationAccountPatAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatabricksIntegrationAccountPatAuthType: valid values are %v", v, allowedDatabricksIntegrationAccountPatAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatabricksIntegrationAccountPatAuthType) IsValid() bool {
	for _, existing := range allowedDatabricksIntegrationAccountPatAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatabricksIntegrationAccountPatAuthType value.
func (v DatabricksIntegrationAccountPatAuthType) Ptr() *DatabricksIntegrationAccountPatAuthType {
	return &v
}
