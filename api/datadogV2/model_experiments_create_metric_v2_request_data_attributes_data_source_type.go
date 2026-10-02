// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateMetricV2RequestDataAttributesDataSourceType Source of the data backing this metric.
type ExperimentsCreateMetricV2RequestDataAttributesDataSourceType string

// List of ExperimentsCreateMetricV2RequestDataAttributesDataSourceType.
const (
	EXPERIMENTSCREATEMETRICV2REQUESTDATAATTRIBUTESDATASOURCETYPE_DATADOG                 ExperimentsCreateMetricV2RequestDataAttributesDataSourceType = "DATADOG"
	EXPERIMENTSCREATEMETRICV2REQUESTDATAATTRIBUTESDATASOURCETYPE_DATADOG_REFERENCE_TABLE ExperimentsCreateMetricV2RequestDataAttributesDataSourceType = "DATADOG_REFERENCE_TABLE"
	EXPERIMENTSCREATEMETRICV2REQUESTDATAATTRIBUTESDATASOURCETYPE_CUSTOMER_WAREHOUSE      ExperimentsCreateMetricV2RequestDataAttributesDataSourceType = "CUSTOMER_WAREHOUSE"
)

var allowedExperimentsCreateMetricV2RequestDataAttributesDataSourceTypeEnumValues = []ExperimentsCreateMetricV2RequestDataAttributesDataSourceType{
	EXPERIMENTSCREATEMETRICV2REQUESTDATAATTRIBUTESDATASOURCETYPE_DATADOG,
	EXPERIMENTSCREATEMETRICV2REQUESTDATAATTRIBUTESDATASOURCETYPE_DATADOG_REFERENCE_TABLE,
	EXPERIMENTSCREATEMETRICV2REQUESTDATAATTRIBUTESDATASOURCETYPE_CUSTOMER_WAREHOUSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsCreateMetricV2RequestDataAttributesDataSourceType) GetAllowedValues() []ExperimentsCreateMetricV2RequestDataAttributesDataSourceType {
	return allowedExperimentsCreateMetricV2RequestDataAttributesDataSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsCreateMetricV2RequestDataAttributesDataSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsCreateMetricV2RequestDataAttributesDataSourceType(value)
	return nil
}

// NewExperimentsCreateMetricV2RequestDataAttributesDataSourceTypeFromValue returns a pointer to a valid ExperimentsCreateMetricV2RequestDataAttributesDataSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsCreateMetricV2RequestDataAttributesDataSourceTypeFromValue(v string) (*ExperimentsCreateMetricV2RequestDataAttributesDataSourceType, error) {
	ev := ExperimentsCreateMetricV2RequestDataAttributesDataSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsCreateMetricV2RequestDataAttributesDataSourceType: valid values are %v", v, allowedExperimentsCreateMetricV2RequestDataAttributesDataSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsCreateMetricV2RequestDataAttributesDataSourceType) IsValid() bool {
	for _, existing := range allowedExperimentsCreateMetricV2RequestDataAttributesDataSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsCreateMetricV2RequestDataAttributesDataSourceType value.
func (v ExperimentsCreateMetricV2RequestDataAttributesDataSourceType) Ptr() *ExperimentsCreateMetricV2RequestDataAttributesDataSourceType {
	return &v
}
