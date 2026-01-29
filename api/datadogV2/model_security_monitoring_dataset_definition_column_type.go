// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetDefinitionColumnType The type of the column in a dataset definition.
type SecurityMonitoringDatasetDefinitionColumnType string

// List of SecurityMonitoringDatasetDefinitionColumnType.
const (
	SECURITYMONITORINGDATASETDEFINITIONCOLUMNTYPE_STRING  SecurityMonitoringDatasetDefinitionColumnType = "string"
	SECURITYMONITORINGDATASETDEFINITIONCOLUMNTYPE_INTEGER SecurityMonitoringDatasetDefinitionColumnType = "integer"
	SECURITYMONITORINGDATASETDEFINITIONCOLUMNTYPE_DOUBLE  SecurityMonitoringDatasetDefinitionColumnType = "double"
	SECURITYMONITORINGDATASETDEFINITIONCOLUMNTYPE_BOOLEAN SecurityMonitoringDatasetDefinitionColumnType = "boolean"
)

var allowedSecurityMonitoringDatasetDefinitionColumnTypeEnumValues = []SecurityMonitoringDatasetDefinitionColumnType{
	SECURITYMONITORINGDATASETDEFINITIONCOLUMNTYPE_STRING,
	SECURITYMONITORINGDATASETDEFINITIONCOLUMNTYPE_INTEGER,
	SECURITYMONITORINGDATASETDEFINITIONCOLUMNTYPE_DOUBLE,
	SECURITYMONITORINGDATASETDEFINITIONCOLUMNTYPE_BOOLEAN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringDatasetDefinitionColumnType) GetAllowedValues() []SecurityMonitoringDatasetDefinitionColumnType {
	return allowedSecurityMonitoringDatasetDefinitionColumnTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringDatasetDefinitionColumnType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringDatasetDefinitionColumnType(value)
	return nil
}

// NewSecurityMonitoringDatasetDefinitionColumnTypeFromValue returns a pointer to a valid SecurityMonitoringDatasetDefinitionColumnType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringDatasetDefinitionColumnTypeFromValue(v string) (*SecurityMonitoringDatasetDefinitionColumnType, error) {
	ev := SecurityMonitoringDatasetDefinitionColumnType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringDatasetDefinitionColumnType: valid values are %v", v, allowedSecurityMonitoringDatasetDefinitionColumnTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringDatasetDefinitionColumnType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringDatasetDefinitionColumnTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringDatasetDefinitionColumnType value.
func (v SecurityMonitoringDatasetDefinitionColumnType) Ptr() *SecurityMonitoringDatasetDefinitionColumnType {
	return &v
}
