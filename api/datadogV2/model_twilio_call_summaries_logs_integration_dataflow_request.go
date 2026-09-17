// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioCallSummariesLogsIntegrationDataflowRequest Twilio Call Summary resource logs, covering the metadata and performance of the calls made from your Twilio account. Requires Voice Insights Advanced Features to be enabled on the Twilio account; without it this dataflow collects no data.
type TwilioCallSummariesLogsIntegrationDataflowRequest struct {
	// Whether Datadog collects this data. Defaults to `false`; set to `true` to start collection.
	Enabled *bool `json:"enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTwilioCallSummariesLogsIntegrationDataflowRequest instantiates a new TwilioCallSummariesLogsIntegrationDataflowRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTwilioCallSummariesLogsIntegrationDataflowRequest() *TwilioCallSummariesLogsIntegrationDataflowRequest {
	this := TwilioCallSummariesLogsIntegrationDataflowRequest{}
	return &this
}

// NewTwilioCallSummariesLogsIntegrationDataflowRequestWithDefaults instantiates a new TwilioCallSummariesLogsIntegrationDataflowRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTwilioCallSummariesLogsIntegrationDataflowRequestWithDefaults() *TwilioCallSummariesLogsIntegrationDataflowRequest {
	this := TwilioCallSummariesLogsIntegrationDataflowRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TwilioCallSummariesLogsIntegrationDataflowRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioCallSummariesLogsIntegrationDataflowRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TwilioCallSummariesLogsIntegrationDataflowRequest) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TwilioCallSummariesLogsIntegrationDataflowRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TwilioCallSummariesLogsIntegrationDataflowRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TwilioCallSummariesLogsIntegrationDataflowRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled *bool `json:"enabled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.Enabled = all.Enabled

	return nil
}
