// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelPricingType Model pricing resource type.
type AiGatewayModelPricingType string

// List of AiGatewayModelPricingType.
const (
	AIGATEWAYMODELPRICINGTYPE_MODEL_PRICING AiGatewayModelPricingType = "model-pricing"
)

var allowedAiGatewayModelPricingTypeEnumValues = []AiGatewayModelPricingType{
	AIGATEWAYMODELPRICINGTYPE_MODEL_PRICING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayModelPricingType) GetAllowedValues() []AiGatewayModelPricingType {
	return allowedAiGatewayModelPricingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayModelPricingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayModelPricingType(value)
	return nil
}

// NewAiGatewayModelPricingTypeFromValue returns a pointer to a valid AiGatewayModelPricingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayModelPricingTypeFromValue(v string) (*AiGatewayModelPricingType, error) {
	ev := AiGatewayModelPricingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayModelPricingType: valid values are %v", v, allowedAiGatewayModelPricingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayModelPricingType) IsValid() bool {
	for _, existing := range allowedAiGatewayModelPricingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayModelPricingType value.
func (v AiGatewayModelPricingType) Ptr() *AiGatewayModelPricingType {
	return &v
}
