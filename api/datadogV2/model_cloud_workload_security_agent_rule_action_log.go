// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentRuleActionLog The log action applied when the rule is triggered.
type CloudWorkloadSecurityAgentRuleActionLog struct {
	// The level of the log action.
	Level *string `json:"level,omitempty"`
	// The message of the log action.
	Message *string `json:"message,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudWorkloadSecurityAgentRuleActionLog instantiates a new CloudWorkloadSecurityAgentRuleActionLog object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudWorkloadSecurityAgentRuleActionLog() *CloudWorkloadSecurityAgentRuleActionLog {
	this := CloudWorkloadSecurityAgentRuleActionLog{}
	return &this
}

// NewCloudWorkloadSecurityAgentRuleActionLogWithDefaults instantiates a new CloudWorkloadSecurityAgentRuleActionLog object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudWorkloadSecurityAgentRuleActionLogWithDefaults() *CloudWorkloadSecurityAgentRuleActionLog {
	this := CloudWorkloadSecurityAgentRuleActionLog{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionLog) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionLog) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionLog) HasLevel() bool {
	return o != nil && o.Level != nil
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *CloudWorkloadSecurityAgentRuleActionLog) SetLevel(v string) {
	o.Level = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionLog) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionLog) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionLog) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CloudWorkloadSecurityAgentRuleActionLog) SetMessage(v string) {
	o.Message = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudWorkloadSecurityAgentRuleActionLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudWorkloadSecurityAgentRuleActionLog) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Level   *string `json:"level,omitempty"`
		Message *string `json:"message,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"level", "message"})
	} else {
		return err
	}
	o.Level = all.Level
	o.Message = all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
