// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRateLimitWindowLimitMetric The metric that a window limit counts. `cost` counts priced spend in microcents.
type AiGatewayRateLimitWindowLimitMetric string

// List of AiGatewayRateLimitWindowLimitMetric.
const (
	AIGATEWAYRATELIMITWINDOWLIMITMETRIC_REQUESTS      AiGatewayRateLimitWindowLimitMetric = "requests"
	AIGATEWAYRATELIMITWINDOWLIMITMETRIC_INPUT_TOKENS  AiGatewayRateLimitWindowLimitMetric = "input_tokens"
	AIGATEWAYRATELIMITWINDOWLIMITMETRIC_OUTPUT_TOKENS AiGatewayRateLimitWindowLimitMetric = "output_tokens"
	AIGATEWAYRATELIMITWINDOWLIMITMETRIC_COST          AiGatewayRateLimitWindowLimitMetric = "cost"
)

var allowedAiGatewayRateLimitWindowLimitMetricEnumValues = []AiGatewayRateLimitWindowLimitMetric{
	AIGATEWAYRATELIMITWINDOWLIMITMETRIC_REQUESTS,
	AIGATEWAYRATELIMITWINDOWLIMITMETRIC_INPUT_TOKENS,
	AIGATEWAYRATELIMITWINDOWLIMITMETRIC_OUTPUT_TOKENS,
	AIGATEWAYRATELIMITWINDOWLIMITMETRIC_COST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayRateLimitWindowLimitMetric) GetAllowedValues() []AiGatewayRateLimitWindowLimitMetric {
	return allowedAiGatewayRateLimitWindowLimitMetricEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayRateLimitWindowLimitMetric) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayRateLimitWindowLimitMetric(value)
	return nil
}

// NewAiGatewayRateLimitWindowLimitMetricFromValue returns a pointer to a valid AiGatewayRateLimitWindowLimitMetric
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayRateLimitWindowLimitMetricFromValue(v string) (*AiGatewayRateLimitWindowLimitMetric, error) {
	ev := AiGatewayRateLimitWindowLimitMetric(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayRateLimitWindowLimitMetric: valid values are %v", v, allowedAiGatewayRateLimitWindowLimitMetricEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayRateLimitWindowLimitMetric) IsValid() bool {
	for _, existing := range allowedAiGatewayRateLimitWindowLimitMetricEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayRateLimitWindowLimitMetric value.
func (v AiGatewayRateLimitWindowLimitMetric) Ptr() *AiGatewayRateLimitWindowLimitMetric {
	return &v
}
