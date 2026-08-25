// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioCallSummariesLogsIntegrationDataflowResponse The Twilio call summaries logs dataflow.
type TwilioCallSummariesLogsIntegrationDataflowResponse struct {
	// Whether the Twilio dataflow is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Read-only collection status of a dataflow.
	Status *IntegrationAccountDataflowStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTwilioCallSummariesLogsIntegrationDataflowResponse instantiates a new TwilioCallSummariesLogsIntegrationDataflowResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTwilioCallSummariesLogsIntegrationDataflowResponse() *TwilioCallSummariesLogsIntegrationDataflowResponse {
	this := TwilioCallSummariesLogsIntegrationDataflowResponse{}
	return &this
}

// NewTwilioCallSummariesLogsIntegrationDataflowResponseWithDefaults instantiates a new TwilioCallSummariesLogsIntegrationDataflowResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTwilioCallSummariesLogsIntegrationDataflowResponseWithDefaults() *TwilioCallSummariesLogsIntegrationDataflowResponse {
	this := TwilioCallSummariesLogsIntegrationDataflowResponse{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TwilioCallSummariesLogsIntegrationDataflowResponse) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioCallSummariesLogsIntegrationDataflowResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TwilioCallSummariesLogsIntegrationDataflowResponse) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TwilioCallSummariesLogsIntegrationDataflowResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TwilioCallSummariesLogsIntegrationDataflowResponse) GetStatus() IntegrationAccountDataflowStatus {
	if o == nil || o.Status == nil {
		var ret IntegrationAccountDataflowStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioCallSummariesLogsIntegrationDataflowResponse) GetStatusOk() (*IntegrationAccountDataflowStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TwilioCallSummariesLogsIntegrationDataflowResponse) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given IntegrationAccountDataflowStatus and assigns it to the Status field.
func (o *TwilioCallSummariesLogsIntegrationDataflowResponse) SetStatus(v IntegrationAccountDataflowStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TwilioCallSummariesLogsIntegrationDataflowResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TwilioCallSummariesLogsIntegrationDataflowResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled *bool                             `json:"enabled,omitempty"`
		Status  *IntegrationAccountDataflowStatus `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	if all.Status != nil && all.Status.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
