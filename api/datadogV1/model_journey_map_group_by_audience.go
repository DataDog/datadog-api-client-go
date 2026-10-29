// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JourneyMapGroupByAudience Audience the group by facet is read from.
type JourneyMapGroupByAudience string

// List of JourneyMapGroupByAudience.
const (
	JOURNEYMAPGROUPBYAUDIENCE_USERS    JourneyMapGroupByAudience = "product_analytics_audience_filters.users"
	JOURNEYMAPGROUPBYAUDIENCE_ACCOUNTS JourneyMapGroupByAudience = "product_analytics_audience_filters.accounts"
)

var allowedJourneyMapGroupByAudienceEnumValues = []JourneyMapGroupByAudience{
	JOURNEYMAPGROUPBYAUDIENCE_USERS,
	JOURNEYMAPGROUPBYAUDIENCE_ACCOUNTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *JourneyMapGroupByAudience) GetAllowedValues() []JourneyMapGroupByAudience {
	return allowedJourneyMapGroupByAudienceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *JourneyMapGroupByAudience) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = JourneyMapGroupByAudience(value)
	return nil
}

// NewJourneyMapGroupByAudienceFromValue returns a pointer to a valid JourneyMapGroupByAudience
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewJourneyMapGroupByAudienceFromValue(v string) (*JourneyMapGroupByAudience, error) {
	ev := JourneyMapGroupByAudience(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for JourneyMapGroupByAudience: valid values are %v", v, allowedJourneyMapGroupByAudienceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v JourneyMapGroupByAudience) IsValid() bool {
	for _, existing := range allowedJourneyMapGroupByAudienceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JourneyMapGroupByAudience value.
func (v JourneyMapGroupByAudience) Ptr() *JourneyMapGroupByAudience {
	return &v
}
