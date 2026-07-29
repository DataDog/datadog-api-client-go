// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationAccountInterfaceId Supported interface (source-type) ids (the `interface_id` path scope).
type IntegrationAccountInterfaceId string

// List of IntegrationAccountInterfaceId.
const (
	INTEGRATIONACCOUNTINTERFACEID_ELASTIC_CLOUD     IntegrationAccountInterfaceId = "elastic-cloud"
	INTEGRATIONACCOUNTINTERFACEID_ELASTIC_CLOUD_CCM IntegrationAccountInterfaceId = "elastic-cloud-ccm"
	INTEGRATIONACCOUNTINTERFACEID_TWILIO            IntegrationAccountInterfaceId = "twilio"
)

var allowedIntegrationAccountInterfaceIdEnumValues = []IntegrationAccountInterfaceId{
	INTEGRATIONACCOUNTINTERFACEID_ELASTIC_CLOUD,
	INTEGRATIONACCOUNTINTERFACEID_ELASTIC_CLOUD_CCM,
	INTEGRATIONACCOUNTINTERFACEID_TWILIO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IntegrationAccountInterfaceId) GetAllowedValues() []IntegrationAccountInterfaceId {
	return allowedIntegrationAccountInterfaceIdEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IntegrationAccountInterfaceId) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IntegrationAccountInterfaceId(value)
	return nil
}

// NewIntegrationAccountInterfaceIdFromValue returns a pointer to a valid IntegrationAccountInterfaceId
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIntegrationAccountInterfaceIdFromValue(v string) (*IntegrationAccountInterfaceId, error) {
	ev := IntegrationAccountInterfaceId(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IntegrationAccountInterfaceId: valid values are %v", v, allowedIntegrationAccountInterfaceIdEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IntegrationAccountInterfaceId) IsValid() bool {
	for _, existing := range allowedIntegrationAccountInterfaceIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntegrationAccountInterfaceId value.
func (v IntegrationAccountInterfaceId) Ptr() *IntegrationAccountInterfaceId {
	return &v
}
