// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayKnownProvider Known provider implementation.
type AiGatewayKnownProvider string

// List of AiGatewayKnownProvider.
const (
	AIGATEWAYKNOWNPROVIDER_CUSTOM        AiGatewayKnownProvider = "custom"
	AIGATEWAYKNOWNPROVIDER_ANTHROPIC     AiGatewayKnownProvider = "anthropic"
	AIGATEWAYKNOWNPROVIDER_OPENAI        AiGatewayKnownProvider = "openai"
	AIGATEWAYKNOWNPROVIDER_AZURE_OPENAI  AiGatewayKnownProvider = "azure_openai"
	AIGATEWAYKNOWNPROVIDER_BEDROCK       AiGatewayKnownProvider = "bedrock"
	AIGATEWAYKNOWNPROVIDER_GEMINI        AiGatewayKnownProvider = "gemini"
	AIGATEWAYKNOWNPROVIDER_BASETEN       AiGatewayKnownProvider = "baseten"
	AIGATEWAYKNOWNPROVIDER_FIREWORKS     AiGatewayKnownProvider = "fireworks"
	AIGATEWAYKNOWNPROVIDER_TOGETHER      AiGatewayKnownProvider = "together"
	AIGATEWAYKNOWNPROVIDER_TYPESAFE      AiGatewayKnownProvider = "typesafe"
	AIGATEWAYKNOWNPROVIDER_AZURE_FOUNDRY AiGatewayKnownProvider = "azure_foundry"
)

var allowedAiGatewayKnownProviderEnumValues = []AiGatewayKnownProvider{
	AIGATEWAYKNOWNPROVIDER_CUSTOM,
	AIGATEWAYKNOWNPROVIDER_ANTHROPIC,
	AIGATEWAYKNOWNPROVIDER_OPENAI,
	AIGATEWAYKNOWNPROVIDER_AZURE_OPENAI,
	AIGATEWAYKNOWNPROVIDER_BEDROCK,
	AIGATEWAYKNOWNPROVIDER_GEMINI,
	AIGATEWAYKNOWNPROVIDER_BASETEN,
	AIGATEWAYKNOWNPROVIDER_FIREWORKS,
	AIGATEWAYKNOWNPROVIDER_TOGETHER,
	AIGATEWAYKNOWNPROVIDER_TYPESAFE,
	AIGATEWAYKNOWNPROVIDER_AZURE_FOUNDRY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayKnownProvider) GetAllowedValues() []AiGatewayKnownProvider {
	return allowedAiGatewayKnownProviderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayKnownProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayKnownProvider(value)
	return nil
}

// NewAiGatewayKnownProviderFromValue returns a pointer to a valid AiGatewayKnownProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayKnownProviderFromValue(v string) (*AiGatewayKnownProvider, error) {
	ev := AiGatewayKnownProvider(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayKnownProvider: valid values are %v", v, allowedAiGatewayKnownProviderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayKnownProvider) IsValid() bool {
	for _, existing := range allowedAiGatewayKnownProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayKnownProvider value.
func (v AiGatewayKnownProvider) Ptr() *AiGatewayKnownProvider {
	return &v
}
