// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationPublishError The reason why the resource could not be published.
type SecurityMonitoringBatchMigrationPublishError struct {
	// The error code returned on rule creation.
	Code *string `json:"code,omitempty"`
	// The rule validation errors returned on rule creation.
	Details []SecurityMonitoringBatchMigrationPublishErrorDetail `json:"details,omitempty"`
	// The error message.
	Message string `json:"message"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringBatchMigrationPublishError instantiates a new SecurityMonitoringBatchMigrationPublishError object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringBatchMigrationPublishError(message string) *SecurityMonitoringBatchMigrationPublishError {
	this := SecurityMonitoringBatchMigrationPublishError{}
	this.Message = message
	return &this
}

// NewSecurityMonitoringBatchMigrationPublishErrorWithDefaults instantiates a new SecurityMonitoringBatchMigrationPublishError object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringBatchMigrationPublishErrorWithDefaults() *SecurityMonitoringBatchMigrationPublishError {
	this := SecurityMonitoringBatchMigrationPublishError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationPublishError) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationPublishError) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationPublishError) HasCode() bool {
	return o != nil && o.Code != nil
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SecurityMonitoringBatchMigrationPublishError) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationPublishError) GetDetails() []SecurityMonitoringBatchMigrationPublishErrorDetail {
	if o == nil || o.Details == nil {
		var ret []SecurityMonitoringBatchMigrationPublishErrorDetail
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationPublishError) GetDetailsOk() (*[]SecurityMonitoringBatchMigrationPublishErrorDetail, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return &o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationPublishError) HasDetails() bool {
	return o != nil && o.Details != nil
}

// SetDetails gets a reference to the given []SecurityMonitoringBatchMigrationPublishErrorDetail and assigns it to the Details field.
func (o *SecurityMonitoringBatchMigrationPublishError) SetDetails(v []SecurityMonitoringBatchMigrationPublishErrorDetail) {
	o.Details = v
}

// GetMessage returns the Message field value.
func (o *SecurityMonitoringBatchMigrationPublishError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationPublishError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *SecurityMonitoringBatchMigrationPublishError) SetMessage(v string) {
	o.Message = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringBatchMigrationPublishError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	toSerialize["message"] = o.Message

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringBatchMigrationPublishError) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Code    *string                                              `json:"code,omitempty"`
		Details []SecurityMonitoringBatchMigrationPublishErrorDetail `json:"details,omitempty"`
		Message *string                                              `json:"message"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"code", "details", "message"})
	} else {
		return err
	}
	o.Code = all.Code
	o.Details = all.Details
	o.Message = *all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
