// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ResourceFilterProviderEnum The name of the cloud provider.
type ResourceFilterProviderEnum string

// List of ResourceFilterProviderEnum.
const (
	RESOURCEFILTERPROVIDERENUM_AWS   ResourceFilterProviderEnum = "aws"
	RESOURCEFILTERPROVIDERENUM_GCP   ResourceFilterProviderEnum = "gcp"
	RESOURCEFILTERPROVIDERENUM_AZURE ResourceFilterProviderEnum = "azure"
)

var allowedResourceFilterProviderEnumEnumValues = []ResourceFilterProviderEnum{
	RESOURCEFILTERPROVIDERENUM_AWS,
	RESOURCEFILTERPROVIDERENUM_GCP,
	RESOURCEFILTERPROVIDERENUM_AZURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ResourceFilterProviderEnum) GetAllowedValues() []ResourceFilterProviderEnum {
	return allowedResourceFilterProviderEnumEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ResourceFilterProviderEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ResourceFilterProviderEnum(value)
	return nil
}

// NewResourceFilterProviderEnumFromValue returns a pointer to a valid ResourceFilterProviderEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewResourceFilterProviderEnumFromValue(v string) (*ResourceFilterProviderEnum, error) {
	ev := ResourceFilterProviderEnum(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ResourceFilterProviderEnum: valid values are %v", v, allowedResourceFilterProviderEnumEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ResourceFilterProviderEnum) IsValid() bool {
	for _, existing := range allowedResourceFilterProviderEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResourceFilterProviderEnum value.
func (v ResourceFilterProviderEnum) Ptr() *ResourceFilterProviderEnum {
	return &v
}
