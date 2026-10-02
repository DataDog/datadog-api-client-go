// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason Reason that the statistical result is marked as unreliable.
type ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason string

// List of ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason.
const (
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSUNRELIABLEREASON_CONTROL_DENOMINATOR_NEAR_ZERO                ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason = "CONTROL_DENOMINATOR_NEAR_ZERO"
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSUNRELIABLEREASON_TREATMENT_DENOMINATOR_NEAR_ZERO              ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason = "TREATMENT_DENOMINATOR_NEAR_ZERO"
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSUNRELIABLEREASON_CONTROL_AND_TREATMENT_DENOMINATORS_NEAR_ZERO ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason = "CONTROL_AND_TREATMENT_DENOMINATORS_NEAR_ZERO"
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSUNRELIABLEREASON_CONTROL_MEAN_NEAR_ZERO                       ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason = "CONTROL_MEAN_NEAR_ZERO"
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSUNRELIABLEREASON_ZERO_VARIANCE                                ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason = "ZERO_VARIANCE"
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSUNRELIABLEREASON_UNKNOWN                                      ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason = "UNKNOWN"
)

var allowedExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReasonEnumValues = []ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason{
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSUNRELIABLEREASON_CONTROL_DENOMINATOR_NEAR_ZERO,
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSUNRELIABLEREASON_TREATMENT_DENOMINATOR_NEAR_ZERO,
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSUNRELIABLEREASON_CONTROL_AND_TREATMENT_DENOMINATORS_NEAR_ZERO,
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSUNRELIABLEREASON_CONTROL_MEAN_NEAR_ZERO,
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSUNRELIABLEREASON_ZERO_VARIANCE,
	EXPERIMENTSVARIANTRESULTSV2DTODATAATTRIBUTESMETRICSITEMSANALYSESITEMSUNRELIABLEREASON_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason) GetAllowedValues() []ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason {
	return allowedExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReasonEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason(value)
	return nil
}

// NewExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReasonFromValue returns a pointer to a valid ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReasonFromValue(v string) (*ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason, error) {
	ev := ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason: valid values are %v", v, allowedExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReasonEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason) IsValid() bool {
	for _, existing := range allowedExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason value.
func (v ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason) Ptr() *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsUnreliableReason {
	return &v
}
