// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod Statistical method used to calculate the experiment results.
type ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod string

// List of ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod.
const (
	EXPERIMENTSANALYSISPLANV2DTODATAATTRIBUTESCONFIDENCEINTERVALMETHOD_SEQUENTIAL            ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod = "Sequential"
	EXPERIMENTSANALYSISPLANV2DTODATAATTRIBUTESCONFIDENCEINTERVALMETHOD_FIXEDSAMPLE           ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod = "FixedSample"
	EXPERIMENTSANALYSISPLANV2DTODATAATTRIBUTESCONFIDENCEINTERVALMETHOD_BAYESIAN              ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod = "Bayesian"
	EXPERIMENTSANALYSISPLANV2DTODATAATTRIBUTESCONFIDENCEINTERVALMETHOD_SEQUENTIALFIXEDHYBRID ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod = "SequentialFixedHybrid"
)

var allowedExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethodEnumValues = []ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod{
	EXPERIMENTSANALYSISPLANV2DTODATAATTRIBUTESCONFIDENCEINTERVALMETHOD_SEQUENTIAL,
	EXPERIMENTSANALYSISPLANV2DTODATAATTRIBUTESCONFIDENCEINTERVALMETHOD_FIXEDSAMPLE,
	EXPERIMENTSANALYSISPLANV2DTODATAATTRIBUTESCONFIDENCEINTERVALMETHOD_BAYESIAN,
	EXPERIMENTSANALYSISPLANV2DTODATAATTRIBUTESCONFIDENCEINTERVALMETHOD_SEQUENTIALFIXEDHYBRID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod) GetAllowedValues() []ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod {
	return allowedExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethodEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod(value)
	return nil
}

// NewExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethodFromValue returns a pointer to a valid ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethodFromValue(v string) (*ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod, error) {
	ev := ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod: valid values are %v", v, allowedExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethodEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod) IsValid() bool {
	for _, existing := range allowedExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod value.
func (v ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod) Ptr() *ExperimentsAnalysisPlanV2DTODataAttributesConfidenceIntervalMethod {
	return &v
}
