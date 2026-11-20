// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityRuntimeStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type SecurityRuntimeStreamQueryDataSource string

// List of SecurityRuntimeStreamQueryDataSource.
const (
	SECURITYRUNTIMESTREAMQUERYDATASOURCE_SECURITY_RUNTIME_STREAM SecurityRuntimeStreamQueryDataSource = "security_runtime_stream"
)

var allowedSecurityRuntimeStreamQueryDataSourceEnumValues = []SecurityRuntimeStreamQueryDataSource{
	SECURITYRUNTIMESTREAMQUERYDATASOURCE_SECURITY_RUNTIME_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityRuntimeStreamQueryDataSource) GetAllowedValues() []SecurityRuntimeStreamQueryDataSource {
	return allowedSecurityRuntimeStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityRuntimeStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityRuntimeStreamQueryDataSource(value)
	return nil
}

// NewSecurityRuntimeStreamQueryDataSourceFromValue returns a pointer to a valid SecurityRuntimeStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityRuntimeStreamQueryDataSourceFromValue(v string) (*SecurityRuntimeStreamQueryDataSource, error) {
	ev := SecurityRuntimeStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityRuntimeStreamQueryDataSource: valid values are %v", v, allowedSecurityRuntimeStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityRuntimeStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedSecurityRuntimeStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityRuntimeStreamQueryDataSource value.
func (v SecurityRuntimeStreamQueryDataSource) Ptr() *SecurityRuntimeStreamQueryDataSource {
	return &v
}
