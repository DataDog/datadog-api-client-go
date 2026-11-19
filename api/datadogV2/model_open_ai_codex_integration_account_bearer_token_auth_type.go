// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAICodexIntegrationAccountBearerTokenAuthType The authentication method type.
type OpenAICodexIntegrationAccountBearerTokenAuthType string

// List of OpenAICodexIntegrationAccountBearerTokenAuthType.
const (
	OPENAICODEXINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN OpenAICodexIntegrationAccountBearerTokenAuthType = "bearer_token"
)

var allowedOpenAICodexIntegrationAccountBearerTokenAuthTypeEnumValues = []OpenAICodexIntegrationAccountBearerTokenAuthType{
	OPENAICODEXINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OpenAICodexIntegrationAccountBearerTokenAuthType) GetAllowedValues() []OpenAICodexIntegrationAccountBearerTokenAuthType {
	return allowedOpenAICodexIntegrationAccountBearerTokenAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OpenAICodexIntegrationAccountBearerTokenAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OpenAICodexIntegrationAccountBearerTokenAuthType(value)
	return nil
}

// NewOpenAICodexIntegrationAccountBearerTokenAuthTypeFromValue returns a pointer to a valid OpenAICodexIntegrationAccountBearerTokenAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOpenAICodexIntegrationAccountBearerTokenAuthTypeFromValue(v string) (*OpenAICodexIntegrationAccountBearerTokenAuthType, error) {
	ev := OpenAICodexIntegrationAccountBearerTokenAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OpenAICodexIntegrationAccountBearerTokenAuthType: valid values are %v", v, allowedOpenAICodexIntegrationAccountBearerTokenAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OpenAICodexIntegrationAccountBearerTokenAuthType) IsValid() bool {
	for _, existing := range allowedOpenAICodexIntegrationAccountBearerTokenAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OpenAICodexIntegrationAccountBearerTokenAuthType value.
func (v OpenAICodexIntegrationAccountBearerTokenAuthType) Ptr() *OpenAICodexIntegrationAccountBearerTokenAuthType {
	return &v
}
