// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DomainAllowlistType Email domain allowlist allowlist type.
type DomainAllowlistType string

// List of DomainAllowlistType.
const (
	DOMAINALLOWLISTTYPE_DOMAIN_ALLOWLIST DomainAllowlistType = "domain_allowlist"
)

var allowedDomainAllowlistTypeEnumValues = []DomainAllowlistType{
	DOMAINALLOWLISTTYPE_DOMAIN_ALLOWLIST,
}

// GetAllowedValues returns the list of possible values.
func (v *DomainAllowlistType) GetAllowedValues() []DomainAllowlistType {
	return allowedDomainAllowlistTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DomainAllowlistType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DomainAllowlistType(value)
	return nil
}

// NewDomainAllowlistTypeFromValue returns a pointer to a valid DomainAllowlistType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDomainAllowlistTypeFromValue(v string) (*DomainAllowlistType, error) {
	ev := DomainAllowlistType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DomainAllowlistType: valid values are %v", v, allowedDomainAllowlistTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DomainAllowlistType) IsValid() bool {
	for _, existing := range allowedDomainAllowlistTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DomainAllowlistType value.
func (v DomainAllowlistType) Ptr() *DomainAllowlistType {
	return &v
}
