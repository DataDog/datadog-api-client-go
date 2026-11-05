// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayZDRMode Zero data retention requirement for the organization's requests.
type AiGatewayZDRMode string

// List of AiGatewayZDRMode.
const (
	AIGATEWAYZDRMODE_OFF      AiGatewayZDRMode = "off"
	AIGATEWAYZDRMODE_REQUIRED AiGatewayZDRMode = "required"
)

var allowedAiGatewayZDRModeEnumValues = []AiGatewayZDRMode{
	AIGATEWAYZDRMODE_OFF,
	AIGATEWAYZDRMODE_REQUIRED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayZDRMode) GetAllowedValues() []AiGatewayZDRMode {
	return allowedAiGatewayZDRModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayZDRMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayZDRMode(value)
	return nil
}

// NewAiGatewayZDRModeFromValue returns a pointer to a valid AiGatewayZDRMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayZDRModeFromValue(v string) (*AiGatewayZDRMode, error) {
	ev := AiGatewayZDRMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayZDRMode: valid values are %v", v, allowedAiGatewayZDRModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayZDRMode) IsValid() bool {
	for _, existing := range allowedAiGatewayZDRModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayZDRMode value.
func (v AiGatewayZDRMode) Ptr() *AiGatewayZDRMode {
	return &v
}
