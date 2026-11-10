// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueBitsInvestigationType Type of the object.
type IssueBitsInvestigationType string

// List of IssueBitsInvestigationType.
const (
	ISSUEBITSINVESTIGATIONTYPE_BITS_INVESTIGATION IssueBitsInvestigationType = "bits_investigation"
)

var allowedIssueBitsInvestigationTypeEnumValues = []IssueBitsInvestigationType{
	ISSUEBITSINVESTIGATIONTYPE_BITS_INVESTIGATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueBitsInvestigationType) GetAllowedValues() []IssueBitsInvestigationType {
	return allowedIssueBitsInvestigationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueBitsInvestigationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueBitsInvestigationType(value)
	return nil
}

// NewIssueBitsInvestigationTypeFromValue returns a pointer to a valid IssueBitsInvestigationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueBitsInvestigationTypeFromValue(v string) (*IssueBitsInvestigationType, error) {
	ev := IssueBitsInvestigationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueBitsInvestigationType: valid values are %v", v, allowedIssueBitsInvestigationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueBitsInvestigationType) IsValid() bool {
	for _, existing := range allowedIssueBitsInvestigationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueBitsInvestigationType value.
func (v IssueBitsInvestigationType) Ptr() *IssueBitsInvestigationType {
	return &v
}
