// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayProviderCredentialSource Whose credential the gateway sends to the provider. `integration_account` sends the integration
// account's credential. `subscription` forwards each caller's own subscription sign-in (`ChatGPT` for
// `openai`, `claude.ai` for `anthropic`) and needs no integration account.
type AiGatewayProviderCredentialSource string

// List of AiGatewayProviderCredentialSource.
const (
	AIGATEWAYPROVIDERCREDENTIALSOURCE_INTEGRATION_ACCOUNT AiGatewayProviderCredentialSource = "integration_account"
	AIGATEWAYPROVIDERCREDENTIALSOURCE_SUBSCRIPTION        AiGatewayProviderCredentialSource = "subscription"
)

var allowedAiGatewayProviderCredentialSourceEnumValues = []AiGatewayProviderCredentialSource{
	AIGATEWAYPROVIDERCREDENTIALSOURCE_INTEGRATION_ACCOUNT,
	AIGATEWAYPROVIDERCREDENTIALSOURCE_SUBSCRIPTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayProviderCredentialSource) GetAllowedValues() []AiGatewayProviderCredentialSource {
	return allowedAiGatewayProviderCredentialSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayProviderCredentialSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayProviderCredentialSource(value)
	return nil
}

// NewAiGatewayProviderCredentialSourceFromValue returns a pointer to a valid AiGatewayProviderCredentialSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayProviderCredentialSourceFromValue(v string) (*AiGatewayProviderCredentialSource, error) {
	ev := AiGatewayProviderCredentialSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayProviderCredentialSource: valid values are %v", v, allowedAiGatewayProviderCredentialSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayProviderCredentialSource) IsValid() bool {
	for _, existing := range allowedAiGatewayProviderCredentialSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayProviderCredentialSource value.
func (v AiGatewayProviderCredentialSource) Ptr() *AiGatewayProviderCredentialSource {
	return &v
}
