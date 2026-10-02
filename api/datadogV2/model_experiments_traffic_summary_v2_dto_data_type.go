// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsTrafficSummaryV2DTODataType Traffic summary resource type.
type ExperimentsTrafficSummaryV2DTODataType string

// List of ExperimentsTrafficSummaryV2DTODataType.
const (
	EXPERIMENTSTRAFFICSUMMARYV2DTODATATYPE_TRAFFIC_SUMMARY ExperimentsTrafficSummaryV2DTODataType = "traffic-summary"
)

var allowedExperimentsTrafficSummaryV2DTODataTypeEnumValues = []ExperimentsTrafficSummaryV2DTODataType{
	EXPERIMENTSTRAFFICSUMMARYV2DTODATATYPE_TRAFFIC_SUMMARY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsTrafficSummaryV2DTODataType) GetAllowedValues() []ExperimentsTrafficSummaryV2DTODataType {
	return allowedExperimentsTrafficSummaryV2DTODataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsTrafficSummaryV2DTODataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsTrafficSummaryV2DTODataType(value)
	return nil
}

// NewExperimentsTrafficSummaryV2DTODataTypeFromValue returns a pointer to a valid ExperimentsTrafficSummaryV2DTODataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsTrafficSummaryV2DTODataTypeFromValue(v string) (*ExperimentsTrafficSummaryV2DTODataType, error) {
	ev := ExperimentsTrafficSummaryV2DTODataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsTrafficSummaryV2DTODataType: valid values are %v", v, allowedExperimentsTrafficSummaryV2DTODataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsTrafficSummaryV2DTODataType) IsValid() bool {
	for _, existing := range allowedExperimentsTrafficSummaryV2DTODataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsTrafficSummaryV2DTODataType value.
func (v ExperimentsTrafficSummaryV2DTODataType) Ptr() *ExperimentsTrafficSummaryV2DTODataType {
	return &v
}
