// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPatchMetricCollectionV2RequestDataType Metric collections resource type.
type ExperimentsPatchMetricCollectionV2RequestDataType string

// List of ExperimentsPatchMetricCollectionV2RequestDataType.
const (
	EXPERIMENTSPATCHMETRICCOLLECTIONV2REQUESTDATATYPE_METRIC_COLLECTIONS ExperimentsPatchMetricCollectionV2RequestDataType = "metric-collections"
)

var allowedExperimentsPatchMetricCollectionV2RequestDataTypeEnumValues = []ExperimentsPatchMetricCollectionV2RequestDataType{
	EXPERIMENTSPATCHMETRICCOLLECTIONV2REQUESTDATATYPE_METRIC_COLLECTIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsPatchMetricCollectionV2RequestDataType) GetAllowedValues() []ExperimentsPatchMetricCollectionV2RequestDataType {
	return allowedExperimentsPatchMetricCollectionV2RequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsPatchMetricCollectionV2RequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsPatchMetricCollectionV2RequestDataType(value)
	return nil
}

// NewExperimentsPatchMetricCollectionV2RequestDataTypeFromValue returns a pointer to a valid ExperimentsPatchMetricCollectionV2RequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsPatchMetricCollectionV2RequestDataTypeFromValue(v string) (*ExperimentsPatchMetricCollectionV2RequestDataType, error) {
	ev := ExperimentsPatchMetricCollectionV2RequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsPatchMetricCollectionV2RequestDataType: valid values are %v", v, allowedExperimentsPatchMetricCollectionV2RequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsPatchMetricCollectionV2RequestDataType) IsValid() bool {
	for _, existing := range allowedExperimentsPatchMetricCollectionV2RequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsPatchMetricCollectionV2RequestDataType value.
func (v ExperimentsPatchMetricCollectionV2RequestDataType) Ptr() *ExperimentsPatchMetricCollectionV2RequestDataType {
	return &v
}
