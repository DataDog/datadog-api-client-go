// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayProviderCreateCredentialSource Whose credential the gateway sends to the provider. Defaults to `integration_account`, which sends
// the integration account's credential. `subscription` forwards each caller's own subscription sign-in
// (`ChatGPT` for `openai`, `claude.ai` for `anthropic`) and needs no integration account; it is available
// only to organizations with subscription providers enabled. A subscription provider's models are
// optional: callers can use any model their subscription includes, and a model you add sets its
// pricing. Immutable.
type AiGatewayProviderCreateCredentialSource string

// List of AiGatewayProviderCreateCredentialSource.
const (
	AIGATEWAYPROVIDERCREATECREDENTIALSOURCE_INTEGRATION_ACCOUNT AiGatewayProviderCreateCredentialSource = "integration_account"
	AIGATEWAYPROVIDERCREATECREDENTIALSOURCE_SUBSCRIPTION        AiGatewayProviderCreateCredentialSource = "subscription"
)

var allowedAiGatewayProviderCreateCredentialSourceEnumValues = []AiGatewayProviderCreateCredentialSource{
	AIGATEWAYPROVIDERCREATECREDENTIALSOURCE_INTEGRATION_ACCOUNT,
	AIGATEWAYPROVIDERCREATECREDENTIALSOURCE_SUBSCRIPTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayProviderCreateCredentialSource) GetAllowedValues() []AiGatewayProviderCreateCredentialSource {
	return allowedAiGatewayProviderCreateCredentialSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayProviderCreateCredentialSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayProviderCreateCredentialSource(value)
	return nil
}

// NewAiGatewayProviderCreateCredentialSourceFromValue returns a pointer to a valid AiGatewayProviderCreateCredentialSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayProviderCreateCredentialSourceFromValue(v string) (*AiGatewayProviderCreateCredentialSource, error) {
	ev := AiGatewayProviderCreateCredentialSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayProviderCreateCredentialSource: valid values are %v", v, allowedAiGatewayProviderCreateCredentialSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayProviderCreateCredentialSource) IsValid() bool {
	for _, existing := range allowedAiGatewayProviderCreateCredentialSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayProviderCreateCredentialSource value.
func (v AiGatewayProviderCreateCredentialSource) Ptr() *AiGatewayProviderCreateCredentialSource {
	return &v
}
