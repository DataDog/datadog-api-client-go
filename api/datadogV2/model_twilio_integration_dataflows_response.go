// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioIntegrationDataflowsResponse Data Datadog collects from Twilio, keyed by dataflow id.
type TwilioIntegrationDataflowsResponse struct {
	// Twilio Alert resource logs, which detail the errors and warnings raised when Twilio makes a webhook request to your server or when your application calls the Twilio REST API.
	TwilioAlertsLogs *TwilioAlertsLogsIntegrationDataflowResponse `json:"twilio-alerts-logs,omitempty"`
	// Twilio Call Summary resource logs, covering the metadata and performance of the calls made from your Twilio account. Requires Voice Insights Advanced Features to be enabled on the Twilio account; without it this dataflow collects no data.
	TwilioCallSummariesLogs *TwilioCallSummariesLogsIntegrationDataflowResponse `json:"twilio-call-summaries-logs,omitempty"`
	// Your Twilio cost data, so that Twilio spend can be broken down and attributed in [Cloud Cost Management](https://docs.datadoghq.com/cloud_cost_management/).
	TwilioCloudCostMetrics *TwilioCloudCostMetricsIntegrationDataflowResponse `json:"twilio-cloud-cost-metrics,omitempty"`
	// Twilio Event resource logs, which record virtually every action taken in your Twilio account, such as provisioning a phone number, changing account security settings, or deleting a recording. Actions are recorded whether they came from the REST API, a user in the Twilio Console, or Twilio itself. [Cloud SIEM](https://docs.datadoghq.com/security/cloud_siem/) analyzes and correlates these logs to detect threats in real time.
	TwilioEventsLogs *TwilioEventsLogsIntegrationDataflowResponse `json:"twilio-events-logs,omitempty"`
	// Twilio Message resource logs for inbound and outbound messages, used to track delivery and troubleshoot message errors. A log is produced when you send a message through the REST API, when Twilio executes a TwiML instruction, and when someone messages one of your Twilio numbers or channel addresses. Message bodies are never collected.
	TwilioMessagesLogs *TwilioMessagesLogsIntegrationDataflowResponse `json:"twilio-messages-logs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTwilioIntegrationDataflowsResponse instantiates a new TwilioIntegrationDataflowsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTwilioIntegrationDataflowsResponse() *TwilioIntegrationDataflowsResponse {
	this := TwilioIntegrationDataflowsResponse{}
	return &this
}

// NewTwilioIntegrationDataflowsResponseWithDefaults instantiates a new TwilioIntegrationDataflowsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTwilioIntegrationDataflowsResponseWithDefaults() *TwilioIntegrationDataflowsResponse {
	this := TwilioIntegrationDataflowsResponse{}
	return &this
}

// GetTwilioAlertsLogs returns the TwilioAlertsLogs field value if set, zero value otherwise.
func (o *TwilioIntegrationDataflowsResponse) GetTwilioAlertsLogs() TwilioAlertsLogsIntegrationDataflowResponse {
	if o == nil || o.TwilioAlertsLogs == nil {
		var ret TwilioAlertsLogsIntegrationDataflowResponse
		return ret
	}
	return *o.TwilioAlertsLogs
}

// GetTwilioAlertsLogsOk returns a tuple with the TwilioAlertsLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationDataflowsResponse) GetTwilioAlertsLogsOk() (*TwilioAlertsLogsIntegrationDataflowResponse, bool) {
	if o == nil || o.TwilioAlertsLogs == nil {
		return nil, false
	}
	return o.TwilioAlertsLogs, true
}

// HasTwilioAlertsLogs returns a boolean if a field has been set.
func (o *TwilioIntegrationDataflowsResponse) HasTwilioAlertsLogs() bool {
	return o != nil && o.TwilioAlertsLogs != nil
}

// SetTwilioAlertsLogs gets a reference to the given TwilioAlertsLogsIntegrationDataflowResponse and assigns it to the TwilioAlertsLogs field.
func (o *TwilioIntegrationDataflowsResponse) SetTwilioAlertsLogs(v TwilioAlertsLogsIntegrationDataflowResponse) {
	o.TwilioAlertsLogs = &v
}

// GetTwilioCallSummariesLogs returns the TwilioCallSummariesLogs field value if set, zero value otherwise.
func (o *TwilioIntegrationDataflowsResponse) GetTwilioCallSummariesLogs() TwilioCallSummariesLogsIntegrationDataflowResponse {
	if o == nil || o.TwilioCallSummariesLogs == nil {
		var ret TwilioCallSummariesLogsIntegrationDataflowResponse
		return ret
	}
	return *o.TwilioCallSummariesLogs
}

// GetTwilioCallSummariesLogsOk returns a tuple with the TwilioCallSummariesLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationDataflowsResponse) GetTwilioCallSummariesLogsOk() (*TwilioCallSummariesLogsIntegrationDataflowResponse, bool) {
	if o == nil || o.TwilioCallSummariesLogs == nil {
		return nil, false
	}
	return o.TwilioCallSummariesLogs, true
}

// HasTwilioCallSummariesLogs returns a boolean if a field has been set.
func (o *TwilioIntegrationDataflowsResponse) HasTwilioCallSummariesLogs() bool {
	return o != nil && o.TwilioCallSummariesLogs != nil
}

// SetTwilioCallSummariesLogs gets a reference to the given TwilioCallSummariesLogsIntegrationDataflowResponse and assigns it to the TwilioCallSummariesLogs field.
func (o *TwilioIntegrationDataflowsResponse) SetTwilioCallSummariesLogs(v TwilioCallSummariesLogsIntegrationDataflowResponse) {
	o.TwilioCallSummariesLogs = &v
}

// GetTwilioCloudCostMetrics returns the TwilioCloudCostMetrics field value if set, zero value otherwise.
func (o *TwilioIntegrationDataflowsResponse) GetTwilioCloudCostMetrics() TwilioCloudCostMetricsIntegrationDataflowResponse {
	if o == nil || o.TwilioCloudCostMetrics == nil {
		var ret TwilioCloudCostMetricsIntegrationDataflowResponse
		return ret
	}
	return *o.TwilioCloudCostMetrics
}

// GetTwilioCloudCostMetricsOk returns a tuple with the TwilioCloudCostMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationDataflowsResponse) GetTwilioCloudCostMetricsOk() (*TwilioCloudCostMetricsIntegrationDataflowResponse, bool) {
	if o == nil || o.TwilioCloudCostMetrics == nil {
		return nil, false
	}
	return o.TwilioCloudCostMetrics, true
}

// HasTwilioCloudCostMetrics returns a boolean if a field has been set.
func (o *TwilioIntegrationDataflowsResponse) HasTwilioCloudCostMetrics() bool {
	return o != nil && o.TwilioCloudCostMetrics != nil
}

// SetTwilioCloudCostMetrics gets a reference to the given TwilioCloudCostMetricsIntegrationDataflowResponse and assigns it to the TwilioCloudCostMetrics field.
func (o *TwilioIntegrationDataflowsResponse) SetTwilioCloudCostMetrics(v TwilioCloudCostMetricsIntegrationDataflowResponse) {
	o.TwilioCloudCostMetrics = &v
}

// GetTwilioEventsLogs returns the TwilioEventsLogs field value if set, zero value otherwise.
func (o *TwilioIntegrationDataflowsResponse) GetTwilioEventsLogs() TwilioEventsLogsIntegrationDataflowResponse {
	if o == nil || o.TwilioEventsLogs == nil {
		var ret TwilioEventsLogsIntegrationDataflowResponse
		return ret
	}
	return *o.TwilioEventsLogs
}

// GetTwilioEventsLogsOk returns a tuple with the TwilioEventsLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationDataflowsResponse) GetTwilioEventsLogsOk() (*TwilioEventsLogsIntegrationDataflowResponse, bool) {
	if o == nil || o.TwilioEventsLogs == nil {
		return nil, false
	}
	return o.TwilioEventsLogs, true
}

// HasTwilioEventsLogs returns a boolean if a field has been set.
func (o *TwilioIntegrationDataflowsResponse) HasTwilioEventsLogs() bool {
	return o != nil && o.TwilioEventsLogs != nil
}

// SetTwilioEventsLogs gets a reference to the given TwilioEventsLogsIntegrationDataflowResponse and assigns it to the TwilioEventsLogs field.
func (o *TwilioIntegrationDataflowsResponse) SetTwilioEventsLogs(v TwilioEventsLogsIntegrationDataflowResponse) {
	o.TwilioEventsLogs = &v
}

// GetTwilioMessagesLogs returns the TwilioMessagesLogs field value if set, zero value otherwise.
func (o *TwilioIntegrationDataflowsResponse) GetTwilioMessagesLogs() TwilioMessagesLogsIntegrationDataflowResponse {
	if o == nil || o.TwilioMessagesLogs == nil {
		var ret TwilioMessagesLogsIntegrationDataflowResponse
		return ret
	}
	return *o.TwilioMessagesLogs
}

// GetTwilioMessagesLogsOk returns a tuple with the TwilioMessagesLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationDataflowsResponse) GetTwilioMessagesLogsOk() (*TwilioMessagesLogsIntegrationDataflowResponse, bool) {
	if o == nil || o.TwilioMessagesLogs == nil {
		return nil, false
	}
	return o.TwilioMessagesLogs, true
}

// HasTwilioMessagesLogs returns a boolean if a field has been set.
func (o *TwilioIntegrationDataflowsResponse) HasTwilioMessagesLogs() bool {
	return o != nil && o.TwilioMessagesLogs != nil
}

// SetTwilioMessagesLogs gets a reference to the given TwilioMessagesLogsIntegrationDataflowResponse and assigns it to the TwilioMessagesLogs field.
func (o *TwilioIntegrationDataflowsResponse) SetTwilioMessagesLogs(v TwilioMessagesLogsIntegrationDataflowResponse) {
	o.TwilioMessagesLogs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TwilioIntegrationDataflowsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TwilioAlertsLogs != nil {
		toSerialize["twilio-alerts-logs"] = o.TwilioAlertsLogs
	}
	if o.TwilioCallSummariesLogs != nil {
		toSerialize["twilio-call-summaries-logs"] = o.TwilioCallSummariesLogs
	}
	if o.TwilioCloudCostMetrics != nil {
		toSerialize["twilio-cloud-cost-metrics"] = o.TwilioCloudCostMetrics
	}
	if o.TwilioEventsLogs != nil {
		toSerialize["twilio-events-logs"] = o.TwilioEventsLogs
	}
	if o.TwilioMessagesLogs != nil {
		toSerialize["twilio-messages-logs"] = o.TwilioMessagesLogs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TwilioIntegrationDataflowsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TwilioAlertsLogs        *TwilioAlertsLogsIntegrationDataflowResponse        `json:"twilio-alerts-logs,omitempty"`
		TwilioCallSummariesLogs *TwilioCallSummariesLogsIntegrationDataflowResponse `json:"twilio-call-summaries-logs,omitempty"`
		TwilioCloudCostMetrics  *TwilioCloudCostMetricsIntegrationDataflowResponse  `json:"twilio-cloud-cost-metrics,omitempty"`
		TwilioEventsLogs        *TwilioEventsLogsIntegrationDataflowResponse        `json:"twilio-events-logs,omitempty"`
		TwilioMessagesLogs      *TwilioMessagesLogsIntegrationDataflowResponse      `json:"twilio-messages-logs,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"twilio-alerts-logs", "twilio-call-summaries-logs", "twilio-cloud-cost-metrics", "twilio-events-logs", "twilio-messages-logs"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.TwilioAlertsLogs != nil && all.TwilioAlertsLogs.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TwilioAlertsLogs = all.TwilioAlertsLogs
	if all.TwilioCallSummariesLogs != nil && all.TwilioCallSummariesLogs.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TwilioCallSummariesLogs = all.TwilioCallSummariesLogs
	if all.TwilioCloudCostMetrics != nil && all.TwilioCloudCostMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TwilioCloudCostMetrics = all.TwilioCloudCostMetrics
	if all.TwilioEventsLogs != nil && all.TwilioEventsLogs.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TwilioEventsLogs = all.TwilioEventsLogs
	if all.TwilioMessagesLogs != nil && all.TwilioMessagesLogs.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TwilioMessagesLogs = all.TwilioMessagesLogs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
