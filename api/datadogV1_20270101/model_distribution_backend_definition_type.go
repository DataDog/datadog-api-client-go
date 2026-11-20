// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type DistributionBackendDefinitionType string

// List of DistributionBackendDefinitionType.
const (
	DISTRIBUTIONBACKENDDEFINITIONTYPE_DISTRIBUTION DistributionBackendDefinitionType = "distribution"
)

var allowedDistributionBackendDefinitionTypeEnumValues = []DistributionBackendDefinitionType{
	DISTRIBUTIONBACKENDDEFINITIONTYPE_DISTRIBUTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DistributionBackendDefinitionType) GetAllowedValues() []DistributionBackendDefinitionType {
	return allowedDistributionBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DistributionBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DistributionBackendDefinitionType(value)
	return nil
}

// NewDistributionBackendDefinitionTypeFromValue returns a pointer to a valid DistributionBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDistributionBackendDefinitionTypeFromValue(v string) (*DistributionBackendDefinitionType, error) {
	ev := DistributionBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DistributionBackendDefinitionType: valid values are %v", v, allowedDistributionBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DistributionBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedDistributionBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DistributionBackendDefinitionType value.
func (v DistributionBackendDefinitionType) Ptr() *DistributionBackendDefinitionType {
	return &v
}
