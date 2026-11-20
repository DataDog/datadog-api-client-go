// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceSummaryLayout Test-only generated schema for the dashboard versioning proof of concept.
type ServiceSummaryLayout string

// List of ServiceSummaryLayout.
const (
	SERVICESUMMARYLAYOUT_ONE_COLUMN   ServiceSummaryLayout = "one_column"
	SERVICESUMMARYLAYOUT_TWO_COLUMN   ServiceSummaryLayout = "two_column"
	SERVICESUMMARYLAYOUT_THREE_COLUMN ServiceSummaryLayout = "three_column"
)

var allowedServiceSummaryLayoutEnumValues = []ServiceSummaryLayout{
	SERVICESUMMARYLAYOUT_ONE_COLUMN,
	SERVICESUMMARYLAYOUT_TWO_COLUMN,
	SERVICESUMMARYLAYOUT_THREE_COLUMN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceSummaryLayout) GetAllowedValues() []ServiceSummaryLayout {
	return allowedServiceSummaryLayoutEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceSummaryLayout) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceSummaryLayout(value)
	return nil
}

// NewServiceSummaryLayoutFromValue returns a pointer to a valid ServiceSummaryLayout
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceSummaryLayoutFromValue(v string) (*ServiceSummaryLayout, error) {
	ev := ServiceSummaryLayout(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceSummaryLayout: valid values are %v", v, allowedServiceSummaryLayoutEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceSummaryLayout) IsValid() bool {
	for _, existing := range allowedServiceSummaryLayoutEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceSummaryLayout value.
func (v ServiceSummaryLayout) Ptr() *ServiceSummaryLayout {
	return &v
}
