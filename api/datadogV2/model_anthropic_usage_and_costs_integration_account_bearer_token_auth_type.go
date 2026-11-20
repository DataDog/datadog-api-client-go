// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType The authentication method type.
type AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType string

// List of AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType.
const (
	ANTHROPICUSAGEANDCOSTSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType = "bearer_token"
)

var allowedAnthropicUsageAndCostsIntegrationAccountBearerTokenAuthTypeEnumValues = []AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType{
	ANTHROPICUSAGEANDCOSTSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType) GetAllowedValues() []AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType {
	return allowedAnthropicUsageAndCostsIntegrationAccountBearerTokenAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType(value)
	return nil
}

// NewAnthropicUsageAndCostsIntegrationAccountBearerTokenAuthTypeFromValue returns a pointer to a valid AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAnthropicUsageAndCostsIntegrationAccountBearerTokenAuthTypeFromValue(v string) (*AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType, error) {
	ev := AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType: valid values are %v", v, allowedAnthropicUsageAndCostsIntegrationAccountBearerTokenAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType) IsValid() bool {
	for _, existing := range allowedAnthropicUsageAndCostsIntegrationAccountBearerTokenAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType value.
func (v AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType) Ptr() *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType {
	return &v
}
