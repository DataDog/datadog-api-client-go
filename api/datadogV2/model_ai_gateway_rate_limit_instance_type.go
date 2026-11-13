// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRateLimitInstanceType AI Gateway rate limit counter resource type.
type AiGatewayRateLimitInstanceType string

// List of AiGatewayRateLimitInstanceType.
const (
	AIGATEWAYRATELIMITINSTANCETYPE_RATE_LIMIT_INSTANCES AiGatewayRateLimitInstanceType = "rate-limit-instances"
)

var allowedAiGatewayRateLimitInstanceTypeEnumValues = []AiGatewayRateLimitInstanceType{
	AIGATEWAYRATELIMITINSTANCETYPE_RATE_LIMIT_INSTANCES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayRateLimitInstanceType) GetAllowedValues() []AiGatewayRateLimitInstanceType {
	return allowedAiGatewayRateLimitInstanceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayRateLimitInstanceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayRateLimitInstanceType(value)
	return nil
}

// NewAiGatewayRateLimitInstanceTypeFromValue returns a pointer to a valid AiGatewayRateLimitInstanceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayRateLimitInstanceTypeFromValue(v string) (*AiGatewayRateLimitInstanceType, error) {
	ev := AiGatewayRateLimitInstanceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayRateLimitInstanceType: valid values are %v", v, allowedAiGatewayRateLimitInstanceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayRateLimitInstanceType) IsValid() bool {
	for _, existing := range allowedAiGatewayRateLimitInstanceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayRateLimitInstanceType value.
func (v AiGatewayRateLimitInstanceType) Ptr() *AiGatewayRateLimitInstanceType {
	return &v
}
