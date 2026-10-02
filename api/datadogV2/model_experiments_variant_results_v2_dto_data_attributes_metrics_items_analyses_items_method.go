// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod Statistical method used to calculate this result.
type ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod string

// List of ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod.
const (
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSMETHOD_FIXED_SAMPLE            ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod = "FIXED_SAMPLE"
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSMETHOD_BAYESIAN                ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod = "BAYESIAN"
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSMETHOD_SEQUENTIAL              ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod = "SEQUENTIAL"
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSMETHOD_SEQUENTIAL_FIXED_HYBRID ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod = "SEQUENTIAL_FIXED_HYBRID"
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSMETHOD_UNKNOWN                 ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod = "UNKNOWN"
)

var allowedExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethodEnumValues = []ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod{
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSMETHOD_FIXED_SAMPLE,
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSMETHOD_BAYESIAN,
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSMETHOD_SEQUENTIAL,
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSMETHOD_SEQUENTIAL_FIXED_HYBRID,
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSMETHOD_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod) GetAllowedValues() []ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod {
	return allowedExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethodEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod(value)
	return nil
}

// NewExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethodFromValue returns a pointer to a valid ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethodFromValue(v string) (*ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod, error) {
	ev := ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod: valid values are %v", v, allowedExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethodEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod) IsValid() bool {
	for _, existing := range allowedExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod value.
func (v ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod) Ptr() *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsMethod {
	return &v
}
