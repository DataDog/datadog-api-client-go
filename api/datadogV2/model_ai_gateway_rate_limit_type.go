// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRateLimitType AI Gateway rate limit resource type.
type AiGatewayRateLimitType string

// List of AiGatewayRateLimitType.
const (
	AIGATEWAYRATELIMITTYPE_RATE_LIMITS AiGatewayRateLimitType = "rate-limits"
)

var allowedAiGatewayRateLimitTypeEnumValues = []AiGatewayRateLimitType{
	AIGATEWAYRATELIMITTYPE_RATE_LIMITS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayRateLimitType) GetAllowedValues() []AiGatewayRateLimitType {
	return allowedAiGatewayRateLimitTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayRateLimitType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayRateLimitType(value)
	return nil
}

// NewAiGatewayRateLimitTypeFromValue returns a pointer to a valid AiGatewayRateLimitType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayRateLimitTypeFromValue(v string) (*AiGatewayRateLimitType, error) {
	ev := AiGatewayRateLimitType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayRateLimitType: valid values are %v", v, allowedAiGatewayRateLimitTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayRateLimitType) IsValid() bool {
	for _, existing := range allowedAiGatewayRateLimitTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayRateLimitType value.
func (v AiGatewayRateLimitType) Ptr() *AiGatewayRateLimitType {
	return &v
}
