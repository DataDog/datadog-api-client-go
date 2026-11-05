// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureDeleteRequestDataAttributesOperation Proposed operation for preview. Write endpoints infer it from the HTTP method.
type AiGatewayCaptureDeleteRequestDataAttributesOperation string

// List of AiGatewayCaptureDeleteRequestDataAttributesOperation.
const (
	AIGATEWAYCAPTUREDELETEREQUESTDATAATTRIBUTESOPERATION_CREATE  AiGatewayCaptureDeleteRequestDataAttributesOperation = "create"
	AIGATEWAYCAPTUREDELETEREQUESTDATAATTRIBUTESOPERATION_REPLACE AiGatewayCaptureDeleteRequestDataAttributesOperation = "replace"
	AIGATEWAYCAPTUREDELETEREQUESTDATAATTRIBUTESOPERATION_DELETE  AiGatewayCaptureDeleteRequestDataAttributesOperation = "delete"
)

var allowedAiGatewayCaptureDeleteRequestDataAttributesOperationEnumValues = []AiGatewayCaptureDeleteRequestDataAttributesOperation{
	AIGATEWAYCAPTUREDELETEREQUESTDATAATTRIBUTESOPERATION_CREATE,
	AIGATEWAYCAPTUREDELETEREQUESTDATAATTRIBUTESOPERATION_REPLACE,
	AIGATEWAYCAPTUREDELETEREQUESTDATAATTRIBUTESOPERATION_DELETE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayCaptureDeleteRequestDataAttributesOperation) GetAllowedValues() []AiGatewayCaptureDeleteRequestDataAttributesOperation {
	return allowedAiGatewayCaptureDeleteRequestDataAttributesOperationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayCaptureDeleteRequestDataAttributesOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayCaptureDeleteRequestDataAttributesOperation(value)
	return nil
}

// NewAiGatewayCaptureDeleteRequestDataAttributesOperationFromValue returns a pointer to a valid AiGatewayCaptureDeleteRequestDataAttributesOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayCaptureDeleteRequestDataAttributesOperationFromValue(v string) (*AiGatewayCaptureDeleteRequestDataAttributesOperation, error) {
	ev := AiGatewayCaptureDeleteRequestDataAttributesOperation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayCaptureDeleteRequestDataAttributesOperation: valid values are %v", v, allowedAiGatewayCaptureDeleteRequestDataAttributesOperationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayCaptureDeleteRequestDataAttributesOperation) IsValid() bool {
	for _, existing := range allowedAiGatewayCaptureDeleteRequestDataAttributesOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayCaptureDeleteRequestDataAttributesOperation value.
func (v AiGatewayCaptureDeleteRequestDataAttributesOperation) Ptr() *AiGatewayCaptureDeleteRequestDataAttributesOperation {
	return &v
}
