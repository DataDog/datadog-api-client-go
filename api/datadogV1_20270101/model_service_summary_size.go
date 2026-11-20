// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceSummarySize Test-only generated schema for the dashboard versioning proof of concept.
type ServiceSummarySize string

// List of ServiceSummarySize.
const (
	SERVICESUMMARYSIZE_SMALL  ServiceSummarySize = "small"
	SERVICESUMMARYSIZE_LARGE  ServiceSummarySize = "large"
	SERVICESUMMARYSIZE_MEDIUM ServiceSummarySize = "medium"
)

var allowedServiceSummarySizeEnumValues = []ServiceSummarySize{
	SERVICESUMMARYSIZE_SMALL,
	SERVICESUMMARYSIZE_LARGE,
	SERVICESUMMARYSIZE_MEDIUM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceSummarySize) GetAllowedValues() []ServiceSummarySize {
	return allowedServiceSummarySizeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceSummarySize) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceSummarySize(value)
	return nil
}

// NewServiceSummarySizeFromValue returns a pointer to a valid ServiceSummarySize
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceSummarySizeFromValue(v string) (*ServiceSummarySize, error) {
	ev := ServiceSummarySize(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceSummarySize: valid values are %v", v, allowedServiceSummarySizeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceSummarySize) IsValid() bool {
	for _, existing := range allowedServiceSummarySizeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceSummarySize value.
func (v ServiceSummarySize) Ptr() *ServiceSummarySize {
	return &v
}
