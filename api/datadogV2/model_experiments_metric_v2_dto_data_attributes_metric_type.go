// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsMetricV2DTODataAttributesMetricType Type of metric calculation.
type ExperimentsMetricV2DTODataAttributesMetricType string

// List of ExperimentsMetricV2DTODataAttributesMetricType.
const (
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESMETRICTYPE_SIMPLE     ExperimentsMetricV2DTODataAttributesMetricType = "SIMPLE"
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESMETRICTYPE_RATIO      ExperimentsMetricV2DTODataAttributesMetricType = "RATIO"
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESMETRICTYPE_PERCENTILE ExperimentsMetricV2DTODataAttributesMetricType = "PERCENTILE"
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESMETRICTYPE_UNKNOWN    ExperimentsMetricV2DTODataAttributesMetricType = "UNKNOWN"
)

var allowedExperimentsMetricV2DTODataAttributesMetricTypeEnumValues = []ExperimentsMetricV2DTODataAttributesMetricType{
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESMETRICTYPE_SIMPLE,
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESMETRICTYPE_RATIO,
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESMETRICTYPE_PERCENTILE,
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESMETRICTYPE_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsMetricV2DTODataAttributesMetricType) GetAllowedValues() []ExperimentsMetricV2DTODataAttributesMetricType {
	return allowedExperimentsMetricV2DTODataAttributesMetricTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsMetricV2DTODataAttributesMetricType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsMetricV2DTODataAttributesMetricType(value)
	return nil
}

// NewExperimentsMetricV2DTODataAttributesMetricTypeFromValue returns a pointer to a valid ExperimentsMetricV2DTODataAttributesMetricType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsMetricV2DTODataAttributesMetricTypeFromValue(v string) (*ExperimentsMetricV2DTODataAttributesMetricType, error) {
	ev := ExperimentsMetricV2DTODataAttributesMetricType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsMetricV2DTODataAttributesMetricType: valid values are %v", v, allowedExperimentsMetricV2DTODataAttributesMetricTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsMetricV2DTODataAttributesMetricType) IsValid() bool {
	for _, existing := range allowedExperimentsMetricV2DTODataAttributesMetricTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsMetricV2DTODataAttributesMetricType value.
func (v ExperimentsMetricV2DTODataAttributesMetricType) Ptr() *ExperimentsMetricV2DTODataAttributesMetricType {
	return &v
}
