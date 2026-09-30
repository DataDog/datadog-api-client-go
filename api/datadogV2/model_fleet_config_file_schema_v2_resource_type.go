// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetConfigFileSchemaV2ResourceType The type of the configuration file schema resource.
type FleetConfigFileSchemaV2ResourceType string

// List of FleetConfigFileSchemaV2ResourceType.
const (
	FLEETCONFIGFILESCHEMAV2RESOURCETYPE_CONFIG_FILE_SCHEMA FleetConfigFileSchemaV2ResourceType = "config_file_schema"
)

var allowedFleetConfigFileSchemaV2ResourceTypeEnumValues = []FleetConfigFileSchemaV2ResourceType{
	FLEETCONFIGFILESCHEMAV2RESOURCETYPE_CONFIG_FILE_SCHEMA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FleetConfigFileSchemaV2ResourceType) GetAllowedValues() []FleetConfigFileSchemaV2ResourceType {
	return allowedFleetConfigFileSchemaV2ResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FleetConfigFileSchemaV2ResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FleetConfigFileSchemaV2ResourceType(value)
	return nil
}

// NewFleetConfigFileSchemaV2ResourceTypeFromValue returns a pointer to a valid FleetConfigFileSchemaV2ResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFleetConfigFileSchemaV2ResourceTypeFromValue(v string) (*FleetConfigFileSchemaV2ResourceType, error) {
	ev := FleetConfigFileSchemaV2ResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FleetConfigFileSchemaV2ResourceType: valid values are %v", v, allowedFleetConfigFileSchemaV2ResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FleetConfigFileSchemaV2ResourceType) IsValid() bool {
	for _, existing := range allowedFleetConfigFileSchemaV2ResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FleetConfigFileSchemaV2ResourceType value.
func (v FleetConfigFileSchemaV2ResourceType) Ptr() *FleetConfigFileSchemaV2ResourceType {
	return &v
}
