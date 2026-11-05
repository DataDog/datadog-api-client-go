// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureUpdateRequestDataAttributesOperation Proposed operation for preview. Write endpoints infer it from the HTTP method.
type AiGatewayCaptureUpdateRequestDataAttributesOperation string

// List of AiGatewayCaptureUpdateRequestDataAttributesOperation.
const (
	AIGATEWAYCAPTUREUPDATEREQUESTDATAATTRIBUTESOPERATION_CREATE  AiGatewayCaptureUpdateRequestDataAttributesOperation = "create"
	AIGATEWAYCAPTUREUPDATEREQUESTDATAATTRIBUTESOPERATION_REPLACE AiGatewayCaptureUpdateRequestDataAttributesOperation = "replace"
	AIGATEWAYCAPTUREUPDATEREQUESTDATAATTRIBUTESOPERATION_DELETE  AiGatewayCaptureUpdateRequestDataAttributesOperation = "delete"
)

var allowedAiGatewayCaptureUpdateRequestDataAttributesOperationEnumValues = []AiGatewayCaptureUpdateRequestDataAttributesOperation{
	AIGATEWAYCAPTUREUPDATEREQUESTDATAATTRIBUTESOPERATION_CREATE,
	AIGATEWAYCAPTUREUPDATEREQUESTDATAATTRIBUTESOPERATION_REPLACE,
	AIGATEWAYCAPTUREUPDATEREQUESTDATAATTRIBUTESOPERATION_DELETE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayCaptureUpdateRequestDataAttributesOperation) GetAllowedValues() []AiGatewayCaptureUpdateRequestDataAttributesOperation {
	return allowedAiGatewayCaptureUpdateRequestDataAttributesOperationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayCaptureUpdateRequestDataAttributesOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayCaptureUpdateRequestDataAttributesOperation(value)
	return nil
}

// NewAiGatewayCaptureUpdateRequestDataAttributesOperationFromValue returns a pointer to a valid AiGatewayCaptureUpdateRequestDataAttributesOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayCaptureUpdateRequestDataAttributesOperationFromValue(v string) (*AiGatewayCaptureUpdateRequestDataAttributesOperation, error) {
	ev := AiGatewayCaptureUpdateRequestDataAttributesOperation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayCaptureUpdateRequestDataAttributesOperation: valid values are %v", v, allowedAiGatewayCaptureUpdateRequestDataAttributesOperationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayCaptureUpdateRequestDataAttributesOperation) IsValid() bool {
	for _, existing := range allowedAiGatewayCaptureUpdateRequestDataAttributesOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayCaptureUpdateRequestDataAttributesOperation value.
func (v AiGatewayCaptureUpdateRequestDataAttributesOperation) Ptr() *AiGatewayCaptureUpdateRequestDataAttributesOperation {
	return &v
}
