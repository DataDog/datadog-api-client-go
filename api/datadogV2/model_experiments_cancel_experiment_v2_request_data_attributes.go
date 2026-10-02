// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCancelExperimentV2RequestDataAttributes Reason to record when canceling the experiment.
type ExperimentsCancelExperimentV2RequestDataAttributes struct {
	// Reason the experiment was canceled. Stored as the decision reason on the experiment conclusion. Must not
	// be blank.
	Reason string `json:"reason"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsCancelExperimentV2RequestDataAttributes instantiates a new ExperimentsCancelExperimentV2RequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCancelExperimentV2RequestDataAttributes(reason string) *ExperimentsCancelExperimentV2RequestDataAttributes {
	this := ExperimentsCancelExperimentV2RequestDataAttributes{}
	this.Reason = reason
	return &this
}

// NewExperimentsCancelExperimentV2RequestDataAttributesWithDefaults instantiates a new ExperimentsCancelExperimentV2RequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCancelExperimentV2RequestDataAttributesWithDefaults() *ExperimentsCancelExperimentV2RequestDataAttributes {
	this := ExperimentsCancelExperimentV2RequestDataAttributes{}
	return &this
}

// GetReason returns the Reason field value.
func (o *ExperimentsCancelExperimentV2RequestDataAttributes) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCancelExperimentV2RequestDataAttributes) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *ExperimentsCancelExperimentV2RequestDataAttributes) SetReason(v string) {
	o.Reason = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCancelExperimentV2RequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["reason"] = o.Reason

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCancelExperimentV2RequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Reason *string `json:"reason"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Reason == nil {
		return fmt.Errorf("required field reason missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"reason"})
	} else {
		return err
	}
	o.Reason = *all.Reason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
