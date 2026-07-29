// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioBasicAuthType Authentication method discriminator.
type TwilioBasicAuthType string

// List of TwilioBasicAuthType.
const (
	TWILIOBASICAUTHTYPE_BASIC TwilioBasicAuthType = "basic"
)

var allowedTwilioBasicAuthTypeEnumValues = []TwilioBasicAuthType{
	TWILIOBASICAUTHTYPE_BASIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TwilioBasicAuthType) GetAllowedValues() []TwilioBasicAuthType {
	return allowedTwilioBasicAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TwilioBasicAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TwilioBasicAuthType(value)
	return nil
}

// NewTwilioBasicAuthTypeFromValue returns a pointer to a valid TwilioBasicAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTwilioBasicAuthTypeFromValue(v string) (*TwilioBasicAuthType, error) {
	ev := TwilioBasicAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TwilioBasicAuthType: valid values are %v", v, allowedTwilioBasicAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TwilioBasicAuthType) IsValid() bool {
	for _, existing := range allowedTwilioBasicAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TwilioBasicAuthType value.
func (v TwilioBasicAuthType) Ptr() *TwilioBasicAuthType {
	return &v
}
