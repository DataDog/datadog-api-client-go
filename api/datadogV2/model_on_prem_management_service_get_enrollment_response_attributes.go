// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnPremManagementServiceGetEnrollmentResponseAttributes Attributes for the enrollment status.
type OnPremManagementServiceGetEnrollmentResponseAttributes struct {
	// The reason for enrollment failure, if applicable.
	FailureReason string `json:"failure_reason"`
	// The runner identifier, if enrollment was successful.
	RunnerId string `json:"runner_id"`
	// The status of the enrollment.
	Status OnPremManagementServiceGetEnrollmentResponseAttributesStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnPremManagementServiceGetEnrollmentResponseAttributes instantiates a new OnPremManagementServiceGetEnrollmentResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnPremManagementServiceGetEnrollmentResponseAttributes(failureReason string, runnerId string, status OnPremManagementServiceGetEnrollmentResponseAttributesStatus) *OnPremManagementServiceGetEnrollmentResponseAttributes {
	this := OnPremManagementServiceGetEnrollmentResponseAttributes{}
	this.FailureReason = failureReason
	this.RunnerId = runnerId
	this.Status = status
	return &this
}

// NewOnPremManagementServiceGetEnrollmentResponseAttributesWithDefaults instantiates a new OnPremManagementServiceGetEnrollmentResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnPremManagementServiceGetEnrollmentResponseAttributesWithDefaults() *OnPremManagementServiceGetEnrollmentResponseAttributes {
	this := OnPremManagementServiceGetEnrollmentResponseAttributes{}
	return &this
}

// GetFailureReason returns the FailureReason field value.
func (o *OnPremManagementServiceGetEnrollmentResponseAttributes) GetFailureReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value
// and a boolean to check if the value has been set.
func (o *OnPremManagementServiceGetEnrollmentResponseAttributes) GetFailureReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailureReason, true
}

// SetFailureReason sets field value.
func (o *OnPremManagementServiceGetEnrollmentResponseAttributes) SetFailureReason(v string) {
	o.FailureReason = v
}

// GetRunnerId returns the RunnerId field value.
func (o *OnPremManagementServiceGetEnrollmentResponseAttributes) GetRunnerId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RunnerId
}

// GetRunnerIdOk returns a tuple with the RunnerId field value
// and a boolean to check if the value has been set.
func (o *OnPremManagementServiceGetEnrollmentResponseAttributes) GetRunnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunnerId, true
}

// SetRunnerId sets field value.
func (o *OnPremManagementServiceGetEnrollmentResponseAttributes) SetRunnerId(v string) {
	o.RunnerId = v
}

// GetStatus returns the Status field value.
func (o *OnPremManagementServiceGetEnrollmentResponseAttributes) GetStatus() OnPremManagementServiceGetEnrollmentResponseAttributesStatus {
	if o == nil {
		var ret OnPremManagementServiceGetEnrollmentResponseAttributesStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OnPremManagementServiceGetEnrollmentResponseAttributes) GetStatusOk() (*OnPremManagementServiceGetEnrollmentResponseAttributesStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *OnPremManagementServiceGetEnrollmentResponseAttributes) SetStatus(v OnPremManagementServiceGetEnrollmentResponseAttributesStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnPremManagementServiceGetEnrollmentResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["failure_reason"] = o.FailureReason
	toSerialize["runner_id"] = o.RunnerId
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnPremManagementServiceGetEnrollmentResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FailureReason *string                                                       `json:"failure_reason"`
		RunnerId      *string                                                       `json:"runner_id"`
		Status        *OnPremManagementServiceGetEnrollmentResponseAttributesStatus `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FailureReason == nil {
		return fmt.Errorf("required field failure_reason missing")
	}
	if all.RunnerId == nil {
		return fmt.Errorf("required field runner_id missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"failure_reason", "runner_id", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.FailureReason = *all.FailureReason
	o.RunnerId = *all.RunnerId
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
