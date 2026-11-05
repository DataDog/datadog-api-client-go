// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelPricingRequestConfig - Pricing configuration accepted when setting model pricing.
type AiGatewayModelPricingRequestConfig struct {
	AiGatewayModelPricingCustomRequestConfig  *AiGatewayModelPricingCustomRequestConfig
	AiGatewayModelPricingCatalogRequestConfig *AiGatewayModelPricingCatalogRequestConfig

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AiGatewayModelPricingCustomRequestConfigAsAiGatewayModelPricingRequestConfig is a convenience function that returns AiGatewayModelPricingCustomRequestConfig wrapped in AiGatewayModelPricingRequestConfig.
func AiGatewayModelPricingCustomRequestConfigAsAiGatewayModelPricingRequestConfig(v *AiGatewayModelPricingCustomRequestConfig) AiGatewayModelPricingRequestConfig {
	return AiGatewayModelPricingRequestConfig{AiGatewayModelPricingCustomRequestConfig: v}
}

// AiGatewayModelPricingCatalogRequestConfigAsAiGatewayModelPricingRequestConfig is a convenience function that returns AiGatewayModelPricingCatalogRequestConfig wrapped in AiGatewayModelPricingRequestConfig.
func AiGatewayModelPricingCatalogRequestConfigAsAiGatewayModelPricingRequestConfig(v *AiGatewayModelPricingCatalogRequestConfig) AiGatewayModelPricingRequestConfig {
	return AiGatewayModelPricingRequestConfig{AiGatewayModelPricingCatalogRequestConfig: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AiGatewayModelPricingRequestConfig) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AiGatewayModelPricingCustomRequestConfig
	err = datadog.Unmarshal(data, &obj.AiGatewayModelPricingCustomRequestConfig)
	if err == nil {
		if obj.AiGatewayModelPricingCustomRequestConfig != nil && obj.AiGatewayModelPricingCustomRequestConfig.UnparsedObject == nil {
			jsonAiGatewayModelPricingCustomRequestConfig, _ := datadog.Marshal(obj.AiGatewayModelPricingCustomRequestConfig)
			if string(jsonAiGatewayModelPricingCustomRequestConfig) == "{}" { // empty struct
				obj.AiGatewayModelPricingCustomRequestConfig = nil
			} else {
				match++
			}
		} else {
			obj.AiGatewayModelPricingCustomRequestConfig = nil
		}
	} else {
		obj.AiGatewayModelPricingCustomRequestConfig = nil
	}

	// try to unmarshal data into AiGatewayModelPricingCatalogRequestConfig
	err = datadog.Unmarshal(data, &obj.AiGatewayModelPricingCatalogRequestConfig)
	if err == nil {
		if obj.AiGatewayModelPricingCatalogRequestConfig != nil && obj.AiGatewayModelPricingCatalogRequestConfig.UnparsedObject == nil {
			jsonAiGatewayModelPricingCatalogRequestConfig, _ := datadog.Marshal(obj.AiGatewayModelPricingCatalogRequestConfig)
			if string(jsonAiGatewayModelPricingCatalogRequestConfig) == "{}" { // empty struct
				obj.AiGatewayModelPricingCatalogRequestConfig = nil
			} else {
				match++
			}
		} else {
			obj.AiGatewayModelPricingCatalogRequestConfig = nil
		}
	} else {
		obj.AiGatewayModelPricingCatalogRequestConfig = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AiGatewayModelPricingCustomRequestConfig = nil
		obj.AiGatewayModelPricingCatalogRequestConfig = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AiGatewayModelPricingRequestConfig) MarshalJSON() ([]byte, error) {
	if obj.AiGatewayModelPricingCustomRequestConfig != nil {
		return datadog.Marshal(&obj.AiGatewayModelPricingCustomRequestConfig)
	}

	if obj.AiGatewayModelPricingCatalogRequestConfig != nil {
		return datadog.Marshal(&obj.AiGatewayModelPricingCatalogRequestConfig)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AiGatewayModelPricingRequestConfig) GetActualInstance() interface{} {
	if obj.AiGatewayModelPricingCustomRequestConfig != nil {
		return obj.AiGatewayModelPricingCustomRequestConfig
	}

	if obj.AiGatewayModelPricingCatalogRequestConfig != nil {
		return obj.AiGatewayModelPricingCatalogRequestConfig
	}

	// all schemas are nil
	return nil
}
