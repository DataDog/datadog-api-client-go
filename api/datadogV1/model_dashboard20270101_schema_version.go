// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Dashboard20270101SchemaVersion Test-only request-body discriminator for the dashboard schema-versioning proof of concept.
type Dashboard20270101SchemaVersion string

// List of Dashboard20270101SchemaVersion.
const (
	DASHBOARD20270101SCHEMAVERSION_VERSION_2027_01_01 Dashboard20270101SchemaVersion = "2027-01-01"
)

var allowedDashboard20270101SchemaVersionEnumValues = []Dashboard20270101SchemaVersion{
	DASHBOARD20270101SCHEMAVERSION_VERSION_2027_01_01,
}

// GetAllowedValues reeturns the list of possible values.
func (v *Dashboard20270101SchemaVersion) GetAllowedValues() []Dashboard20270101SchemaVersion {
	return allowedDashboard20270101SchemaVersionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *Dashboard20270101SchemaVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = Dashboard20270101SchemaVersion(value)
	return nil
}

// NewDashboard20270101SchemaVersionFromValue returns a pointer to a valid Dashboard20270101SchemaVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboard20270101SchemaVersionFromValue(v string) (*Dashboard20270101SchemaVersion, error) {
	ev := Dashboard20270101SchemaVersion(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for Dashboard20270101SchemaVersion: valid values are %v", v, allowedDashboard20270101SchemaVersionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v Dashboard20270101SchemaVersion) IsValid() bool {
	for _, existing := range allowedDashboard20270101SchemaVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Dashboard20270101SchemaVersion value.
func (v Dashboard20270101SchemaVersion) Ptr() *Dashboard20270101SchemaVersion {
	return &v
}
