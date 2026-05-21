// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentSearchIncidentsSortOrder The ways searched incidents can be sorted.
type IncidentSearchIncidentsSortOrder string

// List of IncidentSearchIncidentsSortOrder.
const (
	INCIDENTSEARCHINCIDENTSSORTORDER_CREATED_ASCENDING   IncidentSearchIncidentsSortOrder = "created"
	INCIDENTSEARCHINCIDENTSSORTORDER_CREATED_DESCENDING  IncidentSearchIncidentsSortOrder = "-created"
	INCIDENTSEARCHINCIDENTSSORTORDER_MODIFIED_ASCENDING  IncidentSearchIncidentsSortOrder = "modified"
	INCIDENTSEARCHINCIDENTSSORTORDER_MODIFIED_DESCENDING IncidentSearchIncidentsSortOrder = "-modified"
)

var allowedIncidentSearchIncidentsSortOrderEnumValues = []IncidentSearchIncidentsSortOrder{
	INCIDENTSEARCHINCIDENTSSORTORDER_CREATED_ASCENDING,
	INCIDENTSEARCHINCIDENTSSORTORDER_CREATED_DESCENDING,
	INCIDENTSEARCHINCIDENTSSORTORDER_MODIFIED_ASCENDING,
	INCIDENTSEARCHINCIDENTSSORTORDER_MODIFIED_DESCENDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentSearchIncidentsSortOrder) GetAllowedValues() []IncidentSearchIncidentsSortOrder {
	return allowedIncidentSearchIncidentsSortOrderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentSearchIncidentsSortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentSearchIncidentsSortOrder(value)
	return nil
}

// NewIncidentSearchIncidentsSortOrderFromValue returns a pointer to a valid IncidentSearchIncidentsSortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentSearchIncidentsSortOrderFromValue(v string) (*IncidentSearchIncidentsSortOrder, error) {
	ev := IncidentSearchIncidentsSortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentSearchIncidentsSortOrder: valid values are %v", v, allowedIncidentSearchIncidentsSortOrderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentSearchIncidentsSortOrder) IsValid() bool {
	for _, existing := range allowedIncidentSearchIncidentsSortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentSearchIncidentsSortOrder value.
func (v IncidentSearchIncidentsSortOrder) Ptr() *IncidentSearchIncidentsSortOrder {
	return &v
}
