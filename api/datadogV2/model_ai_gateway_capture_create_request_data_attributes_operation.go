// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureCreateRequestDataAttributesOperation Proposed operation for preview. Write endpoints infer it from the HTTP method.
type AiGatewayCaptureCreateRequestDataAttributesOperation string

// List of AiGatewayCaptureCreateRequestDataAttributesOperation.
const (
	AIGATEWAYCAPTURECREATEREQUESTDATAATTRIBUTESOPERATION_CREATE  AiGatewayCaptureCreateRequestDataAttributesOperation = "create"
	AIGATEWAYCAPTURECREATEREQUESTDATAATTRIBUTESOPERATION_REPLACE AiGatewayCaptureCreateRequestDataAttributesOperation = "replace"
	AIGATEWAYCAPTURECREATEREQUESTDATAATTRIBUTESOPERATION_DELETE  AiGatewayCaptureCreateRequestDataAttributesOperation = "delete"
)

var allowedAiGatewayCaptureCreateRequestDataAttributesOperationEnumValues = []AiGatewayCaptureCreateRequestDataAttributesOperation{
	AIGATEWAYCAPTURECREATEREQUESTDATAATTRIBUTESOPERATION_CREATE,
	AIGATEWAYCAPTURECREATEREQUESTDATAATTRIBUTESOPERATION_REPLACE,
	AIGATEWAYCAPTURECREATEREQUESTDATAATTRIBUTESOPERATION_DELETE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayCaptureCreateRequestDataAttributesOperation) GetAllowedValues() []AiGatewayCaptureCreateRequestDataAttributesOperation {
	return allowedAiGatewayCaptureCreateRequestDataAttributesOperationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayCaptureCreateRequestDataAttributesOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayCaptureCreateRequestDataAttributesOperation(value)
	return nil
}

// NewAiGatewayCaptureCreateRequestDataAttributesOperationFromValue returns a pointer to a valid AiGatewayCaptureCreateRequestDataAttributesOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayCaptureCreateRequestDataAttributesOperationFromValue(v string) (*AiGatewayCaptureCreateRequestDataAttributesOperation, error) {
	ev := AiGatewayCaptureCreateRequestDataAttributesOperation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayCaptureCreateRequestDataAttributesOperation: valid values are %v", v, allowedAiGatewayCaptureCreateRequestDataAttributesOperationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayCaptureCreateRequestDataAttributesOperation) IsValid() bool {
	for _, existing := range allowedAiGatewayCaptureCreateRequestDataAttributesOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayCaptureCreateRequestDataAttributesOperation value.
func (v AiGatewayCaptureCreateRequestDataAttributesOperation) Ptr() *AiGatewayCaptureCreateRequestDataAttributesOperation {
	return &v
}
