// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType The authentication method type.
type OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType string

// List of OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType.
const (
	OPENAILLMEVALUATIONSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType = "bearer_token"
)

var allowedOpenAILLMEvaluationsIntegrationAccountBearerTokenAuthTypeEnumValues = []OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType{
	OPENAILLMEVALUATIONSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType) GetAllowedValues() []OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType {
	return allowedOpenAILLMEvaluationsIntegrationAccountBearerTokenAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType(value)
	return nil
}

// NewOpenAILLMEvaluationsIntegrationAccountBearerTokenAuthTypeFromValue returns a pointer to a valid OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOpenAILLMEvaluationsIntegrationAccountBearerTokenAuthTypeFromValue(v string) (*OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType, error) {
	ev := OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType: valid values are %v", v, allowedOpenAILLMEvaluationsIntegrationAccountBearerTokenAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType) IsValid() bool {
	for _, existing := range allowedOpenAILLMEvaluationsIntegrationAccountBearerTokenAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType value.
func (v OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType) Ptr() *OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType {
	return &v
}
