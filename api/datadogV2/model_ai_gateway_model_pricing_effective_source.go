// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelPricingEffectiveSource Effective source of the model pricing data.
type AiGatewayModelPricingEffectiveSource string

// List of AiGatewayModelPricingEffectiveSource.
const (
	AIGATEWAYMODELPRICINGEFFECTIVESOURCE_CATALOG_AUTO   AiGatewayModelPricingEffectiveSource = "catalog_auto"
	AIGATEWAYMODELPRICINGEFFECTIVESOURCE_MODEL_OVERRIDE AiGatewayModelPricingEffectiveSource = "model_override"
	AIGATEWAYMODELPRICINGEFFECTIVESOURCE_UNKNOWN        AiGatewayModelPricingEffectiveSource = "unknown"
)

var allowedAiGatewayModelPricingEffectiveSourceEnumValues = []AiGatewayModelPricingEffectiveSource{
	AIGATEWAYMODELPRICINGEFFECTIVESOURCE_CATALOG_AUTO,
	AIGATEWAYMODELPRICINGEFFECTIVESOURCE_MODEL_OVERRIDE,
	AIGATEWAYMODELPRICINGEFFECTIVESOURCE_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayModelPricingEffectiveSource) GetAllowedValues() []AiGatewayModelPricingEffectiveSource {
	return allowedAiGatewayModelPricingEffectiveSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayModelPricingEffectiveSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayModelPricingEffectiveSource(value)
	return nil
}

// NewAiGatewayModelPricingEffectiveSourceFromValue returns a pointer to a valid AiGatewayModelPricingEffectiveSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayModelPricingEffectiveSourceFromValue(v string) (*AiGatewayModelPricingEffectiveSource, error) {
	ev := AiGatewayModelPricingEffectiveSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayModelPricingEffectiveSource: valid values are %v", v, allowedAiGatewayModelPricingEffectiveSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayModelPricingEffectiveSource) IsValid() bool {
	for _, existing := range allowedAiGatewayModelPricingEffectiveSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayModelPricingEffectiveSource value.
func (v AiGatewayModelPricingEffectiveSource) Ptr() *AiGatewayModelPricingEffectiveSource {
	return &v
}
