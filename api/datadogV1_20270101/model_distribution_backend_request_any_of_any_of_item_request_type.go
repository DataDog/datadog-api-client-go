// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionBackendRequestAnyOfAnyOfItemRequestType Test-only generated schema for the dashboard versioning proof of concept.
type DistributionBackendRequestAnyOfAnyOfItemRequestType string

// List of DistributionBackendRequestAnyOfAnyOfItemRequestType.
const (
	DISTRIBUTIONBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE_HISTOGRAM DistributionBackendRequestAnyOfAnyOfItemRequestType = "histogram"
)

var allowedDistributionBackendRequestAnyOfAnyOfItemRequestTypeEnumValues = []DistributionBackendRequestAnyOfAnyOfItemRequestType{
	DISTRIBUTIONBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE_HISTOGRAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DistributionBackendRequestAnyOfAnyOfItemRequestType) GetAllowedValues() []DistributionBackendRequestAnyOfAnyOfItemRequestType {
	return allowedDistributionBackendRequestAnyOfAnyOfItemRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DistributionBackendRequestAnyOfAnyOfItemRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DistributionBackendRequestAnyOfAnyOfItemRequestType(value)
	return nil
}

// NewDistributionBackendRequestAnyOfAnyOfItemRequestTypeFromValue returns a pointer to a valid DistributionBackendRequestAnyOfAnyOfItemRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDistributionBackendRequestAnyOfAnyOfItemRequestTypeFromValue(v string) (*DistributionBackendRequestAnyOfAnyOfItemRequestType, error) {
	ev := DistributionBackendRequestAnyOfAnyOfItemRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DistributionBackendRequestAnyOfAnyOfItemRequestType: valid values are %v", v, allowedDistributionBackendRequestAnyOfAnyOfItemRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DistributionBackendRequestAnyOfAnyOfItemRequestType) IsValid() bool {
	for _, existing := range allowedDistributionBackendRequestAnyOfAnyOfItemRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DistributionBackendRequestAnyOfAnyOfItemRequestType value.
func (v DistributionBackendRequestAnyOfAnyOfItemRequestType) Ptr() *DistributionBackendRequestAnyOfAnyOfItemRequestType {
	return &v
}
