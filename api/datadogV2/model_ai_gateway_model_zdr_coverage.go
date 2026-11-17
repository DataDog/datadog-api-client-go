// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelZdrCoverage Datadog's coverage of the model's operations under zero data retention. A model is eligible when
// coverage is `full` and its provider's `zdr_eligible` is `true`. `partial` means only the
// `covered_operations` can be served: requests on other operations are not dispatched to it.
type AiGatewayModelZdrCoverage string

// List of AiGatewayModelZdrCoverage.
const (
	AIGATEWAYMODELZDRCOVERAGE_FULL    AiGatewayModelZdrCoverage = "full"
	AIGATEWAYMODELZDRCOVERAGE_PARTIAL AiGatewayModelZdrCoverage = "partial"
	AIGATEWAYMODELZDRCOVERAGE_NONE    AiGatewayModelZdrCoverage = "none"
)

var allowedAiGatewayModelZdrCoverageEnumValues = []AiGatewayModelZdrCoverage{
	AIGATEWAYMODELZDRCOVERAGE_FULL,
	AIGATEWAYMODELZDRCOVERAGE_PARTIAL,
	AIGATEWAYMODELZDRCOVERAGE_NONE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayModelZdrCoverage) GetAllowedValues() []AiGatewayModelZdrCoverage {
	return allowedAiGatewayModelZdrCoverageEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayModelZdrCoverage) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayModelZdrCoverage(value)
	return nil
}

// NewAiGatewayModelZdrCoverageFromValue returns a pointer to a valid AiGatewayModelZdrCoverage
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayModelZdrCoverageFromValue(v string) (*AiGatewayModelZdrCoverage, error) {
	ev := AiGatewayModelZdrCoverage(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayModelZdrCoverage: valid values are %v", v, allowedAiGatewayModelZdrCoverageEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayModelZdrCoverage) IsValid() bool {
	for _, existing := range allowedAiGatewayModelZdrCoverageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayModelZdrCoverage value.
func (v AiGatewayModelZdrCoverage) Ptr() *AiGatewayModelZdrCoverage {
	return &v
}
