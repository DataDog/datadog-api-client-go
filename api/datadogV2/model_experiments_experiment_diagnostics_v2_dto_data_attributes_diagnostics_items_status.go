// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus Outcome of an individual diagnostic check.
type ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus string

// List of ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus.
const (
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESDIAGNOSTICSITEMSSTATUS_PASS    ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus = "PASS"
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESDIAGNOSTICSITEMSSTATUS_FAIL    ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus = "FAIL"
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESDIAGNOSTICSITEMSSTATUS_WARN    ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus = "WARN"
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESDIAGNOSTICSITEMSSTATUS_ERROR   ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus = "ERROR"
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESDIAGNOSTICSITEMSSTATUS_SKIPPED ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus = "SKIPPED"
)

var allowedExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatusEnumValues = []ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus{
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESDIAGNOSTICSITEMSSTATUS_PASS,
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESDIAGNOSTICSITEMSSTATUS_FAIL,
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESDIAGNOSTICSITEMSSTATUS_WARN,
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESDIAGNOSTICSITEMSSTATUS_ERROR,
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESDIAGNOSTICSITEMSSTATUS_SKIPPED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus) GetAllowedValues() []ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus {
	return allowedExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus(value)
	return nil
}

// NewExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatusFromValue returns a pointer to a valid ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatusFromValue(v string) (*ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus, error) {
	ev := ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus: valid values are %v", v, allowedExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus) IsValid() bool {
	for _, existing := range allowedExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus value.
func (v ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus) Ptr() *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus {
	return &v
}
