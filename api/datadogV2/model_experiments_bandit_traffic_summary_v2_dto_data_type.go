// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditTrafficSummaryV2DTODataType Bandit traffic summary resource type.
type ExperimentsBanditTrafficSummaryV2DTODataType string

// List of ExperimentsBanditTrafficSummaryV2DTODataType.
const (
	EXPERIMENTSBANDITTRAFFICSUMMARYV2DTODATATYPE_BANDIT_TRAFFIC_SUMMARY ExperimentsBanditTrafficSummaryV2DTODataType = "bandit-traffic-summary"
)

var allowedExperimentsBanditTrafficSummaryV2DTODataTypeEnumValues = []ExperimentsBanditTrafficSummaryV2DTODataType{
	EXPERIMENTSBANDITTRAFFICSUMMARYV2DTODATATYPE_BANDIT_TRAFFIC_SUMMARY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsBanditTrafficSummaryV2DTODataType) GetAllowedValues() []ExperimentsBanditTrafficSummaryV2DTODataType {
	return allowedExperimentsBanditTrafficSummaryV2DTODataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsBanditTrafficSummaryV2DTODataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsBanditTrafficSummaryV2DTODataType(value)
	return nil
}

// NewExperimentsBanditTrafficSummaryV2DTODataTypeFromValue returns a pointer to a valid ExperimentsBanditTrafficSummaryV2DTODataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsBanditTrafficSummaryV2DTODataTypeFromValue(v string) (*ExperimentsBanditTrafficSummaryV2DTODataType, error) {
	ev := ExperimentsBanditTrafficSummaryV2DTODataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsBanditTrafficSummaryV2DTODataType: valid values are %v", v, allowedExperimentsBanditTrafficSummaryV2DTODataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsBanditTrafficSummaryV2DTODataType) IsValid() bool {
	for _, existing := range allowedExperimentsBanditTrafficSummaryV2DTODataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsBanditTrafficSummaryV2DTODataType value.
func (v ExperimentsBanditTrafficSummaryV2DTODataType) Ptr() *ExperimentsBanditTrafficSummaryV2DTODataType {
	return &v
}
