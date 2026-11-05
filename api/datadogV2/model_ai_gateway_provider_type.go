// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayProviderType AI Gateway provider resource type.
type AiGatewayProviderType string

// List of AiGatewayProviderType.
const (
	AIGATEWAYPROVIDERTYPE_PROVIDERS AiGatewayProviderType = "providers"
)

var allowedAiGatewayProviderTypeEnumValues = []AiGatewayProviderType{
	AIGATEWAYPROVIDERTYPE_PROVIDERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayProviderType) GetAllowedValues() []AiGatewayProviderType {
	return allowedAiGatewayProviderTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayProviderType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayProviderType(value)
	return nil
}

// NewAiGatewayProviderTypeFromValue returns a pointer to a valid AiGatewayProviderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayProviderTypeFromValue(v string) (*AiGatewayProviderType, error) {
	ev := AiGatewayProviderType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayProviderType: valid values are %v", v, allowedAiGatewayProviderTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayProviderType) IsValid() bool {
	for _, existing := range allowedAiGatewayProviderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayProviderType value.
func (v AiGatewayProviderType) Ptr() *AiGatewayProviderType {
	return &v
}
