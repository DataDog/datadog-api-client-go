// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RecommendationV2RequestType JSON:API resource type for the SPA v2 recommendation request.
type RecommendationV2RequestType string

// List of RecommendationV2RequestType.
const (
	RECOMMENDATIONV2REQUESTTYPE_RECOMMENDATION_V2_REQUEST RecommendationV2RequestType = "recommendation_v2_request"
)

var allowedRecommendationV2RequestTypeEnumValues = []RecommendationV2RequestType{
	RECOMMENDATIONV2REQUESTTYPE_RECOMMENDATION_V2_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RecommendationV2RequestType) GetAllowedValues() []RecommendationV2RequestType {
	return allowedRecommendationV2RequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RecommendationV2RequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RecommendationV2RequestType(value)
	return nil
}

// NewRecommendationV2RequestTypeFromValue returns a pointer to a valid RecommendationV2RequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRecommendationV2RequestTypeFromValue(v string) (*RecommendationV2RequestType, error) {
	ev := RecommendationV2RequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RecommendationV2RequestType: valid values are %v", v, allowedRecommendationV2RequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RecommendationV2RequestType) IsValid() bool {
	for _, existing := range allowedRecommendationV2RequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RecommendationV2RequestType value.
func (v RecommendationV2RequestType) Ptr() *RecommendationV2RequestType {
	return &v
}
