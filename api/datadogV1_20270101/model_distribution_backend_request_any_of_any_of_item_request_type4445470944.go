// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944 Test-only generated schema for the dashboard versioning proof of concept.
type DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944 string

// List of DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944.
const (
	DISTRIBUTIONBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4445470944_HISTOGRAM DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944 = "histogram"
)

var allowedDistributionBackendRequestAnyOfAnyOfItemRequestType4445470944EnumValues = []DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944{
	DISTRIBUTIONBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4445470944_HISTOGRAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944) GetAllowedValues() []DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944 {
	return allowedDistributionBackendRequestAnyOfAnyOfItemRequestType4445470944EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944(value)
	return nil
}

// NewDistributionBackendRequestAnyOfAnyOfItemRequestType4445470944FromValue returns a pointer to a valid DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDistributionBackendRequestAnyOfAnyOfItemRequestType4445470944FromValue(v string) (*DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944, error) {
	ev := DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944: valid values are %v", v, allowedDistributionBackendRequestAnyOfAnyOfItemRequestType4445470944EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944) IsValid() bool {
	for _, existing := range allowedDistributionBackendRequestAnyOfAnyOfItemRequestType4445470944EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944 value.
func (v DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944) Ptr() *DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944 {
	return &v
}
