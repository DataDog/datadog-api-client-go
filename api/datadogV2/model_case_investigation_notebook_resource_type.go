// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseInvestigationNotebookResourceType Case investigation notebook resource type.
type CaseInvestigationNotebookResourceType string

// List of CaseInvestigationNotebookResourceType.
const (
	CASEINVESTIGATIONNOTEBOOKRESOURCETYPE_NOTEBOOK CaseInvestigationNotebookResourceType = "notebook"
)

var allowedCaseInvestigationNotebookResourceTypeEnumValues = []CaseInvestigationNotebookResourceType{
	CASEINVESTIGATIONNOTEBOOKRESOURCETYPE_NOTEBOOK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseInvestigationNotebookResourceType) GetAllowedValues() []CaseInvestigationNotebookResourceType {
	return allowedCaseInvestigationNotebookResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseInvestigationNotebookResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseInvestigationNotebookResourceType(value)
	return nil
}

// NewCaseInvestigationNotebookResourceTypeFromValue returns a pointer to a valid CaseInvestigationNotebookResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseInvestigationNotebookResourceTypeFromValue(v string) (*CaseInvestigationNotebookResourceType, error) {
	ev := CaseInvestigationNotebookResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseInvestigationNotebookResourceType: valid values are %v", v, allowedCaseInvestigationNotebookResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseInvestigationNotebookResourceType) IsValid() bool {
	for _, existing := range allowedCaseInvestigationNotebookResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseInvestigationNotebookResourceType value.
func (v CaseInvestigationNotebookResourceType) Ptr() *CaseInvestigationNotebookResourceType {
	return &v
}
