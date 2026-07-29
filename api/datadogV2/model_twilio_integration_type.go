// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioIntegrationType Integration discriminator for Twilio.
type TwilioIntegrationType string

// List of TwilioIntegrationType.
const (
	TWILIOINTEGRATIONTYPE_TWILIO TwilioIntegrationType = "twilio"
)

var allowedTwilioIntegrationTypeEnumValues = []TwilioIntegrationType{
	TWILIOINTEGRATIONTYPE_TWILIO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TwilioIntegrationType) GetAllowedValues() []TwilioIntegrationType {
	return allowedTwilioIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TwilioIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TwilioIntegrationType(value)
	return nil
}

// NewTwilioIntegrationTypeFromValue returns a pointer to a valid TwilioIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTwilioIntegrationTypeFromValue(v string) (*TwilioIntegrationType, error) {
	ev := TwilioIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TwilioIntegrationType: valid values are %v", v, allowedTwilioIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TwilioIntegrationType) IsValid() bool {
	for _, existing := range allowedTwilioIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TwilioIntegrationType value.
func (v TwilioIntegrationType) Ptr() *TwilioIntegrationType {
	return &v
}
