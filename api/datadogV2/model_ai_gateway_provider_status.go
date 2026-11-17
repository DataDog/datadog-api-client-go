// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayProviderStatus Read-only. Whether the provider is working, based on the model sync and live calls. Omitted for
// providers that are never checked, which include providers without an integration account, and
// Azure OpenAI and Gemini providers until live traffic records a status.
type AiGatewayProviderStatus struct {
	// Time of the last observation.
	CheckedAt *time.Time `json:"checked_at,omitempty"`
	// Address of the models endpoint that the gateway called, for custom providers.
	Endpoint *string `json:"endpoint,omitempty"`
	// Display text saying how to fix the provider. Clients render it as-is.
	Fix *string `json:"fix,omitempty"`
	// HTTP status code that the provider returned.
	HttpStatus *int64 `json:"http_status,omitempty"`
	// Display text for the reason. Clients render it as-is.
	Label *string `json:"label,omitempty"`
	// Time of the last observation in which the provider worked.
	LastOkAt *time.Time `json:"last_ok_at,omitempty"`
	// The provider's own error message, with credentials removed. Never set for custom providers.
	Message *string `json:"message,omitempty"`
	// Why the provider is invalid, as a stable code such as `credentials_rejected`. Set only when `state` is `invalid`.
	Reason *string `json:"reason,omitempty"`
	// Whether the provider is working. `checking` means no observation has been recorded yet, not that a
	// check is running.
	State AiGatewayProviderStatusState `json:"state"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayProviderStatus instantiates a new AiGatewayProviderStatus object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayProviderStatus(state AiGatewayProviderStatusState) *AiGatewayProviderStatus {
	this := AiGatewayProviderStatus{}
	this.State = state
	return &this
}

// NewAiGatewayProviderStatusWithDefaults instantiates a new AiGatewayProviderStatus object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayProviderStatusWithDefaults() *AiGatewayProviderStatus {
	this := AiGatewayProviderStatus{}
	return &this
}

// GetCheckedAt returns the CheckedAt field value if set, zero value otherwise.
func (o *AiGatewayProviderStatus) GetCheckedAt() time.Time {
	if o == nil || o.CheckedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CheckedAt
}

// GetCheckedAtOk returns a tuple with the CheckedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderStatus) GetCheckedAtOk() (*time.Time, bool) {
	if o == nil || o.CheckedAt == nil {
		return nil, false
	}
	return o.CheckedAt, true
}

// HasCheckedAt returns a boolean if a field has been set.
func (o *AiGatewayProviderStatus) HasCheckedAt() bool {
	return o != nil && o.CheckedAt != nil
}

// SetCheckedAt gets a reference to the given time.Time and assigns it to the CheckedAt field.
func (o *AiGatewayProviderStatus) SetCheckedAt(v time.Time) {
	o.CheckedAt = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *AiGatewayProviderStatus) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderStatus) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *AiGatewayProviderStatus) HasEndpoint() bool {
	return o != nil && o.Endpoint != nil
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *AiGatewayProviderStatus) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetFix returns the Fix field value if set, zero value otherwise.
func (o *AiGatewayProviderStatus) GetFix() string {
	if o == nil || o.Fix == nil {
		var ret string
		return ret
	}
	return *o.Fix
}

// GetFixOk returns a tuple with the Fix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderStatus) GetFixOk() (*string, bool) {
	if o == nil || o.Fix == nil {
		return nil, false
	}
	return o.Fix, true
}

// HasFix returns a boolean if a field has been set.
func (o *AiGatewayProviderStatus) HasFix() bool {
	return o != nil && o.Fix != nil
}

// SetFix gets a reference to the given string and assigns it to the Fix field.
func (o *AiGatewayProviderStatus) SetFix(v string) {
	o.Fix = &v
}

// GetHttpStatus returns the HttpStatus field value if set, zero value otherwise.
func (o *AiGatewayProviderStatus) GetHttpStatus() int64 {
	if o == nil || o.HttpStatus == nil {
		var ret int64
		return ret
	}
	return *o.HttpStatus
}

// GetHttpStatusOk returns a tuple with the HttpStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderStatus) GetHttpStatusOk() (*int64, bool) {
	if o == nil || o.HttpStatus == nil {
		return nil, false
	}
	return o.HttpStatus, true
}

// HasHttpStatus returns a boolean if a field has been set.
func (o *AiGatewayProviderStatus) HasHttpStatus() bool {
	return o != nil && o.HttpStatus != nil
}

// SetHttpStatus gets a reference to the given int64 and assigns it to the HttpStatus field.
func (o *AiGatewayProviderStatus) SetHttpStatus(v int64) {
	o.HttpStatus = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *AiGatewayProviderStatus) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderStatus) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *AiGatewayProviderStatus) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *AiGatewayProviderStatus) SetLabel(v string) {
	o.Label = &v
}

// GetLastOkAt returns the LastOkAt field value if set, zero value otherwise.
func (o *AiGatewayProviderStatus) GetLastOkAt() time.Time {
	if o == nil || o.LastOkAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastOkAt
}

// GetLastOkAtOk returns a tuple with the LastOkAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderStatus) GetLastOkAtOk() (*time.Time, bool) {
	if o == nil || o.LastOkAt == nil {
		return nil, false
	}
	return o.LastOkAt, true
}

// HasLastOkAt returns a boolean if a field has been set.
func (o *AiGatewayProviderStatus) HasLastOkAt() bool {
	return o != nil && o.LastOkAt != nil
}

// SetLastOkAt gets a reference to the given time.Time and assigns it to the LastOkAt field.
func (o *AiGatewayProviderStatus) SetLastOkAt(v time.Time) {
	o.LastOkAt = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AiGatewayProviderStatus) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderStatus) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AiGatewayProviderStatus) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AiGatewayProviderStatus) SetMessage(v string) {
	o.Message = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AiGatewayProviderStatus) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderStatus) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AiGatewayProviderStatus) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AiGatewayProviderStatus) SetReason(v string) {
	o.Reason = &v
}

// GetState returns the State field value.
func (o *AiGatewayProviderStatus) GetState() AiGatewayProviderStatusState {
	if o == nil {
		var ret AiGatewayProviderStatusState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderStatus) GetStateOk() (*AiGatewayProviderStatusState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *AiGatewayProviderStatus) SetState(v AiGatewayProviderStatusState) {
	o.State = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayProviderStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CheckedAt != nil {
		if o.CheckedAt.Nanosecond() == 0 {
			toSerialize["checked_at"] = o.CheckedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["checked_at"] = o.CheckedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Fix != nil {
		toSerialize["fix"] = o.Fix
	}
	if o.HttpStatus != nil {
		toSerialize["http_status"] = o.HttpStatus
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.LastOkAt != nil {
		if o.LastOkAt.Nanosecond() == 0 {
			toSerialize["last_ok_at"] = o.LastOkAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["last_ok_at"] = o.LastOkAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["state"] = o.State

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayProviderStatus) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CheckedAt  *time.Time                    `json:"checked_at,omitempty"`
		Endpoint   *string                       `json:"endpoint,omitempty"`
		Fix        *string                       `json:"fix,omitempty"`
		HttpStatus *int64                        `json:"http_status,omitempty"`
		Label      *string                       `json:"label,omitempty"`
		LastOkAt   *time.Time                    `json:"last_ok_at,omitempty"`
		Message    *string                       `json:"message,omitempty"`
		Reason     *string                       `json:"reason,omitempty"`
		State      *AiGatewayProviderStatusState `json:"state"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"checked_at", "endpoint", "fix", "http_status", "label", "last_ok_at", "message", "reason", "state"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CheckedAt = all.CheckedAt
	o.Endpoint = all.Endpoint
	o.Fix = all.Fix
	o.HttpStatus = all.HttpStatus
	o.Label = all.Label
	o.LastOkAt = all.LastOkAt
	o.Message = all.Message
	o.Reason = all.Reason
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
