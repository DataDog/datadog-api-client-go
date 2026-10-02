// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType Whether the reported lift is relative or absolute.
type ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType string

// List of ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType.
const (
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSLIFTTYPE_RELATIVE ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType = "RELATIVE"
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSLIFTTYPE_ABSOLUTE ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType = "ABSOLUTE"
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSLIFTTYPE_UNKNOWN  ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType = "UNKNOWN"
)

var allowedExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftTypeEnumValues = []ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType{
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSLIFTTYPE_RELATIVE,
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSLIFTTYPE_ABSOLUTE,
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSLIFTTYPE_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType) GetAllowedValues() []ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType {
	return allowedExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType(value)
	return nil
}

// NewExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftTypeFromValue returns a pointer to a valid ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftTypeFromValue(v string) (*ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType, error) {
	ev := ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType: valid values are %v", v, allowedExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType) IsValid() bool {
	for _, existing := range allowedExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType value.
func (v ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType) Ptr() *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsLiftType {
	return &v
}
