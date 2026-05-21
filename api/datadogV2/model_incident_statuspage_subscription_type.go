// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentStatuspageSubscriptionType Statuspage email subscription resource type.
type IncidentStatuspageSubscriptionType string

// List of IncidentStatuspageSubscriptionType.
const (
	INCIDENTSTATUSPAGESUBSCRIPTIONTYPE_STATUSPAGE_EMAIL_SUBSCRIPTION IncidentStatuspageSubscriptionType = "statuspage_email_subscription"
)

var allowedIncidentStatuspageSubscriptionTypeEnumValues = []IncidentStatuspageSubscriptionType{
	INCIDENTSTATUSPAGESUBSCRIPTIONTYPE_STATUSPAGE_EMAIL_SUBSCRIPTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentStatuspageSubscriptionType) GetAllowedValues() []IncidentStatuspageSubscriptionType {
	return allowedIncidentStatuspageSubscriptionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentStatuspageSubscriptionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentStatuspageSubscriptionType(value)
	return nil
}

// NewIncidentStatuspageSubscriptionTypeFromValue returns a pointer to a valid IncidentStatuspageSubscriptionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentStatuspageSubscriptionTypeFromValue(v string) (*IncidentStatuspageSubscriptionType, error) {
	ev := IncidentStatuspageSubscriptionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentStatuspageSubscriptionType: valid values are %v", v, allowedIncidentStatuspageSubscriptionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentStatuspageSubscriptionType) IsValid() bool {
	for _, existing := range allowedIncidentStatuspageSubscriptionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentStatuspageSubscriptionType value.
func (v IncidentStatuspageSubscriptionType) Ptr() *IncidentStatuspageSubscriptionType {
	return &v
}
