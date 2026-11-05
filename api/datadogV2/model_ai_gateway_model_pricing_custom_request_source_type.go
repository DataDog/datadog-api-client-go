// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelPricingCustomRequestSourceType Custom pricing source type.
type AiGatewayModelPricingCustomRequestSourceType string

// List of AiGatewayModelPricingCustomRequestSourceType.
const (
	AIGATEWAYMODELPRICINGCUSTOMREQUESTSOURCETYPE_CUSTOM AiGatewayModelPricingCustomRequestSourceType = "custom"
)

var allowedAiGatewayModelPricingCustomRequestSourceTypeEnumValues = []AiGatewayModelPricingCustomRequestSourceType{
	AIGATEWAYMODELPRICINGCUSTOMREQUESTSOURCETYPE_CUSTOM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayModelPricingCustomRequestSourceType) GetAllowedValues() []AiGatewayModelPricingCustomRequestSourceType {
	return allowedAiGatewayModelPricingCustomRequestSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayModelPricingCustomRequestSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayModelPricingCustomRequestSourceType(value)
	return nil
}

// NewAiGatewayModelPricingCustomRequestSourceTypeFromValue returns a pointer to a valid AiGatewayModelPricingCustomRequestSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayModelPricingCustomRequestSourceTypeFromValue(v string) (*AiGatewayModelPricingCustomRequestSourceType, error) {
	ev := AiGatewayModelPricingCustomRequestSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayModelPricingCustomRequestSourceType: valid values are %v", v, allowedAiGatewayModelPricingCustomRequestSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayModelPricingCustomRequestSourceType) IsValid() bool {
	for _, existing := range allowedAiGatewayModelPricingCustomRequestSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayModelPricingCustomRequestSourceType value.
func (v AiGatewayModelPricingCustomRequestSourceType) Ptr() *AiGatewayModelPricingCustomRequestSourceType {
	return &v
}
