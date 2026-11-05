// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelSource How a model catalog entry is managed.
type AiGatewayModelSource string

// List of AiGatewayModelSource.
const (
	AIGATEWAYMODELSOURCE_DISCOVERED AiGatewayModelSource = "discovered"
	AIGATEWAYMODELSOURCE_MANUAL     AiGatewayModelSource = "manual"
)

var allowedAiGatewayModelSourceEnumValues = []AiGatewayModelSource{
	AIGATEWAYMODELSOURCE_DISCOVERED,
	AIGATEWAYMODELSOURCE_MANUAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayModelSource) GetAllowedValues() []AiGatewayModelSource {
	return allowedAiGatewayModelSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayModelSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayModelSource(value)
	return nil
}

// NewAiGatewayModelSourceFromValue returns a pointer to a valid AiGatewayModelSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayModelSourceFromValue(v string) (*AiGatewayModelSource, error) {
	ev := AiGatewayModelSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayModelSource: valid values are %v", v, allowedAiGatewayModelSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayModelSource) IsValid() bool {
	for _, existing := range allowedAiGatewayModelSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayModelSource value.
func (v AiGatewayModelSource) Ptr() *AiGatewayModelSource {
	return &v
}
