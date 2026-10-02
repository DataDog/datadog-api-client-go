// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsMetricV2DTODataAttributesDataSourceType Source of the data used to calculate the metric.
type ExperimentsMetricV2DTODataAttributesDataSourceType string

// List of ExperimentsMetricV2DTODataAttributesDataSourceType.
const (
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESDATASOURCETYPE_DATADOG                 ExperimentsMetricV2DTODataAttributesDataSourceType = "DATADOG"
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESDATASOURCETYPE_DATADOG_REFERENCE_TABLE ExperimentsMetricV2DTODataAttributesDataSourceType = "DATADOG_REFERENCE_TABLE"
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESDATASOURCETYPE_CUSTOMER_WAREHOUSE      ExperimentsMetricV2DTODataAttributesDataSourceType = "CUSTOMER_WAREHOUSE"
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESDATASOURCETYPE_IMPORTED                ExperimentsMetricV2DTODataAttributesDataSourceType = "IMPORTED"
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESDATASOURCETYPE_UNKNOWN                 ExperimentsMetricV2DTODataAttributesDataSourceType = "UNKNOWN"
)

var allowedExperimentsMetricV2DTODataAttributesDataSourceTypeEnumValues = []ExperimentsMetricV2DTODataAttributesDataSourceType{
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESDATASOURCETYPE_DATADOG,
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESDATASOURCETYPE_DATADOG_REFERENCE_TABLE,
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESDATASOURCETYPE_CUSTOMER_WAREHOUSE,
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESDATASOURCETYPE_IMPORTED,
	EXPERIMENTSMETRICV2DTODATAATTRIBUTESDATASOURCETYPE_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsMetricV2DTODataAttributesDataSourceType) GetAllowedValues() []ExperimentsMetricV2DTODataAttributesDataSourceType {
	return allowedExperimentsMetricV2DTODataAttributesDataSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsMetricV2DTODataAttributesDataSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsMetricV2DTODataAttributesDataSourceType(value)
	return nil
}

// NewExperimentsMetricV2DTODataAttributesDataSourceTypeFromValue returns a pointer to a valid ExperimentsMetricV2DTODataAttributesDataSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsMetricV2DTODataAttributesDataSourceTypeFromValue(v string) (*ExperimentsMetricV2DTODataAttributesDataSourceType, error) {
	ev := ExperimentsMetricV2DTODataAttributesDataSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsMetricV2DTODataAttributesDataSourceType: valid values are %v", v, allowedExperimentsMetricV2DTODataAttributesDataSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsMetricV2DTODataAttributesDataSourceType) IsValid() bool {
	for _, existing := range allowedExperimentsMetricV2DTODataAttributesDataSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsMetricV2DTODataAttributesDataSourceType value.
func (v ExperimentsMetricV2DTODataAttributesDataSourceType) Ptr() *ExperimentsMetricV2DTODataAttributesDataSourceType {
	return &v
}
