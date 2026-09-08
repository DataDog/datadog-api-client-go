// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RecommendationsFilterRequestScope Recommendations scope. Defaults to `ccm`; use `experiment` for experimental recommendations or `*` for both.
type RecommendationsFilterRequestScope string

// List of RecommendationsFilterRequestScope.
const (
	RECOMMENDATIONSFILTERREQUESTSCOPE_CCM        RecommendationsFilterRequestScope = "ccm"
	RECOMMENDATIONSFILTERREQUESTSCOPE_EXPERIMENT RecommendationsFilterRequestScope = "experiment"
	RECOMMENDATIONSFILTERREQUESTSCOPE_ALL        RecommendationsFilterRequestScope = "*"
)

var allowedRecommendationsFilterRequestScopeEnumValues = []RecommendationsFilterRequestScope{
	RECOMMENDATIONSFILTERREQUESTSCOPE_CCM,
	RECOMMENDATIONSFILTERREQUESTSCOPE_EXPERIMENT,
	RECOMMENDATIONSFILTERREQUESTSCOPE_ALL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RecommendationsFilterRequestScope) GetAllowedValues() []RecommendationsFilterRequestScope {
	return allowedRecommendationsFilterRequestScopeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RecommendationsFilterRequestScope) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RecommendationsFilterRequestScope(value)
	return nil
}

// NewRecommendationsFilterRequestScopeFromValue returns a pointer to a valid RecommendationsFilterRequestScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRecommendationsFilterRequestScopeFromValue(v string) (*RecommendationsFilterRequestScope, error) {
	ev := RecommendationsFilterRequestScope(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RecommendationsFilterRequestScope: valid values are %v", v, allowedRecommendationsFilterRequestScopeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RecommendationsFilterRequestScope) IsValid() bool {
	for _, existing := range allowedRecommendationsFilterRequestScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RecommendationsFilterRequestScope value.
func (v RecommendationsFilterRequestScope) Ptr() *RecommendationsFilterRequestScope {
	return &v
}
