// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayProviderInitialBindingRelation Provider access relation. Editor also grants View.
type AiGatewayProviderInitialBindingRelation string

// List of AiGatewayProviderInitialBindingRelation.
const (
	AIGATEWAYPROVIDERINITIALBINDINGRELATION_VIEWER AiGatewayProviderInitialBindingRelation = "viewer"
	AIGATEWAYPROVIDERINITIALBINDINGRELATION_EDITOR AiGatewayProviderInitialBindingRelation = "editor"
)

var allowedAiGatewayProviderInitialBindingRelationEnumValues = []AiGatewayProviderInitialBindingRelation{
	AIGATEWAYPROVIDERINITIALBINDINGRELATION_VIEWER,
	AIGATEWAYPROVIDERINITIALBINDINGRELATION_EDITOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayProviderInitialBindingRelation) GetAllowedValues() []AiGatewayProviderInitialBindingRelation {
	return allowedAiGatewayProviderInitialBindingRelationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayProviderInitialBindingRelation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayProviderInitialBindingRelation(value)
	return nil
}

// NewAiGatewayProviderInitialBindingRelationFromValue returns a pointer to a valid AiGatewayProviderInitialBindingRelation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayProviderInitialBindingRelationFromValue(v string) (*AiGatewayProviderInitialBindingRelation, error) {
	ev := AiGatewayProviderInitialBindingRelation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayProviderInitialBindingRelation: valid values are %v", v, allowedAiGatewayProviderInitialBindingRelationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayProviderInitialBindingRelation) IsValid() bool {
	for _, existing := range allowedAiGatewayProviderInitialBindingRelationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayProviderInitialBindingRelation value.
func (v AiGatewayProviderInitialBindingRelation) Ptr() *AiGatewayProviderInitialBindingRelation {
	return &v
}
