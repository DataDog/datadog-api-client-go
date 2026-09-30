// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityContextRevisionsMode Which revisions to return for each entity: `latest` returns only the latest revision of each entity as of `to`,
// and `all` returns every revision in the requested time range.
type EntityContextRevisionsMode string

// List of EntityContextRevisionsMode.
const (
	ENTITYCONTEXTREVISIONSMODE_LATEST EntityContextRevisionsMode = "latest"
	ENTITYCONTEXTREVISIONSMODE_ALL    EntityContextRevisionsMode = "all"
)

var allowedEntityContextRevisionsModeEnumValues = []EntityContextRevisionsMode{
	ENTITYCONTEXTREVISIONSMODE_LATEST,
	ENTITYCONTEXTREVISIONSMODE_ALL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntityContextRevisionsMode) GetAllowedValues() []EntityContextRevisionsMode {
	return allowedEntityContextRevisionsModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityContextRevisionsMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityContextRevisionsMode(value)
	return nil
}

// NewEntityContextRevisionsModeFromValue returns a pointer to a valid EntityContextRevisionsMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityContextRevisionsModeFromValue(v string) (*EntityContextRevisionsMode, error) {
	ev := EntityContextRevisionsMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityContextRevisionsMode: valid values are %v", v, allowedEntityContextRevisionsModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityContextRevisionsMode) IsValid() bool {
	for _, existing := range allowedEntityContextRevisionsModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityContextRevisionsMode value.
func (v EntityContextRevisionsMode) Ptr() *EntityContextRevisionsMode {
	return &v
}
