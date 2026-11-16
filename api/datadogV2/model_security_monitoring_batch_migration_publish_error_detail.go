// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationPublishErrorDetail A rule validation error.
type SecurityMonitoringBatchMigrationPublishErrorDetail struct {
	// The error code.
	Code *string `json:"code,omitempty"`
	// The error message.
	Message *string `json:"message,omitempty"`
	// The rule field the error applies to.
	Target *string `json:"target,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringBatchMigrationPublishErrorDetail instantiates a new SecurityMonitoringBatchMigrationPublishErrorDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringBatchMigrationPublishErrorDetail() *SecurityMonitoringBatchMigrationPublishErrorDetail {
	this := SecurityMonitoringBatchMigrationPublishErrorDetail{}
	return &this
}

// NewSecurityMonitoringBatchMigrationPublishErrorDetailWithDefaults instantiates a new SecurityMonitoringBatchMigrationPublishErrorDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringBatchMigrationPublishErrorDetailWithDefaults() *SecurityMonitoringBatchMigrationPublishErrorDetail {
	this := SecurityMonitoringBatchMigrationPublishErrorDetail{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationPublishErrorDetail) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationPublishErrorDetail) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationPublishErrorDetail) HasCode() bool {
	return o != nil && o.Code != nil
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SecurityMonitoringBatchMigrationPublishErrorDetail) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationPublishErrorDetail) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationPublishErrorDetail) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationPublishErrorDetail) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SecurityMonitoringBatchMigrationPublishErrorDetail) SetMessage(v string) {
	o.Message = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationPublishErrorDetail) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationPublishErrorDetail) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationPublishErrorDetail) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *SecurityMonitoringBatchMigrationPublishErrorDetail) SetTarget(v string) {
	o.Target = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringBatchMigrationPublishErrorDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringBatchMigrationPublishErrorDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Code    *string `json:"code,omitempty"`
		Message *string `json:"message,omitempty"`
		Target  *string `json:"target,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"code", "message", "target"})
	} else {
		return err
	}
	o.Code = all.Code
	o.Message = all.Message
	o.Target = all.Target

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
