// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelMetadataLifecycle Model lifecycle state.
type AiGatewayModelMetadataLifecycle string

// List of AiGatewayModelMetadataLifecycle.
const (
	AIGATEWAYMODELMETADATALIFECYCLE_ACTIVE     AiGatewayModelMetadataLifecycle = "active"
	AIGATEWAYMODELMETADATALIFECYCLE_ALPHA      AiGatewayModelMetadataLifecycle = "alpha"
	AIGATEWAYMODELMETADATALIFECYCLE_BETA       AiGatewayModelMetadataLifecycle = "beta"
	AIGATEWAYMODELMETADATALIFECYCLE_DEPRECATED AiGatewayModelMetadataLifecycle = "deprecated"
)

var allowedAiGatewayModelMetadataLifecycleEnumValues = []AiGatewayModelMetadataLifecycle{
	AIGATEWAYMODELMETADATALIFECYCLE_ACTIVE,
	AIGATEWAYMODELMETADATALIFECYCLE_ALPHA,
	AIGATEWAYMODELMETADATALIFECYCLE_BETA,
	AIGATEWAYMODELMETADATALIFECYCLE_DEPRECATED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayModelMetadataLifecycle) GetAllowedValues() []AiGatewayModelMetadataLifecycle {
	return allowedAiGatewayModelMetadataLifecycleEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayModelMetadataLifecycle) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayModelMetadataLifecycle(value)
	return nil
}

// NewAiGatewayModelMetadataLifecycleFromValue returns a pointer to a valid AiGatewayModelMetadataLifecycle
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayModelMetadataLifecycleFromValue(v string) (*AiGatewayModelMetadataLifecycle, error) {
	ev := AiGatewayModelMetadataLifecycle(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayModelMetadataLifecycle: valid values are %v", v, allowedAiGatewayModelMetadataLifecycleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayModelMetadataLifecycle) IsValid() bool {
	for _, existing := range allowedAiGatewayModelMetadataLifecycleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayModelMetadataLifecycle value.
func (v AiGatewayModelMetadataLifecycle) Ptr() *AiGatewayModelMetadataLifecycle {
	return &v
}
