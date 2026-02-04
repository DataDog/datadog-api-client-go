// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormSubmissionAttributes
type FormSubmissionAttributes struct {
	// The data submitted with the form.
	SubmissionData interface{} `json:"submission_data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormSubmissionAttributes instantiates a new FormSubmissionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormSubmissionAttributes(submissionData interface{}) *FormSubmissionAttributes {
	this := FormSubmissionAttributes{}
	this.SubmissionData = submissionData
	return &this
}

// NewFormSubmissionAttributesWithDefaults instantiates a new FormSubmissionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormSubmissionAttributesWithDefaults() *FormSubmissionAttributes {
	this := FormSubmissionAttributes{}
	return &this
}

// GetSubmissionData returns the SubmissionData field value.
func (o *FormSubmissionAttributes) GetSubmissionData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SubmissionData
}

// GetSubmissionDataOk returns a tuple with the SubmissionData field value
// and a boolean to check if the value has been set.
func (o *FormSubmissionAttributes) GetSubmissionDataOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubmissionData, true
}

// SetSubmissionData sets field value.
func (o *FormSubmissionAttributes) SetSubmissionData(v interface{}) {
	o.SubmissionData = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormSubmissionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["submission_data"] = o.SubmissionData

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormSubmissionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SubmissionData *interface{} `json:"submission_data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SubmissionData == nil {
		return fmt.Errorf("required field submission_data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"submission_data"})
	} else {
		return err
	}
	o.SubmissionData = *all.SubmissionData

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
