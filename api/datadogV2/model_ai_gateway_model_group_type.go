// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelGroupType AI Gateway model group resource type.
type AiGatewayModelGroupType string

// List of AiGatewayModelGroupType.
const (
	AIGATEWAYMODELGROUPTYPE_MODEL_GROUPS AiGatewayModelGroupType = "model-groups"
)

var allowedAiGatewayModelGroupTypeEnumValues = []AiGatewayModelGroupType{
	AIGATEWAYMODELGROUPTYPE_MODEL_GROUPS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayModelGroupType) GetAllowedValues() []AiGatewayModelGroupType {
	return allowedAiGatewayModelGroupTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayModelGroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayModelGroupType(value)
	return nil
}

// NewAiGatewayModelGroupTypeFromValue returns a pointer to a valid AiGatewayModelGroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayModelGroupTypeFromValue(v string) (*AiGatewayModelGroupType, error) {
	ev := AiGatewayModelGroupType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayModelGroupType: valid values are %v", v, allowedAiGatewayModelGroupTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayModelGroupType) IsValid() bool {
	for _, existing := range allowedAiGatewayModelGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayModelGroupType value.
func (v AiGatewayModelGroupType) Ptr() *AiGatewayModelGroupType {
	return &v
}
