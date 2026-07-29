// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationAccountIntegrationId Supported integration ids (the `integration_id` path scope).
type IntegrationAccountIntegrationId string

// List of IntegrationAccountIntegrationId.
const (
	INTEGRATIONACCOUNTINTEGRATIONID_ELASTIC_CLOUD IntegrationAccountIntegrationId = "elastic-cloud"
	INTEGRATIONACCOUNTINTEGRATIONID_TWILIO        IntegrationAccountIntegrationId = "twilio"
)

var allowedIntegrationAccountIntegrationIdEnumValues = []IntegrationAccountIntegrationId{
	INTEGRATIONACCOUNTINTEGRATIONID_ELASTIC_CLOUD,
	INTEGRATIONACCOUNTINTEGRATIONID_TWILIO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IntegrationAccountIntegrationId) GetAllowedValues() []IntegrationAccountIntegrationId {
	return allowedIntegrationAccountIntegrationIdEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IntegrationAccountIntegrationId) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IntegrationAccountIntegrationId(value)
	return nil
}

// NewIntegrationAccountIntegrationIdFromValue returns a pointer to a valid IntegrationAccountIntegrationId
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIntegrationAccountIntegrationIdFromValue(v string) (*IntegrationAccountIntegrationId, error) {
	ev := IntegrationAccountIntegrationId(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IntegrationAccountIntegrationId: valid values are %v", v, allowedIntegrationAccountIntegrationIdEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IntegrationAccountIntegrationId) IsValid() bool {
	for _, existing := range allowedIntegrationAccountIntegrationIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntegrationAccountIntegrationId value.
func (v IntegrationAccountIntegrationId) Ptr() *IntegrationAccountIntegrationId {
	return &v
}
