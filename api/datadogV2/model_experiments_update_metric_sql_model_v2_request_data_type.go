// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsUpdateMetricSQLModelV2RequestDataType Metric SQL models resource type.
type ExperimentsUpdateMetricSQLModelV2RequestDataType string

// List of ExperimentsUpdateMetricSQLModelV2RequestDataType.
const (
	EXPERIMENTSUPDATEMETRICSQLMODELV2REQUESTDATATYPE_METRIC_SQL_MODELS ExperimentsUpdateMetricSQLModelV2RequestDataType = "metric-sql-models"
)

var allowedExperimentsUpdateMetricSQLModelV2RequestDataTypeEnumValues = []ExperimentsUpdateMetricSQLModelV2RequestDataType{
	EXPERIMENTSUPDATEMETRICSQLMODELV2REQUESTDATATYPE_METRIC_SQL_MODELS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsUpdateMetricSQLModelV2RequestDataType) GetAllowedValues() []ExperimentsUpdateMetricSQLModelV2RequestDataType {
	return allowedExperimentsUpdateMetricSQLModelV2RequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsUpdateMetricSQLModelV2RequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsUpdateMetricSQLModelV2RequestDataType(value)
	return nil
}

// NewExperimentsUpdateMetricSQLModelV2RequestDataTypeFromValue returns a pointer to a valid ExperimentsUpdateMetricSQLModelV2RequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsUpdateMetricSQLModelV2RequestDataTypeFromValue(v string) (*ExperimentsUpdateMetricSQLModelV2RequestDataType, error) {
	ev := ExperimentsUpdateMetricSQLModelV2RequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsUpdateMetricSQLModelV2RequestDataType: valid values are %v", v, allowedExperimentsUpdateMetricSQLModelV2RequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsUpdateMetricSQLModelV2RequestDataType) IsValid() bool {
	for _, existing := range allowedExperimentsUpdateMetricSQLModelV2RequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsUpdateMetricSQLModelV2RequestDataType value.
func (v ExperimentsUpdateMetricSQLModelV2RequestDataType) Ptr() *ExperimentsUpdateMetricSQLModelV2RequestDataType {
	return &v
}
