// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsVariantResultsV2DTODataType Experiment variant results resource type.
type ExperimentsVariantResultsV2DTODataType string

// List of ExperimentsVariantResultsV2DTODataType.
const (
	EXPERIMENTSVARIANTRESULTSV2DTODATATYPE_EXPERIMENT_VARIANT_RESULTS ExperimentsVariantResultsV2DTODataType = "experiment-variant-results"
)

var allowedExperimentsVariantResultsV2DTODataTypeEnumValues = []ExperimentsVariantResultsV2DTODataType{
	EXPERIMENTSVARIANTRESULTSV2DTODATATYPE_EXPERIMENT_VARIANT_RESULTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsVariantResultsV2DTODataType) GetAllowedValues() []ExperimentsVariantResultsV2DTODataType {
	return allowedExperimentsVariantResultsV2DTODataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsVariantResultsV2DTODataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsVariantResultsV2DTODataType(value)
	return nil
}

// NewExperimentsVariantResultsV2DTODataTypeFromValue returns a pointer to a valid ExperimentsVariantResultsV2DTODataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsVariantResultsV2DTODataTypeFromValue(v string) (*ExperimentsVariantResultsV2DTODataType, error) {
	ev := ExperimentsVariantResultsV2DTODataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsVariantResultsV2DTODataType: valid values are %v", v, allowedExperimentsVariantResultsV2DTODataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsVariantResultsV2DTODataType) IsValid() bool {
	for _, existing := range allowedExperimentsVariantResultsV2DTODataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsVariantResultsV2DTODataType value.
func (v ExperimentsVariantResultsV2DTODataType) Ptr() *ExperimentsVariantResultsV2DTODataType {
	return &v
}
