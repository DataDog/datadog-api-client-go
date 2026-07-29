// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioDataflowId Identifier of a Twilio dataflow.
type TwilioDataflowId string

// List of TwilioDataflowId.
const (
	TWILIODATAFLOWID_CLOUD_COST_METRICS  TwilioDataflowId = "twilio-cloud-cost-metrics"
	TWILIODATAFLOWID_EVENTS_LOGS         TwilioDataflowId = "twilio-events-logs"
	TWILIODATAFLOWID_MESSAGES_LOGS       TwilioDataflowId = "twilio-messages-logs"
	TWILIODATAFLOWID_ALERTS_LOGS         TwilioDataflowId = "twilio-alerts-logs"
	TWILIODATAFLOWID_CALL_SUMMARIES_LOGS TwilioDataflowId = "twilio-call-summaries-logs"
)

var allowedTwilioDataflowIdEnumValues = []TwilioDataflowId{
	TWILIODATAFLOWID_CLOUD_COST_METRICS,
	TWILIODATAFLOWID_EVENTS_LOGS,
	TWILIODATAFLOWID_MESSAGES_LOGS,
	TWILIODATAFLOWID_ALERTS_LOGS,
	TWILIODATAFLOWID_CALL_SUMMARIES_LOGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TwilioDataflowId) GetAllowedValues() []TwilioDataflowId {
	return allowedTwilioDataflowIdEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TwilioDataflowId) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TwilioDataflowId(value)
	return nil
}

// NewTwilioDataflowIdFromValue returns a pointer to a valid TwilioDataflowId
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTwilioDataflowIdFromValue(v string) (*TwilioDataflowId, error) {
	ev := TwilioDataflowId(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TwilioDataflowId: valid values are %v", v, allowedTwilioDataflowIdEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TwilioDataflowId) IsValid() bool {
	for _, existing := range allowedTwilioDataflowIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TwilioDataflowId value.
func (v TwilioDataflowId) Ptr() *TwilioDataflowId {
	return &v
}
