// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRateLimitInstancesSummaryType AI Gateway rate limit counter summary resource type.
type AiGatewayRateLimitInstancesSummaryType string

// List of AiGatewayRateLimitInstancesSummaryType.
const (
	AIGATEWAYRATELIMITINSTANCESSUMMARYTYPE_RATE_LIMIT_INSTANCES_SUMMARY AiGatewayRateLimitInstancesSummaryType = "rate-limit-instances-summary"
)

var allowedAiGatewayRateLimitInstancesSummaryTypeEnumValues = []AiGatewayRateLimitInstancesSummaryType{
	AIGATEWAYRATELIMITINSTANCESSUMMARYTYPE_RATE_LIMIT_INSTANCES_SUMMARY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayRateLimitInstancesSummaryType) GetAllowedValues() []AiGatewayRateLimitInstancesSummaryType {
	return allowedAiGatewayRateLimitInstancesSummaryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayRateLimitInstancesSummaryType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayRateLimitInstancesSummaryType(value)
	return nil
}

// NewAiGatewayRateLimitInstancesSummaryTypeFromValue returns a pointer to a valid AiGatewayRateLimitInstancesSummaryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayRateLimitInstancesSummaryTypeFromValue(v string) (*AiGatewayRateLimitInstancesSummaryType, error) {
	ev := AiGatewayRateLimitInstancesSummaryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayRateLimitInstancesSummaryType: valid values are %v", v, allowedAiGatewayRateLimitInstancesSummaryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayRateLimitInstancesSummaryType) IsValid() bool {
	for _, existing := range allowedAiGatewayRateLimitInstancesSummaryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayRateLimitInstancesSummaryType value.
func (v AiGatewayRateLimitInstancesSummaryType) Ptr() *AiGatewayRateLimitInstancesSummaryType {
	return &v
}
