// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionBackendRequestAnyOfAnyOfItemRequestType4445472224 Test-only generated schema for the dashboard versioning proof of concept.
type DistributionBackendRequestAnyOfAnyOfItemRequestType4445472224 string

// List of DistributionBackendRequestAnyOfAnyOfItemRequestType4445472224.
const (
	DISTRIBUTIONBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4445472224_HISTOGRAM DistributionBackendRequestAnyOfAnyOfItemRequestType4445472224 = "histogram"
)

var allowedDistributionBackendRequestAnyOfAnyOfItemRequestType4445472224EnumValues = []DistributionBackendRequestAnyOfAnyOfItemRequestType4445472224{
	DISTRIBUTIONBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4445472224_HISTOGRAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DistributionBackendRequestAnyOfAnyOfItemRequestType4445472224) GetAllowedValues() []DistributionBackendRequestAnyOfAnyOfItemRequestType4445472224 {
	return allowedDistributionBackendRequestAnyOfAnyOfItemRequestType4445472224EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DistributionBackendRequestAnyOfAnyOfItemRequestType4445472224) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DistributionBackendRequestAnyOfAnyOfItemRequestType4445472224(value)
	return nil
}

// NewDistributionBackendRequestAnyOfAnyOfItemRequestType4445472224FromValue returns a pointer to a valid DistributionBackendRequestAnyOfAnyOfItemRequestType4445472224
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDistributionBackendRequestAnyOfAnyOfItemRequestType4445472224FromValue(v string) (*DistributionBackendRequestAnyOfAnyOfItemRequestType4445472224, error) {
	ev := DistributionBackendRequestAnyOfAnyOfItemRequestType4445472224(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DistributionBackendRequestAnyOfAnyOfItemRequestType4445472224: valid values are %v", v, allowedDistributionBackendRequestAnyOfAnyOfItemRequestType4445472224EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DistributionBackendRequestAnyOfAnyOfItemRequestType4445472224) IsValid() bool {
	for _, existing := range allowedDistributionBackendRequestAnyOfAnyOfItemRequestType4445472224EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DistributionBackendRequestAnyOfAnyOfItemRequestType4445472224 value.
func (v DistributionBackendRequestAnyOfAnyOfItemRequestType4445472224) Ptr() *DistributionBackendRequestAnyOfAnyOfItemRequestType4445472224 {
	return &v
}
