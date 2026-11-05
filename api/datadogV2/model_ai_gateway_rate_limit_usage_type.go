// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRateLimitUsageType AI Gateway rate limit usage resource type.
type AiGatewayRateLimitUsageType string

// List of AiGatewayRateLimitUsageType.
const (
	AIGATEWAYRATELIMITUSAGETYPE_RATE_LIMIT_USAGE AiGatewayRateLimitUsageType = "rate-limit-usage"
)

var allowedAiGatewayRateLimitUsageTypeEnumValues = []AiGatewayRateLimitUsageType{
	AIGATEWAYRATELIMITUSAGETYPE_RATE_LIMIT_USAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayRateLimitUsageType) GetAllowedValues() []AiGatewayRateLimitUsageType {
	return allowedAiGatewayRateLimitUsageTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayRateLimitUsageType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayRateLimitUsageType(value)
	return nil
}

// NewAiGatewayRateLimitUsageTypeFromValue returns a pointer to a valid AiGatewayRateLimitUsageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayRateLimitUsageTypeFromValue(v string) (*AiGatewayRateLimitUsageType, error) {
	ev := AiGatewayRateLimitUsageType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayRateLimitUsageType: valid values are %v", v, allowedAiGatewayRateLimitUsageTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayRateLimitUsageType) IsValid() bool {
	for _, existing := range allowedAiGatewayRateLimitUsageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayRateLimitUsageType value.
func (v AiGatewayRateLimitUsageType) Ptr() *AiGatewayRateLimitUsageType {
	return &v
}
