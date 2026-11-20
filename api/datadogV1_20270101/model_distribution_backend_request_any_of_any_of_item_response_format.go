// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionBackendRequestAnyOfAnyOfItemResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type DistributionBackendRequestAnyOfAnyOfItemResponseFormat string

// List of DistributionBackendRequestAnyOfAnyOfItemResponseFormat.
const (
	DISTRIBUTIONBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT_SCALAR DistributionBackendRequestAnyOfAnyOfItemResponseFormat = "scalar"
)

var allowedDistributionBackendRequestAnyOfAnyOfItemResponseFormatEnumValues = []DistributionBackendRequestAnyOfAnyOfItemResponseFormat{
	DISTRIBUTIONBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DistributionBackendRequestAnyOfAnyOfItemResponseFormat) GetAllowedValues() []DistributionBackendRequestAnyOfAnyOfItemResponseFormat {
	return allowedDistributionBackendRequestAnyOfAnyOfItemResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DistributionBackendRequestAnyOfAnyOfItemResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DistributionBackendRequestAnyOfAnyOfItemResponseFormat(value)
	return nil
}

// NewDistributionBackendRequestAnyOfAnyOfItemResponseFormatFromValue returns a pointer to a valid DistributionBackendRequestAnyOfAnyOfItemResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDistributionBackendRequestAnyOfAnyOfItemResponseFormatFromValue(v string) (*DistributionBackendRequestAnyOfAnyOfItemResponseFormat, error) {
	ev := DistributionBackendRequestAnyOfAnyOfItemResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DistributionBackendRequestAnyOfAnyOfItemResponseFormat: valid values are %v", v, allowedDistributionBackendRequestAnyOfAnyOfItemResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DistributionBackendRequestAnyOfAnyOfItemResponseFormat) IsValid() bool {
	for _, existing := range allowedDistributionBackendRequestAnyOfAnyOfItemResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DistributionBackendRequestAnyOfAnyOfItemResponseFormat value.
func (v DistributionBackendRequestAnyOfAnyOfItemResponseFormat) Ptr() *DistributionBackendRequestAnyOfAnyOfItemResponseFormat {
	return &v
}
