// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioIntegrationDataflowsRequest Data Datadog collects from Twilio, keyed by dataflow id. Each dataflow turns on a distinct kind of collection: set `enabled` to start or stop it. Defaults listed on each dataflow apply when the account is created; on update, omitted fields keep their current values. Where a dataflow depends on a Twilio feature being enabled, that prerequisite is noted on the dataflow; a dataflow enabled without it is stored but collects no data.
type TwilioIntegrationDataflowsRequest struct {
	// Twilio Alert resource logs, which detail the errors and warnings raised when Twilio makes a webhook request to your server or when your application calls the Twilio REST API. This is the one kind of data the integration collects by default.
	TwilioAlertsLogs *TwilioAlertsLogsIntegrationDataflowRequest `json:"twilio-alerts-logs,omitempty"`
	// Twilio Call Summary resource logs, covering the metadata and performance of the calls made from your Twilio account. Requires Voice Insights Advanced Features to be enabled on the Twilio account; without it this dataflow collects no data.
	TwilioCallSummariesLogs *TwilioCallSummariesLogsIntegrationDataflowRequest `json:"twilio-call-summaries-logs,omitempty"`
	// Your Twilio cost data, so that Twilio spend can be broken down and attributed in [Cloud Cost Management](https://docs.datadoghq.com/cloud_cost_management/). Cost data appears in Cloud Cost Management within 24 hours of enabling this dataflow.
	TwilioCloudCostMetrics *TwilioCloudCostMetricsIntegrationDataflowRequest `json:"twilio-cloud-cost-metrics,omitempty"`
	// Twilio Event resource logs, which record virtually every action taken in your Twilio account, such as provisioning a phone number, changing account security settings, or deleting a recording. Actions are recorded whether they came from the REST API, a user in the Twilio Console, or Twilio itself. [Cloud SIEM](https://docs.datadoghq.com/security/cloud_siem/) analyzes and correlates these logs to detect threats in real time.
	TwilioEventsLogs *TwilioEventsLogsIntegrationDataflowRequest `json:"twilio-events-logs,omitempty"`
	// Twilio Message resource logs for inbound and outbound messages, used to track delivery and troubleshoot message errors. A log is produced when you send a message through the REST API, when Twilio executes a TwiML instruction, and when someone messages one of your Twilio numbers or channel addresses. Message bodies are never collected.
	TwilioMessagesLogs *TwilioMessagesLogsIntegrationDataflowRequest `json:"twilio-messages-logs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTwilioIntegrationDataflowsRequest instantiates a new TwilioIntegrationDataflowsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTwilioIntegrationDataflowsRequest() *TwilioIntegrationDataflowsRequest {
	this := TwilioIntegrationDataflowsRequest{}
	return &this
}

// NewTwilioIntegrationDataflowsRequestWithDefaults instantiates a new TwilioIntegrationDataflowsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTwilioIntegrationDataflowsRequestWithDefaults() *TwilioIntegrationDataflowsRequest {
	this := TwilioIntegrationDataflowsRequest{}
	return &this
}

// GetTwilioAlertsLogs returns the TwilioAlertsLogs field value if set, zero value otherwise.
func (o *TwilioIntegrationDataflowsRequest) GetTwilioAlertsLogs() TwilioAlertsLogsIntegrationDataflowRequest {
	if o == nil || o.TwilioAlertsLogs == nil {
		var ret TwilioAlertsLogsIntegrationDataflowRequest
		return ret
	}
	return *o.TwilioAlertsLogs
}

// GetTwilioAlertsLogsOk returns a tuple with the TwilioAlertsLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationDataflowsRequest) GetTwilioAlertsLogsOk() (*TwilioAlertsLogsIntegrationDataflowRequest, bool) {
	if o == nil || o.TwilioAlertsLogs == nil {
		return nil, false
	}
	return o.TwilioAlertsLogs, true
}

// HasTwilioAlertsLogs returns a boolean if a field has been set.
func (o *TwilioIntegrationDataflowsRequest) HasTwilioAlertsLogs() bool {
	return o != nil && o.TwilioAlertsLogs != nil
}

// SetTwilioAlertsLogs gets a reference to the given TwilioAlertsLogsIntegrationDataflowRequest and assigns it to the TwilioAlertsLogs field.
func (o *TwilioIntegrationDataflowsRequest) SetTwilioAlertsLogs(v TwilioAlertsLogsIntegrationDataflowRequest) {
	o.TwilioAlertsLogs = &v
}

// GetTwilioCallSummariesLogs returns the TwilioCallSummariesLogs field value if set, zero value otherwise.
func (o *TwilioIntegrationDataflowsRequest) GetTwilioCallSummariesLogs() TwilioCallSummariesLogsIntegrationDataflowRequest {
	if o == nil || o.TwilioCallSummariesLogs == nil {
		var ret TwilioCallSummariesLogsIntegrationDataflowRequest
		return ret
	}
	return *o.TwilioCallSummariesLogs
}

// GetTwilioCallSummariesLogsOk returns a tuple with the TwilioCallSummariesLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationDataflowsRequest) GetTwilioCallSummariesLogsOk() (*TwilioCallSummariesLogsIntegrationDataflowRequest, bool) {
	if o == nil || o.TwilioCallSummariesLogs == nil {
		return nil, false
	}
	return o.TwilioCallSummariesLogs, true
}

// HasTwilioCallSummariesLogs returns a boolean if a field has been set.
func (o *TwilioIntegrationDataflowsRequest) HasTwilioCallSummariesLogs() bool {
	return o != nil && o.TwilioCallSummariesLogs != nil
}

// SetTwilioCallSummariesLogs gets a reference to the given TwilioCallSummariesLogsIntegrationDataflowRequest and assigns it to the TwilioCallSummariesLogs field.
func (o *TwilioIntegrationDataflowsRequest) SetTwilioCallSummariesLogs(v TwilioCallSummariesLogsIntegrationDataflowRequest) {
	o.TwilioCallSummariesLogs = &v
}

// GetTwilioCloudCostMetrics returns the TwilioCloudCostMetrics field value if set, zero value otherwise.
func (o *TwilioIntegrationDataflowsRequest) GetTwilioCloudCostMetrics() TwilioCloudCostMetricsIntegrationDataflowRequest {
	if o == nil || o.TwilioCloudCostMetrics == nil {
		var ret TwilioCloudCostMetricsIntegrationDataflowRequest
		return ret
	}
	return *o.TwilioCloudCostMetrics
}

// GetTwilioCloudCostMetricsOk returns a tuple with the TwilioCloudCostMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationDataflowsRequest) GetTwilioCloudCostMetricsOk() (*TwilioCloudCostMetricsIntegrationDataflowRequest, bool) {
	if o == nil || o.TwilioCloudCostMetrics == nil {
		return nil, false
	}
	return o.TwilioCloudCostMetrics, true
}

// HasTwilioCloudCostMetrics returns a boolean if a field has been set.
func (o *TwilioIntegrationDataflowsRequest) HasTwilioCloudCostMetrics() bool {
	return o != nil && o.TwilioCloudCostMetrics != nil
}

// SetTwilioCloudCostMetrics gets a reference to the given TwilioCloudCostMetricsIntegrationDataflowRequest and assigns it to the TwilioCloudCostMetrics field.
func (o *TwilioIntegrationDataflowsRequest) SetTwilioCloudCostMetrics(v TwilioCloudCostMetricsIntegrationDataflowRequest) {
	o.TwilioCloudCostMetrics = &v
}

// GetTwilioEventsLogs returns the TwilioEventsLogs field value if set, zero value otherwise.
func (o *TwilioIntegrationDataflowsRequest) GetTwilioEventsLogs() TwilioEventsLogsIntegrationDataflowRequest {
	if o == nil || o.TwilioEventsLogs == nil {
		var ret TwilioEventsLogsIntegrationDataflowRequest
		return ret
	}
	return *o.TwilioEventsLogs
}

// GetTwilioEventsLogsOk returns a tuple with the TwilioEventsLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationDataflowsRequest) GetTwilioEventsLogsOk() (*TwilioEventsLogsIntegrationDataflowRequest, bool) {
	if o == nil || o.TwilioEventsLogs == nil {
		return nil, false
	}
	return o.TwilioEventsLogs, true
}

// HasTwilioEventsLogs returns a boolean if a field has been set.
func (o *TwilioIntegrationDataflowsRequest) HasTwilioEventsLogs() bool {
	return o != nil && o.TwilioEventsLogs != nil
}

// SetTwilioEventsLogs gets a reference to the given TwilioEventsLogsIntegrationDataflowRequest and assigns it to the TwilioEventsLogs field.
func (o *TwilioIntegrationDataflowsRequest) SetTwilioEventsLogs(v TwilioEventsLogsIntegrationDataflowRequest) {
	o.TwilioEventsLogs = &v
}

// GetTwilioMessagesLogs returns the TwilioMessagesLogs field value if set, zero value otherwise.
func (o *TwilioIntegrationDataflowsRequest) GetTwilioMessagesLogs() TwilioMessagesLogsIntegrationDataflowRequest {
	if o == nil || o.TwilioMessagesLogs == nil {
		var ret TwilioMessagesLogsIntegrationDataflowRequest
		return ret
	}
	return *o.TwilioMessagesLogs
}

// GetTwilioMessagesLogsOk returns a tuple with the TwilioMessagesLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationDataflowsRequest) GetTwilioMessagesLogsOk() (*TwilioMessagesLogsIntegrationDataflowRequest, bool) {
	if o == nil || o.TwilioMessagesLogs == nil {
		return nil, false
	}
	return o.TwilioMessagesLogs, true
}

// HasTwilioMessagesLogs returns a boolean if a field has been set.
func (o *TwilioIntegrationDataflowsRequest) HasTwilioMessagesLogs() bool {
	return o != nil && o.TwilioMessagesLogs != nil
}

// SetTwilioMessagesLogs gets a reference to the given TwilioMessagesLogsIntegrationDataflowRequest and assigns it to the TwilioMessagesLogs field.
func (o *TwilioIntegrationDataflowsRequest) SetTwilioMessagesLogs(v TwilioMessagesLogsIntegrationDataflowRequest) {
	o.TwilioMessagesLogs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TwilioIntegrationDataflowsRequest) MarshalJSON() ([]byte, error) {
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
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TwilioIntegrationDataflowsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TwilioAlertsLogs        *TwilioAlertsLogsIntegrationDataflowRequest        `json:"twilio-alerts-logs,omitempty"`
		TwilioCallSummariesLogs *TwilioCallSummariesLogsIntegrationDataflowRequest `json:"twilio-call-summaries-logs,omitempty"`
		TwilioCloudCostMetrics  *TwilioCloudCostMetricsIntegrationDataflowRequest  `json:"twilio-cloud-cost-metrics,omitempty"`
		TwilioEventsLogs        *TwilioEventsLogsIntegrationDataflowRequest        `json:"twilio-events-logs,omitempty"`
		TwilioMessagesLogs      *TwilioMessagesLogsIntegrationDataflowRequest      `json:"twilio-messages-logs,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
