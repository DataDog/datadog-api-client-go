// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// InboxRuleResponse Response object which includes an inbox rule.
type InboxRuleResponse struct {
	// Inbox rules are used to prioritize and add relevant vulnerabilities to your Security Inbox.
	// An inbox rule is composed of a rule UUID, a rule type, and the rule attributes. All fields are required.
	Data *InboxRule `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInboxRuleResponse instantiates a new InboxRuleResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInboxRuleResponse() *InboxRuleResponse {
	this := InboxRuleResponse{}
	return &this
}

// NewInboxRuleResponseWithDefaults instantiates a new InboxRuleResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInboxRuleResponseWithDefaults() *InboxRuleResponse {
	this := InboxRuleResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InboxRuleResponse) GetData() InboxRule {
	if o == nil || o.Data == nil {
		var ret InboxRule
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboxRuleResponse) GetDataOk() (*InboxRule, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InboxRuleResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given InboxRule and assigns it to the Data field.
func (o *InboxRuleResponse) SetData(v InboxRule) {
	o.Data = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InboxRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InboxRuleResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *InboxRule `json:"data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
