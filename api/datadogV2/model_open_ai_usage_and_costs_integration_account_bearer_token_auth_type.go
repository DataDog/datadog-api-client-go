// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType The authentication method type.
type OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType string

// List of OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType.
const (
	OPENAIUSAGEANDCOSTSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType = "bearer_token"
)

var allowedOpenAIUsageAndCostsIntegrationAccountBearerTokenAuthTypeEnumValues = []OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType{
	OPENAIUSAGEANDCOSTSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType) GetAllowedValues() []OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType {
	return allowedOpenAIUsageAndCostsIntegrationAccountBearerTokenAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType(value)
	return nil
}

// NewOpenAIUsageAndCostsIntegrationAccountBearerTokenAuthTypeFromValue returns a pointer to a valid OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOpenAIUsageAndCostsIntegrationAccountBearerTokenAuthTypeFromValue(v string) (*OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType, error) {
	ev := OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType: valid values are %v", v, allowedOpenAIUsageAndCostsIntegrationAccountBearerTokenAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType) IsValid() bool {
	for _, existing := range allowedOpenAIUsageAndCostsIntegrationAccountBearerTokenAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType value.
func (v OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType) Ptr() *OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType {
	return &v
}
