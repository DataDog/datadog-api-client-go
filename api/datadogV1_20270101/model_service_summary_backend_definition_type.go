// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceSummaryBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type ServiceSummaryBackendDefinitionType string

// List of ServiceSummaryBackendDefinitionType.
const (
	SERVICESUMMARYBACKENDDEFINITIONTYPE_TRACE_SERVICE ServiceSummaryBackendDefinitionType = "trace_service"
)

var allowedServiceSummaryBackendDefinitionTypeEnumValues = []ServiceSummaryBackendDefinitionType{
	SERVICESUMMARYBACKENDDEFINITIONTYPE_TRACE_SERVICE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceSummaryBackendDefinitionType) GetAllowedValues() []ServiceSummaryBackendDefinitionType {
	return allowedServiceSummaryBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceSummaryBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceSummaryBackendDefinitionType(value)
	return nil
}

// NewServiceSummaryBackendDefinitionTypeFromValue returns a pointer to a valid ServiceSummaryBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceSummaryBackendDefinitionTypeFromValue(v string) (*ServiceSummaryBackendDefinitionType, error) {
	ev := ServiceSummaryBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceSummaryBackendDefinitionType: valid values are %v", v, allowedServiceSummaryBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceSummaryBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedServiceSummaryBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceSummaryBackendDefinitionType value.
func (v ServiceSummaryBackendDefinitionType) Ptr() *ServiceSummaryBackendDefinitionType {
	return &v
}
