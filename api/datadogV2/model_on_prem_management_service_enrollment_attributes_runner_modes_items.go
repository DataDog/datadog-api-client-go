// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnPremManagementServiceEnrollmentAttributesRunnerModesItems
type OnPremManagementServiceEnrollmentAttributesRunnerModesItems string

// List of OnPremManagementServiceEnrollmentAttributesRunnerModesItems.
const (
	ONPREMMANAGEMENTSERVICEENROLLMENTATTRIBUTESRUNNERMODESITEMS_WORKFLOW_AUTOMATION OnPremManagementServiceEnrollmentAttributesRunnerModesItems = "workflow_automation"
	ONPREMMANAGEMENTSERVICEENROLLMENTATTRIBUTESRUNNERMODESITEMS_APP_BUILDER         OnPremManagementServiceEnrollmentAttributesRunnerModesItems = "app_builder"
)

var allowedOnPremManagementServiceEnrollmentAttributesRunnerModesItemsEnumValues = []OnPremManagementServiceEnrollmentAttributesRunnerModesItems{
	ONPREMMANAGEMENTSERVICEENROLLMENTATTRIBUTESRUNNERMODESITEMS_WORKFLOW_AUTOMATION,
	ONPREMMANAGEMENTSERVICEENROLLMENTATTRIBUTESRUNNERMODESITEMS_APP_BUILDER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnPremManagementServiceEnrollmentAttributesRunnerModesItems) GetAllowedValues() []OnPremManagementServiceEnrollmentAttributesRunnerModesItems {
	return allowedOnPremManagementServiceEnrollmentAttributesRunnerModesItemsEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnPremManagementServiceEnrollmentAttributesRunnerModesItems) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnPremManagementServiceEnrollmentAttributesRunnerModesItems(value)
	return nil
}

// NewOnPremManagementServiceEnrollmentAttributesRunnerModesItemsFromValue returns a pointer to a valid OnPremManagementServiceEnrollmentAttributesRunnerModesItems
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnPremManagementServiceEnrollmentAttributesRunnerModesItemsFromValue(v string) (*OnPremManagementServiceEnrollmentAttributesRunnerModesItems, error) {
	ev := OnPremManagementServiceEnrollmentAttributesRunnerModesItems(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnPremManagementServiceEnrollmentAttributesRunnerModesItems: valid values are %v", v, allowedOnPremManagementServiceEnrollmentAttributesRunnerModesItemsEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnPremManagementServiceEnrollmentAttributesRunnerModesItems) IsValid() bool {
	for _, existing := range allowedOnPremManagementServiceEnrollmentAttributesRunnerModesItemsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnPremManagementServiceEnrollmentAttributesRunnerModesItems value.
func (v OnPremManagementServiceEnrollmentAttributesRunnerModesItems) Ptr() *OnPremManagementServiceEnrollmentAttributesRunnerModesItems {
	return &v
}
