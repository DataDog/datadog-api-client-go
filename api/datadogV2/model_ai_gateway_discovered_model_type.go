// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayDiscoveredModelType Discovered provider model resource type.
type AiGatewayDiscoveredModelType string

// List of AiGatewayDiscoveredModelType.
const (
	AIGATEWAYDISCOVEREDMODELTYPE_PROVIDER_MODELS AiGatewayDiscoveredModelType = "provider-models"
)

var allowedAiGatewayDiscoveredModelTypeEnumValues = []AiGatewayDiscoveredModelType{
	AIGATEWAYDISCOVEREDMODELTYPE_PROVIDER_MODELS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayDiscoveredModelType) GetAllowedValues() []AiGatewayDiscoveredModelType {
	return allowedAiGatewayDiscoveredModelTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayDiscoveredModelType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayDiscoveredModelType(value)
	return nil
}

// NewAiGatewayDiscoveredModelTypeFromValue returns a pointer to a valid AiGatewayDiscoveredModelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayDiscoveredModelTypeFromValue(v string) (*AiGatewayDiscoveredModelType, error) {
	ev := AiGatewayDiscoveredModelType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayDiscoveredModelType: valid values are %v", v, allowedAiGatewayDiscoveredModelTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayDiscoveredModelType) IsValid() bool {
	for _, existing := range allowedAiGatewayDiscoveredModelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayDiscoveredModelType value.
func (v AiGatewayDiscoveredModelType) Ptr() *AiGatewayDiscoveredModelType {
	return &v
}
