// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFormulaRetentionRequestType The resource type identifier for a retention scalar or retention timeseries request.
type ProductAnalyticsFormulaRetentionRequestType string

// List of ProductAnalyticsFormulaRetentionRequestType.
const (
	PRODUCTANALYTICSFORMULARETENTIONREQUESTTYPE_FORMULA_RETENTION_REQUEST ProductAnalyticsFormulaRetentionRequestType = "formula_retention_request"
)

var allowedProductAnalyticsFormulaRetentionRequestTypeEnumValues = []ProductAnalyticsFormulaRetentionRequestType{
	PRODUCTANALYTICSFORMULARETENTIONREQUESTTYPE_FORMULA_RETENTION_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsFormulaRetentionRequestType) GetAllowedValues() []ProductAnalyticsFormulaRetentionRequestType {
	return allowedProductAnalyticsFormulaRetentionRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsFormulaRetentionRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsFormulaRetentionRequestType(value)
	return nil
}

// NewProductAnalyticsFormulaRetentionRequestTypeFromValue returns a pointer to a valid ProductAnalyticsFormulaRetentionRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsFormulaRetentionRequestTypeFromValue(v string) (*ProductAnalyticsFormulaRetentionRequestType, error) {
	ev := ProductAnalyticsFormulaRetentionRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsFormulaRetentionRequestType: valid values are %v", v, allowedProductAnalyticsFormulaRetentionRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsFormulaRetentionRequestType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsFormulaRetentionRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsFormulaRetentionRequestType value.
func (v ProductAnalyticsFormulaRetentionRequestType) Ptr() *ProductAnalyticsFormulaRetentionRequestType {
	return &v
}
