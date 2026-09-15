// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3RepositoryKind The definition of Entity V3 Repository Kind object.
type EntityV3RepositoryKind string

// List of EntityV3RepositoryKind.
const (
	ENTITYV3REPOSITORYKIND_REPOSITORY EntityV3RepositoryKind = "repository"
)

var allowedEntityV3RepositoryKindEnumValues = []EntityV3RepositoryKind{
	ENTITYV3REPOSITORYKIND_REPOSITORY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntityV3RepositoryKind) GetAllowedValues() []EntityV3RepositoryKind {
	return allowedEntityV3RepositoryKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityV3RepositoryKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityV3RepositoryKind(value)
	return nil
}

// NewEntityV3RepositoryKindFromValue returns a pointer to a valid EntityV3RepositoryKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityV3RepositoryKindFromValue(v string) (*EntityV3RepositoryKind, error) {
	ev := EntityV3RepositoryKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityV3RepositoryKind: valid values are %v", v, allowedEntityV3RepositoryKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityV3RepositoryKind) IsValid() bool {
	for _, existing := range allowedEntityV3RepositoryKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityV3RepositoryKind value.
func (v EntityV3RepositoryKind) Ptr() *EntityV3RepositoryKind {
	return &v
}
