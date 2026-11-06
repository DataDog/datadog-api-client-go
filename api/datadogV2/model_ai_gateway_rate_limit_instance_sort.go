// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRateLimitInstanceSort Sort order for rate limit counters.
type AiGatewayRateLimitInstanceSort string

// List of AiGatewayRateLimitInstanceSort.
const (
	AIGATEWAYRATELIMITINSTANCESORT_UTILIZATION_DESC AiGatewayRateLimitInstanceSort = "utilization_desc"
	AIGATEWAYRATELIMITINSTANCESORT_LAST_SEEN_DESC   AiGatewayRateLimitInstanceSort = "last_seen_desc"
	AIGATEWAYRATELIMITINSTANCESORT_KEY_ASC          AiGatewayRateLimitInstanceSort = "key_asc"
)

var allowedAiGatewayRateLimitInstanceSortEnumValues = []AiGatewayRateLimitInstanceSort{
	AIGATEWAYRATELIMITINSTANCESORT_UTILIZATION_DESC,
	AIGATEWAYRATELIMITINSTANCESORT_LAST_SEEN_DESC,
	AIGATEWAYRATELIMITINSTANCESORT_KEY_ASC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayRateLimitInstanceSort) GetAllowedValues() []AiGatewayRateLimitInstanceSort {
	return allowedAiGatewayRateLimitInstanceSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayRateLimitInstanceSort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayRateLimitInstanceSort(value)
	return nil
}

// NewAiGatewayRateLimitInstanceSortFromValue returns a pointer to a valid AiGatewayRateLimitInstanceSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayRateLimitInstanceSortFromValue(v string) (*AiGatewayRateLimitInstanceSort, error) {
	ev := AiGatewayRateLimitInstanceSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayRateLimitInstanceSort: valid values are %v", v, allowedAiGatewayRateLimitInstanceSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayRateLimitInstanceSort) IsValid() bool {
	for _, existing := range allowedAiGatewayRateLimitInstanceSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayRateLimitInstanceSort value.
func (v AiGatewayRateLimitInstanceSort) Ptr() *AiGatewayRateLimitInstanceSort {
	return &v
}
