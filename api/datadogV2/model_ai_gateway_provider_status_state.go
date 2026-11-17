// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayProviderStatusState Whether the provider is working. `checking` means no observation has been recorded yet, not that a
// check is running.
type AiGatewayProviderStatusState string

// List of AiGatewayProviderStatusState.
const (
	AIGATEWAYPROVIDERSTATUSSTATE_OK       AiGatewayProviderStatusState = "ok"
	AIGATEWAYPROVIDERSTATUSSTATE_INVALID  AiGatewayProviderStatusState = "invalid"
	AIGATEWAYPROVIDERSTATUSSTATE_CHECKING AiGatewayProviderStatusState = "checking"
)

var allowedAiGatewayProviderStatusStateEnumValues = []AiGatewayProviderStatusState{
	AIGATEWAYPROVIDERSTATUSSTATE_OK,
	AIGATEWAYPROVIDERSTATUSSTATE_INVALID,
	AIGATEWAYPROVIDERSTATUSSTATE_CHECKING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayProviderStatusState) GetAllowedValues() []AiGatewayProviderStatusState {
	return allowedAiGatewayProviderStatusStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayProviderStatusState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayProviderStatusState(value)
	return nil
}

// NewAiGatewayProviderStatusStateFromValue returns a pointer to a valid AiGatewayProviderStatusState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayProviderStatusStateFromValue(v string) (*AiGatewayProviderStatusState, error) {
	ev := AiGatewayProviderStatusState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayProviderStatusState: valid values are %v", v, allowedAiGatewayProviderStatusStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayProviderStatusState) IsValid() bool {
	for _, existing := range allowedAiGatewayProviderStatusStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayProviderStatusState value.
func (v AiGatewayProviderStatusState) Ptr() *AiGatewayProviderStatusState {
	return &v
}
