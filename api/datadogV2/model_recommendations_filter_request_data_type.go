// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RecommendationsFilterRequestDataType Legacy JSON:API resource type required by the cost recommendations search decoder.
type RecommendationsFilterRequestDataType string

// List of RecommendationsFilterRequestDataType.
const (
	RECOMMENDATIONSFILTERREQUESTDATATYPE_RECOMMENDATIONS_FILTER RecommendationsFilterRequestDataType = "recommendations_filter"
)

var allowedRecommendationsFilterRequestDataTypeEnumValues = []RecommendationsFilterRequestDataType{
	RECOMMENDATIONSFILTERREQUESTDATATYPE_RECOMMENDATIONS_FILTER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RecommendationsFilterRequestDataType) GetAllowedValues() []RecommendationsFilterRequestDataType {
	return allowedRecommendationsFilterRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RecommendationsFilterRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RecommendationsFilterRequestDataType(value)
	return nil
}

// NewRecommendationsFilterRequestDataTypeFromValue returns a pointer to a valid RecommendationsFilterRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRecommendationsFilterRequestDataTypeFromValue(v string) (*RecommendationsFilterRequestDataType, error) {
	ev := RecommendationsFilterRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RecommendationsFilterRequestDataType: valid values are %v", v, allowedRecommendationsFilterRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RecommendationsFilterRequestDataType) IsValid() bool {
	for _, existing := range allowedRecommendationsFilterRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RecommendationsFilterRequestDataType value.
func (v RecommendationsFilterRequestDataType) Ptr() *RecommendationsFilterRequestDataType {
	return &v
}
