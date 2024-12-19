// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Ecosystem The related vulnerability asset ecosystem.
type Ecosystem string

// List of Ecosystem.
const (
	ECOSYSTEM_PYPI      Ecosystem = "PyPI"
	ECOSYSTEM_MAVEN     Ecosystem = "Maven"
	ECOSYSTEM_NUGET     Ecosystem = "NuGet"
	ECOSYSTEM_NPM       Ecosystem = "Npm"
	ECOSYSTEM_RUBYGEMS  Ecosystem = "RubyGems"
	ECOSYSTEM_GO        Ecosystem = "Go"
	ECOSYSTEM_PACKAGIST Ecosystem = "Packagist"
	ECOSYSTEM_DDEB      Ecosystem = "Ddeb"
	ECOSYSTEM_RPM       Ecosystem = "Rpm"
	ECOSYSTEM_APK       Ecosystem = "Apk"
	ECOSYSTEM_WINDOWS   Ecosystem = "Windows"
)

var allowedEcosystemEnumValues = []Ecosystem{
	ECOSYSTEM_PYPI,
	ECOSYSTEM_MAVEN,
	ECOSYSTEM_NUGET,
	ECOSYSTEM_NPM,
	ECOSYSTEM_RUBYGEMS,
	ECOSYSTEM_GO,
	ECOSYSTEM_PACKAGIST,
	ECOSYSTEM_DDEB,
	ECOSYSTEM_RPM,
	ECOSYSTEM_APK,
	ECOSYSTEM_WINDOWS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *Ecosystem) GetAllowedValues() []Ecosystem {
	return allowedEcosystemEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *Ecosystem) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = Ecosystem(value)
	return nil
}

// NewEcosystemFromValue returns a pointer to a valid Ecosystem
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEcosystemFromValue(v string) (*Ecosystem, error) {
	ev := Ecosystem(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for Ecosystem: valid values are %v", v, allowedEcosystemEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v Ecosystem) IsValid() bool {
	for _, existing := range allowedEcosystemEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Ecosystem value.
func (v Ecosystem) Ptr() *Ecosystem {
	return &v
}
