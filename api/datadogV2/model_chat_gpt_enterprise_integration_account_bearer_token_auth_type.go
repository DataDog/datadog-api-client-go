// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChatGPTEnterpriseIntegrationAccountBearerTokenAuthType The authentication method type.
type ChatGPTEnterpriseIntegrationAccountBearerTokenAuthType string

// List of ChatGPTEnterpriseIntegrationAccountBearerTokenAuthType.
const (
	CHATGPTENTERPRISEINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN ChatGPTEnterpriseIntegrationAccountBearerTokenAuthType = "bearer_token"
)

var allowedChatGPTEnterpriseIntegrationAccountBearerTokenAuthTypeEnumValues = []ChatGPTEnterpriseIntegrationAccountBearerTokenAuthType{
	CHATGPTENTERPRISEINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ChatGPTEnterpriseIntegrationAccountBearerTokenAuthType) GetAllowedValues() []ChatGPTEnterpriseIntegrationAccountBearerTokenAuthType {
	return allowedChatGPTEnterpriseIntegrationAccountBearerTokenAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChatGPTEnterpriseIntegrationAccountBearerTokenAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChatGPTEnterpriseIntegrationAccountBearerTokenAuthType(value)
	return nil
}

// NewChatGPTEnterpriseIntegrationAccountBearerTokenAuthTypeFromValue returns a pointer to a valid ChatGPTEnterpriseIntegrationAccountBearerTokenAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChatGPTEnterpriseIntegrationAccountBearerTokenAuthTypeFromValue(v string) (*ChatGPTEnterpriseIntegrationAccountBearerTokenAuthType, error) {
	ev := ChatGPTEnterpriseIntegrationAccountBearerTokenAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChatGPTEnterpriseIntegrationAccountBearerTokenAuthType: valid values are %v", v, allowedChatGPTEnterpriseIntegrationAccountBearerTokenAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChatGPTEnterpriseIntegrationAccountBearerTokenAuthType) IsValid() bool {
	for _, existing := range allowedChatGPTEnterpriseIntegrationAccountBearerTokenAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChatGPTEnterpriseIntegrationAccountBearerTokenAuthType value.
func (v ChatGPTEnterpriseIntegrationAccountBearerTokenAuthType) Ptr() *ChatGPTEnterpriseIntegrationAccountBearerTokenAuthType {
	return &v
}
