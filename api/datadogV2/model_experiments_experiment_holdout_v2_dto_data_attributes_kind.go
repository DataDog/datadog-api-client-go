// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentHoldoutV2DTODataAttributesKind Whether Datadog serves the holdout. Datadog serves MANAGED holdouts. EXTERNAL holdouts record what your own feature flagging system holds back, and Datadog does not serve them.
type ExperimentsExperimentHoldoutV2DTODataAttributesKind string

// List of ExperimentsExperimentHoldoutV2DTODataAttributesKind.
const (
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESKIND_MANAGED  ExperimentsExperimentHoldoutV2DTODataAttributesKind = "MANAGED"
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESKIND_EXTERNAL ExperimentsExperimentHoldoutV2DTODataAttributesKind = "EXTERNAL"
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESKIND_UNKNOWN  ExperimentsExperimentHoldoutV2DTODataAttributesKind = "UNKNOWN"
)

var allowedExperimentsExperimentHoldoutV2DTODataAttributesKindEnumValues = []ExperimentsExperimentHoldoutV2DTODataAttributesKind{
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESKIND_MANAGED,
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESKIND_EXTERNAL,
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESKIND_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsExperimentHoldoutV2DTODataAttributesKind) GetAllowedValues() []ExperimentsExperimentHoldoutV2DTODataAttributesKind {
	return allowedExperimentsExperimentHoldoutV2DTODataAttributesKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsExperimentHoldoutV2DTODataAttributesKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsExperimentHoldoutV2DTODataAttributesKind(value)
	return nil
}

// NewExperimentsExperimentHoldoutV2DTODataAttributesKindFromValue returns a pointer to a valid ExperimentsExperimentHoldoutV2DTODataAttributesKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsExperimentHoldoutV2DTODataAttributesKindFromValue(v string) (*ExperimentsExperimentHoldoutV2DTODataAttributesKind, error) {
	ev := ExperimentsExperimentHoldoutV2DTODataAttributesKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsExperimentHoldoutV2DTODataAttributesKind: valid values are %v", v, allowedExperimentsExperimentHoldoutV2DTODataAttributesKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsExperimentHoldoutV2DTODataAttributesKind) IsValid() bool {
	for _, existing := range allowedExperimentsExperimentHoldoutV2DTODataAttributesKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsExperimentHoldoutV2DTODataAttributesKind value.
func (v ExperimentsExperimentHoldoutV2DTODataAttributesKind) Ptr() *ExperimentsExperimentHoldoutV2DTODataAttributesKind {
	return &v
}
