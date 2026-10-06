// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioIntegrationAccountBasicAuthType The authentication method type.
type TwilioIntegrationAccountBasicAuthType string

// List of TwilioIntegrationAccountBasicAuthType.
const (
	TWILIOINTEGRATIONACCOUNTBASICAUTHTYPE_BASIC TwilioIntegrationAccountBasicAuthType = "basic"
)

var allowedTwilioIntegrationAccountBasicAuthTypeEnumValues = []TwilioIntegrationAccountBasicAuthType{
	TWILIOINTEGRATIONACCOUNTBASICAUTHTYPE_BASIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TwilioIntegrationAccountBasicAuthType) GetAllowedValues() []TwilioIntegrationAccountBasicAuthType {
	return allowedTwilioIntegrationAccountBasicAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TwilioIntegrationAccountBasicAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TwilioIntegrationAccountBasicAuthType(value)
	return nil
}

// NewTwilioIntegrationAccountBasicAuthTypeFromValue returns a pointer to a valid TwilioIntegrationAccountBasicAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTwilioIntegrationAccountBasicAuthTypeFromValue(v string) (*TwilioIntegrationAccountBasicAuthType, error) {
	ev := TwilioIntegrationAccountBasicAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TwilioIntegrationAccountBasicAuthType: valid values are %v", v, allowedTwilioIntegrationAccountBasicAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TwilioIntegrationAccountBasicAuthType) IsValid() bool {
	for _, existing := range allowedTwilioIntegrationAccountBasicAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TwilioIntegrationAccountBasicAuthType value.
func (v TwilioIntegrationAccountBasicAuthType) Ptr() *TwilioIntegrationAccountBasicAuthType {
	return &v
}
