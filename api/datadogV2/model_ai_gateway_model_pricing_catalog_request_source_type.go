// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelPricingCatalogRequestSourceType Catalog pricing source type.
type AiGatewayModelPricingCatalogRequestSourceType string

// List of AiGatewayModelPricingCatalogRequestSourceType.
const (
	AIGATEWAYMODELPRICINGCATALOGREQUESTSOURCETYPE_CATALOG AiGatewayModelPricingCatalogRequestSourceType = "catalog"
)

var allowedAiGatewayModelPricingCatalogRequestSourceTypeEnumValues = []AiGatewayModelPricingCatalogRequestSourceType{
	AIGATEWAYMODELPRICINGCATALOGREQUESTSOURCETYPE_CATALOG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayModelPricingCatalogRequestSourceType) GetAllowedValues() []AiGatewayModelPricingCatalogRequestSourceType {
	return allowedAiGatewayModelPricingCatalogRequestSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayModelPricingCatalogRequestSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayModelPricingCatalogRequestSourceType(value)
	return nil
}

// NewAiGatewayModelPricingCatalogRequestSourceTypeFromValue returns a pointer to a valid AiGatewayModelPricingCatalogRequestSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayModelPricingCatalogRequestSourceTypeFromValue(v string) (*AiGatewayModelPricingCatalogRequestSourceType, error) {
	ev := AiGatewayModelPricingCatalogRequestSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayModelPricingCatalogRequestSourceType: valid values are %v", v, allowedAiGatewayModelPricingCatalogRequestSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayModelPricingCatalogRequestSourceType) IsValid() bool {
	for _, existing := range allowedAiGatewayModelPricingCatalogRequestSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayModelPricingCatalogRequestSourceType value.
func (v AiGatewayModelPricingCatalogRequestSourceType) Ptr() *AiGatewayModelPricingCatalogRequestSourceType {
	return &v
}
