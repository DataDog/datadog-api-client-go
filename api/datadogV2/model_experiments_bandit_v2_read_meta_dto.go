// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditV2ReadMetaDTO Information about the bandit response.
type ExperimentsBanditV2ReadMetaDTO struct {
	// Things the caller should know about the response. Absent when there are none.
	Warnings []ExperimentsBanditV2WarningDTO `json:"warnings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditV2ReadMetaDTO instantiates a new ExperimentsBanditV2ReadMetaDTO object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditV2ReadMetaDTO() *ExperimentsBanditV2ReadMetaDTO {
	this := ExperimentsBanditV2ReadMetaDTO{}
	return &this
}

// NewExperimentsBanditV2ReadMetaDTOWithDefaults instantiates a new ExperimentsBanditV2ReadMetaDTO object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditV2ReadMetaDTOWithDefaults() *ExperimentsBanditV2ReadMetaDTO {
	this := ExperimentsBanditV2ReadMetaDTO{}
	return &this
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ExperimentsBanditV2ReadMetaDTO) GetWarnings() []ExperimentsBanditV2WarningDTO {
	if o == nil || o.Warnings == nil {
		var ret []ExperimentsBanditV2WarningDTO
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2ReadMetaDTO) GetWarningsOk() (*[]ExperimentsBanditV2WarningDTO, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ExperimentsBanditV2ReadMetaDTO) HasWarnings() bool {
	return o != nil && o.Warnings != nil
}

// SetWarnings gets a reference to the given []ExperimentsBanditV2WarningDTO and assigns it to the Warnings field.
func (o *ExperimentsBanditV2ReadMetaDTO) SetWarnings(v []ExperimentsBanditV2WarningDTO) {
	o.Warnings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditV2ReadMetaDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditV2ReadMetaDTO) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Warnings []ExperimentsBanditV2WarningDTO `json:"warnings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"warnings"})
	} else {
		return err
	}
	o.Warnings = all.Warnings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
